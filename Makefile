VERSION = v1beta
DIR_PATH = bin/
BINARY = ${DIR_PATH}netcoco-${VERSION}
GOARCH = amd64
LDFLAGS = "-X 'main.product_mode=production' -X 'main.version=v1beta'"
# Build the project
build: clean windows linux
windows:
	CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=${GOARCH} go build -ldflags ${LDFLAGS} -v -o ${BINARY}.exe ./runtime ;
	-chmod +x ${BINARY}
linux:
	GOOS=linux GOARCH=${GOARCH} go build -ldflags ${LDFLAGS} -v -o ${BINARY} ./runtime ;
	
clean:
	-rm -f ${DIR_PATH}*