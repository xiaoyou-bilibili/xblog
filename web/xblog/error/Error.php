<?php

namespace xblog\error;
use xblog\ext\Load;

/**
 * Class Error
 * @description 错误处理板块
 * @author 小游
 * @date 2020/4/30
 * @package xblog\error
 */
class Error
{
    /**
     *  404错误
     */
    public static function Code404()
    {
        Load::showError(404);
    }

    public static function Code500(){
        Load::showError(500);
    }
}