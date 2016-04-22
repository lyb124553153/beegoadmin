package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
	"github.com/astaxie/beego"
)

type Strategy struct {
	Id   int64
	StrategyName string  `orm:"size(100)" form:"StrategyName"  valid:"Required"`
	TimeLimit string  `orm:"size(100)" form:"TimeLimit" valid:"Required"`
        Machine []*Machine `orm:"reverse(many)"`
}

func (s *Strategy) TableName() string {
	return beego.AppConfig.String("rbac_strategy_table")
}

func init() {
	orm.RegisterModel(new(Strategy))
}

func checkStrategy(s *Strategy) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&s)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

//get Strategy list
func GetStrategylist(page int64, page_size int64, sort string) (strategys []orm.Params, count int64) {
	o := orm.NewOrm()
	strategy := new(Strategy)
	qs := o.QueryTable(strategy)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&strategys)
	count, _ = qs.Count()
	return strategys, count
}

func AddStrategy(s *Strategy) (int64, error) {
	if err := checkStrategy(s); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	strategy := new(Strategy)
	strategy.StrategyName = s.StrategyName
	strategy.TimeLimit = s.TimeLimit

	id, err := o.Insert(s)
	return id, err
}


func UpdateStrategy(s *Strategy) (int64, error) {
	if err := checkStrategy(s); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	strategy := make(orm.Params)
	if len(s.StrategyName) > 0 {
		strategy["StrategyName"] = s.StrategyName
	}
	if len(s.TimeLimit) > 0 {
		strategy["TimeLimit"] = s.TimeLimit
	}
	if len(strategy) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Strategy
	num, err := o.QueryTable(table).Filter("Id", s.Id).Update(strategy)
	return num, err
}

func DelStrategyById(DelId int64)(int64,error){
	o := orm.NewOrm()


	status, err := o.Delete(&Strategy{Id:DelId})

	return status, err
}

func GetMachineByTrategyId(TrategyId int64)(machines []orm.Params, count int64){
	o := orm.NewOrm()
	machine := new(Machine)
	count, _ = o.QueryTable(machine).Filter("Strategy__Strategy__Id", TrategyId).Values(&machines)
	return machines, count
}

func DelMachineStrategy(strategyid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("machine_strategys").Filter("strategy_id", strategyid).Delete()
	return err
}
func AddStrategyMachine(strategyid int64, machineid int64) (int64, error) {
	o := orm.NewOrm()
	strategy := Strategy{Id: strategyid}
	machine := Machine{Id: machineid}
	m2m := o.QueryM2M(&machine, "Strategy")
	num, err := m2m.Add(&strategy)
	return num, err
}