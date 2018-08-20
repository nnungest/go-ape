FROM centos

RUN yum -y update && yum clean all

RUN mkdir -p /go && chmod -R 777 /go && \
    yum -y install git golang && yum clean all

ENV GOPATH /go

WORKDIR /go/src/go-ape


COPY main.go /go/src/go-ape/

EXPOSE 8080

RUN go build

ENTRYPOINT ["./go-ape"]
