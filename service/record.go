package service

import (
	"chat/model"
	"encoding/base64"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type RecordService struct{}

//持久化聊天记录
func (s *RecordService) Record(ownerid, dstobj int64, cate int, text []byte) (err error) {
	tmp := model.Record{}
	//model.DbEngin.Where("id = ?", userId).Get(&tmp)

	encoded := base64.StdEncoding.EncodeToString(text)
	tmp.Ownerid = ownerid
	tmp.Dstobj = dstobj
	tmp.Cate = cate
	tmp.Text = encoded
	tmp.Createat = time.Now()

	_, err = DbEngin.InsertOne(&tmp)
	//数据库连接操作失败
	return err
}
