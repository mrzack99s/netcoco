@ECHO OFF

set WIN_PROD_LDFLAGS="-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production' -X 'github.com/mrzack99s/netcoco/pkgs/system.Os=windows'"
set LINUX_PROD_LDFLAGS="-X 'github.com/mrzack99s/netcoco/pkgs/system.Product_mode=production'"
set DEV_LDFLAGS="-X 'github.com/mrzack99s/netcoco/pkgs/system.Os=windows'"

if "%1" NEQ "" (
    CALL :%1
)

CALL :end

:build
  del /s /q  netcoco-windows.exe
  del /s /q  netcoco-linux-amd64
  del /s /q  netcoco-dev-windows.exe

  set GOOS=windows
  go build -ldflags %WIN_PROD_LDFLAGS% -o netcoco-windows.exe .\netcoco\main.go

  set GOOS=linux
  go build -ldflags %LINUX_PROD_LDFLAGS% -o netcoco-linux-amd64 .\netcoco\main.go

  CALL :end

:run
  del /s /q  netcoco-windows.exe
  del /s /q  netcoco-linux-amd64
  del /s /q  netcoco-dev-windows.exe

  go build -ldflags %DEV_LDFLAGS% -o netcoco-dev-windows.exe .\netcoco\main.go
  .\netcoco-dev-windows.exe

:end
   exit