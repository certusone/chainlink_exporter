abi/aggregator.go: contracts/Aggregator.abi
	 abigen --abi ./contracts/Aggregator.abi --pkg abi --type Aggregator --out abi/aggregator.go

abi/oracle.go: contracts/Oracle.abi
	 abigen --abi ./contracts/Oracle.abi --pkg abi --type Oracle --out abi/oracle.go

abi/link.go: contracts/LinkToken.abi
	 abigen --abi ./contracts/LinkToken.abi --pkg abi --type ERC --out abi/link.go

.PHONY: abi
abi: abi/aggregator.go abi/oracle.go abi/link.go

.PHONY: build
build: abi
	CGO_ENABLED=0 go build -o bin/chainlink_exporter ./cmd/chainlink_exporter
