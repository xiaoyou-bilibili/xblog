<?php
get_header("plugin");
// 加载CSS样式
more_css();
// 加载js样式
more_script();
?>
    <div id="bg" style="background-image: url(<?php more_background(); ?>)" ></div>
<?php
more_content();
get_footer();
