#!/bin/bash

# Set variables
RPC_ENDPOINT="http://localhost:8080"
CONTRACT_ADDRESS="0x4F30fc8417172821c004a571Cb70085EA3c3C888"
POOL_MANAGER_ADDRESS="0x75E7c1Fd26DeFf28C7d1e82564ad5c24ca10dB14"
TOKEN0_ADDRESS="0x5FbDB2315678afecb367f032d93F642f64180aa3"
TOKEN1_ADDRESS="0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"

# Function to make RPC calls
make_rpc_call() {
    local method=$1
    local params=$2
    local id=$3
    
    echo "Sending request to $RPC_ENDPOINT:"
    echo "Method: $method"
    echo "Params: $params"
    echo "ID: $id"
    
    response=$(curl -v -X POST $RPC_ENDPOINT -H "Content-Type: application/json" -d '{
        "jsonrpc":"2.0",
        "method":"'$method'",
        "params":'$params',
        "id":'$id'
    }' 2>&1)
    
    if [ $? -ne 0 ]; then
        echo "Error: curl command failed"
        echo "Response:"
        echo "$response"
        return 1
    fi
    
    echo "Response:"
    echo "$response"
    
    # Check if the response is valid JSON
    if ! echo "$response" | jq . >/dev/null 2>&1; then
        echo "Error: Invalid JSON response"
        return 1
    fi
    
    echo "$response"
}

# 1. Initialize Pool
echo "Initializing Pool..."
init_pool_response=$(make_rpc_call "initializePool" '{
    "poolManager": "'$POOL_MANAGER_ADDRESS'",
    "currency0": "'$TOKEN0_ADDRESS'",
    "currency1": "'$TOKEN1_ADDRESS'",
    "fee": "3000",
    "tickSpacing": "60",
    "sqrtPriceX96": "79228162514264337593543950336"
}' 1)
echo "$init_pool_response"

# 2. Approve tokens
echo "Approving Token0..."
approve_token0_response=$(make_rpc_call "approveToken" '{
    "token": "'$TOKEN0_ADDRESS'",
    "spender": "'$CONTRACT_ADDRESS'",
    "amount": "1000000000000000000000"
}' 2)
echo "$approve_token0_response"

echo "Approving Token1..."
approve_token1_response=$(make_rpc_call "approveToken" '{
    "token": "'$TOKEN1_ADDRESS'",
    "spender": "'$CONTRACT_ADDRESS'",
    "amount": "1000000000000000000000"
}' 3)
echo "$approve_token1_response"

# 3. Add Liquidity
echo "Adding Liquidity..."
add_liquidity_response=$(make_rpc_call "addLiquidity" '{
    "key": {
        "currency0": "'$TOKEN0_ADDRESS'",
        "currency1": "'$TOKEN1_ADDRESS'",
        "fee": "3000",
        "tickSpacing": 60,
        "hooks": "'$CONTRACT_ADDRESS'"
    },
    "amountEach": "1000000000000000000"
}' 4)
echo "$add_liquidity_response"

# Extract transaction hash
tx_hash=$(echo "$add_liquidity_response" | jq -r '.result')

if [ "$tx_hash" != "null" ]; then
    echo "Transaction hash: $tx_hash"

    # 4. Get transaction receipt
    echo "Getting transaction receipt..."
    receipt_response=$(make_rpc_call "eth_getTransactionReceipt" '["'$tx_hash'"]' 5)
    echo "$receipt_response"

    # 5. Debug trace transaction
    echo "Getting debug trace..."
    debug_trace_response=$(make_rpc_call "debug_traceTransaction" '["'$tx_hash'", {}]' 6)
    echo "$debug_trace_response"
else
    echo "Failed to add liquidity. No transaction hash returned."
fi

# 6. Check Balances
echo "Checking Token0 Balance..."
balance0_response=$(make_rpc_call "getBalance" '{
    "token": "'$TOKEN0_ADDRESS'",
    "account": "'$CONTRACT_ADDRESS'"
}' 7)
echo "$balance0_response"

echo "Checking Token1 Balance..."
balance1_response=$(make_rpc_call "getBalance" '{
    "token": "'$TOKEN1_ADDRESS'",
    "account": "'$CONTRACT_ADDRESS'"
}' 8)
echo "$balance1_response"