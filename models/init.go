// Copyright 2016 GoCms Author. All Rights Reserved.
// Author Mofree<mofree@mofree.org>
// Licensed under the Apache License, Version 2.0 (the "License");
// license that can be found in the LICENSE file.

package models

import (
	"github.com/astaxie/beego"
)

func TableName(name string) string {
	return beego.AppConfig.String("mysqlpre") + name
}
