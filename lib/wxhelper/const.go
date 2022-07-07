package wxhelper

import (
	. "active/common"
)

const (
	CACHE_WX_ACCESS_TOKEN = "wx_access_token_"
	CACHE_WX_JS_TICKET    = "wx_js_ticket_"
)

var (
	CacheServiceInst = &CacheService{}
)
