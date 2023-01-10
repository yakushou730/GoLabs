SHELL := /bin/bash

build-go-app:
	docker build -t image-gallery .

run-go-app-with-network-nat:
	docker run -d -p 802:80 --network nat image-gallery

compare-image-reference:
	docker image ls -f reference=access-log -f reference=image-gallery
