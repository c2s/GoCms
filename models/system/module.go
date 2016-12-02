package system

import (
	"GoCms/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type Module struct {
	Id      int64 `orm:"pk;column(id);"`
	Name    string
	Ename   string
	Pid     string
	Sort    int64
	Url     string
	Status     int
	Created int64
	Updated int64
	Type    int
}

func (this *Module) TableName() string {
	return models.TableName("system_module")
}
func init() {
	orm.RegisterModel(new(Module))
}

//项目列表
func ListModule(condArr map[string]string, page int, offset int) (num int64, err error, user []Module) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("system_module"))
	cond := orm.NewCondition()
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	if condArr["name"] != "" {
		cond = cond.AndCond(cond.And("name", condArr["name"]).Or("name", condArr["name"]))
	}
	start := (page - 1) * offset
	qs = qs.RelatedSel()

	var modules []Module
	qs = qs.OrderBy("created")
	num, err1 := qs.Limit(offset, start).All(&modules)
	return num, err1, modules
}

//统计数量
func CountModule(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("system_module"))
	log.Print(qs)
	qs = qs.RelatedSel()
	cond := orm.NewCondition()
	if condArr["name"] != "" {
		cond = cond.AndCond(cond.And("name", condArr["name"]).Or("name", condArr["name"]))
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}

//添加用户
func AddModule(updPro Module) error {
	o := orm.NewOrm()
	o.Using("fault")
	data := new(Module)

	data.Id = updPro.Id
	data.Name = updPro.Name
	data.Ename = updPro.Ename
	data.Sort = updPro.Sort
	data.Url = updPro.Url
	data.Status = 1
	data.Created = time.Now().Unix()
	_, err := o.Insert(data)
	return err
}


