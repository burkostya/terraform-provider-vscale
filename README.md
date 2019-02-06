Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Building The Provider
---------------------

* export PROVIDER_NAME=vscale
* mkdir -p $GOPATH/src/github.com/terraform-providers
* cd $GOPATH/src/github.com/terraform-providers
* git clone https://github.com/vganyn/terraform-provider-vscale
* cd $GOPATH/src/github.com/terraform-providers/terraform-provider-$PROVIDER_NAME
* go get
* go build
* mkdir -p ~/.terraform.d/plugins/
* mv terraform-provider-$PROVIDER_NAME ~/.terraform.d/plugins/

Using the provider
----------------------

See the [DigitalOcean Provider documentation](https://www.terraform.io/docs/providers/do/index.html) to get started using the DigitalOcean provider.
