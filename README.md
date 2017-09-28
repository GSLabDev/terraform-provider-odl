# Terraform Open Daylight Provider

This is the repository for the Terraform [Open Daylight][1] Provider, which one can use
with Terraform to work with Open Daylight.

[1]: https://www.opendaylight.org/

Coverage is currently only limited to Virtual Tennant Network, Virtual Bridge and Virtual Interface but in the coming months we are planning release coverage for most essential Open Daylight workflows.
Watch this space!

For general information about Terraform, visit the [official website][3] and the
[GitHub project page][4].

[3]: https://terraform.io/
[4]: https://github.com/hashicorp/terraform

# Using the Provider

The current version of this provider requires Terraform v0.10.2 or higher to
run.

Note that you need to run `terraform init` to fetch the provider before
deploying. Read about the provider split and other changes to TF v0.10.0 in the
official release announcement found [here][4].

[4]: https://www.hashicorp.com/blog/hashicorp-terraform-0-10/

## Full Provider Documentation

The provider is usefull in adding Virtual Tenant Network, Bridges and Interfaces using Open Daylight.

### Example
```hcl
# Configure the Open Daylight Provider
provider "odl" {
  server_ip     = "${var.odl_server_ip}"
  port          = "${var.odl_server_port}"
  user_name     = "${var.odl_username}"
  user_password = "${var.odl_password}"
}

# Add a Virtual Tennant Network
resource "odl_virtual_tenant_network" "firstVtn" {
  tenant_name  = "vtn12"
  operation    = "ADD"
  description  = "operation can be ADD or SET only"
  idle_timeout = 56
  hard_timeout = 58
}

# Add a Virtual Bridge
resource "odl_virtual_bridge" "firstVbr" {
  tenant_name  = "${odl_virtual_tenant_network.firstVtn.tenant_name}"
  bridge_name  = "vbr6"
  operation    = "SET"
  description  = "operation can be ADD or SET only"
  age_interval = 577
}

# Add a Virtual Interface
resource "odl_virtual_interface" "firstInterface" {
  tenant_name    = "${odl_virtual_tenant_network.firstVtn.tenant_name}"
  bridge_name    = "${odl_virtual_bridge.firstVbr.bridge_name}"
  description    = "operation can be ADD or SET only"
  interface_name = "interface1"
  enabled        = true
  terminal_name  = "ter1"
}
```

# Building The Provider

**NOTE:** Unless you are [developing][7] or require a pre-release bugfix or feature,
you will want to use the officially released version of the provider (see [the
section above][8]).

[7]: #developing-the-provider
[8]: #using-the-provider


## Cloning the Project

First, you will want to clone the repository to
`$GOPATH/src/github.com/terraform-providers/terraform-provider-odl`:

```sh
mkdir -p $GOPATH/src/github.com/terraform-providers
cd $GOPATH/src/github.com/terraform-providers
git clone git@github.com:terraform-providers/terraform-provider-odl
```

## Running the Build

After the clone has been completed, you can enter the provider directory and
build the provider.

```sh
cd $GOPATH/src/github.com/terraform-providers/terraform-provider-odl
make build
```

## Installing the Local Plugin

After the build is complete, copy the `terraform-provider-odl` binary into
the same path as your `terraform` binary, and re-run `terraform init`.

After this, your project-local `.terraform/plugins/ARCH/lock.json` (where `ARCH`
matches the architecture of your machine) file should contain a SHA256 sum that
matches the local plugin. Run `shasum -a 256` on the binary to verify the values
match.

# Developing the Provider

If you wish to work on the provider, you'll first need [Go][9] installed on your
machine (version 1.9+ is **required**). You'll also need to correctly setup a
[GOPATH][10], as well as adding `$GOPATH/bin` to your `$PATH`.

[9]: https://golang.org/
[10]: http://golang.org/doc/code.html#GOPATH

See [Building the Provider][11] for details on building the provider.

[11]: #building-the-provider

# Testing the Provider

**NOTE:** Testing the Open Daylight provider is currently a complex operation as it
requires having a Open Daylight Server to test against.

## Configuring Environment Variables

Most of the tests in this provider require a comprehensive list of environment
variables to run. See the individual `*_test.go` files in the
[`odl/`](odl/) directory for more details. The next section also
describes how you can manage a configuration file of the test environment
variables.

### Using the `.tf-odl-devrc.mk` file

The [`tf-odl-devrc.mk.example`](tf-odl-devrc.mk.example) file contains
an up-to-date list of environment variables required to run the acceptance
tests. Copy this to `$HOME/.tf-odl-devrc.mk` and change the permissions to
something more secure (ie: `chmod 600 $HOME/.tf-odl-devrc.mk`), and
configure the variables accordingly.

## Running the Acceptance Tests

After this is done, you can run the acceptance tests by running:

```sh
$ make testacc
```

If you want to run against a specific set of tests, run `make testacc` with the
`TESTARGS` parameter containing the run mask as per below:

```sh
make testacc TESTARGS="-run=TestAccOdl"
```

This following example would run all of the acceptance tests matching
`TestAccOdl`. Change this for the specific tests you want to
run.