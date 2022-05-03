<?php get_header('post'); ?>
    <link rel="stylesheet" type="text/css" href="<?php echo setting_template()."/static/css/style.css"?>">
    <div class="main-part">
        <?php $show = xy_option("theme_sakura_show_side");if($show):?>
            <div class="col-xl-2 col-lg-3 col-md-10 col-sm-3 col-0 div-box-content">
                <div id="sidebar-left" class="giligili-left">
                    <?php xy_left_side(); ?>
                </div>
            </div>
        <?php endif;?>
        <div class="col-xl-5 col-lg-5 col-md-10 col-sm-2 col-11  site-content">
            <div class="pattern-center single-center">
                <div class="pattern-attachment-img">
                    <img src="" data-src="<?php post_image(); ?>#lazyload-blur" class="lazyload" onerror="imgError(this,3)" style="width: 100%; height: 100%; object-fit: cover; pointer-events: none;">
                </div>
                <header class="pattern-header single-header">
                    <h1 class="entry-title"><?php post_title(); ?></h1>
                    <p class="entry-census"><?php post_date(); ?><span class="bull">·</span><?php post_view(); ?> 次阅读<span class="bull">·</span><?php post_comment(); ?> 条评论</p>
                </header>
            </div>
            <div id="content" class="site-content">
                <main id="main" class="site-main" role="main">
                    <article id="post-4440" class="post-4440 post type-post status-publish format-standard has-post-thumbnail hentry category-hacking tag-ssh tag-vscode tag-windows">
                        <div class="entry-content">
                            <?php if (get_post_encrypt()): ?>
                                <form action="" class="post-password-form" method="get">
                                    <p>这是一篇受密码保护的文章，您需要提供访问密码：</p>
                                    <p><label for="pwbox-7">密码： <input name="password" id="pwbox-7" type="password" size="20"></label> <input type="submit" value="提交"></p>
                                </form>
                            <?php else: post_content(); endif;?>
                        </div>
                        <div class="single-reward">
                            <div class="reward-open">赏
                                <div class="reward-main">
                                    <ul class="reward-row">
                                        <li class="alipay-code">
                                            <img src="<?php post_alipay(); ?>">
                                        </li>
                                        <li class="wechat-code">
                                            <img src="<?php post_wechat(); ?>">
                                        </li>
                                    </ul>
                                </div>
                            </div>
                        </div>
                        <footer class="post-footer">
                            <div class="post-lincenses"><a href="https://creativecommons.org/licenses/by-nc-sa/4.0/deed.zh" target="_blank" rel="nofollow"><i class="fa fa-creative-commons" aria-hidden="true"></i> 知识共享署名-非商业性使用-相同方式共享 4.0 国际许可协议</a></div>
                            <div class="post-tags">
                                <i class="iconfont icon-tags"></i>
                                <?php foreach (get_post_tags() as $tag){
                                    echo "<a href='/?tag=${tag["link"]}' rel='tag'>${tag["name"]}</a>";
                                }?>
                            </div>
                            <div class="post-share">
                                <div class="social-share sharehidden"></div>
                                <i class="iconfont show-share icon-forward"></i>
                            </div>
                        </footer>
                    </article>
                    <div class="toc-container">
                        <div class="toc"></div>
                    </div>
                </main>
            </div>
        </div>
        <?php if($show):?>
            <div class="col-xl-2 col-lg-3 col-md-3 col-sm-3 col-0 div-box-content right-box">
                <div id="sidebar-right" class="giligili-left sidebar">
                    <?php xy_right_side(); ?>
                </div>
            </div>
        <?php endif;?>
    </div>
<?php xy_comments(get_post_id());  get_footer(); ?>