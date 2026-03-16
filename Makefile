HOST     ?= localhost
PORT     ?= 8080
PHP      := php

.DEFAULT_GOAL := help

.PHONY: help serve new-post version clear-cache

help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

serve: ## Start PHP built-in dev server (HOST=localhost PORT=8080)
	@echo "Starting dev server at http://$(HOST):$(PORT)"
	$(PHP) -S $(HOST):$(PORT)

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
