# Hoxy

Hoxy is a proxy server designed to intercept, decrypt, and process Girl's Frontline game data.
Modified data can not be written back to the request or response at this point and only the global server is supported.
Hoxy also blocks requests to telemetry or ad domains that are frequently contacted if connecting from an android emulator. The block list is currently hardcoded in `/proxy/filters.go` but may be configurable through a config file in the future.


## Modules

Modules for hoxy should be placed in a subfolder of the mods folder.
They must be initialized by an empty import in `mods/mods.go`.
Each module should contain an `init()` function that registers the module, providing an init function that will initialize the module state, and a slice of hooks.
Here's an example from the [constructioninfo](https://github.com/kyoukaya/hoxy/blob/master/mods/constructioninfo/constructinfo.go) mod:

```go
func init() {
	const modName = "Construction Info"
	initFunc := func(userCtx *proxy.UserCtx) ([]*proxy.PacketHook, proxy.ShutdownCb, error) {
		mod := &constructInfo{}
		hooks := []*proxy.PacketHook{
			proxy.NewPacketHook(modName, "SGun/developGun", 0, false, mod.handleSingle),
			proxy.NewPacketHook(modName, "SGun/developMultiGun", 0, false, mod.handleMulti),
		}
		return hooks, func(bool) {}, nil
	}
	proxy.RegisterMod(modName, initFunc)
}
```

## Packet Definitions

If you plan on using hoxy, please note that there will be many breaking changes to the packet definitions as JSON handling will be changed.

Definitions that start with the character "S" indicate that the packet originates from the server,
likewise, a "C" prefix indicates that the packet originates from the client.

Definitions are discovered, not reversed from the game client, so definitions may be incomplete or missing. Please feel to contribute additional definitions or fixes to current definitions.
