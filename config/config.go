package config

type Mode string

const (
	DebugMode   = Mode("debug")
	ReleaseMode = Mode("release")
)

type KafkaProducer struct {
	Broker string `json:"broker"                        validate:"required"`
}

type KafkaConsumer struct {
	Broker          string `json:"broker"               validate:"required"`
	GroupId         string `json:"group_id"             validate:"required"`
	AutoOffsetReset string `json:"auto_offset_reset"    validate:"required"`
}

type Kafka struct {
	Producer KafkaProducer `json:"producer"`
	Consumer KafkaConsumer `json:"consumer"`
}

type Config struct {
	Mode        Mode   `json:"mode"                     validate:"required"`
	ServiceName string `json:"service_name"             validate:"required"`
	HttpPort    int    `json:"http_port"                validate:"required"`
	Kafka       Kafka  `json:"kafka"`
}

var config *Config

func init() {
	config = &Config{}
}

func GetConfig() Config {
	return *config
}
