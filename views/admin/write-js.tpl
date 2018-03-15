<script src="/adresource/js/timepicker.js?v=17.10.30"></script>
<script src="/adresource/js/tokeninput.js?v=17.10.30"></script>
<script>
$(document).ready(function() {
    // 日期时间控件
    $('#date').mask('9999-99-99 99:99').datetimepicker({
        currentText     :   '现在',
        prevText        :   '上一月',
        nextText        :   '下一月',
        monthNames      :   ['一月', '二月', '三月', '四月',
            '五月', '六月', '七月', '八月',
            '九月', '十月', '十一月', '十二月'],
        dayNames        :   ['星期日', '星期一', '星期二',
            '星期三', '星期四', '星期五', '星期六'],
        dayNamesShort   :   ['周日', '周一', '周二', '周三',
            '周四', '周五', '周六'],
        dayNamesMin     :   ['日', '一', '二', '三',
            '四', '五', '六'],
        closeText       :   '完成',
        timeOnlyTitle   :   '选择时间',
        timeText        :   '时间',
        hourText        :   '时',
        amNames         :   ['上午', 'A'],
        pmNames         :   ['下午', 'P'],
        minuteText      :   '分',
        secondText      :   '秒',

        dateFormat      :   'yy-mm-dd',
        timezone        :   28800 / 60,
        hour            :   (new Date()).getHours(),
        minute          :   (new Date()).getMinutes()
    });

    // 聚焦
    $('#title').select();

    // text 自动拉伸
    Typecho.editorResize('text', 'http://localhost/typecho/index.php/action/ajax?do=editorResize&_=d585f8eb5769324c17e304cbd6052dd2');

    // tag autocomplete 提示
    var tags = $('#tags'), tagsPre = [];
    
    if (tags.length > 0) {
        var items = tags.val().split(','), result = [];
        for (var i = 0; i < items.length; i ++) {
            var tag = items[i];

            if (!tag) {
                continue;
            }

            tagsPre.push({
                id      :   tag,
                tags    :   tag
            });
        }

        tags.tokenInput([], {
            propertyToSearch:   'tags',
            tokenValue      :   'tags',
            searchDelay     :   0,
            preventDuplicates   :   true,
            animateDropdown :   false,
            hintText        :   '请输入标签名',
            noResultsText   :   '此标签不存在, 按回车创建',
            prePopulate     :   tagsPre,

            onResult        :   function (result, query, val) {
                if (!query) {
                    return result;
                }

                if (!result) {
                    result = [];
                }

                if (!result[0] || result[0]['id'] != query) {
                    result.unshift({
                        id      :   val,
                        tags    :   val
                    });
                }

                return result.slice(0, 5);
            }
        });

        // tag autocomplete 提示宽度设置
        $('#token-input-tags').focus(function() {
            var t = $('.token-input-dropdown'),
                offset = t.outerWidth() - t.width();
            t.width($('.token-input-list').outerWidth() - offset);
        });
    }

    // 缩略名自适应宽度
    var slug = $('#slug');

    if (slug.length > 0) {
        var wrap = $('<div />').css({
            'position'  :   'relative',
            'display'   :   'inline-block'
        }),
        justifySlug = $('<pre />').css({
            'display'   :   'block',
            'visibility':   'hidden',
            'height'    :   slug.height(),
            'padding'   :   '0 2px',
            'margin'    :   0
        }).insertAfter(slug.wrap(wrap).css({
            'left'      :   0,
            'top'       :   0,
            'minWidth'  :   '5px',
            'position'  :   'absolute',
            'width'     :   '100%'
        })), originalWidth = slug.width();

        function justifySlugWidth() {
            var val = slug.val();
            justifySlug.text(val.length > 0 ? val : '     ');
        }

        slug.bind('input propertychange', justifySlugWidth);
        justifySlugWidth();
    }

    // 原始的插入图片和文件
    Typecho.insertFileToEditor = function (file, url, isImage) {
        var textarea = $('#text'), sel = textarea.getSelection(),
            html = isImage ? '<img src="' + url + '" alt="' + file + '" />'
                : '<a href="' + url + '">' + file + '</a>',
            offset = (sel ? sel.start : 0) + html.length;

        textarea.replaceSelection(html);
        textarea.setSelection(offset, offset);
    };

    var submitted = false, form = $('form[name=write_post],form[name=write_page]').submit(function () {
        submitted = true;
    }), savedData = null;

    // 计算夏令时偏移
    var dstOffset = (function () {
        var d = new Date(),
            jan = new Date(d.getFullYear(), 0, 1),
            jul = new Date(d.getFullYear(), 6, 1),
            stdOffset = Math.max(jan.getTimezoneOffset(), jul.getTimezoneOffset());

        return stdOffset - d.getTimezoneOffset();
    })();
    
    if (dstOffset > 0) {
        $('<input name="dst" type="hidden" />').appendTo(form).val(dstOffset);
    }

    // 时区
    $('<input name="timezone" type="hidden" />').appendTo(form).val(- (new Date).getTimezoneOffset() * 60);

    // 自动保存

    // 自动检测离开页
    var lastData = form.serialize();

    $(window).bind('beforeunload', function () {
        if (!!savedData) {
            lastData = savedData;
        }

        if (form.serialize() != lastData && !submitted) {
            return '内容已经改变尚未保存, 您确认要离开此页面吗?';
        }
    });

    // 控制选项和附件的切换
    var fileUploadInit = false;
    $('#edit-secondary .typecho-option-tabs li').click(function() {
        $('#edit-secondary .typecho-option-tabs li').removeClass('active');
        $(this).addClass('active');
        $(this).parents('#edit-secondary').find('.tab-content').addClass('hidden');
        
        var selected_tab = $(this).find('a').attr('href'),
            selected_el = $(selected_tab).removeClass('hidden');

        if (!fileUploadInit) {
            selected_el.trigger('init');
            fileUploadInit = true;
        }

        return false;
    });

    // 高级选项控制
    $('#advance-panel-btn').click(function() {
        $('#advance-panel').toggle();
        return false;
    });

    // 自动隐藏密码框
    $('#visibility').change(function () {
        var val = $(this).val(), password = $('#post-password');
        console.log(val);

        if ('password' == val) {
            password.removeClass('hidden');
        } else {
            password.addClass('hidden');
        }
    });
    
    // 草稿删除确认
    $('.edit-draft-notice a').click(function () {
        if (confirm('您确认要删除这份草稿吗?')) {
            window.location.href = $(this).attr('href');
        }

        return false;
    });
});
</script>