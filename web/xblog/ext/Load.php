<?php

namespace xblog\ext;

use xblog\error\Error;
use xblog\request\Api;
use Exception;


// 加载必要文件
require __DIR__ . "/../../themes/index.php";
require __DIR__ . "/setting.php";
require __DIR__ . "/posts.php";
require __DIR__ . "/widget.php";
require __DIR__ . "/more.php";
require __DIR__ . "/plugins.php";
require __DIR__ . "/lib.php";
require __DIR__ . "/tools.php";
require __DIR__ . "/user.php";
require __DIR__ . "/access.php";
require __DIR__ . "/index.php";


class Load
{
    // 主题路径
    const PATH = __DIR__ . "/../../themes/".THEME;

    /**
     * 加载页面
     * @param string $dir
     * @return bool
     */
    public static function loadPage(string $dir): bool
    {
        $filename = self::PATH.$dir;
        if (is_file($filename)){
            require $filename;
            return true;
        } else {
            return false;
        }
    }


    /**
     * 显示一些页面信息
     * @param string $message 信息内容
     */
    private static function showMessage(string $message){
        echo "<h1>".$message."</h1>";
    }


    /**
     * 错误信息处理
     * @param Exception $e 错误信息
     */
     private static function errProcess(Exception $e){
        if ($e->getCode()==404){
            Error::Code404();
        }
    }


    /**
     *  显示主页信息
     */
    public static function showIndex(){
        self::loadPage("/index.php");
    }

    /**
     * 显示文章界面
     * @param string $id 文章id
     */
    public static function showPost(string $id){
        global $post;
        // 发送请求获取主页数据
        try {
            $post = Api::newAPi()->getPostContent($id);
            // 判断是否有密码
            if (isset($_GET["password"])){
                $password = $_GET["password"];
                try {
                    $content = Api::newAPi()->getPostEncryptContent($id,$password);
                    // 覆盖内容
                    $post["encrypt"] = false;
                    $post["content"] = $content["content"];
                }catch (Exception $e){}
            }
            self::loadPage("/post.php");
        } catch (Exception $e) {
            self::errProcess($e);
        }
    }

    /**
     *  显示日记界面
     */
    public static function showDiary(){
        self::loadPage("/diary.php");
    }

    /**
     *  显示文档界面
     */
    public static function showDoc($id){
        try {
            if ($id!=0){
                global $docs;
                // 获取文档内容
                $docs = Api::newAPi()->getPluginsDocContent($id);
            }
            self::loadPage("/doc.php");
        } catch (Exception $e) {
            self::errProcess($e);
        }
    }

    /**
     * 显示更多界面
     * @param string $name 界面名字
     */
    public static function showMore(string $name){
        // 先判断这个文件是否存在，不存在则从接口获取
        $tmp = "/more/".$name.".php";
        $filename = self::PATH.$tmp;
        if (is_file($filename)){
            self::loadPage($tmp);
            return;
        }
        global $more;
        // 发送请求获取主页数据
        try {
            $more = Api::newAPi()->getSettingPlugin($name);
            self::loadPage("/more.php");
        } catch (Exception $e) {
            self::errProcess($e);
        }
    }


    /**
     *  显示登录界面
     */
    public static function showLogin($option){
        // 判断用户是否登录
        if (user_is_login()){
            header("location: ".setting_web());
            exit();
        }
        global $errorMessage,$okMessage,$accessPageName;
        if ($option == 'login') {
            // 获取参数
            if (key_post_not_empty("username")) {
                // 获取主页设置
                setting_index('access');
                if (!key_post_not_empty("password")){
                        $errorMessage = "用户名或密码不能为空!";
                } else {
                    // 登录
                    try{
                        $token = Api::newAPi()->userLogin($_POST);
                        // 保存token信息
                        if (key_post_not_empty("remember")){
                            $time = time()+60*60*24*30;
                            setcookie("token",json_encode($token),$time,"/");
                        } else {
                            setcookie("token",json_encode($token),0,"/");
                        }
                        $okMessage = "登录成功！正在跳转主页！";
                        // 重定向到主页
                        echo '<META HTTP-EQUIV="refresh" CONTENT="2;url='.setting_web().'">';
                    }catch (Exception $e) {
                        $errorMessage = $e->getMessage();
                    }
                }
            }
        } else if ($option == 'register'){
            // 获取参数
            if (key_post_not_empty("username")) {
                // 先判断参数
                if (!key_post_not_empty("nickname")){
                    $errorMessage = "昵称不能为空！";
                } else if (!key_post_not_empty("password") || !key_post_not_empty("confirm") || $_POST["password"]!=$_POST["confirm"]){
                    $errorMessage = "两次密码不一致！";
                } else if (!key_post_not_empty("password")){
                    $errorMessage = "邮箱不能为空！";
                } else {
                    // 先判断用户名是否注册过
                    try{
                        $_POST["user"] = $_POST["username"];
                        Api::newAPi()->getUserUsername($_POST);
                        $errorMessage = "用户名或邮箱已注册！";
                    }catch (Exception $e){
                        if ($e->getMessage()!="用户名或邮箱未注册"){
                            $errorMessage = $e->getMessage();
                        } else {
                            // 开始注册
                            try {
                                Api::newAPi()->userRegister($_POST);
                                $okMessage = "注册成功!";
                            }catch (Exception $e){
                                $errorMessage = $e->getMessage();
                            }
                        }
                    }
                }
            }
        } else if ($option == 'forget'){
            try {
                Api::newAPi()->userSendForgetEmail($_POST);
                $okMessage = "邮件已发送，请注意查收!";
            }catch (Exception $e){
                $errorMessage = $e->getMessage();
            }
        }
        $accessPageName = $option;
        self::loadPage("/access.php");
    }

    /**
     *  显示站点地图
     */
    public static function showSiteMap(){
        global $sitemap;
        try {
            $sitemap = Api::newAPi()->getToolsSitemap();
            self::loadPage("/sitemap.php");
        } catch (Exception $e) {
            Error::Code500();
        }
    }

    /**
     * 显示错误页面
     * @param int $code 错误码
     */
    public static function showError(int $code){
        if (!self::loadPage("/error/".$code.".php")){
            self::showMessage($code."错误");
        }
    }

}