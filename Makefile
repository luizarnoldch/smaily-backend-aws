TEMPLATE_FILE := templates/main.yml
STACK_NAME := crud-arnold-test

init:
	go mod init main
update:
	go mod tidy
mock:
	mockery --all --output ./tests/mocks
build:
	./scripts/build.sh
test:
	go test ./tests/...
f_test:
	./scripts/func_test.sh
deploy:
	sam deploy --template-file $(TEMPLATE_FILE) --stack-name $(STACK_NAME) --capabilities CAPABILITY_NAMED_IAM --resolve-s3
destroy:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
e2e:
	make destroy
	sleep 5
	make test
	make build
	make deploy
	sleep 3
	make f_test