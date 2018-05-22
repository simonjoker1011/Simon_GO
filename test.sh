# !/bin/sh
export GOPATH=$GOPATH:`pwd`
#echo 'GOPATH: '$GOPATH

if [[ ! -n $1 ]]; then
	#statements
	echo 'Available packages:'
	go list ./...
else
	go test -v $1
fi

