FROM golang:1.7.3-alpine

RUN 		apk update && \
				apk upgrade && \
				apk add git

ADD 		. 	/build
WORKDIR 		/build/cmd/prom-scaler

RUN 		go get
RUN 		go build -i -o prom-scaler .

ENTRYPOINT ["prom-scaler"]
