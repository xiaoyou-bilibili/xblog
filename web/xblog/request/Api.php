<?php
namespace xblog\request;

use Exception;

/**
 * Class Api
 * @description 所有的API信息，这里设计为单例模式，目的是为了复用curl链接
 * @date 2020/4/30
 * @author 小游
 * @package xblog\request
 */
class Api {
    static public $api = null;
    const API_VERSION = API."/api/v3/";
    const SETTING = self::API_VERSION."settings";
    const POSTS = self::API_VERSION."posts";
    const PLUGINS = self::API_VERSION."plugins";
    const TOOLS = self::API_VERSION."tools";
    const USER = self::API_VERSION."user";

    const GET_SETTING_INDEX = self::SETTING."/site/index";
    const GET_SETTING_PLUGIN = self::SETTING."/site/plugins/";
    const GET_SETTING_LOGIN = self::SETTING."/site/login";
    const GET_THEME_SETTING = self::SETTING."/themes/";


    const GET_POST_LIST = self::POSTS;
    const GET_POST_CONTENT = self::GET_POST_LIST."/";
    const GET_POST_CATEGORY = self::POSTS."/category";

    const GET_PLUGINS_DIARY = self::PLUGINS."/diary";
    const GET_PLUGINS_DOC = self::PLUGINS."/docs";
    const GET_PLUGINS_Friends = self::PLUGINS."/friends";

    const GET_TOOLS_SITEMAP = self::TOOLS."/sitemap";
    const GET_TOOLS_SMILES = self::TOOLS."/smiles";


    const GET_USER_USERNAME = self::USER."/username";
    const POST_USER_TOKEN = self::USER."/token";
    const POST_USER_FORGET_EMAIL = self::USER."/password/email";



    public $client;

    public function __construct()
    {
        $this->client = new Request();
    }

    public static function newAPi(): ?Api
    {
        // 判断一下$api是否存在,不存在则初始化
        if (self::$api==null){
            self::$api = new Api();
//            self::$api->client = new Request();
        }
        return self::$api;
    }


    /* 设置板块 */
    /**
     *  获取主页设置
     * @param  $page mixed 页面类型
     * @return array
     * @throws Exception
     */
    function getSettingIndex($page): array
    {
        return $this->client->Get(self::GET_SETTING_INDEX,["page"=>$page]);
    }

    /**
     *  获取插件的设置
     * @param $name
     * @return array
     * @throws Exception
     */
    function getSettingPlugin($name): array
    {
        return $this->client->Get(self::GET_SETTING_PLUGIN.$name);
    }

    /**
     * 获取注册界面设置
     * @return array
     * @throws Exception
     */
    function getSettingLogin():array
    {
        return $this->client->Get(self::GET_SETTING_LOGIN);
    }

    /**
     *  获取主题设置
     * @param mixed $name 设置的值
     * @return array
     * @throws Exception
     */
    function getSettingThemes($name):array{
        return $this->client->Get(self::GET_THEME_SETTING.$name);
    }


    /* 文章板块 */
    /**
     * 获取文章列表
     * @param array $param 参数
     * @return array
     * @throws Exception
     */
    public function getPostList(array $param): array
    {
        return $this->client->Get(self::GET_POST_LIST,$param);
    }

    /**
     * 获取文章内容
     * @param string $id 文章id
     * @return array 返回请求结果
     * @throws Exception 请求出现错误
     */
    function getPostContent(string $id): array
    {
        return $this->client->Get(self::GET_POST_CONTENT.$id);
    }

    /**
     * 获取加密文章内容
     * @param string $id
     * @param string $password
     * @return array
     * @throws Exception
     */
    function getPostEncryptContent(string $id,string $password): array
    {
        return $this->client->Get(self::GET_POST_CONTENT.$id."/encryption",["password"=>$password]);
    }

    /**
     * 获取文章分类
     * @return array
     * @throws Exception
     */
    function getPostCategory()
    {
        return $this->client->Get(self::GET_POST_CATEGORY);
    }

    /**
     *  获取文章评论
     * @param $id
     * @return array
     * @throws Exception
     */
    function getPostComment($id): array
    {
        return $this->client->Get(self::GET_POST_CONTENT.$id."/comments");
    }

    /* 插件板块 */
    /**
     * 获取日记
     * @param $param
     * @return array
     * @throws Exception
     */
    function getPluginsDiary($param): array
    {
        return $this->client->Get(self::GET_PLUGINS_DIARY,$param);
    }

    /**
     * 获取友链
     * @return array
     * @throws Exception
     */
    function getFriends(): array
    {
        return $this->client->Get(self::GET_PLUGINS_Friends);
    }

    /**
     * 获取文档内容
     * @param int $id 文档id
     * @return array
     * @throws Exception
     */
    function getPluginsDocContent(int $id): array
    {
        return $this->client->Get(self::GET_PLUGINS_DOC."/$id");
    }

    /* 工具板块 */
    /**
     * 获取网站的站点地图
     * @return array
     * @throws Exception
     */
    function getToolsSitemap(): array
    {
        return $this->client->Get(self::GET_TOOLS_SITEMAP);
    }
    /**
     * 获取表情数据
     * @return array
     * @throws Exception
     */
    function getToolsSmile(): array
    {
        return $this->client->Get(self::GET_TOOLS_SMILES);
    }

    /*用户板块*/
    /**
     * 判断用户是否注册过了
     * @param $param
     * @return array
     * @throws Exception
     */
    function getUserUsername($param): array
    {
        return $this->client->Get(self::GET_USER_USERNAME,$param);
    }

    /**
     * 用户注册
     * @param $param
     * @return array
     * @throws Exception
     */
    function userRegister($param): array
    {
        return $this->client->PostJson(self::USER,$param);
    }

    /**
     * 用户登录
     * @param $param
     * @return array
     * @throws Exception
     */
    function userLogin($param): array
    {
        return $this->client->PostJson(self::POST_USER_TOKEN,$param);
    }


    /**
     * 发送找回密码的邮件
     * @param $param
     * @return array
     * @throws Exception
     */
    function userSendForgetEmail($param): array
    {
        return $this->client->PostJson(self::POST_USER_FORGET_EMAIL,$param);
    }

    /**
     * 获取用户信息
     * @return array
     * @throws Exception
     */
    function getInfo(): array
    {
        if (isset($_COOKIE["token"])){
            $cookie = json_decode($_COOKIE["token"],true);
            $id = $cookie["user_id"];
            $token = $cookie["token"];
            $head = array("user_id: ".$id, "token: ".$token,);
            return $this->client->Get(self::USER,[],$head);
        } else {
            throw new Exception("没有token信息");
        }
    }

}