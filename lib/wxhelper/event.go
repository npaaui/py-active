package wxhelper

/////////////////////////EventType 事件类型////////////////////////////////

// EventType 事件类型
type EventType string


// EventBase 事件基础类
type EventBase struct {
	ToUserName   string    // 开发者微信号
	FromUserName string    // 发送方帐号（一个OpenID）
	CreateTime   string    // 消息创建时间（整型）
	MsgType      MsgType   // 消息类型，event
	Event        EventType // 事件类型
}
