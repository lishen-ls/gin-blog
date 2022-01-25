package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type DB struct {
	DBInterface
	// 数据库连接
	conn *sql.DB
}

// DBInterface 数据库连接接口
type DBInterface interface {
	Query(query string, params ...interface{})
	Insert(query string, params ...interface{})
	BatchInsert(query string, params []interface{})
	Execute(query string, params ...interface{})
}

// DO 结果集结构
type DO struct {
	DOInterface
	rows *sql.Rows
}

// DOInterface 结果集接口
type DOInterface interface {
	Fetch() interface{}
	FetchAll() []interface{}
}

// ConnectConfig 数据库连接配置
type ConnectConfig struct {
	username string
	password string
	host     string
	database string
}

// Open
// 开启一个数据库连接
func Open(config ConnectConfig) (*DB, error) {
	username := config.username
	password := config.password
	host := strings.Trim(config.host, " ")
	database := config.database
	var connectStr string

	if host == "" {
		connectStr = fmt.Sprintf("%s:%s@/%s", username, password, database)
	} else {
		connectStr = fmt.Sprintf("%s:%s@[%s]/%s", username, password, host, database)
	}

	db, err := sql.Open(
		"mysql",
		connectStr,
	)

	return &DB{conn: db}, err
}

// Query 搜索
func (db *DB) Query(query string, params ...interface{}) (*DO, error) {
	result, err := db.conn.Query(query, params...)
	do := &DO{rows: result}
	return do, err
}

func (do *DO) Fetch() (map[string]interface{}, error) {
	columns, err := do.rows.Columns()

	if err != nil {
		return nil, err
	}

	columnsCount := len(columns)

	cache := make([]interface{}, columnsCount)

	for index := range cache {
		var pointer interface{}
		cache[index] = &pointer
	}

	if do.rows.Next() {
		_ = do.rows.Scan(cache...)
		item := make(map[string]interface{})
		for index, value := range cache {
			item[columns[index]] = *value.(*interface{})
		}
		return item, nil
	}

	return nil, nil
}

func main() {
	db, _ := Open(ConnectConfig{
		username: "root",
		password: "970414",
		host:     "",
		database: "data",
	})
	do, _ := db.Query("SELECT title, content FROM article WHERE id = ?", 3)
	result, _ := do.Fetch()
	fmt.Println(result)

	//db, err := sql.Open("mysql", "root:970414@/")
	//if err != nil {
	//	println("error")
	//	panic(err)
	//}
	//rows, _ := db.Query("SELECT id, title FROM article")
	//
	//for rows.Next() {
	//	var id int
	//	var title string
	//	err = rows.Scan(&id, &title)
	//	println(id, title)
	//}
}
