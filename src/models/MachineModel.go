package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"time"

)

type Machine struct {
	Id            int64
	MachineNo   string  `orm:"size(100)" form:"MachineNo"  valid:"Required"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	SiteName  string  `orm:"size(100)" form:"SiteName"  valid:"Required"`
	SiteUrl  string  `orm:"size(100)" form:"SiteUrl"  valid:"Required"`
	ChargeName  string  `orm:"size(100)" form:"ChargeName"  valid:"Required"`
	ContactWay  string  `orm:"size(100)" form:"ContactWay"  valid:"Required"`
	Registtime    time.Time `orm:"type(datetime);auto_now_add" `

}

func (m *Machine) TableName() string {
	return beego.AppConfig.String("machine_table")
}

func init() {
	orm.RegisterModel(new(Machine))
}


func GetMachinelist(page int64, page_size int64, sort string) (machines []orm.Params, count int64) {
	o := orm.NewOrm()
	machine := new(Machine)
	qs := o.QueryTable(machine)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&machines)
	count, _ = qs.Count()
	return machines, count
}


//添加设备
func InsertMachine(m *Machine) (int64, error) {
	//TODO 设备验证
	/*if err := checkMachine(m); err != nil {
		return 0, err
	}*/
	o := orm.NewOrm()
	machine := new(Machine)
	machine.MachineNo = m.MachineNo
	machine.Status = m.Status
	machine.SiteName = m.SiteName
	machine.ChargeName = m.ChargeName
	machine.SiteUrl = m.SiteUrl
	machine.ContactWay = m.ContactWay
	id,err :=  o.Insert(machine)

	return id, err
}


func DeleteMachineById(Id int64)(int64,error){
	o := orm.NewOrm()


	status, err := o.Delete(&Machine{Id:Id})

	return status, err
}
//修改状态
func UpdateMachine(m *Machine) (int64, error) {
	/*if err := checkUser(u); err != nil {
		return 0, err
	}*/
	o := orm.NewOrm()
	machine := make(orm.Params)
	if len(m.MachineNo) > 0 {
		machine["MachineNo"] = m.MachineNo
	}
	if len(m.SiteName) > 0{
		machine["SiteName"] = m.SiteName
	}
	if len(m.SiteUrl) > 0 {
		machine["SiteUrl"] = m.SiteUrl
	}
	if len(m.ChargeName) >0 {
		machine["ChargeName"] = m.ChargeName
	}
	if len(m.ContactWay) >0{
		machine["ContactWay"] = m.ContactWay
	}
	if m.Status != 0 {
		machine["Status"] = m.Status
	}
	if len(machine) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Machine
	num, err := o.QueryTable(table).Filter("Id", m.Id).Update(machine)
	return num, err
}


func 	GetMachineById(id int64) (m Machine) {
	m  = Machine{Id: id}
	o := orm.NewOrm()
	o.Read(&m, "Id")
	return m
}