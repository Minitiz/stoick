package db

import (
	"database/sql"
	"fmt"
)

// DB représente l'interface vers la base de données PostgreSQL
var DBClient *sql.DB

type Client interface {
	InsertURL(urlShortener Data) error
	ReadDataAndIncrementAccess(urlShortener Data) ([]Data, error)
}

type DbInfoStruct struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

var DbInfo DbInfoStruct

func New(host string,
	port string,
	user string,
	password string,
	dbName string) DbInfoStruct {
	return DbInfoStruct{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbName:   dbName}
}

func (info DbInfoStruct) CreateClient() error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.host, info.port, info.user, info.password, info.dbName)

	client, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	if err := client.Ping(); err != nil {
		client.Close()
		return err
	}
	DBClient = client
	return nil
}
