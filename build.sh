# !/bin/sh
rm -rf bin/*
export GOPATH=$GOPATH:`pwd`

echo 'GOPATH: '$GOPATH

go install -v -gcflags "-N -l" ./...
