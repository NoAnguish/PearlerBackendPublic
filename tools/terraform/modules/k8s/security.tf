resource "yandex_kms_symmetric_key" "pearler_k8s_key" {
  name              = "PearlerK8sKey"
  default_algorithm = "AES_128"
  rotation_period   = "2190h" // equal to 3 months
}

resource "yandex_iam_service_account" "pearler_k8s_sa" {
  name = var.sa_name // camel case is not available
}

resource "yandex_resourcemanager_folder_iam_binding" "pearler_k8s_access_to_registry" {
  folder_id = var.folder_id
  role      = "container-registry.images.puller"
  members   = [
    "serviceAccount:${yandex_iam_service_account.pearler_k8s_sa.id}"
  ]
}