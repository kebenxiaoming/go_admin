<div id="left">
	<div class="left_tb"><img src="/static/admin/images/sy_07.jpg" width="248" height="154" /></div>
	<div class="left_nr">
		<div class="sideMenu">
			<ul>
				{{range $index,$module:=.sidebar}}
				{{if $module.MenuList}}
				{{if eq $module.Module_id $.current_module_id}}
				<li class="nLi on">
					<h3 style="background-color:#541b86">{{$module.Module_name}}</h3>
					<ul class="sub">
						{{else}}
						<li class="nLi">
							<h3>{{$module.Module_name}}</h3>
							<ul class="sub">
								{{end}}
								{{range $menu_k,$menu_list:=$module.MenuList}}
								{{if eq  $.content_header.Menu_id $menu_list.Menu_id}}
								<li style="background-color:#412a54"><a href="{{$menu_list.Menu_url}}">{{$menu_list.Menu_name}}</a></li>
								{{else}}
								<li><a href="{{$menu_list.Menu_url}}">{{$menu_list.Menu_name}}</a></li>
								{{end}}
								{{end}}
							</ul>
						</li>
						{{end}}
						{{end}}
					</ul>
		</div>
	</div>
</div>
<!--- 以上为左侧菜单栏 sidebar --->

<div id="right">
	<div class="right_nr">
		<table width="855" border="0" cellspacing="0" cellpadding="0" class="right_nr_bt">
			<tr>
				<td width="151" height="50" align="center" style="background-color:#541b86">{{.content_header.Menu_name}}</td>
				<td width="574">
					{{if  eq .content_header.Menu_id 11}}
				</td>
				<td width="130"><div class="cp_xz"><a href="{:url('Module/add')}">添加模块</a></div></td>
				{{else if eq .content_header.Menu_id 2}}

				</td>
				<td width="130"><div class="cp_xz"><a href="{:url('User/add')}">添加用户</a></div></td>
				{{else if eq .content_header.Menu_id 14}}

				</td>
				<td width="130"><div class="cp_xz"><a href="{:url('Menu/add')}">添加菜单</a></div></td>
				{{else if eq .content_header.Menu_id 7}}

				</td>
				<td width="130"><div class="cp_xz"><a href="{:url('Group/add')}">添加账号组</a></div></td>
				{{end}}
			</tr>
		</table>
