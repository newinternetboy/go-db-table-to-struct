package main

import (
	"fmt"
	"log"

	"github.com/newinternetboy/go-db-struct-convert/converter"
)

func main() {
	tableName := "kv_config"
	prefix := ""
	enableJson := false
	packageName := "datastructs"
	tagKey := "mysqlfield"
	savePath := "./dao/datastructs/KvConfig.go"
	//staging 数据库连接
	dsn := ""
	//table to struct
	t2s := converter.NewTable2Struct()
	//配置
	t2s.Config(
		&converter.T2tConfig{
			RmTagIfUcFirsted: false,
			TagToLower:       true,
			UcFirstOnly:      false,
		},
	)

	//开始迁移
	err := t2s.Table(tableName).Prefix(prefix).EnableJsonTag(enableJson).PackageName(packageName).TagKey(tagKey).SavePath(savePath).Dsn(dsn).Run()
	if err != nil {
		log.Fatal("err:%v", err)
	}
	fmt.Println("imigrate success")
}
