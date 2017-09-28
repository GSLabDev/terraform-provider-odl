---
layout: "odl"
page_title: "Open Daylight: odl_virtual_tenant_network"
sidebar_current: "docs-odl-resource-inventory-folder"
description: |-
  Provides a Open Daylight Virtual Tenant Network resource. This can be used to create and delete Virtual Tenant Network.
---

# odl\_virtual\_tenant\_network

Provides a Open Daylight Virtual Tenant Network resource. This can be used to create and delete Virtual Tenant Network from Open Daylight.

## Example Usage

```hcl
resource "odl_virtual_tenant_network" "firstVtn" {
  tenant_name  = "vtn12"
  operation    = "ADD"
  description  = "operation can be ADD or SET only"
  idle_timeout = 56
  hard_timeout = 58
}
```

## Argument Reference

The following arguments are supported:

* `tenant_name` - (Required) The Tenant Name of the Virtual Network
* `operation` - (Optional) The operation to be performed SET or ADD
* `description` - (Optional) The description of the network
* `idle_timeout` - (Optional) The idle timeout of the network
* `hard_timeout` - (Optional) The hard timeout of the network