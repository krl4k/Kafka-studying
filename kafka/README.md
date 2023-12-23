# Kafka

Pull-model - когда клиент сам забирает данные из брокера. Kafka ничего не отправляет сама.

# link

    https://dzone.com/articles/how-to-deploy-apache-kafka-with-kubernetes


```bash

kubectl apply -f 00-namespace.yaml
kubectl apply -f 01-zookeper.yaml
kubectl apply -f 02-kafka.yaml

# открыть порт для доступа к kafka извне
kubectl port-forward kafka-broker-d4fb477fd-7kwmw 9092 -n kafka 

# писать в топик test
kcat -P -b localhost:9092 -t test

# читать из топика test
kcat -C -b localhost:9092 -t test

```


### Создание топика и партиций
Лучше создавать минимум 4 партиции на сервис, чтобы можно было масштабировать сервис.


```bash
# создать топик test с 3 партициями
kafka-topics.sh --create --topic test --partitions 3 --replication-factor 1 --bootstrap-server localhost:9092
```    