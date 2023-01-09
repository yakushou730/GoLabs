```shell
kubectl apply -f deployment.yaml
kubectl port-forward deploy/whoami-1 8089:80

# 查看 hostname
kubectl exec deploy/whoami-1 -- sh -c 'hostname' 
```
