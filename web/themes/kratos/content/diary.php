<div class="diary-card">
    <img alt class="user-icon" src="<?php plugins_diary_avatar(); ?>" />
    <div class="arrow"></div>
    <div class="diary-content">
        <div class="diary-head">
            <div class="diary-head-name">
                <?php plugins_diary_nickname(); ?>
            </div><div class="diary-head-time">
                <?php plugins_diary_date(); ?>
            </div>
        </div>
        <?php if (!get_plugins_diary_encrypt()):?>
        <div id="post" class="diary-body">
            <?php plugins_diary_content(); ?>
        </div>
        <?php else: ?>
        <div class="lock-face">
            <div class="lock-text">
                这篇日记已被上锁！
            </div>
        </div>
        <?php endif;?>
        <div class="diary-footer">
            <div class="diary-comment">
                <i class="far fa-comment"></i><?php plugins_diary_comment(); ?>条评论
            </div>
            <div class="diary-like">
                <i class="far fa-thumbs-up"></i><?php plugins_diary_good(); ?>人点赞
            </div>
        </div>
    </div>
</div>

