package main

import (
  "fmt"
  "time"
  "context"
  "net/http"
  "backend/ent"
  "go.uber.org/zap"
  "backend/ent/migrate"
  "github.com/gin-gonic/gin"
  ginzap "github.com/gin-contrib/zap"
  helper "backend/src/helpers"

  "database/sql"
  "entgo.io/ent/dialect"
  entsql "entgo.io/ent/dialect/sql"
  _ "github.com/jackc/pgx/v5/stdlib"
)

func replaceLogger(r *gin.Engine){
  /*
    Referencing https://github.com/gin-contrib/zap

    Building a zap logger

  */
  config := zap.NewProductionConfig()
  config.OutputPaths = []string{fmt.Sprintf("/usr/src/app/logs/%s.log", helper.GetEnvWithDefault("env", "development")), "stdout"}
  zapLogger, _ := config.Build()
  r.Use(ginzap.GinzapWithConfig(zapLogger, &ginzap.Config{
    TimeFormat: time.RFC3339,
    UTC: true,
  }))
}

func databaseOpen(databaseUrl string) *ent.Client {
    db, err := sql.Open("pgx", databaseUrl)
    if err != nil {
        fmt.Println(err)
    }

    // Create an ent.Driver from `db`.
    drv := entsql.OpenDB(dialect.Postgres, db)
    return ent.NewClient(ent.Driver(drv))
}

func runMigrations() (*ent.Client, context.Context){
  /*

    Adding migrations

  */
  database_url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
                                helper.GetEnv("DATABASE_USER"),
                                helper.GetEnv("DATABASE_PASSWORD"),
                                helper.GetEnv("DATABASE_HOST"),
                                helper.GetEnv("DATABASE_PORT"),
                                helper.GetEnv("DATABASE_NAME"))

  client := databaseOpen(database_url)
  defer client.Close()
  ctx := context.Background()
  err := client.Schema.Create(
      ctx,
      migrate.WithDropIndex(true),
      migrate.WithDropColumn(true),
  )
  if err != nil {
      fmt.Printf("failed creating schema resources: %v \n", err)
  }
  return client, ctx
}

func main() {
  r := gin.Default()

  replaceLogger(r)
  runMigrations()

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run()
}