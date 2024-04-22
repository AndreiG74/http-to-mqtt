package main

import (
	"fmt"
	"log"
	"net/http"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient MQTT.Client

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	// Анализируем параметры запроса
	r.ParseForm()
	topic := r.FormValue("topic")
	message := r.FormValue("message")

	if topic == "" || message == "" {
		http.Error(w, "Не хватает параметров запроса", http.StatusBadRequest)
		return
	}

	// Отправляем сообщение на MQTT-сервер
	token := mqttClient.Publish(topic, 0, false, message)
	token.Wait()

	// Отправляем ответ клиенту
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Сообщение успешно отправлено на тему %s: %s", topic, message)
}

func main() {
	// Настройка клиента MQTT
	opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt.broker.address:1883")
	opts.SetClientID("mqtt_http_bridge")
	mqttClient = MQTT.NewClient(opts)

	// Подключаемся к брокеру MQTT
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Не удалось подключиться к MQTT-брокеру:", token.Error())
	}

	// Настройка HTTP-обработчика
	http.HandleFunc("/sendmqtt", handleHTTP)
	fmt.Println("Сервер HTTP слушает на порту :8080...")

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("Ошибка при запуске HTTP-сервера: ", err)
	}
}
