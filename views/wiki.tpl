 <!DOCTYPE html>
{{template "header"}}
<title>Wiki列表 - 水利设计CMS系统</title>
<style>
i#delete
{
color:#DC143C;
}
</style>
</head>

<body>
<div class="navbar navba-default navbar-fixed-top">
  <div class="container-fill">{{template "navbar" .}}</div>
</div>


<div class="container">
<div class="row">

<div class="col-md-8">
<div class="content">
  <div>
  <h3 class="pull-left">欢迎来到HydroCMS Wiki</h3>
  <ul class="nav nav-tabs pull-right">
    <li class="active"><a href="javascript:;">最近回复</a></li>
    <li><a href="/wikis/latest">最近发布</a></li>
    <li><a href="/wikis/no_reply">尚未回复</a></li>
    
  </ul>
  </div>
<div class="clearfix"></div>
  <dl class="topics">

{{range $index, $elem := .Wikis}}
     {{if lt $index 20}}
  <dd>
      <a href="" class="pull-left">
    <img class="img-rounded" src="" alt="">
    </a>
    <a class="badge pull-right" href="">{{.ReplyCount}}</a>
    <h4><a href="/wiki/view/{{.Id}}" class="title">{{.Title}}
    <span class="glyphicon glyphicon-pushpin"></span></a></h4> 
    <div class="space"></div>
    <div class="info">
    <a class="label label-info" href="/go/share">分享</a> •
    <a href=""><strong>{{.Author}}</strong></a> •
    <time datetime="{{dateformat .Updated "2006-01-02 T 15:04:05"}}" title="{{dateformat .Updated "2006-01-02 T 15:04:05"}}">{{dateformat .Updated "2006-01-02 T 15:04:05"}}</time> • 最后回复来自 <a href="">{{.ReplyLastUserName}}</a>
    </div>
    <div class="clear"></div>
  </dd>
{{end}}
    {{end}}

  </dl>
      <!-- <ul class="pager">
        <li class="number">1/178</li>
        <li class="next">
          <a href="/?p=2">下一页 →</a>
        </li>
      </ul> -->
      <div style="text-align:center;padding-left: 100px;margin-top: -24px;float: right;" class="pagination">
      {{if .paginator}}
        {{if gt .paginator.PageNums 1}}
        <ul class="pagination pagination-sm">
          {{if .paginator.HasPrev}}
             <li>
               <a href="{{.paginator.PageLinkFirst}}">首页</a>
             </li>
             <li>
               <a href="{{.paginator.PageLinkPrev}}">上一页</a>
             </li>
             {{else}}
             <li class="disabled">
               <a>首页</a>
             </li>
             <li class="disabled">
               <a>上一页</a>
             </li>
             {{end}}
                   {{range $index, $page := .paginator.Pages}}
             <li{{if $.paginator.IsActive .}} class="active"{{end}}>
               <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
             </li>
             {{end}}
                   {{if .paginator.HasNext}}
             <li>
               <a href="{{.paginator.PageLinkNext}}">下一页</a>
             </li>
             <li>
               <a href="{{.paginator.PageLinkLast}}">末页</a>
             </li>
             {{else}}
             <li class="disabled">
               <a>下一页</a>
             </li>
             <li class="disabled">
               <a>末页</a>
             </li>
             {{end}}
             <li class="disabled">
               <a>
                 共{{.paginator.Nums}}条数据 每页{{.paginator.PerPageNums}}条 当前{{.paginator.Page}}/{{.paginator.PageNums}}页
               </a>
             </li>
             <li>
               <input type="text" type="submit" id="p" name="p" placeholder="跳转页" style="width: 47      px;height: 30px;border: 1px solid #dddddd;border-left: 0px;border-radius: 0px 4px 4px        0px;text-align: center;"/>  
             </li>
        </ul>
      {{end}} 
      {{end}}
    </div>


  </div>
  </div>

<div class="col-md-4 sidebar">
  <div class="panel panel-default">
  <div class="panel-body">
    <h4><small>致力于打造HydroCMS交流和分享的地方</small></h4>
    <hr>

