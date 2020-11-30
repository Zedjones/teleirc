CURRENT_VERSION=`git describe --tags`

if [ $# -ge 1 ] && { [ $1 = "-a" ] || [ $1 = "--all" ]; }
then
    # Building different OSes for x86-64
    GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$CURRENT_VERSION" -o output/teleirc-x86-linux cmd/teleirc.go
    GOOS=openbsd GOARCH=amd64 go build -ldflags "-X main.version=$CURRENT_VERSION" -o output/teleirc-openbsd-linux cmd/teleirc.go
    GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$CURRENT_VERSION" -o output/teleirc-x86-windows.exe cmd/teleirc.go
    GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$CURRENT_VERSION" -o output/teleirc-x86-darwin cmd/teleirc.go
    # Building arm binary for Linux
    GOOS=linux GOARCH=arm go build -ldflags "-X main.version=$CURRENT_VERSION" -o output/teleirc-arm-linux cmd/teleirc.go
else
    go build -ldflags "-X main.version=$CURRENT_VERSION" cmd/teleirc.go
fi