FROM golang:alpine AS builder

RUN apk add build-base
COPY cmd /src/cmd
COPY internal /src/internal
COPY go.mod /src/
COPY go.sum /src/
RUN cd /src/cmd/LightAlert/ && CGO_ENABLED=0 go build -o /LightAlert .


FROM scratch

WORKDIR /data/LightAlert
WORKDIR /app

COPY --from=builder /LightAlert /app/

ENTRYPOINT ["./LightAlert"]