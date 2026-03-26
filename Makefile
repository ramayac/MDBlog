HOST      ?= localhost
PORT      ?= 8080
PHP       := php
TAG       ?= latest
REGISTRY  ?= ghcr.io/ramayac/mdblog

.DEFAULT_GOAL := help

.PHONY: help serve lint new-post version clear-cache utf8-fix docker-build docker-run docker-run-release docker-stop docker-push docker-pull

help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

serve: ## Start PHP built-in dev server (HOST=localhost PORT=8080)
	@echo "Starting dev server at http://$(HOST):$(PORT)"
	$(PHP) -S $(HOST):$(PORT)

lint: ## Check all PHP files for syntax errors
	@errors=0; \
	for f in $$(find . -name '*.php' | sort); do \
		$(PHP) -l "$$f" || errors=$$((errors+1)); \
	done; \
	[ $$errors -eq 0 ] && echo "All PHP files OK." || { echo "$$errors file(s) failed."; exit 1; }

clear-cache: ## Delete all cached .json files from the cache/ folder
	@find cache/ -name '*.json' -type f -delete
	@echo "Cache cleared."

version: ## Bake current git commit+tag into version.php (run before FTP upload or Docker build)
	$(eval _COMMIT  := $(shell git log -1 --format="%h" 2>/dev/null))
	$(eval _DATE    := $(shell git log -1 --format="%ad" --date=short 2>/dev/null))
	$(eval _TAG     := $(shell git describe --tags --abbrev=0 2>/dev/null || true))
	$(eval _VERSION := $(if $(_TAG),$(_TAG),$(_COMMIT)))
	@$(PHP) -r 'file_put_contents("version.php", "<?php return [\"commit\"=>\"$(_COMMIT)\",\"date\"=>\"$(_DATE)\",\"version\"=>\"$(_VERSION)\"];\n");'
	@echo "Written: version.php  (commit=$(_COMMIT), version=$(_VERSION), date=$(_DATE))"

new-post: ## Create a new post template: make new-post TITLE="my post title" [CATEGORY=slug] [TAGS="tag1, tag2"]
	$(eval DATE    := $(shell date +%Y-%m-%d))
	$(eval SLUG    := $(shell echo "$(TITLE)" | tr '[:upper:]' '[:lower:]' | tr ' ' '-' | tr -cd '[:alnum:]-'))
	$(eval DIR     := $(if $(CATEGORY),posts/$(CATEGORY),posts))
	$(eval FILE    := $(DIR)/$(DATE)-$(SLUG).md)
	$(eval AUTHOR  := $(shell $(PHP) -r '$$c=include("config.php"); echo $$c["author_name"];'))
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

docker-build: ## Build the Docker image (bakes version.php first)
	$(MAKE) version
	docker build -t mdblog:latest .

docker-run: ## Start blog via Docker Compose at http://localhost:8080
	docker compose up

docker-stop: ## Stop and remove Docker Compose containers
	docker compose down

docker-push: ## Push image to registry (REGISTRY=ghcr.io/ramayac/mdblog)
	docker tag mdblog:latest $(REGISTRY):latest
	docker push $(REGISTRY):latest

docker-pull: ## Pull a release image from registry: make docker-pull [TAG=1.2.3]
	docker pull $(REGISTRY):$(TAG)
	docker tag $(REGISTRY):$(TAG) mdblog:latest
	@echo "Pulled $(REGISTRY):$(TAG) and tagged as mdblog:latest — run with: make docker-run-release"

docker-run-release: ## Run the pulled release image without rebuilding (use after docker-pull)
	docker compose up --no-build

utf8-fix: ## Re-encode any non-UTF-8 .md files in posts/ to UTF-8 (fixes Bref JSON encoding errors)
	@echo "Scanning posts/ for non-UTF-8 encoded Markdown files..."
	@fixed=0; \
	for f in $$(find posts/ -name '*.md' | sort); do \
		if ! $(PHP) -r 'exit(mb_check_encoding(file_get_contents("'"$$f"'"), "UTF-8") ? 0 : 1);' 2>/dev/null; then \
			$(PHP) -r 'file_put_contents("'"$$f"'", mb_convert_encoding(file_get_contents("'"$$f"'"), "UTF-8", "auto"));'; \
			echo "  Converted: $$f"; \
			fixed=$$((fixed+1)); \
		fi; \
	done; \
	if [ $$fixed -eq 0 ]; then echo "All files are already valid UTF-8."; \
	else echo "$$fixed file(s) converted to UTF-8."; fi
