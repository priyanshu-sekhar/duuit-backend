// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"database/sql"
	"duuit/controllers"
	"duuit/models/dao"
	"duuit/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	sqlConn, _ := beego.AppConfig.String("sqlconn")
	connectionString := fmt.Sprintf(sqlConn)
	var err error
	sqlDB, err := sql.Open("mysql", connectionString)
	utils.LogError(err)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	utils.LogError(err)


	migrationError := db.
		Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&dao.Goal{},
			&dao.User{},
		)
	utils.LogError(migrationError)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/goal",
			beego.NSInclude(
				&controllers.GoalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}


