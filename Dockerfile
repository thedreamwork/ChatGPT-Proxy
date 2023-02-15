FROM golang:1-alpine as builder
RUN apk update && apk add make
WORKDIR /build
ADD . .
RUN make build

FROM alpine
COPY --from=builder /build/ChatGPT-PROXY /bin/ChatGPT-PROXY
RUN chmod +x /bin/ChatGPT-PROXY

ENTRYPOINT ["/bin/ChatGPT-PROXY"]