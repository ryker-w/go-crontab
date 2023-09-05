import{i as Z}from"./vue-demi-71ba0ef2.js";import{ao as H,e as D,an as I,ad as J,B as k,aw as O,Y as $,h as A,A as G,f as T,_ as tt,g as et,d as st,n as nt,$ as ot,j as ct}from"./@vue-021a4eb1.js";/*!
  * pinia v2.0.34
  * (c) 2023 Eduardo San Martin Morote
  * @license MIT
  */let N;const L=t=>N=t,W=Symbol();function C(t){return t&&typeof t=="object"&&Object.prototype.toString.call(t)==="[object Object]"&&typeof t.toJSON!="function"}var R;(function(t){t.direct="direct",t.patchObject="patch object",t.patchFunction="patch function"})(R||(R={}));function ht(){const t=H(!0),n=t.run(()=>D({}));let s=[],e=[];const r=I({install(u){L(r),r._a=u,u.provide(W,r),u.config.globalProperties.$pinia=r,e.forEach(a=>s.push(a)),e=[]},use(u){return!this._a&&!Z?e.push(u):s.push(u),this},_p:s,_a:null,_e:t,_s:new Map,state:n});return r}const Y=()=>{};function B(t,n,s,e=Y){t.push(n);const r=()=>{const u=t.indexOf(n);u>-1&&(t.splice(u,1),e())};return!s&&et()&&st(r),r}function j(t,...n){t.slice().forEach(s=>{s(...n)})}function x(t,n){t instanceof Map&&n instanceof Map&&n.forEach((s,e)=>t.set(e,s)),t instanceof Set&&n instanceof Set&&n.forEach(t.add,t);for(const s in n){if(!n.hasOwnProperty(s))continue;const e=n[s],r=t[s];C(r)&&C(e)&&t.hasOwnProperty(s)&&!k(e)&&!O(e)?t[s]=x(r,e):t[s]=e}return t}const rt=Symbol();function ut(t){return!C(t)||!t.hasOwnProperty(rt)}const{assign:y}=Object;function ft(t){return!!(k(t)&&t.effect)}function at(t,n,s,e){const{state:r,actions:u,getters:a}=n,f=s.state.value[t];let g;function b(){f||(s.state.value[t]=r?r():{});const v=ot(s.state.value[t]);return y(v,u,Object.keys(a||{}).reduce((d,m)=>(d[m]=I(ct(()=>{L(s);const _=s._s.get(t);return a[m].call(_,_)})),d),{}))}return g=q(t,b,n,s,e,!0),g}function q(t,n,s={},e,r,u){let a;const f=y({actions:{}},s),g={deep:!0};let b,v,d=I([]),m=I([]),_;const p=e.state.value[t];!u&&!p&&(e.state.value[t]={}),D({});let F;function V(c){let o;b=v=!1,typeof c=="function"?(c(e.state.value[t]),o={type:R.patchFunction,storeId:t,events:_}):(x(e.state.value[t],c),o={type:R.patchObject,payload:c,storeId:t,events:_});const h=F=Symbol();nt().then(()=>{F===h&&(b=!0)}),v=!0,j(d,o,e.state.value[t])}const z=u?function(){const{state:o}=s,h=o?o():{};this.$patch(S=>{y(S,h)})}:Y;function K(){a.stop(),d=[],m=[],e._s.delete(t)}function M(c,o){return function(){L(e);const h=Array.from(arguments),S=[],w=[];function U(i){S.push(i)}function X(i){w.push(i)}j(m,{args:h,name:c,store:l,after:U,onError:X});let E;try{E=o.apply(this&&this.$id===t?this:l,h)}catch(i){throw j(w,i),i}return E instanceof Promise?E.then(i=>(j(S,i),i)).catch(i=>(j(w,i),Promise.reject(i))):(j(S,E),E)}}const Q={_p:e,$id:t,$onAction:B.bind(null,m),$patch:V,$reset:z,$subscribe(c,o={}){const h=B(d,c,o.detached,()=>S()),S=a.run(()=>T(()=>e.state.value[t],w=>{(o.flush==="sync"?v:b)&&c({storeId:t,type:R.direct,events:_},w)},y({},g,o)));return h},$dispose:K},l=tt(Q);e._s.set(t,l);const P=e._e.run(()=>(a=H(),a.run(()=>n())));for(const c in P){const o=P[c];if(k(o)&&!ft(o)||O(o))u||(p&&ut(o)&&(k(o)?o.value=p[c]:x(o,p[c])),e.state.value[t][c]=o);else if(typeof o=="function"){const h=M(c,o);P[c]=h,f.actions[c]=o}}return y(l,P),y(J(l),P),Object.defineProperty(l,"$state",{get:()=>e.state.value[t],set:c=>{V(o=>{y(o,c)})}}),e._p.forEach(c=>{y(l,a.run(()=>c({store:l,app:e._a,pinia:e,options:f})))}),p&&u&&s.hydrate&&s.hydrate(l.$state,p),b=!0,v=!0,l}function bt(t,n,s){let e,r;const u=typeof n=="function";typeof t=="string"?(e=t,r=u?s:n):(r=t,e=t.id);function a(f,g){const b=A();return f=f||b&&G(W,null),f&&L(f),f=N,f._s.has(e)||(u?q(e,n,r,f):at(e,r,f)),f._s.get(e)}return a.$id=e,a}function yt(t){{t=J(t);const n={};for(const s in t){const e=t[s];(k(e)||O(e))&&(n[s]=$(t,s))}return n}}export{ht as c,bt as d,yt as s};
