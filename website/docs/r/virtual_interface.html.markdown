---
layout: "odl"
page_title: "Open Daylight: odl_virtual_interface"
sidebar_current: "docs-odl-resource-inventory-folder"
description: |-
  Provides a Open Daylight Virtual Interface resource. This can be used to create and delete Virtual Interface in Virtual Bridge of a Virtual Tenant Network.
---

# odl\_virtual\_interface

Provides a Open Daylight Virtual Interface resource. This can be used to create and delete Virtual Interface in Virtual Bridge of a Virtual Tenant Network

## Example Usage

```hcl
resource "odl_virtual_interface" "firstInterface" {
  tenant_name    = "${odl_virtual_tenant_network.firstVtn.tenant_name}"
  bridge_name    = "${odl_virtual_bridge.firstVbr.bridge_name}"
  interface_name = "interface1"
  description    = "operation can be ADD or SET only"
  enabled        = true
  terminal_name  = "ter1"
}
```

## Argument Reference

The following arguments are supported:

* `tenant_name` - (Required) The Tenant Name of the Virtual Tenant Network to which Virtual Interface needs to be attached
* `bridge_name` - (Required) The Virtual Bridge Name which is in Virtual Tenant Network and needs to be attached Virtual Interface
* `interface_name` - (Required) The Virtual Interface Name
* `operation` - (Optional) The operation to be performed SET or ADD
* `description` - (Optional) The description of the Virtual Interface
* `enabled` - (Optional) Virtual Interface is enabled
* `terminal_name` - (Optional) Name of the terminal