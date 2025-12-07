@echo off
setlocal enabledelayedexpansion

set VERSION=0.9.0

echo ====================================
echo   Silk Release Packager v%VERSION%
echo ====================================
echo.

REM Create releases directory
if not exist "releases" mkdir releases
cd releases

REM Clean up old releases
if exist "silk-console-v%VERSION%" rmdir /s /q "silk-console-v%VERSION%"
if exist "silk-tray-v%VERSION%" rmdir /s /q "silk-tray-v%VERSION%"
if exist "silk-tray-only-v%VERSION%" rmdir /s /q "silk-tray-only-v%VERSION%"
if exist "silk-console-v%VERSION%.zip" del "silk-console-v%VERSION%.zip"
if exist "silk-tray-v%VERSION%.zip" del "silk-tray-v%VERSION%.zip"
if exist "silk-tray-only-v%VERSION%.zip" del "silk-tray-only-v%VERSION%.zip"

echo [1/3] Packaging silk-console-v%VERSION%...
mkdir "silk-console-v%VERSION%"
copy "..\silk.exe" "silk-console-v%VERSION%\" >nul
xcopy "..\web" "silk-console-v%VERSION%\web\" /E /I /Y >nul
copy "..\README.md" "silk-console-v%VERSION%\" >nul
copy "..\LICENSE" "silk-console-v%VERSION%\" >nul
copy "..\VLC_SETUP.md" "silk-console-v%VERSION%\" >nul
tar -a -c -f "silk-console-v%VERSION%.zip" "silk-console-v%VERSION%"
echo    ✓ silk-console-v%VERSION%.zip created

echo [2/3] Packaging silk-tray-v%VERSION%...
mkdir "silk-tray-v%VERSION%"
copy "..\silk-tray.exe" "silk-tray-v%VERSION%\" >nul
copy "..\icon.ico" "silk-tray-v%VERSION%\" >nul
xcopy "..\web" "silk-tray-v%VERSION%\web\" /E /I /Y >nul
copy "..\README.md" "silk-tray-v%VERSION%\" >nul
copy "..\LICENSE" "silk-tray-v%VERSION%\" >nul
copy "..\VLC_SETUP.md" "silk-tray-v%VERSION%\" >nul
tar -a -c -f "silk-tray-v%VERSION%.zip" "silk-tray-v%VERSION%"
echo    ✓ silk-tray-v%VERSION%.zip created

echo [3/3] Packaging silk-tray-only-v%VERSION%...
mkdir "silk-tray-only-v%VERSION%"
copy "..\silk-tray-only.exe" "silk-tray-only-v%VERSION%\" >nul
copy "..\icon.ico" "silk-tray-only-v%VERSION%\" >nul
xcopy "..\web" "silk-tray-only-v%VERSION%\web\" /E /I /Y >nul
copy "..\README.md" "silk-tray-only-v%VERSION%\" >nul
copy "..\LICENSE" "silk-tray-only-v%VERSION%\" >nul
copy "..\VLC_SETUP.md" "silk-tray-only-v%VERSION%\" >nul
tar -a -c -f "silk-tray-only-v%VERSION%.zip" "silk-tray-only-v%VERSION%"
echo    ✓ silk-tray-only-v%VERSION%.zip created

echo.
echo ====================================
echo Release packages created in releases/
echo ====================================
echo.
echo Files ready for upload:
echo   - silk-console-v%VERSION%.zip
echo   - silk-tray-v%VERSION%.zip
echo   - silk-tray-only-v%VERSION%.zip
echo.

cd ..
pause