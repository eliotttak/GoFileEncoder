@echo off

set GOOS=%1
set GOARCH=%2
set EXT=

if "%GOOS%" == "windows" (
    set EXT=.exe
)

echo Building from .\setups\setup\%GOOS%\ to .\bin\setups\GoFileEncoder_setup_%GOOS%_%GOARCH%%EXT%
mkdir .\bin\setups\

rsrc -arch=%GOARCH% -ico icons\icon_16.ico,icons\icon_32.ico,icons\icon_64.ico,icons\icon_128.ico,icons\icon_256.ico -o .\setups\setup_%GOOS%\rsrc.syso


go-bindata -pkg assets -o assets/bindata.go LICENSE bin\portables\GoFileEncoder_portable_%GOOS%_%GOARCH%%EXT%
go build -o .\bin\setups\GoFileEncoder_setup_%GOOS%_%GOARCH%%EXT% .\setups\setup_%GOOS%\