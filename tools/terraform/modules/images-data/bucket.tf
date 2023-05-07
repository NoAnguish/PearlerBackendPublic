resource "yandex_iam_service_account_static_access_key" "sa_static_key" {
  service_account_id = var.deploy_sa_id
  description        = "static access key for creating pearler images object storage"
}

resource "yandex_storage_bucket" "image_bucket" {
  access_key = yandex_iam_service_account_static_access_key.sa_static_key.access_key
  secret_key = yandex_iam_service_account_static_access_key.sa_static_key.secret_key

  bucket = var.bucket_name
  acl    = "public-read"
}