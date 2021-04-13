FROM golang:1.16
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOBIN=/bin
WORKDIR /build
COPY ./cmd .
RUN go install ./...
ENTRYPOINT ["cmd"]