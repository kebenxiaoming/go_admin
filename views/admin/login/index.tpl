{{template "admin/public/simple_header.tpl" .}}
<body>
<div class="login1"></div>
<div class="login2">
    <div class="login2_nr">
        <form name="loginForm" method="post" action="">
            <table width="356" border="0" cellspacing="0" cellpadding="0" class="login2_nr_xx">
                <tr>
                    <td height="62" colspan="2" valign="top"><input type="text" name="user_name" value="{{.user_name}}" id="textfield1" placeholder="Username" class="login2_nr_sr"></td>
                </tr>
                <tr>
                    <td height="61" colspan="2" valign="top"><input type="password" name="password" id="textfield2" value = "{{.password}}" placeholder="Password" class="login2_nr_sr"></td>
                </tr>
                <tr>
                    <td width="177" height="53" valign="top"><input type="text" name="captcha" id="textfield" placeholder="输入验证码" class="login2_nr_yz"></td>
                    <td width="240" valign="top">{{create_captcha}}<a href="#"></a><span style="color:red;font-size:10px;">{{.err}}</span></td>
                </tr>
                <tr>
                    <td height="58" valign="top"><input type="checkbox" name="remember" value="remember-me" id="checkbox"  class="login2_nr_xz"></td>
                    <td align="right" valign="top"><input name="loginSubmit" type="submit" value="登 录"   class="login2_nr_dl"></td>
                </tr>
            </table>
        </form>
    </div>
</div>
<div class="login3"> Copyright © 1989 - 2017  Sunnier.All Rights Reserved.<br />
    sunnier 版权所有</div>
</body>
</html>


