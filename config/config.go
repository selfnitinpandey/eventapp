package config

type AppConfig struct {
	MONGO_URI    string
	DATABASENAME string
	SERVER       string
	PORT         string
}

func LoadConfig() AppConfig {
	// err := godotenv.Load("./.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// appConfig := AppConfig{
	// 	os.Getenv("MONGO_URI"), os.Getenv("DATABASE_NAME"), os.Getenv("SERVER"), os.Getenv("PORT"),
	// }

	appConfig := AppConfig{"mongodb://localhost:27017", "eventapp", "localhost", "27017"}
	return appConfig
}
