<?php if (get_widget_title()==""): ?>
    <div class="notice-panel">
        <div class="textwidget">
            <?php widget_html(); ?>
        </div>
    </div>
<?php else: ?>
    <div class="notice-panel tools-item" >
        <h4 class="widget-title">
            <?php widget_title(); ?>
        </h4>
        <div class="textwidget side-padding">
            <?php widget_html(); ?>
        </div>
    </div>
<?php endif; ?>