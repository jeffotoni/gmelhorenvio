FROM golang:1.19 as builder
WORKDIR /go/src/gmelhorenvio
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o gmelhorenvio
RUN cp gmelhorenvio /go/bin/

FROM alpine:latest AS final
RUN apk update
RUN apk add --no-cache tzdata
RUN apk add --no-cache ca-certificates
ENV TZ="America/Sao_Paulo"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /go/bin/gmelhorenvio /
CMD ["/gmelhorenvio"]
