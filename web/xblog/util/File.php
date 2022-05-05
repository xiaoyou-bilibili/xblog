<?php
/**
 * Class File
 * @description 文件读取函数
 * @author 小游
 * @date 2021/07/14
 * @package /xblog/util
 */
namespace xblog\util;

use Exception;
use ZipArchive;

class File
{
    /**
     * 读取文件
     * @param string $filename 文件名
     * @param bool $parse 是否解析为JSON格式
     * @return mixed
     * @throws Exception
     */
    static function readFile(string $filename,bool $parse=false){
        // 读取主题的json文件
        if(is_file($filename)){
            // 读取文件
            $file = fopen($filename,"r");
            $content = fread($file,filesize($filename));
            fclose($file);
            if ($parse){
                // 解析为JSON数据
                return json_decode($content,true);
            } else {
                return $content;
            }
        }
        throw new Exception("文件不存在!");
    }

    /**
     * 下载文件
     * @param string $url 下载地址
     * @param string $file 文件名
     */
    static function downloadFile(string $url, string $file)
    {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_POST, 0);
        curl_setopt($ch,CURLOPT_URL,$url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        // 服务端验证证书
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
        $file_content = curl_exec($ch);
        curl_close($ch);
        $downloaded_file = fopen($file, 'w');
        fwrite($downloaded_file, $file_content);
        fclose($downloaded_file);
    }

    /**
     * 解压文件
     * @param $fromName
     * @param $toName
     * @return bool
     */
    static function unzip($fromName, $toName): bool
    {
        if(!file_exists($fromName)){
            return false;
        }
        $zipArc = new ZipArchive();
        if(!$zipArc->open($fromName)){
            return false;
        }
        if(!$zipArc->extractTo($toName)){
            $zipArc->close();
            return false;
        }
        return $zipArc->close();
    }


    /**
     * 删除文件
     * @param $dirName
     * @param false $delSelf
     * @return bool
     */
    static function delFile($dirName, bool $delSelf=false): bool
    {
        if(file_exists($dirName) && $handle = opendir($dirName)){
            while(false !==($item = readdir( $handle))){
                if($item != '.' && $item != '..'){
                    if(file_exists($dirName.'/'.$item) && is_dir($dirName.'/'.$item)){
                        delFile($dirName.'/'.$item);
                    }else{
                        if(!unlink($dirName.'/'.$item)){
                            return false;
                        }
                    }
                }
            }
            closedir($handle);
            if($delSelf){
                if(!rmdir($dirName)){
                    return false;
                }
            }
        }else{
            return false;
        }
        return true;
    }


}