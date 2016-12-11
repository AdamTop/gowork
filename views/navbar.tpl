{{define "navbar"}}
	<nav class="navbar navbar-default navbar-fixed-top" role="navigation">
		<div class="container">
			<div class="navbar-header">
					<a class="navbar-brand" href="../">Bootstrap</a>
			</div>
			<ul class="nav navbar-nav">
				<li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
				<li {{if .Category}}class="active"{{end}}><a href="/category">分类</a></li>
				<li {{if .Topic}}class="active"{{end}}><a href="/topic">文章</a></li>
			</ul>
			
			<ul class="nav navbar-nav navbar-right">
				{{if .Islogin}}
					<li><a href="/login?exit=true">退出</a></li>
				{{else}}
					<li><a href="/login">管理员登陆</a></li>
				{{end}}	
			</ul>
		</div>
		
	</nav>
{{end}}