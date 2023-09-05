function X(t){return t.split("-")[1]}function Z(t){return t==="y"?"height":"width"}function Y(t){return t.split("-")[0]}function _(t){return["top","bottom"].includes(Y(t))?"x":"y"}function G(t,e,o){let{reference:n,floating:l}=t;const c=n.x+n.width/2-l.width/2,r=n.y+n.height/2-l.height/2,i=_(e),f=Z(i),s=n[f]/2-l[f]/2,u=i==="x";let a;switch(Y(e)){case"top":a={x:c,y:n.y-l.height};break;case"bottom":a={x:c,y:n.y+n.height};break;case"right":a={x:n.x+n.width,y:r};break;case"left":a={x:n.x-l.width,y:r};break;default:a={x:n.x,y:n.y}}switch(X(e)){case"start":a[i]-=s*(o&&u?-1:1);break;case"end":a[i]+=s*(o&&u?-1:1)}return a}const ut=async(t,e,o)=>{const{placement:n="bottom",strategy:l="absolute",middleware:c=[],platform:r}=o,i=c.filter(Boolean),f=await(r.isRTL==null?void 0:r.isRTL(e));let s=await r.getElementRects({reference:t,floating:e,strategy:l}),{x:u,y:a}=G(s,n,f),d=n,h={},y=0;for(let g=0;g<i.length;g++){const{name:m,fn:p}=i[g],{x,y:w,data:C,reset:v}=await p({x:u,y:a,initialPlacement:n,placement:d,strategy:l,middlewareData:h,rects:s,platform:r,elements:{reference:t,floating:e}});u=x??u,a=w??a,h={...h,[m]:{...h[m],...C}},v&&y<=50&&(y++,typeof v=="object"&&(v.placement&&(d=v.placement),v.rects&&(s=v.rects===!0?await r.getElementRects({reference:t,floating:e,strategy:l}):v.rects),{x:u,y:a}=G(s,d,f)),g=-1)}return{x:u,y:a,placement:d,strategy:l,middlewareData:h}};function dt(t){return typeof t!="number"?function(e){return{top:0,right:0,bottom:0,left:0,...e}}(t):{top:t,right:t,bottom:t,left:t}}function tt(t){return{...t,top:t.y,left:t.x,right:t.x+t.width,bottom:t.y+t.height}}const ht=Math.min,gt=Math.max;function pt(t,e,o){return gt(t,ht(e,o))}const bt=t=>({name:"arrow",options:t,async fn(e){const{element:o,padding:n=0}=t||{},{x:l,y:c,placement:r,rects:i,platform:f,elements:s}=e;if(o==null)return{};const u=dt(n),a={x:l,y:c},d=_(r),h=Z(d),y=await f.getDimensions(o),g=d==="y",m=g?"top":"left",p=g?"bottom":"right",x=g?"clientHeight":"clientWidth",w=i.reference[h]+i.reference[d]-a[d]-i.floating[h],C=a[d]-i.reference[d],v=await(f.getOffsetParent==null?void 0:f.getOffsetParent(o));let O=v?v[x]:0;O&&await(f.isElement==null?void 0:f.isElement(v))||(O=s.floating[x]||i.floating[h]);const ft=w/2-C/2,F=u[m],q=O-y[h]-u[p],E=O/2-y[h]/2+ft,M=pt(F,E,q),at=X(r)!=null&&E!=M&&i.reference[h]/2-(E<F?u[m]:u[p])-y[h]/2<0;return{[d]:a[d]-(at?E<F?F-E:q-E:0),data:{[d]:M,centerOffset:E-M}}}}),yt=["top","right","bottom","left"];yt.reduce((t,e)=>t.concat(e,e+"-start",e+"-end"),[]);const vt=function(t){return t===void 0&&(t=0),{name:"offset",options:t,async fn(e){const{x:o,y:n}=e,l=await async function(c,r){const{placement:i,platform:f,elements:s}=c,u=await(f.isRTL==null?void 0:f.isRTL(s.floating)),a=Y(i),d=X(i),h=_(i)==="x",y=["left","top"].includes(a)?-1:1,g=u&&h?-1:1,m=typeof r=="function"?r(c):r;let{mainAxis:p,crossAxis:x,alignmentAxis:w}=typeof m=="number"?{mainAxis:m,crossAxis:0,alignmentAxis:null}:{mainAxis:0,crossAxis:0,alignmentAxis:null,...m};return d&&typeof w=="number"&&(x=d==="end"?-1*w:w),h?{x:x*g,y:p*y}:{x:p*y,y:x*g}}(e,t);return{x:o+l.x,y:n+l.y,data:l}}}};function b(t){var e;return((e=t.ownerDocument)==null?void 0:e.defaultView)||window}function T(t){return b(t).getComputedStyle(t)}function et(t){return t instanceof b(t).Node}function D(t){return et(t)?(t.nodeName||"").toLowerCase():""}let P;function nt(){if(P)return P;const t=navigator.userAgentData;return t&&Array.isArray(t.brands)?(P=t.brands.map(e=>e.brand+"/"+e.version).join(" "),P):navigator.userAgent}function L(t){return t instanceof b(t).HTMLElement}function R(t){return t instanceof b(t).Element}function I(t){return typeof ShadowRoot>"u"?!1:t instanceof b(t).ShadowRoot||t instanceof ShadowRoot}function S(t){const{overflow:e,overflowX:o,overflowY:n,display:l}=T(t);return/auto|scroll|overlay|hidden|clip/.test(e+n+o)&&!["inline","contents"].includes(l)}function mt(t){return["table","td","th"].includes(D(t))}function N(t){const e=/firefox/i.test(nt()),o=T(t),n=o.backdropFilter||o.WebkitBackdropFilter;return o.transform!=="none"||o.perspective!=="none"||!!n&&n!=="none"||e&&o.willChange==="filter"||e&&!!o.filter&&o.filter!=="none"||["transform","perspective"].some(l=>o.willChange.includes(l))||["paint","layout","strict","content"].some(l=>{const c=o.contain;return c!=null&&c.includes(l)})}function j(){return/^((?!chrome|android).)*safari/i.test(nt())}function z(t){return["html","body","#document"].includes(D(t))}const J=Math.min,$=Math.max,V=Math.round;function ot(t){const e=T(t);let o=parseFloat(e.width),n=parseFloat(e.height);const l=L(t),c=l?t.offsetWidth:o,r=l?t.offsetHeight:n,i=V(o)!==c||V(n)!==r;return i&&(o=c,n=r),{width:o,height:n,fallback:i}}function it(t){return R(t)?t:t.contextElement}const rt={x:1,y:1};function W(t){const e=it(t);if(!L(e))return rt;const o=e.getBoundingClientRect(),{width:n,height:l,fallback:c}=ot(e);let r=(c?V(o.width):o.width)/n,i=(c?V(o.height):o.height)/l;return r&&Number.isFinite(r)||(r=1),i&&Number.isFinite(i)||(i=1),{x:r,y:i}}function k(t,e,o,n){var l,c;e===void 0&&(e=!1),o===void 0&&(o=!1);const r=t.getBoundingClientRect(),i=it(t);let f=rt;e&&(n?R(n)&&(f=W(n)):f=W(t));const s=i?b(i):window,u=j()&&o;let a=(r.left+(u&&((l=s.visualViewport)==null?void 0:l.offsetLeft)||0))/f.x,d=(r.top+(u&&((c=s.visualViewport)==null?void 0:c.offsetTop)||0))/f.y,h=r.width/f.x,y=r.height/f.y;if(i){const g=b(i),m=n&&R(n)?b(n):n;let p=g.frameElement;for(;p&&n&&m!==g;){const x=W(p),w=p.getBoundingClientRect(),C=getComputedStyle(p);w.x+=(p.clientLeft+parseFloat(C.paddingLeft))*x.x,w.y+=(p.clientTop+parseFloat(C.paddingTop))*x.y,a*=x.x,d*=x.y,h*=x.x,y*=x.y,a+=w.x,d+=w.y,p=b(p).frameElement}}return tt({width:h,height:y,x:a,y:d})}function A(t){return((et(t)?t.ownerDocument:t.document)||window.document).documentElement}function B(t){return R(t)?{scrollLeft:t.scrollLeft,scrollTop:t.scrollTop}:{scrollLeft:t.pageXOffset,scrollTop:t.pageYOffset}}function lt(t){return k(A(t)).left+B(t).scrollLeft}function H(t){if(D(t)==="html")return t;const e=t.assignedSlot||t.parentNode||I(t)&&t.host||A(t);return I(e)?e.host:e}function ct(t){const e=H(t);return z(e)?e.ownerDocument.body:L(e)&&S(e)?e:ct(e)}function st(t,e){var o;e===void 0&&(e=[]);const n=ct(t),l=n===((o=t.ownerDocument)==null?void 0:o.body),c=b(n);return l?e.concat(c,c.visualViewport||[],S(n)?n:[]):e.concat(n,st(n))}function K(t,e,o){let n;if(e==="viewport")n=function(r,i){const f=b(r),s=A(r),u=f.visualViewport;let a=s.clientWidth,d=s.clientHeight,h=0,y=0;if(u){a=u.width,d=u.height;const g=j();(!g||g&&i==="fixed")&&(h=u.offsetLeft,y=u.offsetTop)}return{width:a,height:d,x:h,y}}(t,o);else if(e==="document")n=function(r){const i=A(r),f=B(r),s=r.ownerDocument.body,u=$(i.scrollWidth,i.clientWidth,s.scrollWidth,s.clientWidth),a=$(i.scrollHeight,i.clientHeight,s.scrollHeight,s.clientHeight);let d=-f.scrollLeft+lt(r);const h=-f.scrollTop;return T(s).direction==="rtl"&&(d+=$(i.clientWidth,s.clientWidth)-u),{width:u,height:a,x:d,y:h}}(A(t));else if(R(e))n=function(r,i){const f=k(r,!0,i==="fixed"),s=f.top+r.clientTop,u=f.left+r.clientLeft,a=L(r)?W(r):{x:1,y:1};return{width:r.clientWidth*a.x,height:r.clientHeight*a.y,x:u*a.x,y:s*a.y}}(e,o);else{const r={...e};if(j()){var l,c;const i=b(t);r.x-=((l=i.visualViewport)==null?void 0:l.offsetLeft)||0,r.y-=((c=i.visualViewport)==null?void 0:c.offsetTop)||0}n=r}return tt(n)}function Q(t,e){return L(t)&&T(t).position!=="fixed"?e?e(t):t.offsetParent:null}function U(t,e){const o=b(t);if(!L(t))return o;let n=Q(t,e);for(;n&&mt(n)&&T(n).position==="static";)n=Q(n,e);return n&&(D(n)==="html"||D(n)==="body"&&T(n).position==="static"&&!N(n))?o:n||function(l){let c=H(l);for(;L(c)&&!z(c);){if(N(c))return c;c=H(c)}return null}(t)||o}function xt(t,e,o){const n=L(e),l=A(e),c=k(t,!0,o==="fixed",e);let r={scrollLeft:0,scrollTop:0};const i={x:0,y:0};if(n||!n&&o!=="fixed")if((D(e)!=="body"||S(l))&&(r=B(e)),L(e)){const f=k(e,!0);i.x=f.x+e.clientLeft,i.y=f.y+e.clientTop}else l&&(i.x=lt(l));return{x:c.left+r.scrollLeft-i.x,y:c.top+r.scrollTop-i.y,width:c.width,height:c.height}}const wt={getClippingRect:function(t){let{element:e,boundary:o,rootBoundary:n,strategy:l}=t;const c=o==="clippingAncestors"?function(s,u){const a=u.get(s);if(a)return a;let d=st(s).filter(m=>R(m)&&D(m)!=="body"),h=null;const y=T(s).position==="fixed";let g=y?H(s):s;for(;R(g)&&!z(g);){const m=T(g),p=N(g);m.position==="fixed"&&(h=null),(y?p||h:p||m.position!=="static"||!h||!["absolute","fixed"].includes(h.position))?h=m:d=d.filter(x=>x!==g),g=H(g)}return u.set(s,d),d}(e,this._c):[].concat(o),r=[...c,n],i=r[0],f=r.reduce((s,u)=>{const a=K(e,u,l);return s.top=$(a.top,s.top),s.right=J(a.right,s.right),s.bottom=J(a.bottom,s.bottom),s.left=$(a.left,s.left),s},K(e,i,l));return{width:f.right-f.left,height:f.bottom-f.top,x:f.left,y:f.top}},convertOffsetParentRelativeRectToViewportRelativeRect:function(t){let{rect:e,offsetParent:o,strategy:n}=t;const l=L(o),c=A(o);if(o===c)return e;let r={scrollLeft:0,scrollTop:0},i={x:1,y:1};const f={x:0,y:0};if((l||!l&&n!=="fixed")&&((D(o)!=="body"||S(c))&&(r=B(o)),L(o))){const s=k(o);i=W(o),f.x=s.x+o.clientLeft,f.y=s.y+o.clientTop}return{width:e.width*i.x,height:e.height*i.y,x:e.x*i.x-r.scrollLeft*i.x+f.x,y:e.y*i.y-r.scrollTop*i.y+f.y}},isElement:R,getDimensions:function(t){return ot(t)},getOffsetParent:U,getDocumentElement:A,getScale:W,async getElementRects(t){let{reference:e,floating:o,strategy:n}=t;const l=this.getOffsetParent||U,c=this.getDimensions;return{reference:xt(e,await l(o),n),floating:{x:0,y:0,...await c(o)}}},getClientRects:t=>Array.from(t.getClientRects()),isRTL:t=>T(t).direction==="rtl"},Lt=(t,e,o)=>{const n=new Map,l={platform:wt,...o},c={...l.platform,_c:n};return ut(t,e,{...l,platform:c})};export{vt as D,bt as u,Lt as z};
