SHELL := /bin/bash

run-hello-kiamol:
# --restart -> 指重新啟動的參數
# --image   -> 指定要使用哪個 image
	kubectl run hello-kiamol --image=kiamol/ch02-hello-kiamol --restart=Never

wait-hello-kiamol:
# 等待 hello-kiamol 的 condition 變成 Ready
	kubectl wait --for=condition=Ready pod hello-kiamol

get-pods:
# 列出這個 cluster 上所有的 pods
	kubectl get pods

get-hello-kiamol:
	kubectl get pod hello-kiamol

get-custom-hello-kiamol:
# --output -> -o ( 印出資訊 )
# 印出自訂欄位訊息
	kubectl get pod hello-kiamol --output custom-columns=NAME:metadata.name,NODE_IP:status.hostIP,POD_IP:status.podIP

get-jsonp-hello-kiamol:
# 以 json 路徑印出對應的資訊
	kubectl get pod hello-kiamol -o jsonpath='{.status.containerStatuses[0].containerID}'

describe-hello-kiamol:
# 顯示對應的 pod 的詳細資訊
	kubectl describe pod hello-kiamol

forward-port-to-hello-kiamol:
# 暫時開通連接 port 的對應 [host local port]:[pod port]
	kubectl port-forward pod/hello-kiamol 8080:80

forward-port-to-deployment-hello-kiamol:
	kubectl port-forward deploy/hello-kiamol-2 8080:80

create-deploy-hello-kiamol:
	kubectl create deployment hello-kiamol-2 --image=kiamol/ch02-hello-kiamol

get-deploy-hello-kiamol-label:
	kubectl get deploy hello-kiamol-2 -o jsonpath='{.spec.template.metadata.labels}'

get-pods-with-label:
# --selector -> -l ( 表示標籤選擇器 )
	kubectl get pods -l app-hello-kiamol-2

overwrite-label:
	kubectl label pods -l app=hello-kiamol-2 --overwrite app=hello-kiamol-x

run-interactive-shell-in-container:
	kubectl exec -it hello-kiamol-3 -- sh
#   wget -O - http://localhost | head -n 4

logs-hello-kiamol:
# --tail ( 倒數 N 行 )
	kubectl logs --tail=5 hello-kiamol-3

exec-by-deploy:
# deployment would run those command in pod
	kubectl exec deploy/hello-kiamol-4 -- sh -c 'wget -O - http://localhost > /dev/null'

cp-file-from-container-to-local:
# 把 container 內的 usr/share/nginx/html/index.html 複製到本機 tmp/index.html
# kubectl cp [source] [destination]
	kubectl cp hello-kiamol-3:usr/share/nginx/html/index.html tmp/index.html

delete-pod-hello-kiamol:
	kubectl delete pod hello-kiamol-3

delete-deploy-hello-kiamol:
	kubectl delete deploy hello-kiamol-2

get-all:
	kubectl get all
