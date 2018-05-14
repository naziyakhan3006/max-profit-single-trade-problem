FROM golang:1.9

WORKDIR /go/src
ENV GOBIN /go/bin
RUN  mkdir -p problem.solving/
COPY . problem.solving/

RUN go env
RUN go install -v ./...
RUN go test -v ./...

CMD ["cmd"]
