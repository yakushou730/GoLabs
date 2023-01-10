SHELL := /bin/bash

build-java-app:
	docker build -t java-app .

network-create-nat:
	docker network create nat

run-java-app-with-network-nat:
	docker run --name iotd -d -p 800:80 --network nat java-app
