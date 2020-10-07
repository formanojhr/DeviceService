.PHONY: build

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-d -s -w" -o .aws-sam/build/devicefunction/deviceget endpoint/rest/handlers/get/main.go
	env GOOS=linux go build -ldflags="-d -s -w" -o .aws-sam/build/devicefunction/deviceregister endpoint/rest/handlers/register/main.go
	chmod 777 .aws-sam/build/devicefunction/deviceregister
	chmod 777 .aws-sam/build/devicefunction/deviceget
	#cp template.yaml .aws-sam/build/devicefunction
	#zip -j .aws-sam/build/devicefunction/deviceservice.zip .aws-sam/build/devicefunction/device
	#sam build

#Deloys local AGW, service
.PHONY: api
api: build
	sam local start-api

# deploy using serverless command
deploy: build
	sls deploy --verbose
