@echo off

:: bash -c "inotifywatch -e modify -e create -e moved_to -t 1 '%SRC%'"
set REL_PATH=.\..\..\frontend\src
set FRONTEND_DIR=
set START_PATH=%CD%
set WWW_PATH=%~pd0www
:: this is used for the bash script to append the directory
set BASH_MNT=\mnt\

:: Need to get the drive letter so this repo can be cloned on any drive letter
:: we also need to make sure its lowercase because linux cant handle case mismath
set DRIVE=%~d0
set DRIVE=%DRIVE:~0,1%

:: on googling how to lowercase in batch it looked like cancer so good old 
:: calling a linux command to save the day (Microsoft plz enhance user experience for this)
for /f "delims=" %%i in ('bash -c "echo '%DRIVE%' | tr [:upper:] [:lower:]"') do set output=%%i

set DRIVE=%output: =%

:: Move to bat file directory
pushd %~pd0
:: Save current directory and change to target directory
pushd %REL_PATH%

:: Save value of CD variable (current directory)
set FRONTEND_DIR=%CD%
:: Save the bash version and convert all the slashes to unix style so it doesnt freak out
set BASH_FRONTEND_DIR=%BASH_MNT%%DRIVE%\%CD:~3%
set BASH_FRONTEND_DIR=%BASH_FRONTEND_DIR:\=/%

:: Restore original directory
popd
popd



:: lets now begin the loop that will auto rebuild on change and copy to the 
:loop

bash -c "inotifywait -e modify -e create -e delete -e moved_to -e moved_from -r '%BASH_FRONTEND_DIR%' 2>/dev/null"
:: when inotify Returns we know something got modified so lets rebuild than copy.
::timeout /t 5 /nobreak
pushd %FRONTEND_DIR%
nwb build-react-app

xcopy /s /y %FRONTEND_DIR%\dist %WWW_PATH%
goto loop