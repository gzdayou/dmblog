{{template "themes/default/header.tpl" .}}

<main>
    <div class="wrap min">
        <section class="home-title">
            <h1>Hello World</h1>
            <span>Just So So ...</span>
        </section>
        <section class="home-posts">
            {{range .list}}
            <div class="post-item">
                <h2>
                    <a href="/article/{{.Cid}}">{{.Title}}</a>
                </h2>
                <p>{{strreplace .Text "<!--markdown-->" ""}}</p>
                <div class="post-meta">
                    <time class="date"> {{stampToDatetime .Created}}</time>
                    <span class="category"><a href="http://localhost/typecho/index.php/category/default/">默认分类</a></span>
                    <span class="comments">{{.Views}} °C</span>
                </div>
            </div>
            {{end}}
        </section>
    </div>
</main>

{{template "themes/default/footer.tpl" .}}