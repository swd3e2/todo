package application

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client *mongo.Client
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Connect(config *Config) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	s.client = client
	// db, err := sql.Open("mysql", fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%s)/%s",
	// 	config.Store.User,
	// 	config.Store.Pwd,
	// 	config.Store.Dsn,
	// 	config.Store.Port,
	// 	config.Store.Database,
	// ))

	// if err != nil {
	// 	return err
	// }

	// if err := db.Ping(); err != nil {
	// 	return err
	// }

	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	// this.Db = db

	return nil
}

func (s *Store) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}
