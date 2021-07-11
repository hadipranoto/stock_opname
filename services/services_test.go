package services

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)



func SqlConnect() (*sql.DB, error) {
	user := "root"
	password := "tarosabi"
	host := "127.0.0.1"
	port := "3306"
	dbName := "stock_opname"

	sqlSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user, password,host,port,dbName)

	db, err := sql.Open("mysql", sqlSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestSqlConnection (t *testing.T){
	conn, err := SqlConnect()
	if err != nil {
		panic(err)
	}
	assert.True(t, conn != nil, "sql connection should not nil")
}
func TestService (t *testing.T){
	conn, err := SqlConnect()
	if err != nil {
		panic(err)
	}
	serv := NewOpnameService(conn)
	data, err := serv.SummaryOpname()
	if err != nil {
		panic(err)
	}	
	assert.True(t, len(data.([]map[string]interface{})) > 0, "Data opname should contain at least 1 row of product")
	

}

