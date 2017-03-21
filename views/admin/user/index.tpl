
{{template "admin/public/header.tpl" .}}
{{template "admin/public/navibar.tpl" .}}
{{template "admin/public/sidebar.tpl" .}}

<!--&lt;!&ndash;- START 以上内容不需更改，保证该TPL页内的标签匹配即可 -&ndash;&gt;-->

{{template "admin/public/message.tpl" .}}
<style type="text/css">
    table tr{
        height:50px;
        text-align: left;
        border:1px solid white;
    }
</style>
<div class="btn-toolbar" style="margin-bottom:2px;">
    <a data-toggle="collapse" data-target="#search"  href="#" title= "检索"><button class="btn btn-primary" style="margin-left:5px"><i class="icon-search"></i></button></a>
</div>
    <div class="block">
        <div id="page-stats" class="block-body collapse in">
            <table width="855" border="0" cellspacing="0" cellpadding="0" class="cpgl_cplb_bt">
                <tr>
                    <td width="60" height="60" align="center">#</td>
                    <td width="111" align="center">用户名</td>
                    <td width="111" align="center">真名</td>
                    <td width="124" align="center">手机</td>
                    <td width="109" align="center">登录时间</td>
                    <td width="100" align="center">登录IP</td>
                    <td width="100" align="center">Group#</td>
                    <td width="100" align="center">操作</td>
                </tr>
            </table>
                {{ range $index,$user_info:=.user_infos}}
                    <table width="855" border="0" cellspacing="0" cellpadding="0" class="cpgl_cplb_nr">
                <tr>
                    <td width="60" align="center">{{$user_info.User_id}}</td>
                    <td width="111" align="center">{{$user_info.User_name}}</td>
                    <td width="124" align="center">{{$user_info.Real_name}}</td>
                    <td width="100" align="center">{{$user_info.Mobile}}</td>
                    <td width="100" align="center">{{$user_info.Login_time}}</td>
                    <td width="100" align="center">{{$user_info.Login_ip}}</td>
                    <td width="100" align="center">{{$user_info.User_group}}</td>
                    <td width="100" align="center">
                        <a href="{{url "User/edit/uid/"}}{{$user_info.User_id}}" title= "修改" >修改</a>
                        &nbsp;
                        {{if ne $user_info.User_id 1}}
                        <a data-toggle="modal" href="#myModal" title= "删除" >
                            <i class="icon-remove" goto="{{url "User/del/uid/"}}{{$user_info.User_id}}" >删除</i></a>
                       {{end}}
                    </td>
                </tr>
                        </table>
                {{end}}
            <!--- START 分页模板--->

            <!--- END --->
        </div>
    </div>

<!--- END 以下内容不需更改，请保证该TPL页内的标签匹配即可 --->
{{template "admin/public/footer.tpl" .}}