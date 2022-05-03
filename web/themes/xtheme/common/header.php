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
    <meta name="referrer" content="no-referrer">
    <link rel="icon" type="image/x-icon" href="<?php setting_head_icon(); ?>">
    <title><?php setting_head_title(); ?></title>
    <meta name="robots" content="max-image-preview:large">
    <?php
        lib_css_font_awesome();
        lib_css_bootstrap();
        lib_css_element();
        lib_script_vue();
        lib_script_element();
        lib_script_jquery();
        lib_script_xiao_you();
    ?>
    <link rel="stylesheet" type="text/css" href="<?php echo setting_template()."/static/css/theme.css"?>">
    <style>
        .to-top {background-image: url(<?php echo setting_template()."/static/image/space-to-top.png"?>);}
        @font-face{font-family:DiaryFont;src:url(<?php echo setting_template()."/static/fonts/diary.ttf"?>)}
        .m-level{background-image:url(<?php echo setting_template()."/static/image/level.png"?>)}
    </style>
    <script>
        xy.net.server = '<?php echo "/api/v3/";?>';
        let globalIsLogin = <?php tools_print_bool(user_is_login()); ?>;
        let globalUserInfo = '<?php echo json_encode($GLOBALS["userInfo"]) ?>';
    </script>
<body>
<div>
    <div id="navigation">
        <div class="mini-header m-header">
            <div class="mini-header__content mini-header--login">
                <div class="nav-link">
                    <ul class="nav-link-ul">
                        <?php foreach (setting_index_head_nav() as $head): ?>
                            <li  class="nav-link-item">
                                <?php echo '<a class="link" href="'.$head["link"].'">'.$head["title"].'</a>' ?>
                            </li>
                        <?php endforeach; ?>
                    </ul>
                </div>
                <div id="search-box" class="nav-search-box">
                    <div class="nav-search">
                        <div id="nav_searchform">
                            <input v-model="keyword" type="text" placeholder="输入你想搜索的内容" class="nav-search-keyword" />
                            <div class="nav-search-btn">
                                <button type="submit" class="nav-search-submit" @click='search'>
                                    <i class="el-icon-search"></i>
                                </button>
                            </div>
                        </div>
                        <ul class="search-suggest suggest" style="display: none">
                            <template v-for="(item,index) in keywords">
                                <li :key="index" class="suggest-item">
                                    <a :href="'/archives/'+item.id">
                                        {{ item.title }}
                                    </a>
                                </li>
                            </template>
                        </ul>
                    </div>
                </div>
                <div>
                    <?php if (user_is_login()): ?>
                    <div>
                        <div class="nav-user-center">
                            <div class="user-con signin">
                                <div class="item">
                                    <div id="loginavatar" class="mini-avatar van-popover__reference">
                                        <img alt src="<?php user_avatar(); ?>">
                                    </div>
                                </div>
<!--                                <div class="item">-->
<!--                                <span>-->
<!--                                  <div class="mini-vip van-popover__reference" aria-describedby="van-popover-2930" tabindex="0">-->
<!--                                    <span class="name">新功能开发中</span>-->
<!--                                  </div>-->
<!--                                </span>-->
<!--                                </div>-->
                            </div>
                            <div>
                              <span>
                                <a href="/admin"><span class="mini-upload van-popover__reference">个人中心</span></a>
                              </span>
                            </div>
                        </div>
                    </div>
                    <?php else: ?>
                    <div>
                        <div class="nav-user-center">
                            <div class="user-con logout">
                                <div class="item">
                                    <div class="mini-login van-popover__reference" aria-describedby="van-popover-6081" tabindex="0">
                                      <span class="name">
<!--                                        <img alt src="//static.hdslb.com/images/akari.jpg" class="logout-face">-->
                                        <a href="/access/login" class="txt name">登录</a>
                                        <a href="/access/register" class="txt name">注册</a>
                                      </span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <?php  endif;?>
                </div>
            </div>
        </div>
        <div style="background-image:url(<?php setting_index_navigation_background(); ?>)" class="bili-banner">
            <div class="taper-line" ></div>
            <a href="#" target="_blank" class="banner-link"></a>
        </div>
    </div>

