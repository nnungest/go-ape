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

## Build and run via Docker

To run this via its supplied Dockerfile instead:
(after cloning and working from root of directory)

$ docker build . -t go-ape/latest

$ docker run -p 8080:8080 go-ape

## Testing Via goss/dgoss
Included is the goss.yaml file. Run it against the container via:

$ dgoss run -p 8080:8080 go-ape/latest

This checks that the app is running, and that endpoints are being returned thereby testing liveness on port 8080.

## Jenkins
Included is the jenkinsfile to build and push image to dockerhub.
