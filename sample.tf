provider "odl"{
        server_ip = "192.168.56.106"
        port = 8080
        user_name = "admin"
        user_password = "admin"
}

resource "odl_vtn" "firstVtn"{
        tenant_name = "vtn5"
        operation = "ADD"
        description = "operation can be ADD or SET only"
        idle_timeout = 56
        hard_timeout = 67
}

resource "odl_vbr" "firstVbr"{
        tenant_name = "${odl_vtn.firstVtn.tenant_name}"
        bridge_name = "vbr6"
        operation = "SET"
        description = "operation can be ADD or SET only"
        age_interval = 577
}

resource "odl_vinterface" "firstInterface"{
        tenant_name = "${odl_vtn.firstVtn.tenant_name}"
        bridge_name = "${odl_vbr.firstVbr.bridge_name}"
        description = "operation can be ADD or SET only"
        interface_name = "interface1"
        enabled = true
        terminal_name = "ter1"
}
