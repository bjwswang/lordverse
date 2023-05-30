#!/bin/bash

# Copyright 2023 The Bestchains Author
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e


echo "Create Voters"

lordversed tx dao create-voter 1  --from alice --yes
lordversed tx dao create-voter 1  --from bob --yes

echo "List Voters"

lordversed query dao list-voter

echo "Create Warehouse with threshold 2 and voters 0,1"

lordversed tx dao  create-warehouse 0,1 2 --from alice --yes

echo "Get Warehouse"

lordversed query dao list-warehouse

echo "Sign Warehouse from alice"
lordversed tx dao sign-warehouse 0 0 --from alice --yes 
lordversed tx dao sign-warehouse 0 1 --from bob --yes 

echo "Get Warehouse"
lordversed query dao list-warehouse

echo "Create a proposal"
lordversed tx dao create-proposal  0 "create-network-1" "I want to create a network" 0days0hours10minutes0seconds  --from alice --yes

echo "Get Proposal"
lordversed query dao list-proposal

echo "Update Proposal"
lordversed tx dao update-proposal  0 "I want to update a network" --from alice --yes


echo "Vote Proposal"
lordversed tx dao vote-proposal 1 0 1  --from alice --yes
lordversed tx dao vote-proposal 1 1 1  --from bob --yes