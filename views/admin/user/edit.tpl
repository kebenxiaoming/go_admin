{{template "admin/public/header.tpl" .}}
{{template "admin/public/navibar.tpl" .}}
{{template "admin/public/sidebar.tpl" .}}

<!--&lt;!&ndash;- START 以上内容不需更改，保证该TPL页内的标签匹配即可 -&ndash;&gt;-->

{{template "admin/public/message.tpl" .}}

<form id="tab" method="post" action="">
    <div class="wzgl_czjl">

        <table width="855" border="0" cellspacing="0" cellpadding="0" class="right_nr_zhzgl">
            <tr>
                <td height="45" colspan="2" style="font-size:18px">修改用户资料</td>
            </tr>
            <input type="hidden" name="user_id" value="{{.user.User_id}}"/>
            <tr>
                <td width="117" height="53">用户名</td>
                <td width="738"><input name="user_name" type="text" required="true"  value="{{.user.User_name}}" class="right_nr_zhzgl_sr"></td>
            </tr>
            <tr>
                <td width="117" height="53">密码</td>
                <td width="738"><input name="password" type="text"  value="" class="right_nr_zhzgl_sr"><span>&nbsp;不填写则默认不修改</span></td>
            </tr>
            <tr>
                <td width="117" height="53">姓名</td>
                <td width="738"><input name="real_name" type="text" required="true"  value="{{.user.Real_name}}" class="right_nr_zhzgl_sr"></td>
            </tr>
            <tr>
                <td width="117" height="53">手机</td>
                <td width="738"><input name="mobile" type="text" required="true"  value="{{.user.Mobile}}" class="right_nr_zhzgl_sr"></td>
            </tr>
            <tr>
                <td width="117" height="53">邮箱</td>
                <td width="738"><input name="email" type="text" required="true"  value="{{.user.Email}}" class="right_nr_zhzgl_sr"></td>
            </tr>
            <tr>
                <td height="212">描述</td>
                <td><textarea name="user_desc" cols="" rows="" class="right_nr_zhzgl_xl">{{.user.User_desc}}</textarea></td>
            </tr>
            <tr>
                <td height="117">账号组</td>
                <td>
                    <div class="wz_qxgl_xl">
                        <select name="user_group" style="margin:5px 0px 0px">
                            {{range $index,$group:=.group_option_list}}
                                {{if eq $.group_id $group.Group_id}}
                                        <option value="{{$group.Group_id}}" selected>{{$group.Group_name}}</option>
                                {{else}}
                                        <option value="{{$group.Group_id}}">{{$group.Group_name}}</option>
                                {{end}}
                            {{end}}
                        </select>
                    </div>
                </td>
            </tr>
        </table>
</form>

<div class="dsjj_nr_menu">
    <input name="" type="submit" value="提交">
</div>


<!-- END 以下内容不需更改，请保证该TPL页内的标签匹配即可 -->
{{template "admin/public/footer.tpl" .}}