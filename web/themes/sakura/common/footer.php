</div>
<footer id="colophon"  class="site-footer" role="contentinfo">
    <div class="site-info">
        <div class="footertext">
            <?php // 这个是加载效果 ?>
            <div class="img-preload">
                <img src="<?php echo setting_template().'/static/image/wordpress-rotating-ball-o.svg'?>"/>
                <img src="<?php echo setting_template().'/static/image/disqus-preloader.svg'?>"/>
            </div>
            <p class="foo-logo" style="background-image: url('<?php echo setting_template().'/static/image/sakura.svg'?>');"></p>
            <p style="font-family: 'Ubuntu', sans-serif">
                <span style="color: #666666">
                    原主题
                    <a target="_blank" href="https://2heng.xin/theme-sakura/">Sakura</a>
                <i class="fa fa-heart animated" style="color: #e74c3c"></i>
                    Power By
                    <a rel="me" href="https://xblog.xiaoyou66.com/" target="_blank">XBlog</a>
              </span>
            </p>
            <p>
                &copy; 2021 <?php setting_index_site_name(); ?>
                <a href="http://www.beian.miit.gov.cn" target="_blank"> <?php setting_index_bei_an(); ?></a>
                <a href="/sitemap" target="_blank">SITEMAP</a>
            </p>
        </div>
        <div class="footer-device"></div>
        <?php
        // 这个是底部的赞助信息，如果想加自己注释掉这个代码
//        <p id="footer-sponsor">
//            <a href="#" target="_blank" alt="CDN Sponsor" rel="nofollow">
//                <img src="https://2heng.xin/wp-content/themes/Sakura/images/Tencent_Cloud_logo.svg" alt="CDN Sponsor" style="height: 1.5em" />
//            </a>
//        </p>
        ?>
    </div>
</footer>
<div class="openNav no-select">
    <div class="iconflat no-select">
        <div class="icon"></div>
    </div>
    <div class="site-branding">
        <h1 class="site-title">
            <a href="<?php setting_index_site_url(); ?>"><?php setting_index_site_name(); ?></a>
        </h1>
    </div>
</div>
</section>
<?php // 下面这一部分是手机版视图 ?>
<div id="mo-nav">
    <div class="m-avatar">
        <img src="<?php echo xy_option("theme_sakura_avatar"); ?>" />
    </div>
    <p style="text-align: center;color: #333;font-weight: 900;font-family: 'Ubuntu', sans-serif;letter-spacing: 1.5px;">
        <?php echo xy_option("theme_sakura_sign"); ?>
    </p>
    <div class="m-search">
        <form class="m-search-form" method="get" action="/" role="search">
            <input class="m-search-input" type="search" name="q" placeholder="搜索..." required />
        </form>
    </div>
    <ul id="menu-new-1" class="menu">
        <?php foreach (setting_index_head_nav() as $head): ?>
            <li>
                <?php echo '<a href="'.$head["link"].'"><span class="faa-parent animated-hover">'.$head["title"].'</span></a>' ?>
                <?php if (count($head["children"])>0): ?>
                    <ul class="sub-menu">
                        <?php foreach ($head["children"] as $head1): ?>
                            <?php echo '<li><a href="'.$head1["link"].'">'.$head1["title"].'</a></li>' ?>
                        <?php endforeach; ?>
                    </ul>
                <?php endif; ?>
            </li>
        <?php endforeach; ?>
    </ul>
    <p style="text-align: center; font-size: 13px; color: #b9b9b9">
        &copy; 2021 <?php setting_index_site_name(); ?>
    </p>
</div>
<a href="#" class="cd-top faa-float animated"></a>
<button onclick="topFunction()" class="mobile-cd-top" id="moblieGoTop" title="Go to top">
    <i class="fa fa-chevron-up" aria-hidden="true"></i>
</button>
<form class="js-search search-form search-form--modal" method="get" action="/" role="search">
    <div class="search-form__inner">
        <div>
            <p class="micro mb-">想要找点什么呢？</p>
            <i class="iconfont icon-search"></i>
            <input class="text-input" type="search" name="q" placeholder="Search" required/>
        </div>
    </div>
    <div class="search_close"></div>
</form>
<!--视频加载库-->
<script type="text/javascript" src="https://cdn.jsdelivr.net/gh/moeplayer/hermit-x@2.9.7/assets/js/hermit-load.min.js?ver=2.9.7"></script>

<script type="text/javascript" src="<?php echo setting_template()."/static/js/lib.min.js"?>"></script>
<script type="text/javascript">
    var Poi = {
        pjax: "1",
        movies: {
            url: "<?php echo xy_option("theme_sakura_index_movie"); ?>",
            name: '主页视频',
            live: "close",
        },
        windowheight: "auto",
        codelamp: "close",
        ajaxurl: "",
        order: "desc",
        formpostion: "bottom",
        bing: "<?php setting_index_background();?>"
    };
</script>
<?php if(xy_option("theme_sakura_show_live2d")): ?>
<div class="prpr">
    <div class="mashiro-tips"></div>
    <canvas id="live2d" width="280" height="250" class="live2d"></canvas>
</div>
<div class="live2d-tool hide-live2d no-select" onclick="hide_live2d()">
    <div class="keys">Hide</div>
</div>
<div class="live2d-tool switch-live2d no-select" onclick="switch_pio()">
    <div class="keys">Switch</div>
</div>
<div class="live2d-tool save-live2d no-select" onclick="save_pio()">
    <div class="keys">Save</div>
</div>
<?php
endif;
// 主题切换 ?>
<div class="changeSkin-gear no-select">
    <div class="keys">
        <span id="open-skinMenu">
          切换主题 | SCHEME TOOL &nbsp;<i class="iconfont icon-gear inline-block rotating"></i>
        </span>
    </div>
</div>
<div class="skin-menu no-select">
    <div class="theme-controls row-container">
        <ul class="menu-list">
            <li id="white-bg">
                <i class="fa fa-television" aria-hidden="true"></i>
            </li>
            <li id="sakura-bg"><i class="iconfont icon-sakura"></i></li>
            <li id="gribs-bg"><i class="fa fa-slack" aria-hidden="true"></i></li>
            <li id="KAdots-bg"><i class="iconfont icon-dots"></i></li>
            <li id="totem-bg">
                <i class="fa fa-optin-monster" aria-hidden="true"></i>
            </li>
            <li id="pixiv-bg"><i class="iconfont icon-pixiv"></i></li>
            <li id="bing-bg"><i class="iconfont icon-bing"></i></li>
            <li id="dark-bg"><i class="fa fa-moon-o" aria-hidden="true"></i></li>
        </ul>
    </div>
    <div class="font-family-controls row-container">
        <button type="button" class="control-btn-serif selected" data-mode="serif" onclick="mashiro_global.font_control.change_font()">
            Serif
        </button>
        <button type="button" class="control-btn-sans-serif" data-mode="sans-serif" onclick="mashiro_global.font_control.change_font()">
            Sans Serif
        </button>
    </div>
</div>
<canvas id="night-mode-cover"></canvas>
<script type="text/javascript" src="<?php echo setting_template()."/static/js/theme.js"?>"></script>
</body>
</html>
