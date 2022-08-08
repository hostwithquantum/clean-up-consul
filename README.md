# clean-up-consul

Sometimes, you end up with synced services, that you didn't want. These _scripts_ helps clean it up.

I built this to reverse a catalog sync (of k8s services) when we got rid off it. And to learn Consul's API a bit.

## Usage

All _commands_ have a `--help` switch.

### deregister (deleting) services in consul

Three flags:

- `consul`
- `service` (optional, default: `""`)
- `tag` (default: `"k8s"`)

The tag is used to filter services, unless you provide a `--service`.

Run with: `go run ./cmd/cleanup/cleanup.go --consul http://server:port ...`

### delete all of what k8s-sync added

Two flags:

- `consul`
- `node` (default: `"k8s-sync"`)

You can execute: `go run ./cmd/delete-node/delete-node.go --consul http://server:port --node ...`
