SHELL := /bin/bash

build-js-app:
	docker build -t access-log .

run-js-app-with-network-nat:
	docker run --name accesslog -d -p 801:80 --network nat access-log
