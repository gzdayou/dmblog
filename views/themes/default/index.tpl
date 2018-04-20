{{template "themes/default/header.tpl" .}}

<main>
    <div class="wrap min">
        <section class="home-posts">
            {{range .list}}
            <div class="post-item">
                <h2>
                    <a href="/article/{{.Cid}}">{{.Title}}</a>
                </h2>
                <p>{{strreplace .Text "<!--markdown-->" "" | toMarkdown | htmlunquote | trimHTML | subList}}。。。</p>
                <div class="post-meta">
                    <time class="date"> {{stampToDatetime .Created}}</time>
                    {{getCatHTML .Cid | str2html}}
                    <span class="comments">{{.Views}} °C</span>
                </div>
            </div>
            {{end}}
        </section>
    </div>
</main>

{{template "themes/default/footer.tpl" .}}