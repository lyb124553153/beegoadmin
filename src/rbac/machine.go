package rbac

import (
	m "github.com/beego/admin/src/models"
	"github.com/astaxie/beego/logs"

	"encoding/json"
)

type MachineController struct {
	CommonController
}

func (this *MachineController) Index() {
	page, _ := this.GetInt64("page")
	page_size, _ := this.GetInt64("rows")
	sort := this.GetString("sort")
	order := this.GetString("order")
	if len(order) > 0 {
		if order == "desc" {
			sort = "-" + sort
		}
	} else {
		sort = "Id"
	}
	machineList, count := m.GetMachinelist(page, page_size, sort)
	if this.IsAjax() {
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &machineList}
		this.ServeJSON()
		return
	} else {
		tree := this.GetTree()
		this.Data["tree"] = &tree
		this.Data["machine"] = &machineList
		if this.GetTemplatetype() != "easyui" {
			this.Layout = this.GetTemplatetype() + "/public/layout.tpl"
		}
		this.TplName = this.GetTemplatetype() + "/rbac/machine.tpl"
	}

}

func (this *MachineController) AddMachine() {
	u := m.Machine{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.InsertMachine(&u)
	if err == nil && id > 0 {
		log := logs.NewLogger(10000)
		log.SetLogger("file", `{"filename":"machine.log"}`)
		ujon,_ := json.Marshal(u)
		log.Info("注册设备",string(ujon))
		log.Close()
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}




func (this *MachineController) DelMachine() {
	Id, _ := this.GetInt64("Id")
	delm :=   m.GetMachineById(Id);
	status, err:= m.DeleteMachineById(Id)
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"machine.log"}`)
	machininfo,_ := json.Marshal(delm)
	log.Info("尝试删除设备",string(machininfo))
	this.Rsp(true, "Success")

	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		log.Info("删除设备成功")
		return
	} else {
		log.Warn(err.Error())
		this.Rsp(false, err.Error())
		return
	}
	log.Close()
}

func(this *MachineController) UpdateMachine(){
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"machine.log"}`)

	machine := m.Machine{}

	if err  := this.ParseForm(&machine); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}

	Id, _ := this.GetInt64("Id")
	UpdateM :=   m.GetMachineById(Id);
	updateM,_ := json.Marshal(UpdateM)
	log.Info("原设备信息",string(updateM))
	machininfo,_ := json.Marshal(machine)
	log.Info("更新信息",string(machininfo))

	id, err := m.UpdateMachine(&machine)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		log.Info("更新成功")
		return
	} else {
		this.Rsp(false, err.Error())
		log.Warn("更新失败")
		log.Warn(err.Error())
		return
	}
	log.Close()
}