<?php setting_access(); ?>
<html lang="zh-CN"><head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title><?php setting_access_site_name(); ?></title>
    <meta name="robots" content="max-image-preview:large, noindex, noarchive">
    <meta name="viewport" content="width=device-width">
    <?php
        lib_css_element();
    ?>
    <link rel="stylesheet" type="text/css" href="<?php echo setting_template()."/static/css/access.css"?>">
</head>
<body>
<div id="app">
    <div id="background-slideshow">
        <div class="slide">
            <div class="wallpaper" style="background-image:url(<?php setting_access_background(); ?>)" ></div>
        </div>
    </div>
    <div class="signup-form ">
        <div class="signup-form__logo-box">
            <div class="signup-form__logo" style="background-image:url(<?php setting_access_logo(); ?>)" ></div>
            <div class="signup-form__catchphrase">
                <?php setting_access_web_text(); ?>
            </div>
        </div>
        <?php xy_access(); ?>
    </div>
    <div id="footer">
        <div class="footer-item">
            <a href="/" target="_blank">2021©<?php setting_access_site_name(); ?></a>
        </div>
    </div>
</div>
<?php
lib_script_jquery();
lib_script_vue();
lib_script_element();
lib_script_xiao_you();
?>
<script>
    xy.net.server = '<?php echo "/api/v3/";?>';
</script>
<script type="text/javascript" src="<?php echo setting_template()."/static/js/access.js"?>"></script>
</body>
</html>