VERSION = v1beta
DIR_PATH = bin/
BINARY = ${DIR_PATH}netcoco-${VERSION}
GOARCH = amd64
LDFLAGS = "-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production' -X 'github.com/mrzack99s/netcoco/pkgs/system.Version=v1beta'"
WIN_LDFLAGS = "-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production' -X 'github.com/mrzack99s/netcoco/pkgs/system.Version=v1beta' -X 'github.com/mrzack99s/netcoco/pkgs/system.Os=windows'"
# Build the project
build: clean windows linux
windows:
	CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=${GOARCH} go build -ldflags ${WIN_LDFLAGS} -v -o ${BINARY}.exe ./runtime ;
	-chmod +x ${BINARY}.exe
linux:
	GOOS=linux GOARCH=${GOARCH} go build -ldflags ${LDFLAGS} -v -o ${BINARY} ./runtime ;
	-chmod +x ${BINARY}

clean:
	-rm -f ${DIR_PATH}netcoco-${VERSION}*