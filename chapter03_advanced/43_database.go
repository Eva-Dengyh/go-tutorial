// 本课目标：了解 database/sql 的基本 CRUD 流程。
//
// 知识点：
//   - sql.Open 打开连接（需导入驱动：_ "github.com/glebarez/sqlite"）
//   - db.Query / Exec；rows.Scan 读结果
//   - 使用占位符 ? 防止 SQL 注入
//
// 运行：go run ./chapter03_advanced/43_database.go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT);
		INSERT INTO users (name) VALUES (?), (?);
	`, "Alice", "Bob")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Printf("id=%d name=%s\n", id, name)
	}
}
