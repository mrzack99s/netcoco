WIN_BINARY = netcoco-windows.exe
LINUX_BINARY = netcoco-linux-amd64
GOARCH = amd64

LDFLAGS = "-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production'"
WIN_LDFLAGS = "-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production' -X 'github.com/mrzack99s/netcoco/pkgs/system.Os=windows'"

# Build binary
build: install tidy linux

install:
	go install ./netcoco
tidy:
	go mod tidy
windows:
	CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=${GOARCH} go build -ldflags ${WIN_LDFLAGS} -v -o ${WIN_BINARY} ./netcoco ;
	-chmod +x ${WIN_BINARY}
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -ldflags ${LDFLAGS} -v -o ${LINUX_BINARY} ./netcoco ;
	-chmod +x ${LINUX_BINARY}