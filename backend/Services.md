sudo snap install microk8s --classic
microk8s enable dashboard dns registry istio jaeger prometheus helm3 ingress metrics-server prometheus registry storage

kubectl get svc -n istio-system


