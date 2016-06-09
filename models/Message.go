package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

type Message struct {
	Id         int64
	Resplay    string
	Name       string
	Status     int64 `orm:"default(2)" form:"Status" valid:"Range(1)"`
	Statusname string
	User       string
	Addtime    time.Time `orm:"type(datetime);auto_now_add"`
	Statuszhi  string
}

func (r *Message) TableName() string {
	return beego.AppConfig.String("message_table")
}

func init() {
	orm.RegisterModel(new(Message))
}

func CheckMessage(r *Message) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&r)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func GetMessageList(page int64, page_size int64, sort string) (messages []orm.Params, count int64) {
	o := orm.NewOrm()
	message := new(Message)

	qs := o.QueryTable(message)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&messages)
	count, _ = qs.Count()
	return messages, count
}

func GetAllMassage() (messages []orm.Params, count int64) {
	o := orm.NewOrm()
	message := new(Message)
	qs := o.QueryTable(message)
	count, err := qs.Values(&messages)
	if err != nil {
		fmt.Println(err)
	}
	return messages, count
}

func UpdateMessage(r *Message) (int64, error) {
	if err := CheckMessage(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	message := make(orm.Params)
	if len(r.Name) > 0 {
		message["Name"] = r.Name
	}
	if r.Status != 0 {
		message["Status"] = r.Status
	}
	if len(r.Statusname) > 0 {
		message["Statusname"] = r.Statusname
	}
	if len(r.Statuszhi) > 0 {
		message["Statuszhi"] = r.Statuszhi
	}
	var table Message
	num, err := o.QueryTable(table).Filter("Id", r.Id).Update(message)

	if num <=0 {
		return num, err
	}
	if err != nil {
		fmt.Println("err is ", err)
	}
	fmt.Println("num is ", num)
	return num, err
}

func Allpre(r *Message) (int64, error) {
	if err := CheckMessage(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	message := make(orm.Params)
	if r.Status != 0 {
		message["Status"] = r.Status
	}
	var table Message
	num, err := o.QueryTable(table).Filter("Id", r.Id).Update(message)
	if err != nil {
		fmt.Println("err is ", err)
	}
	fmt.Println("num is ", num)
	return num, err
}



func AddMessage(r *Message) (int64, error) {
	o := orm.NewOrm()
	message := new(Message)
	message.Name = r.Name
	message.User = r.User
	message.Status = r.Status
	message.Statusname = r.Statusname
	message.Resplay = r.Resplay

	id, err := o.Insert(message)
	return id, err
}

func GetMessageByName(Name string) (err error, message Message) {
	o := orm.NewOrm()
	message = Message{Name: Name}
	err = o.Read(&message, "Name")
	return err, message
}

func DelMessageById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Message{Id: Id})
	return status, err
}

func GetMessageByMessageId(messageId int64) (messages Message) {
	fmt.Printf("Message id  is %d", messageId)
	o := orm.NewOrm()
	message := new(Message)
	err := o.QueryTable(message).Filter("Message__Message__Id", messageId).One(&messages)
	if err != nil {
		fmt.Println(err)
	}
	return messages
}
