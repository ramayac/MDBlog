HOST     ?= localhost
PORT     ?= 8080
PHP      := php

.DEFAULT_GOAL := help

.PHONY: help serve new-post

help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

serve: ## Start PHP built-in dev server (HOST=localhost PORT=8080)
	@echo "Starting dev server at http://$(HOST):$(PORT)"
	$(PHP) -S $(HOST):$(PORT)

new-post: ## Create a new post template: make new-post TITLE="my post title" TAGS="tag1, tag2"
	$(eval DATE    := $(shell date +%Y-%m-%d))
	$(eval SLUG    := $(shell echo "$(TITLE)" | tr '[:upper:]' '[:lower:]' | tr ' ' '-' | tr -cd '[:alnum:]-'))
	$(eval FILE    := posts/$(DATE)-$(SLUG).md)
	$(eval AUTHOR  := $(shell $(PHP) -r "$$c=include('config.php'); echo $$c['author_name'];"))
	@if [ -z "$(TITLE)" ]; then \
		echo "Usage: make new-post TITLE=\"my post title\" TAGS=\"tag1, tag2\""; exit 1; \
	fi
	@if [ -f "$(FILE)" ]; then \
		echo "File already exists: $(FILE)"; exit 1; \
	fi
	@printf -- '---\ntitle: $(TITLE)\ndate: $(DATE)\nauthor: $(AUTHOR)\ntags: $(TAGS)\ndescription: \n---\n\n# $(TITLE)\n' > "$(FILE)"
	@echo "Created: $(FILE)"
