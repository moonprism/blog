import{s as K,L as W,e as q,a as F,c as B,b as U,g as O,f as v,m as j,i as D,h as L,O as ee,k as z,P as H,N as Te,t as E,d as M,j as oe,B as ce,U as Re,n as J,x as Ee,Q as Me,o as Le,l as Ve,aa as Pe,ac as ze,J as je}from"../chunks/scheduler.CjD0NZ8o.js";import{S as Y,i as X,f as te,c as _,a as g,m as d,t as c,b as p,d as h,g as ke,e as Ce}from"../chunks/index.or77jMZA.js";import{j as Ne,k as qe,P as Be,C as Ae,l as He,m as Ue,n as Ge,o as Qe,p as Je,c as re,R as Ke,T as We,a as Ye,b as me,d as Xe,e as Ze,f as ye,g as xe,h as et,i as tt,D as nt}from"../chunks/date.DCN9688O.js";import{C as rt,D as pe,T as ot}from"../chunks/table.Bbh0fYby.js";import{B as se,c as _e}from"../chunks/index.Cy1MAQMk.js";import{d as st}from"../chunks/system.DiN6QSf-.js";import{i as at,f as x,I as lt,L as ft,a as ut}from"../chunks/fetch.UBuoTI4j.js";import{w as Z}from"../chunks/index.BLWFmFSp.js";import{C as it}from"../chunks/table.C6UAzvKa.js";import{g as $t,a as ct}from"../chunks/spread.CgU5AtxT.js";import{z as ne,R as mt,s as pt,a as _t,d as gt,b as dt,D as ht,c as wt,F as le,e as bt,C as ie,f as $e,g as De,h as vt}from"../chunks/zod.Dn88xz39.js";import"../chunks/entry.B9pAj-Df.js";import{M as kt,A as Ct}from"../chunks/attachment-icon.G75shQak.js";import{F as Dt,C as It,c as St,m as Ft,r as Ot}from"../chunks/data.BGGdSGzA.js";import{e as ge}from"../chunks/each.BFF0fzu3.js";import{D as Tt,a as Rt}from"../chunks/DotsHorizontal.pqD2To3M.js";import{a as Et}from"../chunks/alert.PramJLuh.js";import{R as Mt,a as Lt}from"../chunks/row-title.DX9KX_JJ.js";import{a as Vt}from"../chunks/Toaster.svelte_svelte_type_style_lang.DUqWXG7W.js";const de=ne.object({lang:ne.string().min(1).max(10),title:ne.string().min(1).max(100),content:ne.string()});function Pt(a){let e,r,t,o,n,s;function l(u){a[5](u)}let f={carta:a[3],theme:"gist",mode:"tabs"};return a[0]!==void 0&&(f.value=a[0]),r=new kt({props:f}),W.push(()=>te(r,"value",l)),n=new Dt({props:{open:a[2],callback:a[6]}}),{c(){e=q("div"),_(r.$$.fragment),o=F(),_(n.$$.fragment),this.h()},l(u){e=B(u,"DIV",{class:!0});var i=U(e);g(r.$$.fragment,i),o=O(i),g(n.$$.fragment,i),i.forEach(v),this.h()},h(){j(e,"class","gist")},m(u,i){D(u,e,i),d(r,e,null),L(e,o),d(n,e,null),s=!0},p(u,[i]){const $={};!t&&i&1&&(t=!0,$.value=u[0],ee(()=>t=!1)),r.$set($);const m={};i&2&&(m.callback=u[6]),n.$set(m)},i(u){s||(c(r.$$.fragment,u),c(n.$$.fragment,u),s=!0)},o(u){p(r.$$.fragment,u),p(n.$$.fragment,u),s=!1},d(u){u&&v(e),h(r),h(n)}}}function zt(a,e,r){let t,o=Z(!1);z(a,o,k=>r(7,t=k));let n;const l={icons:[{id:"attachment",action:k=>{r(1,n=k),H(o,t=!0,t)},component:Ct}],transformers:Ft},f=new It({disableIcons:!0,sanitizer:!1,extensions:[St(),l]});let{value:u=""}=e;const i=async()=>await f.render(u);function $(k){u=k,r(0,u)}const m=k=>{n.insertAt(n.getSelection().start,`![](${k.key})`),n.update()};return a.$$set=k=>{"value"in k&&r(0,u=k.value)},[u,n,o,f,i,$,m]}class jt extends Y{constructor(e){super(),X(this,e,zt,Pt,K,{value:0,render:4})}get render(){return this.$$.ctx[4]}}function he(a,e,r){const t=a.slice();return t[8]=e[r],t}function Nt(a){let e,r,t,o;return t=new rt({props:{class:"ml-2 h-4 w-4 shrink-0 opacity-50"}}),{c(){e=E(a[3]),r=F(),_(t.$$.fragment)},l(n){e=M(n,a[3]),r=O(n),g(t.$$.fragment,n)},m(n,s){D(n,e,s),D(n,r,s),d(t,n,s),o=!0},p(n,s){(!o||s&8)&&oe(e,n[3])},i(n){o||(c(t.$$.fragment,n),o=!0)},o(n){p(t.$$.fragment,n),o=!1},d(n){n&&(v(e),v(r)),h(t,n)}}}function qt(a){let e,r;return e=new se({props:{builders:[a[11]],variant:"outline",role:"combobox","aria-expanded":a[2],class:"w-[150px] justify-between",$$slots:{default:[Nt]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},p(t,o){const n={};o&2048&&(n.builders=[t[11]]),o&4&&(n["aria-expanded"]=t[2]),o&4104&&(n.$$scope={dirty:o,ctx:t}),e.$set(n)},i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function Bt(a){let e;return{c(){e=E("No item found.")},l(r){e=M(r,"No item found.")},m(r,t){D(r,e,t)},d(r){r&&v(e)}}}function At(a){let e,r,t=a[8].label+"",o,n,s;return e=new Je({props:{class:_e("mr-2 h-4 w-4",a[0]!==a[8].value&&"text-transparent")}}),{c(){_(e.$$.fragment),r=F(),o=E(t),n=F()},l(l){g(e.$$.fragment,l),r=O(l),o=M(l,t),n=O(l)},m(l,f){d(e,l,f),D(l,r,f),D(l,o,f),D(l,n,f),s=!0},p(l,f){const u={};f&3&&(u.class=_e("mr-2 h-4 w-4",l[0]!==l[8].value&&"text-transparent")),e.$set(u),(!s||f&2)&&t!==(t=l[8].label+"")&&oe(o,t)},i(l){s||(c(e.$$.fragment,l),s=!0)},o(l){p(e.$$.fragment,l),s=!1},d(l){l&&(v(r),v(o),v(n)),h(e,l)}}}function we(a){let e,r;function t(...o){return a[5](a[7],...o)}return e=new Qe({props:{class:"cursor-pointer",value:a[8].value,onSelect:t,$$slots:{default:[At]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment)},l(o){g(e.$$.fragment,o)},m(o,n){d(e,o,n),r=!0},p(o,n){a=o;const s={};n&2&&(s.value=a[8].value),n&129&&(s.onSelect=t),n&4099&&(s.$$scope={dirty:n,ctx:a}),e.$set(s)},i(o){r||(c(e.$$.fragment,o),r=!0)},o(o){p(e.$$.fragment,o),r=!1},d(o){h(e,o)}}}function Ht(a){let e,r,t=ge(a[1]),o=[];for(let s=0;s<t.length;s+=1)o[s]=we(he(a,t,s));const n=s=>p(o[s],1,1,()=>{o[s]=null});return{c(){for(let s=0;s<o.length;s+=1)o[s].c();e=ce()},l(s){for(let l=0;l<o.length;l+=1)o[l].l(s);e=ce()},m(s,l){for(let f=0;f<o.length;f+=1)o[f]&&o[f].m(s,l);D(s,e,l),r=!0},p(s,l){if(l&147){t=ge(s[1]);let f;for(f=0;f<t.length;f+=1){const u=he(s,t,f);o[f]?(o[f].p(u,l),c(o[f],1)):(o[f]=we(u),o[f].c(),c(o[f],1),o[f].m(e.parentNode,e))}for(ke(),f=t.length;f<o.length;f+=1)n(f);Ce()}},i(s){if(!r){for(let l=0;l<t.length;l+=1)c(o[l]);r=!0}},o(s){o=o.filter(Boolean);for(let l=0;l<o.length;l+=1)p(o[l]);r=!1},d(s){s&&v(e),Re(o,s)}}}function Ut(a){let e,r,t,o,n,s;return e=new He({props:{placeholder:"Search ...",class:"h-9"}}),t=new Ue({props:{$$slots:{default:[Bt]},$$scope:{ctx:a}}}),n=new Ge({props:{$$slots:{default:[Ht]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment),r=F(),_(t.$$.fragment),o=F(),_(n.$$.fragment)},l(l){g(e.$$.fragment,l),r=O(l),g(t.$$.fragment,l),o=O(l),g(n.$$.fragment,l)},m(l,f){d(e,l,f),D(l,r,f),d(t,l,f),D(l,o,f),d(n,l,f),s=!0},p(l,f){const u={};f&4096&&(u.$$scope={dirty:f,ctx:l}),t.$set(u);const i={};f&4227&&(i.$$scope={dirty:f,ctx:l}),n.$set(i)},i(l){s||(c(e.$$.fragment,l),c(t.$$.fragment,l),c(n.$$.fragment,l),s=!0)},o(l){p(e.$$.fragment,l),p(t.$$.fragment,l),p(n.$$.fragment,l),s=!1},d(l){l&&(v(r),v(o)),h(e,l),h(t,l),h(n,l)}}}function Gt(a){let e,r;return e=new Ae({props:{$$slots:{default:[Ut]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},p(t,o){const n={};o&4227&&(n.$$scope={dirty:o,ctx:t}),e.$set(n)},i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function Qt(a){let e,r,t,o;return e=new qe({props:{asChild:!0,$$slots:{default:[qt,({builder:n})=>({11:n}),({builder:n})=>n?2048:0]},$$scope:{ctx:a}}}),t=new Be({props:{class:"w-[180px] p-0",$$slots:{default:[Gt]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment),r=F(),_(t.$$.fragment)},l(n){g(e.$$.fragment,n),r=O(n),g(t.$$.fragment,n)},m(n,s){d(e,n,s),D(n,r,s),d(t,n,s),o=!0},p(n,s){const l={};s&6156&&(l.$$scope={dirty:s,ctx:n}),e.$set(l);const f={};s&4227&&(f.$$scope={dirty:s,ctx:n}),t.$set(f)},i(n){o||(c(e.$$.fragment,n),c(t.$$.fragment,n),o=!0)},o(n){p(e.$$.fragment,n),p(t.$$.fragment,n),o=!1},d(n){n&&v(r),h(e,n),h(t,n)}}}function Jt(a){let e,r,t;function o(s){a[6](s)}let n={$$slots:{default:[Qt,({ids:s})=>({7:s}),({ids:s})=>s?128:0]},$$scope:{ctx:a}};return a[2]!==void 0&&(n.open=a[2]),e=new Ne({props:n}),W.push(()=>te(e,"open",o)),{c(){_(e.$$.fragment)},l(s){g(e.$$.fragment,s)},m(s,l){d(e,s,l),t=!0},p(s,[l]){const f={};l&4239&&(f.$$scope={dirty:l,ctx:s}),!r&&l&4&&(r=!0,f.open=s[2],ee(()=>r=!1)),e.$set(f)},i(s){t||(c(e.$$.fragment,s),t=!0)},o(s){p(e.$$.fragment,s),t=!1},d(s){h(e,s)}}}function Kt(a,e,r){let t,{items:o}=e,{value:n}=e,s=!1;function l(i){r(2,s=!1),Te().then(()=>{var $;($=document.getElementById(i))==null||$.focus()})}const f=(i,$)=>{r(0,n=$),l(i.trigger)};function u(i){s=i,r(2,s)}return a.$$set=i=>{"items"in i&&r(1,o=i.items),"value"in i&&r(0,n=i.value)},a.$$.update=()=>{var i;a.$$.dirty&3&&r(3,t=((i=o.find($=>$.value===n))==null?void 0:i.label)??"Select ...")},[n,o,s,t,l,f,u]}class Wt extends Y{constructor(e){super(),X(this,e,Kt,Jt,K,{items:1,value:0})}}function Yt(a){let e=a[9]?"New":"Edit",r,t;return{c(){r=E(e),t=E(" Gist")},l(o){r=M(o,e),t=M(o," Gist")},m(o,n){D(o,r,n),D(o,t,n)},p:J,d(o){o&&(v(r),v(t))}}}function Xt(a){let e,r;return e=new bt({props:{$$slots:{default:[Yt]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},p(t,o){const n={};o&1048576&&(n.$$scope={dirty:o,ctx:t}),e.$set(n)},i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function Zt(a){let e;return{c(){e=E("标题")},l(r){e=M(r,"标题")},m(r,t){D(r,e,t)},d(r){r&&v(e)}}}function yt(a){let e,r,t,o,n;e=new De({props:{$$slots:{default:[Zt]},$$scope:{ctx:a}}});const s=[a[19],{autocomplete:"off"}];function l(u){a[11](u)}let f={};for(let u=0;u<s.length;u+=1)f=Le(f,s[u]);return a[3].title!==void 0&&(f.value=a[3].title),t=new lt({props:f}),W.push(()=>te(t,"value",l)),{c(){_(e.$$.fragment),r=F(),_(t.$$.fragment)},l(u){g(e.$$.fragment,u),r=O(u),g(t.$$.fragment,u)},m(u,i){d(e,u,i),D(u,r,i),d(t,u,i),n=!0},p(u,i){const $={};i&1048576&&($.$$scope={dirty:i,ctx:u}),e.$set($);const m=i&524288?$t(s,[ct(u[19]),s[1]]):{};!o&&i&8&&(o=!0,m.value=u[3].title,ee(()=>o=!1)),t.$set(m)},i(u){n||(c(e.$$.fragment,u),c(t.$$.fragment,u),n=!0)},o(u){p(e.$$.fragment,u),p(t.$$.fragment,u),n=!1},d(u){u&&v(r),h(e,u),h(t,u)}}}function xt(a){let e,r,t,o;return e=new ie({props:{$$slots:{default:[yt,({attrs:n})=>({19:n}),({attrs:n})=>n?524288:0]},$$scope:{ctx:a}}}),t=new $e({}),{c(){_(e.$$.fragment),r=F(),_(t.$$.fragment)},l(n){g(e.$$.fragment,n),r=O(n),g(t.$$.fragment,n)},m(n,s){d(e,n,s),D(n,r,s),d(t,n,s),o=!0},p(n,s){const l={};s&1572872&&(l.$$scope={dirty:s,ctx:n}),e.$set(l)},i(n){o||(c(e.$$.fragment,n),c(t.$$.fragment,n),o=!0)},o(n){p(e.$$.fragment,n),p(t.$$.fragment,n),o=!1},d(n){n&&v(r),h(e,n),h(t,n)}}}function en(a){let e;return{c(){e=E("语言")},l(r){e=M(r,"语言")},m(r,t){D(r,e,t)},d(r){r&&v(e)}}}function tn(a){let e,r,t,o,n;e=new De({props:{$$slots:{default:[en]},$$scope:{ctx:a}}});function s(f){a[12](f)}let l={items:Oe};return a[1]!==void 0&&(l.value=a[1]),t=new Wt({props:l}),W.push(()=>te(t,"value",s)),{c(){_(e.$$.fragment),r=F(),_(t.$$.fragment)},l(f){g(e.$$.fragment,f),r=O(f),g(t.$$.fragment,f)},m(f,u){d(e,f,u),D(f,r,u),d(t,f,u),n=!0},p(f,u){const i={};u&1048576&&(i.$$scope={dirty:u,ctx:f}),e.$set(i);const $={};!o&&u&2&&(o=!0,$.value=f[1],ee(()=>o=!1)),t.$set($)},i(f){n||(c(e.$$.fragment,f),c(t.$$.fragment,f),n=!0)},o(f){p(e.$$.fragment,f),p(t.$$.fragment,f),n=!1},d(f){f&&v(r),h(e,f),h(t,f)}}}function nn(a){let e,r,t,o;return e=new ie({props:{$$slots:{default:[tn,({attrs:n})=>({19:n}),({attrs:n})=>n?524288:0]},$$scope:{ctx:a}}}),t=new $e({}),{c(){_(e.$$.fragment),r=F(),_(t.$$.fragment)},l(n){g(e.$$.fragment,n),r=O(n),g(t.$$.fragment,n)},m(n,s){d(e,n,s),D(n,r,s),d(t,n,s),o=!0},p(n,s){const l={};s&1048578&&(l.$$scope={dirty:s,ctx:n}),e.$set(l)},i(n){o||(c(e.$$.fragment,n),c(t.$$.fragment,n),o=!0)},o(n){p(e.$$.fragment,n),p(t.$$.fragment,n),o=!1},d(n){n&&v(r),h(e,n),h(t,n)}}}function rn(a){let e,r,t;function o(s){a[13](s)}let n={};return a[3].content!==void 0&&(n.value=a[3].content),e=new jt({props:n}),W.push(()=>te(e,"value",o)),a[14](e),{c(){_(e.$$.fragment)},l(s){g(e.$$.fragment,s)},m(s,l){d(e,s,l),t=!0},p(s,l){const f={};!r&&l&8&&(r=!0,f.value=s[3].content,ee(()=>r=!1)),e.$set(f)},i(s){t||(c(e.$$.fragment,s),t=!0)},o(s){p(e.$$.fragment,s),t=!1},d(s){a[14](null),h(e,s)}}}function on(a){let e,r,t,o;return e=new ie({props:{$$slots:{default:[rn,({attrs:n})=>({19:n}),({attrs:n})=>n?524288:0]},$$scope:{ctx:a}}}),t=new $e({}),{c(){_(e.$$.fragment),r=F(),_(t.$$.fragment)},l(n){g(e.$$.fragment,n),r=O(n),g(t.$$.fragment,n)},m(n,s){d(e,n,s),D(n,r,s),d(t,n,s),o=!0},p(n,s){const l={};s&1048588&&(l.$$scope={dirty:s,ctx:n}),e.$set(l)},i(n){o||(c(e.$$.fragment,n),c(t.$$.fragment,n),o=!0)},o(n){p(e.$$.fragment,n),p(t.$$.fragment,n),o=!1},d(n){n&&v(r),h(e,n),h(t,n)}}}function sn(a){let e,r;return e=new vt({props:{class:"w-full",$$slots:{default:[ln]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function an(a){let e,r;return e=new se({props:{class:"w-full",$$slots:{default:[fn]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function ln(a){let e;return{c(){e=E("保存")},l(r){e=M(r,"保存")},m(r,t){D(r,e,t)},d(r){r&&v(e)}}}function fn(a){let e,r;return e=new ft({props:{class:"mr-2 h-4 w-4 animate-spin"}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},p:J,i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function un(a){let e,r,t,o,n,s,l,f,u,i,$,m,k,I,T,N,R;t=new wt({props:{$$slots:{default:[Xt]},$$scope:{ctx:a}}}),l=new le({props:{form:a[6],name:"title",class:"w-full p-1",$$slots:{default:[xt]},$$scope:{ctx:a}}}),u=new le({props:{form:a[6],name:"lang",class:"p-1",$$slots:{default:[nn]},$$scope:{ctx:a}}}),$=new le({props:{form:a[6],name:"content",$$slots:{default:[on]},$$scope:{ctx:a}}});const G=[an,sn],V=[];function Q(C,S){return C[5]?0:1}return k=Q(a),I=V[k]=G[k](a),{c(){e=q("input"),r=F(),_(t.$$.fragment),o=F(),n=q("form"),s=q("div"),_(l.$$.fragment),f=F(),_(u.$$.fragment),i=F(),_($.$$.fragment),m=F(),I.c(),this.h()},l(C){e=B(C,"INPUT",{class:!0,type:!0}),r=O(C),g(t.$$.fragment,C),o=O(C),n=B(C,"FORM",{method:!0,class:!0});var S=U(n);s=B(S,"DIV",{class:!0});var P=U(s);g(l.$$.fragment,P),f=O(P),g(u.$$.fragment,P),P.forEach(v),i=O(S),g($.$$.fragment,S),m=O(S),I.l(S),S.forEach(v),this.h()},h(){j(e,"class","fixed left-0 top-0 h-0 w-0"),j(e,"type","checkbox"),e.autofocus=!0,j(s,"class","flex flex-1 space-x-1"),j(n,"method","POST"),j(n,"class","space-y-2 w-full overflow-auto")},m(C,S){D(C,e,S),D(C,r,S),d(t,C,S),D(C,o,S),D(C,n,S),L(n,s),d(l,s,null),L(s,f),d(u,s,null),L(n,i),d($,n,null),L(n,m),V[k].m(n,null),T=!0,e.focus(),N||(R=Me(a[8].call(null,n)),N=!0)},p(C,S){const P={};S&1048576&&(P.$$scope={dirty:S,ctx:C}),t.$set(P);const w={};S&1048584&&(w.$$scope={dirty:S,ctx:C}),l.$set(w);const b={};S&1048578&&(b.$$scope={dirty:S,ctx:C}),u.$set(b);const y={};S&1048588&&(y.$$scope={dirty:S,ctx:C}),$.$set(y);let ae=k;k=Q(C),k!==ae&&(ke(),p(V[ae],1,1,()=>{V[ae]=null}),Ce(),I=V[k],I||(I=V[k]=G[k](C),I.c()),c(I,1),I.m(n,null))},i(C){T||(c(t.$$.fragment,C),c(l.$$.fragment,C),c(u.$$.fragment,C),c($.$$.fragment,C),c(I),T=!0)},o(C){p(t.$$.fragment,C),p(l.$$.fragment,C),p(u.$$.fragment,C),p($.$$.fragment,C),p(I),T=!1},d(C){C&&(v(e),v(r),v(o),v(n)),h(t,C),h(l),h(u),h($),V[k].d(),N=!1,R()}}}function $n(a){let e,r;return e=new ht({props:{class:"sm:max-w-[540px]",$$slots:{default:[un]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},p(t,o){const n={};o&1048622&&(n.$$scope={dirty:o,ctx:t}),e.$set(n)},i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function cn(a){let e,r,t;function o(s){a[15](s)}let n={closeOnEscape:!1,$$slots:{default:[$n]},$$scope:{ctx:a}};return a[4]!==void 0&&(n.open=a[4]),e=new mt({props:n}),W.push(()=>te(e,"open",o)),{c(){_(e.$$.fragment)},l(s){g(e.$$.fragment,s)},m(s,l){d(e,s,l),t=!0},p(s,[l]){const f={};l&1048622&&(f.$$scope={dirty:l,ctx:s}),!r&&l&16&&(r=!0,f.open=s[4],ee(()=>r=!1)),e.$set(f)},i(s){t||(c(e.$$.fragment,s),t=!0)},o(s){p(e.$$.fragment,s),t=!1},d(s){h(e,s)}}}function mn(a,e,r){let t,o,n,s=J,l=()=>(s(),s=Ee(k,w=>r(4,n=w)),k),f;z(a,A,w=>r(16,o=w)),z(a,at,w=>r(5,f=w)),a.$$.on_destroy.push(()=>s());const u=pt(gt(dt(de)),{validators:_t(de),SPA:!0,onUpdate({form:w}){w.valid&&N()},resetForm:!1}),{form:i,enhance:$}=u;z(a,i,w=>r(3,t=w));let{formData:m}=e,{formOpen:k}=e;l();const I=m.id===0;let T;H(i,t=m,t);async function N(){if(t.lang!=="md"){const b=t.content.trim().split(/\r?\n/);b.pop(),b.shift(),H(i,t.content=b.join(`
`),t)}const w={lang:t.lang,title:t.title,content:t.content,html:await T.render()};I?x.post("gist",w).then(b=>{b.ok&&(H(A,o=[b.data,...o],o),be())}):x.put(`gist/${m.id}`,w).then(b=>{b.ok&&(r(10,m=t),r(10,m.html=w.html,m),r(10,m.updated=Date.parse(new Date().toString())/1e3,m),H(A,o[o.findIndex(y=>y.id===m.id)]=m,o),be())})}let R=I?"md":t.lang;function G(){const w=t.content.trim().split(/\r?\n/),b=`\`\`\`${R}`;w[0].startsWith("```")?R!=="md"&&(w[0]=b):R!=="md"&&w.unshift(b),w[w.length-1]!=="```"&&R!=="md"&&w.push("```"),H(i,t.lang=R,t),H(i,t.content=w.join(`
`),t)}function V(w){a.$$.not_equal(t.title,w)&&(t.title=w,i.set(t))}function Q(w){R=w,r(1,R)}function C(w){a.$$.not_equal(t.content,w)&&(t.content=w,i.set(t))}function S(w){W[w?"unshift":"push"](()=>{T=w,r(2,T)})}function P(w){n=w,k.set(n)}return a.$$set=w=>{"formData"in w&&r(10,m=w.formData),"formOpen"in w&&l(r(0,k=w.formOpen))},a.$$.update=()=>{a.$$.dirty&2&&G()},[k,R,T,t,n,f,u,i,$,I,m,V,Q,C,S,P]}class pn extends Y{constructor(e){super(),X(this,e,mn,cn,K,{formData:10,formOpen:0})}}const A=Z([]),_n=Z(0),gn=Z(["content"]),Ie=Z([]);function dn(){x.get("group/gist-lang").then(a=>{a.ok&&Ie.set(a.data.map(e=>({id:e.lang,label:bn(e.lang),icon:re(it,{text:String(e.count)})})))})}dn();const Se=Z("");function hn(a=""){x.get(`gist?q=${a}`).then(e=>{e.ok&&(A.set(e.data.data),_n.set(e.data.pagination.count),Se.set(a))})}function wn(){return{id:0,title:"",lang:"",content:""}}let fe;const ue=Z(!1);function Fe(a){a||(a=wn()),fe&&fe.$destroy(),fe=new pn({target:document.body,props:{formData:a,formOpen:ue}}),ue.set(!0)}function be(){ue.set(!1)}function bn(a=""){var e;return(e=Oe.find(r=>r.value===a))==null?void 0:e.label}const Oe=[{value:"md",label:"Markdown"},{value:"lua",label:"Lua"},{value:"go",label:"Golang"},{value:"js",label:"JavaScript"},{value:"sh",label:"Shell"},{value:"yml",label:"YAML"},{value:"css",label:"CSS"},{value:"php",label:"PHP"},{value:"sql",label:"SQL"},{value:"html",label:"HTML"}];function vn(a){let e,r,t,o="Open Menu",n;return e=new Tt({props:{class:"h-4 w-4"}}),{c(){_(e.$$.fragment),r=F(),t=q("span"),t.textContent=o,this.h()},l(s){g(e.$$.fragment,s),r=O(s),t=B(s,"SPAN",{class:!0,"data-svelte-h":!0}),Ve(t)!=="svelte-1kgmsd2"&&(t.textContent=o),this.h()},h(){j(t,"class","sr-only")},m(s,l){d(e,s,l),D(s,r,l),D(s,t,l),n=!0},p:J,i(s){n||(c(e.$$.fragment,s),n=!0)},o(s){p(e.$$.fragment,s),n=!1},d(s){s&&(v(r),v(t)),h(e,s)}}}function kn(a){let e,r;return e=new se({props:{variant:"ghost",builders:[a[5]],class:"flex h-8 w-8 p-0 data-[state=open]:bg-muted",$$slots:{default:[vn]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},p(t,o){const n={};o&32&&(n.builders=[t[5]]),o&64&&(n.$$scope={dirty:o,ctx:t}),e.$set(n)},i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function Cn(a){let e;return{c(){e=E("编辑")},l(r){e=M(r,"编辑")},m(r,t){D(r,e,t)},d(r){r&&v(e)}}}function Dn(a){let e;return{c(){e=E("⌘⌫")},l(r){e=M(r,"⌘⌫")},m(r,t){D(r,e,t)},d(r){r&&v(e)}}}function In(a){let e,r,t;return r=new Rt({props:{$$slots:{default:[Dn]},$$scope:{ctx:a}}}),{c(){e=E(`删除
      `),_(r.$$.fragment)},l(o){e=M(o,`删除
      `),g(r.$$.fragment,o)},m(o,n){D(o,e,n),d(r,o,n),t=!0},p(o,n){const s={};n&64&&(s.$$scope={dirty:n,ctx:o}),r.$set(s)},i(o){t||(c(r.$$.fragment,o),t=!0)},o(o){p(r.$$.fragment,o),t=!1},d(o){o&&v(e),h(r,o)}}}function Sn(a){let e,r,t,o;return e=new pe({props:{$$slots:{default:[Cn]},$$scope:{ctx:a}}}),e.$on("click",a[1]),t=new pe({props:{$$slots:{default:[In]},$$scope:{ctx:a}}}),t.$on("click",a[3]),{c(){_(e.$$.fragment),r=F(),_(t.$$.fragment)},l(n){g(e.$$.fragment,n),r=O(n),g(t.$$.fragment,n)},m(n,s){d(e,n,s),D(n,r,s),d(t,n,s),o=!0},p(n,s){const l={};s&64&&(l.$$scope={dirty:s,ctx:n}),e.$set(l);const f={};s&64&&(f.$$scope={dirty:s,ctx:n}),t.$set(f)},i(n){o||(c(e.$$.fragment,n),c(t.$$.fragment,n),o=!0)},o(n){p(e.$$.fragment,n),p(t.$$.fragment,n),o=!1},d(n){n&&v(r),h(e,n),h(t,n)}}}function Fn(a){let e,r,t,o;return e=new We({props:{asChild:!0,$$slots:{default:[kn,({builder:n})=>({5:n}),({builder:n})=>n?32:0]},$$scope:{ctx:a}}}),t=new Ye({props:{class:"w-[160px]",align:"end",$$slots:{default:[Sn]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment),r=F(),_(t.$$.fragment)},l(n){g(e.$$.fragment,n),r=O(n),g(t.$$.fragment,n)},m(n,s){d(e,n,s),D(n,r,s),d(t,n,s),o=!0},p(n,s){const l={};s&96&&(l.$$scope={dirty:s,ctx:n}),e.$set(l);const f={};s&65&&(f.$$scope={dirty:s,ctx:n}),t.$set(f)},i(n){o||(c(e.$$.fragment,n),c(t.$$.fragment,n),o=!0)},o(n){p(e.$$.fragment,n),p(t.$$.fragment,n),o=!1},d(n){n&&v(r),h(e,n),h(t,n)}}}function On(a){let e,r;return e=new Ke({props:{$$slots:{default:[Fn]},$$scope:{ctx:a}}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},p(t,[o]){const n={};o&65&&(n.$$scope={dirty:o,ctx:t}),e.$set(n)},i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function Tn(a,e,r){let t;z(a,A,f=>r(4,t=f));let{row:o}=e;function n(){Fe(o)}function s(){x.delete(`gist/${o.id}`).then(f=>{f.ok&&H(A,t=t.filter(u=>u.id!==o.id),t)})}const l=()=>{Et(`删除 [ ${o.title} ] .${o.lang}`).then(()=>{s()})};return a.$$set=f=>{"row"in f&&r(0,o=f.row)},[o,n,s,l]}class Rn extends Y{constructor(e){super(),X(this,e,Tn,On,K,{row:0})}}function En(a){let e,r;return{c(){e=q("div"),r=new Pe(!1),this.h()},l(t){e=B(t,"DIV",{class:!0});var o=U(e);r=ze(o,!1),o.forEach(v),this.h()},h(){r.a=null,j(e,"class","max-w-[85vh] max-h-[160px] overflow-auto")},m(t,o){D(t,e,o),r.m(a[0],e)},p(t,[o]){o&1&&r.p(t[0])},i:J,o:J,d(t){t&&v(e)}}}function Mn(a,e,r){let{html:t=""}=e;return a.$$set=o=>{"html"in o&&r(0,t=o.html)},[t]}class Ln extends Y{constructor(e){super(),X(this,e,Mn,En,K,{html:0})}}function Vn(a){let e,r;return e=new ot({props:{tableModel:a[1],filters:a[0],viewOption:a[2]}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},p:J,i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function Pn(a,e,r){let t,o,n,s,l,f;z(a,Se,b=>r(16,f=b));const u=!ut,$=Xe(A,{page:Ze({}),filter:ye({serverSide:u,fn:({filterValue:b,value:y})=>y.toLowerCase().includes(b.toLowerCase())}),colFilter:xe({serverSide:u}),hide:et(),sort:tt({serverSide:u,toggleOrder:["asc","desc"]})}),m=$.createColumns([$.column({accessor:"id",header:"ID"}),$.column({accessor:"title",header:"Title",cell:({value:b})=>re(Mt,{text:b})}),$.column({accessor:"lang",header:"Lang"}),$.column({accessor:"content",header:"Content"}),$.column({accessor:"html",header:"HTML",cell:({value:b})=>re(Ln,{html:b})}),$.column({accessor:"created",header:"Created",cell:({value:b})=>me(b),plugins:{filter:{exclude:!0}}}),$.column({accessor:"updated",header:"Updated",cell:({value:b})=>me(b),plugins:{filter:{exclude:!0}}}),$.display({id:"actions",header:()=>"",cell:({row:b})=>b.isData()&&b.original?re(Rn,{row:b.original}):""})]),k=[{name:"lang",options:Ie}],I=$.createViewModel(m),T={type:"hideColumn",selected:gn,options:[]},{pluginStates:N}=I,{pageIndex:R,pageSize:G}=N.page;z(a,R,b=>r(13,l=b)),z(a,G,b=>r(12,s=b));const V=N.filter.filterValue;z(a,V,b=>r(11,n=b));const{filterValues:Q}=N.colFilter;z(a,Q,b=>r(10,o=b));const{sortKeys:C}=N.sort;z(a,C,b=>r(9,t=b));let S=!1;je(()=>{r(8,S=!0)});let P;const w=st(()=>{const b=encodeURIComponent(JSON.stringify(P));f!==b&&hn(b)},200);return a.$$.update=()=>{a.$$.dirty&16128&&S&&(P={page:l,page_size:s,filter_text:n,filter_values:o,sort_key:t[0]},w())},[k,I,T,R,G,V,Q,C,S,t,o,n,s,l]}class zn extends Y{constructor(e){super(),X(this,e,Pn,Vn,K,{})}}function ve(a){let e,r,t,o;return{c(){e=q("div"),r=E(a[1]),t=E("/"),o=E(a[0]),this.h()},l(n){e=B(n,"DIV",{class:!0});var s=U(e);r=M(s,a[1]),t=M(s,"/"),o=M(s,a[0]),s.forEach(v),this.h()},h(){j(e,"class","text-sm")},m(n,s){D(n,e,s),L(e,r),L(e,t),L(e,o)},p(n,s){s&2&&oe(r,n[1]),s&1&&oe(o,n[0])},d(n){n&&v(e)}}}function jn(a){let e,r;return e=new Lt({props:{class:"h-4 w-4 "+(a[2]?"animate-spin":"")}}),{c(){_(e.$$.fragment)},l(t){g(e.$$.fragment,t)},m(t,o){d(e,t,o),r=!0},p(t,o){const n={};o&4&&(n.class="h-4 w-4 "+(t[2]?"animate-spin":"")),e.$set(n)},i(t){r||(c(e.$$.fragment,t),r=!0)},o(t){p(e.$$.fragment,t),r=!1},d(t){h(e,t)}}}function Nn(a){let e,r,t,o,n,s,l,f,u,i;t=new nt({props:{text:"New Gist",onClick:a[4]}});let $=a[2]&&ve(a);return l=new se({props:{variant:"link",size:"icon",$$slots:{default:[jn]},$$scope:{ctx:a}}}),l.$on("click",a[3]),u=new zn({}),{c(){e=q("div"),r=q("div"),_(t.$$.fragment),o=F(),n=q("div"),$&&$.c(),s=F(),_(l.$$.fragment),f=F(),_(u.$$.fragment),this.h()},l(m){e=B(m,"DIV",{class:!0});var k=U(e);r=B(k,"DIV",{class:!0});var I=U(r);g(t.$$.fragment,I),o=O(I),n=B(I,"DIV",{class:!0});var T=U(n);$&&$.l(T),s=O(T),g(l.$$.fragment,T),T.forEach(v),I.forEach(v),f=O(k),g(u.$$.fragment,k),k.forEach(v),this.h()},h(){j(n,"class","flex items-center"),j(r,"class","flex items-center justify-between"),j(e,"class","container mx-auto")},m(m,k){D(m,e,k),L(e,r),d(t,r,null),L(r,o),L(r,n),$&&$.m(n,null),L(n,s),d(l,n,null),L(e,f),d(u,e,null),i=!0},p(m,[k]){m[2]?$?$.p(m,k):($=ve(m),$.c(),$.m(n,s)):$&&($.d(1),$=null);const I={};k&132&&(I.$$scope={dirty:k,ctx:m}),l.$set(I)},i(m){i||(c(t.$$.fragment,m),c(l.$$.fragment,m),c(u.$$.fragment,m),i=!0)},o(m){p(t.$$.fragment,m),p(l.$$.fragment,m),p(u.$$.fragment,m),i=!1},d(m){m&&v(e),h(t),$&&$.d(),h(l),h(u)}}}function qn(a,e,r){let t;z(a,A,i=>r(5,t=i));let o=0,n=0,s=!1;async function l(i,$){const m=await Ot($);(await x.put(`gist/${i}`,{html:m})).ok&&(t.forEach(I=>{I.id===i&&(I.html=m)}),r(1,n++,n),n===o&&(r(1,n=0),r(2,s=!1),A.set(t),Vt.success(`同步完成 ${o} 条`)))}function f(){if(s)return;const i=t;r(0,o=i.length),i.forEach(async $=>{let m=$.content;$.lang!=="md"&&(m=`\`\`\`${$.lang}
${$.content}
\`\`\``),await l($.id,m)}),r(2,s=!0)}return[o,n,s,f,()=>Fe()]}class sr extends Y{constructor(e){super(),X(this,e,qn,Nn,K,{})}}export{sr as component};
