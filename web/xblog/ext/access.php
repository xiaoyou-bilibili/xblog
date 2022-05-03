<?php
/**
 * @description 登录界面板块常用函数
 * @author 小游
 * @date 2021/05/16
 */
use xblog\ext\Load;

// 登录界面还是注册界面
$accessPageName = "login";
// 获取登录注册界面
function xy_access(){
    global $accessPageName;
    Load::loadPage("/content/$accessPageName.php");
}
