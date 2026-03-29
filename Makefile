HOST      ?= localhost
PORT      ?= 8080
TAG       ?= latest
REGISTRY  ?= ghcr.io/ramayac/mdblog

# Version info injected into binaries via -ldflags
COMMIT  := $(shell git log -1 --format="%h" 2>/dev/null || echo unknown)
DATE    := $(shell git log -1 --format="%ad" --date=short 2>/dev/null || echo unknown)
_TAG    := $(shell git describe --tags --abbrev=0 2>/dev/null || true)
VERSION := $(if $(_TAG),$(_TAG),$(COMMIT))
LDFLAGS := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)

.DEFAULT_GOAL := help

.PHONY: help serve build build-embed build-index lint test new-post render \
        docker-build docker-build-embed docker-run docker-run-release \
        docker-stop docker-push docker-pull

help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# ── Development ───────────────────────────────────────────────────────────────

serve: ## Start local HTTP server (HOST=localhost PORT=8080)
	@echo "Starting dev server at http://$(HOST):$(PORT)"
	PORT=$(PORT) go run -ldflags "$(LDFLAGS)" ./cmd/mdblog serve

build: ## Compile production binaries to bin/
	@mkdir -p bin
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o bin/mdblog   ./cmd/mdblog
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o bin/lambda   ./cmd/lambda
	@echo "Built: bin/mdblog  bin/lambda"

build-embed: ## Compile embed-variant Lambda binary to bin/lambda-embed
	@mkdir -p bin
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o bin/lambda-embed ./cmd/lambda-embed
	@echo "Built: bin/lambda-embed (templates + assets embedded)"

build-index: ## Generate post metadata index (writes posts/posts.index.json)
	@echo "Building post metadata index..."
	go run ./cmd/mdblog build-index

lint: ## Run go vet on all packages
	go vet ./...

test: build-index ## Run the Go test suite
	go test ./...

# Allow extra arguments to `make render`
ifeq (render,$(firstword $(MAKECMDGOALS)))
  RENDER_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RENDER_ARGS):;@:)
endif

render: ## Render a post to HTML: make render random | make render [category] random | make render filename.md
	go run ./cmd/mdblog render $(RENDER_ARGS)

new-post: ## Scaffold a new post: make new-post TITLE="title" [CATEGORY=slug] [TAGS="tag1, tag2"]
	$(eval DATE    := $(shell date +%Y-%m-%d))
	$(eval SLUG    := $(shell echo "$(TITLE)" | tr '[:upper:]' '[:lower:]' | tr ' ' '-' | tr -cd '[:alnum:]-'))
	$(eval DIR     := $(if $(CATEGORY),posts/$(CATEGORY),posts))
	$(eval FILE    := $(DIR)/$(DATE)-$(SLUG).md)
	$(eval AUTHOR  := $(shell grep -oP 'author_name\s*=\s*"\K[^"]+' config.toml 2>/dev/null || echo "Author"))
	@if [ -z "$(TITLE)" ]; then \
		echo "Usage: make new-post TITLE=\"my post title\" [CATEGORY=slug] [TAGS=\"tag1, tag2\"]"; exit 1; \
	fi
	@if [ -n "$(CATEGORY)" ] && [ ! -d "posts/$(CATEGORY)" ]; then \
		echo "Category folder not found: posts/$(CATEGORY)"; exit 1; \
	fi
	@if [ -f "$(FILE)" ]; then \
		echo "File already exists: $(FILE)"; exit 1; \
	fi
	@printf -- '---\ntitle: $(TITLE)\ndate: $(DATE)\nauthor: $(AUTHOR)\ntags: $(TAGS)\ndescription: \n---\n\n# $(TITLE)\n' > "$(FILE)"
	@echo "Created: $(FILE)"

# ── Docker ────────────────────────────────────────────────────────────────────

docker-build: ## Build the production Docker image (FROM scratch, Lambda-ready)
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg DATE=$(DATE) \
		-t mdblog:latest .

docker-build-embed: ## Build the embed-variant Docker image (templates+assets inside binary)
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg DATE=$(DATE) \
		-f Dockerfile.embed \
		-t mdblog-embed:latest .

docker-run: ## Build and start blog via Docker Compose at http://localhost:8080
	docker compose up --build

docker-stop: ## Stop and remove Docker Compose containers
	docker compose down

docker-push: ## Push image to registry (REGISTRY=ghcr.io/ramayac/mdblog TAG=latest)
	docker tag mdblog:latest $(REGISTRY):$(TAG)
	docker push $(REGISTRY):$(TAG)

docker-pull: ## Pull a release image from registry: make docker-pull [TAG=1.2.3]
	docker pull $(REGISTRY):$(TAG)
	docker tag $(REGISTRY):$(TAG) mdblog:latest
	@echo "Pulled $(REGISTRY):$(TAG) → mdblog:latest — run with: make docker-run-release"

docker-run-release: ## Run the pulled release image without rebuilding (use after docker-pull)
	docker compose up --no-build
