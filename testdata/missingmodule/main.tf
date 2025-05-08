module "does-not-exist" {
  source         = "registry.coder.com/modules/does-not-exist/coder"
}

module "does-not-exist-2" {
  count = 0
  source         = "registry.coder.com/modules/does-not-exist/coder"
}

module "one" {
  source = "./one"
}