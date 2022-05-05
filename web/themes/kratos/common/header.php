<head>
    <meta charset="UTF-8">
    <!--      pjax强制重载-->
    <meta http-equiv="Cache-Control" content="no-transform">
    <meta http-equiv="Cache-Control" content="no-siteapp">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1,maximum-scale=1">
    <meta name="format-detection" content="telphone=no,email=no">
    <meta name="keywords" content="<?php setting_head_keyword(); ?>">
    <meta name="description" itemprop="description" content="<?php setting_head_description(); ?>">
    <meta property="og:title" content="<?php setting_head_title(); ?>">
    <meta property="og:site_name" content="<?php setting_head_title(); ?>">
    <meta property="og:type" content="website">
    <meta property="og:description" content="<?php setting_head_description(); ?>">
    <meta property="og:url" content="<?php setting_head_url(); ?>">
    <meta name="twitter:title" content="<?php setting_head_title(); ?>">
    <meta name="twitter:description" content="<?php setting_head_description(); ?>">
    <meta name="twitter:card" content="summary">
    <!-- 允许访问站外资源 -->
    <meta name="referrer" content="same-origin">
    <link rel="icon" type="image/x-icon" href="<?php setting_head_icon(); ?>">
    <title><?php setting_head_title(); ?></title>
    <meta name="robots" content="max-image-preview:large">
    <?php
        lib_css_font_awesome();
        lib_css_bootstrap();
        lib_css_layui();
        lib_script_jquery();
        lib_script_highlight();
        lib_script_layui();
        lib_script_xiao_you();
    ?>
    <link rel="stylesheet" href="<?php echo setting_template()."/static/css/kratos.min.css"?>" type="text/css">
    <style>#offcanvas-menu{background:#3a3f51}.affix{top:61px}.kratos-cover.kratos-cover_2{background-image:url(<?php setting_index_navigation_background(); ?>)}@media(max-width:768px){.kratos-cover.kratos-cover_2{background-image:url(<?php setting_index_background(); ?>)}}@media(max-width:768px){#kratos-header-section{display:none}nav#offcanvas-menu{top:0;padding-top:190px;}.kratos-cover .desc.desc2{margin-top:-55px}}@media(min-width:768px){.pagination>li>a{background-color:rgba(255,255,255,.9)}.comment-list .children li{background-color:rgba(255,253,232,.7)!important}.theme-bg{background-image:url(<?php setting_index_background(); ?>);background-size:cover;background-attachment:fixed}}    </style>
    <link rel="stylesheet" type="text/css" href="<?php echo setting_template()."/static/css/prism.css"?>">
    <link rel="stylesheet" type="text/css" href="<?php echo setting_template()."/static/css/animate.min.css"?>">
    <script>
        xy.net.server = '<?php echo "/api/v3/";?>';
    </script>
<body>
<div class="theme-bg"></div>
<div id="kratos-wrapper" style="left: 0; position: relative;">
    <div id="kratos-page">
        <div id="navigation">
            <div id="sider-bar" style="visibility: hidden;">
                <nav id="offcanvas-menu" class=" " style="height: 969px; left: -240px;">
                    <ul id="" class="ul-me"><li class="current-menu-item">
                            <a href="http://127.0.0.1/" aria-current="page">首页</a></li>
                        <li><a href="http://xiaoyou66.com">个人网站</a>
                            <ul class="sub-menu">
                                <li><a href="http://测试">网站子页面</a></li>
                            </ul>
                        </li>
                    </ul>
                </nav>
                <div class="clearfix  text-center hide  show" id="aside-user">
<!--                    <div class="dropdown wrapper">-->
<!--                        <div ui-nav="">-->
<!--                            <a href="https://xiaoyou66.com/about/">-->
<!--                                    <span class="thumb-lg w-auto-folded avatar m-t-sm">-->
<!--                                        <img src="http://cdn.xiaoyou66.com/theme/admin_avatar-96x96.jpg" class="img-full">-->
<!--                                    </span>-->
<!--                            </a>-->
<!--                        </div>-->
<!--                        <span class="clear">-->
<!--                              <span class="block m-t-sm">-->
<!--                                <strong class="font-bold text-lt">小游</strong>-->
<!--                              </span><br>-->
<!--                              <span class="text-muted text-xs block">二次元技术宅</span>-->
<!--                            </span>-->
<!--                    </div>-->
                    <div class="line dk hidden-folded"></div>
                </div>
            </div>
            <div id="kratos-header">
                <div class="nav-toggle"><a class="kratos-nav-toggle js-kratos-nav-toggle"><i></i></a></div>
                <header id="kratos-header-section">
                    <div class="container">
                        <div class="nav-header">
                            <nav id="kratos-menu-wrap" class="menu-menu-1-container"><ul id="kratos-primary-menu" class="sf-menu">
                                    <?php foreach (setting_index_head_nav() as $head): ?>
                                    <li>
                                        <?php echo '<a href="'.$head["link"].'">'.$head["title"].'</a>' ?>
                                        <ul class="sub-menu">
                                            <li>
                                                <?php foreach ($head["children"] as $head1): ?>
                                                    <?php echo '<a href="'.$head1["link"].'">'.$head1["title"].'</a>' ?>
                                                <?php endforeach; ?>
                                            </li>
                                        </ul>
                                    </li>
                                    <?php endforeach; ?>
                                </ul>
                            </nav>
                        </div>
                    </div>
                </header>
            </div>
            <div class="kratos-start kratos-hero-2">
                <div class="kratos-overlay"></div>
                <div class="kratos-cover kratos-cover_2 text-center">
                    <div class="desc desc2 animate-box">
                        <a href="<?php setting_index_site_url(); ?>"><h2><?php setting_index_site_name(); ?></h2><br><span><?php setting_index_description(); ?></span></a>
                    </div>
                </div>
            </div>
        </div>