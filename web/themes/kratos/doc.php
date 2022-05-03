<?php get_header('doc');lib_css_element(); ?>
<div id="app-doc" style="background: white;" class="window-container">
    <div class="window-body with-sidebar">
        <div class="sidebar">
            <div class="sidebar-header">
                <a href="javascript:void(0);" data-no-pjax="true" class="title">个人文档系统</a>
                <div class="search-form">
                    <div class="ui small fluid icon input">
                        <input v-model="filterText" type="text" placeholder="请输入搜索关键词..."><i class="fa fa-search"></i>
                    </div>
                </div>
            </div>
            <div class="sidebar-body">
                <div class="catalog-body">
                    <el-tree
                        id="tree"
                        ref="tree"
                        :data="docList"
                        :default-expanded-keys="key"
                        :default-checked-keys="check"
                        :highlight-current="true"
                        :expand-on-click-node="false"
                        :filter-node-method="filterNode"
                        @node-click="click"
                    />
                </div>
            </div>
        </div>
        <div class="workspace">
            <div id="post" class="article">
                <div class="article-head">
                    <div class="left tools">
                        <a class="item icon" @click="sideBar">
                            <i class="fa fa-th-list"></i>
                        </a>
                    </div>
                    <h1 v-if="docContent.title!==''">{{docContent.title}}</h1>
                    <h1 v-else><?php plugins_doc_title(); ?></h1>
                </div>
                <div class="content-div">
                    <div v-if="docContent.content!==''" class="article-body kancloud-markdown-body" v-html="docContent.content"></div>
                    <div v-else class="article-body kancloud-markdown-body">
                        <?php plugins_doc_content(); ?>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<?php lib_script_vue();lib_script_element(); ?>
<?php get_footer();?>
