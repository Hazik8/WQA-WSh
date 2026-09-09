@echo off
setlocal

title WinDroid WQA + WSh Updater

echo ========================================
echo       WinDroid WQA + WSh Updater
echo ========================================
echo.

cd /d C:\WQA

if not exist "go.mod" (
    echo [ERROR] C:\WQA\go.mod not found.
    pause
    exit /b 1
)

echo [1/6] Building WQA...
go build -o wqa.exe .\cmd\wqa
if errorlevel 1 (
    echo [ERROR] WQA build failed.
    pause
    exit /b 1
)

echo [OK] WQA built.
echo.

echo [2/6] Building WSh...
go build -o wsh.exe .\cmd\wsh
if errorlevel 1 (
    echo [ERROR] WSh build failed.
    pause
    exit /b 1
)

echo [OK] WSh built.
echo.

echo [3/6] Creating WinDroid directories...

if not exist "C:\WinDroid\Tools" mkdir "C:\WinDroid\Tools"
if not exist "C:\WinDroid\Bin" mkdir "C:\WinDroid\Bin"

echo [OK] Directories ready.
echo.

echo [4/6] Installing WQA...

copy /Y "C:\WQA\wqa.exe" "C:\WinDroid\Tools\wqa.exe" >nul
if errorlevel 1 (
    echo [ERROR] Failed to install WQA to Tools.
    pause
    exit /b 1
)

copy /Y "C:\WQA\wqa.exe" "C:\WinDroid\Bin\wqa.exe" >nul
if errorlevel 1 (
    echo [ERROR] Failed to install WQA to Bin.
    pause
    exit /b 1
)

echo [OK] WQA installed.
echo.

echo [5/6] Installing WSh...

copy /Y "C:\WQA\wsh.exe" "C:\WinDroid\Tools\wsh.exe" >nul
if errorlevel 1 (
    echo [ERROR] Failed to install WSh to Tools.
    pause
    exit /b 1
)

copy /Y "C:\WQA\wsh.exe" "C:\WinDroid\Bin\wsh.exe" >nul
if errorlevel 1 (
    echo [ERROR] Failed to install WSh to Bin.
    pause
    exit /b 1
)

echo [OK] WSh installed.
echo.

echo [6/6] Checking installation...
echo.

echo WQA:
"C:\WinDroid\Tools\wqa.exe" --version

echo.
echo WSh:
"C:\WinDroid\Tools\wsh.exe" --version

echo.
echo ========================================
echo       Installation completed!
echo ========================================
echo.
echo WQA:
echo   C:\WinDroid\Tools\wqa.exe
echo   C:\WinDroid\Bin\wqa.exe
echo.
echo WSh:
echo   C:\WinDroid\Tools\wsh.exe
echo   C:\WinDroid\Bin\wsh.exe
echo.

pause
endlocal