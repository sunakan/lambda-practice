################################################################################
# 変数
################################################################################
ifdef DYNAMO_LOCAL_ENDPOINT
  DB_LOCAL_ENDPOINT := ${DYNAMO_LOCAL_ENDPOINT}
else
  DB_LOCAL_ENDPOINT := http://localhost:8000
endif

################################################################################
# cgoを使わないGoのクロスコンパイル時に -installsuffix cgo が不要になってた
# https://christina04.hatenablog.com/entry/installsuffix-cgo-is-no-longer-required
# 参考にしたコードではinstallsuffix cgoがついていたのでそのままとする
# 後に勉強して不要と判断したら消す
################################################################################
.PHONY: build-handlers
build-handlers: ## handlers以下のmain.goをコンパイル
	go get -u -t ./...
	find adapter/handlers/ -name main.go -type f \
		| xargs dirname \
		| xargs -I {path} bash -c 'cd {path} && CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo -o main . && pwd && ls'

.PHONY: wait-dynamodb-local
wait-dynamodb-local:
	sh scripts/wait-dynamodb-local.sh $(DB_LOCAL_ENDPOINT)

.PHONY: test
test:
#	DYNAMO_LOCAL_ENDPOINT=$(DB_LOCAL_ENDPOINT) DISABLE_ENV_DECRYPT=1 go test ./adapter/ -v
	DYNAMO_LOCAL_ENDPOINT=$(DB_LOCAL_ENDPOINT) DISABLE_ENV_DECRYPT=1 go test ./adapter/controller/ -v
