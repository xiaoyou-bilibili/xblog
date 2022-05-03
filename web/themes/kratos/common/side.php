<?php if (get_widget_title()==""): ?>
<aside id="kratos_about-2" class="widget widget_kratos_about clearfix">
    <?php widget_html(); ?>
</aside>
<?php else: ?>
<aside id="kratos_comments-2" class="widget widget_kratos_comments clearfix">
    <h4 class="widget-title"><?php widget_title(); ?></h4>
    <div class="recentcomments">
        <?php widget_html(); ?>
    </div>
</aside>
<?php endif; ?>