(function(t){function a(a){for(var s,c,r=a[0],o=a[1],d=a[2],u=0,v=[];u<r.length;u++)c=r[u],Object.prototype.hasOwnProperty.call(e,c)&&e[c]&&v.push(e[c][0]),e[c]=0;for(s in o)Object.prototype.hasOwnProperty.call(o,s)&&(t[s]=o[s]);l&&l(a);while(v.length)v.shift()();return i.push.apply(i,d||[]),n()}function n(){for(var t,a=0;a<i.length;a++){for(var n=i[a],s=!0,r=1;r<n.length;r++){var o=n[r];0!==e[o]&&(s=!1)}s&&(i.splice(a--,1),t=c(c.s=n[0]))}return t}var s={},e={app:0},i=[];function c(a){if(s[a])return s[a].exports;var n=s[a]={i:a,l:!1,exports:{}};return t[a].call(n.exports,n,n.exports,c),n.l=!0,n.exports}c.m=t,c.c=s,c.d=function(t,a,n){c.o(t,a)||Object.defineProperty(t,a,{enumerable:!0,get:n})},c.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},c.t=function(t,a){if(1&a&&(t=c(t)),8&a)return t;if(4&a&&"object"===typeof t&&t&&t.__esModule)return t;var n=Object.create(null);if(c.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:t}),2&a&&"string"!=typeof t)for(var s in t)c.d(n,s,function(a){return t[a]}.bind(null,s));return n},c.n=function(t){var a=t&&t.__esModule?function(){return t["default"]}:function(){return t};return c.d(a,"a",a),a},c.o=function(t,a){return Object.prototype.hasOwnProperty.call(t,a)},c.p="/";var r=window["webpackJsonp"]=window["webpackJsonp"]||[],o=r.push.bind(r);r.push=a,r=r.slice();for(var d=0;d<r.length;d++)a(r[d]);var l=o;i.push([0,"chunk-vendors"]),n()})({0:function(t,a,n){t.exports=n("56d7")},1:function(t,a){},1395:function(t,a,n){},2395:function(t,a,n){},"3f4f":function(t,a,n){"use strict";n("4f90")},"48bb":function(t,a,n){},"4f90":function(t,a,n){},5405:function(t,a,n){"use strict";n("48bb")},"56d7":function(t,a,n){"use strict";n.r(a);n("e260"),n("e6cf"),n("cca6"),n("a79d");var s=n("2b0e"),e=function(){var t=this,a=t.$createElement,n=t._self._c||a;return n("div",{attrs:{id:"app"}},[n("Index")],1)},i=[],c=function(){var t=this,a=t.$createElement,n=t._self._c||a;return n("div",{attrs:{id:"data-view"}},[n("dv-full-screen-container",[n("div",{staticClass:"main-header"},[n("div",{staticClass:"mh-left"}),n("div",{staticClass:"mh-middle"},[t._v("流量掌柜")]),n("div",{staticClass:"mh-right"},[t.hasData?n("dv-border-box-2",{staticClass:"box"},[t._v(" "+t._s(t.data.updated_at)+" ")]):t._e()],1)]),n("dv-border-box-1",{staticClass:"main-container"},[n("dv-border-box-3",{staticClass:"left-chart-container"}),n("div",{staticClass:"right-main-container"},[n("div",{staticClass:"rmc-top-container"},[n("dv-border-box-3",{staticClass:"rmctc-left-container"},[t.hasData?n("Center",{attrs:{data:t.data}}):t._e(),t.qrcodeLatest?n("Qrcode",{attrs:{src:t.src}}):t._e(),t.loading?n("DvLoading",[t._v(t._s(t.loadingText))]):t._e()],1),n("div",{staticClass:"rmctc-right-container"},[n("dv-border-box-3",{staticClass:"rmctc-chart-1"},[n("UV",{attrs:{oldUV:t.data.uv}})],1),n("dv-border-box-4",{staticClass:"rmctc-chart-2",attrs:{reverse:!0}},[t.hasData?n("Chart",{attrs:{data:t.data}}):t._e()],1)],1)],1),n("dv-border-box-4",{staticClass:"rmc-bottom-container"})],1)],1)],1)],1)},r=[],o=n("1da1"),d=(n("96cf"),function(){var t=this,a=t.$createElement,n=t._self._c||a;return n("div",{staticClass:"center-cmp"},[n("div",{staticClass:"cc-header"},[n("dv-decoration-1",{staticStyle:{width:"200px",height:"50px"}}),n("div",[t._v(t._s(this.data.title))]),n("dv-decoration-1",{staticStyle:{width:"200px",height:"50px"}})],1),n("div",{staticClass:"cc-details"},[n("div",[t._v("GMV")]),t._l(this.data.gmv,(function(a,s){return n("div",{staticClass:"card"},[t._v(t._s(a))])}))],2),n("div",{staticClass:"cc-main-container"},[n("div",{staticClass:"ccmc-left"},[n("div",{staticClass:"station-info"},[t._v(" 实时刷单金额"),n("span",[t._v(t._s(this.data.sssd))])]),n("div",{staticClass:"station-info"},[t._v(" 实时uv价值"),n("span",[t._v(t._s(this.data.suv))])]),n("div",{staticClass:"station-info"},[t._v(" 订单转化率"),n("span",[t._v(t._s(this.data.ozhl))])]),n("div",{staticClass:"station-info"},[t._v(" 成交人数转化率"),n("span",[t._v(t._s(this.data.cjrszhl))])]),n("div",{staticClass:"station-info"},[t._v(" 转粉率"),n("span",[t._v(t._s(this.data.zfl))])])]),n("dv-active-ring-chart",{staticClass:"ccmc-middle",attrs:{config:t.config}}),n("div",{staticClass:"ccmc-right"},[n("div",{staticClass:"station-info"},[n("span",[t._v(t._s(this.data.kdj))]),t._v("客单价 ")]),n("div",{staticClass:"station-info"},[n("span",[t._v(t._s(this.data.pay_ucnt))]),t._v("成交人数 ")]),n("div",{staticClass:"station-info"},[n("span",[t._v(t._s(this.data.rjkbsc))]),t._v("人均看播时长 ")]),n("div",{staticClass:"station-info"},[n("span",[t._v(t._s(this.data.cjfszb))]),t._v("成交粉丝占比 ")]),n("div",{staticClass:"station-info"},[n("span",[t._v(t._s(this.data.online_user_ucnt))]),t._v("累计观看人数 ")])])],1)])}),l=[],u=(n("a9e3"),{name:"Center",components:{},props:["data"],data:function(){return{config:{data:[{name:"成交件数",value:Number(this.data.pay_cnt)}],color:["#00e9ff"],lineWidth:30,radius:"60%",activeRadius:"60%",showOriginValue:!0,digitalFlopStyle:{fontSize:25,fill:"#fff"}}}}}),v=u,f=(n("5405"),n("2877")),p=Object(f["a"])(v,d,l,!1,null,null,null),h=p.exports,_=function(){var t=this,a=t.$createElement,n=t._self._c||a;return n("div",{staticClass:"right-chart-1"},[t._m(0),n("div",{staticClass:"rc1-container"},[n("div",{staticClass:"rc1-form"},[n("input",{directives:[{name:"model",rawName:"v-model",value:t.uv,expression:"uv"}],attrs:{type:"number"},domProps:{value:t.uv},on:{input:[function(a){a.target.composing||(t.uv=a.target.value)},t.setUv]}})])])])},m=[function(){var t=this,a=t.$createElement,n=t._self._c||a;return n("div",{staticClass:"r1-header"},[n("span",{staticClass:"r1-span"},[t._v("预期UV")])])}],b={name:"UV",props:["oldUV"],data:function(){return{uv:0}},created:function(){this.uv=this.oldUV},methods:{setUv:function(){this.axios.post("api/set/uv",this.qs.stringify({uv:this.uv})).then((function(t){})).catch((function(t){alert("设置uv异常! 信息: "+t)}))}}},g=b,C=(n("5b1c"),Object(f["a"])(g,_,m,!1,null,null,null)),x=C.exports,y=function(){var t=this,a=t.$createElement,n=t._self._c||a;return n("div",{staticClass:"right-chart-2"},[n("div",{staticClass:"ccmc-left"},[n("div",{staticClass:"station-info"},[t._v(" 商品曝光人数"),n("span",[t._v(t._s(this.data.exposure))])]),n("div",{staticClass:"station-info"},[t._v(" 商品点击人数"),n("span",[t._v(t._s(this.data.click))])]),n("div",{staticClass:"station-info"},[t._v(" 新增粉丝数"),n("span",[t._v(t._s(this.data.incr_fans_cnt))])]),n("div",{staticClass:"station-info"},[t._v(" 购物车点击率"),n("span",[t._v(t._s(this.data.gwcdjl))])])])])},j=[],w={name:"Chart",props:["data"],data:function(){return{}}},O=w,k=(n("3f4f"),Object(f["a"])(O,y,j,!1,null,null,null)),q=k.exports,D=function(){var t=this,a=t.$createElement,n=t._self._c||a;return n("div",{staticClass:"qrcode"},[n("img",{attrs:{src:t.src}})])},z=[],S={name:"Qrcode",props:["src"],methods:{}},T=S,U=(n("beb3"),Object(f["a"])(T,D,z,!1,null,null,null)),V=U.exports,P={name:"Index",components:{UV:x,Center:h,Chart:q,Qrcode:V},data:function(){return{loading:!0,loadingText:"初始化中...",qrcodeLatest:!1,src:"",hasData:!1,data:{updated_at:"",pay_cnt:"",pay_ucnt:"",incr_fans_cnt:"",online_user_ucnt:"",gmv:0,exposure:"",click:"",yin_liu:"",f_yin_liu:"",sssd:"",uv:"",suv:"",ozhl:"",cjrszhl:"",zfl:"",gwcdjl:"",kdj:"",cjfszb:"",rjkbsc:""}}},mounted:function(){this.getData()},methods:{getData:function(){var t=this;this.axios.get("/api/get/last/data").then(function(){var a=Object(o["a"])(regeneratorRuntime.mark((function a(n){return regeneratorRuntime.wrap((function(a){while(1)switch(a.prev=a.next){case 0:if(n=n.data,4003!==n.code){a.next=6;break}return a.next=4,t.login();case 4:a.next=7;break;case 6:n.data&&""!==n.data.title?(t.loading=!1,t.hasData=!0,t.data=n.data):(t.hasData=!1,t.loading=!0,t.loadingText="后台正在拼命拉取数据中");case 7:setTimeout((function(){t.getData()}),500);case 8:case"end":return a.stop()}}),a)})));return function(t){return a.apply(this,arguments)}}()).catch((function(t){alert("获取数据异常! 信息: "+t)}))},login:function(){var t=this;this.axios.get("api/get/qrcode").then((function(a){a=a.data,200===a.code?(t.qrcodeLatest=a.qrcodeLatest,t.qrcodeLatest&&(t.loading=!1,t.src="tmp/qrcode.png?time="+(new Date).getTime())):(t.qrcodeLatest=!1,t.loading=!0,t.loadingText="正在检测登录")})).catch((function(t){alert("检测登录异常! 信息: "+t)}))}}},$=P,E=(n("65d3"),Object(f["a"])($,c,r,!1,null,null,null)),L=E.exports,M={name:"app",components:{Index:L},data:function(){return{}}},I=M,Q=(n("7c55"),Object(f["a"])(I,e,i,!1,null,null,null)),R=Q.exports,J=(n("1395"),n("6c29")),N=n("bc3a"),F=n.n(N),G=n("4328"),W=n.n(G);s["a"].prototype.axios=F.a,s["a"].prototype.qs=W.a,s["a"].config.productionTip=!1,s["a"].use(J["a"]),new s["a"]({render:function(t){return t(R)}}).$mount("#app")},"5b0d":function(t,a,n){},"5b1c":function(t,a,n){"use strict";n("5bd4")},"5bd4":function(t,a,n){},"65d3":function(t,a,n){"use strict";n("c316")},"7c55":function(t,a,n){"use strict";n("2395")},beb3:function(t,a,n){"use strict";n("5b0d")},c316:function(t,a,n){}});