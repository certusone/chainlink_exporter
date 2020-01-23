## Chainlink oracle exporter

This prometheus exporter watches a Chainlink oracle smart contract on the Ethereum blockchain and measures metrics on
request fulfillment.

In order to track whether the exporter is alive and following the chain we export `cl_mon_height` which indicates the
last block the exporter has seen.

### How to build

Running `make build` will build artifacts to `bin`.

The `Dockerfile` can also be used to build a container.

We provide a prebuilt container image on [DockerHub](https://hub.docker.com/r/certusone/chainlink_exporter): `certusone/chainlink_exporter`

### Configuration

The configuration needs to be passed in via the environment.

| Name | Description |
|------|-------------|
| LADDR | Listening address (e.g. `:8080`).
| RPC | Websocket URL of the ethereum node to connect to. |
| ADDRESS | The address of the oracle contract to watch. |
| NODE_ADDRESS | The address of the node that's fulfilling the requests. |
| LINK_ADDRESS | The address of the LINK ERC20 token contract. Defaults to the mainnet contract. |

### Metrics

| Name | Type | Description |
|------|-------------|----------|
| cl_mon_height | gauge | Last processed block number. |
| cl_mon_last_request | gauge | Block number in which the last request was received. |
| cl_mon_last_response | gauge | Block number in which the last response was sent from the oracle. |
| cl_mon_response_time_bucket | histogram | Number of blocks taken to fulfill requests. Histograms are partitioned by the label `spec_id`. |
| cl_mon_missed | counter | Number of missed requests. Labels indicate job/spec id, requester address. |
| cl_mon_fulfilled | counter | Number of fulfilled requests. Labels indicate job/spec id, requester address. |
| cl_mon_rewards | counter | Rewards collected in LINK. Labels indicate job/spec id, requester address and whether the request containing this payment was fulfilled successfully. **Only payments of fulfilled requests are withdrawable.** |
| cl_mon_eth_balance | gauge | Eth balance of the node account. |
| cl_mon_link_balance | gauge | LINK balance of the oracle contract. The value with `type=balance` is the ERC20 balance. The value with `type=withdrawable` is the withdrawable balance. |

### Error handling

In case of errors during startup the program will panic. Errors during runtime are printed to the console and might
lead to the exporter not processing blocks. This will be visible in prometheus as `cl_mon_height` will stop increasing.

The client will automatically try to reconnect and -subscribe once the endpoint becomes available again.
