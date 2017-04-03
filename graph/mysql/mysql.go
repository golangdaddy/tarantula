package mysql

import (
	"fmt"
	"sync"
	"database/sql"
	//
	"github.com/golangdaddy/tarantula/log"
	"github.com/golangdaddy/tarantula/log/gcp"
	"github.com/golangdaddy/tarantula/graph"
	//
	mySql "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	SERVICE_NAME = "MYSQL"
)

type Credentials struct {
	ProjectId  string
	Connection string
	Database   string
	Username   string
	Password   string
}

func (credentials *Credentials) ProjectID() string {
	return credentials.ProjectId
}

func (credentials *Credentials) ServiceName() string {
	return SERVICE_NAME
}

func (credentials *Credentials) DatabaseName() string {
	return credentials.Database
}

type Client struct {
	Log logging.Logger
	DB *graph.Database
	Credentials graph.Credentials
	client *sql.DB
	sync.RWMutex
}

func (client *Client) SetDB(db *graph.Database) {

	client.Lock()

		client.DB = db

	client.Unlock()
}

func NewClient(credentials *Credentials) *Client {
	
	log := logs.NewClient(credentials.ProjectId).NewLogger(SERVICE_NAME)

	client, err := mySql.DialPassword(credentials.Connection, credentials.Username, credentials.Password)

	if err != nil { log.Panic(fmt.Sprintf("Could not open db: %v", err)) }

	return &Client{
		log,
		nil,
		credentials,
		client,
		sync.RWMutex{},
	}
}
