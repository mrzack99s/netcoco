@ECHO OFF

set PROD_LDFLAGS="-X 'main.product_mode=production' -X 'main.version=v1beta'"
set DEV_LDFLAGS="-X 'main.product_mode=development' -X 'main.version=v1beta'"

if "%1" NEQ "" (
    CALL :%1
)

CALL :end

:build
  del /s /q  .\bin\*.*

  set GOOS=windows
  go build -ldflags %PROD_LDFLAGS% -o bin/netcoco.exe .\runtime
  CALL :end

:run
  del /s /q  .\bin\*.*
  go build -ldflags %DEV_LDFLAGS% -o bin/netcoco.exe .\runtime
  .\bin\netcoco.exe

:end
   exit