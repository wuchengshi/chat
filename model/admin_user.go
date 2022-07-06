package model

type AdminUser struct {
	//用户ID
	Id      int64  `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Account string `xorm:"varchar(20)" form:"account" json:"account"`
	Type    string `xorm:"varchar(5)" form:"type" json:"type"` // 什么角色
}
