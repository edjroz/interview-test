# Interview Test For Pocket Network
The purpose of this app is to test the candidates ability to conform to
the Pocket Network workflow, coding style, and process.

This project wraps an [Ethereum](github.com/ethereum/go-ethereum) Node. 

### Help Wanted
Please satisfy all TODO's in the project

### RPC 
Runs on port 1050

### Dependencies
This project uses [go modules](https://blog.golang.org/using-go-modules)

Go version > 1.12

`GO111MODULE="on" go get github.com/goware/modvendor`

`GO111MODULE="on" go mod vendor`

`modvendor -copy="**/*.c **/*.h"`

`go test ./...`

### RPC

Endpoint `localhost:1050/block/byNumber`

Params: height="0"

curl -X POST --data 'height=0' http://localhost:1050/block/byNumber

### Build

go build -o pockethereum github.com/andrewnguyen22/pocket-interview-test/app