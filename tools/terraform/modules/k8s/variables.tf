variable "network_id" {
  type = string
}

variable "subnet_ids" {
  type = object({
      zone_a = string,
      zone_b = string,
      zone_c = string,
  })
}

variable "service_account_id" {
  type = string
}

variable "folder_id" {
  type = string
}

variable "sa_name" {
  type = string
}