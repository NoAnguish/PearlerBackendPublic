resource "yandex_container_registry" "pearler_registry" {
  name      = "pearler-registry"
  folder_id = var.folder_id
}