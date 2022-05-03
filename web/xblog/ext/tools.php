<?php
/**
 * @description 工具板块处理函数
 * @author 小游
 * @date 2021/05/02
 */

use xblog\request\Api;

$sitemap = [];

/*站点地图*/
function tools_sitemap_site(){
    echo $GLOBALS["sitemap"]["site"];
}
function tools_sitemap_map(){
    echo $GLOBALS["sitemap"]["map"];
}
function get_tools_sitemap_post(){
    return $GLOBALS["sitemap"]["post"];
}
function get_tools_sitemap_doc(){
    return $GLOBALS["sitemap"]["doc"];
}
function get_tools_sitemap_category(){
    return $GLOBALS["sitemap"]["category"];
}
function get_tools_sitemap_tag(){
    return $GLOBALS["sitemap"]["tag"];
}

// 获取表情
function tools_get_smile()
{
    try {
        return Api::newAPi()->getToolsSmile();
    } catch (Exception $e) {
        return [];
    }
}

// 打印布尔值
function tools_print_bool(bool $bool){
    if ($bool){
        echo 'true';
    } else {
        echo 'false';
    }
}
// 键值是否存在并不为空
function key_post_not_empty($key): bool
{
    if (isset($_POST[$key]) && !empty($_POST[$key])){
        return true;
    }
    return false;
}