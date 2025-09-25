@echo off

set GOOS=%1
set GOARCH=%2
set EXT=
set BUILDOPTIONS=

set count=0
for %%a in (%*) do (
    set /a count+=1
    if !count! GEQ 3 set BUILDOPTIONS="!BUILDOPTIONS! %%A"
)

if "%GOOS%" == "" (
    FOR /F "delims=" %%A IN ('go env GOOS') DO SET GOOS=%%A
)

if "%GOARCH%" == "" (
    FOR /F "delims=" %%A IN ('go env GOARCH') DO SET GOARCH=%%A
)

if "%GOOS%" == "windows" (
    set EXT=.exe
)

set TARGET=.\bin\portables\GoFileEncoder_portable_%GOOS%_%GOARCH%%EXT%


echo Building from .\pkg\ to %TARGET%

mkdir .\bin\portables\

rm %TARGET%

echo Creating assets
go-bindata -pkg assets -o assets/bindata.go LICENSE

echo Building
go build -o %TARGET% .\pkg\

if "%GOOS%" == "windows" (
    echo Creating icon
    svg_to_ico --input icons\icon.svg --output icons\icon.ico

    echo Adding icons
    resourcehacker -open %TARGET% -save %TARGET% -action addoverwrite -resource icons\icon.ico -mask ICONGROUP,MAINICON,

    rm icons\icon.ico
)