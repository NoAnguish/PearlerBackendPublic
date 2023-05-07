resource "yandex_kubernetes_cluster" "pearler_cluster" {
  name        = "pearlerk8scluster"

  network_id = var.network_id

  master {
    version = "1.20"
    zonal {
      zone      = "ru-central1-b"
      subnet_id = var.subnet_ids.zone_b
    }

    public_ip = true

    maintenance_policy {
      auto_upgrade = true

      maintenance_window {
        start_time = "04:00"
        duration   = "2h"
      }
    }
  }

  service_account_id      = var.service_account_id
  node_service_account_id = yandex_iam_service_account.pearler_k8s_sa.id

  release_channel = "REGULAR"

  kms_provider {
    key_id = yandex_kms_symmetric_key.pearler_k8s_key.id
  }

  depends_on = [
    yandex_iam_service_account.pearler_k8s_sa,
  ]
}