{{template "themes/default/header.tpl" .}}
<link href="/themepth/js/highlight/styles/tomorrow-night-eighties.css" rel="stylesheet">  
<script src="/themepth/js/highlight/highlight.pack.js"></script>  
<script >hljs.initHighlightingOnLoad();</script>
<script type="text/javascript">
(function () {
    window.TypechoComment = {
        dom : function (id) {
            return document.getElementById(id);
        },
    
        create : function (tag, attr) {
            var el = document.createElement(tag);
        
            for (var key in attr) {
                el.setAttribute(key, attr[key]);
            }
        
            return el;
        },

        reply : function (cid, coid) {
            var comment = this.dom(cid), parent = comment.parentNode,
                response = this.dom('respond-post-{{.art.Cid}}'), input = this.dom('comment-parent'),
                form = 'form' == response.tagName ? response : response.getElementsByTagName('form')[0],
                textarea = response.getElementsByTagName('textarea')[0];

            if (null == input) {
                input = this.create('input', {
                    'type' : 'hidden',
                    'name' : 'parent',
                    'id'   : 'comment-parent'
                });

                form.appendChild(input);
            }

            input.setAttribute('value', coid);

            if (null == this.dom('comment-form-place-holder')) {
                var holder = this.create('div', {
                    'id' : 'comment-form-place-holder'
                });

                response.parentNode.insertBefore(holder, response);
            }

            comment.appendChild(response);
            this.dom('cancel-comment-reply-link').style.display = '';

            if (null != textarea && 'text' == textarea.name) {
                textarea.focus();
            }

            return false;
        },

        cancelReply : function () {
            var response = this.dom('respond-post-{{.art.Cid}}'),
            holder = this.dom('comment-form-place-holder'), input = this.dom('comment-parent');

            if (null != input) {
                input.parentNode.removeChild(input);
            }

            if (null == holder) {
                return true;
            }

            this.dom('cancel-comment-reply-link').style.display = 'none';
            holder.parentNode.insertBefore(response, holder);
            return false;
        }
    };
})();
</script>

<main>
    <div class="wrap min">
        <section class="post-title">
            <h2>{{.art.Title}}</h2>
            <div class="post-meta">
                <time class="date">{{stampToDatetime .art.Created}}</time>
                {{getCatHTML .art.Cid | str2html}}
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
            <li>下一篇: 没有更多啦 (つд⊂)</li>
            {{end}}
        </ul>
        <section id="comments" class="post-comments">
            <h3>评论</h3>
            <div id="respond-post-{{.art.Cid}}" class="respond">
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
                    {{str2html .xsrfdata }}
                    <input type="hidden" name="cid" value="{{.art.Cid}}">
                </form>
            </div>

            {{str2html .cmtlist}}

        </section>
    </div>
</main>

{{template "themes/default/footer.tpl" .}}