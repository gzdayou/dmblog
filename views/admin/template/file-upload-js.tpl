<script src="/adresource/js/moxie.js?v=17.10.30"></script>
<script src="/adresource/js/plupload.js?v=17.10.30"></script>
<script>
$(document).ready(function() {
    function updateAttacmentNumber () {
        var btn = $('#tab-files-btn'),
            balloon = $('.balloon', btn),
            count = $('#file-list li .insert').length;

        if (count > 0) {
            if (!balloon.length) {
                btn.html($.trim(btn.html()) + ' ');
                balloon = $('<span class="balloon"></span>').appendTo(btn);
            }

            balloon.html(count);
        } else if (0 == count && balloon.length > 0) {
            balloon.remove();
        }
    }

    $('.upload-area').bind({
        dragenter   :   function () {
            $(this).parent().addClass('drag');
        },

        dragover    :   function (e) {
            $(this).parent().addClass('drag');
        },

        drop        :   function () {
            $(this).parent().removeClass('drag');
        },
        
        dragend     :   function () {
            $(this).parent().removeClass('drag');
        },

        dragleave   :   function () {
            $(this).parent().removeClass('drag');
        }
    });

    updateAttacmentNumber();

    function fileUploadStart (file) {
        $('<li id="' + file.id + '" class="loading">'
            + file.name + '</li>').appendTo('#file-list');
    }

    function fileUploadError (error) {
        var file = error.file, code = error.code, word; 
        
        switch (code) {
            case plupload.FILE_SIZE_ERROR:
                word = '文件大小超过限制';
                break;
            case plupload.FILE_EXTENSION_ERROR:
                word = '文件扩展名不被支持';
                break;
            case plupload.FILE_DUPLICATE_ERROR:
                word = '文件已经上传过';
                break;
            case plupload.HTTP_ERROR:
            default:
                word = '上传出现错误';
                break;
        }

        var fileError = '%s 上传失败'.replace('%s', file.name),
            li, exist = $('#' + file.id);

        if (exist.length > 0) {
            li = exist.removeClass('loading').html(fileError);
        } else {
            li = $('<li>' + fileError + '<br />' + word + '</li>').appendTo('#file-list');
        }

        li.effect('highlight', {color : '#FBC2C4'}, 2000, function () {
            $(this).remove();
        });

        // fix issue #341
        this.removeFile(file);
    }

    var completeFile = null;
    function fileUploadComplete (id, url, data) {
        var li = $('#' + id).removeClass('loading').data('cid', data.cid)
            .data('url', data.url)
            .data('image', data.isImage)
            .html('<input type="hidden" name="attachment[]" value="' + data.cid + '" />'
                + '<a class="insert" target="_blank" href="###" title="点击插入文件">' + data.title + '</a><div class="info">' + data.bytes
                + ' <a class="file" target="_blank" href="http://localhost/typecho/admin/media.php?cid=' 
                + data.cid + '" title="编辑"><i class="i-edit"></i></a>'
                + ' <a class="delete" href="###" title="删除"><i class="i-delete"></i></a></div>')
            .effect('highlight', 1000);
            
        attachInsertEvent(li);
        attachDeleteEvent(li);
        updateAttacmentNumber();

        if (!completeFile) {
            completeFile = data;
        }
    }

    $('#tab-files').bind('init', function () {
        var uploader = new plupload.Uploader({
            browse_button   :   $('.upload-file').get(0),
            url             :   'http://localhost/typecho/index.php/action/upload?_=d585f8eb5769324c17e304cbd6052dd2',
            runtimes        :   'html5,flash,html4',
            flash_swf_url   :   'http://localhost/typecho/admin/js/Moxie.swf',
            drop_element    :   $('.upload-area').get(0),
            filters         :   {
                max_file_size       :   '20mb',
                mime_types          :   [{'title' : '允许上传的文件', 'extensions' : 'gif,jpg,jpeg,png,tiff,bmp'}],
                prevent_duplicates  :   true
            },

            init            :   {
                FilesAdded      :   function (up, files) {
                    for (var i = 0; i < files.length; i ++) {
                        fileUploadStart(files[i]);
                    }

                    completeFile = null;
                    uploader.start();
                },

                UploadComplete  :   function () {
                    if (completeFile) {
                        Typecho.uploadComplete(completeFile);
                    }
                },

                FileUploaded    :   function (up, file, result) {
                    if (200 == result.status) {
                        var data = $.parseJSON(result.response);

                        if (data) {
                            fileUploadComplete(file.id, data[0], data[1]);
                            uploader.removeFile(file);
                            return;
                        }
                    }

                    fileUploadError.call(uploader, {
                        code : plupload.HTTP_ERROR,
                        file : file
                    });
                },

                Error           :   function (up, error) {
                    fileUploadError.call(uploader, error);
                }
            }
        });

        uploader.init();
    });

    function attachInsertEvent (el) {
        $('.insert', el).click(function () {
            var t = $(this), p = t.parents('li');
            Typecho.insertFileToEditor(t.text(), p.data('url'), p.data('image'));
            return false;
        });
    }

    function attachDeleteEvent (el) {
        var file = $('a.insert', el).text();
        $('.delete', el).click(function () {
            if (confirm('确认要删除文件 %s 吗?'.replace('%s', file))) {
                var cid = $(this).parents('li').data('cid');
                $.post('http://localhost/typecho/index.php/action/contents-attachment-edit?_=d585f8eb5769324c17e304cbd6052dd2',
                    {'do' : 'delete', 'cid' : cid},
                    function () {
                        $(el).fadeOut(function () {
                            $(this).remove();
                            updateAttacmentNumber();
                        });
                    });
            }

            return false;
        });
    }

    $('#file-list li').each(function () {
        attachInsertEvent(this);
        attachDeleteEvent(this);
    });
});
</script>