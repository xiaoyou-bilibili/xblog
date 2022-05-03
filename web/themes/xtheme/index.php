<?php get_header('index');?>
<div id="main" class="show">
    <!--设置背景-->
    <div id="bg" style="background-image: url(<?php setting_index_background(); ?>)" ></div>
    <div class="mx-auto" id="main-body">
        <div class="col-xl-2 col-lg-3 col-md-10 col-sm-3 col-0 div-box-content">
            <div id="sidebar-left" class="giligili-left">
                <?php xy_left_side(); ?>
            </div>
        </div>
        <div class="col-xl-5 col-lg-5 col-md-10 col-sm-2 col-11 div-box-content px-0">
            <div id="post-list" class="giligili-left">
                <div class="b-wrap">
                    <div class="primary-menu-itnl">
                        <div id="primaryPageTab" class="page-tab report-wrap-module">
                            <ul class="con">
                                <li>
                                    <a href="/">
                                        <div class="navigation-home round home">
                                            <i class="fa fa-home"></i>
                                        </div>
                                        <span>首页</span>
                                    </a>
                                </li>
                            </ul>
                        </div>
                        <span class="tab-line-itnl" ></span>
                        <div id="post-list-category" class="channel-menu-itnl report-wrap-module">
                            <?php $category = get_index_category();foreach ($category["parent"] as $parent): ?>
                            <div class="item">
                                <el-popover
                                        placement="top-start"
                                        width="100"
                                        trigger="hover"
                                >
                                    <?php foreach ($category["child"] as $child){
                                        if($parent["id"]==$child["parent"]){
                                            echo  "<a href=\"/?category=${child["link"]}\">${child["name"]}</a>";
                                        }
                                    }
                                    echo  "<a slot=\"reference\" class=\"name\" href=\"javascript:void(0);\"><span>${parent["name"]}<em>${parent["count"]}</em></span></a>"?>
                                </el-popover>
                            </div>
                            <?php endforeach; ?>
                        </div>
                    </div>
                </div>
                <div id="blog-post" class="blog-post">
                    <?php xy_posts(); ?>
                    <div id="post-list-pagination">
                        <el-pagination
                                background
                                layout="prev, pager, next"
                                :page-size="1"
                                :total="total"
                                :current-page="current"
                                @current-change="getPosts"
                        ></el-pagination>
                    </div>
                </div>
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
/*设置当前页数和总页数*/
let globalTotal =  <?php current_total(); ?>;
let globalCurrent =  <?php current_current(); ?>;
</script>
<?php get_footer(); ?>