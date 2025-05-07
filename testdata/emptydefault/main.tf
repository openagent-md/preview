terraform {
  required_providers {
    coder = {
      source = "coder/coder"
      version = "2.4.0-pre0"
    }
  }
}

data "coder_parameter" "word" {
  name        = "word"
  description = "Select something"
  type        = "string"
  order       = 1
  # No default selected

  option {
    name  = "Bird"
    value = "bird"
    description = "An animal that can fly."
  }
  option {
    name  = "Boat"
    value = "boat"
  }
}
