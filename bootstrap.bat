@echo off
set day=%1
.\bin\dl.bat %day%
.\bin\hygen.exe day new --day %day%