import{H as h,e as y,j as d,f as m,ar as I,o as n,c as o,a as w,F as k,a8 as O,u as c,Q as p,V as v,P as x,W as L,a1 as P,n as R}from"./@vue-021a4eb1.js";import{u as V}from"./vue-router-5dc1aa30.js";const b={class:"layout-padding layout-padding-unset layout-iframe"},B={class:"layout-padding-auto layout-padding-view"},S=["src","data-url"],T=h({name:"layoutIframeView"}),E=h({...T,props:{refreshKey:{type:String,default:()=>""},name:{type:String,default:()=>"slide-right"},list:{type:Array,default:()=>[]}},setup(l){const s=l,f=y(),i=V(),g=d(()=>s.list.filter(e=>{var t;return(t=e.meta)==null?void 0:t.isIframeOpen})),_=d(()=>i.path),u=(e,t)=>{R(()=>{if(!f.value)return!1;f.value.forEach(r=>{r.dataset.url===e&&(r.onload=()=>{var a;(a=t.meta)!=null&&a.isIframeOpen&&t.meta.loading&&(t.meta.loading=!1)})})})};return m(()=>i.fullPath,e=>{const t=s.list.find(r=>r.path===e);if(!t)return!1;t.meta.isIframeOpen||(t.meta.isIframeOpen=!0),u(e,t)},{immediate:!0}),m(()=>s.refreshKey,()=>{const e=s.list.find(t=>t.path===i.path);if(!e)return!1;e.meta.isIframeOpen&&(e.meta.isIframeOpen=!1),setTimeout(()=>{e.meta.isIframeOpen=!0,e.meta.loading=!0,u(i.fullPath,e)})},{deep:!0}),(e,t)=>{const r=I("loading");return n(),o("div",b,[w("div",B,[(n(!0),o(k,null,O(c(g),a=>p((n(),o("div",{class:"w100",key:a.path,"element-loading-background":"white"},[v(P,{name:l.name},{default:x(()=>[p((n(),o("iframe",{src:a.meta.isLink,key:a.path,frameborder:"0",height:"100%",width:"100%",style:{position:"absolute"},"data-url":a.path,ref_for:!0,ref_key:"iframeRef",ref:f},null,8,S)),[[L,c(_)===a.path]])]),_:2},1032,["name"])])),[[r,a.meta.loading]])),128))])])}}});export{E as default};
