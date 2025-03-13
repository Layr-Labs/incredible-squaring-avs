
RPC_URL=http://localhost:8545
export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# cd to the directory of this script so that this can be run from anywhere
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
cd "$parent_path"

root_dir=$(realpath "$parent_path/../..")

cd "$root_dir/contracts"
forge script script/CreateQuorum.s.sol --rpc-url $RPC_URL --broadcast --slow --private-key $PRIVATE_KEY
