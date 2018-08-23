# Jenkins Readme

Feel free to use these files to spin up jenkins. If you'd like to spin up a singleton on a provider here are the generic steps to do it.
## Bash commands
UBUNTU: 16.04
```
sudo apt-get update

sudo apt-get install openjdk-8-jre

sudo apt-get install jenkins docker.io -y

sudo groupmod -aG docker jenkins

sudo systemctl start jenkins

sudo systemctl start docker

# install goss/dgoss

curl -L https://github.com/aelsabbahy/goss/releases/download/v0.3.6/goss-linux-amd64 -o /usr/local/bin/goss

chmod +rx /usr/local/bin/goss

# (optional) dgoss docker wrapper (use 'master' for latest version)

curl -L https://raw.githubusercontent.com/aelsabbahy/goss/v0.3.6/extras/dgoss/dgoss -o /usr/local/bin/dgoss

chmod +rx /usr/local/bin/dgoss

# set ufw rules (add your public-client ipaddress here.)

sudo ufw allow from IPADDRESSHERE to any port 8080
sudo ufw allow from IPADDRESSHERE to any port 8084
sudo ufw allow from IPADDRESSHERE to any port 22

# allow hooks from github https://api.github.com/meta
sudo ufw allow from 192.30.252.0/22
sudo ufw allow from 185.199.108/22
sudo ufw allow from 140.82.112.0/22
sudo ufw start
```
## Login and Setup

Hit the default port 8080 and login.. it should be something like http://IPADDRESS:8080

Youll need the default pass via
`sudo cat /var/lib/jenkins/secrets/initialAdminPassword`

Once in install suggested plugins and setup the inital user. (remember this user/pass)

## Default port and Vi primer :)

After all that you'll need to change port its running on. by default its set to 8080 which may conflict with any testing steps during the pipeline run. lets change it to port 8084

Using Vi (no cheating!) Type exactly what is written here. Don't type or arrow once in file.
Note: if you aren't root user for some reason you'll have to run vi as `sudo` eg `sudo vi`

we will first take a backup of the file then edit.

```
cp /etc/defaultjenkins jenkins_bak
vi /etc/default/jenkins
```
type `/HTTP_PORT`

hit `enter`

using h j k or l move around (it's the old school wasd)

Once you find the line HTTP_PORT:8080, type `i` and you'll be in insert mode. here you can actually type and edit.

Once done with your editing hit `esc` which will throw you back in normal mode. from there edit type `:wq`

in Jenkins change your your to http://jenkins-ip-address:8080/restart

## Conclusion

That should get you a base setup. 
