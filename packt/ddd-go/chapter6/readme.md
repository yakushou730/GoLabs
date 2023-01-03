## Scene
```
You work for a travel comparison website. Your team is responsible for making recommendations on where 
a customer might be able to travel, given their budget and other factors. Your team is known internally 
as the recommendations team. Your team has been asked to expose your recommendations via an API 
so that other teams in the company may use it to build their own products.

There is another team in your company that is responsible for working with travel providers to onboard them
and aggregate their costs and offer information to your system. They are known as the partnership team.
```

using the makefile to run docker compose

then the partnership service will be running on port 3031
```shell
make docker-compose-up
```

API example for partnership service
```shell
curl --location --request GET 'http://localhost:3031/partnerships?location=UK'
```
