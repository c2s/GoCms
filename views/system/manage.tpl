<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
  <script src="/static/js/jquery.nestable.js"></script>
<link href="/static/css/table-responsive.css" rel="stylesheet">
<link href="/static/css/nestable.css" rel="stylesheet">
</head>

<body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 模块管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">首页</a> </li>
        <li> <a href="/module/manage">模块管理</a> </li>
        <li class="active"> 模块 </li>
      </ul>
      <div class="pull-right"><a href="/module/add" class="btn btn-success">+新模块</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 模块 / 总数：{{.countProject}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="project-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>模块名称</th>
                        <th>模块英文名</th>
						<th>访问地址</th>
						<th>图标</th>
                        <th>排序</th>
                        <th>状态</th>
                        <th>创建时间</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    
                    {{range $k,$v := .modules}}
                    <tr>
                      <td>{{$v.Name}}</td>
                      <td>{{$v.Ename}}</td>
                      <td>{{$v.Url}}</td>
                      <td>{{$v.Icons}}</td>
                      <td>{{$v.Sort}}</td>
                      <td>{{$v.Status}}</td>
                      <td>{{getDate $v.Created}}</td>
                      <td><div class="btn-group">
                          <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                          <ul class="dropdown-menu">
                            <li><a href="/module/edit/{{$v.Id}}">编辑</a></li>
                            <li role="separator" class="divider"></li>
                            <li><a href="javascript:;" class="js-project-single" data-id="{{$v.Id}}" data-status="1">删除</a></li>
                          </ul>
                        </div>
                        &nbsp;&nbsp;<a href="/module/manage?pid={{$v.Id}}">子模块</a>
                      </td>
                    </tr>
                    {{end}}
                    </tbody>
                    
                  </table>
                </form>
                {{template "inc/page.tpl" .}}
				 </section>
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
</body>
</html>


<script>
  $(document).ready(function(){
    $('.dd').nestable();
    update_out('#list2',"#reorder");

    $('#list2').on('change', function() {
      var out = $('#list2').nestable('serialize');
      $('#reorder').val(JSON.stringify(out));

    });
    $('.ext-link').hide();

    $('.menutype input:radio').on('ifClicked', function() {
      val = $(this).val();
      mType(val);

    });

    mType('<?php echo $row['menu_type'];?>');


  });

  function mType( val )
  {
    if(val == 'external') {
      $('.ext-link').show();
      $('.int-link').hide();
    } else {
      $('.ext-link').hide();
      $('.int-link').show();
    }
  }


  function update_out(selector, sel2){

    var out = $(selector).nestable('serialize');
    $(sel2).val(JSON.stringify(out));

  }
</script>
