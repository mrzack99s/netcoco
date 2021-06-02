VERSION = v1.0-alpha
DIR_PATH = bin/
WIN_BINARY = ${DIR_PATH}netcoco-windows.exe
LINUX_BINARY = ${DIR_PATH}netcoco-linux-amd64
GOARCH = amd64
LDFLAGS = "-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production' -X 'github.com/mrzack99s/netcoco/pkgs/system.Version=${VERSION}'"
WIN_LDFLAGS = "-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production' -X 'github.com/mrzack99s/netcoco/pkgs/system.Version=${VERSION}' -X 'github.com/mrzack99s/netcoco/pkgs/system.Os=windows'"
# Build the project
build: clean windows linux
windows:
	CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=${GOARCH} go build -ldflags ${WIN_LDFLAGS} -v -o ${WIN_BINARY} ./runtime ;
	-chmod +x ${WIN_BINARY}
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -ldflags ${LDFLAGS} -v -o ${LINUX_BINARY} ./runtime ;
	-chmod +x ${LINUX_BINARY}

clean:
	-rm -f ${DIR_PATH}netcoco-*