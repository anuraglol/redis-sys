# rediss

tiny redis-like in-mem store. go. listens on 127.0.0.1:8000.

## what it is

key/value store over tcp. talks resp. ascii client like `redis-cli` works.

## run

```sh
go build ./...
go run .
```

server prints `hello world` then starts.

## test

```sh
go test ./...
```

runs resp decode tests in `core/` and a concurrent smoke test in `main_test.go` (needs `redis-cli` in `PATH`).

## layout

```
main.go            entry. calls server.RunAsyncTCPServer.
server/            tcp listeners.
  async-tcp.go     epoll-based, nonblock, ~20k clients. runs expiry cron.
  sync-tcp.go      blocking, one goroutine per conn.
core/              resp + cmds + store.
  resp.go          resp2 decode/encode.
  cmd.go           RedisCmd, RedisCmds.
  comm.go          FDComm: Read/Write over raw fd.
  eval.go          all command impls.
  store.go         in-mem map + expiry map. Put/Get/Del.
  object.go        Obj, type/encoding flags.
  typeencoding.go  type/encoding bit helpers.
  type-string.go   string type/encoding picker.
  expire.go        lazy + active expiry, DeleteExpiredKeys.
  eviction.go      lru eviction.
  eviction-pool.go sorted pool, max 16.
  aof.go           aof rewrite dump.
  stats.go         atomic counters.
  resp_test.go     decode tests.
config/            knobs.
  config.go        AOF path, eviction ratio, key limit.
client/            empty for now.
```

## commands

| cmd | args | notes |
|---|---|---|
| PING | `[msg]` | `+PONG` if no arg, else bulk echo. |
| SET | `key value [EX seconds]` | `+OK`. picks int / embstr / raw encoding. |
| GET | `key` | bulk, or `$-1` nil / expired. |
| TTL | `key` | secs. `:-2` no key, `:-1` no expiry. |
| DEL | `key [key ...]` | int count deleted. |
| EXPIRE | `key seconds` | `:1` / `:0`. |
| INCR | `key` | int, creates `0` if missing. type=string, enc=int. |
| BGREWRITEAOF | - | dumps all keys to aof file as SET. |
| LRU | - | forces lru eviction pass. |
| STATS | - | array of counters. |
| CLIENT | - | stub, `+OK`. |
| LATENCY | - | stub, empty array. |

unknown cmd -> falls through to PING.

## resp

encode + decode for:
- `+` simple string
- `-` error
- `:` int64
- `$` bulk string
- `*` array (nested ok)

shortcuts: `RESP_NIL`, `RESP_OK`, `RESP_ZERO`, `RESP_ONE`, `RESP_MINUS_1`, `RESP_MINUS_2`.

## expiry

lazy on `Get`. active via cron in `async-tcp.go` every 1s -> `core.DeleteExpiredKeys`. samples 20 keys, loops while expired ratio >= 0.25.

## eviction

`Put` checks `len(store) >= config.KeysLimit` -> calls `evict()` -> `evictAllkeysLRU()`. pool max 16. evicts `EvictionRatio * KeysLimit` keys. 24-bit clock for `LastAccessedAt` (ms mod 2^24, wraps). set in `getCurrentClock`.

## aof

file: `./neko-main.aof` (config). `BGREWRITEAOF` walks `store`, writes resp `SET key value` per key. not loaded back on boot (yet).

## stats

atomic counters: total, set, get, ping, expire. read via `STATS`.

## deps

none. stdlib only. go 1.25.9.

## known bits

- `dumpKey` in `aof.go` uses `%S` (not a go verb). rewrite will panic.
- `evalSET` reads `args[3]` for EX seconds but loop index is `i`; off-by-one after `i++`. broken EX path.
- `RunSyncTCPServer` defined but unused. `main` calls async only.
- `client/` dir is empty.
- no persistence load on start.
- no tests for cmd eval, store, eviction, aof.
