{{template "themes/default/header.tpl" .}}

<main>
    <div class="wrap min">
        <section class="post-title">
            <h2>未命名文档</h2>
            <div class="post-meta">
                <time class="date">2018.03.15</time>
                <span class="category"><a href="http://localhost/typecho/index.php/category/default/">默认分类</a></span>
                <span class="comments">0 °C</span>
            </div>
        </section>
        <article class="post-content">
            <p>asdfasdf</p><blockquote><p>引用文字</p></blockquote>        
        </article>
        <ul class="post-near">
            <li>上一篇: <a href="http://localhost/typecho/index.php/archives/1/" title="欢迎使用 Typecho">欢迎使用 Typecho</a></li>
            <li>下一篇: 看完啦 (つд⊂)</li>
        </ul>
        <section id="comments" class="post-comments">
            <h3>没有评论</h3>
            <div id="respond-post-3" class="respond">
            <span class="cancel-comment-reply">
                <a id="cancel-comment-reply-link" href="http://localhost/typecho/index.php/archives/3/#respond-post-3" rel="nofollow" style="display:none" onclick="return TypechoComment.cancelReply();">取消回复</a>        </span>
            <form class="bk-form" method="post" action="http://localhost/typecho/index.php/archives/3/comment" id="comment-form" role="form">
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
            </form>
        </div>
        </section>
    </div>
</main>

{{template "themes/default/footer.tpl" .}}