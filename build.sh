go fmt *.go
go fmt ./models/*.go
go fmt ./resources/icecreams/*.go
$GOPATH/bin/golint *.go
$GOPATH/bin/golint ./models/*.go
$GOPATH/bin/golint ./resources/icecreams/*.go
go vet github.com/zalora_icecream
go install
cp ./zalora.cfg $GOPATH/bin
go test ./resources/icecreams
go test ./models
go test ./

