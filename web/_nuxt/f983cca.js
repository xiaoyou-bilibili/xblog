(window.webpackJsonp=window.webpackJsonp||[]).push([[8,34],{1300:function(e,t,n){"use strict";n(936)},1301:function(e,t,n){var r=n(8)(!1);r.push([e.i,"img.project-img[data-v-11fc4368]{width:100%;height:100px;border-radius:5px}.vxe-button+.vxe-button[data-v-11fc4368]{margin-left:0!important}",""]),e.exports=r},1330:function(e,t,n){"use strict";n.r(t);n(67),n(119),n(95);var r=n(651),o=n(85),c=[{label:"是",value:!0},{label:"否",value:!1}],l={layout:"admin",mixins:[r.default],data:function(){var e=this;return{selects:"",tableProxy:{ajax:{query:function(e){var t=e.page,form=e.form;return Object(o.d)({url:"/projects",type:"get",data:{page:t.currentPage,page_size:t.pageSize,search_type:form.search_type,search_key:form.search_key}})},delete:function(t){return e.deleteProject(e.selects)}}},tableForm:{items:[{field:"search_type",title:"搜索类型",itemRender:{name:"$select",options:[{label:"项目名字",value:"name"},{label:"项目介绍",value:"description"},{label:"是否置顶",value:"is_top"}]}},{field:"search_key",visibleMethod:function(e){return"is_top"!==e.data.search_type},title:"关键词",itemRender:{name:"$input",attrs:{placeholder:"请输入关键词"}}},{field:"search_key",visibleMethod:function(e){return"is_top"===e.data.search_type},title:"置顶",itemRender:{name:"$select",options:c}},{itemRender:{name:"$button",props:{content:"查询",type:"submit",status:"primary"}}},{itemRender:{name:"$button",props:{content:"重置",type:"reset"}}}]},tableToolbar:{buttons:[{code:"insert_actived",name:"添加",icon:"fa fa-plus",status:"primary"},{code:"delete",name:"删除选中",icon:"fa fa-trash",status:"danger"}],zoom:!0,custom:!0},editConfig:{trigger:"manual",mode:"row",autoClear:!1,icon:"none"},tableColumn:[{type:"checkbox",width:50},{field:"name",title:"项目名称",editRender:{name:"input"}},{field:"img",width:180,title:"项目图片",slots:{default:"img"},editRender:{name:"input"}},{field:"make_time",title:"制作时间",editRender:{name:"$input",props:{type:"date"}}},{field:"description",title:"项目描述",editRender:{name:"textarea"}},{field:"video_url",title:"视频地址",editRender:{name:"input"}},{field:"blog_url",title:"博客地址",editRender:{name:"input"}},{field:"code_url",title:"代码地址",editRender:{name:"input"}},{field:"link",title:"轮播图地址",editRender:{name:"input"}},{field:"is_top",title:"轮播图置顶",formatter:"formatTop",width:100,editRender:{name:"$select",options:c}},{title:"操作",slots:{default:"option"}}]}},methods:{checkboxChangeEvent:function(data){var e=[];data.records.map((function(t){return e.push(t.id)})),this.selects=e.toString()},editRowEvent:function(e){this.$refs.grid.setActiveRow(e)},saveRowEvent:function(e){var t=this;this.$refs.grid.clearActived().then((function(){e._id.includes("row")?t.$store.dispatch("admin-plugins/pluginRequest",{url:"/projects",type:"post",data:e}).then((function(e){return t.updateData()})).catch((function(e){t.$message.error(e),t.updateData(!1)})):t.$store.dispatch("admin-plugins/pluginRequest",{url:"/projects/".concat(e._id),type:"put",data:e}).then((function(e){return t.updateData()})).catch((function(e){return t.$message.error(e)}))}))},updateData:function(){var e=!(arguments.length>0&&void 0!==arguments[0])||arguments[0];e&&this.$message.success("更新成功"),this.$refs.grid.commitProxy("query")},deleteProject:function(e){var t=this;""===e?this.$message.warning("请选择内容！"):this.$store.dispatch("admin-plugins/pluginRequest",{url:"/projects/".concat(e),type:"delete",data:null}).then((function(e){return t.updateData()})).catch((function(e){return t.$message.error(e)}))}}},d=(n(1300),n(5)),component=Object(d.a)(l,(function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("vxe-grid",{ref:"grid",attrs:{border:"",resizable:"",height:"780","row-id":"id","pager-config":{pageSize:10},"proxy-config":e.tableProxy,"form-config":e.tableForm,"toolbar-config":e.tableToolbar,columns:e.tableColumn,"edit-config":e.editConfig},on:{"checkbox-change":e.checkboxChangeEvent,"checkbox-all":e.checkboxChangeEvent},scopedSlots:e._u([{key:"img",fn:function(e){var t=e.row;return[n("img",{staticClass:"project-img",attrs:{src:t.img}})]}},{key:"option",fn:function(t){var r=t.row;return[e.$refs.grid.isActiveByRow(r)?n("vxe-button",{attrs:{icon:"fa fa-save",status:"primary",title:"保存",circle:""},on:{click:function(t){return e.saveRowEvent(r)}}}):n("vxe-button",{attrs:{icon:"fa fa-edit",title:"编辑",circle:""},on:{click:function(t){return e.editRowEvent(r)}}}),e._v(" "),n("el-popconfirm",{attrs:{title:"确定要删除这个项目？"},on:{onConfirm:function(t){return e.deleteProject(r._id)}}},[n("vxe-button",{attrs:{slot:"reference",icon:"fa fa-trash",title:"删除",circle:""},slot:"reference"})],1)]}}])})],1)}),[],!1,null,"11fc4368",null);t.default=component.exports},651:function(e,t,n){"use strict";n.r(t),t.default={head:function(){return{title:"XBlog后台管理系统"}}}},936:function(e,t,n){var content=n(1301);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[e.i,content,""]]),content.locals&&(e.exports=content.locals);(0,n(9).default)("60b2a3c7",content,!0,{sourceMap:!1})}}]);