import{c as $,_ as A}from"./index-35ac8c1e.js";import{s as S}from"./pinia-a8c871c6.js";import{E as V,a as O}from"./element-plus-f169eafc.js";import{H as v,ay as I,e as L,_ as N,i as P,ah as a,ar as F,o as w,c as H,V as e,P as t,a as m,T as n,Q,O as U,U as c,u as j}from"./@vue-021a4eb1.js";import"./vue-router-5dc1aa30.js";import"./js-cookie-cf83ad76.js";import"./@element-plus-4a2045cc.js";import"./nprogress-5ada7e0c.js";import"./call-bind-1df9733d.js";import"./get-intrinsic-b9397c9a.js";import"./has-symbols-e8f3ca0e.js";import"./function-bind-22e7ee79.js";import"./has-26d28e02.js";import"./axios-707ed124.js";import"./qs-a8db372b.js";import"./side-channel-9112baf6.js";import"./object-inspect-d11bccf2.js";import"./mitt-f7ef348c.js";import"./vue-demi-71ba0ef2.js";import"./lodash-es-36eb724a.js";import"./@vueuse-8a759679.js";import"./@popperjs-c75af06c.js";import"./@ctrl-f8748455.js";import"./dayjs-1798568c.js";import"./async-validator-dee29e8b.js";import"./memoize-one-297ddbcb.js";import"./escape-html-1d60d822.js";import"./normalize-wheel-es-ed76fb12.js";import"./@floating-ui-463e90e0.js";const q={class:"system-menu-container layout-pd"},G={class:"system-menu-search mb15"},J={class:"ml10"},K=v({name:"systemMenu"}),Te=v({...K,setup(W){const b=I(()=>A(()=>import("./dialog-c50f66b5.js"),["./dialog-c50f66b5.js","./index-35ac8c1e.js","./@vue-021a4eb1.js","./pinia-a8c871c6.js","./vue-demi-71ba0ef2.js","./vue-router-5dc1aa30.js","./element-plus-f169eafc.js","./lodash-es-36eb724a.js","./@vueuse-8a759679.js","./@element-plus-4a2045cc.js","./@popperjs-c75af06c.js","./@ctrl-f8748455.js","./dayjs-1798568c.js","./call-bind-1df9733d.js","./get-intrinsic-b9397c9a.js","./has-symbols-e8f3ca0e.js","./function-bind-22e7ee79.js","./has-26d28e02.js","./async-validator-dee29e8b.js","./memoize-one-297ddbcb.js","./escape-html-1d60d822.js","./normalize-wheel-es-ed76fb12.js","./@floating-ui-463e90e0.js","./js-cookie-cf83ad76.js","./nprogress-5ada7e0c.js","..\\css\\nprogress-8b89e2e0.css","./axios-707ed124.js","./qs-a8db372b.js","./side-channel-9112baf6.js","./object-inspect-d11bccf2.js","./mitt-f7ef348c.js","..\\css\\index-c2520709.css"],import.meta.url)),y=$(),{routesList:g}=S(y),d=L(),r=N({tableData:{data:[],loading:!0}}),_=()=>{r.tableData.loading=!0,r.tableData.data=g.value,setTimeout(()=>{r.tableData.loading=!1},500)},u=s=>{d.value.openDialog(s)},D=(s,i)=>{d.value.openDialog(s,i)},x=s=>{V.confirm(`此操作将永久删除路由：${s.path}, 是否继续?`,"提示",{confirmButtonText:"删除",cancelButtonText:"取消",type:"warning"}).then(()=>{O.success("删除成功"),_()}).catch(()=>{})};return P(()=>{_()}),(s,i)=>{const C=a("el-input"),k=a("ele-Search"),f=a("el-icon"),p=a("el-button"),T=a("ele-FolderAdd"),z=a("SvgIcon"),l=a("el-table-column"),B=a("el-tag"),E=a("el-table"),M=a("el-card"),R=F("loading");return w(),H("div",q,[e(M,{shadow:"hover"},{default:t(()=>[m("div",G,[e(C,{size:"default",placeholder:"请输入菜单名称",style:{"max-width":"180px"}}),e(p,{size:"default",type:"primary",class:"ml10"},{default:t(()=>[e(f,null,{default:t(()=>[e(k)]),_:1}),n(" 查询 ")]),_:1}),e(p,{size:"default",type:"success",class:"ml10",onClick:u},{default:t(()=>[e(f,null,{default:t(()=>[e(T)]),_:1}),n(" 新增菜单 ")]),_:1})]),Q((w(),U(E,{data:r.tableData.data,style:{width:"100%"},"row-key":"path","tree-props":{children:"children",hasChildren:"hasChildren"}},{default:t(()=>[e(l,{label:"菜单名称","show-overflow-tooltip":""},{default:t(o=>[e(z,{name:o.row.meta.icon},null,8,["name"]),m("span",J,c(o.row.meta.title),1)]),_:1}),e(l,{prop:"path",label:"路由路径","show-overflow-tooltip":""}),e(l,{label:"组件路径","show-overflow-tooltip":""},{default:t(o=>[m("span",null,c(o.row.component),1)]),_:1}),e(l,{label:"权限标识","show-overflow-tooltip":""},{default:t(o=>[m("span",null,c(o.row.meta.roles),1)]),_:1}),e(l,{label:"排序","show-overflow-tooltip":"",width:"80"},{default:t(o=>[n(c(o.$index),1)]),_:1}),e(l,{label:"类型","show-overflow-tooltip":"",width:"80"},{default:t(o=>[e(B,{type:"success",size:"small"},{default:t(()=>[n(c(o.row.xx)+"菜单",1)]),_:2},1024)]),_:1}),e(l,{label:"操作","show-overflow-tooltip":"",width:"140"},{default:t(o=>[e(p,{size:"small",text:"",type:"primary",onClick:i[0]||(i[0]=h=>u("add"))},{default:t(()=>[n("新增")]),_:1}),e(p,{size:"small",text:"",type:"primary",onClick:h=>D("edit",o.row)},{default:t(()=>[n("修改")]),_:2},1032,["onClick"]),e(p,{size:"small",text:"",type:"primary",onClick:h=>x(o.row)},{default:t(()=>[n("删除")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"])),[[R,r.tableData.loading]])]),_:1}),e(j(b),{ref_key:"menuDialogRef",ref:d,onRefresh:i[1]||(i[1]=o=>_())},null,512)])}}});export{Te as default};
