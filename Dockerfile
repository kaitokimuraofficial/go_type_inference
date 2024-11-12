FROM golang:1.23.0-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o go_type_inference

# ---------------------------------------------------

FROM debian:bullseye-slim AS deploy

RUN apt-get update

COPY --from=deploy-builder /app/go_type_inference .

CMD ["./go_type_inference"]
