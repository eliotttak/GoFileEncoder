@echo off

set GOOS=%1
set GOARCH=%2
set EXT=

if "%GOOS%" == "" (
    FOR /F "delims=" %i IN ('go env GOOS') DO set GOOS=%%i
)

if "%GOARCH%" == "" (
    FOR /F "delims=" %i IN ('go env GOARCH') DO set GOARCH=%%i
)



if "%GOOS%" == "windows" (
    set EXT=.exe
)

echo Building from .\src\ to .\bin\portables\GoFileEncoder_portable_%GOOS%_%GOARCH%%EXT%

mkdir .\bin\portables\

rsrc -arch=%GOARCH% -ico icons\icon_16.ico,icons\icon_32.ico,icons\icon_64.ico,icons\icon_128.ico,icons\icon_256.ico -o .\src\rsrc.syso

go-bindata -pkg assets -o assets/bindata.go LICENSE
go build -o .\bin\portables\GoFileEncoder_portable_%GOOS%_%GOARCH%%EXT% .\src\