<!DOCTYPE html>
<html lang="en-US">
<head>
    <meta charset="UTF-8" />
    <meta content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" name="viewport"/>
    <meta name="description" content="<?php setting_head_description(); ?>" />
    <meta name="keywords" content="Code, Economics, ACG" />
    <link rel="shortcut icon" href="<?php setting_head_icon(); ?>" type="image/x-icon" />
    <link rel="apple-touch-icon" href="<?php setting_head_icon(); ?>" />
    <link rel="icon" type="image/png" href="<?php setting_head_icon(); ?>" />
    <?php $theme = xy_option("theme_sakura_theme_color") ?>
    <meta name="msapplication-TileColor" content="#ffffff" />
    <meta name="msapplication-TileImage" content="/ms-icon-144x144.png" />
    <meta name="theme-color" content="<?php echo $theme?>" />
    <!-- https://developer.mozilla.org/zh-CN/docs/Web/Performance/dns-prefetch -->
    <link rel="dns-prefetch" href="//cdn.jsdelivr.net" />
    <?php lib_css_bootstrap(); ?>
    <link media="all" rel="stylesheet" href="<?php echo setting_template()."/static/css/theme.css"?>">
    <title><?php setting_head_title(); ?></title>
    <meta name="description" content="<?php setting_head_description(); ?>" />
    <meta property="og:locale" content="en_US" />
    <meta property="og:type" content="website" />
    <meta property="og:title" content="<?php setting_head_title(); ?>" />
    <meta property="og:description" content="<?php setting_head_description(); ?>" />
    <meta property="og:url" content="<?php setting_head_url(); ?>" />
    <meta property="og:site_name" content="<?php setting_index_site_name(); ?>" />
    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:site" content="<?php setting_head_title(); ?>" />
    <link rel="dns-prefetch" href="//cdn.jsdelivr.net" />
    <link rel="dns-prefetch" href="//s.w.org" />
    <!-- 允许访问站外资源 -->
    <meta name="referrer" content="no-referrer">
    <style type="text/css">
        /*导航栏样式控制*/
        .site-top .lower nav{display:block!important}
        #pagination a:hover,.author-profile i,.float-content i:hover,.post-content a:hover,.post-like a,.post-more i:hover,.post-share .show-share,.sub-text,.we-info a,span.sitename{color:<?php echo $theme?>;text-decoration: none;}
        .ar-time i,.comment .comment-reply-link,.download,.feature i,.links ul li:before,.navigator i:hover,.object,.siren-checkbox-radio:checked+.siren-checkbox-radioInput:after,span.ar-circle{background:<?php echo $theme?>}
        ::-webkit-scrollbar-thumb{background:}
        #pagination a:hover,.comment-respond input[type=submit]:hover,.download,.link-title,.links ul li:hover,.navigator i:hover{border-color:<?php echo $theme?>}
        #archives-temp h3,#comments-navi a.next,#comments-navi a.prev,.comment h4 a,.comment h4 a:hover,.comment-respond input[type=submit]:hover,.entry-content a:hover,.entry-title a:hover,.site-info a:hover,.site-title a:hover,.site-top ul li a:hover,.sorry li a:hover,i.iconfont.js-toggle-search.iconsearch:hover,span.page-numbers.current{color:<?php echo $theme?>}
        .comments .comments-main{display:block!important}
        .comments .comments-hidden{display:none!important}
        /*背景图片 todo 这里到时候替换一下*/
        .centerbg{background-image:url(<?php setting_index_navigation_background(); ?>);background-position:center center;background-attachment:inherit}
    </style>
    <?php
    lib_css_fancybox();
    lib_css_font_awesome();
    lib_css_font_awesome4();
    lib_script_jquery();
    lib_script_fancybox();
    lib_script_xiao_you();
    ?>
    <script>
        // /*Initial Variables*/
        var mashiro_global = new Object();
        // 主题设置
        var mashiro_option = new (function () {
            this.NProgressON = true;
            this.qzone_autocomplete = false;
            this.site_name = "<?php setting_index_site_name(); ?>";
            this.author_name = "<?php setting_index_site_name(); ?>";
            this.template_url = "<?php echo setting_template()?>";
            this.site_url = "<?php setting_index_site_url();?>";
            this.islogin = <?php tools_print_bool(user_is_login()); ?>;
            this.userInfo ='<?php echo json_encode($GLOBALS["userInfo"]) ?>';
        })();
        xy.net.server = '<?php echo "/api/v3/";?>';
        mashiro_option.jsdelivr_css_src = "<?php echo setting_template()."/static/css/lib.css"?>";
        mashiro_option.head_notice = "off";
        /*End of Initial Variables*/
        // 关闭console输出
        // console.log = function () {};
        console.info("%c 小游 %c", "background:#24272A; color:#ffffff", "", "https://xiaoyou66.com/");
        console.info("%c Github %c", "background:#24272A; color:#ffffff", "", "https://github.com/xiaoyou66");
        console.info("为什么控制台一直在报错呢？QAQ");
        mashiro_option.land_at_home = true;
    </script>
    <script type="text/javascript">
        if (!!window.ActiveXObject || "ActiveXObject" in window) {
            //is IE?
            alert("朋友，IE浏览器未适配哦~（QQ、360浏览器请关闭 IE 模式访问~）");
        }
    </script>
