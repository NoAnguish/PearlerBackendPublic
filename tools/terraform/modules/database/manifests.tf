locals {
  manifest_values = {
    postgres_docker_image = var.postgres_docker_image,
    namespace             = var.namespace,
  }
  pv_manifest_values = {
    namespace   = var.namespace,
    pv_path     = var.storage_config.pv_path,
    pv_size     = var.storage_config.pv_size,
    match_label = var.storage_config.match_label,
    pv_name     = var.storage_config.pv_name
  }
}

resource kubernetes_manifest "pearler_database_init_configmap" {
  manifest   = yamldecode(templatefile("${path.module}/manifests/init.configmap.yaml", local.manifest_values))
  field_manager { force_conflicts = true }
}

resource kubernetes_manifest "pearler_database_persistent_volume" {
  manifest   = yamldecode(templatefile("${path.module}/manifests/persistent-volume.yaml", local.pv_manifest_values))
  field_manager { force_conflicts = true }
  
  lifecycle {
    prevent_destroy = true 
  }
}

resource kubernetes_manifest "pearler_database_persistent_volume_claim" {
  manifest   = yamldecode(templatefile("${path.module}/manifests/persistent-volume-claim.yaml", local.pv_manifest_values))
  field_manager { force_conflicts = true }
}

resource kubernetes_manifest "pearler_database_deployment" {
  manifest   = yamldecode(templatefile("${path.module}/manifests/deployment.yaml", local.manifest_values))
  field_manager { force_conflicts = true }

  depends_on = [
    kubernetes_manifest.pearler_database_init_configmap,
    kubernetes_manifest.pearler_database_persistent_volume,
    kubernetes_manifest.pearler_database_persistent_volume_claim,
  ]
}

resource kubernetes_manifest "pearler_database_service" {
  manifest   = yamldecode(templatefile("${path.module}/manifests/service.yaml", local.manifest_values))
  field_manager { force_conflicts = true }

  depends_on = [
    kubernetes_manifest.pearler_database_deployment,
  ]
}