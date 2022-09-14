#!/bin/bash
set -x
### Install k3s
### Given permissions
# permissions 644 are given to /etc/rancher/k3s/k3s.yaml
curl -sfL https://get.k3s.io | sh -s - --write-kubeconfig-mode 644
### Confirm all is good and running
kubectl version --short
kubectl get nodes
### Install prerequisites for awx
sudo apt update
sudo apt install git build-essential -y
### Clone the operator install
git clone https://github.com/ansible/awx-operator.git
export NAMESPACE=awx
kubectl create ns ${NAMESPACE}
sudo kubectl config set-context --current --namespace=$NAMESPACE
cd awx-operator/
sudo apt install curl jq -y
RELEASE_TAG=`curl -s https://api.github.com/repos/ansible/awx-operator/releases/latest | grep tag_name | cut -d '"' -f 4`
echo $RELEASE_TAG
git checkout $RELEASE_TAG
make deploy
### Install Ansible AWX on Ubuntu20.04 using operator
#### Create Static data PVC
cat <<EOF | kubectl create -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: static-data-pvc
  namespace: awx
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: local-path
  resources:
    requests:
      storage: 5Gi
EOF
### Check PVC state
kubectl get pvc -n awx
#### paste the following content at the end of awx-deploy.yaml
cat <<EOF > awx-deploy.yml
---
apiVersion: awx.ansible.com/v1beta1
kind: AWX
metadata:
  name: awx
spec:
  service_type: nodeport
  projects_persistence: true
  projects_storage_access_mode: ReadWriteOnce
  web_extra_volume_mounts: |
    - name: static-data
      mountPath: /var/lib/projects
  extra_volumes: |
    - name: static-data
      persistentVolumeClaim:
        claimName: static-data-pvc
EOF
#### apply configuration
kubectl apply -f awx-deploy.yml
########### How to get port on which it's reachable
kubectl get svc -n awx
########## get admin credentials
kubectl get secret awx-admin-password -o jsonpath="{.data.password}" | base64 --decode
echo "Open a web browser and connect to the dashboard with creds admin/passwd_retrieved"
