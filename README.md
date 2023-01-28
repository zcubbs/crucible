# Crucible

An all-in-one CLI for managing your deployment control node with AWX. 
Enables install of a single node K3s cluster, installs helm and AWX-Operator. Ones the operator is ready it depoloy one or more AWX instances. The auto configures the 
AWX templates, credentials, and inventories through yaml config. The cli can also run job templates.

---
![](_assets/crucible.png)

---

## CLI

```cmd
 _____ ______ _   _ _____ ___________ _      _____ 
/  __ \| ___ \ | | /  __ \_   _| ___ \ |    |  ___|
| /  \/| |_/ / | | | /  \/ | | | |_/ / |    | |__  
| |    |    /| | | | |     | | | ___ \ |    |  __|
| \__/\| |\ \| |_| | \__/\_| |_| |_/ / |____| |___
 \____/\_| \_|\___/ \____/\___/\____/\_____/\____/

> crucible -h

Available Commands:

  about       Print the info about crucible-cli
  completion  Generate the autocompletion script for the specified shell
  config      list cli configuration
  helm        Helm Helper Commands
  help        Help about any command
  info        info is a palette that contains system info commands
  k3s         K3s Helper Commands
  os          OS Helper Commands

```

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
