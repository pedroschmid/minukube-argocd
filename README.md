![ArgoCD](https://i0.wp.com/rodrigolira.eti.br/wp-content/uploads/2022/03/argoprojio-ar21.png?fit=1200%2C600&ssl=1)
# ArgoCD 

### Install ArgoCD using Helm
```bash
helm repo add argo https://argoproj.github.io/argo-helm
helm repo update
helm install argo-cd argo/argo-cd --create-namespace --namespace argocd --values resources/argocd/values.yaml
```

### Retrieve ArgoCD first password
```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

### Port foward ArgoCD to 8080
```bash
kubectl port-forward service/argo-cd-argocd-server -n argocd 8080:443
```

### Apply app of apps (application that load all other applications stored in config folder)
```bash
kubectl apply -n argocd -f resources/app-of-apps.yaml
```

### Expose kong to http (Enable http protocol)
```bash
kubectl -n argocd patch deployment argo-cd-argocd-server --type json \
    -p='[ { "op": "replace", "path":"/spec/template/spec/containers/0/command","value": ["argocd-server","--staticassets","/shared/app","--repo-server","argo-cd-argocd-repo-server:8081","--dex-server","http://argo-cd-argocd-dex-server:5556","--logformat","text","--loglevel","info","--redis","argo-cd-argocd-redis:6379","--insecure"] }]'
```
-----

# Kong API

### Deploy Prometheus plugin (Namespace not required because its a global plugin)
```bash
kubectl apply -f resource/plugins/kong/prometheus.yaml
```

### Deploy kong rate limiting plugin (In echo namespace)
```bash
kubectl apply -n echo -f resources/kong/plugins/rate-limiting.yaml
```
-----
# License
[MIT](https://choosealicense.com/licenses/mit/)