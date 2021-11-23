package datastructs

type KvConfig struct {
	Id             int    `mysqlfield:"id"`          // 自增PK
	Type           int    `mysqlfield:"type"`        // 类型。用于归组
	Description    string `mysqlfield:"description"` // 描述该配置的业务意义及value的类型等
	KeyName        string `mysqlfield:"key_name"`    // 配置键值
	ValueConfig    string `mysqlfield:"value_config"`
	ValueType      string `mysqlfield:"value_type"`      // 值的类型。例如List,
	Priority       int    `mysqlfield:"priority"`        // 同一个type的数据的排序优先级，值越小越优先，即值小的排前面
	IsDelete       int    `mysqlfield:"is_delete"`       // 是否已删除。0否，1是
	CreateTime     string `mysqlfield:"create_time"`     // 创建时间
	UpdateTime     string `mysqlfield:"update_time"`     // 创建时间
	CreateOperator string `mysqlfield:"create_operator"` // 创建者
	UpdateOperator string `mysqlfield:"update_operator"` // 更新者
	Version        int    `mysqlfield:"version"`         // 版本号
}
