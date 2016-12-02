package system

import (
	"GoCms/controllers"
	. "GoCms/models/system"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"fmt"
	"GoCms/utils"
	"log"
)

type MyModuleController struct {
	controllers.BaseController
}

type ModuleManageController struct {
	controllers.BaseController
}

func (this *ModuleManageController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "project-manage") {
		this.Redirect("/my/task", 302)
		return
		//this.Abort("401")
	}
	page, err := this.GetInt("p")
	name := this.GetString("keywords")
	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	condArr := make(map[string]string)
	condArr["name"] = name

	countProject := CountModule(condArr)
	paginator := pagination.SetPaginator(this.Ctx, offset, countProject)
	_, _, module := ListModule(condArr, page, offset)

	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["modules"] = module
	this.Data["countProject"] = countProject

	this.TplName = "system/manage.tpl"
}

//添加模块
type AddModuleController struct {
	controllers.BaseController
}

func (this *AddModuleController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "user-add") {
		this.Abort("401")
	}
	condArr := make(map[string]string)
	condArr["status"] = "1"
	this.TplName = "system/module-form.tpl"
}

func (this *AddModuleController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "user-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}

	Name := this.GetString("Name")
	if "" == Name {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写名称"}
		this.ServeJSON()
		return
	}

	Ename := this.GetString("Ename")
	if "" == Ename {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写英文名称"}
		this.ServeJSON()
		return
	}

	Sort, _ := this.GetInt64("Sort")
	if Sort <= 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "参数出错"}
		this.ServeJSON()
		return
	}

	Url := this.GetString("Url")
	if "" == Url {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写英文名称"}
		this.ServeJSON()
		return
	}
	var err error
	//雪花算法ID生成
	id := utils.SnowFlakeId()

	log.Print(id)

	var data Module
	data.Id  = id
	data.Name  = Name
	data.Ename = Ename
	data.Url   = Url
	data.Sort  = Sort

	err = AddModule(data)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "模块添加成功", "id": fmt.Sprintf("%d", id)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "模块添加失败"}
	}
	this.ServeJSON()
}