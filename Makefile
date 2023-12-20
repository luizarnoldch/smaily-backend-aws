TEMPLATE_FILE := deployments/cloudformation/main.yml
STACK_NAME := crud-arnold-test

DYNAMO-LOCAL := deployments/docker/dynamo-local.yml

init:
	go mod init main
update:
	go mod tidy
mock:
	mockery --all --output ./test/mocks
build:
	./scripts/build.sh
unit:
	go test ./test/local/...
integration:
	go test ./tests/integration/...
f_test:
	./scripts/func_test.sh
deploy:
	sam deploy --template-file $(TEMPLATE_FILE) --stack-name $(STACK_NAME) --capabilities CAPABILITY_NAMED_IAM --resolve-s3
destroy:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
dynamo-local:
	docker-compose -f ${DYNAMO-LOCAL} up
e2e:
	make destroy
	sleep 5
	make test
	make build
	make deploy
	sleep 3
	make f_test
pull:
	git pull
	make build
test:
	make unit
	make deploy
	sleep 3
	make integration
ab:
	make pull
	make test
