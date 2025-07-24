package database
import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)
func DatabaseConnection(connectionString string) (*sql.DB,error) {
	db,err:=sql.Open("postgres",connectionString)
	if err!=nil{
		return nil,errors.New("Error: Database Connection")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.New("Error: Cannot reach database")
	}
	return db,nil
}