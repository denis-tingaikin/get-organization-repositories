FROM golang:1.15-buster as go
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOBIN=/bin
WORKDIR /build
COPY ./cmd .
RUN go build -o /bin/app .

FROM alpine as runtime
COPY --from=go /bin/app /bin/app
CMD /bin/app

ENTRYPOINT /bin/app