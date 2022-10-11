FROM golang as build
WORKDIR /go/src/app
COPY go.sum .
COPY go.mod .
COPY main.go .
COPY convert/ ./convert/
WORKDIR /go/src/app
ENV GOPATH= CGO_ENABLED=0
RUN go get .
RUN go build -a -ldflags '-s' -o /hcl2json .

##################
FROM scratch
LABEL maintainer="Thayne McCombs <https://github.com/tmccombs>"
COPY --from=build /hcl2json /hcl2json
ENTRYPOINT [ "/hcl2json" ]
