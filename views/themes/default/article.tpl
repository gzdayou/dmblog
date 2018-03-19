{{template "themes/default/header.tpl" .}}

<main>
    <div class="wrap min">
        <section class="post-title">
            <h2>{{.art.Title}}</h2>
            <div class="post-meta">
                <time class="date">{{stampToDatetime .art.Created}}</time>
                <span class="category"><a href="http://localhost/typecho/index.php/category/default/">默认分类</a></span>
                <span class="comments">{{.art.Views}} °C</span>
            </div>
        </section>
        <article class="post-content">
            
            {{str2html .text}}

        </article>
        <ul class="post-near">
            {{if ne .preCid 0}}
            <li>上一篇: <a href="/article/{{.preCid}}" title="{{.preTitle}}">{{.preTitle}}</a></li>
            {{else}}
            <li>上一篇: 没有上一篇啦 (つд⊂)</li>
            {{end}}
            {{if ne .nextCid 0}}
            <li>下一篇: <a href="/article/{{.nextCid}}" title="{{.nextTitle}}">{{.nextTitle}}</a></li>
            {{else}}
            <li>下一篇: 看完啦 (つд⊂)</li>
            {{end}}
        </ul>
        <section id="comments" class="post-comments">
            <h3>没有评论</h3>
            <div id="respond-post-3" class="respond">
                <span class="cancel-comment-reply">
                    <a id="cancel-comment-reply-link" href="/article/{{.art.Cid}}/#respond-post-{{.art.Cid}}" rel="nofollow" style="display:none" onclick="return TypechoComment.cancelReply();">取消回复</a>
                </span>
                <form class="bk-form" method="post" action="/AddComments" id="comment-form" role="form">
                    <div class="row">
                        <fieldset class="col-m-6">
                            <input type="text" name="author" id="author" placeholder="昵称 *：" value="" required="">
                            <input type="email" name="mail" id="mail" placeholder="电邮 *：" value="" required>
                            <input type="url" name="url" id="url" placeholder="http://" value="">
                        </fieldset>
                        <fieldset class="col-m-6">
                            <textarea rows="3" name="text" id="textarea" placeholder="快来评论吧 (*≧ω≦)ﾉ" required=""></textarea>
                            <button type="submit" class="btn small">写好了~</button>
                        </fieldset>
                    </div>
                    <input type="hidden" name="cid" value="{{.art.Cid}}">
                </form>
            </div>

            {{str2html .cmtlist}}

        </section>
    </div>
</main>

{{template "themes/default/footer.tpl" .}}