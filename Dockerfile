FROM golang:1.7.3-alpine

RUN 				apk update && \
						apk upgrade && \
						apk add git

COPY 		. 	/go/src/github.com/chronojam/prom-scaler
WORKDIR 		/go/src/github.com/chronojam/prom-scaler

RUN 				ls -lrt
RUN 				go build -i -o prom-scaler cmd/prom-scaler/*.go

ENTRYPOINT ["./prom-scaler"]
