################################################################################
# 変数
################################################################################
DYNAMO_ENDPOINT_URL := http://localhost:8000
DYNAMO_TABLE_NAME   := ResourceTable
DYNAMO_PK_NAME      := PK
DYNAMO_SK_NAME      := SK

################################################################################
# タスク
################################################################################
.PHONY: run
run:
	go get -u -t ./...
	go run main.go

.PHONY: create-table
create-table:
	AWS_PAGER='' aws dynamodb create-table \
		--table-name $(DYNAMO_TABLE_NAME) \
		--attribute-definitions AttributeName=$(DYNAMO_PK_NAME),AttributeType=S AttributeName=$(DYNAMO_SK_NAME),AttributeType=S \
		--key-schema AttributeName=$(DYNAMO_PK_NAME),KeyType=HASH AttributeName=$(DYNAMO_SK_NAME),KeyType=RANGE \
		--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
		--endpoint-url $(DYNAMO_ENDPOINT_URL) \
		| jq '.'

.PHONY: delete-table
delete-table:
	AWS_PAGER='' aws dynamodb delete-table --table-name $(DYNAMO_TABLE_NAME) --endpoint-url $(DYNAMO_ENDPOINT_URL) > /dev/null 2>&1 \
	|| echo '$(DYNAMO_TABLE_NAME) がないので消せません'
