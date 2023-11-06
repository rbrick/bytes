bytes.onl
---
> Backend API for aggregating price data for ERC20 tokens
### Building

```sh
go build
```

### Testing

```sh
go test
```

### Commands

##### abifix
Description: 

Used to generate proper ABI (application binary interface) files for the Solidity compiler tool, solc. It removes
unncessary information from these files and writes them to a new file suffixed with `_fixed`. 

Usage:

```sh
./bytes abifix [file]
```