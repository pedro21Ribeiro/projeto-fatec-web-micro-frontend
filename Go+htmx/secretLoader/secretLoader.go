package secretloader

import(
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBEnvVars struct {
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPassword string
}

type AuthEnvVars struct {
	SALT string
	SECRET string
}

var dbEnv *DBEnvVars
var authEnv *AuthEnvVars

func LoadDBEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}

	dbEnv = &DBEnvVars{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}

func LoadAuthEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}

	authEnv = &AuthEnvVars{
		SALT: os.Getenv("SALT"),
		SECRET: os.Getenv("SECRET"),
	}
}

func GetDBEnv() *DBEnvVars {
	if(dbEnv == nil){
		LoadDBEnv()
	}

	return dbEnv
}

func GetAuthEnv() *AuthEnvVars {
	if(authEnv == nil){
		LoadAuthEnv()
	}

	return authEnv
}