import{s as S,v as ce,c as ie,e as b,a as j,b as w,d as O,i as L,g as E,h as $,n as g,w as _e,J as ge,t as he,f as be,I as W}from"../chunks/scheduler.CZixQE4Q.js";import{S as R,i as P,d as m,v as N,a as h,g as v,b as z,c as _,l as H,s as we,e as T,f as C,m as D,t as k,h as x,j as y}from"../chunks/index.BQyUQ7-O.js";import{k as I,j as V,l as J,n as G,o as K,p as Q,q as X,D as le,U as ue}from"../chunks/date.D6FUEdcB.js";import{T as de}from"../chunks/table.B0uzwLsW.js";import{B as fe}from"../chunks/index.7rvypQW0.js";import"../chunks/Toaster.svelte_svelte_type_style_lang.31cgKyMx.js";import{t as p,o as B,i as Y,s as Z}from"../chunks/data.aZv_tajY.js";import{f as ee}from"../chunks/fetch.D4GHD56D.js";import{a as te}from"../chunks/alert.BsUfUaeA.js";import{g as ne,b as oe}from"../chunks/entry.BgRXhY7P.js";const U="src/routes/admin/tag/(components)/table-actions.svelte";function A(a){let e,n,i="Edit",c,o,r="Delete",s,l;const d={c:function(){e=b("div"),n=b("button"),n.textContent=i,c=j(),o=b("button"),o.textContent=r,this.h()},l:function(t){e=w(t,"DIV",{class:!0});var f=O(e);n=w(f,"BUTTON",{class:!0,"data-svelte-h":!0}),L(n)!=="svelte-s77gdn"&&(n.textContent=i),c=E(f),o=w(f,"BUTTON",{class:!0,"data-svelte-h":!0}),L(o)!=="svelte-133nxbs"&&(o.textContent=r),f.forEach(h),this.h()},h:function(){v(n,"class","btn s-yw_WgadpAdwp"),$(n,U,22,2,455),v(o,"class","btn bg-foreground text-background s-yw_WgadpAdwp"),$(o,U,23,2,507),v(e,"class","flex h-8 items-center justify-center gap-1"),$(e,U,21,0,396)},m:function(t,f){z(t,e,f),_(e,n),_(e,c),_(e,o),s||(l=[H(n,"click",a[1],!1,!1,!1,!1),H(o,"click",a[3],!1,!1,!1,!1)],s=!0)},p:g,i:g,o:g,d:function(t){t&&h(e),s=!1,_e(l)}};return m("SvelteRegisterBlock",{block:d,id:A.name,type:"component",source:"",ctx:a}),d}function $e(a,e,n){let i;ce(p,"tableData"),ie(a,p,t=>n(4,i=t));let{$$slots:c={},$$scope:o}=e;N("Table_actions",c,[]);let{row:r}=e;function s(){B(r)}function l(){ee.delete(`tag/${r.id}`).then(t=>{t.ok&&ge(p,i=i.filter(f=>f.id!==r.id),i)})}a.$$.on_mount.push(function(){r===void 0&&!("row"in e||a.$$.bound[a.$$.props.row])&&console.warn("<Table_actions> was created without expected prop 'row'")});const d=["row"];Object.keys(e).forEach(t=>{!~d.indexOf(t)&&t.slice(0,2)!=="$$"&&t!=="slot"&&console.warn(`<Table_actions> was created with unknown prop '${t}'`)});const u=()=>{te(`删除标签：${r.name}`).then(()=>{l()})};return a.$$set=t=>{"row"in t&&n(0,r=t.row)},a.$capture_state=()=>({openForm:B,tableData:p,fet:ee,alertDialog:te,row:r,edit:s,del:l,$tableData:i}),a.$inject_state=t=>{"row"in t&&n(0,r=t.row)},e&&"$$inject"in e&&a.$inject_state(e.$$inject),[r,s,l,u]}class ae extends R{constructor(e){super(e),P(this,e,$e,A,S,{row:0}),m("SvelteRegisterComponent",{component:this,tagName:"Table_actions",options:e,id:A.name})}get row(){throw new Error("<Table_actions>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}set row(e){throw new Error("<Table_actions>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}}const re="src/routes/admin/tag/(components)/table-row-color.svelte";function F(a){let e,n,i,c;const o={c:function(){e=b("div"),n=b("div"),i=j(),c=he(a[0]),this.h()},l:function(s){e=w(s,"DIV",{class:!0});var l=O(e);n=w(l,"DIV",{class:!0,style:!0}),O(n).forEach(h),i=E(l),c=be(l,a[0]),l.forEach(h),this.h()},h:function(){v(n,"class","mr-1 h-4 w-4 rounded-sm shadow"),W(n,"background-color",a[0]),$(n,re,5,2,103),v(e,"class","flex items-center text-muted-foreground"),$(e,re,4,0,47)},m:function(s,l){z(s,e,l),_(e,n),_(e,i),_(e,c)},p:function(s,[l]){l&1&&W(n,"background-color",s[0]),l&1&&we(c,s[0])},i:g,o:g,d:function(s){s&&h(e)}};return m("SvelteRegisterBlock",{block:o,id:F.name,type:"component",source:"",ctx:a}),o}function ve(a,e,n){let{$$slots:i={},$$scope:c}=e;N("Table_row_color",i,[]);let{color:o}=e;a.$$.on_mount.push(function(){o===void 0&&!("color"in e||a.$$.bound[a.$$.props.color])&&console.warn("<Table_row_color> was created without expected prop 'color'")});const r=["color"];return Object.keys(e).forEach(s=>{!~r.indexOf(s)&&s.slice(0,2)!=="$$"&&s!=="slot"&&console.warn(`<Table_row_color> was created with unknown prop '${s}'`)}),a.$$set=s=>{"color"in s&&n(0,o=s.color)},a.$capture_state=()=>({color:o}),a.$inject_state=s=>{"color"in s&&n(0,o=s.color)},e&&"$$inject"in e&&a.$inject_state(e.$$inject),[o]}class se extends R{constructor(e){super(e),P(this,e,ve,F,S,{color:0}),m("SvelteRegisterComponent",{component:this,tagName:"Table_row_color",options:e,id:F.name})}get color(){throw new Error("<Table_row_color>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}set color(e){throw new Error("<Table_row_color>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'")}}function M(a){let e,n;e=new de({props:{tableModel:a[0],viewOption:a[1]},$$inline:!0});const i={c:function(){T(e.$$.fragment)},l:function(o){C(e.$$.fragment,o)},m:function(o,r){D(e,o,r),n=!0},p:g,i:function(o){n||(k(e.$$.fragment,o),n=!0)},o:function(o){x(e.$$.fragment,o),n=!1},d:function(o){y(e,o)}};return m("SvelteRegisterBlock",{block:i,id:M.name,type:"component",source:"",ctx:a}),i}function Te(a,e,n){let i;ce(p,"tableData"),ie(a,p,t=>n(2,i=t));let{$$slots:c={},$$scope:o}=e;N("Table",c,[]),i.length===0&&Y();const r=J(p,{page:K(),filter:X({fn:({filterValue:t,value:f})=>f.toLowerCase().includes(t.toLowerCase())}),hide:G(),sort:Q({toggleOrder:["asc","desc"]})}),s=r.createColumns([r.column({accessor:"id",header:"ID"}),r.column({accessor:"name",header:"Name"}),r.column({accessor:"color",header:"Color",cell:({value:t})=>I(se,{color:t}),plugins:{filter:{exclude:!0}}}),r.column({accessor:"created",header:"Created",cell:({value:t})=>V(t),plugins:{filter:{exclude:!0}}}),r.column({accessor:"updated",header:"Updated",cell:({value:t})=>V(t),plugins:{filter:{exclude:!0}}}),r.display({id:"actions",header:()=>"",cell:({row:t})=>t.isData()&&t.original?I(ae,{row:t.original}):""})]),l=r.createViewModel(s),d={type:"hideColumn",selected:Z,options:[]},u=[];return Object.keys(e).forEach(t=>{!~u.indexOf(t)&&t.slice(0,2)!=="$$"&&t!=="slot"&&console.warn(`<Table> was created with unknown prop '${t}'`)}),a.$capture_state=()=>({createRender:I,createTable:J,addHiddenColumns:G,addPagination:K,addSortBy:Q,addTableFilter:X,DataTable:de,TableActions:ae,initTableData:Y,selectedViewOption:Z,tableData:p,TableRowColor:se,DateFormat:V,table:r,columns:s,tableModel:l,viewOption:d,$tableData:i}),[l,d]}class me extends R{constructor(e){super(e),P(this,e,Te,M,S,{}),m("SvelteRegisterComponent",{component:this,tagName:"Table",options:e,id:M.name})}}const Ce="src/routes/admin/tag/+page.svelte";function pe(a){let e,n;e=new ue({props:{class:"h-4 w-4"},$$inline:!0});const i={c:function(){T(e.$$.fragment)},l:function(o){C(e.$$.fragment,o)},m:function(o,r){D(e,o,r),n=!0},p:g,i:function(o){n||(k(e.$$.fragment,o),n=!0)},o:function(o){x(e.$$.fragment,o),n=!1},d:function(o){y(e,o)}};return m("SvelteRegisterBlock",{block:i,id:pe.name,type:"slot",source:'(11:2) <Button     variant=\\"ghost\\"     size=\\"icon\\"     on:click={() => {       goto(`${base}/admin/article`)     }}   >',ctx:a}),i}function q(a){let e,n,i,c,o,r,s;n=new fe({props:{variant:"ghost",size:"icon",$$slots:{default:[pe]},$$scope:{ctx:a}},$$inline:!0}),n.$on("click",a[0]),c=new le({props:{text:"New Tag",onClick:a[1]},$$inline:!0}),r=new me({$$inline:!0});const l={c:function(){e=b("div"),T(n.$$.fragment),i=j(),T(c.$$.fragment),o=j(),T(r.$$.fragment),this.h()},l:function(u){e=w(u,"DIV",{class:!0});var t=O(e);C(n.$$.fragment,t),i=E(t),C(c.$$.fragment,t),o=E(t),C(r.$$.fragment,t),t.forEach(h),this.h()},h:function(){v(e,"class","container mx-auto"),$(e,Ce,10,0,375)},m:function(u,t){z(u,e,t),D(n,e,null),_(e,i),D(c,e,null),_(e,o),D(r,e,null),s=!0},p:function(u,[t]){const f={};t&4&&(f.$$scope={dirty:t,ctx:u}),n.$set(f)},i:function(u){s||(k(n.$$.fragment,u),k(c.$$.fragment,u),k(r.$$.fragment,u),s=!0)},o:function(u){x(n.$$.fragment,u),x(c.$$.fragment,u),x(r.$$.fragment,u),s=!1},d:function(u){u&&h(e),y(n),y(c),y(r)}};return m("SvelteRegisterBlock",{block:l,id:q.name,type:"component",source:"",ctx:a}),l}function De(a,e,n){let{$$slots:i={},$$scope:c}=e;N("Page",i,[]);const o=[];Object.keys(e).forEach(l=>{!~o.indexOf(l)&&l.slice(0,2)!=="$$"&&l!=="slot"&&console.warn(`<Page> was created with unknown prop '${l}'`)});const r=()=>{ne(`${oe}/admin/article`)},s=()=>B();return a.$capture_state=()=>({Table:me,DataActionButton:le,openForm:B,Button:fe,Undo:ue,goto:ne,base:oe}),[r,s]}class Ne extends R{constructor(e){super(e),P(this,e,De,q,S,{}),m("SvelteRegisterComponent",{component:this,tagName:"Page",options:e,id:q.name})}}export{Ne as component};