{{if .IsLogin}}
        <div align="center">
        <a href="/member/{{.Uname}}"><img class="gravatar img-rounded" src="http://77fkk5.com1.z0.glb.clouddn.com/avatar/dcb8d7fd628811e5b9bc525400020562.jpg?imageView2/2/w/48/h/48/q/100"></a>
        <h4><a href="/member/{{.Uname}}">{{.Uname}}</a><br><small></small></h4>
        <a href="/login?exit=true">退出</a>
        <div class="clearfix"></div>
        </div>

        <!-- <li class="dropdown">
        <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Uname}} <b class="caret"></b></a>
        <ul class="dropdown-menu">
          <li><a href="/user/getuserbyusername?username={{.Uname}}">用户资料</a></li>
          <li><a href="/category/viewbyuname?uname={{.Uname}}">项目列表</a></li>
          <li><a href="/topic/viewbyuname?uname={{.Uname}}">成果列表</a></li>
          <li class="divider"></li>
          <li><a href="/login?exit=true">退出</a></li>
        </ul>
        </li> -->

        {{else}}
        <div align="center">
          <a href="/regist" class="btn btn-info">注册</a>
          </div>
          <div class="cleanfix"></div>
          <div align="center">已注册用户：<a href="/login">登录</a></div>
<!--         <li>
          <a href="/login">登陆</a>
        </li> -->
        {{end}}

  </div>
  </div>
  <a href="/wiki/add" class="btn btn-success btn-lg btn-block" style="margin-bottom: 20px;"><i class="glyphicon glyphicon-edit"></i> 发表话题</a>
  <div class="panel panel-default">
  <div class="panel-heading">
    <h3 class="panel-title">友情社区</h3>
  </div>
  <ul class="list-group">
    <li class="list-group-item">
    <a href="http://studygolang.com" target="_blank" title="Golang中文社区(Go语言构建) | Go语言学习园地">
      Go语言中文网
    </a>
      </li>
      <li class="list-group-item">
    <a href="http://www.golangtc.com/" target="_blank" title="Golang中国社区">
      Golang中国社区
    </a>
      </li> 
  </ul>

  <div class="panel-footer">
    <a href="/link" class="pull-right">更多</a>
    <div class="clearfix"></div>
  </div>
  </div>
  <div class="panel panel-default">
  <div class="panel-heading">
    <h3 class="panel-title">社区状态</h3>
  </div>
  <table width="100%" class="status">
    <thead>
    <tr>
      <th>&nbsp;</th>
      <th></th>
    </tr>
    </thead>
    <tbody>
    <tr>
      <td class="status-label">注册会员</td>
      <td class="value">{{.UsersCounts}}</td>
    </tr>
    <tr>
      <td class="status-label">主题</td>
      <td class="value">{{.Length}}</td>
    </tr>
    <tr>
      <td class="status-label">回复</td>
      <td class="value">{{.ReplyCounts}}</td>
    </tr>
    </tbody>
  </table>
  </div>
</div>

</div>

<!-- </div> -->

  <!-- float: right;调整位置 -->
  

  <!--  <input type="hidden" id="p" name="p" value="2" />
  <button type="submit" class="btn btn-default" >第2页</button>
  -->
  <!-- </form>-->
  </div>



<footer class="footer">
    <div class="container">
      <ul class="footer-links">
          <li><a href="https://github.com/3xxx" target="_blank">GitHub</a></li>
        </ul>
      <p>Copyright © 2016 HydroCMS</p>
      <p>Website built by <a href="https://github.com/3xxx" target="_blank" rel="nofollow">3xxx</a>, powered by <a href="http://golang.org/" target="_blank" rel="nofollow">Golang</a>, <a href="http://www.mongodb.org/" target="_blank" rel="nofollow">MongoDB</a>.</p>
      <p>
        Build version go1.6.2.
        Based on <a target="_blank" href="http://getbootstrap.com/" rel="nofollow">Bootstrap</a>.
        Icons from <a target="_blank" href="http://glyphicons.com/" rel="nofollow">Glyphicons</a>.
        59ms.
      </p>

    </div>
  </footer>






<script type="text/javascript">
  $(document).ready(function() {
  $("table").tablesorter();
  // $("#ajax-append").click(function() {
  //    $.get("assets/ajax-content.html", function(html) {
  //     // append the "ajax'd" data to the table body
  //     $("table tbody").append(html);
  //     // let the plugin know that we made a update
  //     $("table").trigger("update");
  //     // set sorting column and direction, this will sort on the first and third column
  //     var sorting = [[2,1],[0,0]];
  //     // sort on the first column
  //     $("table").trigger("sorton",[sorting]);
  //   });
  //   return false;
  // });
});
</script>

</body>
</html> 