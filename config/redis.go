package config

type redisConfig struct {
	Status      bool   `json:"status" bson:"status"`
	Protocol    string `json:"protocol" bson:"protocol"`
	Host        string `json:"host" bson:"host"`
	Port        int    `json:"port" bson:"port"`
	Password    string `json:"password" bson:"password"`
	Db          int    `json:"db" bson:"db"`
	MaxIdle     int    `json:"max_idle" bson:"max_idle"`
	MaxActice   int    `json:"max_actice" bson:"max_actice"`
	IdleTimeOut int64  `json:"idle_time_out" bson:"idle_time_out"`
	Wait        bool   `json:"wait" bson:"wait"`
}
