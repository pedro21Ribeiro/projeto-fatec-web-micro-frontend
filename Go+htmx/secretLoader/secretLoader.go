package secretloader

import(
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPassword string
}

func LoadEnv() EnvVars {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}

	return EnvVars{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}