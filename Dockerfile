# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8

# Copy the local package files to the container's workspace.
COPY . /go/src/eventmapper

# setup dependencies
RUN go get github.com/WajoxSoftware/middleware
RUN go get github.com/streadway/amqp
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/validator.v2
RUN go get gopkg.in/yaml.v2
RUN go get golang.org/x/net/context
RUN go get github.com/golang/protobuf/proto
RUN go get google.golang.org/grpc

RUN go install eventmapper

VOLUME ["/var/log/eventmapper"]
# Run the command by default when the container starts.
ENTRYPOINT /go/bin/eventmapper /var/log/eventmapper/log.txt /go/src/eventmapper/config.yml

# Document that the service listens on port 8000.
EXPOSE 8000

# Environment variables
ENV GOGC 50
