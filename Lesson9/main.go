// К приложению из практической части предыдущего урока добавьте возможность читать данные из файлов.
// Конфигурация может быть задана в форматах yaml или json. Также по желанию вы можете добавить и другие форматы.
// Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
// Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).

package main

import (
	configuration "Lesson9/Configuration"
	"fmt"
	"os"
)

func main() {
	os.WriteFile("store.json", []byte(`{
	"port": 9090, 
	"db_url": "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable/new-url",
	"jaeger_url": "http://jaeger:19999",
	"sentry_url": "http://sentry:9999",
	"kafka_broker": "kafka:1999",
	"some_app_id": "testid_new",
	"some_app_key": "testkey_new"
	"wrong_key": "wrong"
	}`), 0777)

	fmt.Println(" ")
	file, err := os.Open("store.json")
	if err != nil {
		panic(err)
	}

	defer os.Remove("store.json")
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		panic(err)
	}

	fileData := string(data)

	cfg := configuration.GetConfiguration(fileData)

	fmt.Printf("port: 		%d \n", cfg.Port)
	fmt.Printf("db_url: 	%s \n", cfg.Db_url)
	fmt.Printf("jaeger_url: 	%s \n", cfg.Jaeger_url)
	fmt.Printf("sentry_url: 	%s \n", cfg.Sentry_url)
	fmt.Printf("kafka_broker: 	%s \n", cfg.Kafka_broker)
	fmt.Printf("some_app_id: 	%s \n", cfg.Some_app_id)
	fmt.Printf("some_app_key: 	%s \n", cfg.Some_app_key)

	// fmt.Println(fileData)
	// fmt.Println(string(data))

}
