package config

type dbConfig struct {
	Status          bool   `json:"status" bson:"status"`
	TableCheck      bool   `json:"table_check" bson:"table_check"`
	Type            string `json:"type" bson:"type"`
	Host            string `json:"host" bson:"host"`
	Port            int    `json:"port" bson:"port"`
	Db              string `json:"db" bson:"db"`
	User            string `json:"user" bson:"user"`
	Password        string `json:"password" bson:"password"`
	Charset         string `json:"charset" bson:"charset"`
	DSN             string `json:"dsn" bson:"dsn"`
	Prefix          string `json:"prefix" bson:"prefix"`
	SetMaxIdleConns int    `json:"set_max_idle_conns" bson:"set_max_idle_conns"`
	SetMaxOpenConns int    `json:"set_max_open_conns" bson:"set_max_open_conns"`
}
