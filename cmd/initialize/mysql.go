package initialize

import (
	"image/internal/global"
	"image/pkg/pMysql"
)

func InitMysql() {
	msq := global.GVA_CONFIG.Mysql
	mysqlList, err := pMysql.Client(msq)
	if err != nil {
		panic("mysql初始化失败")
	}

	system := global.GVA_CONFIG.System

	db, ok := mysqlList[system.DefaultDB]
	if !ok {
		panic("默认数据库尚未配置")
	}

	global.GVA_DB_List = mysqlList
	global.GVA_DB = db
}

func ReleaseMysql() {
	dbList := pMysql.Connects()
	for _, db := range dbList {
		if conn, err := db.DB(); err == nil {
			_ = conn.Close()
		}
	}
}
