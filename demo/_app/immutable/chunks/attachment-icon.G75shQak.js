function __vite__mapDeps(indexes) {
  if (!__vite__mapDeps.viteFileDeps) {
    __vite__mapDeps.viteFileDeps = ["./highlight.DGAr_9BG.js","./preload-helper.D6kgxu3v.js"]
  }
  return indexes.map((i) => __vite__mapDeps.viteFileDeps[i])
}
import{s as te,o as Y,p as Ve,q as He,u as Ie,r as ze,v as Ce,e as V,aa as ht,a as U,c as B,b as C,ac as _t,g as j,f as T,m as w,i as H,h as I,H as L,J as x,T as gt,D as dt,L as F,l as pt,z as Be,ag as qe,R as Le,w as ae,E as Me,F as Re,n as ve,B as q,W as we,ah as ke,t as Ee,d as Ne,j as Ae,U as be,M as G,ai as Se,aj as De,K as Pe,O as de}from"./scheduler.CjD0NZ8o.js";import{S as le,i as ne,c as P,a as K,m as W,t as b,b as k,d as O,g as S,e as D,f as me}from"./index.or77jMZA.js";import{e as Q}from"./each.BFF0fzu3.js";import{g as y,a as ee}from"./spread.CgU5AtxT.js";import{d as ge,a as bt}from"./data.BGGdSGzA.js";import{_ as We}from"./preload-helper.D6kgxu3v.js";import{I as vt}from"./Icon.BC0WXsRo.js";function wt(i){let e;const l=i[2].default,n=He(l,i,i[3],null);return{c(){n&&n.c()},l(t){n&&n.l(t)},m(t,o){n&&n.m(t,o),e=!0},p(t,o){n&&n.p&&(!e||o&8)&&Ie(n,l,t,t[3],e?Ce(l,t[3],o,null):ze(t[3]),null)},i(t){e||(b(n,t),e=!0)},o(t){k(n,t),e=!1},d(t){n&&n.d(t)}}}function kt(i){let e,l;const n=[{name:"image-plus"},i[1],{iconNode:i[0]}];let t={$$slots:{default:[wt]},$$scope:{ctx:i}};for(let o=0;o<n.length;o+=1)t=Y(t,n[o]);return e=new vt({props:t}),{c(){P(e.$$.fragment)},l(o){K(e.$$.fragment,o)},m(o,a){W(e,o,a),l=!0},p(o,[a]){const r=a&3?y(n,[n[0],a&2&&ee(o[1]),a&1&&{iconNode:o[0]}]):{};a&8&&(r.$$scope={dirty:a,ctx:o}),e.$set(r)},i(o){l||(b(e.$$.fragment,o),l=!0)},o(o){k(e.$$.fragment,o),l=!1},d(o){O(e,o)}}}function Tt(i,e,l){let{$$slots:n={},$$scope:t}=e;const o=[["path",{d:"M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7"}],["line",{x1:"16",x2:"22",y1:"5",y2:"5"}],["line",{x1:"19",x2:"19",y1:"2",y2:"8"}],["circle",{cx:"9",cy:"9",r:"2"}],["path",{d:"m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"}]];return i.$$set=a=>{l(1,e=Y(Y({},e),Ve(a))),"$$scope"in a&&l(3,t=a.$$scope)},e=Ve(e),[o,e,n,t]}class Et extends le{constructor(e){super(),ne(this,e,Tt,kt,te,{})}}function Oe(i){let e;const l=i[6].default,n=He(l,i,i[5],null);return{c(){n&&n.c()},l(t){n&&n.l(t)},m(t,o){n&&n.m(t,o),e=!0},p(t,o){n&&n.p&&(!e||o&32)&&Ie(n,l,t,t[5],e?Ce(l,t[5],o,null):ze(t[5]),null)},i(t){e||(b(n,t),e=!0)},o(t){k(n,t),e=!1},d(t){n&&n.d(t)}}}function Nt(i){let e,l,n,t,o,a,r=i[1]&&Oe(i);return{c(){e=V("div"),l=new ht(!1),n=U(),r&&r.c(),this.h()},l(s){e=B(s,"DIV",{class:!0});var f=C(e);l=_t(f,!1),n=j(f),r&&r.l(f),f.forEach(T),this.h()},h(){l.a=n,w(e,"class","carta-renderer markdown-body svelte-r6n2gn")},m(s,f){H(s,e,f),l.m(i[2],e),I(e,n),r&&r.m(e,null),i[8](e),t=!0,o||(a=L(e,"scroll",i[7]),o=!0)},p(s,[f]){(!t||f&4)&&l.p(s[2]),s[1]?r?(r.p(s,f),f&2&&b(r,1)):(r=Oe(s),r.c(),b(r,1),r.m(e,null)):r&&(S(),k(r,1,1,()=>{r=null}),D())},i(s){t||(b(r),t=!0)},o(s){k(r),t=!1},d(s){s&&T(e),r&&r.d(),i[8](null),o=!1,a()}}}function At(i,e,l){let{$$slots:n={},$$scope:t}=e,{carta:o}=e,{value:a}=e,{elem:r}=e,s=!1,f=o.renderSSR(a);const c=ge(u=>{o.render(u).then(_=>{l(2,f=""),l(2,f=_)}).then(()=>h("render",void 0))},o.rendererDebounce??300),d=u=>{c(u)};x(()=>o.$setRenderer(r)),x(()=>l(1,s=!0));const h=gt();function m(u){dt.call(this,i,u)}function g(u){F[u?"unshift":"push"](()=>{r=u,l(0,r)})}return i.$$set=u=>{"carta"in u&&l(3,o=u.carta),"value"in u&&l(4,a=u.value),"elem"in u&&l(0,r=u.elem),"$$scope"in u&&l(5,t=u.$$scope)},i.$$.update=()=>{i.$$.dirty&18&&s&&d(a)},[r,s,f,o,a,t,n,m,g]}class Ht extends le{constructor(e){super(),ne(this,e,At,Nt,te,{carta:3,value:4,elem:0})}}function It(i,e,l){const n=bt(e.replaceAll(`\r
`,`
`),l.replaceAll(`\r
`,`
`)),t=zt(i);if(t.length==0){const d=Fe(i);if(!d)return;t.push(d)}let o=0,a=0;const r=d=>{let h=0;for(;h<d&&o<t.length;){const u=(t[o].textContent??"").length-a,_=Math.min(u,d-h);h+=_,a+=_,_==u&&(o++,a=0)}},s=d=>{r(d.value.length)},f=d=>{const h=t.at(o)??Fe(i);if(!h)return;const m=h.textContent??"",g=m.slice(0,a),u=m.slice(a);h.textContent=g+d.value+u,r(d.value.length)},c=d=>{let h=d.value.length;for(;h>0;){const m=t.at(o);if(!m)return;const g=m.textContent??"",u=g.length-a,_=Math.min(u,h);m.textContent=g.slice(0,a)+g.slice(a+_),h-=_,_==u&&(o++,a=0)}};for(const d of n)d.added?f(d):d.removed?c(d):s(d)}function zt(i){const e=[],l=document.createTreeWalker(i,NodeFilter.SHOW_TEXT);for(;l.nextNode();)e.push(l.currentNode);return e}function Fe(i){const e=Array.from(i.querySelectorAll(".line")).at(-1);if(!e)return null;const l=document.createTextNode("");return e.appendChild(l),e.setAttribute("data-temp-node","true"),l}function Ue(i){let e;const l=i[14].default,n=He(l,i,i[13],null);return{c(){n&&n.c()},l(t){n&&n.l(t)},m(t,o){n&&n.m(t,o),e=!0},p(t,o){n&&n.p&&(!e||o&8192)&&Ie(n,l,t,t[13],e?Ce(l,t[13],o,null):ze(t[13]),null)},i(t){e||(b(n,t),e=!0)},o(t){k(n,t),e=!1},d(t){n&&n.d(t)}}}function Ct(i){let e,l="Press ESC then TAB to move the focus off the field",n,t,o,a,r,s,f,c,d,h,m,g=[{spellcheck:"false"},{class:"carta-font-code"},{"aria-multiline":"true"},{"aria-describedby":"editor-unfocus-suggestion"},{tabindex:"0"},{placeholder:i[2]},i[3]],u={};for(let p=0;p<g.length;p+=1)u=Y(u,g[p]);let _=i[8]&&Ue(i);return{c(){e=V("div"),e.textContent=l,n=U(),t=V("div"),o=V("div"),a=V("div"),r=new ht(!1),s=U(),f=V("textarea"),c=U(),_&&_.c(),this.h()},l(p){e=B(p,"DIV",{role:!0,id:!0,class:!0,"data-svelte-h":!0}),pt(e)!=="svelte-167gk2c"&&(e.textContent=l),n=j(p),t=B(p,"DIV",{role:!0,tabindex:!0,class:!0});var E=C(t);o=B(E,"DIV",{class:!0});var z=C(o);a=B(z,"DIV",{class:!0,tabindex:!0,"aria-hidden":!0});var R=C(a);r=_t(R,!1),R.forEach(T),s=j(z),f=B(z,"TEXTAREA",{spellcheck:!0,class:!0,"aria-multiline":!0,"aria-describedby":!0,tabindex:!0,placeholder:!0}),C(f).forEach(T),z.forEach(T),c=j(E),_&&_.l(E),E.forEach(T),this.h()},h(){w(e,"role","tooltip"),w(e,"id","editor-unfocus-suggestion"),w(e,"class","svelte-1fa8oqk"),r.a=null,w(a,"class","carta-highlight carta-font-code svelte-1fa8oqk"),w(a,"tabindex","-1"),w(a,"aria-hidden","true"),Be(f,u),qe(f,"svelte-1fa8oqk",!0),w(o,"class","carta-input-wrapper svelte-1fa8oqk"),w(t,"role","textbox"),w(t,"tabindex","-1"),w(t,"class","carta-input svelte-1fa8oqk")},m(p,E){H(p,e,E),H(p,n,E),H(p,t,E),I(t,o),I(o,a),r.m(i[7],a),i[16](a),I(o,s),I(o,f),f.autofocus&&f.focus(),Le(f,i[0]),i[18](f),i[21](o),I(t,c),_&&_.m(t,null),i[22](t),d=!0,h||(m=[L(f,"input",i[17]),L(f,"scroll",i[19]),L(f,"keydown",i[20]),L(t,"click",i[10]),L(t,"keydown",i[10]),L(t,"scroll",i[15])],h=!0)},p(p,[E]){(!d||E&128)&&r.p(p[7]),Be(f,u=y(g,[{spellcheck:"false"},{class:"carta-font-code"},{"aria-multiline":"true"},{"aria-describedby":"editor-unfocus-suggestion"},{tabindex:"0"},(!d||E&4)&&{placeholder:p[2]},E&8&&p[3]])),E&1&&Le(f,p[0]),qe(f,"svelte-1fa8oqk",!0),p[8]?_?(_.p(p,E),E&256&&b(_,1)):(_=Ue(p),_.c(),b(_,1),_.m(t,null)):_&&(S(),k(_,1,1,()=>{_=null}),D())},i(p){d||(b(_),d=!0)},o(p){k(_),d=!1},d(p){p&&(T(e),T(n),T(t)),i[16](null),i[18](null),i[21](null),_&&_.d(),i[22](null),h=!1,ae(m)}}}function St(i,e,l){let{$$slots:n={},$$scope:t}=e,{carta:o}=e,{value:a=""}=e,{placeholder:r=""}=e,{elem:s}=e,{props:f={}}=e,c,d,h,m=a,g=!1,u=a;const _=()=>{if(!g||!c||(l(4,c.style.height="0",c),l(4,c.style.minHeight="0",c),l(4,c.style.height=c.scrollHeight+"px",c),l(4,c.style.minHeight=h.clientHeight+"px",c),l(4,c.scrollTop=0,c),!(document.activeElement===c))||!o.input)return;const v=o.input.getCursorXY();v&&(v.top<0||v.top+o.input.getRowHeight()>=s.scrollTop+s.clientHeight)&&s.scrollTo({top:v==null?void 0:v.top,behavior:"instant"})},p=()=>{var v;(v=window.getSelection())!=null&&v.toString()||c==null||c.focus()},z=ge(async N=>{const v=await o.highlighter();if(!v)return;let M;const X=await We(()=>import("./highlight.DGAr_9BG.js"),__vite__mapDeps([0,1]),import.meta.url),{isSingleTheme:_e}=X;_e(v.theme)?M=v.codeToHtml(N,{lang:v.lang,theme:v.theme}):M=v.codeToHtml(N,{lang:v.lang,themes:v.theme}),o.sanitizer?l(7,m=o.sanitizer(M)):l(7,m=M),requestAnimationFrame(_)},250),R=ge(async N=>{const v=await o.highlighter(),M=await We(()=>import("./highlight.DGAr_9BG.js"),__vite__mapDeps([0,1]),import.meta.url),{loadNestedLanguages:X}=M;if(!v)return;const{updated:_e}=await X(v,N);_e&&z(N)},300),Z=N=>{d&&(It(d,u,N),requestAnimationFrame(_)),z(N),R(N)};x(()=>{l(8,g=!0),requestAnimationFrame(_)}),x(()=>{o.$setInput(c,s,()=>{l(0,a=c.value)})});function ie(N){dt.call(this,i,N)}function se(N){F[N?"unshift":"push"](()=>{d=N,l(5,d)})}function oe(){a=this.value,l(0,a)}function $(N){F[N?"unshift":"push"](()=>{c=N,l(4,c)})}const fe=()=>l(4,c.scrollTop=0,c),ce=()=>l(9,u=a);function ue(N){F[N?"unshift":"push"](()=>{h=N,l(6,h)})}function he(N){F[N?"unshift":"push"](()=>{s=N,l(1,s)})}return i.$$set=N=>{"carta"in N&&l(11,o=N.carta),"value"in N&&l(0,a=N.value),"placeholder"in N&&l(2,r=N.placeholder),"elem"in N&&l(1,s=N.elem),"props"in N&&l(3,f=N.props),"$$scope"in N&&l(13,t=N.$$scope)},i.$$.update=()=>{i.$$.dirty&1&&Z(a)},[a,s,r,f,c,d,h,m,g,u,p,o,_,t,n,ie,se,oe,$,fe,ce,ue,he]}class Dt extends le{constructor(e){super(),ne(this,e,St,Ct,te,{carta:11,value:0,placeholder:2,elem:1,props:3,resize:12})}get resize(){return this.$$.ctx[12]}}const Vt={writeTab:"Write",previewTab:"Preview",iconsLabels:{}};function pe(i){var t;if(i.key!=="ArrowLeft"&&i.key!=="ArrowRight")return;i.preventDefault();const e=(t=i.currentTarget.parentElement)==null?void 0:t.children;if(!e)return;const l=i.currentTarget.nextElementSibling??e[0],n=i.currentTarget.previousElementSibling??e[e.length-1];i.key==="ArrowRight"?l.focus():n.focus()}function Bt(i){let e,l;return{c(){e=Me("svg"),l=Me("path"),this.h()},l(n){e=Re(n,"svg",{xmlns:!0,width:!0,height:!0,viewBox:!0});var t=C(e);l=Re(t,"path",{fill:!0,d:!0}),C(l).forEach(T),t.forEach(T),this.h()},h(){w(l,"fill","currentColor"),w(l,"d","M4 8a2 2 0 1 1-3.999.001A2 2 0 0 1 4 8m6 0a2 2 0 1 1-3.999.001A2 2 0 0 1 10 8m6 0a2 2 0 1 1-3.999.001A2 2 0 0 1 16 8"),w(e,"xmlns","http://www.w3.org/2000/svg"),w(e,"width","1em"),w(e,"height","1em"),w(e,"viewBox","0 0 16 16")},m(n,t){H(n,e,t),I(e,l)},p:ve,i:ve,o:ve,d(n){n&&T(e)}}}class qt extends le{constructor(e){super(),ne(this,e,null,Bt,te,{})}}function je(i,e,l){const n=i.slice();n[28]=e[l];const t=n[3].iconsLabels[n[28].id]??n[28].label;return n[29]=t,n}function Te(i){const e=i.slice(),l=e[3].iconsLabels.menu??"Menu";return e[29]=l,e}function Xe(i,e,l){const n=i.slice();n[28]=e[l],n[33]=l;const t=n[3].iconsLabels[n[28].id]??n[28].label;return n[29]=t,n}function Ke(i){let e,l=i[3].writeTab+"",n,t,o,a,r=i[3].previewTab+"",s,f,c,d;return{c(){e=V("button"),n=Ee(l),o=U(),a=V("button"),s=Ee(r),this.h()},l(h){e=B(h,"BUTTON",{type:!0,tabindex:!0,class:!0});var m=C(e);n=Ne(m,l),m.forEach(T),o=j(h),a=B(h,"BUTTON",{type:!0,tabindex:!0,class:!0});var g=C(a);s=Ne(g,r),g.forEach(T),this.h()},h(){w(e,"type","button"),w(e,"tabindex",0),w(e,"class",t=i[0]==="write"?"carta-active":""),w(a,"type","button"),w(a,"tabindex",-1),w(a,"class",f=i[0]==="preview"?"carta-active":"")},m(h,m){H(h,e,m),I(e,n),H(h,o,m),H(h,a,m),I(a,s),c||(d=[L(e,"click",i[15]),L(e,"keydown",pe),L(a,"click",i[16]),L(a,"keydown",pe)],c=!0)},p(h,m){m[0]&8&&l!==(l=h[3].writeTab+"")&&Ae(n,l),m[0]&1&&t!==(t=h[0]==="write"?"carta-active":"")&&w(e,"class",t),m[0]&8&&r!==(r=h[3].previewTab+"")&&Ae(s,r),m[0]&1&&f!==(f=h[0]==="preview"?"carta-active":"")&&w(a,"class",f)},d(h){h&&(T(e),T(o),T(a)),c=!1,ae(d)}}}function Je(i){let e,l,n,t=Q(i[4]),o=[];for(let s=0;s<t.length;s+=1)o[s]=Ye(Xe(i,t,s));const a=s=>k(o[s],1,1,()=>{o[s]=null});let r=i[11]&&Ge(Te(i));return{c(){for(let s=0;s<o.length;s+=1)o[s].c();e=U(),r&&r.c(),l=q()},l(s){for(let f=0;f<o.length;f+=1)o[f].l(s);e=j(s),r&&r.l(s),l=q()},m(s,f){for(let c=0;c<o.length;c+=1)o[c]&&o[c].m(s,f);H(s,e,f),r&&r.m(s,f),H(s,l,f),n=!0},p(s,f){if(f[0]&538){t=Q(s[4]);let c;for(c=0;c<t.length;c+=1){const d=Xe(s,t,c);o[c]?(o[c].p(d,f),b(o[c],1)):(o[c]=Ye(d),o[c].c(),b(o[c],1),o[c].m(e.parentNode,e))}for(S(),c=t.length;c<o.length;c+=1)a(c);D()}s[11]?r?(r.p(Te(s),f),f[0]&2048&&b(r,1)):(r=Ge(Te(s)),r.c(),b(r,1),r.m(l.parentNode,l)):r&&(S(),k(r,1,1,()=>{r=null}),D())},i(s){if(!n){for(let f=0;f<t.length;f+=1)b(o[f]);b(r),n=!0}},o(s){o=o.filter(Boolean);for(let f=0;f<o.length;f+=1)k(o[f]);k(r),n=!1},d(s){s&&(T(e),T(l)),be(o,s),r&&r.d(s)}}}function Ye(i){let e,l,n,t,o,a,r,s;var f=i[28].component;function c(h,m){return{}}f&&(l=G(f,c()));function d(){return i[19](i[28])}return{c(){e=V("button"),l&&P(l.$$.fragment),this.h()},l(h){e=B(h,"BUTTON",{class:!0,tabindex:!0,title:!0,"aria-label":!0});var m=C(e);l&&K(l.$$.fragment,m),m.forEach(T),this.h()},h(){w(e,"class","carta-icon svelte-1c77udu"),w(e,"tabindex",i[33]==0?0:-1),w(e,"title",n=i[29]),w(e,"aria-label",t=i[29]),we(()=>i[18].call(e))},m(h,m){H(h,e,m),l&&W(l,e,null),o=ke(e,i[18].bind(e)),a=!0,r||(s=[L(e,"click",Se(De(d))),L(e,"keydown",pe)],r=!0)},p(h,m){if(i=h,m[0]&16&&f!==(f=i[28].component)){if(l){S();const g=l;k(g.$$.fragment,1,0,()=>{O(g,1)}),D()}f?(l=G(f,c()),P(l.$$.fragment),b(l.$$.fragment,1),W(l,e,null)):l=null}(!a||m[0]&24&&n!==(n=i[29]))&&w(e,"title",n),(!a||m[0]&24&&t!==(t=i[29]))&&w(e,"aria-label",t)},i(h){a||(l&&b(l.$$.fragment,h),a=!0)},o(h){l&&k(l.$$.fragment,h),a=!1},d(h){h&&T(e),l&&O(l),o(),r=!1,ae(s)}}}function Ge(i){let e,l,n,t,o,a,r;return l=new qt({}),{c(){e=V("button"),P(l.$$.fragment),this.h()},l(s){e=B(s,"BUTTON",{class:!0,tabindex:!0,title:!0,"aria-label":!0});var f=C(e);K(l.$$.fragment,f),f.forEach(T),this.h()},h(){w(e,"class","carta-icon svelte-1c77udu"),w(e,"tabindex",-1),w(e,"title",n=i[29]),w(e,"aria-label",t=i[29])},m(s,f){H(s,e,f),W(l,e,null),o=!0,a||(r=[L(e,"keydown",pe),L(e,"click",Se(De(i[20])))],a=!0)},p(s,f){(!o||f[0]&8&&n!==(n=s[29]))&&w(e,"title",n),(!o||f[0]&8&&t!==(t=s[29]))&&w(e,"aria-label",t)},i(s){o||(b(l.$$.fragment,s),o=!0)},o(s){k(l.$$.fragment,s),o=!1},d(s){s&&T(e),O(l),a=!1,ae(r)}}}function Qe(i){let e,l,n=Q(i[1].icons.filter(i[24])),t=[];for(let a=0;a<n.length;a+=1)t[a]=Ze(je(i,n,a));const o=a=>k(t[a],1,1,()=>{t[a]=null});return{c(){e=V("div");for(let a=0;a<t.length;a+=1)t[a].c();this.h()},l(a){e=B(a,"DIV",{class:!0,style:!0});var r=C(e);for(let s=0;s<t.length;s+=1)t[s].l(r);r.forEach(T),this.h()},h(){w(e,"class","carta-icons-menu svelte-1c77udu"),Pe(e,"top",i[10]+"px")},m(a,r){H(a,e,r);for(let s=0;s<t.length;s+=1)t[s]&&t[s].m(e,null);i[26](e),l=!0},p(a,r){if(r[0]&4122){n=Q(a[1].icons.filter(a[24]));let s;for(s=0;s<n.length;s+=1){const f=je(a,n,s);t[s]?(t[s].p(f,r),b(t[s],1)):(t[s]=Ze(f),t[s].c(),b(t[s],1),t[s].m(e,null))}for(S(),s=n.length;s<t.length;s+=1)o(s);D()}(!l||r[0]&1024)&&Pe(e,"top",a[10]+"px")},i(a){if(!l){for(let r=0;r<n.length;r+=1)b(t[r]);l=!0}},o(a){t=t.filter(Boolean);for(let r=0;r<t.length;r+=1)k(t[r]);l=!1},d(a){a&&T(e),be(t,a),i[26](null)}}}function Ze(i){let e,l,n,t,o=i[29]+"",a,r,s,f,c,d;var h=i[28].component;function m(u,_){return{}}h&&(l=G(h,m()));function g(){return i[25](i[28])}return{c(){e=V("button"),l&&P(l.$$.fragment),n=U(),t=V("span"),a=Ee(o),r=U(),this.h()},l(u){e=B(u,"BUTTON",{class:!0,"aria-label":!0});var _=C(e);l&&K(l.$$.fragment,_),n=j(_),t=B(_,"SPAN",{});var p=C(t);a=Ne(p,o),p.forEach(T),r=j(_),_.forEach(T),this.h()},h(){w(e,"class","carta-icon-full svelte-1c77udu"),w(e,"aria-label",s=i[29])},m(u,_){H(u,e,_),l&&W(l,e,null),I(e,n),I(e,t),I(t,a),I(e,r),f=!0,c||(d=[L(e,"click",Se(De(g))),L(e,"keydown",pe)],c=!0)},p(u,_){if(i=u,_[0]&18&&h!==(h=i[28].component)){if(l){S();const p=l;k(p.$$.fragment,1,0,()=>{O(p,1)}),D()}h?(l=G(h,m()),P(l.$$.fragment),b(l.$$.fragment,1),W(l,e,n)):l=null}(!f||_[0]&26)&&o!==(o=i[29]+"")&&Ae(a,o),(!f||_[0]&26&&s!==(s=i[29]))&&w(e,"aria-label",s)},i(u){f||(l&&b(l.$$.fragment,u),f=!0)},o(u){l&&k(l.$$.fragment,u),f=!1},d(u){u&&T(e),l&&O(l),c=!1,ae(d)}}}function Lt(i){let e,l,n,t,o,a,r,s,f,c,d,h,m,g=i[2]=="tabs"&&Ke(i),u=!(i[2]==="tabs"&&i[0]==="preview")&&Je(i),_=i[12]&&i[11]&&Qe(i);return{c(){e=V("div"),l=V("div"),g&&g.c(),n=U(),t=V("div"),a=U(),r=V("div"),u&&u.c(),f=U(),_&&_.c(),c=q(),this.h()},l(p){e=B(p,"DIV",{class:!0,role:!0});var E=C(e);l=B(E,"DIV",{class:!0});var z=C(l);g&&g.l(z),z.forEach(T),n=j(E),t=B(E,"DIV",{class:!0}),C(t).forEach(T),a=j(E),r=B(E,"DIV",{class:!0});var R=C(r);u&&u.l(R),R.forEach(T),E.forEach(T),f=j(p),_&&_.l(p),c=q(),this.h()},h(){w(l,"class","carta-toolbar-left svelte-1c77udu"),w(t,"class","carta-filler svelte-1c77udu"),we(()=>i[17].call(t)),w(r,"class","carta-toolbar-right svelte-1c77udu"),w(e,"class","carta-toolbar svelte-1c77udu"),w(e,"role","toolbar"),we(()=>i[22].call(e))},m(p,E){H(p,e,E),I(e,l),g&&g.m(l,null),I(e,n),I(e,t),o=ke(t,i[17].bind(t)),I(e,a),I(e,r),u&&u.m(r,null),i[21](r),s=ke(e,i[22].bind(e)),i[23](e),H(p,f,E),_&&_.m(p,E),H(p,c,E),d=!0,h||(m=[L(window,"resize",i[13]),L(window,"click",i[14])],h=!0)},p(p,E){p[2]=="tabs"?g?g.p(p,E):(g=Ke(p),g.c(),g.m(l,null)):g&&(g.d(1),g=null),p[2]==="tabs"&&p[0]==="preview"?u&&(S(),k(u,1,1,()=>{u=null}),D()):u?(u.p(p,E),E[0]&5&&b(u,1)):(u=Je(p),u.c(),b(u,1),u.m(r,null)),p[12]&&p[11]?_?(_.p(p,E),E[0]&6144&&b(_,1)):(_=Qe(p),_.c(),b(_,1),_.m(c.parentNode,c)):_&&(S(),k(_,1,1,()=>{_=null}),D())},i(p){d||(b(u),b(_),d=!0)},o(p){k(u),k(_),d=!1},d(p){p&&(T(e),T(f),T(c)),g&&g.d(),o(),u&&u.d(),i[21](null),s(),i[23](null),_&&_.d(p),h=!1,ae(m)}}}const Mt=8;function Rt(i,e,l){let{carta:n}=e,{mode:t}=e,{tab:o}=e,{labels:a}=e,r,s,f,c=[...n.icons],d=0,h=0,m=0,g=!1,u=!1;const _=()=>new Promise(requestAnimationFrame),p=ge(async()=>{if(!r||!f)return;const v=()=>r.scrollWidth-r.clientWidth>0;for(;v();)c.pop(),l(4,c),await _();const M=()=>d>2*h+Mt;for(;c.length<n.icons.length&&M();)c.push(n.icons[c.length]),l(4,c),await _()},100);function E(v){const M=v.target;s&&!s.contains(M)&&l(12,u=!1)}x(p);const z=()=>l(0,o="write"),R=()=>l(0,o="preview");function Z(){d=this.clientWidth,l(8,d)}function ie(){h=this.clientWidth,l(9,h)}const se=v=>{var M,X;n.input&&v.action(n.input),(M=n.input)==null||M.update(),(X=n.input)==null||X.textarea.focus()},oe=()=>l(12,u=!u);function $(v){F[v?"unshift":"push"](()=>{f=v,l(7,f)})}function fe(){m=this.clientHeight,l(10,m)}function ce(v){F[v?"unshift":"push"](()=>{r=v,l(5,r)})}const ue=v=>!c.includes(v),he=v=>{var M,X;n.input&&v.action(n.input),(M=n.input)==null||M.update(),(X=n.input)==null||X.textarea.focus(),l(12,u=!1)};function N(v){F[v?"unshift":"push"](()=>{s=v,l(6,s)})}return i.$$set=v=>{"carta"in v&&l(1,n=v.carta),"mode"in v&&l(2,t=v.mode),"tab"in v&&l(0,o=v.tab),"labels"in v&&l(3,a=v.labels)},i.$$.update=()=>{i.$$.dirty[0]&18&&l(11,g=c.length!==n.icons.length)},[o,n,t,a,c,r,s,f,d,h,m,g,u,p,E,z,R,Z,ie,se,oe,$,fe,ce,ue,he,N]}class Pt extends le{constructor(e){super(),ne(this,e,Rt,Lt,te,{carta:1,mode:2,tab:0,labels:3},null,[-1,-1])}}function ye(i,e,l){const n=i.slice();return n[32]=e[l].component,n[33]=e[l].props,n}function $e(i,e,l){const n=i.slice();return n[32]=e[l].component,n[33]=e[l].props,n}function xe(i,e,l){const n=i.slice();return n[32]=e[l].component,n[33]=e[l].props,n}function et(i){let e,l,n;function t(a){i[20](a)}let o={carta:i[2],labels:i[15],mode:i[9]};return i[0]!==void 0&&(o.tab=i[0]),e=new Pt({props:o}),F.push(()=>me(e,"tab",t)),{c(){P(e.$$.fragment)},l(a){K(e.$$.fragment,a)},m(a,r){W(e,a,r),n=!0},p(a,r){const s={};r[0]&4&&(s.carta=a[2]),r[0]&512&&(s.mode=a[9]),!l&&r[0]&1&&(l=!0,s.tab=a[0],de(()=>l=!1)),e.$set(s)},i(a){n||(b(e.$$.fragment,a),n=!0)},o(a){k(e.$$.fragment,a),n=!1},d(a){O(e,a)}}}function tt(i){let e,l,n,t,o;function a(c){i[21](c)}function r(c){i[22](c)}function s(c){i[23](c)}let f={carta:i[2],placeholder:i[6],props:i[7],$$slots:{default:[Wt]},$$scope:{ctx:i}};return i[1]!==void 0&&(f.value=i[1]),i[10]!==void 0&&(f.resize=i[10]),i[11]!==void 0&&(f.elem=i[11]),e=new Dt({props:f}),F.push(()=>me(e,"value",a)),F.push(()=>me(e,"resize",r)),F.push(()=>me(e,"elem",s)),e.$on("scroll",i[16]),{c(){P(e.$$.fragment)},l(c){K(e.$$.fragment,c)},m(c,d){W(e,c,d),o=!0},p(c,d){const h={};d[0]&4&&(h.carta=c[2]),d[0]&64&&(h.placeholder=c[6]),d[0]&128&&(h.props=c[7]),d[0]&8196|d[1]&512&&(h.$$scope={dirty:d,ctx:c}),!l&&d[0]&2&&(l=!0,h.value=c[1],de(()=>l=!1)),!n&&d[0]&1024&&(n=!0,h.resize=c[10],de(()=>n=!1)),!t&&d[0]&2048&&(t=!0,h.elem=c[11],de(()=>t=!1)),e.$set(h)},i(c){o||(b(e.$$.fragment,c),o=!0)},o(c){k(e.$$.fragment,c),o=!1},d(c){O(e,c)}}}function lt(i){let e,l,n=Q(i[2].components.filter(ft)),t=[];for(let a=0;a<n.length;a+=1)t[a]=nt(xe(i,n,a));const o=a=>k(t[a],1,1,()=>{t[a]=null});return{c(){for(let a=0;a<t.length;a+=1)t[a].c();e=q()},l(a){for(let r=0;r<t.length;r+=1)t[r].l(a);e=q()},m(a,r){for(let s=0;s<t.length;s+=1)t[s]&&t[s].m(a,r);H(a,e,r),l=!0},p(a,r){if(r[0]&4){n=Q(a[2].components.filter(ft));let s;for(s=0;s<n.length;s+=1){const f=xe(a,n,s);t[s]?(t[s].p(f,r),b(t[s],1)):(t[s]=nt(f),t[s].c(),b(t[s],1),t[s].m(e.parentNode,e))}for(S(),s=n.length;s<t.length;s+=1)o(s);D()}},i(a){if(!l){for(let r=0;r<n.length;r+=1)b(t[r]);l=!0}},o(a){t=t.filter(Boolean);for(let r=0;r<t.length;r+=1)k(t[r]);l=!1},d(a){a&&T(e),be(t,a)}}}function nt(i){let e,l,n;const t=[{carta:i[2]},i[33]];var o=i[32];function a(r,s){let f={};for(let c=0;c<t.length;c+=1)f=Y(f,t[c]);return s!==void 0&&s[0]&4&&(f=Y(f,y(t,[{carta:r[2]},ee(r[33])]))),{props:f}}return o&&(e=G(o,a(i))),{c(){e&&P(e.$$.fragment),l=q()},l(r){e&&K(e.$$.fragment,r),l=q()},m(r,s){e&&W(e,r,s),H(r,l,s),n=!0},p(r,s){if(s[0]&4&&o!==(o=r[32])){if(e){S();const f=e;k(f.$$.fragment,1,0,()=>{O(f,1)}),D()}o?(e=G(o,a(r,s)),P(e.$$.fragment),b(e.$$.fragment,1),W(e,l.parentNode,l)):e=null}else if(o){const f=s[0]&4?y(t,[{carta:r[2]},ee(r[33])]):{};e.$set(f)}},i(r){n||(e&&b(e.$$.fragment,r),n=!0)},o(r){e&&k(e.$$.fragment,r),n=!1},d(r){r&&T(l),e&&O(e,r)}}}function Wt(i){let e,l,n=i[13]&&lt(i);return{c(){n&&n.c(),e=q()},l(t){n&&n.l(t),e=q()},m(t,o){n&&n.m(t,o),H(t,e,o),l=!0},p(t,o){t[13]?n?(n.p(t,o),o[0]&8192&&b(n,1)):(n=lt(t),n.c(),b(n,1),n.m(e.parentNode,e)):n&&(S(),k(n,1,1,()=>{n=null}),D())},i(t){l||(b(n),l=!0)},o(t){k(n),l=!1},d(t){t&&T(e),n&&n.d(t)}}}function it(i){let e,l,n;function t(a){i[24](a)}let o={carta:i[2],value:i[1],$$slots:{default:[Ot]},$$scope:{ctx:i}};return i[12]!==void 0&&(o.elem=i[12]),e=new Ht({props:o}),F.push(()=>me(e,"elem",t)),e.$on("scroll",i[16]),e.$on("render",i[25]),{c(){P(e.$$.fragment)},l(a){K(e.$$.fragment,a)},m(a,r){W(e,a,r),n=!0},p(a,r){const s={};r[0]&4&&(s.carta=a[2]),r[0]&2&&(s.value=a[1]),r[0]&8196|r[1]&512&&(s.$$scope={dirty:r,ctx:a}),!l&&r[0]&4096&&(l=!0,s.elem=a[12],de(()=>l=!1)),e.$set(s)},i(a){n||(b(e.$$.fragment,a),n=!0)},o(a){k(e.$$.fragment,a),n=!1},d(a){O(e,a)}}}function rt(i){let e,l,n=Q(i[2].components.filter(ct)),t=[];for(let a=0;a<n.length;a+=1)t[a]=at($e(i,n,a));const o=a=>k(t[a],1,1,()=>{t[a]=null});return{c(){for(let a=0;a<t.length;a+=1)t[a].c();e=q()},l(a){for(let r=0;r<t.length;r+=1)t[r].l(a);e=q()},m(a,r){for(let s=0;s<t.length;s+=1)t[s]&&t[s].m(a,r);H(a,e,r),l=!0},p(a,r){if(r[0]&4){n=Q(a[2].components.filter(ct));let s;for(s=0;s<n.length;s+=1){const f=$e(a,n,s);t[s]?(t[s].p(f,r),b(t[s],1)):(t[s]=at(f),t[s].c(),b(t[s],1),t[s].m(e.parentNode,e))}for(S(),s=n.length;s<t.length;s+=1)o(s);D()}},i(a){if(!l){for(let r=0;r<n.length;r+=1)b(t[r]);l=!0}},o(a){t=t.filter(Boolean);for(let r=0;r<t.length;r+=1)k(t[r]);l=!1},d(a){a&&T(e),be(t,a)}}}function at(i){let e,l,n;const t=[{carta:i[2]},i[33]];var o=i[32];function a(r,s){let f={};for(let c=0;c<t.length;c+=1)f=Y(f,t[c]);return s!==void 0&&s[0]&4&&(f=Y(f,y(t,[{carta:r[2]},ee(r[33])]))),{props:f}}return o&&(e=G(o,a(i))),{c(){e&&P(e.$$.fragment),l=q()},l(r){e&&K(e.$$.fragment,r),l=q()},m(r,s){e&&W(e,r,s),H(r,l,s),n=!0},p(r,s){if(s[0]&4&&o!==(o=r[32])){if(e){S();const f=e;k(f.$$.fragment,1,0,()=>{O(f,1)}),D()}o?(e=G(o,a(r,s)),P(e.$$.fragment),b(e.$$.fragment,1),W(e,l.parentNode,l)):e=null}else if(o){const f=s[0]&4?y(t,[{carta:r[2]},ee(r[33])]):{};e.$set(f)}},i(r){n||(e&&b(e.$$.fragment,r),n=!0)},o(r){e&&k(e.$$.fragment,r),n=!1},d(r){r&&T(l),e&&O(e,r)}}}function Ot(i){let e,l,n=i[13]&&rt(i);return{c(){n&&n.c(),e=q()},l(t){n&&n.l(t),e=q()},m(t,o){n&&n.m(t,o),H(t,e,o),l=!0},p(t,o){t[13]?n?(n.p(t,o),o[0]&8192&&b(n,1)):(n=rt(t),n.c(),b(n,1),n.m(e.parentNode,e)):n&&(S(),k(n,1,1,()=>{n=null}),D())},i(t){l||(b(n),l=!0)},o(t){k(n),l=!1},d(t){t&&T(e),n&&n.d(t)}}}function st(i){let e,l,n=Q(i[2].components.filter(ut)),t=[];for(let a=0;a<n.length;a+=1)t[a]=ot(ye(i,n,a));const o=a=>k(t[a],1,1,()=>{t[a]=null});return{c(){for(let a=0;a<t.length;a+=1)t[a].c();e=q()},l(a){for(let r=0;r<t.length;r+=1)t[r].l(a);e=q()},m(a,r){for(let s=0;s<t.length;s+=1)t[s]&&t[s].m(a,r);H(a,e,r),l=!0},p(a,r){if(r[0]&4){n=Q(a[2].components.filter(ut));let s;for(s=0;s<n.length;s+=1){const f=ye(a,n,s);t[s]?(t[s].p(f,r),b(t[s],1)):(t[s]=ot(f),t[s].c(),b(t[s],1),t[s].m(e.parentNode,e))}for(S(),s=n.length;s<t.length;s+=1)o(s);D()}},i(a){if(!l){for(let r=0;r<n.length;r+=1)b(t[r]);l=!0}},o(a){t=t.filter(Boolean);for(let r=0;r<t.length;r+=1)k(t[r]);l=!1},d(a){a&&T(e),be(t,a)}}}function ot(i){let e,l,n;const t=[{carta:i[2]},i[33]];var o=i[32];function a(r,s){let f={};for(let c=0;c<t.length;c+=1)f=Y(f,t[c]);return s!==void 0&&s[0]&4&&(f=Y(f,y(t,[{carta:r[2]},ee(r[33])]))),{props:f}}return o&&(e=G(o,a(i))),{c(){e&&P(e.$$.fragment),l=q()},l(r){e&&K(e.$$.fragment,r),l=q()},m(r,s){e&&W(e,r,s),H(r,l,s),n=!0},p(r,s){if(s[0]&4&&o!==(o=r[32])){if(e){S();const f=e;k(f.$$.fragment,1,0,()=>{O(f,1)}),D()}o?(e=G(o,a(r,s)),P(e.$$.fragment),b(e.$$.fragment,1),W(e,l.parentNode,l)):e=null}else if(o){const f=s[0]&4?y(t,[{carta:r[2]},ee(r[33])]):{};e.$set(f)}},i(r){n||(e&&b(e.$$.fragment,r),n=!0)},o(r){e&&k(e.$$.fragment,r),n=!1},d(r){r&&T(l),e&&O(e,r)}}}function Ft(i){let e,l,n,t,o,a,r,s,f,c,d=!i[5]&&et(i),h=(i[9]=="split"||i[0]=="write")&&tt(i),m=(i[9]=="split"||i[0]=="preview")&&it(i),g=i[13]&&st(i);return{c(){e=V("div"),d&&d.c(),l=U(),n=V("div"),t=V("div"),h&&h.c(),o=U(),m&&m.c(),r=U(),g&&g.c(),this.h()},l(u){e=B(u,"DIV",{class:!0});var _=C(e);d&&d.l(_),l=j(_),n=B(_,"DIV",{class:!0});var p=C(n);t=B(p,"DIV",{class:!0});var E=C(t);h&&h.l(E),o=j(E),m&&m.l(E),E.forEach(T),p.forEach(T),r=j(_),g&&g.l(_),_.forEach(T),this.h()},h(){w(t,"class",a="carta-container mode-"+i[9]+" svelte-11jlii3"),w(n,"class","carta-wrapper"),w(e,"class",s="carta-editor carta-theme__"+i[3]+" svelte-11jlii3"),we(()=>i[27].call(e))},m(u,_){H(u,e,_),d&&d.m(e,null),I(e,l),I(e,n),I(n,t),h&&h.m(t,null),I(t,o),m&&m.m(t,null),I(e,r),g&&g.m(e,null),i[26](e),f=ke(e,i[27].bind(e)),c=!0},p(u,_){u[5]?d&&(S(),k(d,1,1,()=>{d=null}),D()):d?(d.p(u,_),_[0]&32&&b(d,1)):(d=et(u),d.c(),b(d,1),d.m(e,l)),u[9]=="split"||u[0]=="write"?h?(h.p(u,_),_[0]&513&&b(h,1)):(h=tt(u),h.c(),b(h,1),h.m(t,o)):h&&(S(),k(h,1,1,()=>{h=null}),D()),u[9]=="split"||u[0]=="preview"?m?(m.p(u,_),_[0]&513&&b(m,1)):(m=it(u),m.c(),b(m,1),m.m(t,null)):m&&(S(),k(m,1,1,()=>{m=null}),D()),(!c||_[0]&512&&a!==(a="carta-container mode-"+u[9]+" svelte-11jlii3"))&&w(t,"class",a),u[13]?g?(g.p(u,_),_[0]&8192&&b(g,1)):(g=st(u),g.c(),b(g,1),g.m(e,null)):g&&(S(),k(g,1,1,()=>{g=null}),D()),(!c||_[0]&8&&s!==(s="carta-editor carta-theme__"+u[3]+" svelte-11jlii3"))&&w(e,"class",s)},i(u){c||(b(d),b(h),b(m),b(g),c=!0)},o(u){k(d),k(h),k(m),k(g),c=!1},d(u){u&&T(e),d&&d.d(),h&&h.d(),m&&m.d(),g&&g.d(),i[26](null),f()}}}function Ut(i){const e=i.scrollHeight-i.clientHeight;return i.scrollTop*(1+i.clientHeight/e)/i.scrollHeight}const ft=({parent:i})=>[i].flat().includes("input"),ct=({parent:i})=>[i].flat().includes("renderer"),ut=({parent:i})=>[i].flat().includes("editor");function jt(i,e,l){let{carta:n}=e,{theme:t="default"}=e,{value:o=""}=e,{mode:a="auto"}=e,{scroll:r="sync"}=e,{disableToolbar:s=!1}=e,{placeholder:f=""}=e,{textarea:c={}}=e,{selectedTab:d="write"}=e,{labels:h={}}=e;const m={...Vt,...h};let g,u,_=!1,p,E,z,R,Z,ie=0;const se=ge(()=>{Z=null},1e3);function oe(A){const[J,re]=A.target==z?[z,R]:[R,z];u=="split"&&r=="sync"&&$(J,re)}function $(A,J){const re=Ut(A);if(ie=re,Z&&Z!=A)return;Z=A;const mt=J.scrollHeight-J.clientHeight;J.scrollTo({top:re*mt,behavior:"smooth"}),se()}function fe(A){if(u!=="tabs")return;const J=A==="write"?z:R;if(!J)return;const re=J.scrollHeight-J.clientHeight;J.scroll({top:re*ie,behavior:"instant"})}x(()=>n.$setElement(E)),x(()=>l(13,_=!0));function ce(A){d=A,l(0,d)}function ue(A){o=A,l(1,o)}function he(A){p=A,l(10,p)}function N(A){z=A,l(11,z)}function v(A){R=A,l(12,R)}const M=()=>{u=="split"&&r=="sync"&&$(z,R)};function X(A){F[A?"unshift":"push"](()=>{E=A,l(14,E)})}function _e(){g=this.clientWidth,l(8,g)}return i.$$set=A=>{"carta"in A&&l(2,n=A.carta),"theme"in A&&l(3,t=A.theme),"value"in A&&l(1,o=A.value),"mode"in A&&l(18,a=A.mode),"scroll"in A&&l(4,r=A.scroll),"disableToolbar"in A&&l(5,s=A.disableToolbar),"placeholder"in A&&l(6,f=A.placeholder),"textarea"in A&&l(7,c=A.textarea),"selectedTab"in A&&l(0,d=A.selectedTab),"labels"in A&&l(19,h=A.labels)},i.$$.update=()=>{i.$$.dirty[0]&262400&&l(9,u=a==="auto"?g>768?"split":"tabs":a),i.$$.dirty[0]&1536&&p&&p(),i.$$.dirty[0]&6145&&fe(d)},[d,o,n,t,r,s,f,c,g,u,p,z,R,_,E,m,oe,$,a,h,ce,ue,he,N,v,M,X,_e]}class $t extends le{constructor(e){super(),ne(this,e,jt,Ft,te,{carta:2,theme:3,value:1,mode:18,scroll:4,disableToolbar:5,placeholder:6,textarea:7,selectedTab:0,labels:19},null,[-1,-1])}}function Xt(i){let e,l;return e=new Et({props:{class:"h-5 w-5"}}),{c(){P(e.$$.fragment)},l(n){K(e.$$.fragment,n)},m(n,t){W(e,n,t),l=!0},p:ve,i(n){l||(b(e.$$.fragment,n),l=!0)},o(n){k(e.$$.fragment,n),l=!1},d(n){O(e,n)}}}class xt extends le{constructor(e){super(),ne(this,e,null,Xt,te,{})}}export{xt as A,$t as M};
