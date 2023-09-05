import{H as h,e as P,_ as $,ah as s,o as u,c as x,V as e,P as l,a as F,T as V,U as E,F as O,a8 as A,O as p}from"./@vue-021a4eb1.js";const H={class:"system-dic-dialog-container"},I=F("span",{class:"ml10"},"字段",-1),L={class:"dialog-footer"},Y=h({name:"systemDicDialog"}),G=h({...Y,emits:["refresh"],setup(j,{expose:y,emit:D}){const w=P(),o=$({ruleForm:{dicName:"",fieldName:"",status:!0,list:[],describe:""},dialog:{isShowDialog:!1,type:"",title:"",submitTxt:""}}),v=(r,a)=>{r==="edit"?(a.fieldName==="SYS_UERINFO"?a.list=[{id:Math.random(),label:"sex",value:"1"},{id:Math.random(),label:"sex",value:"0"}]:a.list=[{id:Math.random(),label:"role",value:"admin"},{id:Math.random(),label:"role",value:"common"},{id:Math.random(),label:"roleName",value:"超级管理员"},{id:Math.random(),label:"roleName",value:"普通用户"}],o.ruleForm=a,o.dialog.title="修改字典",o.dialog.submitTxt="修 改"):(o.dialog.title="新增字典",o.dialog.submitTxt="新 增"),o.dialog.isShowDialog=!0},f=()=>{o.dialog.isShowDialog=!1},N=()=>{f()},U=()=>{f(),D("refresh")},C=()=>{o.ruleForm.list.push({id:Math.random(),label:"",value:""})},S=r=>{o.ruleForm.list.splice(r,1)};return y({openDialog:v}),(r,a)=>{const k=s("el-alert"),i=s("el-input"),m=s("el-form-item"),d=s("el-col"),M=s("el-switch"),T=s("ele-Plus"),g=s("el-icon"),c=s("el-button"),z=s("ele-Delete"),b=s("el-row"),R=s("el-form"),B=s("el-dialog");return u(),x("div",H,[e(B,{title:o.dialog.title,modelValue:o.dialog.isShowDialog,"onUpdate:modelValue":a[4]||(a[4]=t=>o.dialog.isShowDialog=t),width:"769px"},{footer:l(()=>[F("span",L,[e(c,{onClick:N,size:"default"},{default:l(()=>[V("取 消")]),_:1}),e(c,{type:"primary",onClick:U,size:"default"},{default:l(()=>[V(E(o.dialog.submitTxt),1)]),_:1})])]),default:l(()=>[e(k,{title:"半成品，交互过于复杂，请自行扩展！",type:"warning",closable:!1,class:"mb20"}),e(R,{ref_key:"dicDialogFormRef",ref:w,model:o.ruleForm,size:"default","label-width":"90px"},{default:l(()=>[e(b,{gutter:35},{default:l(()=>[e(d,{xs:24,sm:12,md:12,lg:12,xl:12,class:"mb20"},{default:l(()=>[e(m,{label:"字典名称"},{default:l(()=>[e(i,{modelValue:o.ruleForm.dicName,"onUpdate:modelValue":a[0]||(a[0]=t=>o.ruleForm.dicName=t),placeholder:"请输入字典名称",clearable:""},null,8,["modelValue"])]),_:1})]),_:1}),e(d,{xs:24,sm:12,md:12,lg:12,xl:12,class:"mb20"},{default:l(()=>[e(m,{label:"字段名"},{default:l(()=>[e(i,{modelValue:o.ruleForm.fieldName,"onUpdate:modelValue":a[1]||(a[1]=t=>o.ruleForm.fieldName=t),placeholder:"请输入字段名，拼接 ruleForm.list",clearable:""},null,8,["modelValue"])]),_:1})]),_:1}),e(d,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:l(()=>[e(m,{label:"字典状态"},{default:l(()=>[e(M,{modelValue:o.ruleForm.status,"onUpdate:modelValue":a[2]||(a[2]=t=>o.ruleForm.status=t),"inline-prompt":"","active-text":"启","inactive-text":"禁"},null,8,["modelValue"])]),_:1})]),_:1}),e(d,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:l(()=>[(u(!0),x(O,null,A(o.ruleForm.list,(t,n)=>(u(),p(b,{gutter:35,key:n},{default:l(()=>[e(d,{xs:24,sm:12,md:12,lg:12,xl:12,class:"mb20"},{default:l(()=>[e(m,{prop:`list[${n}].label`},{label:l(()=>[n===0?(u(),p(c,{key:0,type:"primary",circle:"",size:"small",onClick:C},{default:l(()=>[e(g,null,{default:l(()=>[e(T)]),_:1})]),_:1})):(u(),p(c,{key:1,type:"danger",circle:"",size:"small",onClick:_=>S(n)},{default:l(()=>[e(g,null,{default:l(()=>[e(z)]),_:1})]),_:2},1032,["onClick"])),I]),default:l(()=>[e(i,{modelValue:t.label,"onUpdate:modelValue":_=>t.label=_,style:{width:"100%"},placeholder:"请输入字段名"},null,8,["modelValue","onUpdate:modelValue"])]),_:2},1032,["prop"])]),_:2},1024),e(d,{xs:24,sm:12,md:12,lg:12,xl:12,class:"mb20"},{default:l(()=>[e(m,{label:"属性",prop:`list[${n}].value`},{default:l(()=>[e(i,{modelValue:t.value,"onUpdate:modelValue":_=>t.value=_,style:{width:"100%"},placeholder:"请输入属性值"},null,8,["modelValue","onUpdate:modelValue"])]),_:2},1032,["prop"])]),_:2},1024)]),_:2},1024))),128))]),_:1}),e(d,{xs:24,sm:24,md:24,lg:24,xl:24,class:"mb20"},{default:l(()=>[e(m,{label:"字典描述"},{default:l(()=>[e(i,{modelValue:o.ruleForm.describe,"onUpdate:modelValue":a[3]||(a[3]=t=>o.ruleForm.describe=t),type:"textarea",placeholder:"请输入字典描述",maxlength:"150"},null,8,["modelValue"])]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])]),_:1},8,["title","modelValue"])])}}});export{G as default};