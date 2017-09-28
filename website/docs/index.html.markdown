---
layout: "odl"
page_title: "Provider: Open Daylight"
sidebar_current: "docs-odl-index"
description: |-
  The Open Daylight provider is used to interact with the resources supported by
  Open Daylight. The provider needs to be configured with the proper credentials
  before it can be used.
---

# Open Daylight Provider

The Open Daylight provider is used to interact with the resources supported by
Open Daylight.
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

~> **NOTE:** The Open Daylight Provider currently represents _initial support_
and therefore may undergo significant changes as the community improves it. This
provider at this time only supports adding Virtual Tenant Network, Virtual Bridge and Virtual Interface Resource

## Example Usage

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

## Argument Reference

The following arguments are used to configure the Active Directory Provider:

* `user_name` - (Required) This is the username for Open Daylight operations. Can also
  be specified with the `ODL_SERVER_USER` environment variable.
* `user_password` - (Required) This is the password for Open Daylight API operations. Can
  also be specified with the `ODL_SERVER_PASSWORD` environment variable.
* `server_ip` - (Required) This is the Open Daylight server ip for Open Daylight Api
  operations. Can also be specified with the `ODL_SERVER_IP` environment
  variable.
* `port` - (Required) This is the port for API operations of the Open Daylight.

## Acceptance Tests

The Active Directory provider's acceptance tests require the above provider
configuration fields to be set using the documented environment variables.

Once all these variables are in place, the tests can be run like this:

```
make testacc TEST=./builtin/providers/odl
```
