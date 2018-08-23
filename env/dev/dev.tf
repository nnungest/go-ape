# set token in .tfvars file.

provider "digitalocean" {
  token = "${var.do_token}"
}

resource "digitalocean_droplet" "k8s-master" {
  count = 1
  image  = "ubuntu-18-04-x64"
  name   = "k8s-master"
  region = "nyc1"
  size   = "4GB"
  ssh_keys = ["${digitalocean_ssh_key.ssh.id}"]
}

resource "digitalocean_droplet" "k8s-worker" {
  count = 1
  image  = "ubuntu-18-04-x64"
  name   = "k8s-worker"
  region = "nyc1"
  size   = "4GB"
  ssh_keys = ["${digitalocean_ssh_key.ssh.id}"]
} 
