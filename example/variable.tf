# ####

# Default region

variable "vscale_msk" {
  description = "vscale MSK data"
  default     = "msk0"
}

# Default Os

variable "vscale_centos_7" {
  description = "centos"
  default     = "centos_7_64_001_master"
}

# plans

variable "vscale_rplan" {
  type = "map"
  default = {
    "s"   = "small"
    "m"   = "medium"
    "l"   = "large"
    "xl"  = "huge"
    "xxl" = "monster"
  }
}

# dns record prefix

variable "devs" {

  type    = "list"

  default = ["dev1.vganin", "dev2.vganin"]

}
