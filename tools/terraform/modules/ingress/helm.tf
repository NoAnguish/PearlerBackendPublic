resource "helm_release" "nginx_ingress" {
  name = "pearler-ingress"

  repository = "https://kubernetes.github.io/ingress-nginx"
  chart      = "ingress-nginx"
  version    = "4.2.5"
  
  cleanup_on_fail  = true
  create_namespace = true 
  namespace        = "pearler-ingress" 
  
  wait          = true
  wait_for_jobs = true  
}