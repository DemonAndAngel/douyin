(function(e) {
        function t(t) {
            var n = t[0];
            var c = t[1];
            var i = t[2];
            var u, s, l = 0, d = [];
            for (; l < n.length; l++) {
                s = n[l];
                if (Object.prototype.hasOwnProperty.call(a, s) && a[s])
                    d.push(a[s][0]);
                a[s] = 0
            }
            for (u in c)
                if (Object.prototype.hasOwnProperty.call(c, u))
                    e[u] = c[u];
            if (f)
                f(t);
            while (d.length)
                d.shift()();
            o.push.apply(o, i || []);
            return r()
        }
        function r() {
            var e;
            for (var t = 0; t < o.length; t++) {
                var r = o[t];
                var n = true;
                for (var c = 1; c < r.length; c++) {
                    var u = r[c];
                    if (0 !== a[u])
                        n = false
                }
                if (n) {
                    o.splice(t--, 1);
                    e = i(i.s = r[0])
                }
            }
            return e
        }
        var n = {};
        var a = {
            6: 0
        };
        var o = [];
        function c(e) {
            return i.p + "common/" + ({
                1: "TaskLanding",
                2: "agreement",
                3: "authority",
                4: "college",
                5: "home",
                7: "message",
                8: "organization",
                9: "rules",
                10: "topic",
                11: "topicdetail",
                12: "weekly"
            }[e] || e) + "." + {
                1: "f0a5b09e",
                2: "d7e49477",
                3: "150ea3f5",
                4: "67658a41",
                5: "1d4a68fa",
                7: "9d4cb134",
                8: "52b9a9be",
                9: "465d75b4",
                10: "67f56f12",
                11: "fc94426e",
                12: "3f5b1603"
            }[e] + ".js"
        }
        function i(t) {
            if (n[t])
                return n[t].exports;
            var r = n[t] = {
                i: t,
                l: false,
                exports: {}
            };
            e[t].call(r.exports, r, r.exports, i);
            r.l = true;
            return r.exports
        }
        i.e = function e(t) {
            var r = [];
            if (!l || !l.prototype.then) {
                function n(e) {
                    var t = this.constructor;
                    return this.then((function(r) {
                            return t.resolve(e()).then((function() {
                                    return r
                                }
                            ))
                        }
                    ), (function(r) {
                            return t.resolve(e()).then((function() {
                                    return t.reject(r)
                                }
                            ))
                        }
                    ))
                }
                var o = setTimeout;
                function u() {}
                function s(e, t) {
                    return function() {
                        e.apply(t, arguments)
                    }
                }
                function l(e) {
                    if (!(this instanceof l))
                        throw new TypeError("Promises must be constructed via new");
                    if ("function" !== typeof e)
                        throw new TypeError("not a function");
                    this._state = 0;
                    this._handled = false;
                    this._value = void 0;
                    this._deferreds = [];
                    v(e, this)
                }
                function f(e, t) {
                    while (3 === e._state)
                        e = e._value;
                    if (0 === e._state) {
                        e._deferreds.push(t);
                        return
                    }
                    e._handled = true;
                    l._immediateFn((function() {
                            var r = 1 === e._state ? t.onFulfilled : t.onRejected;
                            if (null === r) {
                                (1 === e._state ? d : p)(t.promise, e._value);
                                return
                            }
                            var n;
                            try {
                                n = r(e._value)
                            } catch (a) {
                                p(t.promise, a);
                                return
                            }
                            d(t.promise, n)
                        }
                    ))
                }
                function d(e, t) {
                    try {
                        if (t === e)
                            throw new TypeError("A promise cannot be resolved with itself.");
                        if (t && ("object" === typeof t || "function" === typeof t)) {
                            var r = t.then;
                            if (t instanceof l) {
                                e._state = 3;
                                e._value = t;
                                b(e);
                                return
                            } else if ("function" === typeof r) {
                                v(s(r, t), e);
                                return
                            }
                        }
                        e._state = 1;
                        e._value = t;
                        b(e)
                    } catch (n) {
                        p(e, n)
                    }
                }
                function p(e, t) {
                    e._state = 2;
                    e._value = t;
                    b(e)
                }
                function b(e) {
                    if (2 === e._state && 0 === e._deferreds.length)
                        l._immediateFn((function() {
                                if (!e._handled)
                                    l._unhandledRejectionFn(e._value)
                            }
                        ));
                    for (var t = 0, r = e._deferreds.length; t < r; t++)
                        f(e, e._deferreds[t]);
                    e._deferreds = null
                }
                function m(e, t, r) {
                    this.onFulfilled = "function" === typeof e ? e : null;
                    this.onRejected = "function" === typeof t ? t : null;
                    this.promise = r
                }
                function v(e, t) {
                    var r = false;
                    try {
                        e((function(e) {
                                if (r)
                                    return;
                                r = true;
                                d(t, e)
                            }
                        ), (function(e) {
                                if (r)
                                    return;
                                r = true;
                                p(t, e)
                            }
                        ))
                    } catch (n) {
                        if (r)
                            return;
                        r = true;
                        p(t, n)
                    }
                }
                l.prototype["catch"] = function(e) {
                    return this.then(null, e)
                }
                ;
                l.prototype.then = function(e, t) {
                    var r = new this.constructor(u);
                    f(this, new m(e,t,r));
                    return r
                }
                ;
                l.prototype["finally"] = n;
                l.all = function(e) {
                    return new l((function(t, r) {
                            if (!e || "undefined" === typeof e.length)
                                throw new TypeError("Promise.all accepts an array");
                            var n = Array.prototype.slice.call(e);
                            if (0 === n.length)
                                return t([]);
                            var a = n.length;
                            function o(e, c) {
                                try {
                                    if (c && ("object" === typeof c || "function" === typeof c)) {
                                        var i = c.then;
                                        if ("function" === typeof i) {
                                            i.call(c, (function(t) {
                                                    o(e, t)
                                                }
                                            ), r);
                                            return
                                        }
                                    }
                                    n[e] = c;
                                    if (0 === --a)
                                        t(n)
                                } catch (u) {
                                    r(u)
                                }
                            }
                            for (var c = 0; c < n.length; c++)
                                o(c, n[c])
                        }
                    ))
                }
                ;
                l.resolve = function(e) {
                    if (e && "object" === typeof e && e.constructor === l)
                        return e;
                    return new l((function(t) {
                            t(e)
                        }
                    ))
                }
                ;
                l.reject = function(e) {
                    return new l((function(t, r) {
                            r(e)
                        }
                    ))
                }
                ;
                l.race = function(e) {
                    return new l((function(t, r) {
                            for (var n = 0, a = e.length; n < a; n++)
                                e[n].then(t, r)
                        }
                    ))
                }
                ;
                l._immediateFn = "function" === typeof setImmediate && function(e) {
                        setImmediate(e)
                    }
                    || function(e) {
                        o(e, 0)
                    }
                ;
                l._unhandledRejectionFn = function e(t) {
                    if ("undefined" !== typeof console && console)
                        void 0
                }
            }
            var h = a[t];
            if (0 !== h)
                if (h)
                    r.push(h[2]);
                else {
                    var y = new l((function(e, r) {
                            h = a[t] = [e, r]
                        }
                    ));
                    r.push(h[2] = y);
                    var g = document.createElement("script");
                    var O;
                    g.charset = "utf-8";
                    g.timeout = 120;
                    if (i.nc)
                        g.setAttribute("nonce", i.nc);
                    g.src = c(t);
                    var w = new Error;
                    O = function(e) {
                        g.onerror = g.onload = null;
                        clearTimeout(_);
                        var r = a[t];
                        if (0 !== r) {
                            if (r) {
                                var n = e && ("load" === e.type ? "missing" : e.type);
                                var o = e && e.target && e.target.src;
                                w.message = "Loading chunk " + t + " failed.\n(" + n + ": " + o + ")";
                                w.name = "ChunkLoadError";
                                w.type = n;
                                w.request = o;
                                r[1](w)
                            }
                            a[t] = void 0
                        }
                    }
                    ;
                    var _ = setTimeout((function() {
                            O({
                                type: "timeout",
                                target: g
                            })
                        }
                    ), 12e4);
                    g.onerror = g.onload = O;
                    document.head.appendChild(g)
                }
            return l.all(r)
        }
        ;
        i.m = e;
        i.c = n;
        i.d = function(e, t, r) {
            if (!i.o(e, t))
                Object.defineProperty(e, t, {
                    enumerable: true,
                    get: r
                })
        }
        ;
        i.r = function(e) {
            if ("undefined" !== typeof Symbol && Symbol.toStringTag)
                Object.defineProperty(e, Symbol.toStringTag, {
                    value: "Module"
                });
            Object.defineProperty(e, "__esModule", {
                value: true
            })
        }
        ;
        i.t = function(e, t) {
            if (1 & t)
                e = i(e);
            if (8 & t)
                return e;
            if (4 & t && "object" === typeof e && e && e.__esModule)
                return e;
            var r = Object.create(null);
            i.r(r);
            Object.defineProperty(r, "default", {
                enumerable: true,
                value: e
            });
            if (2 & t && "string" != typeof e)
                for (var n in e)
                    i.d(r, n, function(t) {
                        return e[t]
                    }
                        .bind(null, n));
            return r
        }
        ;
        i.n = function(e) {
            var t = e && e.__esModule ? function t() {
                    return e["default"]
                }
                : function t() {
                    return e
                }
            ;
            i.d(t, "a", t);
            return t
        }
        ;
        i.o = function(e, t) {
            return Object.prototype.hasOwnProperty.call(e, t)
        }
        ;
        i.p = "//sf3-scmcdn-tos.pstatp.com/goofy/ies/douyin_creator/";
        i.oe = function(e) {
            void 0;
            throw e
        }
        ;
        var u = window["webpackJsonp"] = window["webpackJsonp"] || [];
        var s = u.push.bind(u);
        u.push = t;
        u = u.slice();
        for (var l = 0; l < u.length; l++)
            t(u[l]);
        var f = s;
        o.push([0, 0]);
        return r()
    }
)({
    0: function(e, t, r) {
        r("9ae88e6caac879637e63");
        e.exports = r("3177845424933048caec")
    },
    "275976081ce1abf67779": function(e, t) {
        e.exports = React
    },
    "2a1bb827226863258ede": function(e, t, r) {
        "use strict";
        t["a"] = {
            count: function e(t) {
                window.Slardar && window.Slardar("emit", "counter", {
                    name: t,
                    value: 1
                })
            },
            timer: function e(t) {
                window.Slardar && window.Slardar("emit", "timer", {
                    name: t
                })
            }
        }
    },
    "3177845424933048caec": function(e, t, r) {
        "use strict";
        r.r(t);
        (function(e) {
                var t = r("275976081ce1abf67779");
                var n = r.n(t);
                var a = r("de7f2df91c095f0cd45d");
                var o = r.n(a);
                var c = r("66d806c6556aa1e94936");
                var i = r.n(c);
                var u = r("f7979a2d28599d65372d");
                var s = Object(c["hot"])(e)(u["a"]);
                o.a.render(n.a.createElement(s, null), document.getElementById("root"))
            }
        ).call(this, r("f48b41f3ce9a574655af")(e))
    },
    "38eb9cc333ee91a6129c": function(e, t) {
        e.exports = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAFIAAABSCAMAAADw8nOpAAAAnFBMVEVHcEz////////////////////////79/X////j4OD///////8gERT///8jERUcEBUuERUXDxUoERU0EBX3VQv3TA08DhbpLhn6YAj86+T8cQXb1db8gQPkSQ1aTE/IwMG5q63iOwyHfH9JLDDUOg9KHBo4IydvIhN3YGO0nZq3PhNIOT792sqUKRD8xrH1YzD3fEr3oGnYYAj7q42tRjivAAAADHRSTlMAUBCvIIBg/NvscDg5XOXBAAAFbElEQVRYw7XZCXeiPBQGYLtYPC0gskopi4i44lL//3/7sucmEPVrp9eemTMd+vS9CQaIo9Ff18f7y8Ry/ldZk5f3D5P3PJ44P6zJ2+uQOLacX9Rk3ANfJ84va6IFfbOcX5f1pIjOP6k3KT6Jb5Z10c5ITe8XOc4t6qxvvvKuy2I24NlKTW3NRFWUvHc+nmxmrJofMjNwghUmV2uWasLOHvqvbP+Yp6Gs9qz7MTh9LCHO7oJQFSbNaT3L2a61jPwnXb1MaC1j0pClKhpBBVVzlnw0PyheDImqtM1YVa07lLOg0mj0TucGirrnk7IrfvZle5FVyVky8gWM5GBCn5cgrcLVzCkYTT6UBRftKQR9WFtB1nBMJVkwks69q7Y9JErSqXzXHei9ZSQ9SG1bAz3651asVlbr91As6KSeEeTzyGubipg1J91bpCkjwlgBMit8hioxTSQEpYfJg1yny3bIVMkhEY0hEIPtQcZ0yiU02RQppNY2D6nU9gBiorW1ZxpIKcKmA/xCJIjppHUrTQMpQuoREYdrm6aK6WToqtL6IOYg2RMZx8g0Va+kVllWe2nqpNa2zMhVQqa9q3NhJqcgJB9HkDEIKkKmAyQ3TSnVvkHRlL2YZtLuhwQZY/RVWSl+6WYhJ/0G6SsRYwLGiESVWjp6l5Qh4TgiEKe0nNRBbFZUVp+0h0h9JDG23KAqYpLSsRyCZst2U0rSN5CgbxFyvynp/JarmJDsZMyW6Hftt5n1MElDBkt5+yRJC38REh1Q1FWZZYK076QMNhlYIpSUlAzw70XHsjf6bRKnbGs4q0OkR03fv0eyCW8r5TwxpqSmmZQT3lbq2Wwm/QdJr7YeJWHKqZn0Nvrb+Mcp+Vi2pbIkovrxWPKYGwhW9abYbOKBGX80pauGLOlbMTaSt8ZSDmaRgisrWYeGSeW8vPnu8Wp5udqztc2Y0oekMSXoe8PE1W690kmPvyHvk/YylTcpbLVcz+chUlVyYHZMq/pe3p4FggzDMN+VgNRC3iYL2Tdf1rsQV3SUpKeFfJxk5g6DkUIOhjRcx33QOLtWxA0Ro4MgW0+d7zvX8aV4g1ctNdc5EU9y3tgZpIc03RO1YjVPC9xf4HU05EXE34KMj5AeeHAgaYKckuBOPRAZFXFquBkEq4ZFRnNNxNNVfvvUKRlvkCzmMgNm663I5JyOci05JskahhSk6V7d9mt4/7jZnean0+UAvndZJM1aHcihu2D2tEdygph4kg7q/bRzCKMkaVb9tm89pNS3No2scxTmyIzvPqTAZyn1SqHVFc1VuEiSrveIYn7gc213nxnFIz6fIhQzWVPRvfEMycygW90wjycs4s5R61hEh091ko5by82ga3a+bReDpkXFKMSdo5goY/C9i7noDj7iu92iada496o/R+k1ihhJY6JxXH3Pd4H2iK9sRPjdIlnk5NjlRX92OJ6jSJgJjonJcD7f+TNlI+Id7OnY3QJVs9gF6Ld/fl6OUk0pmDAxbHBM356SFZ+ZGSPhps56QclFh0VU58v1iE714/V8CqMIhCRksp5O0fKM0M6Gmzpg62nVLHJMor++vr6IOUcV0uuEIpLBTDrXy+n/ruHWE9/Fq2fxLs+RmTfNNwIxOSckQwfIJl5TMcxXfIMMb2E+s/23PRJzLvKQNKZiMiMhnaO+o5AMBNvGozvCLGZ6wmROQZryk6fkQfFPR4CkIq5zCjYb5ZboBYNUZOSc986Chlyc5wmaoGbBo1/ULVG5cXs4s4y8cTGa2BSqHMucRTzoG7dgw/pwuagm713MPKsFF0/nq1yfn/5yE3w0evoXW/Vvf/2BAj6XfoVa46GPUl7ffv7hzPhZMP8BipOeEL/lB5EAAAAASUVORK5CYII="
    },
    "3bb5129593dfc30c97dc": function(e, t, r) {
        var n = r("b0c5ee3192f445a1763f");
        var a = n(false);
        a.push([e.i, "html{--min-width: 1200px;--max-width: 1300px;--header-height: 84px}html a.link,body a.link,#root a.link{text-decoration:none}input[type=number]::-webkit-inner-spin-button,input[type=number]::-webkit-outer-spin-button{-webkit-appearance:none;margin:0}input[type=number]{-moz-appearance:textfield}body.mounted{opacity:1}.semi-modal-confirm .semi-modal-body .semi-modal-confirm-title-text{font-size:18px}.semi-modal-title{font-size:18px;margin-top:0;text-align:left}.semi-modal-close{right:22px;top:22px}.error-boundary{margin:0;height:100%;display:-moz-box;display:-ms-flexbox;display:flex;-moz-box-orient:vertical;-moz-box-direction:normal;-ms-flex-direction:column;flex-direction:column;-moz-box-pack:center;-ms-flex-pack:center;justify-content:center;-moz-box-align:center;-ms-flex-align:center;align-items:center;color:var(--color-text-2);-webkit-user-select:none;-moz-user-select:none;-ms-user-select:none;user-select:none}.error-boundary pre:nth-last-of-type(1){color:white;-webkit-user-select:all;-moz-user-select:all;-ms-user-select:all;user-select:all;max-width:80vw;white-space:pre-wrap}.error-boundary pre:nth-last-of-type(1)::-moz-selection{background:var(--color-text-0)}.error-boundary pre:nth-last-of-type(1)::selection{background:var(--color-text-0)}\n", ""]);
        e.exports = a
    },
    "500aa1333462095cfbfa": function(e, t, r) {
        "use strict";
        r.d(t, "d", (function() {
                return D
            }
        ));
        r.d(t, "c", (function() {
                return M
            }
        ));
        r.d(t, "e", (function() {
                return U
            }
        ));
        r.d(t, "a", (function() {
                return q
            }
        ));
        var n = r("cbdf0739294e689c1b30");
        var a = r.n(n);
        var o = r("dd9ec66d1326640ab372");
        var c = r.n(o);
        var i = r("4ad597b35ebd8057c3b2");
        var u = r.n(i);
        var s = r("c4ff96cd159000886ec3");
        var l = r.n(s);
        var f = r("baed0ab21797785a14ed");
        var d = r.n(f);
        var p = r("bc45805c8d629bd7e372");
        var b = r.n(p);
        var m = r("a8f8d006001b3a82c300");
        var v = r.n(m);
        var h = r("af47c0de0114689236f0");
        var y = r.n(h);
        var g = r("f9ac2da030989e463e0d");
        var O = r.n(g);
        var w = r("36a0461b5c940c57f431");
        var _ = r.n(w);
        var E = r("242d31511f36a8fa77df");
        var x = r.n(E);
        var j = r("bd183afcc37eabd79225");
        var k = r.n(j);
        function S(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function P(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    S(Object(r), true).forEach((function(t) {
                            A(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    S(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function A(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        var C = k.a.create({
            baseURL: "/aweme/v1/creator/"
        });
        C.defaults.timeout = 5e3;
        C.interceptors.request.use((function(e) {
                var t = e.params;
                return P(P({}, e), {}, {
                    params: P({
                        aid: "2906",
                        app_name: "aweme_creator_platform",
                        device_platform: "web",
                        referer: document.referrer,
                        user_agent: navigator.userAgent,
                        cookie_enabled: navigator.cookieEnabled,
                        screen_width: screen.width,
                        screen_height: screen.height,
                        browser_language: navigator.language,
                        browser_platform: navigator.platform,
                        browser_name: navigator.appCodeName,
                        browser_version: navigator.appVersion,
                        browser_online: navigator.onLine,
                        timezone_name: Intl.DateTimeFormat().resolvedOptions().timeZone
                    }, t)
                })
            }
        ));
        C.interceptors.request.use((function(e) {
                return e
            }
        ), (function(e) {
                return Promise.reject(e)
            }
        ));
        C.interceptors.response.use((function(e) {
                var t = e.status
                    , r = e.data;
                if (200 === t)
                    return r;
                else
                    Promise.reject(r)
            }
        ), (function(e) {
                return Promise.reject(e)
            }
        ));
        t["b"] = C;
        var I = k.a.create({
            baseURL: "/aweme/v1/"
        });
        var T = k.a.create({
            baseURL: "/aweme/v2/school/"
        });
        var R = k.a.create({
            baseURL: "/passport/web/"
        });
        var N = k.a.create({
            baseURL: "/user/vcd/"
        });
        var D = I;
        var L = T;
        var M = R;
        var U = N;
        var q = k.a.create({
            baseURL: "https://www.douyin.com/",
            timeout: 1e4,
            withCredentials: true
        })
    },
    "63fd2a486193ef1e55ef": function(e, t, r) {
        var n = r("aeb79101789bee463010");
        if ("string" === typeof n)
            n = [[e.i, n, ""]];
        var a;
        var o;
        var c = {
            hmr: true
        };
        c.transform = a;
        c.insertInto = void 0;
        var i = r("586b245799805ba7ce84")(n, c);
        if (n.locals)
            e.exports = n.locals;
        if (false)
            ;
    },
    "6b5026880925b775c9d1": function(e, t, r) {
        var n = r("b0c5ee3192f445a1763f");
        var a = n(false);
        a.push([e.i, "body #pc_slide{margin:10px auto;height:auto !important;border:none !important}body #pc_slide #verify-bar-box{background:#fff}body #pc_slide #img-close{display:none}\n", ""]);
        e.exports = a
    },
    "73342a90a6c5104ac5f9": function(e, t, r) {
        var n = r("6b5026880925b775c9d1");
        if ("string" === typeof n)
            n = [[e.i, n, ""]];
        var a;
        var o;
        var c = {
            hmr: true
        };
        c.transform = a;
        c.insertInto = void 0;
        var i = r("586b245799805ba7ce84")(n, c);
        if (n.locals)
            e.exports = n.locals;
        if (false)
            ;
    },
    "7a4d57deb770ff98ef68": function(e, t, r) {
        var n = r("3bb5129593dfc30c97dc");
        if ("string" === typeof n)
            n = [[e.i, n, ""]];
        var a;
        var o;
        var c = {
            hmr: true
        };
        c.transform = a;
        c.insertInto = void 0;
        var i = r("586b245799805ba7ce84")(n, c);
        if (n.locals)
            e.exports = n.locals;
        if (false)
            ;
    },
    "7ee77fab2e2688fdceb3": function(e, t, r) {
        "use strict";
        var n = r("275976081ce1abf67779");
        var a = r.n(n);
        t["a"] = Object(n["createContext"])({})
    },
    "8d3bf0c15b5acded9951": function(e, t) {
        e.exports = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEwAAABMCAMAAADwSaEZAAAAq1BMVEVHcEz////////////////////////49PX///////8AAAD/////AE8A9+8A2dLe/v2gADH/gKceIiIaAAj/3ujf399EEB6V/Pn/m7om+fLeA0b/JGg8ABJnAB9B+vRQUFAAfHf/Q33FAD0Avrf/vtK6/fsAZGAAnZhucXGQkZL+ZJMAREFn9vJGPUCAACe/v7/vAEqQCjHv/v7/EFoAi4YwfnsQi4ffUHxAu7c9LFlTAAAACnRSTlMAcGBfed+/+oAgXFJGkAAAAy1JREFUWMPN2Fl3qjAQAGBsj1JDqgh4ZBdQoCJ1aXuX///LLgghARIMy8OdJ6v0O5MhC4wg5LFcSKNivhSqeBWlkSHOkPUmTRBvTUs87RXQI5T9aVvXZuiv6ycYEPuKe81qX9ZL/AAD4wMJS2FefNoqYHAoZXILgdM6rFBEbK3AxKd5rd5RxLTcRAJ7Xq8MM9xH/LVtu103jG0BD7Yu6wyh3P79WmH7Ppjk07BPhHEkRmI6NCkXiCV26oeJ31BtX3AqsX0/LEvNoayEElN6YpIfUGZHiYG+mGa1JwcYikmaOSEmafcJMcqiGYNJqadMh0np+yH2ookwaZ1mm8hUWM6502FZwCCYEIPwf8GO2oTYxtIHYOfoJ45jr4VBSHI8WHRApxEFC2Rro3FjHj7ZaJgMbBNaiZ6X7xl2RlS6W6+zq802BoATQBRdmFdIbkiOo4UBoCKvA4trVAeWh+1cZJmNFdauebsYWPcW9BijsZamwM4tawS2alnDsccgsSVq2UwaiikrsvZHP7/v30OxKB8kOsx8SM7K/tgBJ6ZZD+Z2dxxHHoIpRGIPy7RpWxAfdsaJ6bklq2A45uFbmScmMzZHPiyuRqnlY1RHYVn903JWZJgDxmJEydTRWFhlFgAwtmZhVbM6puB7k/BhPxXWetaPMJatjDsHFuGFqTcyyxftVzVrLjzLiVjl/k1tbJnErHH41maKH6id+iDJG23z7RoufgT+U7cMYm2oXJtjVZg8TudixRZnO540MACcOy1xLEm71eGAzmP0fcIqGe0McMntf2eUlBFKxG6iAt7TKaydJjvXMIw0xLcFMqZshZHP857b1Vbx2YkptFfEOGRbm9y6MJol1JfX3+tOy2Q8GqKXV7H+9a8veiMqga0tk4gteuFvNIE+NYp1tDqtfdU9uDZ+sZMmVxzIbOuRGKNJoprJkRigXh7IAdP66GzfXCD0Ez2LTWKV5/rNYbeWyPYNpbGkyiYk43ZRn7aphAW7TeVUnnm3OVpe8yfNONXOQlV5m3HCC6rzlqdZ0m4CXdH/z+oNzG3/BqbIboeOb60Kwmx80/cFt5CX83HWomhH/wNfVQq5Ea1tuwAAAABJRU5ErkJggg=="
    },
    "8d6a31f6e5330c22a4e6": function(e, t, r) {
        e.exports = r.p + "svgs/refresh.31a48243.svg"
    },
    "8da2087b1702e53cfde8": function(e, t, r) {
        "use strict";
        r.d(t, "a", (function() {
                return s
            }
        ));
        r.d(t, "b", (function() {
                return l
            }
        ));
        r.d(t, "f", (function() {
                return f
            }
        ));
        r.d(t, "k", (function() {
                return d
            }
        ));
        r.d(t, "h", (function() {
                return p
            }
        ));
        r.d(t, "g", (function() {
                return b
            }
        ));
        r.d(t, "j", (function() {
                return m
            }
        ));
        r.d(t, "c", (function() {
                return v
            }
        ));
        r.d(t, "d", (function() {
                return h
            }
        ));
        r.d(t, "e", (function() {
                return y
            }
        ));
        r.d(t, "i", (function() {
                return g
            }
        ));
        r.d(t, "m", (function() {
                return O
            }
        ));
        r.d(t, "l", (function() {
                return w
            }
        ));
        r.d(t, "n", (function() {
                return _
            }
        ));
        var n = r("2142af01f72a50657d1a");
        var a = r.n(n);
        var o = r("2c0e7a644ece321d90ae");
        var c = r.n(o);
        var i = r("bfae4cb5dae79298d66c");
        var u = r.n(i);
        var s = "https://sf1-cdn-tos.huoshanstatic.com/obj/ies-fe-bee/bee_prod/biz_181/bee_prod_181_bee_publish_";
        var l = location.host.includes("boe");
        var f = ["huoshan", "hotsoon"].some((function(e) {
                return location.host.includes(e)
            }
        ));
        var d = [1104, 1105];
        var p = {
            org: 2,
            user: 1
        };
        var b = {
            douyin: l ? 2909 : 2906,
            hotsoon: 3119
        };
        var m = {
            AUDITING: 0,
            APPROVED: 1,
            REJECTED: 2,
            EDITING: 5
        };
        var v = {
            AUDITING: 0,
            APPROVED: 1,
            REJECTED: 2,
            WITHDRAW: 3,
            WAITING: 4,
            WITHDRAW_REJECTED: 5
        };
        var h = {
            0: "\u5f85\u63a5\u53d7",
            1: "\u5df2\u63a5\u53d7",
            2: "\u5df2\u62d2\u7edd",
            3: "\u64a4\u9500\u901a\u8fc7",
            4: "\u64a4\u9500\u7533\u8bf7\u4e2d",
            5: "\u64a4\u9500\u9a73\u56de"
        };
        var y = {
            "\u4e2d\u56fd": "+86",
            "\u4e2d\u56fd\u9999\u6e2f": "+852",
            "\u4e2d\u56fd\u6fb3\u95e8": "+853",
            "\u4e2d\u56fd\u53f0\u6e7e": "+886",
            "\u963f\u5c14\u5df4\u5c3c\u4e9a": "+355",
            "\u963f\u5c14\u53ca\u5229\u4e9a": "+213",
            "\u963f\u5bcc\u6c57": "+93",
            "\u963f\u6839\u5ef7": "+54",
            "\u963f\u62c9\u4f2f\u8054\u5408\u914b\u957f\u56fd": "+971",
            "\u963f\u9c81\u5df4\u5c9b": "+297",
            "\u963f\u66fc": "+968",
            "\u963f\u585e\u62dc\u7586": "+994",
            "\u963f\u68ee\u677e": "+247",
            "\u57c3\u53ca": "+20",
            "\u57c3\u585e\u4fc4\u6bd4\u4e9a": "+251",
            "\u7231\u5c14\u5170": "+353",
            "\u7231\u6c99\u5c3c\u4e9a": "+372",
            "\u5b89\u9053\u5c14": "+376",
            "\u5b89\u54e5\u62c9": "+244",
            "\u5b89\u572d\u62c9": "+1264",
            "\u5b89\u63d0\u74dc\u548c\u5df4\u5e03\u8fbe": "+1268",
            "\u5965\u5730\u5229": "+43",
            "\u5965\u5170\u7fa4\u5c9b": "+358",
            "\u6fb3\u5927\u5229\u4e9a": "+61",
            "\u5df4\u5df4\u591a\u65af": "+1246",
            "\u5df4\u54c8\u9a6c": "+1242",
            "\u5df4\u57fa\u65af\u5766": "+92",
            "\u5df4\u62c9\u572d": "+595",
            "\u5df4\u52d2\u65af\u5766": "+970",
            "\u5df4\u6797": "+973",
            "\u5df4\u62ff\u9a6c": "+507",
            "\u5df4\u897f": "+55",
            "\u767d\u4fc4\u7f57\u65af": "+375",
            "\u767e\u6155\u5927": "+1441",
            "\u4fdd\u52a0\u5229\u4e9a": "+359",
            "\u5317\u9a6c\u91cc\u4e9a\u7eb3\u7fa4\u5c9b": "+1670",
            "\u8d1d\u5b81": "+229",
            "\u6bd4\u5229\u65f6": "+32",
            "\u51b0\u5c9b": "+354",
            "\u6ce2\u591a\u9ece\u5404": "+1787",
            "\u6ce2\u9ed1": "+387",
            "\u6ce2\u5170": "+48",
            "\u73bb\u5229\u7ef4\u4e9a": "+591",
            "\u4f2f\u5229\u5179": "+501",
            "\u535a\u8328\u74e6\u7eb3": "+267",
            "\u4e0d\u4e39": "+975",
            "\u5e03\u57fa\u62c9\u6cd5\u7d22": "+226",
            "\u5e03\u9686\u8fea": "+257",
            "\u671d\u9c9c": "+850",
            "\u8d64\u9053\u51e0\u5185\u4e9a": "+240",
            "\u4e39\u9ea6": "+45",
            "\u5fb7\u56fd": "+49",
            "\u591a\u54e5": "+228",
            "\u591a\u7c73\u5c3c\u52a0\u5171\u548c\u56fd": "+1809",
            "\u591a\u7c73\u5c3c\u514b": "+1767",
            "\u4fc4\u7f57\u65af": "+7",
            "\u5384\u74dc\u591a\u5c14": "+593",
            "\u5384\u7acb\u7279\u91cc\u4e9a": "+291",
            "\u6cd5\u56fd": "+33",
            "\u6cd5\u7f57\u7fa4\u5c9b": "+298",
            "\u6cd5\u5c5e\u6ce2\u5229\u5c3c\u897f\u4e9a": "+689",
            "\u6cd5\u5c5e\u572d\u4e9a\u90a3": "+594",
            "\u68b5\u8482\u5188": "+379",
            "\u83f2\u5f8b\u5bbe": "+63",
            "\u6590\u6d4e": "+679",
            "\u82ac\u5170": "+358",
            "\u4f5b\u5f97\u89d2": "+238",
            "\u5188\u6bd4\u4e9a": "+220",
            "\u521a\u679c": "+242",
            "\u521a\u679c\u6c11\u4e3b\u5171\u548c\u56fd": "+243",
            "\u54e5\u4f26\u6bd4\u4e9a": "+57",
            "\u54e5\u65af\u8fbe\u9ece\u52a0": "+506",
            "\u683c\u6797\u7eb3\u8fbe": "+1473",
            "\u683c\u9675\u5170\u5c9b": "+299",
            "\u683c\u9c81\u5409\u4e9a": "+995",
            "\u74dc\u5fb7\u7f57\u666e": "+590",
            "\u5173\u5c9b": "+1671",
            "\u572d\u4e9a\u90a3": "+592",
            "\u6d77\u5730": "+509",
            "\u97e9\u56fd": "+82",
            "\u8377\u5170": "+31",
            "\u8377\u5c5e\u5b89\u7684\u5217\u65af": "+599",
            "\u6d2a\u90fd\u62c9\u65af": "+504",
            "\u5409\u5e03\u63d0": "+253",
            "\u52a0\u62ff\u5927": "+1",
            "\u5409\u5c14\u5409\u65af\u65af\u5766": "+996",
            "\u51e0\u5185\u4e9a": "+224",
            "\u52a0\u90a3\u5229\u7fa4\u5c9b": "+3491",
            "\u52a0\u7eb3": "+233",
            "\u52a0\u84ec": "+241",
            "\u67ec\u57d4\u5be8": "+855",
            "\u6377\u514b": "+420",
            "\u5580\u9ea6\u9686": "+237",
            "\u5361\u5854\u5c14": "+974",
            "\u5f00\u66fc\u7fa4\u5c9b": "+1345",
            "\u79d1\u6469\u7f57": "+269",
            "\u79d1\u7d22\u6c83": "+883",
            "\u79d1\u7279\u8fea\u74e6": "+225",
            "\u79d1\u5a01\u7279": "+965",
            "\u514b\u7f57\u5730\u4e9a": "+385",
            "\u80af\u5c3c\u4e9a": "+254",
            "\u62c9\u8131\u7ef4\u4e9a": "+371",
            "\u83b1\u7d22\u6258": "+266",
            "\u8001\u631d": "+856",
            "\u9ece\u5df4\u5ae9": "+961",
            "\u7acb\u9676\u5b9b": "+370",
            "\u5229\u6bd4\u91cc\u4e9a": "+231",
            "\u5229\u6bd4\u4e9a": "+218",
            "\u5217\u652f\u6566\u58eb\u767b": "+423",
            "\u7559\u5c3c\u65fa\u5c9b": "+262",
            "\u5362\u68ee\u5821": "+352",
            "\u5362\u65fa\u8fbe": "+250",
            "\u7f57\u9a6c\u5c3c\u4e9a": "+40",
            "\u9a6c\u8fbe\u52a0\u65af\u52a0": "+261",
            "\u9a6c\u5c14\u4ee3\u592b": "+960",
            "\u9a6c\u8033\u4ed6": "+356",
            "\u9a6c\u62c9\u7ef4": "+265",
            "\u9a6c\u6765\u897f\u4e9a": "+60",
            "\u9a6c\u91cc": "+223",
            "\u9a6c\u5176\u987f": "+389",
            "\u9a6c\u7ecd\u5c14\u7fa4\u5c9b": "+692",
            "\u9a6c\u63d0\u5c3c\u514b": "+596",
            "\u9a6c\u7ea6\u7279": "+262",
            "\u6bdb\u91cc\u6c42\u65af": "+230",
            "\u6bdb\u91cc\u5854\u5c3c\u4e9a": "+222",
            "\u7f8e\u56fd": "+1",
            "\u7f8e\u5c5e\u8428\u6469\u4e9a": "+1684",
            "\u7f8e\u5c5e\u7ef4\u5c14\u4eac\u7fa4\u5c9b": "+1340",
            "\u8499\u53e4": "+976",
            "\u8499\u585e\u62c9\u7279\u5c9b": "+1664",
            "\u8499\u7279\u5167\u54e5\u7f85": "+382",
            "\u5b5f\u52a0\u62c9\u56fd": "+880",
            "\u79d8\u9c81": "+51",
            "\u5bc6\u514b\u7f57\u5c3c\u897f\u4e9a\u8054\u90a6": "+691",
            "\u7f05\u7538": "+95",
            "\u6469\u5c14\u591a\u74e6": "+373",
            "\u6469\u6d1b\u54e5": "+212",
            "\u6469\u7eb3\u54e5": "+377",
            "\u83ab\u6851\u6bd4\u514b": "+258",
            "\u58a8\u897f\u54e5": "+52",
            "\u7eb3\u7c73\u6bd4\u4e9a": "+264",
            "\u5357\u975e": "+27",
            "\u5c3c\u6cca\u5c14": "+977",
            "\u5c3c\u52a0\u62c9\u74dc": "+505",
            "\u5c3c\u65e5\u5c14": "+227",
            "\u5c3c\u65e5\u5229\u4e9a": "+234",
            "\u632a\u5a01": "+47",
            "\u5e15\u52b3": "+680",
            "\u8461\u8404\u7259": "+351",
            "\u5343\u91cc\u8fbe\u53ca\u6258\u5df4\u54e5": "+1868",
            "\u65e5\u672c": "+81",
            "\u745e\u5178": "+46",
            "\u745e\u58eb": "+41",
            "\u5723\u8bde\u5c9b": "+61",
            "\u5723\u57fa\u8328\u548c\u5c3c\u7ef4\u65af": "+1869",
            "\u5723\u5362\u897f\u4e9a": "+1758",
            "\u5723\u9a6c\u529b\u8bfa": "+223",
            "\u5723\u76ae\u57c3\u5c14\u548c\u5bc6\u514b\u9686\u7fa4\u5c9b": "+508",
            "\u5723\u6587\u68ee\u7279\u548c\u683c\u6797\u7eb3\u4e01\u65af": "+1784",
            "\u65af\u91cc\u5170\u5361": "+94",
            "\u65af\u6d1b\u4f10\u514b": "+421",
            "\u65af\u6d1b\u6587\u5c3c\u4e9a": "+386",
            "\u65af\u5a01\u58eb\u5170": "+268",
            "\u82cf\u4e39": "+249",
            "\u82cf\u91cc\u5357": "+597",
            "\u7d22\u9a6c\u91cc": "+252",
            "\u5854\u5409\u514b\u65af\u5766": "+992",
            "\u6c99\u7279\u963f\u62c9\u4f2f": "+966",
            "\u585e\u820c\u5c14": "+248",
            "\u585e\u6d66\u8def\u65af": "+357",
            "\u585e\u5185\u52a0\u5c14": "+221",
            "\u585e\u62c9\u5229\u6602": "+232",
            "\u585e\u5c14\u7ef4\u4e9a": "+381",
            "\u8428\u5c14\u74e6\u591a": "+503",
            "\u6c64\u52a0": "+676",
            "\u5766\u6851\u5c3c\u4e9a": "+255",
            "\u6cf0\u56fd": "+66",
            "\u7279\u514b\u65af\u548c\u51ef\u79d1\u65af\u7fa4\u5c9b": "+1649",
            "\u7a81\u5c3c\u65af": "+216",
            "\u56fe\u74e6\u5362": "+688",
            "\u571f\u8033\u5176": "+90",
            "\u571f\u5e93\u66fc\u65af\u5766": "+993",
            "\u4e4c\u5179\u522b\u514b\u65af\u5766": "+998",
            "\u4e4c\u62c9\u572d": "+598",
            "\u4e4c\u514b\u5170": "+380",
            "\u4e4c\u5e72\u8fbe": "+256",
            "\u6587\u83b1": "+673",
            "\u59d4\u5185\u745e\u62c9": "+58",
            "\u5371\u5730\u9a6c\u62c9": "+502",
            "\u53d9\u5229\u4e9a": "+963",
            "\u5308\u7259\u5229": "+36",
            "\u65b0\u897f\u5170": "+64",
            "\u65b0\u5580\u91cc\u591a\u5c3c\u4e9a": "+687",
            "\u65b0\u52a0\u5761": "+65",
            "\u5e0c\u814a": "+30",
            "\u897f\u8428\u6469\u4e9a": "+685",
            "\u897f\u6492\u54c8\u62c9": "+212",
            "\u897f\u73ed\u7259": "+34",
            "\u8d8a\u5357": "+84",
            "\u7ea6\u65e6": "+962",
            "\u82f1\u5c5e\u7ef4\u5c14\u4eac\u7fa4\u5c9b": "+1284",
            "\u82f1\u56fd": "+44",
            "\u5370\u5ea6\u5c3c\u897f\u4e9a": "+62",
            "\u5370\u5ea6": "+91",
            "\u610f\u5927\u5229": "+39",
            "\u4ee5\u8272\u5217": "+972",
            "\u4f0a\u6717": "+98",
            "\u4f0a\u62c9\u514b": "+964",
            "\u4e5f\u95e8": "+967",
            "\u4e9a\u7f8e\u5c3c\u4e9a": "+374",
            "\u7259\u4e70\u52a0": "+1876",
            "\u4e2d\u975e": "+236",
            "\u667a\u5229": "+56",
            "\u76f4\u5e03\u7f57\u9640": "+350",
            "\u4e4d\u5f97": "+235",
            "\u8d5e\u6bd4\u4e9a": "+260"
        };
        var g = {
            MEMBER_CONFIG: "member_config",
            DATA_CONFIG: "data_config",
            VIDEO_CONFIG: "video_config"
        };
        var O = {
            DESC: "DESC",
            ASC: "ASC",
            NONE: "NONE"
        };
        var w = {
            DEFAULT: 1,
            FOLLOW_ALL: 2,
            ADD_FOLLOW: 3,
            PUBLISH: 4,
            PLAY: 5,
            LIKE: 6,
            COMMENT: 7
        };
        var _ = {
            DEFAULT: 1,
            PLAY: 2,
            LIKE: 3,
            FOLLOW: 4,
            COMMENT: 5
        }
    },
    "936930c0a95d45899072": function(e, t, r) {
        var n = r("b0c5ee3192f445a1763f");
        var a = n(false);
        a.push([e.i, ".agreement{display:-moz-box;display:-ms-flexbox;display:flex;-ms-flex-wrap:wrap;flex-wrap:wrap;-moz-box-pack:center;-ms-flex-pack:center;justify-content:center;color:var(--color-text-2);margin:12px -2px;width:-webkit-fit-content;width:-moz-fit-content;width:fit-content}.agreement img{cursor:pointer}.agreement a{margin:0 2px}.agreement a.link{cursor:pointer;color:#04498d}\n", ""]);
        e.exports = a
    },
    "9ae88e6caac879637e63": function(e, t, r) {
        "use strict";
        r.r(t);
        var n = r("36a0461b5c940c57f431");
        var a = r.n(n);
        var o = r("242d31511f36a8fa77df");
        var c = r.n(o);
        var i = r("8d749f11ec9298f0ed09");
        var u = r.n(i);
        r("c19b0c3563b16bedf694");
        r("e0e3cd8cc5b7d3495eef");
        r("948c8c91206c63ce5834");
        r("bbb61e2c92096455e2dd");
        r("1a01c2cd55d37ef54066");
        r("1a8309099560b509be18");
        Object.assign = r("83406643bfb209d249f4");
        r("44638868b2286dc1d52e");
        r("4ddc57e2c1441e85127e");
        r("73335c8407f708998cda");
        r("2e1567429e1988dd263d");
        if ("undefined" === typeof Promise) {
            r("9fbadcf2c98d054f245e").enable();
            window.Promise = r("e6615fafe8e11b81e37a")
        }
    },
    a418510e8fecea13f334: function(e, t, r) {
        var n = r("936930c0a95d45899072");
        if ("string" === typeof n)
            n = [[e.i, n, ""]];
        var a;
        var o;
        var c = {
            hmr: true
        };
        c.transform = a;
        c.insertInto = void 0;
        var i = r("586b245799805ba7ce84")(n, c);
        if (n.locals)
            e.exports = n.locals;
        if (false)
            ;
    },
    aeb79101789bee463010: function(e, t, r) {
        var n = r("b0c5ee3192f445a1763f");
        var a = n(false);
        a.push([e.i, ".account-modal .semi-modal-header{display:none}.account-modal .semi-modal-body{min-height:407px;padding-bottom:24px}.account-modal .semi-modal-body .semi-tabs-bar{display:-moz-box;display:-ms-flexbox;display:flex;-ms-flex-pack:distribute;justify-content:space-around}.account-modal .semi-modal-body .semi-tabs-bar .semi-tabs-tab{width:50%;padding:26px 0 19px;text-align:center;margin:0;font-size:18px}.account-modal .semi-modal-body .semi-tabs-content{padding:14px 0 0}.account-modal .semi-modal-body .semi-input-large,.account-modal .semi-modal-body .semi-button-size-large{height:64px;font-size:16px}.account-modal .semi-modal-body .semi-button-size-large{margin-top:20px}.account-modal .semi-modal-body .semi-input-prefix{padding-left:0}.account-modal .semi-modal-body .semi-input-prefix .semi-select-selection-text{font-weight:500;font-size:16px}.account-modal .semi-modal-body .semi-input-prefix .semi-select{height:64px;background:transparent;border:none}.account-modal .semi-modal-body .semi-input-suffix{font-size:16px;margin-right:16px;white-space:nowrap;font-weight:500}.account-modal .semi-modal-body .agreement{width:230px;margin:12px auto 0}.account-modal .account-qrcode{text-align:center}.account-modal .account-qrcode .qrcode-image{width:200px;height:200px;margin:18px auto;position:relative}.account-modal .account-qrcode .qrcode-image>img,.account-modal .account-qrcode .qrcode-image>div{width:100%;height:100%}.account-modal .account-qrcode .qrcode-image>div{position:absolute;top:0;left:0;background:rgba(255,255,255,0.95);display:-moz-box;display:-ms-flexbox;display:flex;-moz-box-pack:center;-ms-flex-pack:center;justify-content:center;-moz-box-align:center;-ms-flex-align:center;align-items:center}.account-modal .account-qrcode .qrcode-image>div img{cursor:pointer}.account-modal .account-qrcode .qrcode-image .logo{width:38px;height:38px;position:absolute;top:50%;left:50%;-webkit-transform:translate(-50%, -50%);-moz-transform:translate(-50%, -50%);transform:translate(-50%, -50%);border-radius:8px;background:white;border:2px solid white}.account-modal .account-qrcode .qrcode-image .logo.hotsoon{border-width:.5px .5px 2px}.account-modal .account-qrcode p{margin-top:0;margin-bottom:28px;padding:0 20px;color:var(--color-text-2);font-size:16px;line-height:28px}.account-modal .account-qrcode p.error{color:var(--color-danger)}.account-modal .account-qrcode p em{color:var(--color-link);font-style:normal}.account-modal .account-qrcode .agreement{width:auto}.account-modal .account-phone .semi-input-suffix{-moz-box-pack:end;-ms-flex-pack:end;justify-content:flex-end;-ms-flex-negative:0;flex-shrink:0;margin-left:16px}.account-modal .account-phone .semi-input-suffix span{cursor:pointer}.account-modal .account-phone .semi-input-suffix span.send-code-timer{color:var(--color-text-2);cursor:default}.account-modal .account-phone p{color:var(--color-text-2)}.account-modal .account-phone p.err{height:14px;margin:0 0 2px;padding-right:130px;color:var(--color-danger)}.account-modal .account-phone .semi-button-disabled,.account-modal .account-phone .semi-button-disabled:hover{background:#c6cacd}.account-modal .account-phone .semi-form-field{padding-bottom:4px}.account-modal .account-phone .toggle{display:-moz-box;display:-ms-flexbox;display:flex;-moz-box-align:center;-ms-flex-align:center;align-items:center;width:-webkit-fit-content;width:-moz-fit-content;width:fit-content;float:right;margin-bottom:12px;cursor:pointer;color:rgba(40,40,40,0.5)}.account-modal.verify .semi-modal-header{display:block;height:66px;-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box;display:-moz-box;display:-ms-flexbox;display:flex;-moz-box-pack:center;-ms-flex-pack:center;justify-content:center;-moz-box-align:center;-ms-flex-align:center;align-items:center}.account-modal.verify .semi-modal-header .semi-button{display:-moz-box;display:-ms-flexbox;display:flex}.account-modal.verify .semi-modal-body{min-height:273px;margin-top:0}.account-modal.verify .semi-modal-body p{width:280px;margin:auto}.account-modal.verify .semi-modal-body p:not(.error){color:var(--color-text-1)}.account-modal.verify .qrcode-image{margin:0 auto 20px}.account-modal.bind .semi-modal-header{display:block}.account-modal.bind .semi-modal-body{min-height:0}.account-modal.type .semi-modal-header{display:block}.account-modal.type .semi-modal-body{min-height:0}.account-modal.type .semi-modal-body p{margin:0;color:#6b7075;font-size:14px;margin-bottom:10px}.account-modal.type .semi-modal-body .btn{width:100%;margin-top:26px;text-align:right}.account-modal.type .semi-modal-body .semi-button{margin:0}.phone-code-select{width:352px}.slider .semi-modal-confirm-content{margin:0 0 30px}.semi-modal-centered{top:50vh !important}\n", ""]);
        e.exports = a
    },
    af806d7aca2bb73a8ac4: function(e, t) {
        e.exports = SemiUI
    },
    b760f20d47b617d144e4: function(e, t, r) {
        e.exports = r.p + "svgs/uncheck.1809bf04.svg"
    },
    d93bd82914f844303b93: function(e, t, r) {
        e.exports = r.p + "svgs/check.90aefaa2.svg"
    },
    de7f2df91c095f0cd45d: function(e, t) {
        e.exports = ReactDOM
    },
    e0eb4ce3159968a79748: function(e, t, r) {
        "use strict";
        var n = r("cbdf0739294e689c1b30");
        var a = r.n(n);
        var o = r("dd9ec66d1326640ab372");
        var c = r.n(o);
        var i = r("4ad597b35ebd8057c3b2");
        var u = r.n(i);
        var s = r("c4ff96cd159000886ec3");
        var l = r.n(s);
        var f = r("a8f8d006001b3a82c300");
        var d = r.n(f);
        var p = r("af47c0de0114689236f0");
        var b = r.n(p);
        var m = r("f9ac2da030989e463e0d");
        var v = r.n(m);
        var h = r("8446bcb970ac73a9d7b7");
        var y = r.n(h);
        var g = r("7503101e3de10c4b9a7b");
        var O = r.n(g);
        var w = r("dafcc91707e0a69ef1b5");
        var _ = r.n(w);
        var E = r("5777342bebabc3495ab9");
        var x = r.n(E);
        var j = r("60310abcca4f44b14e57");
        var k = r.n(j);
        var S = r("242d31511f36a8fa77df");
        var P = r.n(S);
        var A = r("bc45805c8d629bd7e372");
        var C = r.n(A);
        var I = r("01307be32cc5492fccd7");
        var T = r.n(I);
        var R = r("baed0ab21797785a14ed");
        var N = r.n(R);
        var D = r("33711cb8a0e1011e6a3f");
        var L = r.n(D);
        var M = r("275976081ce1abf67779");
        var U = r.n(M);
        var q = r("af806d7aca2bb73a8ac4");
        var B = r.n(q);
        var H = r("bd183afcc37eabd79225");
        var F = r.n(H);
        var G = r("a418510e8fecea13f334");
        var z = r.n(G);
        var V = r("d93bd82914f844303b93");
        var J = r.n(V);
        var W = r("b760f20d47b617d144e4");
        var Y = r.n(W);
        var K = r("8da2087b1702e53cfde8");
        function Q(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function Z(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    Q(Object(r), true).forEach((function(t) {
                            X(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    Q(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function X(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        function $(e, t) {
            return ae(e) || ne(e, t) || te(e, t) || ee()
        }
        function ee() {
            throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
        }
        function te(e, t) {
            if (!e)
                return;
            if ("string" === typeof e)
                return re(e, t);
            var r = Object.prototype.toString.call(e).slice(8, -1);
            if ("Object" === r && e.constructor)
                r = e.constructor.name;
            if ("Map" === r || "Set" === r)
                return Array.from(e);
            if ("Arguments" === r || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))
                return re(e, t)
        }
        function re(e, t) {
            if (null == t || t > e.length)
                t = e.length;
            for (var r = 0, n = new Array(t); r < t; r++)
                n[r] = e[r];
            return n
        }
        function ne(e, t) {
            if ("undefined" === typeof Symbol || !(Symbol.iterator in Object(e)))
                return;
            var r = [];
            var n = true;
            var a = false;
            var o = void 0;
            try {
                for (var c = e[Symbol.iterator](), i; !(n = (i = c.next()).done); n = true) {
                    r.push(i.value);
                    if (t && r.length === t)
                        break
                }
            } catch (u) {
                a = true;
                o = u
            } finally {
                try {
                    if (!n && null != c["return"])
                        c["return"]()
                } finally {
                    if (a)
                        throw o
                }
            }
            return r
        }
        function ae(e) {
            if (Array.isArray(e))
                return e
        }
        t["a"] = function(e) {
            var t = e.className
                , r = e.agreed
                , n = void 0 === r ? false : r
                , a = e.checkState
                , o = e.agreeInfoPrefix
                , c = void 0 === o ? "\u767b\u5f55\uff0c\u5373\u8868\u660e\u4f60\u540c\u610f\u6211\u4eec\u7684" : o
                , i = e.agreementLink
                , u = e.privacyLink
                , s = e.modalLink
                , l = e.modalLinkText;
            var f = a || Object(M["useState"])(n)
                , d = $(f, 2)
                , p = d[0]
                , b = d[1];
            var m = Object(M["useState"])({
                visible: false,
                title: l
            })
                , v = $(m, 2)
                , h = v[0]
                , y = v[1];
            var g = function e() {
                if (!n)
                    b((function(e) {
                            return !e
                        }
                    ))
            };
            var O = function e(t) {
                F.a.get("".concat(K["a"]).concat(t)).then((function(e) {
                        var t = e.data;
                        y((function(e) {
                                return Z(Z({}, e), {}, {
                                    visible: true,
                                    width: 800,
                                    onOk: function e() {
                                        return y((function(e) {
                                                return Z(Z({}, e), {}, {
                                                    visible: false
                                                })
                                            }
                                        ))
                                    },
                                    onCancel: function e() {
                                        return y((function(e) {
                                                return Z(Z({}, e), {}, {
                                                    visible: false
                                                })
                                            }
                                        ))
                                    },
                                    children: U.a.createElement("div", {
                                        dangerouslySetInnerHTML: {
                                            __html: t
                                        }
                                    })
                                })
                            }
                        ))
                    }
                ))
            };
            return U.a.createElement("div", {
                className: "agreement".concat(t ? " ".concat(t) : "")
            }, U.a.createElement("img", {
                src: p ? J.a : Y.a,
                onClick: g
            }), n ? U.a.createElement("span", null, "\u6211\u5df2\u9605\u8bfb\u5e76\u540c\u610f") : U.a.createElement("div", null, c), s && l ? U.a.createElement(U.a.Fragment, null, U.a.createElement("a", {
                className: "link",
                onClick: function e() {
                    return O(s)
                }
            }, l), U.a.createElement(q["Modal"], h)) : U.a.createElement(U.a.Fragment, null, U.a.createElement("a", {
                className: "link",
                target: "_blank",
                rel: "noopener noreferrer",
                href: i || "//www.douyin.com/agreement/"
            }, "\u7528\u6237\u534f\u8bae"), "\u548c", U.a.createElement("a", {
                className: "link",
                target: "_blank",
                rel: "noopener noreferrer",
                href: u || "//www.douyin.com/privacy"
            }, "\u9690\u79c1\u653f\u7b56"), n && U.a.createElement("span", null, ",\u8bf7\u4ed4\u7ec6\u9605\u8bfb")))
        }
    },
    f7979a2d28599d65372d: function(e, t, r) {
        "use strict";
        var n = r("36a0461b5c940c57f431");
        var a = r("8d749f11ec9298f0ed09");
        var o = r("baed0ab21797785a14ed");
        var c = r("f90105b73394ddd1c1ce");
        var i = r("bc45805c8d629bd7e372");
        var u = r("a8f8d006001b3a82c300");
        var s = r("242d31511f36a8fa77df");
        var l = r("af47c0de0114689236f0");
        var f = r("dafcc91707e0a69ef1b5");
        var d = r("2d07b2a40f20ec52c656");
        var p = r("7d923aba6bcac07f74aa");
        var b = r("2c0e7a644ece321d90ae");
        var m = r("bfae4cb5dae79298d66c");
        var v = r("275976081ce1abf67779");
        var h = r.n(v);
        var y = r("a63b0d047588ea783f61");
        var g = r.n(y);
        var O = r("103be25b8913a0141218");
        var w = r("42993e61ed154f509e4c");
        var _ = r("8da2087b1702e53cfde8");
        var E = function e(t, r) {
            if (window[t]) {
                r(window[t]);
                return
            }
            var n = {};
            if (window.XMLHttpRequest)
                n = new XMLHttpRequest;
            else
                n = new ActiveXObject("Microsoft.XMLHTTP");
            n.open("GET", t);
            n.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
            n.send();
            n.onreadystatechange = function() {
                if (4 == n.readyState && 200 == n.status) {
                    var e = n.responseText;
                    try {
                        e = JSON.parse(e);
                        window[t] = {
                            data: e
                        };
                        r && r({
                            data: e
                        })
                    } catch (a) {
                        void 0
                    }
                }
            }
        };
        var x = E;
        var j = r("7ee77fab2e2688fdceb3");
        var k = r("cbdf0739294e689c1b30");
        var S = r("dd9ec66d1326640ab372");
        var P = r("4ad597b35ebd8057c3b2");
        var A = r("c4ff96cd159000886ec3");
        var C = r("f9ac2da030989e463e0d");
        var I = r("7503101e3de10c4b9a7b");
        var T = r("5777342bebabc3495ab9");
        var R = r("60310abcca4f44b14e57");
        var N = r("01307be32cc5492fccd7");
        var D = r("33711cb8a0e1011e6a3f");
        var L = r("ba6d8152ea0357825423");
        var M = r("fa05e5a41e4fc0f3d07f");
        var U = r("8446bcb970ac73a9d7b7");
        var q = r("5fb4a8bb2d828c746bf7");
        var B = r("8c3413d5484a39d78b8e");
        var H = r("6a7b77e0b23af5792ac3");
        var F = r("d87ad826655176a1ce21");
        var G = r("af806d7aca2bb73a8ac4");
        var z = r("500aa1333462095cfbfa");
        var V = r("a1c8a05eb561659fd014");
        var J = r.n(V);
        var W = r("c9dd95be3dee262a6ebc");
        var Y = r("ab01c31194414375af26");
        var K = r("bd183afcc37eabd79225");
        var Q = r.n(K);
        function Z(e) {
            "@babel/helpers - typeof";
            if ("function" === typeof Symbol && "symbol" === typeof Symbol.iterator)
                Z = function e(t) {
                    return typeof t
                }
                ;
            else
                Z = function e(t) {
                    return t && "function" === typeof Symbol && t.constructor === Symbol && t !== Symbol.prototype ? "symbol" : typeof t
                }
                ;
            return Z(e)
        }
        function X(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function $(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    X(Object(r), true).forEach((function(t) {
                            ee(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    X(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function ee(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        function te(e) {
            var t;
            var r = new RegExp("(^| )".concat(e, "=([^;]*)(;|$)"));
            if (t = document.cookie.match(r))
                return unescape(t[2]);
            return null
        }
        function re(e) {
            var t, r, n, a;
            for (n = 1,
                     a = arguments.length; n < a; n++) {
                t = arguments[n];
                for (r in t)
                    if (Object.prototype.hasOwnProperty.call(t, r))
                        e[r] = t[r]
            }
            return e
        }
        function ne(e) {
            var t = location
                , r = t.protocol
                , n = t.host;
            var a = n.includes("boe");
            var o = n.includes("huoshan") || n.includes("hotsoon");
            var c = e.baseURL
                , i = void 0 === c ? "//sso".concat(a ? "-boe" : "", ".").concat(o ? "huoshan" : "douyin", ".com") : c
                , u = e.url;
            var s = e.method || "GET";
            var l = te("s_v_web_id") || void 0;
            var f = te("ttwid") || void 0;
            var d = $($({}, e.data || {}), {}, {
                service: e.data.service || "".concat(r, "//").concat(n),
                is_vcd: 1,
                fp: l,
                ttwid: f
            });
            var p = e.headers || {};
            var b = e.callback || function(e) {
                    return e || {}
                }
            ;
            var m = this;
            var v = {user/i
                baseURL: i,
                url: u,
                method: s,
                headers: p,
                withCredentials: true
            };
            if ("GET" === s)
                v.params = d;
            else {
                v.headers["Content-Type"] = "application/x-www-form-urlencoded";
                var h = [];
                for (var y in d)
                    if (d.hasOwnProperty(y)) {
                        var g = d[y];
                        if (null != g)
                            if ("object" === Z(g))
                                h.push("".concat(encodeURIComponent(y), "=").concat(encodeURIComponent(JSON.stringify(g))));
                            else
                                h.push("".concat(encodeURIComponent(y), "=").concat(encodeURIComponent(g)))
                    }
                v.data = h.join("&")
            }
            return new Promise((function(e, t) {
                    Q()(v).then((function(r) {
                            var n = r.data
                                , a = void 0 === n ? {} : n;
                            if (Number.isNaN(a.error_code))
                                a = a.data;
                            else
                                a = $($({}, a), a.data || {});
                            if (0 === a.error_code)
                                e(re({
                                    status: "success"
                                }, b(a)));
                            else
                                t($($({}, a), {}, {
                                    message: a.description || a.message,
                                    "x-tt-logid": (p || {})["x-tt-logid"] || ""
                                }))
                        }
                    ))["catch"]((function(e) {
                            t({
                                status: "fail",
                                captcha: "",
                                message: m.failText || "\u7f51\u7edc\u8bf7\u6c42\u5931\u8d25\uff0c\u8bf7\u91cd\u8bd5",
                                error_code: -1,
                                "x-tt-logid": (e.response && e.response.headers || {})["x-tt-logid"] || ""
                            })
                        }
                    ))
                }
            ))
        }
        var ae = {
            getCookie: te,
            extend: re,
            request: ne
        };
        function oe(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function ce(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    oe(Object(r), true).forEach((function(t) {
                            ie(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    oe(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function ie(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        function ue(e) {
            "@babel/helpers - typeof";
            if ("function" === typeof Symbol && "symbol" === typeof Symbol.iterator)
                ue = function e(t) {
                    return typeof t
                }
                ;
            else
                ue = function e(t) {
                    return t && "function" === typeof Symbol && t.constructor === Symbol && t !== Symbol.prototype ? "symbol" : typeof t
                }
                ;
            return ue(e)
        }
        var se = {
            OAUTH: {
                BIND: "/passport/web/auth/bind/",
                BIND_V2: "/passport/web/auth/bind/v2/",
                UNBIND: "/passport/web/auth/unbind/",
                LOGIN: "/passport/web/auth/login/",
                LOGIN_ONLY: "/passport/web/auth/login_only/",
                REGISTER_WITH_PWD: "/passport/mobile/register/",
                USER_BIND: "/passport/mobile/user_auth_bind/",
                THIRD_MOBILE_BIND: "/passport/auth/bind_with_mobile_authorize/",
                THIRD_MOBILE_LOGIN: "/passport/auth/bind_with_mobile_login/",
                SWITCH_BIND: "/passport/auth/switch_bind/",
                SWITCH_BIND_WEB: "/passport/web/auth/switch_bind/"
            },
            ACCOUNT: {
                CHECK_LOGIN: "/check_login/",
                INFO: "/passport/web/account/info/",
                SET: "/passport/web/account/set/",
                MOBILE_CODE_LOGIN: "/quick_login/v2/",
                PWD_LOGIN: "/account_login/v2/",
                SSO_LOGIN: "/passport/web/web_login",
                SMS_LOGIN_ONLY: "/passport/mobile/sms_login_only/",
                SMS_LOGIN_CONTINUE: "/passport/mobile/sms_login_continue/",
                LOGOUT: "/passport/web/logout/?account_sdk_source=web",
                EMAIL_VERIFY: "/passport/web/email/register_verify/?account_sdk_source=web",
                MOBILE_BIND: "/bind_mobile/",
                EMAIL_BIND: "/passport/web/email/bind/",
                BIND_LOGIN: "/passport/web/mobile/bind_login/",
                ACCOUNT_SWITCH: "/passport/account/switch/",
                CONNECT_LOGIN: "/connect_login/"
            },
            PASSWORD: {
                EMAIL_RESET: "/passport/web/email/password/reset/?account_sdk_source=web",
                EMAIL_UPDATE: "/passport/web/password/update/?account_sdk_source=web",
                SET: "/passport/web/password/set/?account_sdk_source=web",
                CHANGE: "/passport/web/password/change/?account_sdk_source=web",
                RESET: "/passport/web/password/reset/?account_sdk_source=web"
            },
            CAPTCHA: {
                CODE: "/passport/web/refresh_captcha/?account_sdk_source=web"
            },
            CODE: {
                MOBILE: "/send_activation_code/v2/",
                EMAIL: "/passport/web/email/send_code/?account_sdk_source=web"
            },
            VOICE_CODE: {
                MOBILE: "/passport/mobile/send_voice_code/?account_sdk_source=web"
            },
            QRCODE: {
                GET: "/get_qrcode/",
                CHECK_STATUS: "/check_qrconnect/",
                SCAN: "/passport/web/scan_qrcode/",
                CONFIRM: "/passport/web/confirm_qrcode/"
            },
            MOBILE: {
                VALIDATE_CODE: "/passport/web/validate_code/",
                CHANGE: "/passport/web/mobile/change/"
            }
        };
        J.a.polyfill();
        function le(e) {
            var t = e || {};
            this.failText = t.failText;
            this._default = {
                aid: t.aid,
                language: t.language,
                multi_login: t.multi_login
            };
            var r = ae.request;
            var n = this;
            ae.request = function(e) {
                return r.call(n, e)
            }
        }
        function fe(e) {
            var t = [];
            var r;
            var n = 0;
            var a = 0;
            for (n = 0; n < e.length; n++) {
                r = e.charCodeAt(n);
                if (0 <= r && r <= 127)
                    t.push(r);
                else if (128 <= r && r <= 2047) {
                    t.push(192 | 31 & r >> 6);
                    t.push(128 | 63 & r)
                } else if (2048 <= r && r <= 55295 || 57344 <= r && r <= 65535) {
                    t.push(224 | 15 & r >> 12);
                    t.push(128 | 63 & r >> 6);
                    t.push(128 | 63 & r)
                }
            }
            for (a = 0; a < t.length; a++)
                t[a] &= 255;
            return t
        }
        le.prototype._paramsPack = function(e) {
            if ("object" !== ue(e))
                e = {};
            e = e || {};
            return ae.extend(e, this._default)
        }
        ;
        le.prototype.__encrypt = function(e) {
            var t, r;
            var n = [];
            var a = [];
            if (void 0 === e)
                return "";
            e = String(e);
            a = fe(e);
            for (t = 0,
                     r = a.length; t < r; ++t)
                n.push((5 ^ a[t]).toString(16));
            return n.join("")
        }
        ;
        le.prototype.__encryptParams = function(e, t) {
            var r, n, a;
            var o = 0;
            if ("object" !== ue(e))
                return e;
            e = ae.extend({}, e);
            t = t || [];
            for (r = 0,
                     n = t.length; r < n; ++r) {
                a = e[t[r]];
                if (void 0 !== a) {
                    o |= 1;
                    e[t[r]] = this.__encrypt(a)
                }
            }
            e.mix_mode = o;
            return e
        }
        ;
        le.prototype.checkLogin = function() {
            return ae.request({
                url: se.ACCOUNT.CHECK_LOGIN,
                method: "POST",
                data: {},
                headers: {
                    "X-CSRFToken": ae.getCookie("csrftoken")
                },
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.userLogin = function(e, t) {
            var r = t ? se.ACCOUNT.PWD_LOGIN : se.ACCOUNT.MOBILE_CODE_LOGIN;
            this._paramsPack(e);
            if (t)
                e = this.__encryptParams(e, ["account", "password"]);
            else
                e = this.__encryptParams(e, ["mobile", "code"]);
            return ae.request({
                baseURL: e.baseURL,
                url: r,
                method: "POST",
                data: e,
                headers: {
                    "X-CSRFToken": ae.getCookie("csrftoken")
                },
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.allLogin = function(e, t) {
            this._paramsPack(e);
            var r = "";
            switch (t) {
                case "pwd_login":
                    r = se.ACCOUNT.PWD_LOGIN;
                    e = this.__encryptParams(e, ["account", "password"]);
                    break;
                case "mobile_code_login":
                    r = se.ACCOUNT.MOBILE_CODE_LOGIN;
                    e = this.__encryptParams(e, ["mobile", "code"]);
                    break;
                case "sms_login_only":
                    r = se.ACCOUNT.SMS_LOGIN_ONLY;
                    e = this.__encryptParams(e, ["mobile", "code"]);
                    break;
                case "sms_login_continue":
                    r = se.ACCOUNT.SMS_LOGIN_CONTINUE;
                    e = this.__encryptParams(e, ["mobile", "code"]);
                    break;
                default:
                    return Promise.reject({
                        status: "fail",
                        message: "error login type",
                        captcha: "",
                        error_code: -1,
                        "x-tt-logid": ""
                    })
            }
            return ae.request({
                url: r,
                method: "POST",
                data: e,
                headers: {
                    "X-CSRFToken": ae.getCookie("csrftoken")
                },
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.accountSet = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["username", "password"]);
            return ae.request({
                url: se.ACCOUNT.SET,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.pwdSet = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.PASSWORD.SET,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.pwdChange = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["password", "code", "ticket"]);
            return ae.request({
                url: se.PASSWORD.CHANGE,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.pwdReset = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["code", "password", "mobile"]);
            return ae.request({
                url: se.PASSWORD.RESET,
                method: "POST",
                data: e,
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.getMobileCode = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.CODE.MOBILE,
                data: e
            })
        }
        ;
        le.prototype.bindMobile = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.ACCOUNT.MOBILE_BIND,
                data: e,
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.setCodeStatus = function(e, t) {
            var r = t || 60;
            var n = null;
            return new Promise((function(t) {
                    n = setInterval((function() {
                            if (0 === r) {
                                clearInterval(n);
                                n = null;
                                t();
                                return false
                            }
                            e.text = "".concat(r, "s\u91cd\u65b0\u83b7\u53d6");
                            r--
                        }
                    ), 1e3)
                }
            ))
        }
        ;
        le.prototype.refreshCaptcha = function() {
            return ae.request({
                url: se.CAPTCHA.CODE,
                callback: function e(t) {
                    return t.captcha ? {
                        captcha: t.captcha
                    } : {}
                }
            })
        }
        ;
        le.prototype.logout = function(e) {
            e = e || "/";
            e = encodeURIComponent(e);
            window.location.href = "".concat(se.ACCOUNT.LOGOUT, "&next=").concat(e)
        }
        ;
        le.prototype.emailPwdReset = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["email", "password", "code", "type"]);
            return ae.request({
                url: se.PASSWORD.EMAIL_RESET,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.emailPwdUpdate = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["current_password", "password"]);
            return ae.request({
                url: se.PASSWORD.EMAIL_UPDATE,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.getEmailCode = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["type", "email", "password"]);
            return ae.request({
                url: se.CODE.EMAIL,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.emailVerify = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.ACCOUNT.EMAIL_VERIFY,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.ssoLogin = function(e) {
            this._paramsPack(e);
            e.next = e.next || "/";
            window.location.href = "".concat(se.ACCOUNT.SSO_LOGIN, "?platform_app_id=").concat(e.platform_app_id, "&platform=").concat(e.platform, "&aid=").concat(e.aid, "&next=").concat(encodeURIComponent(e.next)).concat(e.action ? "&action=".concat(e.action) : "")
        }
        ;
        le.prototype.oauthBind = function(e, t) {
            this._paramsPack(e);
            return ae.request({
                url: t ? se.OAUTH.BIND_V2 : se.OAUTH.BIND,
                method: "POST",
                data: e,
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.oauthUnbind = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.OAUTH.UNBIND,
                data: e
            })
        }
        ;
        le.prototype.oauthLogin = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.OAUTH.LOGIN,
                method: "POST",
                data: e,
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.oauthLoginOnly = function(e) {
            this._paramsPack(e);
            if (e.extra_data && 1 == e.mix_mode)
                e = this.__encryptParams(e, ["extra_data"]);
            return ae.request({
                url: se.OAUTH.LOGIN_ONLY,
                method: "POST",
                data: e,
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.getAccountInfo = function() {
            var e = this._paramsPack();
            return ae.request({
                url: se.ACCOUNT.INFO,
                data: e,
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.bindEmail = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["email", "code"]);
            return ae.request({
                url: se.ACCOUNT.EMAIL_BIND,
                method: "POST",
                data: e,
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.bindLogin = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["mobile", "code"]);
            return ae.request({
                url: se.ACCOUNT.BIND_LOGIN,
                method: "POST",
                data: e,
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        }
        ;
        le.prototype.getQrcode = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.QRCODE.GET,
                data: e
            })
        }
        ;
        le.prototype.checkQrcodeStatus = function(e) {
            this._paramsPack(e);
            return ae.request({
                baseURL: e.baseURL,
                url: se.QRCODE.CHECK_STATUS,
                data: e
            }).then((function(e) {
                    var t = {
                        1: "new",
                        2: "scanned",
                        3: "confirmed"
                    };
                    return ce(ce({}, e), {}, {
                        status: t[e.status]
                    })
                }
            ))
        }
        ;
        le.prototype.scanQrcode = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.QRCODE.SCAN,
                data: e
            })
        }
        ;
        le.prototype.confirmQrcode = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.QRCODE.CONFIRM,
                data: e
            })
        }
        ;
        le.prototype.getMobileVoiceCode = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.VOICE_CODE.MOBILE,
                data: e
            })
        }
        ;
        le.prototype.userRegister = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["mobile", "code", "password"]);
            return ae.request({
                url: se.OAUTH.REGISTER_WITH_PWD,
                data: e,
                method: "POST"
            })
        }
        ;
        le.prototype.userOAuthBind = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.OAUTH.USER_BIND,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.thirdMobileBind = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["mobile", "sms_code"]);
            return ae.request({
                url: se.OAUTH.THIRD_MOBILE_BIND,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.thirdMobileLogin = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["mobile", "sms_code"]);
            return ae.request({
                url: se.OAUTH.THIRD_MOBILE_LOGIN,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.validateMobileCode = function(e) {
            this._paramsPack(e);
            if (e.mix_mode)
                e = this.__encryptParams(e, ["mobile", "code"]);
            return ae.request({
                url: se.MOBILE.VALIDATE_CODE,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.changeMobile = function(e) {
            this._paramsPack(e);
            e = this.__encryptParams(e, ["mobile", "code"]);
            return ae.request({
                url: se.MOBILE.CHANGE,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.accountSwitch = function(e) {
            this._paramsPack(e);
            return ae.request({
                url: se.ACCOUNT.ACCOUNT_SWITCH,
                method: "POST",
                data: e
            })
        }
        ;
        le.prototype.switchBind = function(e) {
            this._paramsPack(e);
            var t = e.type;
            var r = "app" === t ? se.OAUTH.SWITCH_BIND : se.OAUTH.SWITCH_BIND_WEB;
            delete e.type;
            return ae.request({
                url: r,
                method: "POST",
                data: e
            })
        }
        ;
        var de = le;
        var pe = location
            , be = pe.host;
        var me = be.includes("boe");
        var ve = {
            D: "//sso".concat(me ? "-boe" : "", ".douyin.com"),
            H: "//sso".concat(me ? "-boe" : "", ".huoshan.com")
        };
        var he = {
            D: me ? "//creator-boe.douyin.com" : "//creator.douyin.com",
            H: me ? "//creator-boe.huoshan.com" : "//creator.huoshan.com"
        };
        var ye = function e(t) {
            var r = t.baseURL
                , n = t.aid
                , a = t.service;
            var o = ae.request;
            var c = function e(t) {
                return o.call({}, t)
            };
            return c({
                baseURL: r,
                url: se.ACCOUNT.CHECK_LOGIN,
                method: "POST",
                data: {
                    aid: n,
                    service: a
                }
            })
        };
        var ge = function e(t) {
            var r = ae.request;
            var n = function e(t) {
                return r.call({}, t)
            };
            return n({
                url: se.ACCOUNT.CONNECT_LOGIN,
                method: "POST",
                data: t,
                callback: function e(t) {
                    return {
                        data: t
                    }
                }
            })
        };
        function Oe(e, t) {
            if (null == e)
                return {};
            var r = we(e, t);
            var n, a;
            if (Object.getOwnPropertySymbols) {
                var o = Object.getOwnPropertySymbols(e);
                for (a = 0; a < o.length; a++) {
                    n = o[a];
                    if (t.indexOf(n) >= 0)
                        continue;
                    if (!Object.prototype.propertyIsEnumerable.call(e, n))
                        continue;
                    r[n] = e[n]
                }
            }
            return r
        }
        function we(e, t) {
            if (null == e)
                return {};
            var r = {};
            var n = Object.keys(e);
            var a, o;
            for (o = 0; o < n.length; o++) {
                a = n[o];
                if (t.indexOf(a) >= 0)
                    continue;
                r[a] = e[a]
            }
            return r
        }
        function _e(e, t, r, n, a, o, c) {
            try {
                var i = e[o](c);
                var u = i.value
            } catch (s) {
                r(s);
                return
            }
            if (i.done)
                t(u);
            else
                Promise.resolve(u).then(n, a)
        }
        function Ee(e) {
            return function() {
                var t = this
                    , r = arguments;
                return new Promise((function(n, a) {
                        var o = e.apply(t, r);
                        function c(e) {
                            _e(o, n, a, c, i, "next", e)
                        }
                        function i(e) {
                            _e(o, n, a, c, i, "throw", e)
                        }
                        c(void 0)
                    }
                ))
            }
        }
        function xe(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function je(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    xe(Object(r), true).forEach((function(t) {
                            ke(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    xe(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function ke(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        function Se(e, t) {
            return Te(e) || Ie(e, t) || Ae(e, t) || Pe()
        }
        function Pe() {
            throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
        }
        function Ae(e, t) {
            if (!e)
                return;
            if ("string" === typeof e)
                return Ce(e, t);
            var r = Object.prototype.toString.call(e).slice(8, -1);
            if ("Object" === r && e.constructor)
                r = e.constructor.name;
            if ("Map" === r || "Set" === r)
                return Array.from(e);
            if ("Arguments" === r || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))
                return Ce(e, t)
        }
        function Ce(e, t) {
            if (null == t || t > e.length)
                t = e.length;
            for (var r = 0, n = new Array(t); r < t; r++)
                n[r] = e[r];
            return n
        }
        function Ie(e, t) {
            if ("undefined" === typeof Symbol || !(Symbol.iterator in Object(e)))
                return;
            var r = [];
            var n = true;
            var a = false;
            var o = void 0;
            try {
                for (var c = e[Symbol.iterator](), i; !(n = (i = c.next()).done); n = true) {
                    r.push(i.value);
                    if (t && r.length === t)
                        break
                }
            } catch (u) {
                a = true;
                o = u
            } finally {
                try {
                    if (!n && null != c["return"])
                        c["return"]()
                } finally {
                    if (a)
                        throw o
                }
            }
            return r
        }
        function Te(e) {
            if (Array.isArray(e))
                return e
        }
        var Re = function e(t) {
            var r = t.routes
                , n = void 0 === r ? [] : r
                , a = t.history;
            var o = Object(v["useContext"])(j["a"])
                , c = Se(o.userInfo, 2)
                , i = c[0]
                , u = c[1]
                , s = Se(o.accountModal, 2)
                , l = s[0]
                , f = s[1];
            var d = i.org_info;
            var p = d || {}
                , b = p.status;
            var m = i || {}
                , y = m.reload
                , w = m.status_code
                , E = m.user_profile
                , x = void 0 === E ? {} : E
                , k = m.is_registered;
            var S = w >= 0;
            var P = Boolean(x.unique_id);
            var A = [_["j"].AUDITING, _["j"].EDITING].includes(b);
            var C = Object(v["useState"])(false)
                , I = Se(C, 2)
                , T = I[0]
                , R = I[1];
            var N = function e() {
                var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : {}
                    , r = t.jump;
                window.history.replaceState({
                    urlPath: "/"
                }, "", "/".concat(r ? "?jump=".concat(r) : ""))
            };
            var D = function e(t) {
                var r = false;
                if (t) {
                    t = decodeURIComponent(t);
                    var n = t.match(/^https?:\/\/([^\/\?]+)/i);
                    var a = n && n[1] || "";
                    r = a.endsWith(".douyin.com") || a.endsWith(".huoshan.com")
                }
                return r ? t : void 0
            };
            var L = function e(t) {
                var r = new RegExp(_["f"] ? "douyin" : "huoshan","g");
                var n = _["f"] ? "huoshan" : "douyin";
                return t.includes("creator") ? t.replace(r, n) : t
            };
            var M = function e() {
                var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : {};
                f(je({
                    visible: true,
                    closable: Boolean(t.next),
                    maskClosable: !t.next
                }, t))
            };
            var U = function() {
                var e = Ee(regeneratorRuntime.mark((function e() {
                        var t, r;
                        return regeneratorRuntime.wrap((function e(n) {
                                while (1)
                                    switch (n.prev = n.next) {
                                        case 0:
                                            R(false);
                                            N();
                                            n.next = 4;
                                            return z["b"].get("/user/info/");
                                        case 4:
                                            t = n.sent;
                                            if (!t.is_confer_login) {
                                                n.next = 13;
                                                break
                                            }
                                            n.next = 8;
                                            return z["b"].post("/permission/confer/logout/");
                                        case 8:
                                            G["Toast"].info("\u5df2\u9000\u51fa\u4ee3\u8fd0\u8425\u72b6\u6001");
                                            R(true);
                                            if (_["f"])
                                                u((function(e) {
                                                        return je(je({}, e), {}, {
                                                            reload: true
                                                        })
                                                    }
                                                ));
                                            else
                                                window.location.href = "/creator-micro/home";
                                            n.next = 15;
                                            break;
                                        case 13:
                                            r = "".concat(location.protocol, "//").concat(location.hostname);
                                            location.href = "//sso".concat(r.includes("boe") ? "-boe" : "", ".").concat(_["f"] ? "huoshan" : "douyin", ".com/logout/?service=").concat(r, "/passport/web/logout/?next=").concat(r);
                                        case 15:
                                        case "end":
                                            return n.stop()
                                    }
                            }
                        ), e)
                    }
                )));
                return function t() {
                    return e.apply(this, arguments)
                }
            }();
            var q = function() {
                var e = Ee(regeneratorRuntime.mark((function e() {
                        var t, r, n, a, o, c, i, s, l, f;
                        return regeneratorRuntime.wrap((function e(d) {
                                while (1)
                                    switch (d.prev = d.next) {
                                        case 0:
                                            t = location,
                                                r = t.protocol,
                                                n = t.hostname;
                                            a = "".concat(r, "//").concat(n);
                                            d.next = 4;
                                            return ye({
                                                service: a,
                                                aid: _["g"][_["f"] ? "hotsoon" : "douyin"]
                                            });
                                        case 4:
                                            o = d.sent;
                                            c = o.has_login;
                                            i = o.redirect_url;
                                            if (!c) {
                                                d.next = 11;
                                                break
                                            }
                                            i && (location.href = i);
                                            d.next = 17;
                                            break;
                                        case 11:
                                            d.next = 13;
                                            return ye({
                                                baseURL: ve[_["f"] ? "D" : "H"],
                                                service: he[_["f"] ? "D" : "H"],
                                                aid: _["g"][_["f"] ? "douyin" : "hotsoon"]
                                            });
                                        case 13:
                                            s = d.sent;
                                            l = s.has_login;
                                            f = s.redirect_url;
                                            if (l)
                                                f && (location.href = f);
                                            else
                                                u({
                                                    status_code: 8
                                                });
                                        case 17:
                                        case "end":
                                            return d.stop()
                                    }
                            }
                        ), e)
                    }
                )));
                return function t() {
                    return e.apply(this, arguments)
                }
            }();
            var B = function e(t, r) {
                ge({
                    aid: 2906,
                    token: t,
                    platform: "aweme_v2",
                    service: "https://creator.douyin.com/creator-micro/content/publish?".concat(r),
                    force_login: 0,
                    is_vcd: 1,
                    login_scene: "connect_login"
                }).then((function(e) {
                        var t = e.message
                            , r = e.data
                            , n = void 0 === r ? {} : r;
                        var a = n.redirect_url;
                        location.assign(a)
                    }
                ))["catch"]((function(e) {
                        var t = e.error_code
                            , r = e.message;
                        if (/^109/.test(t))
                            G["Toast"].warning(r || "\u4f60\u7684\u8d26\u53f7\u88ab\u5c01\u7981");
                        else
                            G["Toast"].warning(r)
                    }
                ))
            };
            var H = function() {
                var e = Ee(regeneratorRuntime.mark((function e(t) {
                        var r, n;
                        return regeneratorRuntime.wrap((function e(a) {
                                while (1)
                                    switch (a.prev = a.next) {
                                        case 0:
                                            r = Q.a.get("/web/api/media/user/info/").then((function(e) {
                                                    var t = e.data;
                                                    return null === t || void 0 === t ? void 0 : t.user
                                                }
                                            ));
                                            n = z["b"].get("/user/info/");
                                            a.next = 4;
                                            return Promise.all([r, n]).then(function() {
                                                var e = Ee(regeneratorRuntime.mark((function e(r) {
                                                        var n, a, o, c, i, s, l, f, d;
                                                        return regeneratorRuntime.wrap((function e(p) {
                                                                while (1)
                                                                    switch (p.prev = p.next) {
                                                                        case 0:
                                                                            n = Se(r, 2),
                                                                                a = n[0],
                                                                                o = void 0 === a ? {} : a,
                                                                                c = n[1];
                                                                            c.user_profile = c.user_profile || c.douyin_user_verify_info || {};
                                                                            i = je(je({}, c.douyin_user_verify_info || {}), c.user_profile);
                                                                            if (!(i.unique_id || i.douyin_unique_id)) {
                                                                                p.next = 17;
                                                                                break
                                                                            }
                                                                            s = (i.avatar_url || "").split("//").pop();
                                                                            i.avatar_url = "".concat(s ? "//".concat(s) : "");
                                                                            l = je(je({}, c), i);
                                                                            z["d"].get("/creator/user_message/unread_count").then((function(e) {
                                                                                    var t = e.data;
                                                                                    u(Object.assign({}, o, l, {
                                                                                        unread_message_count: null === t || void 0 === t ? void 0 : t.count
                                                                                    }));
                                                                                    R(true)
                                                                                }
                                                                            ));
                                                                            if (!c.is_confer_login) {
                                                                                p.next = 15;
                                                                                break
                                                                            }
                                                                            f = g.a.parse(location.search, {
                                                                                ignoreQueryPrefix: true
                                                                            });
                                                                            if (!("jianying" === f.from && f.vid && f.token)) {
                                                                                p.next = 15;
                                                                                break
                                                                            }
                                                                            p.next = 13;
                                                                            return z["b"].post("/permission/confer/logout/");
                                                                        case 13:
                                                                            G["Toast"].info("\u5df2\u9000\u51fa\u4ee3\u8fd0\u8425\u72b6\u6001");
                                                                            window.location.href = "/creator-micro/content/publish?vid=".concat(f.vid, "&from=jianying&token=").concat(f.token, "&meta=").concat(f.meta);
                                                                        case 15:
                                                                            p.next = 26;
                                                                            break;
                                                                        case 17:
                                                                            u(c);
                                                                            R(true);
                                                                            if (!(15384 !== c.status_code)) {
                                                                                p.next = 26;
                                                                                break
                                                                            }
                                                                            if (t) {
                                                                                p.next = 26;
                                                                                break
                                                                            }
                                                                            d = g.a.parse(location.search, {
                                                                                ignoreQueryPrefix: true
                                                                            });
                                                                            if (!("jianying" === d.from && d.vid && d.token)) {
                                                                                p.next = 25;
                                                                                break
                                                                            }
                                                                            B(d.token, g.a.stringify({
                                                                                vid: d.vid,
                                                                                from: d.from,
                                                                                meta: d.meta
                                                                            }));
                                                                            return p.abrupt("return");
                                                                        case 25:
                                                                            q();
                                                                        case 26:
                                                                        case "end":
                                                                            return p.stop()
                                                                    }
                                                            }
                                                        ), e)
                                                    }
                                                )));
                                                return function(t) {
                                                    return e.apply(this, arguments)
                                                }
                                            }())["catch"]((function() {
                                                    u({
                                                        status_code: 8
                                                    });
                                                    R(true)
                                                }
                                            ));
                                        case 4:
                                        case "end":
                                            return a.stop()
                                    }
                            }
                        ), e)
                    }
                )));
                return function t(r) {
                    return e.apply(this, arguments)
                }
            }();
            var F = function() {
                var e = Ee(regeneratorRuntime.mark((function e(t) {
                        var r, n, a;
                        return regeneratorRuntime.wrap((function e(o) {
                                while (1)
                                    switch (o.prev = o.next) {
                                        case 0:
                                            r = t.logintype,
                                                n = t.loginapp,
                                                a = t.jump;
                                            if (!(r && n)) {
                                                o.next = 5;
                                                break
                                            }
                                            o.next = 4;
                                            return z["b"].post("/login/", {
                                                login_type: _["h"][r],
                                                login_app: _["g"][n]
                                            });
                                        case 4:
                                            N({
                                                jump: a
                                            });
                                        case 5:
                                            if (!a) {
                                                o.next = 9;
                                                break
                                            }
                                            a = D(a);
                                            if (a && k)
                                                location.assign(L(a));
                                            return o.abrupt("return");
                                        case 9:
                                        case "end":
                                            return o.stop()
                                    }
                            }
                        ), e)
                    }
                )));
                return function t(r) {
                    return e.apply(this, arguments)
                }
            }();
            var V = function() {
                var e = Ee(regeneratorRuntime.mark((function e() {
                        var t, r, n, o, c, i, u, s, l, f, d;
                        return regeneratorRuntime.wrap((function e(p) {
                                while (1)
                                    switch (p.prev = p.next) {
                                        case 0:
                                            t = g.a.parse(location.search, {
                                                ignoreQueryPrefix: true
                                            }),
                                                r = t.jump,
                                                n = t.swap,
                                                o = t.nocheck,
                                                c = t.noswitch,
                                                i = t.next,
                                                u = t.type,
                                                s = t.app,
                                                l = t.out,
                                                f = t.logintype,
                                                d = t.loginapp;
                                            if (!l) {
                                                p.next = 4;
                                                break
                                            }
                                            U();
                                            return p.abrupt("return");
                                        case 4:
                                            p.next = 6;
                                            return F({
                                                logintype: f,
                                                loginapp: d,
                                                jump: r
                                            });
                                        case 6:
                                            if (S && (i || u && s))
                                                if (i && P)
                                                    location.assign(L(D(i)));
                                                else {
                                                    a.replace("/");
                                                    M({
                                                        next: D(i),
                                                        loginType: u,
                                                        loginApp: s,
                                                        nocheck: o,
                                                        noswitch: c
                                                    })
                                                }
                                            if (n) {
                                                a.replace("/");
                                                M({
                                                    loginType: "user",
                                                    loginApp: n,
                                                    swap: n
                                                })
                                            }
                                            if (!S || y) {
                                                H(o);
                                                if (y)
                                                    a.replace("/")
                                            }
                                            if (A && "/" !== a.location.pathname) {
                                                G["Toast"].warning("\u4f60\u7684\u673a\u6784\u6b63\u5728\u5ba1\u6838\u4e2d");
                                                a.replace("/")
                                            }
                                        case 10:
                                        case "end":
                                            return p.stop()
                                    }
                            }
                        ), e)
                    }
                )));
                return function t() {
                    return e.apply(this, arguments)
                }
            }();
            Object(v["useEffect"])((function() {
                    if (T)
                        document.body.classList.add("mounted");
                    else
                        document.body.classList.remove("mounted")
                }
            ), [T]);
            Object(v["useEffect"])((function() {
                    V()
                }
            ), [S, y]);
            var J = Object(v["useMemo"])((function() {
                    return n.map((function(t, r) {
                            var n = t.auth
                                , a = t.component
                                , o = t.path
                                , c = t.exact
                                , i = t.subRoutes
                                , u = Oe(t, ["auth", "component", "path", "exact", "subRoutes"]);
                            var s = function t(r) {
                                var c = r.match
                                    , s = r.history;
                                if (T && c.path === o) {
                                    window.document.body.scrollTop = 0;
                                    window.document.documentElement.scrollTop = 0;
                                    var l = document.querySelector("#root>.creator-container>.semi-layout");
                                    l && (l.scrollTop = 0);
                                    if (n && (!S || !k)) {
                                        if (!S)
                                            M();
                                        s.push("/");
                                        return null
                                    }
                                }
                                return h.a.createElement(a, u, h.a.createElement(e, {
                                    routes: i
                                }))
                            };
                            return h.a.createElement(O["b"], {
                                key: r,
                                render: s,
                                path: o,
                                exact: c
                            })
                        }
                    ))
                }
            ), [S, k, T]);
            return h.a.createElement(O["d"], null, J)
        };
        var Ne = Object(O["i"])(Re);
        var De = r("3652df46c7e3ce11a2c8");
        var Le = r("63fd2a486193ef1e55ef");
        var Me = r("73342a90a6c5104ac5f9");
        function Ue(e, t) {
            return Ge(e) || Fe(e, t) || Be(e, t) || qe()
        }
        function qe() {
            throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
        }
        function Be(e, t) {
            if (!e)
                return;
            if ("string" === typeof e)
                return He(e, t);
            var r = Object.prototype.toString.call(e).slice(8, -1);
            if ("Object" === r && e.constructor)
                r = e.constructor.name;
            if ("Map" === r || "Set" === r)
                return Array.from(e);
            if ("Arguments" === r || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))
                return He(e, t)
        }
        function He(e, t) {
            if (null == t || t > e.length)
                t = e.length;
            for (var r = 0, n = new Array(t); r < t; r++)
                n[r] = e[r];
            return n
        }
        function Fe(e, t) {
            if ("undefined" === typeof Symbol || !(Symbol.iterator in Object(e)))
                return;
            var r = [];
            var n = true;
            var a = false;
            var o = void 0;
            try {
                for (var c = e[Symbol.iterator](), i; !(n = (i = c.next()).done); n = true) {
                    r.push(i.value);
                    if (t && r.length === t)
                        break
                }
            } catch (u) {
                a = true;
                o = u
            } finally {
                try {
                    if (!n && null != c["return"])
                        c["return"]()
                } finally {
                    if (a)
                        throw o
                }
            }
            return r
        }
        function Ge(e) {
            if (Array.isArray(e))
                return e
        }
        var ze = document.createElement("script");
        ze.src = "//verify.snssdk.com/static/pc_slide.js";
        document.body.appendChild(ze);
        var Ve = function() {
            var e = Object(v["useContext"])(j["a"]) || {}
                , t = Ue(e.sliderCallback, 2)
                , r = t[0]
                , n = t[1]
                , a = Ue(e.accountModal, 1)
                , o = a[0];
            var c = o.loginApp;
            var i = function e() {
                var t = {
                    ele: "pc_slide",
                    host: "https://".concat(_["b"] ? "boe-" : "", "verify.snssdk.com"),
                    aid: _["g"][c],
                    lang: "zh",
                    app_name: "aweme",
                    challenge_code: "1105",
                    toolbarBackColor: "#fff",
                    promptBackColor: "#F0F0F0",
                    promptFontColor: "#808080",
                    refreshFontColor: "#4A90E2",
                    refreshIconColor: "#4A90E2",
                    validatePassBackColor: "#A0CC49",
                    validateFailBackColor: "#EB2F2F",
                    successCb: function e() {
                        setTimeout((function() {
                                "function" === typeof r && r();
                                n(null)
                            }
                        ), 500)
                    },
                    errorCb: function e() {}
                };
                if (window.SliderVerification) {
                    var a = new window.SliderVerification(t);
                    a.init();
                    a.show()
                }
            };
            Object(v["useEffect"])((function() {
                    if (r)
                        i()
                }
            ), [r]);
            return h.a.createElement("div", {
                id: "pc_slide"
            })
        };
        var Je = r("e0eb4ce3159968a79748");
        var We = r("8d6a31f6e5330c22a4e6");
        var Ye = r.n(We);
        var Ke = r("8d3bf0c15b5acded9951");
        var Qe = r.n(Ke);
        var Ze = r("38eb9cc333ee91a6129c");
        var Xe = r.n(Ze);
        var $e = r("2a1bb827226863258ede");
        function et(e, t, r, n, a, o, c) {
            try {
                var i = e[o](c);
                var u = i.value
            } catch (s) {
                r(s);
                return
            }
            if (i.done)
                t(u);
            else
                Promise.resolve(u).then(n, a)
        }
        function tt(e) {
            return function() {
                var t = this
                    , r = arguments;
                return new Promise((function(n, a) {
                        var o = e.apply(t, r);
                        function c(e) {
                            et(o, n, a, c, i, "next", e)
                        }
                        function i(e) {
                            et(o, n, a, c, i, "throw", e)
                        }
                        c(void 0)
                    }
                ))
            }
        }
        function rt(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function nt(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    rt(Object(r), true).forEach((function(t) {
                            at(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    rt(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function at(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        function ot(e, t) {
            return lt(e) || st(e, t) || it(e, t) || ct()
        }
        function ct() {
            throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
        }
        function it(e, t) {
            if (!e)
                return;
            if ("string" === typeof e)
                return ut(e, t);
            var r = Object.prototype.toString.call(e).slice(8, -1);
            if ("Object" === r && e.constructor)
                r = e.constructor.name;
            if ("Map" === r || "Set" === r)
                return Array.from(e);
            if ("Arguments" === r || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))
                return ut(e, t)
        }
        function ut(e, t) {
            if (null == t || t > e.length)
                t = e.length;
            for (var r = 0, n = new Array(t); r < t; r++)
                n[r] = e[r];
            return n
        }
        function st(e, t) {
            if ("undefined" === typeof Symbol || !(Symbol.iterator in Object(e)))
                return;
            var r = [];
            var n = true;
            var a = false;
            var o = void 0;
            try {
                for (var c = e[Symbol.iterator](), i; !(n = (i = c.next()).done); n = true) {
                    r.push(i.value);
                    if (t && r.length === t)
                        break
                }
            } catch (u) {
                a = true;
                o = u
            } finally {
                try {
                    if (!n && null != c["return"])
                        c["return"]()
                } finally {
                    if (a)
                        throw o
                }
            }
            return r
        }
        function lt(e) {
            if (Array.isArray(e))
                return e
        }
        var ft = function(e) {
            var t = e.isActive
                , r = e.forceUseQRCode
                , n = e.setForceUseQRCode;
            var a = null;
            var o = 1e3;
            var c = Object(v["useContext"])(j["a"])
                , i = ot(c.accountSDK, 1)
                , u = i[0]
                , s = ot(c.accountModal, 2)
                , l = s[0]
                , f = s[1];
            var d = l.loginType
                , p = l.loginApp;
            var b = "hotsoon" === p;
            var m = l.next
                , y = void 0 === m ? b ? location.href : "".concat(location.origin, "/creator-micro/home") : m;
            var g = Object(v["useState"])(false)
                , O = ot(g, 2)
                , w = O[0]
                , E = O[1];
            var x = Object(v["useState"])("")
                , k = ot(x, 2)
                , S = k[0]
                , P = k[1];
            var A = Object(v["useState"])("")
                , C = ot(A, 2)
                , I = C[0]
                , T = C[1];
            var R = Object(v["useState"])("")
                , N = ot(R, 2)
                , D = N[0]
                , L = N[1];
            var M = function e(r) {
                var n = r.token
                    , c = r.baseURL
                    , i = r.service;
                var s = "".concat(i || location.href, "?logintype=").concat(d, "&loginapp=").concat(p, "&jump=").concat(y);
                t && u.checkQrcodeStatus({
                    next: y,
                    token: n,
                    baseURL: c,
                    service: s,
                    correct_service: s,
                    vsrt: i ? 1 : void 0
                }).then((function(t) {
                        var r = t.status
                            , c = t.redirect_url;
                        a = setTimeout((function() {
                                if ("new" === r || "scanned" === r) {
                                    e({
                                        token: n
                                    });
                                    T("")
                                } else if ("confirmed" === r) {
                                    try {
                                        window.localStorage.LOGIN_STATUS = JSON.stringify({
                                            logintype: d,
                                            loginapp: p
                                        })
                                    } catch (t) {
                                        void 0
                                    }
                                    $e["a"].count("qrcode_login_success");
                                    f((function(e) {
                                            return nt(nt({}, e), {}, {
                                                visible: false
                                            })
                                        }
                                    ));
                                    location.assign(c)
                                } else {
                                    E(true);
                                    T("\u4e8c\u7ef4\u7801\u5df2\u5931\u6548\uff0c\u70b9\u51fb\u5237\u65b0")
                                }
                            }
                        ), o)
                    }
                ))["catch"]((function(t) {
                        var r = t.error_code
                            , a = t.message;
                        if (18 === r)
                            e({
                                token: n,
                                baseURL: ve[_["f"] ? "D" : "H"],
                                service: he[_["f"] ? "D" : "H"]
                            });
                        else
                            T(a)
                    }
                ))
            };
            var U = function() {
                var e = tt(regeneratorRuntime.mark((function e() {
                        return regeneratorRuntime.wrap((function e(t) {
                                while (1)
                                    switch (t.prev = t.next) {
                                        case 0:
                                            t.next = 2;
                                            return u.getQrcode({
                                                next: y
                                            }).then((function(e) {
                                                    var t = e.qrcode
                                                        , r = e.token;
                                                    T("");
                                                    E(false);
                                                    P("data:image/png;base64,".concat(t));
                                                    a = setTimeout((function() {
                                                            return M({
                                                                token: r
                                                            })
                                                        }
                                                    ), o)
                                                }
                                            ))["catch"]((function(e) {
                                                    return T(e.message)
                                                }
                                            ));
                                        case 2:
                                        case "end":
                                            return t.stop()
                                    }
                            }
                        ), e)
                    }
                )));
                return function t() {
                    return e.apply(this, arguments)
                }
            }();
            Object(v["useEffect"])((function() {
                    U();
                    return function() {
                        return clearTimeout(a)
                    }
                }
            ), []);
            Object(v["useEffect"])((function() {
                    if (r && n) {
                        n(false);
                        L("\u4e3a\u4fdd\u8bc1\u5e10\u53f7\u5b89\u5168\uff0c\u8bf7\u4f7f\u7528\u6296\u97f3\u5ba2\u6237\u7aef\uff0c\u767b\u5f55\u5f53\u524d\u5e10\u53f7\u540e\u626b\u7801\u5b8c\u6210\u9a8c\u8bc1");
                        w && U()
                    }
                }
            ), [r]);
            return h.a.createElement("div", {
                className: "account-qrcode"
            }, h.a.createElement("div", {
                className: "qrcode-image"
            }, S && h.a.createElement(h.a.Fragment, null, h.a.createElement("img", {
                src: S
            }), h.a.createElement("img", {
                className: "logo".concat(b ? " hotsoon" : ""),
                src: b ? Xe.a : Qe.a
            })), w && h.a.createElement("div", {
                onClick: U
            }, h.a.createElement("img", {
                src: Ye.a
            }))), h.a.createElement("p", {
                className: I ? "error" : ""
            }, I || D || (b ? h.a.createElement(h.a.Fragment, null, "\u6253\u5f00", h.a.createElement("em", null, "\u6296\u97f3\u706b\u5c71\u7248App"), " \u626b\u63cf\u4e8c\u7ef4\u7801", h.a.createElement("br", null), "\u70b9\u51fb\u300c\u9996\u9875\u300d-\u300c\u4e2a\u4eba\u5934\u50cf\u300d-\u300c\u641c\u7d22\u300d-\u300c\u626b\u4e00\u626b\u300d") : h.a.createElement(h.a.Fragment, null, "\u6253\u5f00", h.a.createElement("em", null, "\u6296\u97f3App"), " \u626b\u63cf\u4e8c\u7ef4\u7801", h.a.createElement("br", null), "\u70b9\u51fb\u300c\u9996\u9875\u300d-\u300c\u641c\u7d22\u300d-\u300c\u626b\u4e00\u626b\u300d"))))
        };
        function dt(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function pt(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    dt(Object(r), true).forEach((function(t) {
                            bt(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    dt(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function bt(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        function mt(e, t) {
            return Ot(e) || gt(e, t) || ht(e, t) || vt()
        }
        function vt() {
            throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
        }
        function ht(e, t) {
            if (!e)
                return;
            if ("string" === typeof e)
                return yt(e, t);
            var r = Object.prototype.toString.call(e).slice(8, -1);
            if ("Object" === r && e.constructor)
                r = e.constructor.name;
            if ("Map" === r || "Set" === r)
                return Array.from(e);
            if ("Arguments" === r || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))
                return yt(e, t)
        }
        function yt(e, t) {
            if (null == t || t > e.length)
                t = e.length;
            for (var r = 0, n = new Array(t); r < t; r++)
                n[r] = e[r];
            return n
        }
        function gt(e, t) {
            if ("undefined" === typeof Symbol || !(Symbol.iterator in Object(e)))
                return;
            var r = [];
            var n = true;
            var a = false;
            var o = void 0;
            try {
                for (var c = e[Symbol.iterator](), i; !(n = (i = c.next()).done); n = true) {
                    r.push(i.value);
                    if (t && r.length === t)
                        break
                }
            } catch (u) {
                a = true;
                o = u
            } finally {
                try {
                    if (!n && null != c["return"])
                        c["return"]()
                } finally {
                    if (a)
                        throw o
                }
            }
            return r
        }
        function Ot(e) {
            if (Array.isArray(e))
                return e
        }
        var wt = G["Select"].Option;
        var _t = function(e) {
            var t = e.bind
                , r = e.forceUseQRCode;
            var n = Object(v["useContext"])(j["a"])
                , a = mt(n.accountSDK, 1)
                , o = a[0]
                , c = mt(n.accountModal, 2)
                , i = c[0]
                , u = c[1]
                , s = mt(n.sliderCallback, 2)
                , l = s[0]
                , f = s[1]
                , d = mt(n.userInfo, 2)
                , p = d[0]
                , b = d[1];
            var m = i.next
                , y = void 0 === m ? location.href : m
                , g = i.loginType
                , O = i.loginApp;
            var w = "hotsoon" === O;
            var E = Object(v["useState"])(null)
                , x = mt(E, 2)
                , k = x[0]
                , S = x[1];
            var P = Object(v["useState"])("+86")
                , A = mt(P, 2)
                , C = A[0]
                , I = A[1];
            var T = Object(v["useState"])("")
                , R = mt(T, 2)
                , N = R[0]
                , D = R[1];
            var L = Object(v["useState"])(false)
                , M = mt(L, 2)
                , U = M[0]
                , q = M[1];
            var B = Object(v["useState"])(60)
                , H = mt(B, 2)
                , F = H[0]
                , z = H[1];
            var V = Object(v["useState"])(false)
                , J = mt(V, 2)
                , W = J[0]
                , Y = J[1];
            var K = Object(v["useState"])(true)
                , Q = mt(K, 2)
                , Z = Q[0]
                , X = Q[1];
            var $ = Object(v["useState"])("")
                , ee = mt($, 2)
                , te = ee[0]
                , re = ee[1];
            var ne = Object(v["useState"])(false);
            var ae = mt(ne, 1)
                , oe = ae[0];
            var ce = function e(t) {
                if (/^\d{4}$/.test(t))
                    return true;
                else {
                    re("\u9a8c\u8bc1\u7801\u5e94\u4e3a4\u4f4d\u6570\u5b57");
                    return false
                }
            };
            var ie = function e(t) {
                if (!t) {
                    re("\u8bf7\u586b\u5199\u624b\u673a\u53f7");
                    return false
                }
                if ("+86" !== C || /^\d{11}$/.test(t))
                    return true;
                else {
                    re("\u624b\u673a\u53f7\u5e94\u4e3a11\u4f4d\u6570\u5b57");
                    return false
                }
            };
            var ue = function e() {
                if (!ie(N))
                    return;
                o.getMobileCode({
                    mobile: C + N,
                    type: t ? 8 : 24
                }).then((function(t) {
                        var r = t.status
                            , n = t.message
                            , a = t.error_code;
                        if ("success" === r) {
                            z(60);
                            Y(true);
                            re("")
                        } else if (_["k"].includes(a))
                            f((function() {
                                    return e
                                }
                            ));
                        else
                            re(n)
                    }
                ))["catch"]((function(t) {
                        if (_["k"].includes(t.error_code))
                            f((function() {
                                    return e
                                }
                            ));
                        else if (2040 === t.error_code)
                            r && r();
                        else
                            re(t.message)
                    }
                ))
            };
            var se = function e(t) {
                var r = t.number
                    , n = t.code;
                re("");
                D(r);
                if (r && n)
                    X(false);
                else
                    X(true)
            };
            var le = function e(t) {
                var n = t.number
                    , a = t.code
                    , c = t.baseURL
                    , i = t.service;
                o.userLogin({
                    account: U ? C + n : void 0,
                    mobile: U ? void 0 : C + n,
                    password: U ? a : void 0,
                    code: U ? void 0 : a,
                    baseURL: c,
                    service: "".concat(i || location.href, "?logintype=").concat(g, "&loginapp=").concat(O, "&jump=").concat(y),
                    vsrt: i ? 1 : void 0
                }, U).then((function(e) {
                        var t = e.status
                            , r = e.message
                            , n = e.data
                            , a = void 0 === n ? {} : n;
                        var o = a.redirect_url;
                        if ("success" === t) {
                            try {
                                window.localStorage.LOGIN_STATUS = JSON.stringify({
                                    logintype: g,
                                    loginapp: O
                                })
                            } catch (c) {
                                void 0
                            }
                            u((function(e) {
                                    return pt(pt({}, e), {}, {
                                        visible: false
                                    })
                                }
                            ));
                            $e["a"].count("phone_login_success");
                            location.assign(o)
                        } else
                            re(r)
                    }
                ))["catch"]((function(t) {
                        var o = t.error_code
                            , u = t.message;
                        if (18 === o)
                            e({
                                number: n,
                                code: a,
                                baseURL: ve[_["f"] ? "D" : "H"],
                                service: he[_["f"] ? "D" : "H"]
                            });
                        else if (_["k"].includes(o))
                            f((function() {
                                    return function() {
                                        return e({
                                            number: n,
                                            code: a,
                                            baseURL: c,
                                            service: i
                                        })
                                    }
                                }
                            ));
                        else if (1039 === o) {
                            G["Toast"].info("\u4e3a\u4fdd\u62a4\u8d26\u6237\u5b89\u5168\uff0c\u8bf7\u4f7f\u7528\u9a8c\u8bc1\u7801\u767b\u5f55");
                            q(false)
                        } else if (2040 === o) {
                            G["Toast"].info("\u4e3a\u4fdd\u62a4\u8d26\u6237\u5b89\u5168\uff0c\u8bf7\u4f7f\u7528\u626b\u7801\u767b\u5f55");
                            r && r()
                        } else
                            re(u)
                    }
                ))
            };
            var fe = function e(t) {
                var r = t.number
                    , n = t.code;
                o.bindMobile({
                    mobile: C + r,
                    type: 8,
                    code: n
                }).then((function(e) {
                        var t = e.status
                            , r = e.message;
                        if ("success" === t) {
                            G["Toast"].success("\u7ed1\u5b9a\u6210\u529f");
                            u((function(e) {
                                    return pt(pt({}, e), {}, {
                                        visible: false
                                    })
                                }
                            ));
                            b((function(e) {
                                    return pt(pt({}, e), {}, {
                                        user_profile: pt(pt({}, e.user_profile), {}, {
                                            login_with_only_tc: false
                                        })
                                    })
                                }
                            ))
                        } else
                            re(r)
                    }
                ))["catch"]((function(e) {
                        var t = e.error_code
                            , r = e.message;
                        re(r)
                    }
                ))
            };
            var de = function e(r) {
                var n = r.number
                    , a = r.code;
                if (!ie(n))
                    return;
                if (!U && !ce(a))
                    return;
                if (!oe && !t) {
                    re("\u8bf7\u52fe\u9009\u540c\u610f\u7528\u6237\u534f\u8bae\u548c\u9690\u79c1\u653f\u7b56");
                    return
                } else
                    re("");
                if (t)
                    fe({
                        number: n,
                        code: a
                    });
                else {
                    $e["a"].count("click_login_btn");
                    le({
                        number: n,
                        code: a
                    })
                }
            };
            Object(v["useEffect"])((function() {
                    k && k.setValue("code", "");
                    z(60);
                    Y(false);
                    re("")
                }
            ), [U]);
            var pe = function e() {
                return h.a.createElement(G["Select"], {
                    dropdownClassName: "phone-code-select",
                    value: C,
                    onChange: function e(t) {
                        return I(t)
                    },
                    renderSelectedItem: function e(t) {
                        var r = t.value;
                        return r
                    }
                }, Object.keys(_["e"]).map((function(e) {
                        return h.a.createElement(wt, {
                            key: "key",
                            value: _["e"][e]
                        }, "".concat(e, "\uff08").concat(_["e"][e], "\uff09"))
                    }
                )))
            };
            var be = function e() {
                return h.a.createElement("span", {
                    onClick: function e() {
                        return ue()
                    }
                }, "\u53d1\u9001\u9a8c\u8bc1\u7801")
            };
            var me = function e() {
                Object(v["useEffect"])((function() {
                        var e = setTimeout((function() {
                                if (F > 0)
                                    z((function(e) {
                                            return e - 1
                                        }
                                    ));
                                else
                                    Y(false)
                            }
                        ), 1e3);
                        return function() {
                            return clearTimeout(e)
                        }
                    }
                ), [F]);
                return h.a.createElement("span", {
                    className: "send-code-timer"
                }, "".concat(F, "s"))
            };
            return h.a.createElement(G["Form"], {
                className: "account-phone",
                onValueChange: se,
                onSubmit: de,
                getFormApi: function e(t) {
                    return S(t)
                }
            }, t && h.a.createElement("p", null, "\u4e3a\u9632\u6b62\u5e10\u53f7\u4e22\u5931\u4ee5\u53ca\u65b9\u4fbf\u627e\u56de\uff0c\u5efa\u8bae\u4f60\u7ed1\u5b9a\u624b\u673a\u53f7\u7801\u540e\u518d\u4f7f\u7528"), h.a.createElement(G["Form"].Input, {
                field: "number",
                type: "tel",
                size: "large",
                noLabel: true,
                placeholder: "\u8bf7\u8f93\u5165\u624b\u673a\u53f7",
                prefix: h.a.createElement(pe, null)
            }), h.a.createElement(G["Form"].Input, {
                field: "code",
                type: U ? "password" : "tel",
                size: "large",
                noLabel: true,
                placeholder: U ? "\u8bf7\u8f93\u5165\u5bc6\u7801" : "\u8bf7\u8f93\u51654\u4f4d\u9a8c\u8bc1\u7801",
                suffix: U ? null : W ? h.a.createElement(me, null) : h.a.createElement(be, null)
            }), h.a.createElement("p", {
                className: "err"
            }, te), h.a.createElement("div", {
                className: "toggle",
                onClick: function e() {
                    return q((function(e) {
                            return !e
                        }
                    ))
                }
            }, h.a.createElement("span", null, U ? "\u5207\u6362\u4e3a\u9a8c\u8bc1\u7801\u767b\u5f55 " : "\u5207\u6362\u4e3a\u5bc6\u7801\u767b\u5f55 "), h.a.createElement("svg", {
                width: "12",
                height: "12",
                viewBox: "0 0 12 12",
                fill: "none"
            }, h.a.createElement("g", {
                opacity: "0.4"
            }, h.a.createElement("path", {
                d: "M4 10L8 6L4 2",
                stroke: "black",
                strokeLinecap: "round",
                strokeLinejoin: "round"
            })))), h.a.createElement(G["Button"], {
                disabled: Z,
                block: true,
                type: "primary",
                size: "large",
                theme: "solid",
                htmlType: "submit"
            }, t ? "\u5b8c\u6210" : "\u767b\u5f55"), !t && h.a.createElement(Je["a"], {
                checkState: ne,
                agreementLink: w ? "//www.huoshanzhibo.com/agreement" : "",
                privacyLink: w ? "//www.huoshanzhibo.com/privacy" : ""
            }))
        };
        function Et(e, t, r, n, a, o, c) {
            try {
                var i = e[o](c);
                var u = i.value
            } catch (s) {
                r(s);
                return
            }
            if (i.done)
                t(u);
            else
                Promise.resolve(u).then(n, a)
        }
        function xt(e) {
            return function() {
                var t = this
                    , r = arguments;
                return new Promise((function(n, a) {
                        var o = e.apply(t, r);
                        function c(e) {
                            Et(o, n, a, c, i, "next", e)
                        }
                        function i(e) {
                            Et(o, n, a, c, i, "throw", e)
                        }
                        c(void 0)
                    }
                ))
            }
        }
        function jt(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function kt(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    jt(Object(r), true).forEach((function(t) {
                            St(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    jt(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function St(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        function Pt(e, t) {
            return Rt(e) || Tt(e, t) || Ct(e, t) || At()
        }
        function At() {
            throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
        }
        function Ct(e, t) {
            if (!e)
                return;
            if ("string" === typeof e)
                return It(e, t);
            var r = Object.prototype.toString.call(e).slice(8, -1);
            if ("Object" === r && e.constructor)
                r = e.constructor.name;
            if ("Map" === r || "Set" === r)
                return Array.from(e);
            if ("Arguments" === r || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))
                return It(e, t)
        }
        function It(e, t) {
            if (null == t || t > e.length)
                t = e.length;
            for (var r = 0, n = new Array(t); r < t; r++)
                n[r] = e[r];
            return n
        }
        function Tt(e, t) {
            if ("undefined" === typeof Symbol || !(Symbol.iterator in Object(e)))
                return;
            var r = [];
            var n = true;
            var a = false;
            var o = void 0;
            try {
                for (var c = e[Symbol.iterator](), i; !(n = (i = c.next()).done); n = true) {
                    r.push(i.value);
                    if (t && r.length === t)
                        break
                }
            } catch (u) {
                a = true;
                o = u
            } finally {
                try {
                    if (!n && null != c["return"])
                        c["return"]()
                } finally {
                    if (a)
                        throw o
                }
            }
            return r
        }
        function Rt(e) {
            if (Array.isArray(e))
                return e
        }
        var Nt = function() {
            var e = null;
            var t = 1e3;
            var r = Object(v["useContext"])(j["a"])
                , n = Pt(r.userInfo, 1)
                , a = n[0]
                , o = Pt(r.accountSDK, 1)
                , c = o[0]
                , i = Pt(r.accountModal, 2)
                , u = i[0]
                , s = i[1];
            var l = u.verifyCallback;
            var f = a.is_login_hotsoon;
            var d = location
                , p = d.protocol
                , b = d.pathname;
            var m = "".concat(p, "//creator").concat(_["b"] ? "-boe" : "", ".").concat(f ? "huoshan" : "douyin", ".com").concat(b);
            var y = Object(v["useState"])(false)
                , g = Pt(y, 2)
                , O = g[0]
                , w = g[1];
            var E = Object(v["useState"])("")
                , x = Pt(E, 2)
                , k = x[0]
                , S = x[1];
            var P = Object(v["useState"])("\u4e3a\u4fdd\u8bc1\u5e10\u53f7\u5b89\u5168\uff0c\u8bf7\u4f7f\u7528\u6296\u97f3\u5ba2\u6237\u7aef\uff0c\u767b\u5f55\u5f53\u524d\u5e10\u53f7\u540e\u626b\u7801\u5b8c\u6210\u9a8c\u8bc1")
                , A = Pt(P, 2)
                , C = A[0]
                , I = A[1];
            var T = Object(v["useState"])("")
                , R = Pt(T, 2)
                , N = R[0]
                , D = R[1];
            var L = _["g"][f ? "hotsoon" : "douyin"];
            var M = function r(n) {
                z["c"].get("/validate_by_qrcode/", {
                    params: {
                        aid: L,
                        next: m,
                        token: n
                    }
                }).then((function(a) {
                        var o = a.data;
                        var c = o || {}
                            , i = c.data
                            , u = c.message;
                        var f = i || {}
                            , d = f.status
                            , p = f.error_code;
                        e = setTimeout((function() {
                                if ("success" === u)
                                    if ("new" === d || "scanned" === d) {
                                        r(n);
                                        D("")
                                    } else if ("confirmed" === d) {
                                        s((function(e) {
                                                return kt(kt({}, e), {}, {
                                                    visible: false
                                                })
                                            }
                                        ));
                                        "function" === typeof l && l()
                                    } else {
                                        w(true);
                                        D("\u4e8c\u7ef4\u7801\u5df2\u5931\u6548\uff0c\u70b9\u51fb\u5237\u65b0")
                                    }
                                else if (1338 === p) {
                                    D("\u9a8c\u8bc1\u5931\u8d25\uff0c\u8bf7\u4f7f\u7528\u5f53\u524d\u767b\u5f55\u7528\u6237\u626b\u7801\u9a8c\u8bc1");
                                    w(true)
                                }
                            }
                        ), t)
                    }
                ))["catch"]((function(e) {
                        return D(e.message)
                    }
                ))
            };
            var U = function() {
                var r = xt(regeneratorRuntime.mark((function r() {
                        return regeneratorRuntime.wrap((function r(n) {
                                while (1)
                                    switch (n.prev = n.next) {
                                        case 0:
                                            n.next = 2;
                                            return z["c"].get("/get_qrcode/", {
                                                params: {
                                                    next: m,
                                                    aid: L
                                                }
                                            }).then((function(e) {
                                                    var t = e.data;
                                                    return t.data
                                                }
                                            )).then((function(r) {
                                                    var n = r.qrcode
                                                        , a = r.token;
                                                    D("");
                                                    w(false);
                                                    n && S("data:image/png;base64,".concat(n));
                                                    e = setTimeout((function() {
                                                            return M(a)
                                                        }
                                                    ), t)
                                                }
                                            ))["catch"]((function(e) {
                                                    return D(e.message)
                                                }
                                            ));
                                        case 2:
                                        case "end":
                                            return n.stop()
                                    }
                            }
                        ), r)
                    }
                )));
                return function e() {
                    return r.apply(this, arguments)
                }
            }();
            Object(v["useEffect"])((function() {
                    U();
                    return function() {
                        return clearTimeout(e)
                    }
                }
            ), []);
            return h.a.createElement("div", {
                className: "account-qrcode"
            }, h.a.createElement("div", {
                className: "qrcode-image"
            }, k && h.a.createElement("img", {
                src: k
            }), O && h.a.createElement("div", {
                onClick: U
            }, h.a.createElement("img", {
                src: Ye.a
            }))), h.a.createElement("p", {
                className: N ? "error" : ""
            }, N || C))
        };
        function Dt(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function Lt(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    Dt(Object(r), true).forEach((function(t) {
                            Mt(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    Dt(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function Mt(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        function Ut(e, t, r, n, a, o, c) {
            try {
                var i = e[o](c);
                var u = i.value
            } catch (s) {
                r(s);
                return
            }
            if (i.done)
                t(u);
            else
                Promise.resolve(u).then(n, a)
        }
        function qt(e) {
            return function() {
                var t = this
                    , r = arguments;
                return new Promise((function(n, a) {
                        var o = e.apply(t, r);
                        function c(e) {
                            Ut(o, n, a, c, i, "next", e)
                        }
                        function i(e) {
                            Ut(o, n, a, c, i, "throw", e)
                        }
                        c(void 0)
                    }
                ))
            }
        }
        function Bt(e, t) {
            return Vt(e) || zt(e, t) || Ft(e, t) || Ht()
        }
        function Ht() {
            throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
        }
        function Ft(e, t) {
            if (!e)
                return;
            if ("string" === typeof e)
                return Gt(e, t);
            var r = Object.prototype.toString.call(e).slice(8, -1);
            if ("Object" === r && e.constructor)
                r = e.constructor.name;
            if ("Map" === r || "Set" === r)
                return Array.from(e);
            if ("Arguments" === r || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))
                return Gt(e, t)
        }
        function Gt(e, t) {
            if (null == t || t > e.length)
                t = e.length;
            for (var r = 0, n = new Array(t); r < t; r++)
                n[r] = e[r];
            return n
        }
        function zt(e, t) {
            if ("undefined" === typeof Symbol || !(Symbol.iterator in Object(e)))
                return;
            var r = [];
            var n = true;
            var a = false;
            var o = void 0;
            try {
                for (var c = e[Symbol.iterator](), i; !(n = (i = c.next()).done); n = true) {
                    r.push(i.value);
                    if (t && r.length === t)
                        break
                }
            } catch (u) {
                a = true;
                o = u
            } finally {
                try {
                    if (!n && null != c["return"])
                        c["return"]()
                } finally {
                    if (a)
                        throw o
                }
            }
            return r
        }
        function Vt(e) {
            if (Array.isArray(e))
                return e
        }
        var Jt = function() {
            var e = Object(v["useContext"])(j["a"])
                , t = Bt(e.accountSDK, 2)
                , r = t[0]
                , n = t[1]
                , a = Bt(e.accountModal, 2)
                , o = a[0]
                , c = a[1];
            var i = o.next
                , u = o.swap
                , s = o.loginType
                , l = o.loginApp
                , f = o.noswitch;
            var d = function() {
                var e = qt(regeneratorRuntime.mark((function e(t) {
                        var r, a, o, s, l, f, d, p;
                        return regeneratorRuntime.wrap((function e(b) {
                                while (1)
                                    switch (b.prev = b.next) {
                                        case 0:
                                            r = t.type,
                                                a = t.app;
                                            o = _["f"] ? "hotsoon" : "douyin";
                                            s = "hotsoon" === a;
                                            l = o !== a;
                                            if (!l) {
                                                b.next = 17;
                                                break
                                            }
                                            b.next = 7;
                                            return ye({
                                                baseURL: ve[_["f"] ? "D" : "H"],
                                                service: he[_["f"] ? "D" : "H"],
                                                aid: _["g"][_["f"] ? "douyin" : "hotsoon"]
                                            });
                                        case 7:
                                            f = b.sent;
                                            d = f.has_login;
                                            p = "//creator".concat(_["b"] ? "-boe" : "", ".").concat(s ? "huoshan" : "douyin", ".com/");
                                            if (!d) {
                                                b.next = 15;
                                                break
                                            }
                                            location.href = "".concat(p, "?logintype=").concat(r, "&loginapp=").concat(a);
                                            return b.abrupt("return");
                                        case 15:
                                            location.href = "".concat(p, "?type=").concat(r, "&app=").concat(a, "&next=").concat(i || "", "&nocheck=").concat(u || "");
                                            return b.abrupt("return");
                                        case 17:
                                            n(new de({
                                                aid: _["g"][a]
                                            }));
                                            c((function(e) {
                                                    return Lt(Lt({}, e), {}, {
                                                        loginType: r,
                                                        loginApp: a,
                                                        swap: ""
                                                    })
                                                }
                                            ));
                                        case 19:
                                        case "end":
                                            return b.stop()
                                    }
                            }
                        ), e)
                    }
                )));
                return function t(r) {
                    return e.apply(this, arguments)
                }
            }();
            Object(v["useEffect"])((function() {}
            ), []);
            return h.a.createElement("div", null, h.a.createElement(G["Form"], {
                onSubmit: d,
                initValues: {
                    type: s || "user",
                    app: l || (_["f"] ? "hotsoon" : "douyin")
                },
                render: function e(t) {
                    var r = t.values
                        , n = r.type
                        , a = r.app;
                    return h.a.createElement(h.a.Fragment, null, h.a.createElement(G["Form"].RadioGroup, {
                        field: "type",
                        label: "\u9009\u62e9\u767b\u5f55\u65b9\u5f0f"
                    }, h.a.createElement(G["Radio"], {
                        value: "user"
                    }, "\u521b\u4f5c\u8005\u767b\u5f55"), h.a.createElement(G["Radio"], {
                        value: "org",
                        disabled: "hotsoon" === a || f
                    }, "\u673a\u6784\u767b\u5f55")), h.a.createElement("p", null, "\u673a\u6784\u767b\u5f55\u4ec5\u9002\u7528\u4e8e\u4ee3\u4ed6\u4eba\u8fd0\u8425\u7684\u673a\u6784\u7528\u6237"), h.a.createElement(G["Form"].RadioGroup, {
                        field: "app",
                        initValue: ["douyin", "hotsoon"].includes(f) ? f : void 0,
                        label: "\u9009\u62e9\u767b\u5f55\u4ea7\u54c1"
                    }, h.a.createElement(G["Radio"], {
                        value: "douyin",
                        disabled: "hotsoon" === f
                    }, "\u6296\u97f3"), h.a.createElement(G["Radio"], {
                        value: "hotsoon",
                        disabled: "org" === n || "douyin" === f
                    }, "\u6296\u97f3\u706b\u5c71\u7248")), h.a.createElement("div", {
                        className: "btn"
                    }, h.a.createElement(G["Button"], {
                        type: "primary",
                        theme: "solid",
                        htmlType: "submit"
                    }, "\u786e\u8ba4")))
                }
            }))
        };
        function Wt(e) {
            var t = Object(v["useRef"])();
            Object(v["useEffect"])((function() {
                    t.current = e
                }
            ));
            return t.current
        }
        function Yt() {
            Yt = Object.assign || function(e) {
                for (var t = 1; t < arguments.length; t++) {
                    var r = arguments[t];
                    for (var n in r)
                        if (Object.prototype.hasOwnProperty.call(r, n))
                            e[n] = r[n]
                }
                return e
            }
            ;
            return Yt.apply(this, arguments)
        }
        function Kt(e, t) {
            var r = Object.keys(e);
            if (Object.getOwnPropertySymbols) {
                var n = Object.getOwnPropertySymbols(e);
                if (t)
                    n = n.filter((function(t) {
                            return Object.getOwnPropertyDescriptor(e, t).enumerable
                        }
                    ));
                r.push.apply(r, n)
            }
            return r
        }
        function Qt(e) {
            for (var t = 1; t < arguments.length; t++) {
                var r = null != arguments[t] ? arguments[t] : {};
                if (t % 2)
                    Kt(Object(r), true).forEach((function(t) {
                            Zt(e, t, r[t])
                        }
                    ));
                else if (Object.getOwnPropertyDescriptors)
                    Object.defineProperties(e, Object.getOwnPropertyDescriptors(r));
                else
                    Kt(Object(r)).forEach((function(t) {
                            Object.defineProperty(e, t, Object.getOwnPropertyDescriptor(r, t))
                        }
                    ))
            }
            return e
        }
        function Zt(e, t, r) {
            if (t in e)
                Object.defineProperty(e, t, {
                    value: r,
                    enumerable: true,
                    configurable: true,
                    writable: true
                });
            else
                e[t] = r;
            return e
        }
        function Xt(e, t) {
            return nr(e) || rr(e, t) || er(e, t) || $t()
        }
        function $t() {
            throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")
        }
        function er(e, t) {
            if (!e)
                return;
            if ("string" === typeof e)
                return tr(e, t);
            var r = Object.prototype.toString.call(e).slice(8, -1);
            if ("Object" === r && e.constructor)
                r = e.constructor.name;
            if ("Map" === r || "Set" === r)
                return Array.from(e);
            if ("Arguments" === r || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(r))
                return tr(e, t)
        }
        function tr(e, t) {
            if (null == t || t > e.length)
                t = e.length;
            for (var r = 0, n = new Array(t); r < t; r++)
                n[r] = e[r];
            return n
        }
        function rr(e, t) {
            if ("undefined" === typeof Symbol || !(Symbol.iterator in Object(e)))
                return;
            var r = [];
            var n = true;
            var a = false;
            var o = void 0;
            try {
                for (var c = e[Symbol.iterator](), i; !(n = (i = c.next()).done); n = true) {
                    r.push(i.value);
                    if (t && r.length === t)
                        break
                }
            } catch (u) {
                a = true;
                o = u
            } finally {
                try {
                    if (!n && null != c["return"])
                        c["return"]()
                } finally {
                    if (a)
                        throw o
                }
            }
            return r
        }
        function nr(e) {
            if (Array.isArray(e))
                return e
        }
        var ar = function() {
            var e = Object(v["useContext"])(j["a"])
                , t = Xt(e.sliderCallback, 1)
                , r = t[0]
                , n = Xt(e.accountSDK, 2)
                , a = n[0]
                , o = n[1]
                , c = Xt(e.accountModal, 2)
                , i = c[0]
                , u = c[1];
            var s = Object(v["useState"])("qrcode")
                , l = Xt(s, 2)
                , f = l[0]
                , d = l[1];
            var p = Object(v["useState"])(false)
                , b = Xt(p, 2)
                , m = b[0]
                , y = b[1];
            var g = i || {}
                , O = g.visible
                , w = g.closable
                , E = void 0 === w ? false : w
                , x = g.maskClosable
                , k = void 0 === x ? true : x
                , S = g.bind
                , P = g.verify
                , A = g.loginType
                , C = g.loginApp
                , I = g.swap;
            var T = I || !S && !P && !(A && C);
            var R = {
                className: "account-modal".concat(S ? " bind" : P ? " verify" : T ? " type" : ""),
                title: S ? "\u7ed1\u5b9a\u624b\u673a\u53f7" : P ? "\u4e8c\u7ef4\u7801\u9a8c\u8bc1" : I ? "\u5207\u6362\u767b\u5f55\u4ea7\u54c1" : T ? "\u767b\u5f55\u5e10\u53f7" : "",
                closable: E,
                maskClosable: T ? k : true,
                centered: true,
                footer: null,
                width: T ? 471 : 400
            };
            var N = function e() {
                u((function(e) {
                        if (e.nocheck) {
                            location.reload();
                            return e
                        }
                        return Qt(Qt({}, e), {}, {
                            visible: false
                        })
                    }
                ))
            };
            Object(v["useEffect"])((function() {
                    if (!T && !a)
                        o(new de({
                            aid: _["g"][C]
                        }))
                }
            ), [a, T]);
            var D = Wt(O);
            if (!D && O && T)
                $e["a"].count("show_login_type_dialog");
            return h.a.createElement(h.a.Fragment, null, h.a.createElement(G["Modal"], Yt({}, R, {
                visible: O,
                onCancel: N
            }), T ? h.a.createElement(Jt, null) : S ? h.a.createElement(_t, {
                bind: true
            }) : P ? h.a.createElement(Nt, null) : h.a.createElement(h.a.Fragment, null, h.a.createElement(G["Tabs"], {
                type: "line",
                activeKey: f,
                onChange: function e(t) {
                    return d(t)
                }
            }, h.a.createElement(G["TabPane"], {
                tab: "\u626b\u7801\u767b\u5f55",
                itemKey: "qrcode"
            }, h.a.createElement(ft, {
                isActive: "qrcode" === f,
                forceUseQRCode: m,
                setForceUseQRCode: y
            })), h.a.createElement(G["TabPane"], {
                tab: "\u624b\u673a\u53f7\u767b\u5f55",
                itemKey: "phone"
            }, h.a.createElement(_t, {
                forceUseQRCode: function e() {
                    d("qrcode");
                    y(true)
                }
            }))))), h.a.createElement(G["Modal"], {
                visible: O && Boolean(r),
                className: "slider",
                header: null,
                footer: null,
                icon: null,
                centered: true,
                width: 400
            }, h.a.createElement(Ve, null)))
        };
        var or = r("9dd2ac9e2686794e7fee");
        var cr = r("4560f6d1bf73e95da83e");
        var ir = r("ed3d117d0e7678809fef");
        var ur = r("3b9cc50e592d6554d351");
        function sr(e) {
            "@babel/helpers - typeof";
            if ("function" === typeof Symbol && "symbol" === typeof Symbol.iterator)
                sr = function e(t) {
                    return typeof t
                }
                ;
            else
                sr = function e(t) {
                    return t && "function" === typeof Symbol && t.constructor === Symbol && t !== Symbol.prototype ? "symbol" : typeof t
                }
                ;
            return sr(e)
        }
        function lr(e, t) {
            if (!(e instanceof t))
                throw new TypeError("Cannot call a class as a function")
        }
        function fr(e, t) {
            for (var r = 0; r < t.length; r++) {
                var n = t[r];
                n.enumerable = n.enumerable || false;
                n.configurable = true;
                if ("value"in n)
                    n.writable = true;
                Object.defineProperty(e, n.key, n)
            }
        }
        function dr(e, t, r) {
            if (t)
                fr(e.prototype, t);
            if (r)
                fr(e, r);
            return e
        }
        function pr(e, t) {
            if ("function" !== typeof t && null !== t)
                throw new TypeError("Super expression must either be null or a function");
            e.prototype = Object.create(t && t.prototype, {
                constructor: {
                    value: e,
                    writable: true,
                    configurable: true
                }
            });
            if (t)
                br(e, t)
        }
        function br(e, t) {
            br = Object.setPrototypeOf || function e(t, r) {
                t.__proto__ = r;
                return t
            }
            ;
            return br(e, t)
        }
        function mr(e) {
            var t = yr();
            return function r() {
                var n = gr(e), a;
                if (t) {
                    var o = gr(this).constructor;
                    a = Reflect.construct(n, arguments, o)
                } else
                    a = n.apply(this, arguments);
                return vr(this, a)
            }
        }
        function vr(e, t) {
            if (t && ("object" === sr(t) || "function" === typeof t))
                return t;
            return hr(e)
        }
        function hr(e) {
            if (void 0 === e)
                throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
            return e
        }
        function yr() {
            if ("undefined" === typeof Reflect || !Reflect.construct)
                return false;
            if (Reflect.construct.sham)
                return false;
            if ("function" === typeof Proxy)
                return true;
            try {
                Date.prototype.toString.call(Reflect.construct(Date, [], (function() {}
                )));
                return true
            } catch (e) {
                return false
            }
        }
        function gr(e) {
            gr = Object.setPrototypeOf ? Object.getPrototypeOf : function e(t) {
                return t.__proto__ || Object.getPrototypeOf(t)
            }
            ;
            return gr(e)
        }
        var Or = function(e) {
            pr(r, e);
            var t = mr(r);
            function r(e) {
                var n;
                lr(this, r);
                n = t.call(this, e);
                n.state = {
                    hasError: ""
                };
                return n
            }
            dr(r, [{
                key: "componentDidCatch",
                value: function e(t, r) {}
            }, {
                key: "render",
                value: function e() {
                    var t = this.state.hasError;
                    var r = this.props.children;
                    if (t)
                        return h.a.createElement("div", {
                            className: "error-boundary"
                        }, h.a.createElement("pre", null, "\u7cfb\u7edf\u6062\u590d\u4e2d..."), h.a.createElement("pre", null, String(t)));
                    return r
                }
            }], [{
                key: "getDerivedStateFromError",
                value: function e(t) {
                    var r = t.stack || t;
                    if (r)
                        document.body.classList.add("mounted");
                    return {
                        hasError: r
                    }
                }
            }]);
            return r
        }(h.a.Component);
        var wr = r("efda4e481a0822267562");
        var _r = h.a.lazy((function() {
                return Promise.all([r.e(0), r.e(5)]).then(r.bind(null, "cd7096297567607bab4d"))
            }
        ));
        var Er = h.a.lazy((function() {
                return Promise.all([r.e(0), r.e(8)]).then(r.bind(null, "076a2104c67cc24dfd94"))
            }
        ));
        var xr = h.a.lazy((function() {
                return Promise.all([r.e(0), r.e(3)]).then(r.bind(null, "7bb98b840a1a2280ecb1"))
            }
        ));
        var jr = h.a.lazy((function() {
                return r.e(2).then(r.bind(null, "a7afabcd1c3965639082"))
            }
        ));
        var kr = h.a.lazy((function() {
                return r.e(2).then(r.bind(null, "4d5ee06f4d87b4a75697"))
            }
        ));
        var Sr = h.a.lazy((function() {
                return Promise.all([r.e(0), r.e(7)]).then(r.bind(null, "41d26551b76e20760527"))
            }
        ));
        var Pr = h.a.lazy((function() {
                return r.e(10).then(r.bind(null, "baeac90352277538f91b"))
            }
        ));
        var Ar = h.a.lazy((function() {
                return Promise.all([r.e(0), r.e(11)]).then(r.bind(null, "1d3e296683c169a68f19"))
            }
        ));
        var Cr = h.a.lazy((function() {
                return Promise.all([r.e(0), r.e(12)]).then(r.bind(null, "195d6b9ce84f950dbd80"))
            }
        ));
        var Ir = h.a.lazy((function() {
                return r.e(9).then(r.bind(null, "0eda5612a267beddd64c"))
            }
        ));
        var Tr = h.a.lazy((function() {
                return Promise.all([r.e(0), r.e(1)]).then(r.bind(null, "0ed6d6147bd8f6cdefae"))
            }
        ));
        var Rr = h.a.lazy((function() {
                return r.e(4).then(r.bind(null, "f6d8ae17819aa61a3383"))
            }
        ));
        var Nr = [{
            path: "/",
            exact: true,
            component: _r
        }, {
            path: "/organization",
            component: Er
        }, {
            path: "/authority",
            auth: true,
            component: xr
        }, {
            path: "/agreement",
            component: jr
        }, {
            path: "/college/:id",
            component: Rr
        }, {
            path: "/help/:key",
            component: kr
        }, {
            path: "/help",
            component: function e() {
                return h.a.createElement(O["a"], {
                    to: "/help/intro"
                })
            }
        }, {
            path: "/message",
            auth: true,
            component: Sr
        }, {
            path: "/publicity/topic/:name",
            auth: true,
            component: Ar
        }, {
            path: "/publicity/topic",
            auth: true,
            component: Pr
        }, {
            path: "/weekly",
            auth: true,
            component: Cr
        }, {
            path: "/rules/:id",
            component: Ir
        }, {
            path: "/task/landing",
            component: function e() {
                return h.a.createElement(O["a"], {
                    to: "/creator-micro/home"
                })
            }
        }, {
            component: _r
        }];
        var Dr = r("7a4d57deb770ff98ef68");
        function Lr(e, t, r, n, a, o, c) {
            try {
                var i = e[o](c);
                var u = i.value
            } catch (s) {
                r(s);
                return
            }
            if (i.done)
                t(u);
            else
                Promise.resolve(u).then(n, a)
        }
        function Mr(e) {
            return function() {
                var t = this
                    , r = arguments;
                return new Promise((function(n, a) {
                        var o = e.apply(t, r);
                        function c(e) {
                            Lr(o, n, a, c, i, "next", e)
                        }
                        function i(e) {
                            Lr(o, n, a, c, i, "throw", e)
                        }
                        c(void 0)
                    }
                ))
            }
        }
        function Ur() {
            Ur = Object.assign || function(e) {
                for (var t = 1; t < arguments.length; t++) {
                    var r = arguments[t];
                    for (var n in r)
                        if (Object.prototype.hasOwnProperty.call(r, n))
                            e[n] = r[n]
                }
                return e
            }
            ;
            return Ur.apply(this, arguments)
        }
        function qr(e, t) {
            if (null == e)
                return {};
            var r = Br(e, t);
            var n, a;
            if (Object.getOwnPropertySymbols) {
                var o = Object.getOwnPropertySymbols(e);
                for (a = 0; a < o.length; a++) {
                    n = o[a];
                    if (t.indexOf(n) >= 0)
                        continue;
                    if (!Object.prototype.propertyIsEnumerable.call(e, n))
                        continue;
                    r[n] = e[n]
                }
            }
            return r
        }
        function Br(e, t) {
            if (null == e)
                return {};
            var r = {};
            var n = Object.keys(e);
            var a, o;
            for (o = 0; o < n.length; o++) {
                a = n[o];
                if (t.indexOf(a) >= 0)
                    continue;
                r[a] = e[a]
            }
            return r
        }
        var Hr = ["/", "/task/landing"].includes(location.pathname);
        var Fr = function e(t, r) {
            var n = r.userInfo || {}
                , a = n.is_log_as_person
                , o = n.is_registered
                , c = n.org_info;
            var i = c || {}
                , u = i.status;
            var s = [_["j"].REJECTED].includes(u);
            var l = /\/creator-micro\/home$/.test(t);
            var f = /\/organization$/.test(window.location.href);
            var d = s || !a && !o;
            if (l && f && d)
                return true
        };
        var Gr = Object(O["i"])((function(e) {
                var t;
                var r = e.history
                    , n = qr(e, ["history"]);
                var a = r.location;
                var o = function e(t) {
                    var a = t.path;
                    var o = a instanceof Array ? a[0] : a;
                    if (Fr(o, n))
                        return;
                    if (/^(https?:?)\/\//.test(o)) {
                        window.location.href = o;
                        return
                    }
                    r.push(o)
                };
                var c = n.userInfo || {}
                    , i = c.is_confer_login
                    , u = c.is_login_hotsoon;
                var s = ["/", "/creator-micro/home"].includes(a.pathname) || i || u || (null === (t = a.hostname) || void 0 === t ? void 0 : t.includes("huoshan"));
                return h.a.createElement(or["a"], Ur({
                    onChange: o,
                    location: a,
                    showFooter: Hr,
                    hideSidebar: s
                }, n), h.a.createElement(Or, null, h.a.createElement(Ne, {
                    routes: Nr
                }), h.a.createElement(ar, null)))
            }
        ));
        var zr = t["a"] = function() {
            var e = {};
            var t = {};
            try {
                e = JSON.parse(localStorage.BEE_CONFIG || "{}");
                t = JSON.parse(localStorage.LAYOUT_CONFIG || "{}")
            } catch (m) {
                void 0
            }
            var r = Object(v["useState"])(null);
            var n = Object(v["useState"])({});
            var a = Object(v["useState"])({});
            var o = Object(v["useState"])(e);
            var c = Object(v["useState"])(t);
            var i = Object(v["useState"])(false);
            var u = Object(v["useState"])();
            var s = function e() {
                try {
                    x("".concat(_["a"]).concat(_["b"] ? 1144 : 1081, ".json"), (function(e) {
                            var t = e.data;
                            if (t) {
                                o[1](t);
                                localStorage.BEE_CONFIG = JSON.stringify(t)
                            }
                        }
                    ));
                    x("".concat(_["a"]).concat(_["b"] ? 1133 : 1017, ".json"), (function(e) {
                            var t = e.data;
                            if (t) {
                                c[1](t);
                                localStorage.LAYOUT_CONFIG = JSON.stringify(t)
                            }
                        }
                    ))
                } catch (m) {
                    void 0
                }
            };
            var l = function() {
                var e = Mr(regeneratorRuntime.mark((function e() {
                        var t;
                        return regeneratorRuntime.wrap((function e(r) {
                                while (1)
                                    switch (r.prev = r.next) {
                                        case 0:
                                            t = new wr["TtWid"]({
                                                aid: _["g"][location.host.includes("huoshan") ? "hotsoon" : "douyin"],
                                                service: location.host,
                                                region: location.host.includes("boe") ? "boe" : "cn",
                                                needFid: false
                                            });
                                            r.prev = 1;
                                            r.next = 4;
                                            return t.checkWebId();
                                        case 4:
                                            r.next = 9;
                                            break;
                                        case 6:
                                            r.prev = 6;
                                            r.t0 = r["catch"](1);
                                            void 0;
                                        case 9:
                                        case "end":
                                            return r.stop()
                                    }
                            }
                        ), e, null, [[1, 6]])
                    }
                )));
                return function t() {
                    return e.apply(this, arguments)
                }
            }();
            Object(v["useEffect"])((function() {
                    s();
                    l()
                }
            ), []);
            var f = {
                accountSDK: r,
                accountModal: n,
                userInfo: a,
                beeConfig: o,
                layoutConfig: c,
                refreshList: i,
                sliderCallback: u
            };
            var d = function e() {
                $e["a"].count("click_login");
                n[1]({
                    visible: true
                })
            };
            var p = g.a.parse(location.search, {
                ignoreQueryPrefix: true
            })
                , b = p.name;
            return h.a.createElement(w["a"], null, h.a.createElement(j["a"].Provider, {
                value: f
            }, h.a.createElement(Gr, {
                name: b,
                userInfo: a[0],
                config: c[0],
                login: d
            })))
        }
    }
});
//# sourceMappingURL=index.c2ae5a8d.js.map
