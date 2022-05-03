<?php setting_access(); ?>
<html lang="zh-CN"><head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title><?php setting_access_site_name(); ?></title>
    <meta name="robots" content="max-image-preview:large, noindex, noarchive">
    <link rel="stylesheet" id="wp-admin-css" href="<?php echo setting_template().'/static/css/customlogin.min.css'?>" type="text/css">
    <style>body{background:#92C1D1 url(<?php setting_access_background(); ?>) fixed center top no-repeat!important;background-size:cover!important}.login h1 a{background-image:url(<?php setting_access_logo(); ?>)!important}</style>	<meta name="referrer" content="strict-origin-when-cross-origin">
    <meta name="viewport" content="width=device-width">
</head>
<body class="login js login-action-login wp-core-ui  locale-zh-cn">
<?php xy_access() ?>
<div class="clear"></div>
</body></html>