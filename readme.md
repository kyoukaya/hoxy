# Hoxy

Hoxy is a proxy server designed to intercept, decrypt, and process Girl's Frontline game data.
Modified data can not be written back to the request or response at this point and only the global server is supported.

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

Definitions that start with the character "S" indicate that the packet originates from the server,
likewise, a "C" prefix indicates that the packet originates from the client.

Definitions are discovered, not reversed from the game client, so definitions may be incomplete or missing. Please feel to contribute additional definitions or fixes to current definitions.

## Bugs

`MarshalMismatchErr` is returned by the packet unmarshaller frequently due to a few reasons.
- Go maps are not ordered. Fixed in the interim by using a generic ordered map[string]interface{}.
- JSON fields may appear only when they are being used. Possibly fixed by `,omitempty` struct tags.

## TODO

- Move userauth out of the mods folder as it provides core functionality.
- Config file.
- Fix the constant `MarshalMismatchErr`s. Might have to modify the standard json library or use a custom one to implement ordered maps. Conditional JSON fields may need specific custom marshal/unmarshal procedures. 
- Support other game servers. This is limited by the fact that I don't have clients for any server other than the global server.
- Support caching of update data. The game updates off AWS using a https connection, capability to generate a unique CA cert and MITM https connections is already written in and can be activated with the `-https` flag. Further analysis of update traffic is required.
- Order hook execution by their priority.
