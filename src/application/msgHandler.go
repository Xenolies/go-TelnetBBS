package application

import (
	"TelnetBBS/src/inface"
	"TelnetBBS/src/utils"
	"fmt"
)

/*
消息处理模块
*/

type MsgHandle struct {
	// 存放每一个MsgID对应的处理方法
	Apis map[string]inface.IRouter

	// 负责处理Worker取任务的消息队列
	TaskQueue []chan inface.IRequest

	// 业务工作 worker 数量
	WorkerPoolSize uint32
}

// NewMsgHandle 创建消息处理对象
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[string]inface.IRouter),
		WorkerPoolSize: utils.Gc.WorkerPoolSize,                               // 从全局配置中获取最大worker数量
		TaskQueue:      make([]chan inface.IRequest, utils.Gc.WorkerPoolSize), // 最大消息队列数量
	}
}

// DoMsgHandler 调度  执行对应的Router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request inface.IRequest) {
	// 从 Request中找到 MsgID
	handler, ok := mh.Apis[request.GetCommand()]
	if !ok {
		fmt.Println("Api MsgID: ", request.GetCommand(), " Not Found")
		return
	}
	// 根据MsgID执行对应Router业务
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(cmd string, router inface.IRouter) {
	// 判断当前msg绑定的Router是否存在
	if _, ok := mh.Apis[cmd]; ok {
		// ID 已经注册了
		panic("Repeat API, Command: " + cmd)
	}

	// 添加Msg和API的绑定关系
	mh.Apis[cmd] = router
	fmt.Println("Add Api MsgID SUCCESS: ", cmd)
}

// StartWorkPool 启动Worker工作池 (只发生一次)
func (mh *MsgHandle) StartWorkPool() {
	// 根据WorkerPoolSize来启动 Worker, 每个Worker使用一个go承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 启动一个worker
		// 1 当前worker对应的channel消息队列开辟空间
		mh.TaskQueue[i] = make(chan inface.IRequest, utils.Gc.MaxWorkerTaskLen)
		// 2 启动当前worker 阻塞等待Channel传输过来的消息
		go mh.statOneWorker(i, mh.TaskQueue[i])
	}

}

// 启动Worker工作流程
func (mh *MsgHandle) statOneWorker(workrID int, taskQueue chan inface.IRequest) {
	fmt.Println("Start WorkerID: ", workrID, "Start!!")

	// 阻塞等待消息
	for {
		select {
		// 如果队列中有消息, 执行当前Request的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)

		}

	}

}

// SendMsgToTaskQueue 将消息交给 TaskWorker 交给 Worker 处理
func (mh *MsgHandle) SendMsgToTaskQueue(request inface.IRequest) {
	// 将消息平均分配给不通过的 Worker
	// 根据客户端链接 ConnID 分配
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize // 按照一定规则分配给队列
	fmt.Println("ADD ConnId: ", request.GetConnection().GetConnID(), " Command: ", request.GetCommand(), "--> WorkerID: ", workerID)

	// 将消息发送给对应的 Worker 的 TaskQueue
	mh.TaskQueue[workerID] <- request
}
