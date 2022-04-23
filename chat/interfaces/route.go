package interfaces

type Router interface {

	//处理业务之前
	PreHandle(request Request)

	//处理中
	Handle(request Request)
	//处理之后
	PostHandle(request Request)
}
