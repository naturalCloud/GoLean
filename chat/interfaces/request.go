package interfaces

//
type Request interface {
	// 获取链接
	GetConn() Connection
	// 获取数据
	GetData() []byte
	GetMsgType() uint32
	GetRequestId() uint32
}
