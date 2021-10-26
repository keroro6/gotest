package dbtest

import "database/sql"

// recordStats 一个操作db的函数
func recordStats(db *sql.DB, userID, productID int64) (err error) {
	// 开启事务
	// 操作views和product_viewers两张表
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}

	// 更新products表
	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		return
	}

	return
}
