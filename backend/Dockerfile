FROM golang:1.15 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch

ARG COMMIT_SHA
ENV COMMIT_SHA=${COMMIT_SHA}

COPY --from=builder /app/elosimulator /app/elosimulator

ENTRYPOINT [ "/app/elosimulator" ]