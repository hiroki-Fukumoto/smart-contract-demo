setup:
	docker network create geth-dapp-demo-network
	docker-compose -f ./docker-compose.yml build
	cp ./src/.env.example ./src/.env

.PHONY: build
build:
	docker-compose -f ./docker-compose.yml build

up:
	docker-compose -f ./docker-compose.yml up

up-d:
	docker-compose -f ./docker-compose.yml up -d

down:
	docker-compose -f ./docker-compose.yml down

exec-src:
	make up-d
	docker-compose -f ./docker-compose.yml exec api bash || true

exec-geth:
	make up-d
	docker-compose -f ./docker-compose.yml exec geth bash || true

sol=
go=
sol-compile:
	make up-d
	docker-compose -f ./docker-compose.yml exec geth /bin/bash -c "solcjs --optimize --abi ./${sol}.sol -o build && solcjs --optimize --bin ./${sol}.sol -o build && abigen --abi=./build/${sol}_sol_${sol}.abi --bin=./build/${sol}_sol_${sol}.bin --pkg=contracts --out=../src/contracts/${go}.go"

generate-api-doc:
	make up-d
	docker-compose -f ./docker-compose.yml exec api /bin/bash -c "swag init ./main.go"
	cd src/docs && swagger2openapi --outfile ./v3/openapi.yaml ./swagger.yaml
