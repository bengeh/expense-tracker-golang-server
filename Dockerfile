# Two-stage build:
#    first  FROM prepares a binary file in full environment ~780MB
#    second FROM takes only binary file ~10MB

FROM golang:1.9 AS builder

RUN go version

COPY . "/go/src/expense_tracker_golang_api_with_docker"
WORKDIR "/go/src/expense_tracker_golang_api_with_docker"

#RUN go get -v -t  .

RUN set -x && \
    #go get github.com/2tvenom/go-test-teamcity && \  
    go get github.com/golang/dep/cmd/dep && \
    dep init && \
    dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /golang_api_docker

CMD ["/golang_api_docker"]

EXPOSE 8000

