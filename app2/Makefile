################################################################################
# 変数
################################################################################
ifdef DYNAMO_LOCAL_ENDPOINT
  DB_LOCAL_ENDPOINT := ${DYNAMO_LOCAL_ENDPOINT}
else
  DB_LOCAL_ENDPOINT := http://localhost:8000
endif

.PHONY: setup
setup:
	go get -u -t ./...

.PHONY: test
test:
	DYNAMO_LOCAL_ENDPOINT=$(DB_LOCAL_ENDPOINT) DISABLE_ENV_DECRYPT=1 go test ./adapter/ -v
#	DYNAMO_LOCAL_ENDPOINT=$(DB_LOCAL_ENDPOINT) DISABLE_ENV_DECRYPT=1 go test ./adapter/controller/ -v
