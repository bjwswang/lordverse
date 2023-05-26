# Voter

## Design

```protobuf
syntax = "proto3";
package lordverse.dao;

option go_package = "lordverse/x/dao/types";

message Voter {
  uint64 id = 1;
  uint32 weight = 2; 
  string creator = 3;
}
```

## Service

```protobuf
  // CreateWarehouse defines a rpc to create a voter.
  rpc CreateVoter     (MsgCreateVoter    ) returns (MsgCreateVoterResponse    );
  // UpdateVoter defines a rpc to update a voter.
  rpc UpdateVoter     (MsgUpdateVoter    ) returns (MsgUpdateVoterResponse    );
  // DeleteVoter defines a rpc to delete a voter.
  rpc DeleteVoter     (MsgDeleteVoter    ) returns (MsgDeleteVoterResponse    );
```

## Test

### Prerequisite

- `lordversed` chain is running
- `account` has coins(retrieved from Faucet)

### Workflow

1. Create a voter

```shell
lordversed tx dao  create-voter  2 --from bjwswang --keyring-backend test
```

2. Query a voter

```shell
❯ lordversed query dao list-voter
Voter:
- creator: cosmos175qqqusn0lwdax8jy6gncv08lprp5qfp7ptrtj
  id: "0"
  weight: 2
pagination:
  next_key: nu
```