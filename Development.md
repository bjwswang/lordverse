# Cosmos Development

## Quick start with Ignite CLI

Typical steps to use `ignite cli` for cosmos development.

### 1. Create a new chain

```shell
    ignite scaffold chain lordverse --no-module
```

### 2. Create a module

```shell
    ignite scaffold module lordverse --dep bank
```

### 3. Create a Message

`Message` is the core object in cosmos network.`ignite` provides code generation for `CRUD` on `Message`.

```shell
    ignite scaffold list Land owner location grade assessment extra  --module lordverse
```

## Client/Application Development

<https://docs.ignite.com/clients/go-client>

## How to do Tests?

TODO...

## Other Good Projects to help development

1. [awesome-cosmos](https://github.com/cosmos/awesome-cosmos)
2. [ping-pub/explorer](https://github.com/ping-pub/explorer) Light explorer for Cosmos-based Blockchains
