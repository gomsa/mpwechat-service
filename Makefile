dev:
	make build && make run

build:
	protoc -I . --go_out=plugins=micro:. proto/mpwechat/mpwechat.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t mpwechat-service .

run:
	docker run --net="host" -p 12725 \
                    -e MICRO_REGISTRY=mdns  \
                    -e SERVICE_NAME="gomsa.mpwechat"  \
                    -e AUTH_SERVICE="gomsa.auth"  \
                    -e WECHAT_APP_ID="wx15550c1a89d982c8"  \
                    -e WECHAT_APP_SECRET="f9c11f183a5beb592ccd801298ff5533"  \
					-e DB_NAME="npmwechat_service"  \
                    -e DB_HOST="127.0.0.1"  \
                    -e DB_PORT="3306"  \
                    -e DB_USER="root"  \
                    -e DB_PASSWORD="123456"  \
                    mpwechat-service
