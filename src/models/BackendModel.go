package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type Backend struct {
	Id            int64
	MachineNo   string  `orm:"size(64)" form:"MachineNo"  valid:"Required"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	SiteName  string  `orm:"size(100)" form:"SiteName"  valid:"Required"`
	SiteUrl  string  `orm:"size(100)" form:"SiteUrl"  valid:"Required"`
	ChargeName  string  `orm:"size(32)" form:"ChargeName"  valid:"Required"`
	ContactWay  string  `orm:"size(32)" form:"ContactWay"  valid:"Required"`
	Paths  string  `orm:";size(32)" form:"Paths"  valid:"Required"`
	Level  int     `orm:"default(1)" form:"Level"  valid:"Required"`
	Pid    int64   `form:"Pid"  valid:"Required"`
	Strategy  []*Strategy   `orm:"reverse(many)"`
        Machine  []*Machine   `orm:"reverse(many)"`
}

type BackRootPath struct {

	Id            int64
	MachineNo   string  `orm:"size(64)" form:"MachineNo"  valid:"Required"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	MachineName  string  `orm:"size(100)" form:"MachineName"  valid:"Required"`
	ChargeName  string  `orm:"size(32)" form:"ChargeName"  valid:"Required"`
	ContactWay  string  `orm:"size(32)" form:"ContactWay"  valid:"Required"`
	MachineInfo  string  `orm:";size(32)" form:"MachineInfo"  valid:"Required"`
	Level  int     `orm:"default(1)" form:"Level"  valid:"Required"`
	Pid    int64   `form:"Pid"  valid:"Required"`

}

func (b *Backend) TableName() string {
	return beego.AppConfig.String("backend_table")
}

func init() {
	orm.RegisterModel(new(Backend))
}


func AddBackend(m *Backend) (int64, error) {
	//TODO 设备验证
	/*if err := checkMachine(m); err != nil {
		return 0, err
	}*/
	o := orm.NewOrm()
	machine := new(Backend)
	//machine.MachineNo = m.MachineNo
	if  m.Status != 0 {
		machine.Status = m.Status
	}else{
		machine.Status = 2
	}
	machine.SiteName = m.SiteName
	machine.ChargeName = m.ChargeName
	machine.SiteUrl = m.SiteUrl
	machine.ContactWay = m.ContactWay
	machine.MachineNo = m.MachineNo
	machine.Pid = m.Pid
	machine.Level = m.Level
	machine.Paths = m.Paths

	id,err :=  o.Insert(machine)

	return id, err
}

func AddBackRootPath( r *BackRootPath)(int64,error){
	o := orm.NewOrm()
	machine := new(Backend)
	machine.MachineNo = r.MachineNo
	if  r.Status != 0 {
		machine.Status = r.Status
	}else{
		machine.Status = 2
	}
	machine.SiteName = r.MachineName
	machine.ChargeName = r.ChargeName
	machine.SiteUrl = r.MachineInfo
	machine.Pid = 0
	machine.Level = 0
	machine.ContactWay = r.ContactWay
	id,err := o.Insert(machine)
	return id,err
}

func GetBackendById(id int64) (b Backend ,error  error) {
	b  = Backend{Id: id}
	o := orm.NewOrm()
	err := o.Read(&b, "Id")
	return b ,err
}

func GetBackendByPid(Pid int64)(paths []orm.Params,count int64) {
	backend := new(Backend)
	o := orm.NewOrm()
	count, _ = o.QueryTable(backend).Filter("Pid", Pid).Values(&paths)
	return paths, count
}

//修改状态
func UpdateBackend(b *Backend) (int64, error) {

	o := orm.NewOrm()
	backend := make(orm.Params)
	if len(b.MachineNo) > 0 {
		backend["MachineNo"] = b.MachineNo
	}
	if len(b.SiteName) > 0{
		backend["SiteName"] = b.SiteName
	}
	if len(b.SiteUrl) > 0 {
		backend["SiteUrl"] = b.SiteUrl
	}
	if len(b.ChargeName) >0 {
		backend["ChargeName"] = b.ChargeName
	}
	if len(b.ContactWay) >0{
		backend["ContactWay"] = b.ContactWay
	}
	if b.Status != 0 {
		backend["Status"] = b.Status
	}
	if len(backend) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Backend
	num, err := o.QueryTable(table).Filter("Id", b.Id).Update(backend)
	return num, err
}

func DeleteBackendById(Id int64)(int64,error){
	o := orm.NewOrm()
	status, err := o.Delete(&Backend{Id:Id})
	return status, err
}

func DeleteBackendByPid(Pid int64)(int64,error){
	o := orm.NewOrm()
	backend := new(Backend)
	num, err := o.QueryTable(backend).Filter("Pid",Pid).Delete()
	fmt.Printf("删除结果%d 错误信息%s",num ,err )
	return num, err
}

func BindMachinetoBackend(machineid int64, backendid int64) (int64, error) {
	o := orm.NewOrm()
	backend := Backend{Id: backendid}
	machine := Machine{Id: machineid}
	m2m := o.QueryM2M(&backend, "Machine")

	num, err := m2m.Add(&machine)
	return num, err
}

func UnbindBackendMachine(backendid int64) error{
	o :=orm.NewOrm()
	_,err := o.QueryTable("machine_backends").Filter("backend_id",backendid).Delete()
	return  err
}

func GetBackendlist(page int64, page_size int64, sort string) (backends []orm.Params,count int64){
     o := orm.NewOrm()
     backend := new(Backend)
     qs := o.QueryTable(backend)
     var offset int64
     if page<=1{
	     offset = 0
     }else {
	     offset = (page -1)*page_size
     }
    qs.Limit(page_size,offset).OrderBy(sort).Values(&backends)
	count,_ = qs.Count()
	return  backends,count
}


func GetStrategyBackendId(Id int64) (s []orm.Params, count int64) {
	o := orm.NewOrm()
	strategy := new(Strategy)
	count, _ = o.QueryTable(strategy).Filter("backend_id",Id).Values(&s)

	return s, count
}

func GetMachineByBackendId(BackendId int64)(machines []orm.Params, count int64){
	o := orm.NewOrm()
	machine := new(Machine)
	count, _ = o.QueryTable(machine).Filter("Backend__Backend__Id", BackendId).Values(&machines)
	return machines, count
}


func GetStrategyByBackendId(BackendId int64) (s []orm.Params,count int64) {
	o := orm.NewOrm()
	strategy := new(Strategy)
	count,_ =o.QueryTable(strategy).Filter("backend_id", BackendId).Values(&s)
	return s ,count
}
