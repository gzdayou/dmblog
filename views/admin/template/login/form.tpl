<!DOCTYPE HTML>
<html class="no-js">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="renderer" content="webkit">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>登录到Hello World - Hello World - Powered by dmblog</title>
        <meta name="robots" content="noindex, nofollow">
        <link rel="stylesheet" href="/adresource/css/normalize.css?v=17.10.30">
        <link rel="stylesheet" href="/adresource/css/grid.css?v=17.10.30">
        <link rel="stylesheet" href="/adresource/css/style.css?v=17.10.30">
        <!--[if lt IE 9]>
        <script src="/adresource/js/html5shiv.js?v=17.10.30"></script>
        <script src="/adresource/js/respond.js?v=17.10.30"></script>
        <![endif]-->    
    </head>
    <body class="body-100">
    <!--[if lt IE 9]>
        <div class="message error browsehappy" role="dialog">当前网页 <strong>不支持</strong> 你正在使用的浏览器. 为了正常的访问, 请 <a href="http://browsehappy.com/">升级你的浏览器</a>.</div>
    <![endif]-->
<div class="typecho-login-wrap">
    <div class="typecho-login">
        <h1><a href="http://typecho.org" class="i-logo">Typecho</a></h1>
        <form action="/dologin" method="post" name="login" role="form">
            <p>
                <label for="name" class="sr-only">用户名</label>
                <input type="text" id="name" name="name" value="" placeholder="用户名" class="text-l w-100" autofocus />
            </p>
            <p>
                <label for="password" class="sr-only">密码</label>
                <input type="password" id="password" name="password" class="text-l w-100" placeholder="密码" />
            </p>
            <p class="submit">
                <button type="button" id="loginButton" class="btn btn-l w-100 primary">登录</button>
                <input type="hidden" name="referer" value="" />
            </p>
        </form>
        
        <p class="more-link">
            <a href="/">返回首页</a>
        </p>
    </div>
</div>
<script src="/adresource/js/jquery.js?v=17.10.30"></script>
<script src="/adresource/js/jquery-ui.js?v=17.10.30"></script>
<script src="/adresource/js/typecho.js?v=17.10.30"></script>
<script src="/adresource/js/jquery.cookie.js"></script>
<script>
function base64_decode(str){ var c1, c2, c3, c4; var base64DecodeChars = new Array( -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 62, -1, -1, -1, 63, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -1, -1, -1, -1, -1, -1, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, -1, -1, -1, -1, -1, -1, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, -1, -1, -1, -1, -1 ); var i=0, len = str.length, string = ''; while (i < len){ do{ c1 = base64DecodeChars[str.charCodeAt(i++) & 0xff] } while ( i < len && c1 == -1 ); if (c1 == -1) break; do{ c2 = base64DecodeChars[str.charCodeAt(i++) & 0xff] } while ( i < len && c2 == -1 ); if (c2 == -1) break; string += String.fromCharCode((c1 << 2) | ((c2 & 0x30) >> 4)); do{ c3 = str.charCodeAt(i++) & 0xff; if (c3 == 61) return string; c3 = base64DecodeChars[c3] } while ( i < len && c3 == -1 ); if (c3 == -1) break; string += String.fromCharCode(((c2 & 0XF) << 4) | ((c3 & 0x3C) >> 2)); do{ c4 = str.charCodeAt(i++) & 0xff; if (c4 == 61) return string; c4 = base64DecodeChars[c4] } while ( i < len && c4 == -1 ); if (c4 == -1) break; string += String.fromCharCode(((c3 & 0x03) << 6) | c4) } return string; }

$(document).ready(function () {
    $('#name').focus();

    $("#loginButton").click(function(){

        var un = $("#name").val();
        var up = $("#password").val();

        if( un == "" || up == "" ) {
            alert("用户名或密码为空");
            return false;
        }

        var xsrf, xsrflist;
        xsrf = $.cookie("_xsrf");
        xsrflist = xsrf.split("|");
        //args._xsrf = base64_decode(xsrflist[0]);

        $.ajax({
            type: 'POST',
            url: '/dologin',
            data: {name : un, password : up, _xsrf : base64_decode(xsrflist[0])},
            dataType: "json",
            success: function (t) {
                if( t.code == 1 ) {
                    window.location.href = "/AddArticle";
                    return ;
                } else {
                    alert("账号错误");
                    return false;
                }
            }
        });
    });
    
});
</script>
    </body>
</html>
