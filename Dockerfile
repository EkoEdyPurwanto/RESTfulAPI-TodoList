### Build Stage ###
FROM golang:1.20.4-alpine3.16 AS BUILD

### LABEL INSTRUCTION (hanya metadata) ###
LABEL authors="eep"
LABEL company="pondok programmer" reachMe="https://github.com/EchoEdyP"

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Go modules manifests
COPY go.mod go.sum ./

### RUN INSTRUCTION (Build Stage) ###
RUN go mod download && go mod tidy

### Copy the source code into the container (build stage) ###
COPY . .

### RUN INSTRUCTION (Build Stage) ###
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todoApp ./cmd/todolist-api/main.go

### Production Stage ###
FROM alpine:3.16

# Set the working directory inside the container
WORKDIR /app

# Install curl
RUN apk --no-cache add curl

# Copy only the necessary files from the build stage
COPY --from=BUILD /go/src/app/todoApp .
COPY --from=BUILD /go/src/app/internal/database/postgres/migrations/ ./internal/database/postgres/migrations/

# Expose port 1234
EXPOSE 1234

### CMD INSTRUCTION (if container run) ###
CMD ["./todoApp"]
