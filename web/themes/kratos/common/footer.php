<!-- 动画脚本-->
<script src="<?php echo setting_template()."/static/js/wow.min.js"?>"></script>
<footer>
    <div id="footer" style="background:rgba(35,40,45,1)">
        <div class="container">
            <div class="row">
                <div class="col-md-6 col-md-offset-3 footer-list text-center">
                    <p class="kratos-social-icons">
                    </p>
                    <p>
                        © 2021 <a href="<?php setting_index_site_url();?>"><?php setting_index_site_name();?></a>. All Rights Reserved.
                        <br>本站已萌萌哒(✪ω✪)运行<span id="span_dt_dt"></span>
                        <br>Theme
                        <a href="" target="_blank" rel="nofollow">Kratos</a> Power by <a href="https://xblog.xiaoyou66.com/" target="_blank" rel="nofollow">XBlog</a>
                        <br><a href="https://beian.miit.gov.cn/" rel="external nofollow" target="_blank"><i class="govimg"></i><?php setting_index_bei_an(); ?></a>
                        <br><a href="/sitemap" target="_blank">Sitemap</a>
                    </p>
                </div>
            </div>
        </div>
        <div class="cd-tool text-center">
            <div class="gotop-box"><div class="gotop-btn"><span class="fa fa-chevron-up"></span></div></div>
            <div class="search-box">
                <span class="fa fa-search"></span>
                <form class="search-form" role="search" method="get" id="searchform" action="<?php echo setting_web()?>">
                    <input type="text" name="q" id="search" placeholder="请输入关键词..." style="display:none">
                </form>
            </div>
        </div>
    </div>
</footer>
</div>
</div>
<script type="text/javascript" id="kratos-js-extra">
    /* <![CDATA[ */
    var xb = {"api":"","theme":"<?php echo setting_template() ?>","web":"<?php echo setting_web(); ?>","ctime":"<?php setting_index_build_time(); ?>"};
    /* ]]> */
</script>
<?php lib_script_layer(); ?>
<script type="text/javascript" src="<?php echo setting_template().'/static/js/theme.min.js' ?>" id="theme-js"></script>
<script type="text/javascript" src="<?php echo setting_template()."/static/js/kratos.js"?>"></script>
<!--<script type="text/javascript" src="http://127.0.0.1/wp-content/themes/kratos/static/js/pjax.js?ver=1.0" id="pjax-js"></script>-->
</body>