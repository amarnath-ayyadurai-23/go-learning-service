## We specify the base image we need for our
## go application
FROM golang:alpine
## We create an /app directory within our
## image that will hold our application source
## files
RUN apk add build-base git

RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
RUN adduser -S -D -H -h /app appuser
USER appuser

WORKDIR /app/http
## we run go build to compile the binary
## executable of our Go program
RUN go build -o apiserver .
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/http/apiserver"]