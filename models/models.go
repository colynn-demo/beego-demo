package models

import (
	"sync"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User ..
type User struct {
	ID      int `orm:"column(id);int(8);auto" json:"id"`
	Name    string
	WebSite string
	Email   string
}

// DBinit ..
func DBinit() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/beego_demo?charset=utf8")
	orm.RegisterModel(new(User))

	// sync table
	orm.RunSyncdb("default", false, true)
}

var globalOrm orm.Ormer
var once sync.Once

// GetOrmer :set ormer singleton
func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
