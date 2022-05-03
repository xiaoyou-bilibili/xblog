<?php get_header('friend'); ?>
<div class="main-part">
    <div id="content" class="site-content" style="background-color: rgba(255, 255, 255, 0.8);">
        <span class="linkss-title">友人帐</span>
        <article class="post-item post-26 page type-page status-publish hentry">
            <div class="links">
                <ul class="link-items fontSmooth">
                    <?php $friends = more_get_friends(); foreach ($friends as $item ):
                        echo <<<EOF
                <li class="link-item">
                    <a class="link-item-inner effect-apollo" href="${item['url']}" title="${item['name']}" target="_blank" rel="friend">
                        <img class="lazyload" onerror="imgError(this,1)" data-src="" src="${item['avatar']}">
                        <span class="sitename">${item['name']}</span>
                        <div class="linkdes">${item['dec']}</div>
                    </a>
                </li>
EOF;endforeach;?>
                </ul>
            </div>
        </article>
        <div class="have-toc"></div>
        <div class="toc-container">
            <div class="toc">
                <ol class="toc-list "></ol>
            </div>
        </div>
    </div>
</div>
<?php get_footer(); ?>