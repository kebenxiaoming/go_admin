{*
<div id="left">
		<div class="left_tb"><img src="/static/admin/images/sy_07.jpg" width="248" height="154" /></div>
		<div class="left_nr">
			<div class="sideMenu">
				<ul>
				{foreach name="sidebar" item="module"}
					{if condition="count($module['menu_list']) gt 0"}
						{if condition="$module['module_id'] eq $current_module_id"}
							<li class="nLi on">
								<h3 style="background-color:#541b86">{$module.module_name|default=""}</h3>
								<ul class="sub">
									{else /}
									<li class="nLi">
										<h3>{$module.module_name|default=""}</h3>
										<ul class="sub">
						{/if}
						{foreach name="module['menu_list']" item="menu_list"}
							{if condition="strtolower(substr($menu_list['menu_url'],0,4)) eq 'http'"}
								{if condition="$menu_list['menu_id'] eq $content_header['menu_id']"}
									<li style="background-color:#412a54"><a target=_blank href="<?php echo url($menu_list['menu_url']);?>">{$menu_list.menu_name}</a></li>
									{else /}
									<li><a target=_blank href="<?php echo url($menu_list['menu_url']);?>">{$menu_list.menu_name}</a></li>
								{/if}
								{else/}
								{if condition="$menu_list['menu_id'] eq $content_header['menu_id']"}
									<li style="background-color:#412a54"><a href="<?php echo url($menu_list['menu_url']);?>">{$menu_list.menu_name}</a></li>
									{else /}
									<li><a href="<?php echo url($menu_list['menu_url']);?>">{$menu_list.menu_name}</a></li>
								{/if}
							{/if}
						{/foreach}
				</ul>
				</li>
				{/if}
				{/foreach}
				</ul>
			</div>
		</div>
	</div>
	 <!--- 以上为左侧菜单栏 sidebar --->

	<div id="right">
		<div class="right_nr">
			<table width="855" border="0" cellspacing="0" cellpadding="0" class="right_nr_bt">
				<tr>
					<td width="151" height="50" align="center" style="background-color:#541b86">{$content_header.menu_name}</td>
					<td width="574">
					{if condition="$content_header['menu_id'] ==11"/}

					</td>
					<td width="130"><div class="cp_xz"><a href="{:url('Module/add')}">添加模块</a></div></td>
					{elseif condition="$content_header['menu_id'] ==2"/}

					</td>
					<td width="130"><div class="cp_xz"><a href="{:url('User/add')}">添加用户</a></div></td>
					{elseif condition="$content_header['menu_id'] ==14"/}

					</td>
					<td width="130"><div class="cp_xz"><a href="{:url('Menu/add')}">添加菜单</a></div></td>
					{elseif condition="$content_header['menu_id'] ==7"/}

					</td>
					<td width="130"><div class="cp_xz"><a href="{:url('Group/add')}">添加账号组</a></div></td>
					{/if}
				</tr>
			</table>*}
