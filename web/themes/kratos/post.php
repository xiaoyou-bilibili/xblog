<?php get_header('post'); ?>
<div id="kratos-blog-post">
    <div id="container" class="container">
        <!-- 引入js文件-->
        <script src="<?php echo setting_template()."/static/js/prism.js"?>"></script>
        <div class="row">
            <section id="main" class="col-md-8">
                <article>
                    <div class="kratos-hentry kratos-post-inner clearfix">
                        <header class="kratos-entry-header">
                            <div class="wow shake animated" style="visibility: visible; animation-name: shake;">
                                <h1 class="kratos-entry-title text-center"><?php post_title(); ?></h1>
                            </div>
                            <div class="wow lightSpeedIn animated" style="visibility: visible; animation-name: lightSpeedIn;">
                                <div class="kratos-post-meta text-center">
                                    <span>
                                    <i class="fa fa-calendar"></i>  <?php post_date(); ?>
                                        <i class="fa fa-eye"></i> <?php post_view(); ?> 次阅读
                                        <span class="hd">
                                            <i class="fa fa-commenting-o"></i> <?php post_comment(); ?> 条评论
                                            <i class="fa fa-thumbs-o-up"></i> <?php post_good(); ?> 人点赞
                                        </span>
                                    </span>
                                </div>
                            </div>
                        </header>
                        <div class="kratos-post-content">
                            <?php if (!get_post_encrypt()){
                                post_content();
                            } else {
                                $template = setting_template();
                                echo <<<EOF
                                <form class="protected-post-form" action="" method="get">
                                    <div class="panel panel-pwd">
                                        <div class="panel-body text-center">
                                            <img class="post-pwd" src="${template}/static/images/fingerprint.png"><br />
                                            <h4>这是一篇受保护的文章，请输入阅读密码！</h4>
                                            <div class="input-group">
                                                <div class="input-group-addon"><i class="fa fa-key"></i></div>
                                                <p><input class="form-control" placeholder="输入阅读密码" name="password" type="password" size="20"></p>
                                            </div>
                                            <div class="comment-form" style="margin-top:15px;">
                                            <button id="generate" class="btn btn-primary btn-pwd" type="submit">确认</button></div>
                                        </div>
                                    </div>
                                </form>
EOF;
                            } ?>
                        </div>
                        <div class="kratos-copyright text-center clearfix">
                            <h5>本作品采用 <a rel="license nofollow" target="_blank" href="http://creativecommons.org/licenses/by-sa/4.0/">知识共享署名-相同方式共享 4.0 国际许可协议</a> 进行许可</h5>
                        </div>
                        <footer class="kratos-entry-footer clearfix">
                            <div class="post-like-donate text-center clearfix" id="post-like-donate">
                                <a href="javascript:;" class="Donate"><i class="fab fa-bitcoin"></i> 打赏</a>
                                <a href="javascript:;" id="btn-love" data-id="1" class="Love done"><i class="fa fa-thumbs-up"></i> 点赞</a>
                                <a href="javascript:;" class="Share"><i class="fa fa-share-alt"></i>分享</a><div class="share-wrap" style="display: none;">
                                    <div class="share-group">
                                        <a href="javascript:;" class="share-plain twitter" onclick="share('qq');" rel="nofollow">
                                            <div class="icon-wrap">
                                                <i class="fab fa-qq"></i>
                                            </div>
                                        </a>
                                        <a href="javascript:;" class="share-plain weixin" onclick="share('qzone');" rel="nofollow">
                                            <div class="icon-wrap">
                                                <i class="fa fa-star"></i>
                                            </div>
                                        </a>
                                        <a href="javascript:;" class="share-plain weibo" onclick="share('weibo');" rel="nofollow">
                                            <div class="icon-wrap">
                                                <i class="fab fa-weibo"></i>
                                            </div>
                                        </a>
                                        <a href="javascript:;" class="share-plain facebook style-plain" onclick="share('facebook');" rel="nofollow">
                                            <div class="icon-wrap">
                                                <i class="fab fa-facebook"></i>
                                            </div>
                                        </a>
                                        <a href="javascript:;" class="share-plain twitter style-plain" onclick="share('twitter');" rel="nofollow">
                                            <div class="icon-wrap">
                                                <i class="fab fa-twitter"></i>
                                            </div>
                                        </a>
                                    </div>
                                    <script type="text/javascript">
                                        function share(obj){
                                            var qqShareURL="http://connect.qq.com/widget/shareqq/index.html?";
                                            var weiboShareURL="http://service.weibo.com/share/share.php?";
                                            var qzoneShareURL="https://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?";
                                            var facebookShareURL="https://www.facebook.com/sharer/sharer.php?";
                                            var twitterShareURL="https://twitter.com/intent/tweet?";
                                            var host_url="http://127.0.0.1/index.php/2021/04/29/hello-world/";
                                            var title='%E3%80%90%E4%B8%96%E7%95%8C%EF%BC%8C%E6%82%A8%E5%A5%BD%EF%BC%81%E3%80%91';
                                            var qqtitle='%E3%80%90%E4%B8%96%E7%95%8C%EF%BC%8C%E6%82%A8%E5%A5%BD%EF%BC%81%E3%80%91';
                                            var excerpt='%E6%AC%A2%E8%BF%8E%E4%BD%BF%E7%94%A8WordPress%E3%80%82%E8%BF%99%E6%98%AF%E6%82%A8%E7%9A%84%E7%AC%AC%E4%B8%80%E7%AF%87%E6%96%87%E7%AB%A0%E3%80%82%E7%BC%96%E8%BE%91%E6%88%96%E5%88%A0%E9%99%A4%E5%AE%83%EF%BC%8C%E7%84%B6%E5%90%8E%E5%BC%80%E5%A7%8B%E5%86%99%E4%BD%9C%E5%90%A7%EF%BC%81';
                                            var wbexcerpt='%E6%AC%A2%E8%BF%8E%E4%BD%BF%E7%94%A8WordPress%E3%80%82%E8%BF%99%E6%98%AF%E6%82%A8%E7%9A%84%E7%AC%AC%E4%B8%80%E7%AF%87%E6%96%87%E7%AB%A0%E3%80%82%E7%BC%96%E8%BE%91%E6%88%96%E5%88%A0%E9%99%A4%E5%AE%83%EF%BC%8C%E7%84%B6%E5%90%8E%E5%BC%80%E5%A7%8B%E5%86%99%E4%BD%9C%E5%90%A7%EF%BC%81';
                                            var pic="";
                                            var _URL;
                                            if(obj=="qq"){
                                                _URL=qqShareURL+"url="+host_url+"&title="+qqtitle+"&pics="+pic+"&desc=&summary="+excerpt+"&site=vtrois";
                                            }else if(obj=="weibo"){
                                                _URL=weiboShareURL+"url="+host_url+"&title="+title+wbexcerpt+"&pic="+pic;
                                            }else if(obj=="qzone"){
                                                _URL=qzoneShareURL+"url="+host_url+"&title="+qqtitle+"&pics="+pic+"&desc=&summary="+excerpt+"&site=vtrois";
                                            }else if(obj=="facebook"){
                                                _URL=facebookShareURL+"u="+host_url;
                                            }else if(obj=="twitter"){
                                                _URL=twitterShareURL+"text="+title+excerpt+"&url="+host_url;
                                            }
                                            window.open(_URL);
                                        }
                                    </script>
                                </div>                </div>
                            <div class="footer-tag clearfix">
                                <div class="pull-left">
                                    <i class="fa fa-tags"></i>
                                    <?php foreach (get_post_tags() as $tag){
                                        echo "<a href='/?tag=${tag["link"]}'>${tag["name"]}</a>";
                                    }?>
                                </div>
                                <div class="pull-date">
                                    <span>最后编辑：<?php post_modify(); ?></span>
                                </div>
                            </div>
                        </footer>
                    </div>
                    <nav class="navigation post-navigation clearfix" role="navigation">
                    </nav>
                    <!--  评论区 -->
                    <?php xy_comments(get_post_id()); ?>
                </article>
            </section>
            <aside id="kratos-widget-area" class="col-md-4 hidden-xs hidden-sm scrollspy">
                <div id="sidebar" class="affix-top">
                    <?php xy_left_side(); ?>
                </div>
            </aside>
        </div>
    </div>
</div>
<?php lib_script_vue(); ?>
<script>
    let globalPost={alipay:'<?php post_alipay(); ?>',wechat:'<?php post_wechat(); ?>','id':<?php post_id(); ?>,'isLogin': <?php tools_print_bool(user_is_login()); ?>}
    let globalUserInfo = '<?php echo json_encode($GLOBALS["userInfo"]) ?>';
</script>
<?php get_footer(); ?>