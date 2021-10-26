package dbtest

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

//现在我们需要为代码中的recordStats函数编写单元测试，但是又不想在测试过程中连接真实的数据库进行测试。这个时候我们就可以像下面示例代码中那样使用sqlmock工具去mock数据库操作。

func TestShouldUpdateStats(t *testing.T) {
	// mock一个*sql.DB对象，不需要连接真实的数据库
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mock执行指定SQL语句时的返回结果
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats:%v", err)
	}

	// 将mock的DB对象传入我们的函数中
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

// TestShouldRollbackStatUpdatesOnFailure sql执行失败回滚的测试用例
func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	// now we execute our method
	if err = recordStats(db, 2, 3); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations:%s", err)
	}

}
