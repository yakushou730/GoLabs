SHELL := /bin/bash

pull-web-ping:
	docker pull diamol/ch03-web-ping

run-web-ping:
# --name  	     ( name of the container )
# --env    -> -e ( environment value )
	docker run -d --name web-ping -e TARGET=google.com diamol/ch03-web-ping

logs-web-pig:
	docker logs web-ping

rm-web-ping:
	docker rm -f web-ping

image-ls:
# 這裡的 w* 是指對 image 列表做 w* 的 regexp 過濾
	docker image ls w*

run-local-web-ping:
	docker run -e TARGET=docker.com web-ping

history-web-ping:
# 查看這個 image build 的歷史紀錄
	docker history web-ping

system-df:
# 查看 Docker 用到的硬碟空間
	docker system df
