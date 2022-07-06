package model

import "time"

type Record struct {
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//谁的10000
	Ownerid int64 `xorm:"bigint(20)" form:"ownerid" json:"ownerid"` // 记录是谁的
	//对端,10001
	Dstobj   int64     `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`   // 对端信息
	Text     string    `xorm:"text" form:"text" json:"text"`             // 聊天记录
	Cate     int       `xorm:"int(11)" form:"cate" json:"cate"`          // 什么类型
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"` // 创建时间
}
