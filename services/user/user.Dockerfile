# Dockerfile References: https://docs.docker.com/engine/reference/builder/
############################
# STEP 1 build executable binary
############################
FROM golang:1.13.1-alpine AS builder

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Create appuser.
RUN adduser -D -g '' appuser

# Copy project to working directory
RUN mkdir /app
COPY . /app/
WORKDIR /app/services/user/server

RUN apk update && apk add git

# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main server.go plugin.go


############################
# STEP 2 build a small image
############################
FROM scratch

# # Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
# # Copy our static executable.
COPY --from=builder /app/services/user/server/main /app/main
# # Copy our static executable.
COPY --from=builder /app/core/config/app.config.dev.json /app/core/config/.app.config.json

WORKDIR /app

# # Use an unprivileged user.
USER appuser

ENV MICRO_REGISTRY 'kubernetes'

# Run the hello binary.
ENTRYPOINT ["/app/main"]
