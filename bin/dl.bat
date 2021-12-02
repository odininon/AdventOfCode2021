@echo off
set day=%1
set output=day%day%.txt

cd /D "%~dp0"

.\aocdl.exe -day %day% -output ..\inputs\%output%