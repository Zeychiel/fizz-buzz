# Accept the Go version for the image to be set as a build argument.
ARG GO_VERSION=1.18

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder


# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
RUN apk add --no-cache ca-certificates

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
# * GOFLAGS=-mod=vendor to force `go build` to look into the `/vendor` folder.
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN go build \
    -installsuffix 'static' \
    -o /app ./cmd/server

# Create logs folder
RUN mkdir /src/logs

# Final stage: the running container.
FROM scratch AS final


# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the second stage.
COPY --from=builder /app /app

# Copy conf files
COPY --from=builder /src/.env /src/firebase.json ./

# Copy logs folder
COPY --from=builder /src/logs /logs

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080


# Run the compiled binary.
ENTRYPOINT ["/app"]

