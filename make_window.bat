@ECHO OFF

set VERSION="v1.0-alpha"
set BIN_PATH=bin
set WIN_PROD_LDFLAGS="-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production' -X 'github.com/mrzack99s/netcoco/pkgs/system.Version=%VERSION%' -X 'github.com/mrzack99s/netcoco/pkgs/system.Os=windows'"
set LINUX_PROD_LDFLAGS="-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production' -X 'github.com/mrzack99s/netcoco/pkgs/system.Version=%VERSION%'"
set DEV_LDFLAGS="-X 'github.com/mrzack99s/netcoco/pkgs/system.Version=%VERSION%' -X 'github.com/mrzack99s/netcoco/pkgs/system.Os=windows'"

if "%1" NEQ "" (
    CALL :%1
)

CALL :end

:build
  del /s /q  .\%BIN_PATH%\netcoco-*

  set GOOS=windows
  go build -ldflags %WIN_PROD_LDFLAGS% -o %BIN_PATH%/netcoco-windows.exe .\runtime

  set GOOS=linux
  go build -ldflags %LINUX_PROD_LDFLAGS% -o %BIN_PATH%/netcoco-linux-amd64 .\runtime

  CALL :end

:run
  del /s /q  .\%BIN_PATH%\netcoco-*
  go build -ldflags %DEV_LDFLAGS% -o %BIN_PATH%/netcoco-windows.exe .\runtime

  cd %BIN_PATH%
  .\netcoco-windows.exe

:end
   exit