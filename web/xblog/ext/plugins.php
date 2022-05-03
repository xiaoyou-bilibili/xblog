<?php
/**
 * @description 插件板块设置
 * @author 小游
 * @date 2021/05/01
 */

use xblog\ext\Load;
use xblog\request\Api;

// 日记数据
$diary = [];
// 文档数据
$docs = [];

/**
 * 加载日记
 */
function xy_diary(){
    global $diary,$current;
    try {
        // 获取文章列表
        $data = Api::newAPi()->getPluginsDiary($_GET);
        $current = $data;
        foreach ($data["contents"] as $item){
            $diary = $item;
            Load::loadPage("/content/diary.php");
        }
    } catch (Exception $e) {

    }
}


/*日记相关*/
function get_plugins_diary_id(){
    return $GLOBALS['diary']["diary_id"];
}
function plugins_diary_content(){
    echo $GLOBALS['diary']["content"];
}
function plugins_diary_date(){
    echo $GLOBALS['diary']["date"];
}
function plugins_diary_comment(){
    echo $GLOBALS['diary']["comment"];
}
function plugins_diary_good(){
    echo $GLOBALS['diary']["good"];
}
function plugins_diary_avatar(){
    echo $GLOBALS['diary']["avatar"];
}
function plugins_diary_nickname(){
    echo $GLOBALS['diary']["nickname"];
}
function get_plugins_diary_encrypt(){
    return $GLOBALS['diary']["encrypt"];
}

/*文档相关*/
function plugins_doc_title(){
    if (isset($GLOBALS['docs'])){
        echo $GLOBALS['docs']["title"];
    }
}
function plugins_doc_content(){
    if (isset($GLOBALS['docs'])) {
        echo $GLOBALS['docs']["content"];
    }
}