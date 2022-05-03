<?php
/**
 * Class Handle
 * @description 路由处理类
 * @author 小游
 * @date 2020/4/30
 * @package xblog\route
 */
namespace xblog\route;

use xblog\api\ApiServer;
use xblog\error\Error;
use xblog\ext\Load;

class Handle
{
    /**
     * 请求方式
     * @var string
     */
    private $method = '';

    /**
     * 请求地址
     * @var string
     */
    private $url = '';


    public function __construct()
    {
        // 获取相关参数
        $this->url = explode('?', $_SERVER['REQUEST_URI'], 2)[0];
        $this->method = $_SERVER["REQUEST_METHOD"];
        // 处理请求
        $this->urlProcess();
    }

    /**
     * 处理URL地址
     */
    private function urlProcess()
    {
        $result = explode('/', $this->url);
        if (count($result) <= 2 && $result[1] == '') {
            $this->indexProcess();
            return;
        }
        switch ($result[1]) {
            case "archives":
                $this->postProcess($result[2]);
                break;
            case "doc":
                $this->docProcess($result[2]);
                break;
            case "diary":
                $this->diaryProcess();
                break;
            case "more":
                $this->moreProcess($result[2]);
                break;
            case "access":
                $this->loginProcess($result[2]);
                break;
            case "sitemap":
                $this->siteMapProcess();
                break;
            case "php":
                $this->apiServerProcess($this->method.$this->url);
                break;
            default:
                Error::Code404();
        }
    }

    /**
     *  处理主页信息
     */
    private function indexProcess()
    {
        Load::showIndex();
    }

    /**
     * 处理文章请求
     * @param $id int 文章id
     */
    private function postProcess(int $id)
    {
        if (is_numeric($id)) {
            Load::showPost($id);
        } else {
            Error::Code404();
        }
    }

    /**
     * 处理文档请求
     * @param $id mixed 文档id
     */
    private function docProcess($id)
    {
        if (is_numeric($id)) {
            Load::showDoc($id);
        } else {
            Load::showDoc(0);
        }
    }

    /**
     *  处理日记请求
     */
    private function diaryProcess()
    {
        Load::showDiary();
    }

    /**
     * 处理页面请求
     * @param $name string 页面名字
     */
    private function moreProcess(string $name)
    {
        Load::showMore($name);
    }

    /**
     *  处理登录请求
     */
    private function loginProcess($option)
    {
        Load::showLogin($option);
    }

    /**
     *  获取网站的站点地图
     */
    private function siteMapProcess()
    {
        Load::showSiteMap();
    }

    private function apiServerProcess($url){
        new ApiServer($url);
    }

}