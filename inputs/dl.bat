@echo off
set day=%1
set output=day%day%.txt
.\aocdl.exe -day %day% -output %output%