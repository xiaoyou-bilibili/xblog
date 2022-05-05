<?php
    function echoComment($comment,$show = true)
    {
    $postId = get_post_id();
//  这个EOF结束处不能有空格
//  参考 https://www.runoob.com/php/php-eof-heredoc.html
        echo <<<EOF
<div class="contents">
    <div class="comment-arrow">
        <div class="main">
            <div class="profile">
                <a href="javascript: return false;" rel="nofollow">
                    <img src="https://cdn.jsdelivr.net/gh/moezx/cdn@3.0.2/img/svg/loader/trans.ajax-spinner-preloader.svg" onerror="imgError(this,1)" data-src="${comment["avatar"]}" class="lazyload avatar avatar-24 photo" alt="😀" width="24" height="24">
                </a>
            </div>
            <div class="commentinfo">
                <section class="commeta">
                    <div class="left">
                        <h4 class="author">
                          <a href="javascript: return false;" rel="nofollow">
                            ${comment["nickname"]}
                            <span class="showGrade0" title="萌萌哒新人~">
                                <img src="https://cdn.jsdelivr.net/gh/moezx/cdn@3.1.9/img/Sakura/images/level/level_${comment["level"]}.svg" style="height: 1.5em; max-height: 1.5em; display: inline-block;">
                            </span>
                           </a>
                        </h4>
                    </div>
EOF;
        if ($show){
            echo "<a rel=\"nofollow\" class=\"comment-reply-link\" href=\"#comment-${comment["id"]}\" data-commentid=\"${comment["id"]}\" data-postid=\"$postId\" data-belowelement=\"comment-${comment["id"]}\" data-respondelement=\"respond\">Reply</a>";
        }
echo <<<EOF
                    <div class="right">
                        <div class="info">
                            <time datetime="2020-03-01">发布于 ${comment["date"]}</time>
                        </div>
                    </div>
                </section>
            </div>
            <div class="body">
                <p>${comment["content"]}</p>
            </div>
        </div>
        <div class="arrow-left"></div>
    </div>
</div>
EOF;
    }
