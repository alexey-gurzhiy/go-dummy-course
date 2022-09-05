// Разработайте пакет для чтения конфигурации типичного веб-приложения через флаги или переменные окружения.
// Пример конфигурации можно посмотреть здесь. По желанию вы можете задать другие имена полям, сгруппировать их или добавить собственные поля.
// Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
// Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).

// port: 8080
// db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
// jaeger_url: http://jaeger:16686
// sentry_url: http://sentry:9000
// kafka_broker: kafka:9092
// some_app_id: testid
// some_app_key: testkey

package main

import (
	configuration "Lesson8/Configuration"
	"fmt"
	"os"

)


func main() {
	cfg := configuration.GetConfiguration()

	fmt.Printf("port: 		%d \n", cfg.Port)
	fmt.Printf("db_url: 	%s \n", cfg.Db_url)
	fmt.Printf("jaeger_url: 	%s \n", cfg.Jaeger_url)
	fmt.Printf("sentry_url: 	%s \n", cfg.Sentry_url)
	fmt.Printf("kafka_broker: 	%s \n", cfg.Kafka_broker)
	fmt.Printf("some_app_id: 	%s \n", cfg.Some_app_id)
	fmt.Printf("some_app_key: 	%s \n", cfg.Some_app_key)
fmt.Println(" ")
file,err := os.Open("go.mod")
if err != nil {
	panic(err)
}

data := make([]byte, 24)
result, err := file.Read(data)
if err != nil{
	panic(err)
}
fmt.Println(result, string(data))
fmt.Printf("file: %v\n", file)

}
