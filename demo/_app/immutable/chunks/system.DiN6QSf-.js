const s=(t=null)=>{let o,e,l;return t===null?(o=window.scrollY||document.documentElement.scrollTop,e=document.documentElement.scrollHeight,l=document.documentElement.clientHeight):(o=t.scrollTop,e=t.scrollHeight,l=t.clientHeight),o+l>=e-5},r=(t,o=300)=>{let e;return function(...l){clearTimeout(e),e=setTimeout(()=>t.apply(this,l),o)}},u=(t,o=300)=>{let e,l,n;return function(){const c=this,i=arguments;e?(clearTimeout(l),l=setTimeout(()=>{Date.now()-n>=o&&(t.apply(c,i),n=Date.now())},Math.max(o-(Date.now()-n),0))):(t.apply(c,i),n=Date.now(),e=!0)}};export{r as d,s as i,u as t};
