<?php
    function echoComment($comment,$show = true)
    {
//  这个EOF结束处不能有空格
//  参考 https://www.runoob.com/php/php-eof-heredoc.html
        echo <<<EOF
        <div id="comment-${comment["id"]}" class="comment-body">
            <div class="comment-author vcard">
EOF;
        if ($comment["uid"] != ""){
            $href = "//space.bilibili.com/${comment["uid"]}/dynamic";
        } else {
            $href = "javascript:void(0);";
        }
        // 注意要判断一下hang，要不然会重复请求
        if ($comment["hang"]!= ""){
            $hang = "<img class=\"user-decorator\" src=\"${comment["hang"]};\" />";
        } else {
            $hang = "";
        }
        echo <<<EOF
        <a href="$href" target="_blank">
            <img class="user-head c-pointer" style="border-radius: 50%;" src="${comment["avatar"]}" alt=""/>
            $hang
         </a>
EOF;

        echo <<<EOF
                <cite class="fn">${comment["nickname"]}</cite>
                <a href="javascript:void(0);" target="_blank" lvl="${comment["level"]}" class="n-level m-level"></a>
            </div>
            <div class="comment-meta commentmetadata">
                <a id="comment-${comment["id"]}">${comment["date"]}</a>
            </div>
            <div class="comment-content">${comment["content"]}</div>
EOF;
        if ($show){
            echo <<<EOF
            <div class="reply" @click="reply(${comment["id"]})">
                <a class="comment-reply-link">回复</a>
            </div>
EOF;
        }
        echo '</div>';
    }
?>
<div id="comments" class="comments-area">
    <div class="comment-list">
        <?php $comments = get_post_comments();
        foreach ($comments as $comment){
            if ($comment["parent"]==0){
                echo '<div class="comment">';
                echoComment($comment);
                foreach ($comments as $comment2){
                    if ($comment2["parent"]==$comment["id"]){
                        echo '<div class="children"><div class="comment">';
                        echoComment($comment2);
                        foreach ($comments as $comment3) {
                            if ($comment3["parent"] == $comment2["id"]) {
                                echo '<div class="children"><div class="comment">';
                                echoComment($comment3);
                                foreach ($comments as $comment4) {
                                    if ($comment4["parent"] == $comment3["id"]) {
                                        echo '<div class="children"><div class="comment">';
                                        echoComment($comment4);
                                        foreach ($comments as $comment5) {
                                            if ($comment5["parent"] == $comment4["id"]) {
                                                echo '<div class="children"><div class="comment">';
                                                echoComment($comment5,false);
                                                echo '</div></div>';
                                            }
                                        }
                                        echo '</div></div>';
                                    }
                                }
                                echo '</div></div>';
                            }
                        }
                        echo '</div></div>';
                    }
                }
                echo '</div>';
            }
        }
        ?>
    </div>
    <div id="respond" class="comment-respond">
        <h4 id="reply-title" class="comment-reply-title">
            发表评论
            <a id="cancel-comment-reply-link" style="display:none;" @click="cancelComment">放弃治疗</a>
        </h4>
        <div id="comment-info">
            <a href="javascript:void(0);" class="fa">
                <img alt :src="commentData.avatar" width="48" height="48" class="up-face">
                <img
                    v-if="commentData.hang"
                    alt
                    :src="commentData.hang"
                    class="pendant"
                >
            </a>
            <div id="comment-form" class="comment-text">
                <el-form ref="comment" :model="commentData" :rules="rules">
                    <?php if (!user_is_login()): ?>
                    <div>
                        <p class="comment-tip">
                            小提示:输入B站uid可以快速获取昵称和头像哦(*为必填项)
                        </p>
                        <div class="comment-input">
                            <el-form-item class="comment_input">
                                <el-input
                                        id="comment-uid"
                                        v-model="commentData.uid"
                                        placeholder="B站uid"
                                        prefix-icon="el-icon-s-opportunity"
                                ></el-input>
                            </el-form-item>
                            <el-form-item class="comment_input" prop="name">
                                <el-input
                                        id="comment-author"
                                        v-model="commentData.name"
                                        placeholder="昵称 *"
                                        prefix-icon="el-icon-user"
                                ></el-input>
                            </el-form-item>
                            <el-form-item class="comment_input" prop="email">
                                <el-input
                                        id="comment-email"
                                        v-model="commentData.email"
                                        prop="email"
                                        placeholder="邮箱 *"
                                        prefix-icon="el-icon-message"
                                ></el-input>
                            </el-form-item>
                            <el-form-item class="comment_input">
                                <el-input
                                        id="comment-url"
                                        v-model="commentData.site"
                                        placeholder="网址"
                                        prefix-icon="el-icon-attract"
                                ></el-input>
                            </el-form-item>
                        </div>
                    </div>
                    <?php else: ?>
                    <div>
                        <p class="logged-in-as">
                            <a href="javascript:void(0)">已登录为: <?php user_nickname(); ?></a>。
                        </p>
                    </div>
                    <?php endif; ?>
                    <el-form-item class="comment-s" prop="content">
                        <input id="hide-comment" style="display:none" type="text"/>
                        <el-input
                                class="comment-content"
                                id="comment"
                                v-model="commentData.content"
                                type="textarea"
                                placeholder="期待大佬的精彩发言"
                        ></el-input>
                        <div class="OwO">
                        </div>
                        <button type="button" class="comment-submit push-status" @click="comment">
                            发表评论
                        </button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </div>
</div>
