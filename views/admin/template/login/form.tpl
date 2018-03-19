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
<script>
$(document).ready(function () {
    $('#name').focus();

    $("#loginButton").click(function(){

        var un = $("#name").val();
        var up = $("#password").val();

        if( un == "" || up == "" ) {
            alert("用户名或密码为空");
            return false;
        }

        $.ajax({
            type: 'POST',
            url: '/dologin',
            data: {name : un, password : up},
            dataType: "json",
            success: function (t) {
                console.log(t)
            }
        });
    });
    
});
</script>
    </body>
</html>
