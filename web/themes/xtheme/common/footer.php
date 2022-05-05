    <div id="footer" class="footers">
        <div class="footer-item">
            <p>
                ©2020<a href="<?php setting_index_site_url();?>"><?php setting_index_site_name();?></a>. All Rights Reserved.|本站已萌萌哒(✪ω✪)运行<span id="span_dt_dt"></span>
            </p>
        </div>
        <div class="footer-item">
            <p>
                Powered by <a target="_blank" href="https://xblog.xiaoyou66.com">XBlog</a>|
                <a href="/sitemap" rel="external nofollow" target="_blank">SITEMAP |</a>
                <span style="color: rgba(255, 255, 255, .5);">
                    <a href="https://beian.miit.gov.cn/">
                        <?php setting_index_bei_an(); ?>
                    </a>
                    <?php if (get_setting_index_gov_bei_an()!=""):  ?>
                        <span>
                            <img src="<?php echo setting_template().'/static/image/beian.png' ?>">
                            <a href="http://www.beian.gov.cn/portal/index.do" target="_blank"> <?php setting_index_gov_bei_an(); ?></a>
                        </span>
                    <?php endif; ?>
                </span>
            </p>
        </div>
        <div id="to-top" class="to-top" style="background-position-x: -40px;"></div>
    </div>
</div>
<script>
var globalBuildTime = '<?php setting_index_build_time(); ?>';
</script>
<script type="text/javascript" src="<?php echo setting_template()."/static/js/theme.js"?>"></script>
</body>