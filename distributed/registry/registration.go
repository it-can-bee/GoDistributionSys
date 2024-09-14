package registry

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
	/* 动态接收来自注册服务中心的更新 */
	RequiredServices []ServiceName //当前服务所需依赖关系链
	///如果你在一个分布式系统中注册一个服务，这个字段用于告诉注册中心或服务发现系统：为了正常工作，该服务还需要依赖哪些其他服务。可以想象这是一个依赖关系链，用来确保依赖的服务都正常运行
	ServiceUpdateURL string
	HeartbeatURL     string
}

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}

type ServiceName string

// 准备注册
const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("Portald")
)
