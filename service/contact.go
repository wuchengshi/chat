package service

import (
	"chat/model"
	"encoding/base64"
	"errors"
	"log"
	"time"
)

type ContactService struct {
}

//自动添加好友
func (service *ContactService) AddFriend(userid, dstid int64) error {
	//如果加自己
	if userid == dstid {
		return errors.New("不能添加自己为好友啊")
	}
	//判断是否已经加了好友
	tmp := model.Contact{}
	//查询是否已经是好友
	// 条件的链式操作
	DbEngin.Where("ownerid = ?", userid).
		And("dstobj = ?", dstid).
		And("cate = ?", model.CONCAT_CATE_USER).
		Get(&tmp)
	//获得1条记录
	//count()
	//如果存在记录说明已经是好友了不加
	if tmp.Id > 0 {
		return errors.New("该用户已经被添加过啦")
	}
	//事务,
	session := DbEngin.NewSession()
	session.Begin()
	//插自己的
	_, e2 := session.InsertOne(model.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//插对方的
	_, e3 := session.InsertOne(model.Contact{
		Ownerid:  dstid,
		Dstobj:   userid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//没有错误
	if e2 == nil && e3 == nil {
		//提交
		session.Commit()
		return nil
	} else {
		//回滚
		session.Rollback()
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}
}

func (service *ContactService) SearchComunity(userId int64) []model.Community {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	_ = DbEngin.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	_ = DbEngin.In("id", comIds).Find(&coms)

	for k, v := range coms {
		decoded, _ := base64.StdEncoding.DecodeString(string(v.Name))
		coms[k].Name = string(decoded)
	}

	return coms
}
func (service *ContactService) SearchComunityIds(userId int64) (comIds []int64) {
	//todo 获取用户全部群ID
	conconts := make([]model.Contact, 0)
	comIds = make([]int64, 0)

	_ = DbEngin.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	return comIds
}

//加群
func (service *ContactService) JoinCommunity(userId, comId int64) error {
	cot := model.Contact{
		Ownerid: userId,
		Dstobj:  comId,
		Cate:    model.CONCAT_CATE_COMUNITY,
	}
	_, _ = DbEngin.Get(&cot)
	if cot.Id == 0 {
		cot.Createat = time.Now()
		_, err := DbEngin.InsertOne(cot)
		return err
	} else {
		return nil
	}

}

//建群
func (service *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.Ownerid == 0 {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	num, err := DbEngin.Count(&com)

	if num > 5 {
		err = errors.New("一个用户最多只能创建5个群")
		return com, err
	} else {
		comm.Createat = time.Now()
		comm.Name = base64.StdEncoding.EncodeToString([]byte(comm.Name))
		session := DbEngin.NewSession()
		session.Begin()
		_, err = session.InsertOne(&comm)
		if err != nil {
			session.Rollback()
			return com, err
		}
		_, err = session.InsertOne(
			model.Contact{
				Ownerid:  comm.Ownerid,
				Dstobj:   comm.Id,
				Cate:     model.CONCAT_CATE_COMUNITY,
				Createat: time.Now(),
			})
		if err != nil {
			session.Rollback()
		} else {
			session.Commit()
		}
		return comm, err
	}
}

//查找好友
func (service *ContactService) SearchFriend(userId int64) []model.AdminUser {
	conconts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DbEngin.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&conconts)
	for _, v := range conconts {
		objIds = append(objIds, v.Dstobj)
	}
	coms := make([]model.AdminUser, 0)
	if len(objIds) == 0 {
		return coms
	}
	DbEngin.In("id", objIds).Find(&coms)
	return coms
}

//查找全部好友
func (service *ContactService) SearchUsers(userId int64) []model.AdminUser {
	tmp := make([]model.AdminUser, 0)
	_ = DbEngin.Find(&tmp)
	return tmp
}

//查找聊天记录
func (service *ContactService) SearchRecord(userId, destid int64, cate int) []model.Record {
	tmp := make([]model.Record, 0)
	log.Println("catecate::::::", cate)
	if cate == 1 { //单聊
		_ = DbEngin.Where("ownerid = ? and dstobj = ? and cate = ?", userId, destid, 1).Or("ownerid = ? and dstobj = ? and cate = ?", destid, userId, 1).OrderBy(" createat asc ").Find(&tmp)
	} else if cate == 2 {
		_ = DbEngin.Where("dstobj = ? and cate = ?", destid, 2).OrderBy(" createat asc ").Find(&tmp)
	}
	return tmp
}
