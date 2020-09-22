include makefiles/gitignore.mk
include makefiles/rq.mk
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
	docker-compose up
.PHONY: down
down: ## docker-compose
	docker-compose down

#.PHONY: build
#build: ## docker-compose
#	docker-compose build
#.PHONY: bash
#bash: ## docker-compose
#	docker-compose run --service-ports go-app

.PHONY: deploy-docs
deploy-docs: ## ドキュメントをデプロイする
	git subtree push --prefix docs/html/ origin gh-pages
