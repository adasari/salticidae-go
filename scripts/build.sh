#!/bin/bash -e

# Location of this script
SRC_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

source "${SRC_DIR}/env.sh"

if [[ "$OSTYPE" == "linux-gnu" ]]; then
    cd $SALTICIDAE_PATH
    cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX="$SALTICIDAE_PATH/build" .
    make -j4
    make install
elif [[ "$OSTYPE" == "darwin"* ]]; then
    brew install Determinant/salticidae/salticidae
else
    echo "Your operating system is not supported."
    exit 1
fi
