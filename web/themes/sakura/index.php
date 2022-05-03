<?php get_header('index');?>
<?php
// 这个是聚焦的内容，目前不设置这个功能，后续再加上
//    <div class="top-feature-row">
//        <h1 class="fes-title" style="font-family: 'Ubuntu', sans-serif">
//            <i class="fa fa-anchor" aria-hidden="true"></i> 聚焦
//        </h1>
//        <div class="top-feature-v2">
//            <div class="the-feature square from_left_and_right">
//                <a href="/theme-sakura/" target="_blank"
//                ><div class="img">
//                        <img src="https://cdn.jsdelivr.net/gh/moezx/cdn@3.3.2/img/other/sakura.md.png"/>
//                    </div>
//                    <div class="info">
//                        <h3>Sakura</h3>
//                        <p>本站 WordPress 主题</p>
//                    </div>
//                </a>
//            </div>
//        </div>
//    </div>
?>
<div class="main-part">
    <?php $show = xy_option("theme_sakura_show_side");if($show):?>
    <div class="col-xl-2 col-lg-3 col-md-10 col-sm-3 col-0 div-box-content">
        <div id="sidebar-left" class="giligili-left">
            <?php xy_left_side(); ?>
        </div>
    </div>
    <?php endif;?>
    <div id="content" class="col-xl-5 col-lg-5 col-md-10 col-sm-2 col-11  site-content">
        <div id="primary" class="content-area">
            <main id="main" class="site-main" role="main">
                <h1 class="main-title" style="font-family: 'Ubuntu', sans-serif">
                    <i class="fa fa-envira" aria-hidden="true"></i> 最新文章
                </h1>
                <?php xy_posts(); ?>
            </main>
            <div id="pagination">
                <?php
                    $href = "";
                    if (get_current_current()<get_current_count()) {
                        $href = "href=\"?page=".(get_current_current()+1);
                        if (isset($_GET["q"])){ $href.="&q=".$_GET["q"]; }
                        if (isset($_GET["tag"])){ $href.="&tag=".$_GET["tag"]; }
                        if (isset($_GET["category"])){ $href.="&category=".$_GET["category"]; }
                        $href.="\"";
                    }
                    echo "<a ".$href.">下一页</a>"
                ?>
            </div>
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
<?php get_footer(); ?>