package proxy

type hoxyModule struct {
	name     string
	initFunc ModuleInitFunc
}

var (
	modules []*hoxyModule
)

// ShutdownCb will be called when the proxy is shutting down or when a user reconnects.
type ShutdownCb func(shuttingDown bool)

// ModuleInitFunc will be executed when a user authenticates with the server to get
// initialized packethooks and the shutdown callback for a module.
type ModuleInitFunc func(userCtx *UserCtx) ([]*PacketHook, ShutdownCb, error)

// RegisterMod adds a HoxyModule that will be have its hook and shutdown generators run
// when a user authenticates with the game servers.
func RegisterMod(name string, initFunc ModuleInitFunc) {
	modules = append(modules, &hoxyModule{
		name:     name,
		initFunc: initFunc,
	})
}
