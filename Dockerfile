FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/gramma-trainer-server/
COPY . .

RUN go get ./...
RUN go generate
RUN go build -o /go/bin/gramma-trainer-server

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/gramma-trainer-server /go/bin/gramma-trainer-server
# Run the hello binary.
ENTRYPOINT ["/go/bin/gramma-trainer-server"]

