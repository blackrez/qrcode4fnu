FROM golang:1.12 as builder

WORKDIR /go/src/github.com/blackrez/qrcode4fnu
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -v

FROM  gcr.io/distroless/static
COPY --from=builder /go/src/github.com/blackrez/qrcode4fnu/qrcode4fnu /qrcode4fnu

EXPOSE 8080

CMD ["/qrcode4fnu"]


