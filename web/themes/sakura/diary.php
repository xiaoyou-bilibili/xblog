<?php get_header('diary');?>
    <?php lib_script_vue();lib_css_element();lib_script_element(); ?>
    <link rel="stylesheet" type="text/css" href="<?php echo setting_template()."/static/css/post.css"?>">
    <div class="main-part">
        <?php $show = xy_option("theme_sakura_show_side");if($show):?>
            <div class="col-xl-2 col-lg-3 col-md-10 col-sm-3 col-0 div-box-content">
                <div id="sidebar-left" class="giligili-left">
                    <?php xy_left_side(); ?>
                </div>
            </div>
        <?php endif;?>
        <div class="col-xl-5 col-lg-5 col-md-10 col-sm-2 col-11">
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
        <?php if($show):?>
            <div class="col-xl-2 col-lg-3 col-md-3 col-sm-3 col-0 div-box-content right-box">
                <div id="sidebar-right" class="giligili-left sidebar">
                    <?php xy_right_side(); ?>
                </div>
            </div>
        <?php endif;?>
    </div>
    <script>
        /*设置当前页数和总页数*/
        let globalTotal =  <?php current_total(); ?>;
        let globalCurrent =  <?php current_current(); ?>;
        if($("#diary-list").length>0){new Vue({el:'#diary-list',data:{total:globalTotal,current:globalCurrent,password:''},methods:{getDiary(page){window.location.href=xy.tools.updateQueryStringParameter(window.location.href,'page',page)},getEncryptContent(id){window.location='/archives/'+id+'?password='+this.password},gotoDiary(id){window.location='/archives/'+id}}})}
    </script>
<?php get_footer(); ?>