# Kafka streams
Нужен для обработки данных в реальном времени.
Является промежуточным слоем между producer и consumer для обработки данных(преобразования, агрегации, фильтрации и т.д.).

```
+-------------------+                          +-------------------+                          +-------------------+
|                   |                          |                   |                          |                   |                   
|  Producer         |-----1.Send to----------->|  Kafka Streams    |-----2.Send message------>|  Kafka Consumer   |
|                   |       Topic 1            |                   |       to Topic 2         |                   |
+-------------------+                          +-------------------+                          +-------------------+
```


### Use cases for Kafka Streams
Можно использовать для:
* обработки данных в реальном времени
* преобразования данных
* агрегации данных
* преобразования данных из одного формата в другой
* преобразования данных из одной темы в другую


