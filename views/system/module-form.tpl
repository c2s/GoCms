<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 菜单管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">首页</a> </li>
        <li> <a href="/module/manage">菜单管理</a> </li>
        <li class="active"> 菜单 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="module-form">
                <header> <b>帐号信息</b> </header>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">菜单名称</label>
                  <div class="col-sm-10">
                    <input type="text" name="Name"  value="{{.module.Name}}" class="form-control" placeholder="请填写名称">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">菜单英文名称</label>
                  <div class="col-sm-10">
                    <input type="text" name="Ename"  value="{{.module.Ename}}" class="form-control" placeholder="请填写英文名称">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">访问地址</label>
                  <div class="col-sm-10">
                    <input type="text" name="Url"  value="{{.module.Url}}" class="form-control" placeholder="请填写访问地址">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">图标</label>
                  <div class="col-sm-10">
                    <input type="text" name="Icons"  value="{{.module.Icons}}" class="form-control" placeholder="请填写图标">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">排序</label>
                  <div class="col-sm-10">
                    <input type="text" name="Sort"  value="{{.module.Sort}}" class="form-control" placeholder="请填写排序">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" value="{{.module.Id}}">
                    <button type="submit" class="btn btn-primary">提 交</button>
                  </div>
                </div>
              </form>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script src="/static/js/jquery-ui-1.10.3.min.js"></script>
<script src="/static/js/datepicker-zh-CN.js"></script>
<script>
$(function(){
	$('#default-date-picker').datepicker('option', $.datepicker.regional['zh-CN']); 	
	$('#default-date-picker').datepicker({
        dateFormat: 'yy-mm-dd',
		changeMonth: true,
		changeYear: true,
		yearRange:'-60:+0'
    });
})
</script>
</body>
</html>
