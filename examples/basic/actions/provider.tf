terraform {
  required_providers {
    datapower = {
      source = "scottw514/datapower"
    }
  }
}

provider "datapower" {
  hostname = "datapower-dev"
  username = "admin"
  password = "admin"
}
