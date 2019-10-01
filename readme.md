# Hoxy

Hoxy is a proxy server designed to intercept, decrypt, and process Girl's Frontline game data.
Modified data can not be written back to the request or response at this point and only the global server is supported.
Hoxy also blocks requests to telemetry or ad domains that are frequently contacted if connecting from an android emulator. The block list is currently hardcoded in `proxy/filters.go` but may be user configurable in the future.

## Usage

An example program is provided in `cmd/example` which initializes the packetlogger and constructioninfo mods.
Use `go run cmd/example/hoxy.go` to start the proxy server up, and then direct your client to use it.

## Modules

Each module should initialize itself by calling `proxy.RegisterMod(modName, initFunc)` at program start up, either in the `init()` function or in the global scope.
The init function will be called when a user authenticates with the game server to set up the module for that user.
Here's an example from the [constructioninfo](https://github.com/kyoukaya/hoxy/blob/master/mods/constructioninfo/constructioninfo.go) mod:

```go
func init() {
	const modName = "Construction Info"
	initFunc := func(userCtx *proxy.UserCtx) ([]*proxy.PacketHook, proxy.ShutdownCb, error) {
		dollinfo.Init()
		equipinfo.Init()
		mod := &constructInfo{}
		hooks := []*proxy.PacketHook{
			proxy.NewPacketHook(modName, "SGun/developGun", 0, false, mod.handleGunConstruct),
			proxy.NewPacketHook(modName, "SGun/developMultiGun", 0, false, mod.handleGunConstruct),
			proxy.NewPacketHook(modName, "SEquip/develop", 0, false, mod.handleEquipConstruct),
			proxy.NewPacketHook(modName, "SEquip/developMulti", 0, false, mod.handleEquipConstruct),
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
Additionally, helpful comments about packet fields that are not immediately obvious are welcome.
