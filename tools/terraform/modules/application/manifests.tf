locals {
  deployment_manifest_values = {
    docker_image_url = var.docker_image_url,
    namespace        = var.namespace,
  }
  default_manifest_values = {
    namespace = var.namespace,
  }
  ingress_manifest_values = {
    api_host  = var.api_host,
    namespace = var.namespace,
  }
  configmap_manifest_values = {
    image_bucket_name  = var.image_bucket_name,
    namespace          = var.namespace,
    access_key_id      = var.access_key_id,
    secret_access_key  = var.secret_access_key
  }
}

resource kubernetes_manifest "pearler_application_deployment" {
  manifest   = yamldecode(templatefile("${path.module}/manifests/deployment.yaml", local.deployment_manifest_values))
  field_manager { force_conflicts = true }

  depends_on = [
    kubernetes_manifest.pearler_application_configmap,
  ]
}

resource kubernetes_manifest "pearler_application_configmap" {
  manifest   = yamldecode(templatefile("${path.module}/manifests/configmap.yaml", local.configmap_manifest_values))
  field_manager { force_conflicts = true }
}

resource kubernetes_manifest "pearler_application_service" {
  manifest   = yamldecode(templatefile("${path.module}/manifests/service.yaml", local.default_manifest_values))
  field_manager { force_conflicts = true }
}

resource kubernetes_manifest "pearler_ingress" {
  manifest   = yamldecode(templatefile("${path.module}/manifests/ingress.yaml", local.ingress_manifest_values))
  field_manager { force_conflicts = true }

  depends_on = [
    kubernetes_manifest.pearler_application_service,
    kubernetes_manifest.pearler_application_deployment
  ]
}