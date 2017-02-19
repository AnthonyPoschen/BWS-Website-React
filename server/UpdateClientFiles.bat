:: Run this file as administrator or it doesn't work QQ
:: Build the new frontend and then delete the current version and then reset working directory
cd .\..\frontend & nwb build & cd .\..\server\webfront\www & del *.* & cd .\..\..\ & robocopy .\..\frontend\dist .\webfront\www /COPYALL /E