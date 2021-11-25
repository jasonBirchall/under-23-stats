FROM golang:alpine as builder

WORKDIR /go/src/app

# Get reflex for live reload in dev
RUN go get -u github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./report

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY templates templates

# Copy from builder
COPY --from=builder /go/src/app/report /app/report

EXPOSE 8080

CMD ["./report"]