</head>
<body class="home blog hfeed chinese-font">
<div class="scrollbar" id="bar"></div>
<section id="main-container">
    <div class="headertop filter-dot">
        <div id="banner_wave_1"></div>
        <div id="banner_wave_2"></div>
        <figure id="centerbg" class="centerbg">
            <div class="focusinfo">
                <?php if (!xy_option("theme_sakura_show_avatar")): ?>
                    <h1 class="center-text glitch is-glitching Ubuntu-font" data-text="<?php setting_index_site_name(); ?>">
                        <?php setting_index_site_name(); ?>
                    </h1>
                <?php else: ?>
                    <div class="header-tou" ><img src="<?php echo xy_option("theme_sakura_avatar"); ?>"></div>
                <?php endif; ?>
                <div class="header-info">
                    <p><?php setting_index_description(); ?></p>
                    <div class="top-social_v2">
                        <?php
                            $template = setting_template();
                            if (xy_option("theme_sakura_sns_bili")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_bili")."\" target=\"_blank\" class=\"social-bili\" title=\"bilibili\"><img src=\"$template/static/image/sns/bilibili.png\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_csdn")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_csdn")."\" target=\"_blank\" class=\"social-csdn\" title=\"CSDN\"><img src=\"$template/static/image/sns/csdn.png\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_douban")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_douban")."\" target=\"_blank\" class=\"social-douban\" title=\"豆瓣\"><img src=\"$template/static/image/sns/douban.png\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_github")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_github")."\" target=\"_blank\" class=\"social-github\" title=\"github\"><img src=\"$template/static/image/sns/github.png\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_jianshu")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_jianshu")."\" target=\"_blank\" class=\"social-jianshu\" title=\"简书\"><img src=\"$template/static/image/sns/jianshu.png\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_qq")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_qq")."\" target=\"_blank\" class=\"social-qq\" title=\"QQ\"><img src=\"$template/static/image/sns/qq.png\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_qzone")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_qzone")."\" target=\"_blank\" class=\"social-qzone\" title=\"QQ空间\"><img src=\"$template/static/image/sns/qzone.png\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_telegram")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_telegram")."\" target=\"_blank\" class=\"social-telegram\" title=\"telegram\"><img src=\"$template/static/image/sns/telegram.svg\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_twitter")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_twitter")."\" target=\"_blank\" class=\"social-twitter\" title=\"Twitter\"><img src=\"$template/static/image/sns/twitter.png\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_wangyiyun")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_wangyiyun")."\" target=\"_blank\" class=\"social-wangyiyun\" title=\"网易云\"><img src=\"$template/static/image/sns/wangyiyun.png\"/></a></li>";
                            }
                            if (xy_option("theme_sakura_sns_zhihu")!=""){
                                echo "<li><a href=\"".xy_option("theme_sakura_sns_zhihu")."\" target=\"_blank\" class=\"social-zhihu\" title=\"知乎\"><img src=\"$template/static/image/sns/zhihu.png\"/></a></li>";
                            }
                        ?>
                    </div>
                </div>
            </div>
        </figure>
        <div id="video-container" style="">
            <video id="bgvideo" class="video" video-name="" src="" width="auto" preload="auto"></video>
            <div id="video-btn" class="loadvideo videolive"></div>
            <div id="video-add"></div>
            <div class="video-stu"></div>
        </div>
        <div class="headertop-down faa-float animated" onclick="headertop_down()">
            <span> <i class="fa fa-chevron-down" aria-hidden="true"></i> </span>
        </div>
    </div>
    <div id="page" class="site wrapper">
        <header id="navigation" class="site-header no-select is-homepage" role="banner">
            <div class="site-top">
                <div class="site-branding">
              <span class="site-title">
                <span class="logolink moe-mashiro">
                  <a href="<?php setting_index_site_url(); ?>" alt="<?php setting_index_site_name(); ?>">
                    <ruby>
                        <span class="sakuraso"><?php setting_index_site_name(); ?></span>
                    </ruby>
                  </a>
                </span>
              </span>
                </div>
                <div class="lower-cantiner">
                    <div class="lower">
                        <div id="show-nav" class="showNav mobile-fit">
                            <div class="line line1"></div>
                            <div class="line line2"></div>
                            <div class="line line3"></div>
                        </div>
                        <nav class="mobile-fit-control hide">
                            <ul id="menu-new" class="menu">
                                <?php foreach (setting_index_head_nav() as $head): ?>
                                    <li>
                                        <?php echo '<a href="'.$head["link"].'"><span class="faa-parent animated-hover">'.$head["title"].'</span></a>' ?>
                                        <?php if (count($head["children"])>0): ?>
                                        <ul class="sub-menu">
                                            <?php foreach ($head["children"] as $head1): ?>
                                                <?php echo '<li><span class="faa-parent animated-hover"><a href="'.$head1["link"].'">'.$head1["title"].'</span></a></li>' ?>
                                            <?php endforeach; ?>
                                        </ul>
                                        <?php endif; ?>
                                    </li>
                                <?php endforeach; ?>
                            </ul>
                        </nav>
                    </div>
                </div>
                <div class="header-user-avatar">
                    <?php if (user_is_login()): ?>
                        <a href="/admin">
                            <img class="faa-shake animated-hover" src="<?php user_avatar();?>" width="30" height="30"/>
                        </a>
                    <?php else: ?>
                    <a href="/access/login">
                        <img class="faa-shake animated-hover" src="<?php echo setting_template().'/static/image/none.png'?>" width="30" height="30"/>
                    </a>
                    <div class="header-user-menu">
                        <div class="herder-user-name no-logged">
                            <a href="/access/login"  style="color: #333;font-weight: bold;text-decoration: none;">登录</a>
                        </div>
                    </div>
                    <?php endif; ?>
                </div>
                <div class="searchbox">
                    <i class="iconfont js-toggle-search iconsearch icon-search"></i>
                </div>
            </div>
        </header>
        <div class="blank"></div>
