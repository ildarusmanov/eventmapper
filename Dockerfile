# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8

# Copy the local package files to the container's workspace.
COPY . /go/src/eventmapper

# setup dependencies
WORKDIR /go/src/eventmapper
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure


RUN go install eventmapper

# Run the command by default when the container starts.
ENTRYPOINT /go/bin/eventmapper --configfile "/go/src/eventmapper/config.yml"

# Document that the service listens on port 8000.
EXPOSE 8000