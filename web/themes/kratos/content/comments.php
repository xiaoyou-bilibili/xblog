<?php
    function echoChild($comment,$show = true){
        echo <<<EOF
         <div id="div-comment-${comment['id']}" class="comment-body">
            <div class="comment-author vcard">
                <img alt="" src="${comment['avatar']}" class="avatar avatar-50 photo" height="50" width="50">
                <div style="position: absolute;">
                    <a href="//www.bilibili.com/blackboard/help.html#会员等级相关" target="_blank" lvl="0" class="n-level m-level"></a>
                </div>
                <cite class="fn">
                    <a href="http://127.0.0.1" target="_blank" rel="nofollow">${comment['nickname']}</a>
                </cite>
            </div>
            <div class="comment-meta commentmetadata">${comment['date']}</div>
            <p>${comment['content']}</p>   
EOF;
        if($show){
            echo <<<EOF
            <div class="reply">
                <a rel="nofollow" class="comment-reply-link" href="#" data-commentid="${comment['id']}" data-belowelement="div-comment-${comment['id']}" data-respondelement="respond">回复</a>
            </div>
EOF;
        }

echo <<<EOF
        </div> 
EOF;
    }
?>


<div id="comments" class="comments-area">
    <ol class="comment-list">
        <?php $comments = get_post_comments() ?>
        <?php foreach ($comments as $comment):  ?>
            <?php if ($comment["parent"]==0): ?>
            <li class="comment even thread-even depth-1" id="comment-1">
                <div id="div-comment-<?php echo $comment["id"]?>" class="comment-body">
                    <div class="comment-author vcard">
                        <img alt="" src="<?php echo $comment["avatar"] ?>" class="avatar avatar-50 photo" height="50" width="50">
                        <?php if($comment["level"]!=0): ?>
                        <div style="position: absolute;">
                            <a href="//www.bilibili.com/blackboard/help.html#会员等级相关" target="_blank" lvl="<?php echo $comment["level"] ?>" class="n-level m-level"></a>
                        </div>
                        <?php endif; ?>
                        <cite class="fn"><a href="https://wordpress.org/" target="_blank" rel="nofollow"><?php echo $comment["nickname"] ?></a></cite>
                    </div>
                    <div class="comment-meta commentmetadata">
                        <a href="javascript:"><?php echo $comment["date"] ?></a>
                    </div>
                    <p><?php echo $comment["content"] ?></p>
                    <div class="reply">
                        <a rel="nofollow" class="comment-reply-link" href="#" data-commentid="<?php echo $comment["id"]?>" data-postid="1" data-belowelement="div-comment-<?php echo $comment["id"]?>" data-respondelement="respond">回复</a>
                    </div>
                </div>
                <?php foreach ($comments as $comment1): ?>
                <ol class="children">
                    <?php if ($comment["id"]==$comment1["parent"]): ?>
                    <li class="comment byuser comment-author-xiaoyou even thread-even depth-1">
                        <?php echoChild($comment1); ?>
                        <?php foreach ($comments as $comment2): ?>
                            <ol class="children">
                                <?php if ($comment1["id"]==$comment2["parent"]): ?>
                                    <li class="comment byuser comment-author-xiaoyou even thread-even depth-1">
                                        <?php echoChild($comment2); ?>
                                        <?php foreach ($comments as $comment3): ?>
                                            <ol class="children">
                                                <?php if ($comment2["id"]==$comment3["parent"]): ?>
                                                    <li class="comment byuser comment-author-xiaoyou even thread-even depth-1">
                                                        <?php echoChild($comment3); ?>
                                                        <?php foreach ($comments as $comment4): ?>
                                                            <ol class="children">
                                                                <?php if ($comment3["id"]==$comment4["parent"]): ?>
                                                                    <li class="comment byuser comment-author-xiaoyou even thread-even depth-1">
                                                                        <?php echoChild($comment4,false); ?>
                                                                    </li>
                                                                <?php endif; ?>
                                                            </ol>
                                                        <?php endforeach; ?>
                                                    </li>
                                                <?php endif; ?>
                                            </ol>
                                        <?php endforeach; ?>
                                    </li>
                                <?php endif; ?>
                            </ol>
                        <?php endforeach; ?>
                    </li>
                    <?php endif; ?>
                </ol>
                <?php endforeach; ?>
            </li>
            <?php endif;?>
        <?php endforeach; ?>
    </ol>
    <div id="respond" class="comment-respond">
        <h4 id="reply-title" class="comment-reply-title">发表评论
            <small><a rel="nofollow" id="cancel-comment-reply-link" href="/index.php/2021/04/29/hello-world/#respond" style="display: none;">取消回复</a></small>
        </h4>
        <div id="commentform" class="comment-form">
            <?php if (!user_is_login()): ?>
            <p class="comment-notes">昵称和uid可以选填一个，填邮箱必填(留言回复后将会发邮件给你)<br>tips:输入uid可以快速获得你的昵称和头像</p>
            <?php else: ?>
            <p class="logged-in-as">
                <a href="/admin" aria-label="已登录为xiaoyou。编辑您的个人资料。">已登录为<?php user_nickname()?></a>。
            </p>
            <?php endif; ?>
            <div class="comment form-group has-feedback">
                <div class="input-group">
                    <textarea v-model="commentData.content" class="form-control" id="comment" placeholder="期待大佬的精彩发言~φ(>ω<*) " name="comment" rows="5" aria-required="true" required=""></textarea>
                </div>
                <div class="OwO">
                </div>
            </div>
            <?php if (!user_is_login()): ?>
            <div class="wow bounceInLeft animated" style="visibility: visible; animation-name: bounceInLeft;">
                <div class="comment-form-bilibili form-group has-feedback">
                    <div class="input-group">
                        <div class="input-group-addon">
                            <i class="fa fa-id-card" aria-hidden="true"></i>
                        </div>
                        <input v-model="commentData.uid" class="form-control" placeholder="B站uid*" id="uid" name="uid" type="text" value="" size="30">
                    </div>
                </div>
            </div>
            <div class="wow bounceInRight animated" style="visibility: visible; animation-name: bounceInRight;">
                <div class="comment-form-author form-group has-feedback">
                    <div class="input-group">
                        <div class="input-group-addon"><i class="fa fa-user"></i></div>
                        <input  v-model="commentData.name" class="form-control" placeholder="昵称*" id="author" name="author" type="text" value="" size="30">
                    </div>
                </div>
            </div>
            <div class="wow bounceInLeft animated" style="visibility: visible; animation-name: bounceInLeft;">
                <div class="comment-form-email form-group has-feedback wow jackInTheBox animated" style="visibility: visible; animation-name: jackInTheBox;">
                    <div class="input-group">
                        <div class="input-group-addon"><i class="far fa-envelope"></i></div>
                        <input v-model="commentData.email" class="form-control" placeholder="邮箱" id="email" name="email" type="text" value="" size="30">
                    </div>
                </div>
            </div>
            <div class="wow bounceInRight animated" style="visibility: visible; animation-name: bounceInRight;">
                <div class="comment-form-url form-group has-feedback wow jackInTheBox animated" style="visibility: visible; animation-name: jackInTheBox;">
                    <div class="input-group">
                        <div class="input-group-addon"><i class="fa fa-link"></i></div>
                        <input v-model="commentData.site" class="form-control" placeholder="网站" id="url" name="url" type="text" value="" size="30">
                    </div>
                </div>
            </div>
            <?php endif;?>
            <p class="form-submit">
                <input name="submit" id="submit" class="btn btn-primary" value="发表评论" @click="comment">
                <input type="hidden" name="comment_post_ID" value="<?php post_id(); ?>" id="comment_post_ID">
                <input type="hidden" name="comment_parent" id="comment_parent" value="0">
            </p>
            <input type="hidden" id="_wp_unfiltered_html_comment_disabled" name="_wp_unfiltered_html_comment" value="e2c9787f22">
        </div>
    </div>
</div>
