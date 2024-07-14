## [WIP] TEE Account Abstraction

Decentralized Account Abstraction implementation using trusted execution environments ([TEE](https://en.wikipedia.org/wiki/Trusted_execution_environment)) with [SUAVE](https://github.com/flashbots/suave-geth) for custodying the private keys.
https://collective.flashbots.net/t/account-abstraction-leveraging-tee/3670

## Usage

WIP

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Format

```shell
$ forge fmt
```

### Deploy

```shell
$ forge create --rpc-url https://rpc.rigil.suave.flashbots.net --private-key <private-key> src/Account.sol:Account --legacy
```
