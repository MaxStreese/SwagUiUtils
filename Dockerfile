#===============================================================================
# STEP 1 build executable binary
#===============================================================================
FROM golang:alpine as builder

# Install git & SSL ca certificates.
RUN apk update && apk add --no-cache git ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

# Copy all source files and set working directory
COPY . $GOPATH/src/github.com/maxstreese/swaguiutils
WORKDIR $GOPATH/src/github.com/maxstreese/swaguiutils/cmd/swaguiserver

# Download all dependencies
RUN GO111MODULE=on go mod download

# Build the binary
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/swaguiserver

# ==============================================================================
# STEP 2 build a small image
# ==============================================================================
FROM scratch

# Copy security stuff from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy the binary from builder
COPY --from=builder /go/bin/swaguiserver /go/bin/swaguiserver

# Use the previously created unprivileged appuser
USER appuser

# Run the binary
ENTRYPOINT ["/go/bin/swaguiserver"]
