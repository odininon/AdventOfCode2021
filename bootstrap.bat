@echo off
set day=%1

cd /D "%~dp0"
call .\bin\dl.bat %day%

cd /D "%~dp0"
call .\bin\hygen.exe day new --day %day%