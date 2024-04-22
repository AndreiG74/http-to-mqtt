# http to mqtt
 The script starts an http server, which receives requests and publishes their parameters to the mqtt server

The script starts an http server, which receives requests and publishes their parameters to the mqtt server:
When sending an HTTP POST request to http://ip.host:8080/sendmqtt with the topic and message parameters,
the data will be extracted from the request and sent to the MQTT server.
You should replace "mqtt.broker.address:1883" with your MQTT broker's address.

Скрипт запускает http сервер, на который принимает запросы, и публикует их параметры на mqtt  сервер:
при отправке HTTP POST-запроса на http://ip.хоста :8080/sendmqtt с параметрами topic и message,
данные будут извлечены из запроса и отправлены на MQTT-сервер. 
Вы должны заменить "mqtt.broker.address:1883" на адрес вашего MQTT-брокера.	