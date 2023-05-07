resource kubernetes_manifest "pearler_cert_manager_cluster_issuer" {
  manifest   = yamldecode(file("${path.module}/manifests/cluster-issuer.yaml"))

  field_manager { force_conflicts = true }
}
