FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/api/
COPY . .

RUN go get -d -v
RUN go generate ./...
#RUN go build -o /go/bin/api
#RUN vgo build
#RUN chmod 777 /go/bin/api
#RUN go run github.com/99designs/gqlgen
#RUN go run server/server.go -v
# ENTRYPOINT ["/go/bin/api"]

############################
# STEP 2 build a small image
############################
#FROM scratch
# Copy our static executable.
#COPY --from=builder /go/bin/api /go/bin/api
# Run the hello binary.
#RUN ["chmod", "+x", ""]
#RUN chmod 777 /go/bin/api



