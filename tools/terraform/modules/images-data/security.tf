resource "yandex_resourcemanager_folder_iam_binding" "pearler_k8s_sa_access_to_s3_upload" {
  folder_id = var.folder_id
  role      = "storage.uploader"
  members   = [
    "serviceAccount:${var.k8s_sa_id}"
  ]
}