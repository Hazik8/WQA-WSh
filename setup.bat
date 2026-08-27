```bat
@echo off
setlocal EnableExtensions
chcp 65001 >nul
title WQA-WSh Developer Build

echo.
echo ==========================================
echo          WQA-WSh Developer Build
echo ==========================================
echo.

REM ==========================================
REM 1. Update dependencies
REM ==========================================
echo [1/6] Updating Go modules...
go mod tidy

if errorlevel 1 (
    echo.
    echo [ERROR] go mod tidy failed.
    pause
    exit /b 1
)

echo [OK] Dependencies updated.
echo.

REM ==========================================
REM 2. Format source code
REM ==========================================
echo [2/6] Formatting source code...
gofmt -w .

if errorlevel 1 (
    echo.
    echo [ERROR] gofmt failed.
    pause
    exit /b 1
)

echo [OK] Code formatted.
echo.

REM ==========================================
REM 3. Run tests
REM ==========================================
echo [3/6] Running tests...
go test ./...

if errorlevel 1 (
    echo.
    echo [ERROR] Tests failed.
    echo [STOP] Push cancelled.
    echo.
    pause
    exit /b 1
)

echo [OK] Tests passed.
echo.

REM ==========================================
REM 4. Build WSh
REM ==========================================
echo [4/6] Building WSh...
go build -o wsh.exe .\cmd\wsh

if errorlevel 1 (
    echo.
    echo [ERROR] WSh build failed.
    pause
    exit /b 1
)

echo [OK] wsh.exe built.
echo.

REM ==========================================
REM 5. Build WQA
REM ==========================================
echo [5/6] Building WQA...
go build -o wqa.exe .\cmd\wqa

if errorlevel 1 (
    echo.
    echo [ERROR] WQA build failed.
    pause
    exit /b 1
)

echo [OK] wqa.exe built.
echo.

REM ==========================================
REM 6. Commit and Push
REM ==========================================
echo [6/6] Uploading changes to GitHub...
echo.

git add .

git diff --cached --quiet

if not errorlevel 1 (
    echo [INFO] No changes to commit.
    echo.
    echo ==========================================
    echo       Nothing to push.
    echo ==========================================
    echo.
    pause
    exit /b 0
)

set "COMMIT_MESSAGE="
set /p "COMMIT_MESSAGE=Commit message: "

if "%COMMIT_MESSAGE%"=="" (
    set "COMMIT_MESSAGE=Update WQA-WSh"
)

echo.
echo Creating commit...
git commit -m "%COMMIT_MESSAGE%"

if errorlevel 1 (
    echo.
    echo [ERROR] Commit failed.
    pause
    exit /b 1
)

echo.
echo Pushing to GitHub...
git push

if errorlevel 1 (
    echo.
    echo [ERROR] Push failed.
    echo.
    pause
    exit /b 1
)

echo.
echo ==========================================
echo          BUILD COMPLETE
echo ==========================================
echo.
echo WSh:  wsh.exe
echo WQA:  wqa.exe
echo Git:  pushed successfully
echo.
echo Repository:
echo https://github.com/Hazik8/wqa-wsh
echo.
pause
endlocal
```
