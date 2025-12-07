@echo off
echo Building Silk variants...
echo.

echo [1/3] Building silk.exe (Console + Browser)...
go build -o silk.exe
if %errorlevel% neq 0 goto error

echo [2/3] Building silk-tray.exe (System Tray + Browser)...
go build -tags tray -ldflags -H=windowsgui -o silk-tray.exe
if %errorlevel% neq 0 goto error

echo [3/3] Building silk-tray-only.exe (System Tray Only)...
go build -tags trayonly -ldflags -H=windowsgui -o silk-tray-only.exe
if %errorlevel% neq 0 goto error

echo.
echo ✓ All builds completed successfully!
echo.
echo Files created:
echo   - silk.exe           (Console + Browser)
echo   - silk-tray.exe      (System Tray + Browser)  
echo   - silk-tray-only.exe (System Tray Only)
echo.
goto end

:error
echo.
echo ✗ Build failed!
echo.

:end