## We specify the base image we need for our
## go application
FROM golang:1.16.15-alpine
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /http
## We copy everything in the root directory
## into our /app directory
ADD http/ /http
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /http
## we run go build to compile the binary
## executable of our Go program
RUN go build -o apiserver .
## Our start command which kicks off
## our newly created binary executable
CMD ["/http/apiserver"]