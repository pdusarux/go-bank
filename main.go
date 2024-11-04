package main

import (
	"fmt"
	"go-bank/handler"
	"go-bank/logs"
	"go-bank/repository"
	"go-bank/service"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	initConfig()
	initTimezone()
	db := initDB()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	accountRepositoryDB := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)

	r := gin.Default()

	r.SetTrustedProxies([]string{"192.168.1.1", "192.168.1.2"})

	r.GET("/customers", customerHandler.GetCustomers())
	r.GET("/customers/:customer_id", customerHandler.GetCustomer())

	r.POST("/customers/:customer_id/accounts", accountHandler.NewAccount)
	r.GET("/customers/:customer_id/accounts", accountHandler.GetAccounts)

	logs.Info("Starting server on port", zap.Int("port", viper.GetInt("app.port")))
	r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDB() *sqlx.DB {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"),
	)

	db, err := sqlx.Connect(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Duration(viper.GetInt("db.conn_max_lifetime")) * time.Minute)
	db.SetMaxOpenConns(viper.GetInt("db.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("db.max_idle_conns"))

	return db
}
