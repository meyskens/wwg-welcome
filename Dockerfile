FROM golang:1.9 AS builder

ARG GO_ARCH

COPY ./ /go/src/github.com/meyskens/wwg-welcome
WORKDIR /go/src/github.com/meyskens/wwg-welcome

RUN go get 
RUN CGO_ENABLED=0 GOOS=linux GOARCH="${GO_ARCH}" go build -a -installsuffix cgo -o wwg-welcome ./

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/meyskens/wwg-welcome/wwg-welcome /wwg-welcome

CMD [ "/wwg-welcome" ]