DIR=$(cd ../; pwd)
export GOPATH=$DIR:$GOPATH
GOOS=linux   GOARCH=amd64 go build -o ../bin/coslist_linux_amd64       main.go
GOOS=linux   GOARCH=386   go build -o ../bin/coslist_linux_386         main.go
GOOS=windows GOARCH=amd64 go build -o ../bin/coslist_windows_amd64.exe main.go
GOOS=windows GOARCH=386   go build -o ../bin/coslist_windows_386.exe   main.go
GOOS=darwin  GOARCH=amd64 go build -o ../bin/coslist_darwin_amd64      main.go
GOOS=darwin  GOARCH=386   go build -o ../bin/coslist_darwin_386        main.go