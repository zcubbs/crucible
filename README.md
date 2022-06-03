# Crucible

An all-in-one tool for managing your air-gaped deployments with AWX.

---
![](_assets/crucible.png)

---
## Local Dev

### Install AWX on local cluster

```powershell
helm repo add awx-operator https://ansible.github.io/awx-operator/
helm upgrade --install awx-operator awx-operator/awx-operator -n awx --create-namespace
kubectl create ns awx-dev
kubectl apply -f _dev/awx/awx-dev.yaml
```

To reset awx password:
```powershell
kubectl -n awx exec -it awx-dev-<POD_ID> -c awx-dev-web -- awx-manage changepassword admin
```
