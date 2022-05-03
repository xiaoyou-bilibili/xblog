<?php get_header('index'); ?>
<div id="kratos-blog-post">
    <div id="container" class="container">
        <div class="row">
            <section id="main" class="col-md-8">
                <?php xy_posts(); ?>
                <div id="post-page"></div>
            </section>
            <aside class="col-md-4 hidden-xs hidden-sm scrollspy">
                <div id="sidebar" class="affix-top wow bounceInRight animated" style="visibility: visible; animation-name: bounceInRight;">
                    <?php xy_right_side(); ?>
                </div>
            </aside>
        </div>
    </div>
</div>
<script>
    layui.use('laypage',function(){var laypage=layui.laypage;laypage.render({elem:'post-page',theme:'#51aded',count:<?php current_count();?>,limit:<?php current_size();?>,curr:<?php current_current();?>,jump:function(content,first){if(!first){window.location.href=xy.tools.updateQueryStringParameter(window.location.href,'page',content.curr)}}})});
</script>
<?php get_footer(); ?>