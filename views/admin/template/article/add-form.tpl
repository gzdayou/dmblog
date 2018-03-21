<div class="main">
    <div class="body container">
        <div class="typecho-page-title">
    <h2>撰写新文章</h2>
</div>
        <div class="row typecho-page-main typecho-post-area" role="form">
            <form action="/AddArticle" method="post" name="write_post">
                <div class="col-mb-12 col-tb-9" role="main">
                    
                    <p class="title">
                        <label for="title" class="sr-only">标题</label>
                        <input type="text" id="title" name="title" autocomplete="off" value="" placeholder="标题" class="w-100 text title" />
                    </p>
                                        <p class="mono url-slug">
                        <label for="slug" class="sr-only">网址缩略名</label>
                        localhost/typecho/index.php/archives/{cid}/                    </p>
                    <p>
                        <label for="text" class="sr-only">文章内容</label>
                        <textarea style="height: 350px" autocomplete="off" id="text" name="text" class="w-100 mono"></textarea>
                    </p>
                    
                                        <section id="custom-field" class="typecho-post-option fold">
                        <label id="custom-field-expand" class="typecho-label"><a href="##"><i class="i-caret-right"></i> 自定义字段</a></label>
                        <table class="typecho-list-table mono">
                            <colgroup>
                                <col width="25%"/>
                                <col width="10%"/>
                                <col width="55%"/>
                                <col width="10%"/>
                            </colgroup>
                                                                                                                <tr>
                                <td>
                                    <label for="fieldname" class="sr-only">字段名称</label>
                                    <input type="text" name="fieldNames[]" placeholder="字段名称" id="fieldname" class="text-s w-100">
                                </td>
                                <td>
                                    <label for="fieldtype" class="sr-only">字段类型</label>
                                    <select name="fieldTypes[]" id="fieldtype">
                                        <option value="str">字符</option>
                                        <option value="int">整数</option>
                                        <option value="float">小数</option>
                                    </select>
                                </td>
                                <td>
                                    <label for="fieldvalue" class="sr-only">字段值</label>
                                    <textarea name="fieldValues[]" placeholder="字段值" id="fieldvalue" class="text-s w-100" rows="2"></textarea>
                                </td>
                                <td>
                                    <button type="button" class="btn btn-xs">删除</button>
                                </td>
                            </tr>
                                                    </table>
                        <div class="description clearfix">
                            <button type="button" class="btn btn-xs operate-add">+添加字段</button>
                            自定义字段可以扩展你的模板功能, 使用方法参见 <a href="http://docs.typecho.org/help/custom-fields">帮助文档</a>                        </div>
                    </section>

                    <p class="submit clearfix">
                        <span class="right">
                            <input type="hidden" name="cid" value="" />
                            <button type="submit" name="do" value="save" id="btn-save" class="btn">保存草稿</button>
                            <button type="submit" name="do" value="publish" class="btn primary" id="btn-submit">发布文章</button>
                                                        <input type="hidden" name="markdown" value="1" />
                                                        {{str2html .xsrfdata }}
                                                    </span>
                    </p>

                                    </div>

                <div id="edit-secondary" class="col-mb-12 col-tb-3" role="complementary">
                    <ul class="typecho-option-tabs clearfix">
                        <li class="active w-50"><a href="#tab-advance">选项</a></li>
                        <li class="w-50"><a href="#tab-files" id="tab-files-btn">附件</a></li>
                    </ul>


                    <div id="tab-advance" class="tab-content">
                        <section class="typecho-post-option" role="application">
                            <label for="date" class="typecho-label">发布日期</label>
                            <p><input class="typecho-date w-100" type="text" name="date" id="date" value="" /></p>
                        </section>

                        <section class="typecho-post-option category-option">
                            <label class="typecho-label">分类</label>
                                <ul>
                                    {{range .list}}
                                    <li><input type="checkbox" id="category-{{.Mid}}" value="{{.Mid}}" name="category[]" />
                                        <label for="category-{{.Mid}}">{{.Name}}</label>
                                    </li>
                                    {{end}}
                                </ul>
                        </section>

                        <section class="typecho-post-option">
                            <label for="token-input-tags" class="typecho-label">标签</label>
                            <p><input id="tags" name="tags" type="text" value="" class="w-100 text" /></p>
                        </section>

                        
                        <button type="button" id="advance-panel-btn" class="btn btn-xs">高级选项 <i class="i-caret-down"></i></button>
                        <div id="advance-panel">
                                                        <section class="typecho-post-option visibility-option">
                                <label for="visibility" class="typecho-label">公开度</label>
                                <p>
                                <select id="visibility" name="visibility">
                                                                        <option value="publish" selected>公开</option>
                                    <option value="hidden">隐藏</option>
                                    <option value="password">密码保护</option>
                                    <option value="private">私密</option>
                                                                        <option value="waiting">待审核</option>
                                </select>
                                </p>
                                <p id="post-password" class="hidden">
                                    <label for="protect-pwd" class="sr-only">内容密码</label>
                                    <input type="text" name="password" id="protect-pwd" class="text-s" value="" size="16" placeholder="内容密码" />
                                </p>
                            </section>
                            
                            <section class="typecho-post-option allow-option">
                                <label class="typecho-label">权限控制</label>
                                <ul>
                                    <li><input id="allowComment" name="allowComment" type="checkbox" value="1" checked="true" />
                                    <label for="allowComment">允许评论</label></li>
                                    <li><input id="allowPing" name="allowPing" type="checkbox" value="1" checked="true" />
                                    <label for="allowPing">允许被引用</label></li>
                                    <li><input id="allowFeed" name="allowFeed" type="checkbox" value="1" checked="true" />
                                    <label for="allowFeed">允许在聚合中出现</label></li>
                                </ul>
                            </section>
                            
                            <section class="typecho-post-option">
                                <label for="trackback" class="typecho-label">引用通告</label>
                                <p><textarea id="trackback" class="w-100 mono" name="trackback" rows="2"></textarea></p>
                                <p class="description">每一行一个引用地址, 用回车隔开</p>
                            </section>

                                                    </div><!-- end #advance-panel -->

                                            </div><!-- end #tab-advance -->

                    <div id="tab-files" class="tab-content hidden">
                        

                    <div id="upload-panel" class="p">
                        <div class="upload-area" draggable="true">拖放文件到这里<br>或者 <a href="###" class="upload-file">选择文件上传</a></div>
                        <ul id="file-list">
                            </ul>
                    </div>

                    </div><!-- end #tab-files -->
                </div>
            </form>
        </div>
    </div>
</div>

{{template "admin/template/copyright.tpl" .}}
{{template "admin/template/common-js.tpl" .}}
{{template "admin/template/form-js.tpl" .}}
{{template "admin/template/write-js.tpl" .}}
{{template "admin/template/editor-js.tpl" .}}
{{template "admin/template/file-upload-js.tpl" .}}
{{template "admin/template/footer.tpl" .}}