<!DOCTYPE html>

<html>
<head>
  <title>便利店后台管理中心</title><meta charset="UTF-8" />
  <link rel="shortcut icon"  href="/static/img/favicon.ico">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css" />
  <link rel="stylesheet" type="text/css" href="/static/datatable/dataTables.bootstrap.css">

  <script src="http://cdn.staticfile.org/jquery/2.1.1-rc2/jquery.min.js"></script>
  <script src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript" language="javascript" src="/static/datatable/jquery.dataTables.min.js">
  </script>
  <script type="text/javascript" language="javascript" src="/static/datatable/dataTables.bootstrap.js"></script>  
</head>
<script type="text/javascript">
var URL="/cstore"
  
    function refreshTab(title) {
        var tab = $("#tabs").tabs("getTab", title);
        $("#tabs").tabs("update", {tab: tab, options: tab.panel("options")});
    }
    //选择分组
    function selectgroup(group_id){
        $(this).addClass("current");
        vac.ajax(URL+'/index', {group_id:group_id}, 'GET', function(data){
            $("#tree").tree("loadData",data)
        })

    }
</script>

<body >
  <header class="navbar navbar-static-top bs-docs-nav" id="top" role="banner">
  <div class="container" style="width: 100%; height:82px; background:#D9E5FD;">
     
    <nav class="collapse navbar-collapse bs-navbar-collapse" role="navigation">
         <div class="nav navbar-nav"  >
            <h2>后台管理系统</h2>
        </div>
        
         <ul class="nav navbar-nav navbar-right">
        <li><a>欢迎你！ {{.userinfo.Nickname}}</a> </li>
        <li><a href="javascript:void(0);" data-target="#myModal" data-toggle="modal" >修改密码</a></li>
        <li><a href="/cstore/logout" target="_parent">退 出</a></li>
      </ul>
    </nav>
  </div>
</header>

<div region="west" border="false" split="true" title="菜单"  tools="#toolbar" style="width:200px;padding:5px;">
    <ul id="tree"></ul>
</div>
<div region="center" border="false" >
    <div id="tabs" >
    </div>
</div>




<!--右键菜单-->
<div id="mm" style="width: 120px;display:none;">
    <div iconCls='icon-reload' type="refresh">刷新</div>
    <div class="menu-sep"></div>
    <div  type="close">关闭</div>
    <div type="closeOther">关闭其他</div>
    <div type="closeAll">关闭所有</div>
</div>




</body>
</html>