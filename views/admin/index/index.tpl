{{template "admin/public/header.tpl" .}}
{{template "admin/public/navibar.tpl" .}}
{{template "admin/public/sidebar.tpl" .}}

<!--&lt;!&ndash;- START 以上内容不需更改，保证该TPL页内的标签匹配即可 -&ndash;&gt;-->

{{template "admin/public/message.tpl" .}}
{{if eq .user_info.User_type 0}}
  <table width="855" border="0" cellspacing="0" cellpadding="0" class="right_nr_sy_bt">
    <tr>
      <td width="21"><img src="/static/admin/images/sy_06.jpg" width="4" height="18" /></td>
      <td width="834">快捷菜单</td>
    </tr>
  </table>
  <div class="right_nr_kj_cd"><a href="/admin/user/index">账号列表</a></div>
{{end}}
<table width="855" border="0" cellspacing="0" cellpadding="0" class="right_nr_sy_bt">
  <tr>
    <td width="21"><img src="/static/admin/images/sy_06.jpg" width="4" height="18" /></td>
    <td width="834">当前账号信息</td>
  </tr>
</table>
<table width="855" border="0" cellspacing="0" cellpadding="0" class="right_nr_zhxx">
  <tr>
    <td width="101" height="65" align="center" style="background-color:#f4f4f4">用户名 </td>
    <td width="103" align="center" style="background-color:#f4f4f4">真实姓名</td>
    <td width="128" align="center" style="background-color:#f4f4f4">手机号</td>
    <td width="194" align="center" style="background-color:#f4f4f4">Email</td>
    <td width="193" align="center" style="background-color:#f4f4f4">登录时间</td>
    <td width="136" align="center" style="background-color:#f4f4f4">登录IP</td>
  </tr>
  <tr>
    <td height="60" align="center">{{.user_info.User_name}}</td>
    <td align="center">{{.user_info.Real_name}}</td>
    <td align="center">{{.user_info.Mobile}}</td>
    <td align="center">{{.user_info.Email}}</td>
    <td align="center">{{.user_info.Login_time}}</td>
    <td align="center">{{.user_info.Login_ip}}{{.hehe}}</td>
  </tr>
</table>
</div>
</div>
<!--- END 以下内容不需更改，请保证该TPL页内的标签匹配即可 --->
{{template "admin/public/footer.tpl" .}}