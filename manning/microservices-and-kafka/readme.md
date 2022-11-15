1. local 端 安裝 Kafka

https://www.conduktor.io/kafka/how-to-install-apache-kafka-on-mac-with-homebrew
https://medium.com/@at_ishikawa/getting-started-with-kafka-on-mac-f6aa8924fcda

```
CLI

- 建立 topic
kafka-topics --bootstrap-server localhost:9092 --create --topic Orders

-- 列出 topics
kafka-topics --bootstrap-server localhost:9092 --list

- 發送訊息
kafka-console-producer --bootstrap-server localhost:9092 --topic Orders

- 接收訊息
kafka-console-consumer --bootstrap-server localhost:9092 --topic Orders

- 刪除 topic
kafka-topics --bootstrap-server localhost:9092 --delete --topic Orders
```

2. 建立 topics
```
kafka-topics --bootstrap-server localhost:9092 --create --topic Orders --config max.message.bytes=10485760
kafka-topics --bootstrap-server localhost:9092 --create --topic Notifications --config max.message.bytes=10485760

或 (先建立再調設定)
kafka-topics --bootstrap-server localhost:9092 --create --topic Orders
kafka-topics --bootstrap-server localhost:9092 --entity-type topics --entity-name Orders --alter --add-config max.message.bytes=10485760
```
