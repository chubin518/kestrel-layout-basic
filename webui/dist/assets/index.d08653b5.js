import{u as Z,_ as v,o as j,b as ee}from"./index.f512f932.js";import{d as h,c as z,D as f,aF as y,aI as t,aJ as e,aN as d,aO as i,b4 as te,br as E,bs as T,H as p,bt as ae,bu as ne,b5 as M,bv as W,M as q,bw as R,b9 as oe,ba as se,f as B,aV as U,u as C,bx as k,by as x,bz as P,aG as le,bA as ce,bB as re,bC as ie,bD as V,bE as ue,E as F,aL as S,aK as A,bj as H,bF as G,bG as _e,bH as de,bI as pe}from"./arco.fe848dac.js";/* empty css              */import{L as O}from"./chart.59e18053.js";import"./vue.b1b7370f.js";const me=h({__name:"banner",setup(n){const l=Z(),o=z(()=>({name:l.name}));return(a,s)=>{const r=te,m=E,c=T;return f(),y(m,{class:"banner"},{default:t(()=>[e(m,{span:8},{default:t(()=>[e(r,{heading:5,style:{"margin-top":"0"}},{default:t(()=>[d(i(a.$t("workplace.welcome"))+" "+i(o.value.name),1)]),_:1})]),_:1}),e(c,{class:"panel-border"})]),_:1})}}});const fe=v(me,[["__scopeId","data-v-5803762b"]]);const ge={},D=n=>(oe("data-v-861c1f3d"),n=n(),se(),n),ye=D(()=>p("img",{alt:"avatar",src:"//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/288b89194e657603ff40db39e8072640.svg~tplv-49unhts6dw-image.image"},null,-1)),he={class:"unit"},ve=D(()=>p("img",{alt:"avatar",src:"//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/fdc66b07224cdf18843c6076c2587eb5.svg~tplv-49unhts6dw-image.image"},null,-1)),xe={class:"unit"},be=D(()=>p("img",{alt:"avatar",src:"//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/77d74c9a245adeae1ec7fb5d4539738d.svg~tplv-49unhts6dw-image.image"},null,-1)),we={class:"unit"},$e=D(()=>p("img",{alt:"avatar",src:"//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/c8b36e26d2b9bb5dbf9b74dd6d7345af.svg~tplv-49unhts6dw-image.image"},null,-1));function Be(n,l){const o=ae,a=ne,s=M,r=W,m=q,c=T,g=R;return f(),y(g,{cols:24,"row-gap":16,class:"panel"},{default:t(()=>[e(r,{class:"panel-col",span:{xs:12,sm:12,md:12,lg:12,xl:12,xxl:6}},{default:t(()=>[e(s,null,{default:t(()=>[e(o,{size:54,class:"col-avatar"},{default:t(()=>[ye]),_:1}),e(a,{title:n.$t("workplace.onlineContent"),value:373.5,precision:1,"value-from":0,animation:"","show-group-separator":""},{suffix:t(()=>[d(" W+ "),p("span",he,i(n.$t("workplace.pecs")),1)]),_:1},8,["title"])]),_:1})]),_:1}),e(r,{class:"panel-col",span:{xs:12,sm:12,md:12,lg:12,xl:12,xxl:6}},{default:t(()=>[e(s,null,{default:t(()=>[e(o,{size:54,class:"col-avatar"},{default:t(()=>[ve]),_:1}),e(a,{title:n.$t("workplace.putIn"),value:368,"value-from":0,animation:"","show-group-separator":""},{suffix:t(()=>[p("span",xe,i(n.$t("workplace.pecs")),1)]),_:1},8,["title"])]),_:1})]),_:1}),e(r,{class:"panel-col",span:{xs:12,sm:12,md:12,lg:12,xl:12,xxl:6}},{default:t(()=>[e(s,null,{default:t(()=>[e(o,{size:54,class:"col-avatar"},{default:t(()=>[be]),_:1}),e(a,{title:n.$t("workplace.newDay"),value:8874,"value-from":0,animation:"","show-group-separator":""},{suffix:t(()=>[p("span",we,i(n.$t("workplace.pecs")),1)]),_:1},8,["title"])]),_:1})]),_:1}),e(r,{class:"panel-col",span:{xs:12,sm:12,md:12,lg:12,xl:12,xxl:6},style:{"border-right":"none"}},{default:t(()=>[e(s,null,{default:t(()=>[e(o,{size:54,class:"col-avatar"},{default:t(()=>[$e]),_:1}),e(a,{title:n.$t("workplace.newFromYesterday"),value:2.8,precision:1,"value-from":0,animation:""},{suffix:t(()=>[d(" % "),e(m,{class:"up-icon"})]),_:1},8,["title"])]),_:1})]),_:1}),e(r,{span:24},{default:t(()=>[e(c,{class:"panel-border"})]),_:1})]),_:1})}const Ce=v(ge,[["render",Be],["__scopeId","data-v-861c1f3d"]]);function N(n=!1){const l=B(n);return{loading:l,setLoading:s=>{l.value=s},toggle:()=>{l.value=!l.value}}}function ke(){return j.get("/api/content-data")}function Fe(n){return j.get("/api/popular/list",{params:n})}function J(n){const l=ee(),o=z(()=>l.theme==="dark");return{chartOption:z(()=>n(o.value))}}const Ee=h({__name:"content-chart",setup(n){function l(u){return{type:"text",bottom:"8",...u,style:{text:"",textAlign:"center",fill:"#4E5969",fontSize:12}}}const{loading:o,setLoading:a}=N(!0),s=B([]),r=B([]),m=B([l({left:"2.6%"}),l({right:0})]),{chartOption:c}=J(()=>({grid:{left:"2.6%",right:"0",top:"10",bottom:"30"},xAxis:{type:"category",offset:2,data:s.value,boundaryGap:!1,axisLabel:{color:"#4E5969",formatter(u,_){return _===0||_===s.value.length-1?"":`${u}`}},axisLine:{show:!1},axisTick:{show:!1},splitLine:{show:!0,interval:u=>!(u===0||u===s.value.length-1),lineStyle:{color:"#E5E8EF"}},axisPointer:{show:!0,lineStyle:{color:"#23ADFF",width:2}}},yAxis:{type:"value",axisLine:{show:!1},axisLabel:{formatter(u,_){return _===0?u:`${u}k`}},splitLine:{show:!0,lineStyle:{type:"dashed",color:"#E5E8EF"}}},tooltip:{trigger:"axis",formatter(u){const[_]=u;return`<div>
            <p class="tooltip-title">${_.axisValueLabel}</p>
            <div class="content-panel"><span>\u603B\u5185\u5BB9\u91CF</span><span class="tooltip-value">${(Number(_.value)*1e4).toLocaleString()}</span></div>
          </div>`},className:"echarts-tooltip-diy"},graphic:{elements:m.value},series:[{data:r.value,type:"line",smooth:!0,symbolSize:12,emphasis:{focus:"series",itemStyle:{borderWidth:2}},lineStyle:{width:3,color:new O(0,0,1,0,[{offset:0,color:"rgba(30, 231, 255, 1)"},{offset:.5,color:"rgba(36, 154, 255, 1)"},{offset:1,color:"rgba(111, 66, 251, 1)"}])},showSymbol:!1,areaStyle:{opacity:.8,color:new O(0,0,0,1,[{offset:0,color:"rgba(17, 126, 255, 0.16)"},{offset:1,color:"rgba(17, 128, 255, 0)"}])}}]}));return(async()=>{a(!0);try{const{data:u}=await ke();u.forEach((_,b)=>{s.value.push(_.x),r.value.push(_.y),b===0&&(m.value[0].style.text=_.x),b===u.length-1&&(m.value[1].style.text=_.x)})}catch{}finally{a(!1)}})(),(u,_)=>{const b=k,w=U("Chart"),L=x,I=P;return f(),y(I,{loading:C(o),style:{width:"100%"}},{default:t(()=>[e(L,{class:"general-card","header-style":{paddingBottom:0},"body-style":{paddingTop:"20px"},title:u.$t("workplace.contentData")},{extra:t(()=>[e(b,null,{default:t(()=>[d(i(u.$t("workplace.viewMore")),1)]),_:1})]),default:t(()=>[e(w,{height:"289px",option:C(c)},null,8,["option"])]),_:1},8,["title"])]),_:1},8,["loading"])}}});const Se={class:"increases-cell"},Ae=h({__name:"popular-content",setup(n){const l=B("text"),{loading:o,setLoading:a}=N(),s=B(),r=async c=>{try{a(!0);const{data:g}=await Fe({type:c});s.value=g}catch{}finally{a(!1)}},m=c=>{r(c)};return r("text"),(c,g)=>{const u=k,_=ce,b=re,w=ie,L=V,I=q,K=ue,Y=M,Q=x,X=P;return f(),y(X,{loading:C(o),style:{width:"100%"}},{default:t(()=>[e(Q,{class:"general-card","header-style":{paddingBottom:"0"},"body-style":{padding:"17px 20px 21px 20px"}},{title:t(()=>[d(i(c.$t("workplace.popularContent")),1)]),extra:t(()=>[e(u,null,{default:t(()=>[d(i(c.$t("workplace.viewMore")),1)]),_:1})]),default:t(()=>[e(Y,{direction:"vertical",size:10,fill:""},{default:t(()=>[e(b,{"model-value":l.value,"onUpdate:modelValue":g[0]||(g[0]=$=>l.value=$),type:"button",onChange:m},{default:t(()=>[e(_,{value:"text"},{default:t(()=>[d(i(c.$t("workplace.popularContent.text")),1)]),_:1}),e(_,{value:"image"},{default:t(()=>[d(i(c.$t("workplace.popularContent.image")),1)]),_:1}),e(_,{value:"video"},{default:t(()=>[d(i(c.$t("workplace.popularContent.video")),1)]),_:1})]),_:1},8,["model-value","onChange"]),e(K,{data:s.value,pagination:!1,bordered:!1,scroll:{x:"100%",y:"264px"}},{columns:t(()=>[e(w,{title:"\u6392\u540D","data-index":"key"}),e(w,{title:"\u5185\u5BB9\u6807\u9898","data-index":"title"},{cell:t(({record:$})=>[e(L,{ellipsis:{rows:1}},{default:t(()=>[d(i($.title),1)]),_:2},1024)]),_:1}),e(w,{title:"\u70B9\u51FB\u91CF","data-index":"clickNumber"}),e(w,{title:"\u65E5\u6DA8\u5E45","data-index":"increases",sortable:{sortDirections:["ascend","descend"]}},{cell:t(({record:$})=>[p("div",Se,[p("span",null,i($.increases)+"%",1),$.increases!==0?(f(),y(I,{key:0,style:{color:"#f53f3f","font-size":"8px"}})):le("",!0)])]),_:1})]),_:1},8,["data"])]),_:1})]),_:1})]),_:1},8,["loading"])}}});const De=v(Ae,[["__scopeId","data-v-6e29d48b"]]),Le=h({__name:"categories-percent",setup(n){const{loading:l}=N(),{chartOption:o}=J(a=>({legend:{left:"center",data:["\u7EAF\u6587\u672C","\u56FE\u6587\u7C7B","\u89C6\u9891\u7C7B"],bottom:0,icon:"circle",itemWidth:8,textStyle:{color:a?"rgba(255, 255, 255, 0.7)":"#4E5969"},itemStyle:{borderWidth:0}},tooltip:{show:!0,trigger:"item"},graphic:{elements:[{type:"text",left:"center",top:"40%",style:{text:"\u5185\u5BB9\u91CF",textAlign:"center",fill:a?"#ffffffb3":"#4E5969",fontSize:14}},{type:"text",left:"center",top:"50%",style:{text:"928,531",textAlign:"center",fill:a?"#ffffffb3":"#1D2129",fontSize:16,fontWeight:500}}]},series:[{type:"pie",radius:["50%","70%"],center:["50%","50%"],label:{formatter:"{d}%",fontSize:14,color:a?"rgba(255, 255, 255, 0.7)":"#4E5969"},itemStyle:{borderColor:a?"#232324":"#fff",borderWidth:1},data:[{value:[148564],name:"\u7EAF\u6587\u672C",itemStyle:{color:a?"#3D72F6":"#249EFF"}},{value:[334271],name:"\u56FE\u6587\u7C7B",itemStyle:{color:a?"#A079DC":"#313CA9"}},{value:[445694],name:"\u89C6\u9891\u7C7B",itemStyle:{color:a?"#6CAAF5":"#21CCFF"}}]}]}));return(a,s)=>{const r=U("Chart"),m=x,c=P;return f(),y(c,{loading:C(l),style:{width:"100%"}},{default:t(()=>[e(m,{class:"general-card","header-style":{paddingBottom:"0"},"body-style":{padding:"20px"}},{title:t(()=>[d(i(a.$t("workplace.categoriesPercent")),1)]),default:t(()=>[e(r,{height:"310px",option:C(o)},null,8,["option"])]),_:1})]),_:1},8,["loading"])}}}),Ie={style:{"margin-bottom":"-1rem"}},ze={class:"icon"},Te=h({__name:"recently-visited",setup(n){const l=[{text:"workplace.contentManagement",icon:"icon-storage"},{text:"workplace.contentStatistical",icon:"icon-file"},{text:"workplace.advanced",icon:"icon-settings"}];return(o,a)=>{const s=V,r=E,m=G,c=x;return f(),y(c,{class:"general-card",title:o.$t("workplace.recently.visited"),"header-style":{paddingBottom:"0"},"body-style":{paddingTop:"26px"}},{default:t(()=>[p("div",Ie,[e(m,{gutter:8},{default:t(()=>[(f(),F(S,null,A(l,g=>e(r,{key:g.text,span:8,class:"wrapper"},{default:t(()=>[p("div",ze,[(f(),y(H(g.icon)))]),e(s,{class:"text"},{default:t(()=>[d(i(o.$t(g.text)),1)]),_:2},1024)]),_:2},1024)),64))]),_:1})])]),_:1},8,["title"])}}});const Pe=v(Te,[["__scopeId","data-v-242aaa28"]]),Ve={class:"icon"},Ge=h({__name:"quick-operation",setup(n){const l=[{text:"workplace.contentManagement",icon:"icon-file"},{text:"workplace.contentStatistical",icon:"icon-storage"},{text:"workplace.advanced",icon:"icon-settings"},{text:"workplace.onlinePromotion",icon:"icon-mobile"},{text:"workplace.contentPutIn",icon:"icon-fire"}];return(o,a)=>{const s=k,r=V,m=E,c=G,g=T,u=x;return f(),y(u,{class:"general-card",title:o.$t("workplace.quick.operation"),"header-style":{paddingBottom:"0"},"body-style":{padding:"24px 20px 0 20px"}},{extra:t(()=>[e(s,null,{default:t(()=>[d(i(o.$t("workplace.quickOperation.setup")),1)]),_:1})]),default:t(()=>[e(c,{gutter:8},{default:t(()=>[(f(),F(S,null,A(l,_=>e(m,{key:_.text,span:8,class:"wrapper"},{default:t(()=>[p("div",Ve,[(f(),y(H(_.icon)))]),e(r,{class:"text"},{default:t(()=>[d(i(o.$t(_.text)),1)]),_:2},1024)]),_:2},1024)),64))]),_:1}),e(g,{class:"split-line",style:{margin:"0"}})]),_:1},8,["title"])}}}),Ne={class:"item-content"},Oe=h({__name:"announcement",setup(n){const l=[{type:"orangered",label:"\u6D3B\u52A8",content:"\u5185\u5BB9\u6700\u65B0\u4F18\u60E0\u6D3B\u52A8"},{type:"cyan",label:"\u6D88\u606F",content:"\u65B0\u589E\u5185\u5BB9\u5C1A\u672A\u901A\u8FC7\u5BA1\u6838\uFF0C\u8BE6\u60C5\u8BF7\u70B9\u51FB\u67E5\u770B\u3002"},{type:"blue",label:"\u901A\u77E5",content:"\u5F53\u524D\u4EA7\u54C1\u8BD5\u7528\u671F\u5373\u5C06\u7ED3\u675F\uFF0C\u5982\u9700\u7EED\u8D39\u8BF7\u70B9\u51FB\u67E5\u770B\u3002"},{type:"blue",label:"\u901A\u77E5",content:"1\u6708\u65B0\u7CFB\u7EDF\u5347\u7EA7\u8BA1\u5212\u901A\u77E5"},{type:"cyan",label:"\u6D88\u606F",content:"\u65B0\u589E\u5185\u5BB9\u5DF2\u7ECF\u901A\u8FC7\u5BA1\u6838\uFF0C\u8BE6\u60C5\u8BF7\u70B9\u51FB\u67E5\u770B\u3002"}];return(o,a)=>{const s=k,r=_e,m=x;return f(),y(m,{class:"general-card",title:o.$t("workplace.announcement"),"header-style":{paddingBottom:"0"},"body-style":{padding:"15px 20px 13px 20px"}},{extra:t(()=>[e(s,null,{default:t(()=>[d(i(o.$t("workplace.viewMore")),1)]),_:1})]),default:t(()=>[p("div",null,[(f(),F(S,null,A(l,(c,g)=>p("div",{key:g,class:"item"},[e(r,{color:c.type,size:"small"},{default:t(()=>[d(i(c.label),1)]),_:2},1032,["color"]),p("span",Ne,i(c.content),1)])),64))])]),_:1},8,["title"])}}});const je=v(Oe,[["__scopeId","data-v-efbb480f"]]);const Me=["src"],We=h({__name:"carousel",setup(n){const l=["//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/5cc3cd1d994b7ef9db6a1f619a22addd.jpg~tplv-49unhts6dw-image.image","//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/f256cbcc287139e191fecea9d255a1f0.jpg~tplv-49unhts6dw-image.image","//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/b557ff0cd44146a2e471b477af2f30d0.jpg~tplv-49unhts6dw-image.image","//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/665106f4bbd2a2df96eaf7aec52f7bc3.jpg~tplv-49unhts6dw-image.image","//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/ea095a2c9c72b5d8f2f2818040db736d.jpg~tplv-49unhts6dw-image.image"];return(o,a)=>{const s=de,r=pe;return f(),y(r,{"indicator-type":"slider","show-arrow":"hover","auto-play":"",style:{width:"100%",height:"170px","border-radius":"4px",overflow:"hidden"}},{default:t(()=>[(f(),F(S,null,A(l,(m,c)=>e(s,{key:c},{default:t(()=>[p("div",null,[p("img",{src:m,style:{width:"100%"}},null,8,Me)])]),_:2},1024)),64))]),_:1})}}});const qe={};function Re(n,l){const o=k,a=E,s=G,r=x;return f(),y(r,{class:"general-card",title:n.$t("workplace.docs"),"header-style":{paddingBottom:0},"body-style":{paddingTop:0},style:{height:"166px"}},{extra:t(()=>[e(o,null,{default:t(()=>[d(i(n.$t("workplace.viewMore")),1)]),_:1})]),default:t(()=>[e(s,null,{default:t(()=>[e(a,{span:12},{default:t(()=>[e(o,null,{default:t(()=>[d(i(n.$t("workplace.docs.productOverview")),1)]),_:1})]),_:1}),e(a,{span:12},{default:t(()=>[e(o,null,{default:t(()=>[d(i(n.$t("workplace.docs.userGuide")),1)]),_:1})]),_:1}),e(a,{span:12},{default:t(()=>[e(o,null,{default:t(()=>[d(i(n.$t("workplace.docs.workflow")),1)]),_:1})]),_:1}),e(a,{span:12},{default:t(()=>[e(o,null,{default:t(()=>[d(i(n.$t("workplace.docs.interfaceDocs")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["title"])}const Ue=v(qe,[["render",Re],["__scopeId","data-v-86d656a8"]]),He={class:"container"},Je={class:"left-side"},Ke={class:"panel"},Ye={class:"right-side"},Qe={class:"panel moduler-wrap"},Xe={name:"Dashboard"},Ze=h({...Xe,setup(n){return(l,o)=>{const a=W,s=R;return f(),F("div",He,[p("div",Je,[p("div",Ke,[e(fe),e(Ce),e(Ee)]),e(s,{cols:24,"col-gap":16,"row-gap":16,style:{"margin-top":"16px"}},{default:t(()=>[e(a,{span:{xs:24,sm:24,md:24,lg:12,xl:12,xxl:12}},{default:t(()=>[e(De)]),_:1}),e(a,{span:{xs:24,sm:24,md:24,lg:12,xl:12,xxl:12}},{default:t(()=>[e(Le)]),_:1})]),_:1})]),p("div",Ye,[e(s,{cols:24,"row-gap":16},{default:t(()=>[e(a,{span:24},{default:t(()=>[p("div",Qe,[e(Ge),e(Pe)])]),_:1}),e(a,{class:"panel",span:24},{default:t(()=>[e(We)]),_:1}),e(a,{class:"panel",span:24},{default:t(()=>[e(je)]),_:1}),e(a,{class:"panel",span:24},{default:t(()=>[e(Ue)]),_:1})]),_:1})])])}}});const st=v(Ze,[["__scopeId","data-v-4f549a29"]]);export{st as default};
