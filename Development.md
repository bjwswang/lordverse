# Cosmos Development

## Quick start with Ignite CLI

Typical steps to use `ignite cli` for cosmos development.

### 1. Create a new chain

```shell
    ignite scaffold chain lordverse --no-module
```

### 2. Serve the chain in Dev mode

```shell
    ignite chain serve
```

> Note: auto build when changes happend including protobuf

### 3. Create a module

```shell
    ignite scaffold module lordverse --dep bank
```

### 4. Create a List

`Message` is the core object in cosmos network.`ignite` provides code generation for `CRUD` on `Message`.

```shell
    ignite scaffold list Land owner location grade assessment extra  --module lordverse
```

### 4. Define a Type

```shell
    ignite scaffold type Voter Address Weight --module dao
```

## Client/Application Development

<https://docs.ignite.com/clients/go-client>

## Test with CLI

> A CLI `{project}d` will be built automatically with command embeded in `{project}d tx xxx`

## Other Good Projects to help development

1. [awesome-cosmos](https://github.com/cosmos/awesome-cosmos)
2. [ping-pub/explorer](https://github.com/ping-pub/explorer) Light explorer for Cosmos-based Blockchains
3. [evmos](https://docs.evmos.org/)
