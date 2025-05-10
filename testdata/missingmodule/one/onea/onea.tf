terraform {
  required_providers {
    coder = {
      source = "coder/coder"
      version = "2.4.0-pre0"
    }
  }
}

data "null_data_source" "values" {
  inputs = {
    foo = "bar"
  }
}
