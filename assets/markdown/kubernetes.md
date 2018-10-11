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

### kubernetes dashboard token

```
$ kubectl -n kube-system get secret | grep dashboard-token
$ kubectl -n kube-system describe secret 'chosen secret name from step 1'
```

### kubernetes config template

```yml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    server: https://127.0.0.1
  name: gcp_cluster_name
contexts:
- context:
    cluster: gcp_cluster_name
    user: gcp_cluster_name
  name: gcp_cluster_name
current-context: gcp_cluster_name
kind: Config
preferences: {}
users:
- name: gcp_cluster_name
  user:
    token:
      xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    auth-provider:
      config:
        access-token: xxxxxxxxxxxxxxxx
        cmd-args: config config-helper --format=json
        cmd-path: /Users/wayneshen/googlecloud/google-cloud-sdk/bin/gcloud
        expiry: 2018-09-04T07:00:43Z
        expiry-key: '{.credential.token_expiry}'
        token-key: '{.credential.access_token}'
      name: gcp
```