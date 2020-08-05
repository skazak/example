package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skazak/example/internal/app/delivery/http/handler"
	"github.com/skazak/example/internal/app/delivery/http/middleware"
	"github.com/skazak/example/internal/app/driver"
	"github.com/skazak/example/internal/app/repository/postgres"
	"github.com/skazak/example/internal/app/service"
)

func getEnvFatal(key string) string {
	res := os.Getenv(key)
	if res == "" {
		log.Fatalf("please define env variable: %s", key)
	}
	return res
}

func getEnvWithDefault(key string) string {
	res := os.Getenv(key)
	if res == "" {
		return "8080"
	}
	return res
}

func main() {
	port := getEnvFatal("PORT")
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatal("env variable PORT must be int")
	}

	pgURL := getEnvFatal("PG_HOST")
	pgDB := getEnvFatal("PG_DB")
	pgLogin := getEnvFatal("PG_LOGIN")
	pgPass := getEnvFatal("PG_PASS")

	err := driver.ConnectPg(fmt.Sprintf("postgres://%s:%s@%s/%s", pgLogin, pgPass, pgURL, pgDB))
	if err != nil {
		log.Fatal(err)
	}

	cr := postgres.NewPgCategoryRepository(driver.Driver.PgConn)
	ir := postgres.NewPgItemRepository(driver.Driver.PgConn)
	ur := postgres.NewPgUserRepository(driver.Driver.PgConn)

	cs := service.NewCategoryService(cr, ir)
	is := service.NewItemService(ir)
	us := service.NewUserService(ur)

	r := gin.Default()

	loginGroup := r.Group("/")
	handler.InitAuthHandler(loginGroup, us)

	actionGroup := r.Group("/", middleware.AuthenticationRequired())
	handler.InitCategoryHandler(actionGroup, cs)
	handler.InitItemHandler(actionGroup, is)
	handler.InitUserHandler(actionGroup, us)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
