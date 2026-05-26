package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Record map[string]any

type Table[T any] interface {
	GetName() string
	GetColumns() []string
	Read(record *Record) (*T, error)
	ToRecord(entity *T) *Record
}

type Database[T any] interface {
	Query(queryString string, args ...any) ([]*T, error)
	Get(id string) (*T, error)
	Create(entity *T) (string, error)
	Update(id string, entity *T) error
	Delete(id string) error
}

func NewPGConnection(databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return conn, nil
}

type postgresDatabase[T any] struct {
	conn  *pgx.Conn
	table Table[T]
}

func NewPostgresDatabase[T any](conn *pgx.Conn, table Table[T]) Database[T] {
	return &postgresDatabase[T]{conn: conn, table: table}
}

func (db *postgresDatabase[T]) Query(queryString string, args ...any) ([]*T, error) {
	rows, err := db.conn.Query(context.Background(), queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*T
	for rows.Next() {
		record := make(Record)
		err := rows.Scan(&record)
		if err != nil {
			return nil, err
		}

		entity, err := db.table.Read(&record)
		if err != nil {
			return nil, err
		}
		results = append(results, entity)
	}
	return results, nil
}

func (db *postgresDatabase[T]) Get(id string) (*T, error) {
	queryString := "SELECT * FROM " + db.table.GetName() + " WHERE id = $1"
	results, err := db.Query(queryString, id)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no record found with id: %s", id)
	}
	return results[0], nil
}

func (db *postgresDatabase[T]) Create(entity *T) (string, error) {
	record := db.table.ToRecord(entity)
	columns := db.table.GetColumns()
	values := make([]any, len(columns))
	for i, column := range columns {
		values[i] = (*record)[column]
	}
	queryString := "INSERT INTO " + db.table.GetName() + " (" + joinColumns(columns) + ") VALUES (" + placeholders(len(columns)) + ") RETURNING id"
	var id string
	err := db.conn.QueryRow(context.Background(), queryString, values...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (db *postgresDatabase[T]) Update(id string, entity *T) error {
	record := db.table.ToRecord(entity)
	columns := db.table.GetColumns()
	values := make([]any, len(columns))
	for i, column := range columns {
		values[i] = (*record)[column]
	}
	queryString := "UPDATE " + db.table.GetName() + " SET " + setColumns(columns) + " WHERE id = $1"
	_, err := db.conn.Exec(context.Background(), queryString, append(values, id)...)
	return err
}

func (db *postgresDatabase[T]) Delete(id string) error {
	queryString := "DELETE FROM " + db.table.GetName() + " WHERE id = $1"
	_, err := db.conn.Exec(context.Background(), queryString, id)
	return err
}

func joinColumns(columns []string) string {
	result := ""
	for i, column := range columns {
		if i > 0 {
			result += ", "
		}
		result += column
	}
	return result
}

func placeholders(n int) string {
	result := ""
	for i := 1; i <= n; i++ {
		if i > 1 {
			result += ", "
		}
		result += fmt.Sprintf("$%d", i)
	}
	return result
}

func setColumns(columns []string) string {
	result := ""
	for i, column := range columns {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%s = $%d", column, i+1)
	}
	return result
}
