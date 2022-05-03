<?php get_header('doc');lib_script_vue();lib_script_element();lib_css_element(); ?>
<style>
    footer{display: none}
</style>
<div id="doc-app" style="background: white;" class="window-container">
    <div class="window-body with-sidebar">
        <div id="sidebar" class="sidebar">
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
        <div class="workspace" id="workspace">
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
<script>
    // 文档功能
    new Vue({el:'#doc-app',data:{docList:[],filterText:'',key:[],check:[],docContent:{title:'',content:''}},watch:{filterText(val){this.$refs.tree.filter(val)}},mounted(){this.getDocList()},methods:{filterNode(value,data){if(!value){return true}else{return data.label.includes(value)}},sideBar(){if($('.sidebar').css('margin-left')==='0px'){$('.sidebar').css('margin-left','-280px')}else{$('.sidebar').css('margin-left','0')}},click(choose){window.history.pushState({id:1,name:'doc'},'测试','/doc/'+choose.id);this.setDocContent(choose.id)},setDocContent(id){xy.net.request('plugins/docs/'+id,"GET").then((data)=>{window.scroll(0,0);this.docContent=data})},getDocList(){xy.net.request("plugins/docs","GET").then((data)=>{this.docList=this.DataProcess(data)})},DataProcess(data){let trees=[];if(data!=null){for(let i=0;i<data.length;i++){if(data[i].parent===0){trees.push({id:data[i].id,label:data[i].title})}}trees=this.InsertChild(trees,data)}return trees},InsertChild(trees,data){for(let i=0;i<trees.length;i++){for(let j=0;j<data.length;j++){if(trees[i].id===data[j].parent){if(Object.prototype.hasOwnProperty.call(trees[i],'children')){trees[i].children.push({id:data[j].id,label:data[j].title})}else{trees[i].children=[{id:data[j].id,label:data[j].title}]}}}if(Object.prototype.hasOwnProperty.call(trees[i],'children')){trees[i].children=this.InsertChild(trees[i].children,data)}}return trees}}});
    const high=$(document).height()-$('#navigation').height();
    $('#doc-app').css('height', high + 'px')
    $('#sidebar').css('height', high + 'px')
    $('#workspace').css('height', high + 'px')
</script>
<?php get_footer(); ?>
