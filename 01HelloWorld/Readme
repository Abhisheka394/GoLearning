Created a directory

...
    mkdir 01helloWorld
    cd 1HelloWorld
    go mod init demo
...

- Explanation - the name demo is the name of the go module
- Ultimately this name demo becomes the virtual root path of the project

    go run main.go

- To see work directory
 
    go run --work main.go

- To build a binary
 
    go build main.go

    go build -o hello main.go


To cross Compile

    go tool dist list
    go tool dist list | grep darwin
        // darwin/amd64
        // darwin/arm64
    GOOS=darwin && GOARCH=arm64 go build -o macapp main.go