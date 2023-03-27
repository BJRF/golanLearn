package main

import (
	"fmt"
	"github.com/jialanli/windward"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	id   int
	Name string
	Age  int
}

type Config struct {
	Data struct {
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

// 设置表名，可以通过给Food struct类型定义 TableName函数，返回一个字符串作为表名
func (u UserInfo) TableName() string {
	return "UserInfo"
}

/*
ORM的全称是：Object Relational Mapping(对象关系映射)，其主要作用是在编程中，把面向对象的概念跟数据库中表的概念对应起来。
举例来说就是我们定义一个对象，那就对应着一张表，这个对象的实例，就对应着表中的一条记录。
*/
func mysqlTest() {
	file := "./conf/database.yaml" //配置文件所在位置
	var config_paths = []string{file}
	w := windward.GetWindward()
	w.InitConf(config_paths) //初始化自定义的配置文件
	var config Config
	err := w.ReadConfig(config_paths[0], &config)
	if err != nil {
		fmt.Sprintln("读取config失败")
		return
	}
	name := w.GetValString(file, "name")
	pwd := w.GetValString(file, "password")
	addr := w.GetValString(file, "addr") //ip:port
	database := w.GetValString(file, "database")

	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", name, pwd, addr, database)
	// step1 : 连接数据库
	db, err := gorm.Open("mysql", args)
	if err != nil {
		fmt.Printf("连接数据库失败%s", err)
		//do something
		return
	}
	fmt.Println("连接数据库成功")
	//defer是Go语言提供的一种用于注册延迟调用的机制：让函数或语句可以在当前函数执行完毕后（包括通过return正常结束或者panic导致的异常结束）执行。
	//defer语句通常用于一些成对操作的场景：打开连接/关闭连接；加锁/释放锁；打开文件/关闭文件等。
	//defer在一些需要回收资源的场景非常有用，可以很方便地在函数结束前做一些清理操作。
	//在打开资源语句的下一行，直接一句defer就可以在函数返回前关闭资源。
	defer db.Close()
	// step2 : 插入一行记录
	user := UserInfo{id: 1, Name: "hjf", Age: 18}
	err = db.Create(&user).Error
	if err != nil {
		fmt.Println("插入数据失败")
	}
	fmt.Println("插入数据成功")
	// step3 ：查询记录
	var tmpUser UserInfo
	err = db.Where("name = ?", "hjf").First(&tmpUser).Error //查询User并将信息保存到tmpUser
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("查询到数据%v\n", tmpUser)
	err = db.Delete(&user).Error
	if err != nil {
		fmt.Println("删除数据失败")
	}
	fmt.Println("删除数据成功")
}
