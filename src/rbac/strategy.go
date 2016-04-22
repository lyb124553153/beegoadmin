package rbac

import (
	"github.com/astaxie/beego/orm"
	m "github.com/beego/admin/src/models"

	"strings"
	"strconv"
	"github.com/astaxie/beego"
)

type StrategyController struct {
	CommonController
}

func (this *StrategyController) Index() {
	if this.IsAjax() {
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
		roles, count := m.GetStrategylist(page, page_size, sort)
		if len(roles) < 1 {
			roles = []orm.Params{}
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &roles}
		this.ServeJSON()
		return
	} else {
		this.TplName = this.GetTemplatetype() + "/rbac/strategy.html"
	}

}

func (this *StrategyController) AddAndEdit() {
	s := m.Strategy{}
	if err := this.ParseForm(&s); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	var id int64
	var err error
	Sid, _ := this.GetInt64("Id")
	beego.Debug(s)
	beego.Debug(&s)
	beego.Debug(this);
	if Sid > 0 {
		id, err = m.UpdateStrategy(&s)
	} else {
		id, err = m.AddStrategy(&s)
	}
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *StrategyController) DelStrategy() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelStrategyById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}

func (this *StrategyController) Getlist() {
	s, _ := m.GetStrategylist(1, 1000, "Id")
	if len(s) < 1 {
		s = []orm.Params{}
	}
	this.Data["json"] = &s
	this.ServeJSON()
	return
}


func (this *StrategyController) StrategyToMachineList() {
	strategyid, _ := this.GetInt64("Id");
	isajax,_ := this.GetInt64("Isajax");
	if this.IsAjax() {
		machines, count := m.GetMachinelist(1, 1000, "Id")
		list, _ := m.GetMachineByTrategyId(strategyid)
		for i := 0; i < len(machines); i++ {
			for x := 0; x < len(list); x++ {
				if machines[i]["Id"] == list[x]["Id"] {
					machines[i]["checked"] = 1
				}
			}
		}
		if len(machines) < 1 {
			machines = []orm.Params{}
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &machines}
		this.ServeJSON()
		return
	}else if isajax == 1{
		machines, count := m.GetMachinelist(1, 1000, "Id")
		list, _ := m.GetMachineByTrategyId(strategyid)
		for i := 0; i < len(machines); i++ {
			for x := 0; x < len(list); x++ {
				if machines[i]["Id"] == list[x]["Id"] {
					machines[i]["checked"] = 1
				}
			}
		}
		if len(machines) < 1 {
			machines = []orm.Params{}
		}
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &machines}
		this.ServeJSON()
		return

	} else {
		this.Data["strategyid"] = strategyid
		this.TplName = this.GetTemplatetype() + "/rbac/strategytomachine.html"
	}
}

func (this *StrategyController) AddStrategyToMachine() {
	strategyid, _ := this.GetInt64("Id")
	ids := this.GetString("ids")
	machineids := strings.Split(ids, ",")
	err := m.DelMachineStrategy(strategyid)
	if err != nil {
		this.Rsp(false, err.Error())
	}
	if len(ids) > 0 {
		for _, v := range machineids {
			id, _ := strconv.Atoi(v)
			_, err := m.AddStrategyMachine(strategyid, int64(id))
			if err != nil {
				this.Rsp(false, err.Error())
			}
		}
	}
	this.Rsp(true, "success")
}