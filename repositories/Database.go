package repositories

import (
	"context"
	"fmt"
	"log"
	"stori/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDatabase struct {
	connection *pgxpool.Pool
}

const connStringTemplate = "postgresql://%s:%s@%s:%d/%s"

func NewDatabase(config config.DBConfig) (*PostgresDatabase, error) {
	uri := fmt.Sprintf(connStringTemplate, config.UserName, config.Password, config.Host, config.Port, config.DbName)

	dbCfg, err := pgxpool.ParseConfig(uri)
	if err != nil {
		log.Fatalf("DB | Impossible to parse initial configuration %s", err)
	}

	connection, err := pgxpool.ConnectConfig(context.Background(), dbCfg)
	if err != nil {
		log.Fatalf("DB | Impossible to connect to the database %s", err)
	}

	dbCfg.MaxConns = config.PoolSize
	dbCfg.MaxConnLifetime = config.MaxConnLifetime

	return &PostgresDatabase{connection: connection}, nil
}

func (db *PostgresDatabase) Save(query string, args ...interface{}) (int64, error) {
	ctx := context.Background()

	txn, err := db.connection.Begin(ctx)
	if err != nil {
		return 0, err
	}

	res, err := db.connection.Exec(ctx, query, args...)
	if err != nil {
		rollbackErr := txn.Rollback(ctx)
		log.Printf("ERROR | DB | Failed to run transaction %s. RollbackErr: %s", err, rollbackErr)

		return 0, err
	}

	affectedRows := res.RowsAffected()

	return affectedRows, txn.Commit(ctx)
}
