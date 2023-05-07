resource "helm_release" "cert_manager" {
  name = "pearler-cert-manager"

  repository = "https://charts.jetstack.io"
  chart      = "cert-manager"
  version    = "1.10.1"

  set {
    name  = "installCRDs"
    value = "true"
  }
  
  cleanup_on_fail  = true
  create_namespace = true 
  namespace        = "cert-manager" 
  
  wait          = true
  wait_for_jobs = true  
}
