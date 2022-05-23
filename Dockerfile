FROM golang:1.18-alpine as builder
workdir /go/src/devopstom.com/simpleip-api
RUN apk --no-cache add ca-certificates
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simpleip-api .
RUN chmod +x simpleip-api

FROM scratch
WORKDIR /
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/devopstom.com/simpleip-api/simpleip-api ./
ENTRYPOINT ["/simpleip-api"]
CMD ["/simpleip-api"]
