package main

import (
	"fmt"
	"log"
	"time"

	"github.com/newinternetboy/go-db-struct-convert/converter"
)

func main() {
	prefix := ""
	enableJson := false
	packageName := "datastructs"
	tagKey := "mysqlfield"
	saveDir := "./"
	//staging 数据库连接
	dsn := "user_name:passwd@tcp(host:port)/database?charset=ut8"
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

	tableNames := []string{
		"table_name1",
		"table_name2",
	}

	for _, tableName := range tableNames {
		savePath := saveDir + "/" + tableName + ".go"
		//开始迁移
		err := t2s.Table(tableName).Prefix(prefix).EnableJsonTag(enableJson).PackageName(packageName).TagKey(tagKey).SavePath(savePath).Dsn(dsn).Run()
		if err != nil {
			log.Fatalf("err:%v", err)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("imigrate success")
}
