package adapters

import (
	"database-conciliator/internal/db"
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"
)

type PostgreSQLAdapter struct {
	db *sql.DB
}

func NewPostgreSQLAdapter() *PostgreSQLAdapter {
	return &PostgreSQLAdapter{}
}

func (p *PostgreSQLAdapter) Connect() error {

	db, err := sql.Open("postgres", GetConnectionString())
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados PostgreSQL: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("erro ao testar a conexão com o banco de dados PostgreSQL: %v", err)
	}
	p.db = db

	return nil
}

func (p *PostgreSQLAdapter) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

func (p *PostgreSQLAdapter) Exec(query string, args ...interface{}) (sql.Result, error) {
	if p.db == nil {
		return nil, fmt.Errorf("não é possível executar consulta: adaptador não conectado ao banco de dados")
	}

	result, err := p.db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar consulta SQL: %v", err)
	}

	return result, nil
}

func (p *PostgreSQLAdapter) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if p.db == nil {
		return nil, fmt.Errorf("não é possível executar consulta: adaptador não conectado ao banco de dados")
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar consulta SQL: %v", err)
	}

	return rows, nil
}

// Implemente a interface db.Database no adaptador PostgreSQL
var _ db.Database = (*PostgreSQLAdapter)(nil)
