import{_ as V}from"./index-35ac8c1e.js";import{E as B,a as E}from"./element-plus-f169eafc.js";import{H as b,ay as $,e as A,_ as O,i as P,ah as r,ar as H,o as d,c as M,a as f,V as t,P as n,T as i,Q as U,O as u,u as F}from"./@vue-021a4eb1.js";import{_ as I}from"./_plugin-vue_export-helper-c27b6911.js";import"./pinia-a8c871c6.js";import"./vue-demi-71ba0ef2.js";import"./vue-router-5dc1aa30.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-4a2045cc.js";import"./nprogress-5ada7e0c.js";import"./call-bind-1df9733d.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./axios-707ed124.js";import"./qs-a8db372b.js";import"./side-channel-9112baf6.js";import"./object-inspect-d11bccf2.js";import"./mitt-f7ef348c.js";import"./lodash-es-36eb724a.js";import"./@vueuse-8a759679.js";import"./@popperjs-c75af06c.js";import"./@ctrl-f8748455.js";import"./dayjs-1798568c.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";const L={class:"system-role-container layout-padding"},j={class:"system-role-padding layout-padding-auto layout-padding-view"},Q={class:"system-user-search mb15"},q=b({name:"systemRole"}),G=b({...q,setup(J){const v=$(()=>V(()=>import("./dialog-bc03b251.js"),["./dialog-bc03b251.js","./@vue-021a4eb1.js","./_plugin-vue_export-helper-c27b6911.js","..\\css\\dialog-c40da9d0.css"],import.meta.url)),c=A(),a=O({tableData:{data:[],total:0,loading:!1,param:{search:"",pageNum:1,pageSize:10}}}),p=()=>{a.tableData.loading=!0;const l=[];for(let e=0;e<20;e++)l.push({roleName:e===0?"超级管理员":"普通用户",roleSign:e===0?"admin":"common",describe:`测试角色${e+1}`,sort:e,status:!0,createTime:new Date().toLocaleString()});a.tableData.data=l,a.tableData.total=a.tableData.data.length,setTimeout(()=>{a.tableData.loading=!1},500)},y=l=>{c.value.openDialog(l)},D=(l,e)=>{c.value.openDialog(l,e)},w=l=>{B.confirm(`此操作将永久删除角色名称：“${l.roleName}”，是否继续?`,"提示",{confirmButtonText:"确认",cancelButtonText:"取消",type:"warning"}).then(()=>{p(),E.success("删除成功")}).catch(()=>{})},h=l=>{a.tableData.param.pageSize=l,p()},x=l=>{a.tableData.param.pageNum=l,p()};return P(()=>{p()}),(l,e)=>{const z=r("el-input"),C=r("ele-Search"),_=r("el-icon"),m=r("el-button"),k=r("ele-FolderAdd"),s=r("el-table-column"),g=r("el-tag"),N=r("el-table"),S=r("el-pagination"),R=H("loading");return d(),M("div",L,[f("div",j,[f("div",Q,[t(z,{modelValue:a.tableData.param.search,"onUpdate:modelValue":e[0]||(e[0]=o=>a.tableData.param.search=o),size:"default",placeholder:"请输入角色名称",style:{"max-width":"180px"}},null,8,["modelValue"]),t(m,{size:"default",type:"primary",class:"ml10"},{default:n(()=>[t(_,null,{default:n(()=>[t(C)]),_:1}),i(" 查询 ")]),_:1}),t(m,{size:"default",type:"success",class:"ml10",onClick:e[1]||(e[1]=o=>y("add"))},{default:n(()=>[t(_,null,{default:n(()=>[t(k)]),_:1}),i(" 新增角色 ")]),_:1})]),U((d(),u(N,{data:a.tableData.data,style:{width:"100%"}},{default:n(()=>[t(s,{type:"index",label:"序号",width:"60"}),t(s,{prop:"roleName",label:"角色名称","show-overflow-tooltip":""}),t(s,{prop:"roleSign",label:"角色标识","show-overflow-tooltip":""}),t(s,{prop:"sort",label:"排序","show-overflow-tooltip":""}),t(s,{prop:"status",label:"角色状态","show-overflow-tooltip":""},{default:n(o=>[o.row.status?(d(),u(g,{key:0,type:"success"},{default:n(()=>[i("启用")]),_:1})):(d(),u(g,{key:1,type:"info"},{default:n(()=>[i("禁用")]),_:1}))]),_:1}),t(s,{prop:"describe",label:"角色描述","show-overflow-tooltip":""}),t(s,{prop:"createTime",label:"创建时间","show-overflow-tooltip":""}),t(s,{label:"操作",width:"100"},{default:n(o=>[t(m,{disabled:o.row.roleName==="超级管理员",size:"small",text:"",type:"primary",onClick:T=>D("edit",o.row)},{default:n(()=>[i("修改")]),_:2},1032,["disabled","onClick"]),t(m,{disabled:o.row.roleName==="超级管理员",size:"small",text:"",type:"primary",onClick:T=>w(o.row)},{default:n(()=>[i("删除")]),_:2},1032,["disabled","onClick"])]),_:1})]),_:1},8,["data"])),[[R,a.tableData.loading]]),t(S,{onSizeChange:h,onCurrentChange:x,class:"mt15","pager-count":5,"page-sizes":[10,20,30],"current-page":a.tableData.param.pageNum,"onUpdate:currentPage":e[2]||(e[2]=o=>a.tableData.param.pageNum=o),background:"","page-size":a.tableData.param.pageSize,"onUpdate:pageSize":e[3]||(e[3]=o=>a.tableData.param.pageSize=o),layout:"total, sizes, prev, pager, next, jumper",total:a.tableData.total},null,8,["current-page","page-size","total"])]),t(F(v),{ref_key:"roleDialogRef",ref:c,onRefresh:e[4]||(e[4]=o=>p())},null,512)])}}});const Ce=I(G,[["__scopeId","data-v-c3c032b5"]]);export{Ce as default};
