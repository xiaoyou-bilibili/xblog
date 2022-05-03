<style>
    *{margin:0;padding:0;-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box;font-family:Microsoft Yahei,"微软雅黑","Helvetica Neue",Helvetica,Hiragino Sans GB,WenQuanYi Micro Hei,sans-serif}
    html,body{width:100%;height:100%}
    a{text-decoration:none;color:#333;-webkit-transition:.3s ease all;-moz-transition:.3s ease all;-o-transition:.3s ease all;transition:.3s ease all}
    a:focus{outline:none}
    .sitemap-lists a{padding:8px 5px;border-radius:5px}
    .sitemap-lists a:hover{background:#eee}
    img{border:none}
    li{list-style:none}
    .clear-fix{zoom:1}
    .clear-fix:before,.clear-fix:after{display:table;line-height:0;content:""}
    .clear-fix:after{clear:both}
    .hidden{display:none}
    .container{max-width:900px;margin:0 auto;position:relative;padding:5px}
    .page-title{font-weight:600;font-size:30px;text-align:center;padding:40px;position:relative}
    .page-title:after{content:"";border-bottom:3px #bdbdbd solid;position:absolute;left:50%;top:50%;padding-top:60px;transform:translate(-50%,-50%);width:60px;z-index:-1}
    .page-title:hover>a{color:#848484}
    .section-title{font-weight:500;font-size:16px;position:relative;margin:15px 0 10px;color:#fff;background:#565555;display:inline-block;padding:5px 8px;border-radius:5px}
    .post-lists li{padding:4px 0}
    .post-lists li>a{display:block}
    .category-lists li>a,.tag-lists li>a{display:inline-block;float:left;margin-right:4px;margin-bottom:4px}
    .page-footer{text-align:center;padding:10px;font-size:14px;color:#c7c7c7}
</style>
<div class="container">
    <h1 class="page-title">
        <a href="/assets/sitemap.xml" target="_blank">Sitemap</a>
    </h1>
    <h2 class="section-title">
        文章 / Article
    </h2>
    <ul class="sitemap-lists post-lists clear-fix">
        <?php foreach (get_tools_sitemap_post() as $item): ?>
        <li>
            <?php echo "<a href='".$item["url"]."' title='".$item["title"]."' target='_blank'>".$item["title"]."</a>" ?>
        </li>
        <?php endforeach;?>
    </ul>
    <h2 class="section-title">
        文档 / Document
    </h2>
    <ul class="sitemap-lists post-lists clear-fix">
        <?php foreach (get_tools_sitemap_doc() as $item): ?>
            <li>
                <?php echo "<a href='".$item["url"]."' title='".$item["title"]."' target='_blank'>".$item["title"]."</a>" ?>
            </li>
        <?php endforeach;?>
    </ul>
    <h2 class="section-title">
        分类 / Category
    </h2>
    <ul class="sitemap-lists category-lists clear-fix">
        <?php foreach (get_tools_sitemap_category() as $item): ?>
            <li>
                <?php echo "<a href='".$item["url"]."' title='".$item["title"]."' target='_blank'>".$item["title"]."</a>" ?>
            </li>
        <?php endforeach;?>
    </ul>
    <h2 class="section-title">
        标签 / Tag
    </h2>
    <ul class="sitemap-lists tag-lists clear-fix">
        <?php foreach (get_tools_sitemap_tag() as $item): ?>
            <li>
                <?php echo "<a href='".$item["url"]."' title='".$item["title"]."' target='_blank'>".$item["title"]."</a>" ?>
            </li>
        <?php endforeach;?>
    </ul>
</div>

