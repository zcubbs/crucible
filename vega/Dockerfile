FROM --platform=$BUILDPLATFORM golang:1.18-stretch AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environmet variables needed for our image and build the API server.
ARG TARGETOS
ARG TARGETARCH
ARG GH_ACCESS_TOKEN
RUN git config --global url.https://$GH_ACCESS_TOKEN@github.com/.insteadOf https://github.com/
RUN GOWORK=off GOOS=$TARGETOS GOARCH=$TARGETARCH go build \
    -ldflags="-s -w" -o app .

FROM alpine:3.14

RUN mkdir cmd
COPY --from=builder ["/build/app", "/cmd"]

# Export necessary port.
WORKDIR /cmd

# Command to run when starting the container.
ENTRYPOINT ["./app"]
