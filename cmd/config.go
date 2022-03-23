package cmd

type Config struct {
	RestPort *int `env:"REST_PORT,default=8080"`

	Db struct {
		Debug *bool   `env:"DB_DEBUG,default=false"`
		Dsn   *string `env:"DSN,required=true"`
	}
}
