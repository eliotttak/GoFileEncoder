@echo off

set GOOS=%1
set GOARCH=%2
set EXT=

if "%GOOS%" == "" (
    FOR /F "delims=" %%A IN ('go env GOOS') DO SET GOOS=%%A
)

if "%GOARCH%" == "" (
    FOR /F "delims=" %%A IN ('go env GOARCH') DO SET GOARCH=%%A
)

if "%GOOS%" == "windows" (
    set EXT=.exe
)

echo Building from .\pkg\ to .\bin\portables\GoFileEncoder_portable_%GOOS%_%GOARCH%%EXT%

mkdir .\bin\portables\

rsrc -arch=%GOARCH% -ico icons\icon_16.ico,icons\icon_32.ico,icons\icon_64.ico,icons\icon_128.ico,icons\icon_256.ico -o .\pkg\rsrc.syso

go-bindata -pkg assets -o assets/bindata.go LICENSE

go build -o .\bin\portables\GoFileEncoder_portable_%GOOS%_%GOARCH%%EXT% .\pkg\
