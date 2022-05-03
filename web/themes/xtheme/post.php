<?php get_header('post'); ?>
<link rel="stylesheet" type="text/css" href="<?php echo setting_template()."/static/css/post.css"?>">
<link rel="stylesheet" type="text/css" href="<?php echo setting_template()."/static/css/prism.css"?>">
<div id="main" class="show">
    <div id="bg" style="background-image: url(<?php setting_index_background(); ?>)"></div>
    <div id="main-body" class="mx-auto">
        <div class="col-xl-2 col-lg-3 col-md-3 col-sm-3 col-0 div-box-content left-box">
            <div id="sidebar-left" class="giligili-left">
                <?php xy_left_side(); ?>
            </div>
        </div>
        <div class="col-xl-5 col-lg-5 col-md-10 col-sm-2 col-11 div-box-content px-0">
            <div id="post-content-main" class="giligili-left">
                <div id="post-content">
                        <div class="giligili-item breadcrumbs" id="post-breadcrumb">
                            <el-breadcrumb separator="/">
                                <el-breadcrumb-item>
                                    <a href="/">
                                        首页
                                    </a>
                                </el-breadcrumb-item>
                                <el-breadcrumb-item>
                                    <a href="/?category=<?php echo get_post_category()[0]["link"]?>">
                                        <?php echo get_post_category()[0]["name"]?>
                                    </a>
                                </el-breadcrumb-item>
                                <el-breadcrumb-item><?php post_title(); ?></el-breadcrumb-item>
                            </el-breadcrumb>
                        </div>
                        <article id="post" class="giligili-item post">
                            <div class="post-info">
                                <div>
                                    <div class="thumbnail-link" style="background-image:url(<?php post_image(); ?>)"></div>
                                    <div class="thumbnail-shadow"></div>
                                    <canvas id="bubble" width="752" height="200"></canvas>
                                </div>
                                <div class="post-content">
                                    <div class="text-info">
                                        <header class="entry-header">
                                            <h1 class="entry-title text-center">
                                                <?php post_title(); ?>
                                            </h1>
                                            <div class="post-meta text-center">
                                                <span>
                                                    <i class="fa fa-calendar-alt"></i> <?php post_date(); ?>
                                                    <i class="fa fa-book-reader"></i> <?php post_view(); ?>次阅读
                                                    <!-- <font-awesome-icon icon="calculator" /> 共个字-->
                                                  <span class="hd">
                                                      <i class="fa fa-comments"></i><?php post_comment(); ?>条评论
                                                      <i class="fa fa-thumbs-up"></i><?php post_good(); ?>人点赞
                                                  </span>
                                                </span>
                                            </div>
                                        </header>
                                        <?php if (!get_post_encrypt()): ?>
                                        <div id="post-content" class="content">
                                            <?php post_content(); ?>
                                        </div>
                                        <?php else: ?>
                                            <div class="lock-face">
                                                <img class="face-img" alt src="<?php echo setting_template()."/static/image/post/face.jpg"?>">
                                                <div class="lock-text">
                                                    这是一篇受保护的文章，请输入访问密码!
                                                </div>
                                                <div class="lock-input">
                                                    <el-input v-model="password" type="password" placeholder="请输入访问密码">
                                                        <el-button slot="prepend" icon="el-icon-key" ></el-button>
                                                    </el-input>
                                                </div>
                                                <el-button class="lock-btn" type="primary" @click="getEncryptContent(password)">
                                                    点击访问
                                                </el-button>
                                            </div>
                                        <?php endif;?>
                                    </div>
                                </div>
                                <span class="category">
                                    <a href="/?category=<?php echo get_post_category()[0]["link"]?>">
                                        <?php echo get_post_category()[0]["name"]?>
                                    </a>
                                </span>
                            </div>
                            <footer id="post-functions" class="entry-footer clearfix">
                                <div>
                                    <div
                                        id="layer1"
                                        class="layer layer-page"
                                    >
                                        <div class="layer-title" style="">
                                            打赏作者
                                            <span class="layer-setwin">
                                                <a class="layer-close1" href="javascript:void(0);" @click="close"><i class="fa fa-times"></i> </a>
                                        </span>
                                        </div>
                                        <div id="" class="layer-content" style="height: 328px;">
                                            <div class="donate-box">
                                                <div class="meta-pay text-center">
                                                    <strong>扫一扫支付</strong>
                                                </div>
                                                <div class="qr-pay text-center">
                                                    <img id="alipay_qr" alt="" class="pay-img" :src="alipay">
                                                    <img id="wechat_qr" alt="" class="pay-img" :src="wechat">
                                                </div>
                                                <div class="choose-pay text-center mt-2">
                                                    <input id="alipay" type="radio" name="pay-method" checked="" @click="switchpay('aliy')">
                                                    <label for="alipay" class="pay-button">
                                                        <img alt src="<?php echo setting_template().'/static/image/post/alipay.png'?>">
                                                    </label>
                                                    <input id="wechatpay" type="radio" name="pay-method" @click="switchpay('wechat')">
                                                    <label for="wechatpay" class="pay-button">
                                                        <img alt src="<?php echo setting_template().'/static/image/post/wechat.png'?>">
                                                    </label>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div id="post-like-donate" class="post-like-donate text-center clearfix">
                                        <a href="javascript:;" class="Donate" @click="open"><i class="fa fa-donate"></i> 打赏</a>
                                        <a id="btn" href="javascript:;" data-action="love" class="Love" @click="love"
                                        > <i class="fa fa-thumbs-up"></i> 点赞</a>
                                        <a href="javascript:;" class="Share" @click="sharebar"><i class="fa fa-share-alt"></i> 分享</a>
                                        <div class="share-wrap" style="display: none">
                                            <div class="share-group">
                                                <a href="javascript:;" class="share-plain twitter" rel="nofollow" @click="share('qq')">
                                                    <div class="icon-wrap">
                                                        <i class="fab fa-qq"></i>
                                                    </div>
                                                </a>
                                                <a href="javascript:;" class="share-plain qzone" rel="nofollow" @click="share('qzone')">
                                                    <div class="icon-wrap">
                                                        <i class="fa fa-star"></i>
                                                    </div>
                                                </a>
                                                <a href="javascript:;" class="share-plain weixin" rel="nofollow" @click="share('weixin')">
                                                    <div class="icon-wrap">
                                                        <i class="fab fa-weixin"></i>
                                                    </div>
                                                </a>
                                                <a href="javascript:;" class="share-plain weibo" rel="nofollow" @click="share('weibo')">
                                                    <div class="icon-wrap">
                                                        <i class="fab fa-weibo"></i>
                                                    </div>
                                                </a>
                                            </div>
                                        </div>
                                    </div>
                                </div>
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
                        </article>
                    </div>
                <!--评论部分-->
                <?php xy_comments(get_post_id()); ?>
            </div>
        </div>
        <div class="col-xl-2 col-lg-3 col-md-3 col-sm-3 col-0 div-box-content right-box">
            <div id="sidebar-right" class="giligili-left sidebar">
                <?php xy_right_side(); ?>
            </div>
        </div>
    </div>
</div>
<script>
    let globalAlipay = '<?php post_alipay(); ?>';
    let globalWechat = '<?php post_wechat(); ?>';
    let globalPostId = '<?php post_id(); ?>';
    let globalAvatar = '<?php echo setting_template()."/static/image/post/avatar.png"?>';
    // 自动加上行号显示
    $('pre').addClass("line-numbers").css("white-space", "pre-wrap");
</script>
<script type="text/javascript" src="<?php echo setting_template()."/static/js/prism.js"?>"></script>
<?php get_footer(); ?>