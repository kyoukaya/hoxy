package proxy

import (
	"hoxy/authcode"
	"hoxy/log"
	"sync"
)

// UserCtx contains all the state pertaining to an authenticated user connected with
// the game servers.
type UserCtx struct {
	mutex       *sync.Mutex
	Key         string
	Longtoken   string
	UID         string
	Openid      int
	RawHooks    map[string][]*PacketHook
	Hooks       map[string][]*PacketHook
	shutdownCBs []ShutdownCb
}

func insertHook(hookMap map[string][]*PacketHook, hook *PacketHook) error {
	var hookSlice []*PacketHook
	hookSlice, ok := hookMap[hook.target]
	if !ok {
		hookSlice = make([]*PacketHook, 0)
	}

	// TODO: Order by hook priority when inserting
	hookSlice = append(hookSlice, hook)
	hookMap[hook.target] = hookSlice

	return nil
}

func (userCtx *UserCtx) initMods(mods []*hoxyModule) {
	if userCtx == nil {
		return
	}
	log.Infof("%s connected. Loading modules...\n", userCtx.UID)
	for _, mod := range mods {
		hooks, cb, err := mod.initFunc(userCtx)
		if err != nil {
			log.Warnf("%s failed to load.", mod.name)
			continue
		}
		userCtx.initHooks(hooks)
		// Init shutdown callback
		userCtx.shutdownCBs = append(userCtx.shutdownCBs, cb)
		log.Infof("%s loaded.", mod.name)
	}
}

func (userCtx *UserCtx) initHooks(hooks []*PacketHook) {
	if userCtx == nil {
		return
	}
	for _, hook := range hooks {
		var destMap map[string][]*PacketHook
		if hook.raw {
			destMap = userCtx.RawHooks
		} else {
			destMap = userCtx.Hooks
		}

		if err := insertHook(destMap, hook); err != nil {
			log.Warnf("Error while loading %#v: %v", hook, err)
		}
	}
}

// Decode decodes the input string into bytes with the user's key. If the UserCtx ptr
// is nil, decode will use the default key.
func (userCtx *UserCtx) Decode(s string) ([]byte, int64, error) {
	if userCtx != nil {
		return authcode.Decode(s, userCtx.Key)
	}

	return authcode.Decode(s, "")
}
