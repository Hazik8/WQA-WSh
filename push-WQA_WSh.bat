@echo off
setlocal

title WQA Git Push

echo ========================================
echo          WQA Git Push
echo ========================================
echo.

cd /d C:\WQA

if not exist ".git" (
    echo [ERROR] Git repository not found.
    pause
    exit /b 1
)

echo [1/4] Checking Git status...
git status
echo.

set /p "COMMIT_MSG=Enter commit name: "

if "%COMMIT_MSG%"=="" (
    echo.
    echo [ERROR] Commit name cannot be empty.
    pause
    exit /b 1
)

echo.
echo [2/4] Adding changes...
git add .
if errorlevel 1 (
    echo [ERROR] git add failed.
    pause
    exit /b 1
)

echo [OK] Changes added.
echo.

echo [3/4] Creating commit...
git commit -m "%COMMIT_MSG%"
if errorlevel 1 (
    echo.
    echo [ERROR] Commit failed.
    pause
    exit /b 1
)

echo [OK] Commit created.
echo.

echo [4/4] Pushing to GitHub...
git push origin main
if errorlevel 1 (
    echo.
    echo [ERROR] Push failed.
    pause
    exit /b 1
)

echo.
echo ========================================
echo          PUSH COMPLETED!
echo ========================================
echo.
echo Commit: %COMMIT_MSG%
echo Branch: main
echo.

pause
endlocal