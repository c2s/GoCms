// Copyright 2016 GoCms Author. All Rights Reserved.
// Author Mofree<mofree@mofree.org>
// Licensed under the Apache License, Version 2.0 (the "License");
// license that can be found in the LICENSE file.

package system

import (
	"GoCms/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
	"fmt"
	"ApiManager/utils"
)

type Module struct {
	Id       int `orm:"pk;column(id);"`
	Name     string
	Ename    string
	ParentId int
	Icons    string
	Sort     int64
	Url      string
	Status   int
	Created  int64
	Updated  int64
	Type     int
}

func (this *Module) TableName() string {
	return models.TableName("system_module")
}
func init() {
	orm.RegisterModel(new(Module))
}

// 项目列表
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
	// 0 表示父级菜单
	cond = cond.And("parent_id", 0)
	qs = qs.SetCond(cond)

	start := (page - 1) * offset
	qs = qs.RelatedSel()

	var modules []Module
	qs = qs.OrderBy("created")
	num, err1 := qs.Limit(offset, start).All(&modules)
	return num, err1, modules
}

// 统计数量
func CountModule(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("system_module"))
	log.Print(qs)
	qs = qs.RelatedSel()
	cond := orm.NewCondition()
	// 0 表示父级菜单
	cond = cond.And("parent_id", 0)
	num, _ := qs.SetCond(cond).Count()
	return num
}

// 添加用户
func AddModule(updPro Module) error {
	o := orm.NewOrm()
	o.Using("fault")
	data := new(Module)

	data.Name = updPro.Name
	data.Ename = updPro.Ename
	data.Sort = updPro.Sort
	data.Url = updPro.Url
	data.Status = 1
	data.Created = time.Now().Unix()
	_, err := o.Insert(data)
	return err
}


func UpdateModule(id int, updPro Module) error {
	var pro Module
	o := orm.NewOrm()
	pro = Module{Id: id}

	pro.Name = updPro.Name
	pro.Ename = updPro.Ename
	pro.Icons = updPro.Icons
	pro.Url = updPro.Url
	pro.Sort = updPro.Sort
	pro.Updated = time.Now().Unix()
	//pro.Status = updPro.Status
	_, err := o.Update(&pro, "name", "ename", "icons", "url", "sort", "updated")
	return err
}
// 获取菜单
func GetModule(id int) (Module, error) {
	var project Module
	var err error

	//err = utils.GetCache("GetProject.id."+fmt.Sprintf("%d", id), &project)
	//if err != nil {
	o := orm.NewOrm()
	project = Module{Id: id}
	err = o.Read(&project)
	//utils.SetCache("GetProject.id."+fmt.Sprintf("%d", id), project, 600)
	//}
	return project, err
}

func ListModuleTeam(projectId int, page int, offset int) (num int64, err error, ops []Module) {
	var teams []Module
	var errs error

	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	errs = utils.GetCache("Module.id."+fmt.Sprintf("%d", projectId), &teams)
	if errs != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		o := orm.NewOrm()
		o.Using("default")
		qs := o.QueryTable(models.TableName("system_module"))
		cond := orm.NewCondition()
		if projectId > 0 {
			cond = cond.And("id", projectId)
		}
		qs = qs.SetCond(cond)

		qs.Limit(offset, start).All(&teams)
		utils.SetCache("Module.id."+fmt.Sprintf("%d", projectId), teams, cache_expire)
	}
	return num, errs, teams
}




