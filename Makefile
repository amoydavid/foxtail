BINARY="foxtail"
VERSION=1.0.0
BUILD=`date +%FT%T%z`

default:
	go build --tags json1 -o ${BINARY}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: default clean