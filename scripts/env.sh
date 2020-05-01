export GOPATH="$(go env GOPATH)"
export SALTICIDAE_ORG="ava-labs"
export SALTICIDAE_GO_VER="v0.1.0"
export SALTICIDAE_GO_PATH="$GOPATH/src/github.com/$SALTICIDAE_ORG/salticidae-go"

export SALTICIDAE_PATH="$SALTICIDAE_GO_PATH/salticidae"
export CGO_CFLAGS="-I$SALTICIDAE_PATH/build/include/"
export CGO_LDFLAGS="-L$SALTICIDAE_PATH/build/lib/ -lsalticidae -luv -lssl -lcrypto -lstdc++"
