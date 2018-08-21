# Go Ape
A basic app for testing a CI CD pipeline.

**Endpoints**

GET "/" returns "hello world"

GET "/uppercase=true" (or false) returns "HELLO WORLD" or "hello world" , respecitvely.

GET "/reverse=true" (or false) returens "dlrow olleh" or "hellow world"

GET "/hello" returns "hello"

GET /world" returns "world"

## Running
### Option 1: Running from pre-built
A prebuilt binary is included as "go-ape" in this repo. Simply run that and access it over http: $ ./go-ape && curl localhost:8080/

### Option 2: Building the go Application from source

Assuming go is installed:

$ git clone https://github.com/nnungest/go-ape

$ cd go-ape

$ go build

$ ./go-ape

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

Included is the jenkinsfile to build and push image to dockerhub. To get dgoss working in Jenkins one needs to install it on the server. This can be accomplished by running the dockerfile in jenkins/Dockerfile. The tests are buiiltin to the Jenkinsfile for the pipeline so nothing needs done there. 

The jenkins portion can be run via: ```docker build -t jenkins/goss && docker run -p 8084:8080 -p 50000:50000 jenkins/jenkins:lts ```  

Here we are specifying we run jenkins over port 8084 on our local@ localhost:8084 so that we aren't colliding with any local development on port 8080. Fire up a browser and head over to jenkins!

### Jenkins pipeline setup
  Jenkins needs to have credentials set under: Manage Jenkins > credentials > system > global credentials 
   Here you can add bitbucket and github login credentials of the username/password variety.

For you github api keys you'll create those via  localhost:8084/configure ... scroll down to github section and click the jenkins dropdown next to key option. You'll need a github api key for giving jenkins permission to read/write on the repo, which you can get from your github page. Enter this as a "secret text". once done make sure you select it from the key dropdown, click "test" and if all is good, go ahead and save the page. 
