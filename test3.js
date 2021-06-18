function a () {
    for (var t, e, n = [], r = 0; r < arguments.length; r++)
        n[r] = arguments[r];
    return Wn(i, this._url, s) && (t = (e = i.commonParams || {}).bid,
        e = e.web_id,
    t && e && (c = Hn(e, t),
        this.setRequestHeader(Rn, c))),
        or(this)(i, c, a, s),
        this._start = Date.now(),
        this._data = null == n ? void 0 : n[0],
        o.apply(this, n)
}