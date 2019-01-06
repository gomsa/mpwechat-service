dev:
	make build && make run

build:
	protoc -I . --go_out=plugins=micro:. proto/wechat/mp.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t mpwechat-service .

run:
	docker run --net="host" -p 12725 \
                    -e MICRO_REGISTRY=mdns  \
                    -e WECHAT_APP_ID="wx15550c1a89d982c8"  \
                    -e WECHAT_APP_SECRET="f9c11f183a5beb592ccd801298ff5533"  \
                    mpwechat-service
