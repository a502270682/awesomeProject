package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	// 引入数据库驱动注册及初始化
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const TableUser = "user"


// Config defines a cluster of mysql dbs, master and slave
type Config struct {
	Master *ConnectionConfig
	Slave  *ConnectionConfig
}

// ConnectionConfig defines configuration of a db connection
type ConnectionConfig struct {
	Dsn     string
	MaxIdle int
	MaxOpen int
	Name    string // we use this name for logging and reporting, no space allowed
	Debug   bool
}

var (
	ormLog = flag.Bool("ormlog", false, "--ormlog")
	master *gorm.DB
)

func init(){
	_, err := InitializeDao(&Config{
		Master: &ConnectionConfig{
			Dsn: "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
			MaxIdle: 100,
			MaxOpen: 200,
		},
	})
	if err != nil {
		panic(err)
	}
}

func InitializeDao(o *Config) (*gorm.DB, error) {
	masterDB, err := gorm.Open("mysql", o.Master.Dsn)
	if err != nil {
		errStr := fmt.Sprintf("failed to open MySQL master db, error=%v", err)
		return nil, errors.New(errStr)
	}
	master = masterDB
	if *ormLog {
		master = masterDB.Debug()
	}
	master.DB().SetMaxIdleConns(o.Master.MaxIdle)
	master.DB().SetMaxOpenConns(o.Master.MaxOpen)
	master.SetNowFuncOverride(func() time.Time {
		return time.Now().UTC()
	})
	master.SingularTable(true)
	// master.SetLogger(gorm.Logger{})
	return master, nil
}

func getClient() *gorm.DB{
	if master == nil {
		// return InitializeDao() get global config
		return nil
	}
	return master
}

type User struct {
	*gorm.Model
	Name string
	Age int
}

func (u *User) TableName() string {
	return TableUser
}

func creatUser(db *gorm.DB, u *User) error{
	return db.Create(u).Error
}

func updateMany() error {
	var err error
	var users []*User
	db := getClient()
	if err = db.Where("age > (?)", 16).Find(&users).Error; err != nil {
		panic(err)
	}
	for i := range users {
		users[i].Age = 15
	}
	if err = db.Where("age > (?)", 16).UpdateColumns(&users).Error; err != nil {
		panic(err)
	}
	var newUsers []*User
	if err = db.Where("id = (?)", users[0].ID).Find(&newUsers).Error; err != nil {
		panic(err)
	}
	fmt.Println(newUsers[0])
	return nil
}

func main() {

	//if err = creatUser(getClient(), &User{
	//	Name: "lyj",
	//	Age: 17,
	//}); err != nil {
	//	panic(err)
	//}
	updateMany()

}