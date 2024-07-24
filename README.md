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
$ go run ./script deploy
```

### Verify

```shell
forge v --verifier blockscout  --chain 33626250 --verifier-url https://explorer.toliman.suave.flashbots.net/api <address> src/Account.sol:Account
```

## [WIP] Use Cases


### Sign up


```shell
$ go run ./script signup
```