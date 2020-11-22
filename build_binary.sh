CURRENT_VERSION=`git describe --always`
go build -ldflags "-X main.version=$CURRENT_VERSION" cmd/teleirc.go
