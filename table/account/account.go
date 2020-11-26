package account

import (
	"context"
	"database/sql"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"fmt"
)

type Account struct {
	ID           uint
	Username     string
	Full_Name    string
	Password     string
	Gender       bool
	Phone_Number string
	Address      string
	Email        string
	Position     string
	Permission   string
}
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Account) TableName() string {
	return "Account"
}

var server = "localhost"
var port = 1433
var user = "sa"
var password = "1234"
var database = "TX_EquipmentManager"
var (
	ctx context.Context
	db  *sql.DB
)

func AllUsers(c *gin.Context) {
	// sqlDB, err := db.DB()

	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// sqlDB.SetMaxIdleConns(10)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDB.SetMaxOpenConns(100)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDB.SetConnMaxLifetime(time.Hour)

	dsn := "sqlserver://sa:1234@localhost:1433?database=TX_EquipmentManager"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err})

		fmt.Println(err)

	} else {
		account := Account{Username: "Jinzhu", Password: "1234"}
		fmt.Println("Connected !")
		result := db.Create(&account)
		c.JSON(http.StatusOK, gin.H{"message": "create thành công"})
		fmt.Println(result)
	}
}
