# Geth DAPP DEMO

Implementing Smart Contract in the Golang

# Requirement

- Golang
- Gin
- Solidity
- Ganache
- go-ethereum
- swag
- Docker
- Docker Compose

# Set up

```bash
$ make setup
```

In "src/.env", set "GANACHE_HOST" to the network of the "geth" container

```bash
$ docker network inspect geth-dapp-demo-network
```

```
ex)

...
"Containers": {
    ...
    "xxxx" : {
        "Name": "geth-dapp-demo-geth-1", <- This is the name of the network to connect to the "geth" container.
    }
}
```

# Usage

1. Please change your Ganache private key.

You can check Ganache's Private Key with this command.

```bash
$ make exec-geth
$ ganache
```

Delete the leading "0x" of the acquired Private Key and paste it.

```golang
./src/contract_address/contract_address.go

priKey := "Your Private Key"
```

2. Get a Contract address.

```bash
$ make up
```

go to http://localhost:8080/swagger/index.html
call http://localhost:8080/api/v1/contract-address

3. Please change your contract address.

```golang
./src/controller/hello.go

constantAddress = "Your Contract Address"
```

4. call http://localhost:8080/api/v1/hello

If "Hello World" is displayed, it is successful!!

# Compile solidity to Golang

Compile the code in the ./contracts and output the Golang code in the ./src/contracts.

```bash
$ make sol-compile sol={solidity file name} go={golang file name}

ex)
make sol-compile sol=Hello go=hello
```

# Generate API docs

```bash
$ make generate-api-doc
```

SEE: https://github.com/swaggo/swag

# Container up

```bash
$ make up
```

# Container down

```bash
$ make down
```

# Exec src container

```bash
$ make exec-src
```

# Exec geth container

```bash
$ make exec-geth
```
