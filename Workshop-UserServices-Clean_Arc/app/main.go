package main

import (
	"clean/configs"
	"clean/modules/logs"
	"clean/modules/servers"
	"context"
	"log"
	"os"

	databases "clean/pkg/databases/postgresql"

	vault "github.com/hashicorp/vault/api"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

var (
	username      string
	password      string
	host          string
	port          string
	protocol      string
	database      string
	fiber_host    string
	fiber_port    string
	kafka_host    string
	kafka_group   string
	redis_addr    string
)

func main() {
	// Load dotenv config
	if err := godotenv.Load("../config.env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	// Connect Vault
	config := vault.DefaultConfig()

	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	ctx := context.Background()

	// Read a secret
	secret, err := client.KVv2((os.Getenv("VAULT_TYPE"))).Get(ctx, os.Getenv("VAULT_PATH"))
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	password, _ = secret.Data["password"].(string)
	username, _ = secret.Data["username"].(string)
	host, _ = secret.Data["host"].(string)
	port, _ = secret.Data["port"].(string)
	protocol, _ = secret.Data["protocol"].(string)
	database, _ = secret.Data["database"].(string)
	fiber_host, _ = secret.Data["fiber_host"].(string)
	fiber_port, _ = secret.Data["fiber_port"].(string)
	kafka_host, _ = secret.Data["kafka_host"].(string)
	kafka_group, _ = secret.Data["kafka_group"].(string)
	redis_addr, _ = secret.Data["redis_addr"].(string)


	// Fiber configs
	cfg.App.Host = fiber_host
	cfg.App.Port = fiber_port

	// Database Configs
	cfg.PostgreSQL.Host = host
	cfg.PostgreSQL.Port = port
	cfg.PostgreSQL.Protocol = protocol
	cfg.PostgreSQL.Username = username
	cfg.PostgreSQL.Password = password
	cfg.PostgreSQL.Database = database
	log.Println("vault password: ",password)


	// Kafka Config
	cfg.Sarama.Host = []string{kafka_host}
	cfg.Sarama.Group = kafka_group

	//Redis Config
	cfg.Redis.Addr = redis_addr
	log.Println(cfg.Sarama.Host)
	// New Database
	db, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//defer db.Close()

	// Kafka Producer Config

	producer, err := sarama.NewSyncProducer(cfg.Sarama.Host, nil)
	if err != nil {
		logs.Error(err)
	}
	defer producer.Close()

	// Kafka Consumer Config
	consumer, err := sarama.NewConsumerGroup(cfg.Sarama.Host, cfg.Sarama.Group, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	s := servers.NewServer(cfg, db, producer, consumer)
	s.Start()
}
