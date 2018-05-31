# !/bin/sh
export GOPATH=$GOPATH:`pwd`
#echo 'GOPATH: '$GOPATH
arr=($(go list ./...))

function runtest(){
	GOCACHE=off go test -v $1
}

if [[ ! -n $1 ]]; then
	echo 'Available packages: (-a for all)'

	for e in ${arr[@]}; do
		echo $e
	done
elif [[ $1 == '-a' ]]; then
		#statements	
	for e in ${arr[@]}; do
		echo "testing... "$e
		runtest $e
	#	go test -v $e
	done
else
	runtest $1
	#go test -v $1
fi

