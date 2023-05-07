resource "kubernetes_namespace" "pearler" {
  metadata { name = var.namespace }
}