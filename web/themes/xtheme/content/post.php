<div class="panel">
    <?php
        if (get_post_is_top()){ echo '<img class="stick" alt src="'.setting_template()."/static/image/post/sticky.png\"".'>'; }
    ?>
    <div class="index-post-img">
        <a href="<?php echo setting_web()."/archives/".get_post_id();?>" class="link">
            <div style="background-image:url(<?php post_image(); ?>)"  class="item-thumb"></div>
        </a>
    </div>
    <div class="post-meta wrapper-lg p-b-none">
        <h2 class="m-t-none index-post-title">
            <a href="<?php echo setting_web()."/archives/".get_post_id();?>">
                <?php post_title(); ?>
            </a>
        </h2>
        <p class="summary l-h-2x text-muted">
            <?php post_summer(); ?>
        </p>
        <div class="line line-lg b-b b-light"></div>
        <div class="text-muted post-item-foot-icon">
            <span><i class="fa fa-clock"></i><?php post_date(); ?></span>
            <span><i class="fas fa-comments"></i><?php post_comment(); ?>条评论</span>
            <span><i class="fa fa-eye"></i><?php post_view(); ?>次阅读</span>
            <span class="good"><i class="fa fa-thumbs-up"></i><?php post_good(); ?>人点赞</span>
        </div>
    </div>
</div>