FROM jenkins/jenkins:lts
USER root
# from github.com/goss
# See https://github.com/aelsabbahy/goss/releases for release versions
RUN curl -L https://github.com/aelsabbahy/goss/releases/download/v0.3.6/goss-linux-amd64 -o /usr/local/bin/goss
RUN chmod +rx /usr/local/bin/goss

# (optional) dgoss docker wrapper (use 'master' for latest version)
RUN curl -L https://raw.githubusercontent.com/aelsabbahy/goss/v0.3.6/extras/dgoss/dgoss -o /usr/local/bin/dgoss
RUN chmod +rx /usr/local/bin/dgoss
RUN apt-get update && apt-get upgrade docker -y
USER jenkins
