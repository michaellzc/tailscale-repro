# tailscale-shutdown-repro

This reproduces the ungracefully shutdown of tailscale server by [tailscale.com/tsnet](https://pkg.go.dev/tailscale.com) package when memory store is used.

## Usage

Create an ephemeral auth key and configure the `TS_AUTHKEY` env var.

### Good 

This uses the default file store, the device is gone right after the program exits.

```sh
go run ./cmd/good
```

### Bad

This uses the memory store, the device is still there after the program exits for a while. There are logs from tsnet package saying that the device is failed to remove.

```
TryLogout control response: mustRegen=false, newURL=, err=no nodekey to log out
```

```sh
go run ./cmd/bad
```

### Bad with workaround

This uses the memory store, but we attempt to logout manually by using the `LocalClient`. The device is gone right after the program exits.

```sh
go run ./cmd/bad-workaround 
```
