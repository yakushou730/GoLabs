SHELL := /bin/bash

# docker top <container-id>
# 查看對應 container 內的 process

# docker logs <container-id>
# 查看對應 container 的 log

# docker inspect <container-id>
# 查看對應 container 的詳細資訊

# docker stats <container-id>
# 查看對應 container 使用的硬體資源

# docker rm -f <container-id>
# 移除對應的 container (-f 表示強制)

# docker cp [local file path] <container-id>:[container file path]
# copy [local file file] 到 container 內的指定路徑 [container file path]

hello:
	docker run diamol/ch02-hello-diamol

connect:
# --interactive -> -i (告訴 Docker 我們要設定對 container 內的連線)
# --tty         -> -t (告訴 Docker 我們要連線到 container 內的 terminal session)
# 表示 互動式 虛擬終端介面
	docker run -it diamol/base

ps:
# list active containers
	docker ps

ps-a:
# list all containers including Exited
	docker ps -a

run-web:
# --detach  -> -d ( 讓 container 在背景執行 )
# --publish -> -p ( 開 container 對應的 port 給 host) -> [host-port]:[container-port]
	docker run -d -p 8088:80 diamol/ch02-hello-diamol-web

