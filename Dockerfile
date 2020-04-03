FROM golang as build
WORKDIR /go/src/app
COPY go.sum .
COPY go.mod .
COPY main.go .
COPY convert_test.go .
COPY convert.go .
WORKDIR /go/src/app
RUN go get .
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o /hcl2json .

##################
FROM scratch
LABEL maintainer="Thayne McCombs <https://github.com/tmccombs>"
COPY --from=build /hcl2json /hcl2json
ENTRYPOINT [ "/hcl2json" ]
