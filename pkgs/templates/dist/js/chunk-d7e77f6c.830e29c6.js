(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-d7e77f6c"],{"32b5":function(t,e,o){"use strict";o("7663")},7663:function(t,e,o){},"7b3d":function(t,e,o){"use strict";o.r(e);var s=function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",[o("sui-menu",{attrs:{fixed:"",inverted:""}},[o("sui-container",[o("sui-menu-item",{tag:"router-link",staticClass:"header",attrs:{to:"/dashboard"}},[t._v(" NetCoCo ")]),o("sui-menu-item",{tag:"router-link",attrs:{to:"/dashboard"}},[t._v("Dashboard")]),o("sui-menu-item",{tag:"router-link",attrs:{to:"/devices"}},[t._v("Devices")]),o("sui-menu-item",{tag:"router-link",attrs:{to:"/topology"}},[t._v("Topology")]),o("sui-menu-menu",{tag:"div",attrs:{position:"right"}},[o("sui-menu-item",{tag:"span"},[o("sui-button",{attrs:{color:"red"},on:{click:function(e){return t.logout()}}},[t._v("Logout")])],1)],1)],1)],1),o("sui-container",{staticClass:"main"},[o("router-view")],1),o("sui-segment",{staticClass:"footer center aligned",attrs:{vertical:""}},[o("FooterText")],1)],1)},n=[],a=o("2b0e"),i=o("84f0"),r=a["a"].extend({components:{FooterText:i["a"]},methods:{clearSession:function(){delete localStorage.logged,this.$router.push("/login")},logout:function(){var t=this,e=localStorage.logged,o=JSON.parse(e);this.$api_connection.unsecureAPI().post("/logout",{username:o.username,password:o.APISecret}).then((function(){t.$toasted.success("Logout success"),t.clearSession()})).catch((function(){t.$toasted.error("Logout failed!")}))}}}),u=r,c=(o("32b5"),o("2877")),l=Object(c["a"])(u,s,n,!1,null,"5c120c67",null);e["default"]=l.exports},"84f0":function(t,e,o){"use strict";var s=function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("span",[t._v("2021 © Chatdanai Phakaket. "),o("sui-icon",{attrs:{size:"large",name:"github square",link:"",onclick:"location.href='https://www.github.com/mrzack99s'"}})],1)},n=[],a=o("2877"),i={},r=Object(a["a"])(i,s,n,!1,null,null,null);e["a"]=r.exports}}]);
//# sourceMappingURL=chunk-d7e77f6c.830e29c6.js.map