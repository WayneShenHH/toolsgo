# kubectl

### get token

```
$ kubectl config view
```

### setting config

```
$ gcloud container clusters get-credentials fsbs-k8s-1 --zone asia-east1-b --project cool-coral-208703
```

### redis-proxy

```
$ gcloud container clusters get-credentials fsbs-k8s-1 --zone asia-east1-b --project cool-coral-208703 \
 && kubectl port-forward --namespace msg $(kubectl get pod --namespace msg --selector="app=redis" --output jsonpath='{.items[0].metadata.name}') 6380:6379
```

### deploy

```
$ cd /Users/wayneshen/project/sbodds_document/doc-devops/kubernetes
$ kubectl apply -f dev/libgo.yml
```

### kubernetes dashboard

```
$ kubectl port-forward svc/kubernetes-dashboard 3001:443 -n=kube-system
```