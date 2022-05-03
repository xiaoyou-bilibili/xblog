<?php

namespace xblog\request;

// 参考
// https://www.runoob.com/php/func-curl_setopt.html
// https://www.runoob.com/php/php-ref-curl.html
// 获取状态码参考 https://blog.csdn.net/xifeijian/article/details/17739155
use Exception;

/**
 * Class Request
 * @description 发送request请求
 * @author 小游
 * @date 2020/4/30
 * @package xblog\request
 */
class Request
{
    private $ch;
    // curl初始化
    public function __construct()
    {
        $this->ch = curl_init();
        // 设置超时时间
        curl_setopt($this->ch, CURLOPT_TIMEOUT, 200);
        // 服务端验证证书
        curl_setopt($this->ch, CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt($this->ch, CURLOPT_SSL_VERIFYHOST, false);
        // 获取到的值不显示出来
        curl_setopt($this->ch, CURLOPT_RETURNTRANSFER, 1);
    }

    /**
     *  处理请求错误
     * @param $res
     * @throws Exception 返回异常信息
     * @return mixed
     */
    private function errProcess($res)
    {
        $code = curl_getinfo($this->ch,CURLINFO_HTTP_CODE);
        if ($code>=200 && $code<=300){
            return $res;
        } else {
            if(isset($res["message"])){
                throw new Exception($res["message"],$code);
            } else {
                throw new Exception("未知错误",$code);
            }
        }
    }

    /**
     * 发送请求
     * @param string $url 请求地址
     * @return mixed
     */
    private function sendRequest(string $url, $head=[]){
        // 设置地址
        curl_setopt($this->ch,CURLOPT_URL,$url);
        // 设置头部信息
        if (count($head)>0){
            curl_setopt($this->ch, CURLOPT_HTTPHEADER, $head);
        }
        // 解析数据并返回(后面那个参数表示我们解析为数组)
        return json_decode(curl_exec($this->ch),true);
    }

    /**
     * 发送json请求
     * @param string $url 请求地址
     * @param string $method 请求方式
     * @param array $param  请求参数
     * @return mixed
     */
    private function sendJsonRequest(string $url, string $method, array $param=[]){
        $json = json_encode($param);
        // 设置地址
        curl_setopt($this->ch,CURLOPT_URL,$url);
        if ($method=="POST"){
            // 指定POST请求
            curl_setopt($this->ch, CURLOPT_POST, 1);
        }
        // 指定请求体
        curl_setopt($this->ch, CURLOPT_POSTFIELDS, $json);
        // 设置头部为json格式
        curl_setopt($this->ch, CURLOPT_HTTPHEADER, array(
            'Content-Type: application/json; charset=utf-8',
            'Content-Length: ' . strlen($json)
        ));
        // 发送请求
        return json_decode(curl_exec($this->ch),true);
    }

    /** 发送get请求
     * @param string $url 网址
     * @param array $param 参数
     * @throws Exception 请求出现错误
     * @return mixed
     */
    public function Get(string $url, array $param=[], $head=[])
    {
        $data = [];
        // 这里拼接一下参数
        foreach ($param as $key=>$value){
            array_push($data,"$key=$value");
        }
        $res = $this->sendRequest($url."?".join("&",$data),$head);
        return $this->errProcess($res);
    }

    /**
     * 发送JSON的post请求
     * @param string $url 网址
     * @param array $param 参数
     * @throws Exception
     */
    public function PostJson(string $url, array $param=[])
    {
        $res = $this->sendJsonRequest($url,"POST",$param);
        return $this->errProcess($res);
    }


    public function __destruct()
    {
        // 销毁对象时关闭连接
        curl_close($this->ch);
    }

}