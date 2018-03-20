<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>曾祥杰的博客</title>
    <link rel="icon" href="/themepth/img/icon/32.png" sizes="32x32"/>
    <link rel="icon" href="/themepth/img/icon/192.png" sizes="192x192"/>
    <link href="/themepth/css/single.css" rel="stylesheet" type="text/css"/>
    <link href="/themepth/css/font-awesome.min.css" rel="stylesheet" type="text/css"/>
    <meta name="viewport" content="width=device-width, maximum-scale=1, initial-scale=1"/>
    <meta property="og:site_name" content="曾祥杰的博客">
    <meta property="og:title" content=""/>
<meta name="description" content="曾祥杰的个人博客" />
<meta name="keywords" content="go语言博客" />
</head>
<body>
<header>
    <div class="head-title">
        <h4>Hello World</h4>
    </div>
    <div class="toggle-btn"></div>
    <div class="light-btn"></div>
    <div class="search-btn"></div>
    <form class="search-form" method="post" action="">
        <input type="text" name="s" placeholder="搜索什么？">
    </form>
    <ul class="head-menu">
        <li><a href="/">首页</a></li>
        <li class="has-child">
            <a>分类</a>
            {{str2html getHeaderCatlist}}
        </li>
        <li><a href="/start-page.html">关于</a></li>    </ul>
</header> 