package datastructs

type WearGeneralActivityParticipator struct {
	Id                  int    `mysqlfield:"id"`            // 自增主键
	ActivityId          int    `mysqlfield:"activity_id"`   // 活动id
	UserId              int    `mysqlfield:"user_id"`       // 用户id
	UserInfo            string `mysqlfield:"user_info"`     // 用户信息（如：用户名、用户头像、用户性别）
	Status              int    `mysqlfield:"status"`        // 状态: 0-任务进行中、1-已完成、2-未完成、3-中奖、4-未中奖、5-已登记
	ActivityData        string `mysqlfield:"activity_data"` // 活动数据
	ActivityStartTime   int    `mysqlfield:"activity_start_time"`
	ActivityEndTime     int    `mysqlfield:"activity_end_time"`
	ActivityCategoryKey string `mysqlfield:"activity_category_key"`
	ActivityTarget      int    `mysqlfield:"activity_target"`
	CreateTime          int    `mysqlfield:"create_time"`     // 创建时间
	UpdateTime          int    `mysqlfield:"update_time"`     // 编辑时间
	IsDelete            int    `mysqlfield:"is_delete"`       // 是否已删除。0否，1是
	CreateOperator      string `mysqlfield:"create_operator"` // 创建者
	UpdateOperator      string `mysqlfield:"update_operator"` // 更新者
	Version             int    `mysqlfield:"version"`         // 版本号
}
