variable "postgres_docker_image" {
  type = string
}

variable "namespace" {
  type = string
}

variable "storage_config" {
  type = object({
    pv_path     = string,
    pv_size     = string,
    match_label = string,
    pv_name     = string,
  })
}