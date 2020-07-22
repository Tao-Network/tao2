#!/bin/sh

# vars from docker env
# - IDENTITY (default to empty)
# - PASSWORD (default to empty)
# - PRIVATE_KEY (default to empty)
# - BOOTNODES (default to empty)
# - EXTIP (default to empty)
# - VERBOSITY (default to 3)
# - MAXPEERS (default to 25)
# - SYNC_MODE (default to 'full')
# - NETWORK_ID (default to '558')
# - WS_SECRET (default to empty)
# - NETSTATS_HOST (default to 'netstats-server:3000')
# - NETSTATS_PORT (default to 'netstats-server:3000')

# constants
DATA_DIR="data"
KEYSTORE_DIR="keystore"

# variables
genesisPath=""
params=""
accountsCount=$(
  tao account list --datadir $DATA_DIR  --keystore $KEYSTORE_DIR \
  2> /dev/null \
  | wc -l
)

# file to env
for env in IDENTITY PASSWORD PRIVATE_KEY BOOTNODES WS_SECRET NETSTATS_HOST \
           NETSTATS_PORT EXTIP SYNC_MODE NETWORK_ID ANNOUNCE_TXS STORE_REWARD DEBUG_MODE MAXPEERS; do
  file=$(eval echo "\$${env}_FILE")
  if [[ -f $file ]] && [[ ! -z $file ]]; then
    echo "Replacing $env by $file"
    export $env=$(cat $file)
  elif [[ "$env" == "BOOTNODES" ]] && [[ ! -z $file ]]; then
    echo "Bootnodes file is not available. Waiting for it to be provisioned..."
    while true ; do
      if [[ -f $file ]] && [[ $(grep -e enode $file) ]]; then
        echo "Fount bootnode file."
        break
      fi
      echo "Still no bootnodes file, sleeping..."
      sleep 5
    done
    export $env=$(cat $file)
  fi
done

# networkid
if [[ ! -z $NETWORK_ID ]]; then
  case $NETWORK_ID in
    558 )
      genesisPath="mainnet.json"
      ;;
    559 )
      genesisPath="testnet.json"
      params="$params --tao-testnet --gcmode archive --rpcapi db,eth,net,web3,debug,posv"
      ;;
    90 )
      genesisPath="devnet.json"
      ;;
    * )
      echo "network id not supported"
      ;;
  esac
  params="$params --networkid $NETWORK_ID"
fi

# custom genesis path
if [[ ! -z $GENESIS_PATH ]]; then
  genesisPath="$GENESIS_PATH"
fi

# data dir
if [[ ! -d $DATA_DIR/tao ]]; then
  echo "No blockchain data, creating genesis block."
  tao init $genesisPath --datadir $DATA_DIR 2> /dev/null
fi

# identity
if [[ -z $IDENTITY ]]; then
  IDENTITY="unnamed_$(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c6)"
fi

# password file
if [[ ! -f ./password ]]; then
  if [[ ! -z $PASSWORD ]]; then
    echo "Password env is set. Writing into file."
    echo "$PASSWORD" > ./password
  else
    echo "No password set (or empty), generating a new one"
    $(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32} > password)
  fi
fi

# private key
if [[ $accountsCount -le 0 ]]; then
  echo "No accounts found"
  if [[ ! -z $PRIVATE_KEY ]]; then
    echo "Creating account from private key"
    echo "$PRIVATE_KEY" > ./private_key
    tao  account import ./private_key \
      --datadir $DATA_DIR \
      --keystore $KEYSTORE_DIR \
      --password ./password
    rm ./private_key
  else
    echo "Creating new account"
    tao account new \
      --datadir $DATA_DIR \
      --keystore $KEYSTORE_DIR \
      --password ./password
  fi
fi
account=$(
  tao account list --datadir $DATA_DIR  --keystore $KEYSTORE_DIR \
  2> /dev/null \
  | head -n 1 \
  | cut -d"{" -f 2 | cut -d"}" -f 1
)
echo "Using account $account"
params="$params --unlock $account"

# bootnodes
if [[ ! -z $BOOTNODES ]]; then
  params="$params --bootnodes $BOOTNODES"
fi

# extip
if [[ ! -z $EXTIP ]]; then
  params="$params --nat extip:${EXTIP}"
fi

# syncmode
if [[ ! -z $SYNC_MODE ]]; then
  params="$params --syncmode ${SYNC_MODE}"
fi

# netstats
if [[ ! -z $WS_SECRET ]]; then
  echo "Will report to netstats server ${NETSTATS_HOST}:${NETSTATS_PORT}"
  params="$params --ethstats ${IDENTITY}:${WS_SECRET}@${NETSTATS_HOST}:${NETSTATS_PORT}"
else
  echo "WS_SECRET not set, will not report to netstats server."
fi

# annonce txs
if [[ ! -z $ANNOUNCE_TXS ]]; then
  params="$params --announce-txs"
fi

# store reward
if [[ ! -z $STORE_REWARD ]]; then
  params="$params --store-reward"
fi

# debug mode
if [[ ! -z $DEBUG_MODE ]]; then
  params="$params --gcmode archive --rpcapi db,eth,net,web3,debug,posv"
fi

# maxpeers
if [[ -z $MAXPEERS ]]; then
  MAXPEERS=25
fi

# dump
echo "dump: $IDENTITY $account $BOOTNODES"

set -x

echo '[
  "enode://023628d1b2b87066d8c0af658d67fa76d17d4e870d70ba4a37512869357407bad8f485fcbfe145289d16c06096a3b4445d7dd4d537b6cf6be9cce0b5cec32bb3@108.61.242.3:20202",
  "enode://7b3717f1318eb10e0e04a2590741fd9521a80d8bf5bb715d6507618d930f9d94a2a8edeff812e8d0e4406fa573d176c004e71e7f142aeeb088756875852c68c8@149.28.69.174:20202"
]' > $DATA_DIR/tao/static-nodes.json

exec tao $params \
  --verbosity $VERBOSITY \
  --datadir $DATA_DIR \
  --keystore $KEYSTORE_DIR \
  --identity $IDENTITY \
  --maxpeers $MAXPEERS \
  --password ./password \
  --port 20202 \
  --txpool.globalqueue 5000 \
  --txpool.globalslots 5000 \
  --rpc \
  --rpccorsdomain "*" \
  --rpcaddr 0.0.0.0 \
  --rpcport 8545 \
  --rpcvhosts "*" \
  --ws \
  --wsaddr 0.0.0.0 \
  --wsport 8546 \
  --wsorigins "*" \
  --mine \
  --gasprice "250000000" \
  --targetgaslimit "84000000" \
  "$@"
