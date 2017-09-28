---
layout: "odl"
page_title: "Open Daylight: odl_virtual_bridge"
sidebar_current: "docs-odl-resource-inventory-folder"
description: |-
  Provides a Open Daylight Virtual Bridge resource. This can be used to create and delete Virtual Bridge in Virtual Tenant Network.
---

# odl\_virtual\_bridge

Provides a Open Daylight Virtual Bridge resource. This can be used to create and delete Virtual Bridge in Virtual Tenant Network from Open Daylight.

## Example Usage

```hcl
resource "odl_virtual_bridge" "firstVbr" {
  tenant_name  = "${odl_virtual_tenant_network.firstVtn.tenant_name}"
  bridge_name  = "vbr6"
  operation    = "SET"
  description  = "operation can be ADD or SET only"
  age_interval = 577
}
```

## Argument Reference

The following arguments are supported:

* `tenant_name` - (Required) The Virtual Tenant Name to which we want to attach Virtual Bridge
* `bridge_name` - (Required) The Virtual Bridge Name
* `operation` - (Optional) The operation to be performed SET or ADD
* `description` - (Optional) The description of the Virtual Bridge
* `age_interval` - (Optional) The age interval of the Virtual Bridge