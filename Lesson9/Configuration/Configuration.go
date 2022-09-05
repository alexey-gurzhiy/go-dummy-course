package configuration

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
)
//Помню, что нет смысла загонять в декод и оно и так может спарсить, но переделывать и так не успеваю. Уже после курса самостоятельно посмотрю и переделаю декод как в разборе 
func (f FlagConfig) Decode(value string) string {
	parsed, err := url.Parse(value)
	if err != nil {
		fmt.Println("Oops, something went wrong.")
	}
	return parsed.String()
}

type FlagConfig struct {
	Port         int    `json:"port"`
	Db_url       string `json:"db_url"`
	Jaeger_url   string `json:"jaeger_url"`
	Sentry_url   string `json:"sentry_url"`
	Kafka_broker string `json:"kafka_broker"`
	Some_app_id  string `json:"some_app_id"`
	Some_app_key string `json:"some_app_key"`
}

func GetConfiguration(fileData string) FlagConfig {
	cfg := FlagConfig{}
	check := make(map[string]bool)
	check["port"] = false
	check["db_url"] = false
	check["jaeger_url"] = false
	check["sentry_url"] = false
	check["kafka_broker"] = false
	check["some_app_id"] = false
	check["some_app_key"] = false

	flag.IntVar(&cfg.Port, "port", 8080, "Please enter the port to connect to")
	flag.StringVar(&cfg.Db_url, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "Please provide a link to DB")
	flag.StringVar(&cfg.Jaeger_url, "jaeger_url", "http://jaeger:16686", "Please provide an url to jaeger")
	flag.StringVar(&cfg.Sentry_url, "sentry_url", "http://sentry:9000", "Please provide a url to a sentry")
	flag.StringVar(&cfg.Kafka_broker, "kafka_broker", "kafka:9092", "Please provide a kafka broker")
	flag.StringVar(&cfg.Some_app_id, "some_app_id", "testid", "Please provide app id")
	flag.StringVar(&cfg.Some_app_key, "some_app_key", "testkey", "Please provide a key to the app")

	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		fmt.Printf("The flag is %s with a value: %v\n", f.Name, f.Value)
		check[f.Name] = true
	})

	if !check["port"] {
		if os.Getenv("port") != "" {
			port, err := strconv.Atoi(os.Getenv("port"))
			if err != nil {
				panic("Ааааа, все пропало. Ну или просто в порт передали не стрингу в свойствах переменных окружения")
			}
			cfg.Port = port
		}
	}
	if !check["db_url"] {
		if os.Getenv("db_url") != "" {
			cfg.Db_url = os.Getenv("db_url")
		}
	}
	if !check["jaeger_url"] {
		if os.Getenv("jaeger_url") != "" {
			cfg.Jaeger_url = os.Getenv("jaeger_url")
		}
	}
	if !check["sentry_url"] {
		if os.Getenv("sentry_url") != "" {
			cfg.Sentry_url = os.Getenv("sentry_url")
		}
	}
	if !check["kafka_broker"] {
		if os.Getenv("kafka_broker") != "" {
			cfg.Kafka_broker = os.Getenv("kafka_broker")
		}
	}
	if !check["some_app_id"] {
		if os.Getenv("some_app_id") != "" {
			cfg.Some_app_id = os.Getenv("some_app_id")
		}
	}
	if !check["some_app_key"] {
		if os.Getenv("some_app_key") != "" {
			cfg.Some_app_key = os.Getenv("some_app_key")
		}
	}

	cfg.Db_url = cfg.Decode(cfg.Db_url)
	cfg.Jaeger_url = cfg.Decode(cfg.Jaeger_url)
	cfg.Sentry_url = cfg.Decode(cfg.Sentry_url)

	fmt.Println(fileData)
	var flags FlagConfig
	if err := json.Unmarshal([]byte(fileData), &flags); err != nil {
		fmt.Printf("\n\n!!!!!		There was an error. Ommiting json file. Error message: %s	!!!!!\n\n", err)
	} else {
		cfg = flags
	}
	cfg.Db_url = cfg.Decode(cfg.Db_url)
	cfg.Jaeger_url = cfg.Decode(cfg.Jaeger_url)
	cfg.Sentry_url = cfg.Decode(cfg.Sentry_url)
	return cfg
}
