include makefiles/gitignore.mk
include makefiles/help.mk

################################################################################
# 変数
################################################################################

################################################################################
# マクロ
################################################################################

################################################################################
# タスク
################################################################################
.PHONY: up
up: ## docker-compose
	mkdir -p go-pkg
	docker-compose up
.PHONY: down
down: ## docker-compose
	docker-compose down

.PHONY: build
build: ## docker-compose
	docker-compose build

.PHONY: bash
bash: ## docker-compose
	docker-compose exec go-app bash

.PHONY: chown
chown:
	sudo chown -R $(USER):$(USER) ./


.PHONY: deploy-docs
deploy-docs: ## ドキュメントをデプロイする
	git subtree push --prefix docs/html/ origin gh-pages
