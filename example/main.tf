provider "vscale" {
  # You need to set this in your .bashrc
  # export VSCALE_API_TOKEN="Your API TOKEN"
  #
}

# Create a new SSH key
resource "vscale_ssh_key" "user1" {
  name = "user1 key"
  key  = "ssh rsa ..."
}

# Create a new scalet
resource "vscale_scalet" "test" {
  count     = "${length(var.devs)}"
  ssh_keys  = ["${vscale_ssh_key.user1.id}"]
  hostname  = "${var.devs[count.index]}"
  make_from = "${var.vscale_centos_7}"
  location  = "${var.vscale_msk}"
  rplan     = "${var.vscale_rplan["s"]}"
  name      = "${var.devs[count.index]}"

  provisioner "remote-exec" {
    inline = [
      "sudo yum update",
      "hostnamectl set-hostname ${var.devs[count.index]}"
    ]
    connection {
      type     = "ssh"
      user     = "root"
      private_key = "${file("~/.ssh/id_rsa")}"
    }
  }
}

# Write credentials to file
resource "null_resource" "devstxt" {
  count   = "${length(var.devs)}"
  provisioner "local-exec" {
    command = "echo ${var.devs[count.index]} ${vscale_scalet.test.*.public_address[count.index]}) >> ./devs.txt"
  }
}
