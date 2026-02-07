package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mrhb787/hospital-ward-manager/model"
	_ "github.com/lib/pq"
)

type Client struct {
	DB *sql.DB
}

type Service interface {
	NewClient() (*Client, error)
	GetClient() (*Client, error)
	GetUserById(userId uint32) (model.User, error)
	GetUserByPhone(phone string) (model.User, error)
	CreateUserSession(session model.UserSession) (err error)
	GetUserSession(token string, userId int) (session model.UserSession, err error)
}

type service struct {
	dbConn string
	client *Client
}

func NewService(connStr string, client *Client) Service {
	return &service{dbConn: connStr, client: client}
}

func (s *service) NewClient() (*Client, error) {
	db, err := sql.Open("postgres", s.dbConn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established")
	s.client = &Client{DB: db}
	return s.client, nil
}

func (s *service) GetClient() (*Client, error) {
	dbClient := s.client
	if dbClient == nil {
		var err error
		dbClient, err = s.NewClient()
		if err != nil {
			return nil, err
		}
		return dbClient, nil
	}
	return dbClient, nil
}

func (c *Client) Close() error {
	return c.DB.Close()
}

func (c *Client) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := c.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return rows, nil
}

func (c *Client) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := c.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %w", err)
	}
	return result, nil
}

func (c *Client) QueryRow(query string, args ...interface{}) *sql.Row {
	return c.DB.QueryRow(query, args...)
}
