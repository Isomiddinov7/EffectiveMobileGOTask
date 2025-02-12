CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

copy-proto-module:
	rm -rf "${CURRENT_DIR}/protos"
	rsync -rv --exclude=.git "${CURRENT_DIR}/tg_protos"/* "${CURRENT_DIR}/protos"

gen-proto-module:
	./scripts/gen_proto.sh "${CURRENT_DIR}"

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o "${CURRENT_DIR}/bin/${APP}" "${APP_CMD_DIR}/main.go"

swag-init:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://bahodir:1100@0.0.0.0:5432/songs?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://bahodir:1100@0.0.0.0:5432/songs?sslmode=disable' down