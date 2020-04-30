#!/bin/bash -e

# The directory above this one
SALTICIDAE_GO_PATH=${SALTICIDAE_GO_PATH:-$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )}
export SALTICIDAE_PATH="$SALTICIDAE_GO_PATH/salticidae"

if [[ "$OSTYPE" == "linux-gnu" ]]; then
    export CGO_CFLAGS="-I$SALTICIDAE_PATH/build/include/"
    export CGO_LDFLAGS="-L$SALTICIDAE_PATH/build/lib/ -lsalticidae -luv -lssl -lcrypto -lstdc++"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    export CGO_CFLAGS="-I/usr/local/opt/openssl/include"
    export CGO_LDFLAGS="-L/usr/local/opt/openssl/lib/ -lsalticidae -luv -lssl -lcrypto"
else
    echo "Your operating system is not supported."
    exit 1
fi

