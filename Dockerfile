# Build stage,
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd

# Final stage (runtime environnement),
FROM ubuntu:24.04
WORKDIR /app

COPY --from=builder /app/main ./main
COPY template/ ./template/
COPY static/ ./static/
COPY uploads/ ./uploads/
COPY db/ ./db/
# metadata,
LABEL maintainer="mennaas | abalouri | abaid | ranniz | ychatoua"
LABEL version="1.0"
LABEL description="forum"
# the port,
EXPOSE 8080
# the comand to run,
CMD ["./main"]