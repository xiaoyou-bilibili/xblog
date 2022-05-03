<article class="kratos-hentry clearfix wow bounceInUp animated"
         style="visibility: visible; animation-name: bounceInUp;">
    <div class="kratos-entry-border-new clearfix">
        <?php if(get_post_is_top()) echo '<img class="stickyimg" src="'.setting_template().'/static/images/sticky.png"/>'; ?>
        <div class="kratos-post-inner-new"
             style="background-image:url('<?php post_image(); ?>');">
            <div class="kratos-entry-content-new ">
                <header class="kratos-entry-header-new">
                    <h2 class="kratos-entry-title-new">
                        <a href="<?php echo setting_web()."/archives/";post_id(); ?>"><?php post_title(); ?></a>
                    </h2>
                </header>
                <p><?php post_summer(); ?></p>
            </div>
        </div>
        <div class="kratos-post-meta-new">
        <span class="pull-left">
            <a href="javascript:;"><i class="fa fa-calendar"></i> <?php post_date(); ?></a>
            <a href="javascript:;"><i class="fas fa-comments"></i> <?php post_comment(); ?>条评论</a>
        </span>
            <span class="visible-lg visible-md visible-sm pull-left">
                <a href="javascript:;"><i class="fa fa-eye"></i> <?php post_view(); ?>次阅读</a>
                <a href="javascript:;"><i class="fas fa-thumbs-up"></i> <?php post_good(); ?>人点赞</a>
            </span>
            <span class="pull-right">
            <a class="read-more" href="<?php echo setting_web()."/archives/";post_id(); ?>" title="阅读全文">
                阅读全文 <i class="fa fa-chevron-circle-right"></i>
            </a>
        </span>
        </div>
    </div>
</article>