?>
<section id="comments" class="comments">
    <div class="comments-main">
        <h3 id="comments-list-title">Comments | <span class="noticom"><?php post_comment();?> 条评论 </span></h3>
        <div id="loading-comments"><span></span></div>
        <ul class="commentwrap">
            <?php $comments = get_post_comments();
            foreach ($comments as $comment){
                if ($comment["parent"]==0){
                    echo '<li class="comment even thread-even depth-1" id="comment-"'.$comment["id"].'>';
                    echoComment($comment);
                    echo '<hr><ul class="children">';
                    foreach ($comments as $comment2){
                        if ($comment2["parent"]==$comment["id"]){
                            echo '<li class="comment odd alt depth-2" id="comment-"'.$comment2["id"].'>';
                            echoComment($comment2);
                            echo '<hr><ul class="children">';
                            foreach ($comments as $comment3){
                                if ($comment3["parent"]==$comment2["id"]){
                                    echo '<li class="comment odd alt depth-3" id="comment-"'.$comment3["id"].'>';
                                    echoComment($comment3);
                                    echo '<hr><ul class="children">';
                                    foreach ($comments as $comment4){
                                        if ($comment4["parent"]==$comment3["id"]){
                                            echo '<li class="comment odd alt depth-4" id="comment-"'.$comment4["id"].'>';
                                            echoComment($comment4);
                                            echo '<hr><ul class="children">';
                                            foreach ($comments as $comment5){
                                                if ($comment5["parent"]==$comment4["id"]){
                                                    echo '<li class="comment odd alt depth-5" id="comment-"'.$comment5["id"].'>';
                                                    echoComment($comment5,false);
                                                }
                                            }
                                            echo '</ul></li>';
                                        }
                                    }
                                    echo '</ul></li>';
                                }
                            }
                            echo '</ul></li>';
                        }
                    }
                    echo '</ul></li>';
                }
            }
            ?>
        </ul>
        <div id="respond" class="comment-respond">
            <h3 id="reply-title" class="comment-reply-title">
                <small>
                    <a rel="nofollow" id="cancel-comment-reply-link" href="" style="display: none;">Cancel Reply</a>
                </small>
            </h3>
            <form action="/" method="post" id="commentform" class="comment-form">
                <div class="comment-textarea">
                    <textarea placeholder="你是我一生只会遇见一次的惊喜 ..." name="comment" class="commentbody" id="comment" rows="5" tabindex="4"></textarea>
                    <label class="input-label active">你是我一生只会遇见一次的惊喜 ...</label>
                </div>
                <div id="upload-img-show"></div>
                <!--插入表情面版-->
                <p id="emotion-toggle" class="no-select">
                    <span class="emotion-toggle-off">戳我试试 OωO</span>
                    <span class="emotion-toggle-on">嘿嘿嘿 ヾ(≧∇≦*)ゝ</span>
                </p>
                <div class="emotion-box no-select">
                    <?php $smiles = tools_get_smile(); ?>
                    <table class="motion-switcher-table">
                        <tbody>
                        <tr>
                            <th onclick="motionSwitch('.bili')" class="bili-bar on-hover">bilibili~</th>
                            <th onclick="motionSwitch('.tv')" class="tv-bar">小电视</th>
                            <th onclick="motionSwitch('.zhihu')" class="zhihu-bar">知乎</th>
                            <th onclick="motionSwitch('.tieba')" class="tieba-bar">贴吧泡泡</th>
                            <th onclick="motionSwitch('.menhera')" class="menhera-bar">颜文字</th>
                        </tr>
                        </tbody>
                    </table>
                    <div class="bili-container motion-container">
                        <?php foreach ($smiles["bilibili"]["container"] as $item){
                                echo "<span class=\"emotion-secter emotion-item emotion-select-parent\" onclick=\"grin(${item["desc"]})\">${item["icon"]}</span>";
                            } ?>
                    </div>
                    <div class="tv-container motion-container" style="display:none;">
                        <?php foreach ($smiles["小电视"]["container"] as $item){
                            echo "<span class=\"emotion-secter emotion-item emotion-select-parent\" onclick=\"grin(${item["desc"]})\">${item["icon"]}</span>";
                        } ?>
                    </div>
                    <div class="zhihu-container motion-container" style="display:none;">
                        <?php foreach ($smiles["知乎表情"]["container"] as $item){
                            echo "<span class=\"emotion-secter emotion-item emotion-select-parent\" onclick=\"grin(${item["desc"]})\">${item["icon"]}</span>";
                        } ?>
                    </div>
                    <div class="tieba-container motion-container" style="display:none;">
                        <?php foreach ($smiles["贴吧泡泡"]["container"] as $item){
                            echo "<span class=\"emotion-secter emotion-item emotion-select-parent\" onclick=\"grin(${item["desc"]})\">${item["icon"]}</span>";
                        } ?>
                    </div>
                    <div class="menhera-container motion-container" style="display:none;">
                        <?php foreach ($smiles["颜文字"]["container"] as $item){
                            echo "<a class=\"emoji-item\">${item["desc"]}</a>";
                        } ?>
                    </div>
                </div>
                <!--表情面版完-->
                <div class="cmt-info-container">
                    <?php if (!user_is_login()): ?>
                    <div class="comment-user-avatar">
                        <img alt="" src="">
                    </div>
                    <div class="popup cmt-popup cmt-author" onclick="cmt_showPopup(this)">
                        <span class="popuptext" id="thePopup" style="margin-left: -115px;width: 230px;">输入UID将自动拉取昵称和头像</span>
                        <input type="text" placeholder="昵称或B站UID (必须 Name* )" name="author" id="author" value="" size="22" autocomplete="off" tabindex="1" aria-required="true">
                    </div>
                    <div class="popup cmt-popup" onclick="cmt_showPopup(this)">
                        <span class="popuptext" id="thePopup" style="margin-left: -65px;width: 130px;">你将收到回复通知</span>
                        <input type="text" placeholder="邮箱 (必须 Email* )" name="email" id="email" value="" size="22" tabindex="1" autocomplete="off" aria-required="true">
                    </div>
                    <div class="popup cmt-popup" onclick="cmt_showPopup(this)">
                        <span class="popuptext" id="thePopup" style="margin-left: -55px;width: 110px;">禁止小广告😀</span>
                        <input type="text" placeholder="网站 (选填 Site)" name="url" id="url" value="" size="22" autocomplete="off" tabindex="1">
                    </div>
                    <?php else: ?>
                    <div class="popup">
                        已登录为：<?php user_nickname(); ?>
                    </div>
                    <?php endif; ?>
                </div>
                <p class="form-submit">
                    <input name="submit" type="submit" id="submit" class="submit" value="BiuBiuBiu~">
                    <input type="hidden" name="comment_post_ID" value="<?php post_id();?>" id="comment_post_ID">
                    <input type="hidden" name="comment_parent" id="comment_parent" value="0">
                </p>
            </form>
        </div>
    </div>
</section>
