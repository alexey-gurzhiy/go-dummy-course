package configuration

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func (f FlagConfig) Decode(value string) string {
	parsed, err := url.Parse(value)
	if err != nil {
		fmt.Println("Oops, something went wrong.")
	}
	return parsed.String()
}

type FlagConfig struct {
	Port         int
	Db_url       string
	Jaeger_url   string
	Sentry_url   string
	Kafka_broker string
	Some_app_id  string
	Some_app_key string
}

func GetConfiguration() FlagConfig {
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

	return cfg
}
