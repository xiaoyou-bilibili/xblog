<article class="post post-list-thumb <?php echo get_post_no()%2==1?'post-list-thumb-left':'' ?>" itemscope="" itemtype="https://schema.org/BlogPosting">
    <div class="post-thumb">
        <a href="<?php echo setting_web()."/archives/".get_post_id();?>">
            <img class="lazyload" id="post-th-4482"
                 src=""
                 onerror="imgError(this,3)"
                 data-src="<?php post_image(); ?>"
            />
        </a>
    </div>
    <div class="post-content-wrap">
        <div class="post-content">
            <div class="post-date">
                <i class="iconfont icon-time"></i>发布于 <?php post_date(); ?>
            </div>
            <a href="<?php echo setting_web()."/archives/".get_post_id();?>" class="post-title">
                <h3><?php post_title(); ?></h3>
            </a>
            <div class="post-meta">
                <span>
                    <i class="iconfont icon-attention"></i><?php post_view(); ?>热度
                </span>
                <span class="comments-number">
                    <i class="iconfont icon-mark"></i><?php post_comment(); ?> 条评论
                </span>
            </div>
            <div class="float-content">
                <p>
                    <?php post_content(); ?>
                </p>
                <div class="post-bottom">
                    <a href="<?php echo setting_web()."/archives/".get_post_id();?>" class="button-normal">
                        <i class="iconfont icon-caidan"></i>
                    </a>
                </div>
            </div>
        </div>
    </div>
</article>