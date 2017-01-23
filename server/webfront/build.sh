EXE=webfront
CGO_ENABLED=0 GOOS=linux go build -a -o $EXE
docker build -t brainwave-studios/$EXE .
rm -rf $EXE