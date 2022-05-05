<?php get_header('diary');?>
    <link rel="stylesheet" type="text/css" href="<?php echo setting_template()."/static/css/post.css"?>">
    <div id="main" class="show">
        <div id="bg" style="background-image: url(<?php setting_index_background(); ?>)" ></div>
        <div class="mx-auto" id="main-body">
            <div class="col-xl-2 col-lg-3 col-md-3 col-sm-3 col-0 div-box-content">
                <div id="sidebar-left" class="giligili-left">
                    <?php xy_left_side(); ?>
                </div>
            </div>
            <div class="col-xl-5 col-lg-5 col-md-5 col-sm-2 col-11 div-box-content">
                <div id="diary-list">
                    <div id="diary">
                        <?php xy_diary(); ?>
                    </div>
                    <el-pagination
                            background
                            layout="prev, pager, next"
                            :total="total"
                            :current-page="current"
                            :page-size="1"
                            @current-change="getDiary"
                    ></el-pagination>
                </div>
            </div>
            <div class="col-xl-2 col-lg-3 col-md-3 col-sm-3 col-0 div-box-content ">
                <div id="sidebar-right" class="giligili-left sidebar">
                    <?php xy_right_side(); ?>
                </div>
            </div>
        </div>
    </div>
    <script>
        /*设置当前页数和总页数*/
        let globalTotal =  <?php current_total(); ?>;
        let globalCurrent =  <?php current_current(); ?>;
    </script>
<?php get_footer(); ?>