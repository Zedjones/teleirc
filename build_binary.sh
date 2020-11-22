CURRENT_VERSION=`git describe --tags`
go build -ldflags "-X main.version=$CURRENT_VERSION" cmd/teleirc.go
