package main

import (
	httpgoods "goboilerplate-domain-driven/internal/adapter/http/goods"
	goodsRepo "goboilerplate-domain-driven/internal/adapter/repository/goods"
	"goboilerplate-domain-driven/internal/infra"
	serviceGoods "goboilerplate-domain-driven/internal/usecase/goods"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	appHost, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	psqlDB, err := infra.NewPostgresDB(infra.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPort,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAMES"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	goodsRepo := goodsRepo.NewGoodsRepository(psqlDB)

	ucGoods := serviceGoods.NewServiceGoods(goodsRepo)

	httpGoods := httpgoods.NewGoodsHandler(ucGoods)

	httpServer := infra.NewNetHttpServer(os.Getenv("APP_HOST"), uint(appHost))
	httpServer.RouteNetHttp(httpGoods)
	httpServer.NetHttpListen()

}
