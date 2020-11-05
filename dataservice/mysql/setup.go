package mysql

import (
	"fmt"

	"github.com/baby-vibe/domain/model"
	"github.com/jinzhu/gorm"
)

const (
	dbms = "mysql"
	user = "root"
	pass = "`1q2w3e4r"
	db   = "BABY_VIBE"
)

var dbConn *gorm.DB

func Setup() {
	var err error
	conn := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, db)
	dbConn, err = gorm.Open(dbms, conn)
	if err != nil {
		panic(err)
	}

	dbConn.AutoMigrate(
		&model.Signal{},
		&model.Weight{},
	)
	//dbConn.Model(&model.Signal{}).AddForeignKey("writer_id", "users(id)", "CASCADE", "CASCADE")
}
