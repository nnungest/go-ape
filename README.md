# Go Ape
A basic app for testing a CI CD pipeline.

**Endpoints**

GET "/" returns "hello world"
GET "/hello" returns "hello"
GET /world" returns "world"

## Running
### Option 1: Running from pre-built
A prebuilt binary is included as "go-ape" in this repo. Simply run that and access it over http: $ go-ape && curl localhost:8080/

### Option 2: Building the go Application from source
Assuming go is installed:
$ git clone https://github.com/nnungest/go-ape
$ cd go-ape
$ go build
$ go-ape
$ curl localhost:8080/ENDPOINT


TODO: endpoint string manipulation

## Running via Docker

To run this via its supplied Dockerfile instead:
(after cloning and working from root of directory)

$ docker build - < Dockerfile
