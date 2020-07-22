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
  "enode://5f7a617191f84f365a0ddad6d35baa3c63291e94ccad77c65d6c8afa17044d020fa94c80f02f6036e3abd74f99040e4ffd14d5dae9ed1805f8d103fb1adb5b7d@45.32.227.55:20202",
  "enode://0de4801ceb78ff4403005531f0c6a5addc29fef6b669a1798e6fa65f5e7eac69af5f4f9523dd029a51447e0c541d5e54f4bf3c6aa7555f05e6fe30ae7bc1369e@140.82.50.27:20202",
  "enode://2b26d1921e948a296f3108e2a1460d143029d88c413fd7110eac3184ec85560d1615ce82936996eae332ff5b851576b450e19588f5fb317ab7ef638fcfe7da13@149.248.44.254:20202",
  "enode://395efafc2fe425d107cd39aa7bc633c9802d61842a962e5f328776017a8d660dc7566896664bd2949925fe625e07ea5a8a69d4804e425e66f4a0810afe1f947b@45.77.85.124:20202",
  "enode://d292b0b7a2af14934fec213963c40f84a2fcdcd44f13facdafa289574649d254a6d509235a2f8f043ce21028e9f695a739bd497a3b5b97ac3fa9a929b2b9101e@8.3.29.246:20202",
  "enode://8ed7c43edc84933817af1772c3607fa557fa8761b48a1b6cfc08f64b44542df4fc06094029b5fa745b8330152b37ac4706b8e564279077ff2f940706f9c637a4@108.61.242.3:20202"
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
