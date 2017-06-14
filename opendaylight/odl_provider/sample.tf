provider "odl" {
	user = "admin"
  	password = "admin"
  	ip = "192.168.56.102"
  	port = "8181"

    
}

resource "odl_networkTopology" "foo"{
  user = "admin"
  password = "admin"
  ip = "192.168.56.102"
  port = "8181"
  resturl = "/config/network-topology:network-topology/"
}
