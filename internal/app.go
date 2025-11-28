package internal

import (
	"context"
	"embed"
	httpgoods "goboilerplate-domain-driven/internal/adapter/http/goods"
	"goboilerplate-domain-driven/internal/adapter/middleware"
	goodsRepo "goboilerplate-domain-driven/internal/adapter/repository/goods"
	"goboilerplate-domain-driven/internal/infra"
	"goboilerplate-domain-driven/internal/infra/observability"
	serviceGoods "goboilerplate-domain-driven/internal/usecase/goods"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

//go:embed infra/migrations/postgres/*.sql
var embedMigrations embed.FS

func App() {

	ctx := context.Background()

	tracer, err := observability.InitTracer(ctx, observability.Config{
		ServiceName: "goboilerplate-domain-driven",
		OtelMode:    os.Getenv("OTEL_TRACER_MODE"),
		Endpoint:    os.Getenv("OTEL_TRACER_OTLP_ENDPOINT"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer tracer.Shutdown(ctx)

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	psqlDB, err := infra.NewInitDB(infra.DBConfig{
		Driver:   os.Getenv("DB_DRIVER"),
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

	validate := validator.New()

	goodsRepo := goodsRepo.NewGoodsRepository(psqlDB)
	ucGoods := serviceGoods.NewServiceGoods(goodsRepo)
	httpGoods := httpgoods.NewGoodsHandler(ucGoods, validate)

	httpServer := infra.NewNetHttpServer(os.Getenv("APP_HOST"), uint(appPort))
	httpServer.Use(middleware.LoggingMiddleware)
	httpServer.RouteNetHttp(httpGoods)
	httpServer.NetHttpListen()
}
