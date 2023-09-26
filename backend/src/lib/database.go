package lib

import  (
  "fmt"
  "context"
  "backend/ent"
  "database/sql"
  "backend/ent/migrate"
  "entgo.io/ent/dialect"
  helper "backend/src/helpers"
  entsql "entgo.io/ent/dialect/sql"
  _ "github.com/jackc/pgx/v5/stdlib"
)

type dbContext struct {
	Client *ent.Client
	Context context.Context
}

func NewDbContext(c *ent.Client, ctx context.Context) *dbContext{
	return &dbContext{c, ctx}
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

func ConnectToDatabase() *dbContext{
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
  ctx := context.Background()
  err := client.Schema.Create(
      ctx,
      migrate.WithDropIndex(true),
      migrate.WithDropColumn(true),
  )
  if err != nil {
      fmt.Printf("failed creating schema resources: %v \n", err)
  }
  return NewDbContext(client, ctx)
}

var DbCtx *dbContext

func init() {
	DbCtx = ConnectToDatabase()
}

func CloseDBConnection() {
  DbCtx.Client.Close()
}