resource "yandex_kubernetes_node_group" "pearler_node_group" {
  cluster_id  = "${yandex_kubernetes_cluster.pearler_cluster.id}"
  name        = "pearlernodegroup"
  version     = "1.20"

  instance_template {
    platform_id = "standard-v2"

    resources {
      cores         = 4
      memory        = 6
      core_fraction = 50
    }

    metadata = {
        "ssh-keys" = file("${path.module}/ssh-access")
    }

    network_interface {
      nat         = true
      subnet_ids  = [var.subnet_ids.zone_c]
    }

    boot_disk {
      type = "network-hdd"
      size = 64
    }

    container_runtime {
      type = "docker"
    }
  }

  scale_policy {
    fixed_scale {
      size = 2
    }
  }

  maintenance_policy {
    auto_upgrade = false
    auto_repair  = false
  }
}