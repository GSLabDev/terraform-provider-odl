provider "odl"{
        server_ip = "192.168.56.106"
        port = 8080
        user_name = "admin"
        user_password = "admin"
}

resource "odl_vtn" "firstVtn"{
        tenant_name = "vtn5"
}

resource "odl_vbr" "firstVbr"{
        tenant_name = "${odl_vtn.firstVtn.tenant_name}"
        bridge_name = "vbr6"
}
