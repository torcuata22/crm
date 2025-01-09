package database

import (
	"gorm.io/gorm"
	//"gorm.io/driver/sqlite" // Sqlite driver based on CGO
)

var DBConn *gorm.DB
