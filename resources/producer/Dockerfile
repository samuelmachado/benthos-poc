  
FROM golang:1.16-stretch AS builder
WORKDIR /api
COPY . .


RUN go env -w GOPROXY=https://proxy.golang.org
RUN go env -w CGO_ENABLED="0"
RUN go env -w GOPRIVATE=github.com/samuelmachado
RUN go env -w GO111MODULE='on'

# RUN go build 

FROM alpine:latest
RUN mkdir -p /go/bin && \
    mkdir /logs && \
    chmod -R 755 /logs
# COPY --from=builder /api/bin/api .
# ENTRYPOINT /go/bin/api