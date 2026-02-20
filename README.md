# Distributed Cache

A distributed in-memory key-value cache built from scratch in Go.

Goでゼロから構築した分散インメモリKVキャッシュ。

---

## Features / 機能

- In-memory key-value store with TTL-based expiration and automatic eviction
- Consistent hashing with virtual nodes for balanced key distribution
- TCP-based client-server communication with a custom wire protocol
- Cluster-aware request routing (local handling or transparent proxying)
- Node discovery with heartbeat-based health checking
- Graceful shutdown via OS signal handling

---

- TTLベースの有効期限と自動削除を備えたインメモリKVストア
- 仮想ノードによるコンシステントハッシュで均等なキー分散
- 独自ワイヤプロトコルによるTCPベースのクライアント・サーバー通信
- クラスタ対応のリクエストルーティング（ローカル処理または透過的プロキシ）
- ハートビートによるノードディスカバリとヘルスチェック
- OSシグナルによるグレースフルシャットダウン

## Project Structure / プロジェクト構成

```
distributed-cache/
├── main.go              # CLI entry point / エントリーポイント
├── cache/               # In-memory store / インメモリストア
├── consistent/          # Consistent hashing ring / コンシステントハッシュリング
├── discovery/           # Node registry & health checks / ノード登録とヘルスチェック
├── protocol/            # Wire protocol / ワイヤプロトコル
├── server/              # TCP server & routing / TCPサーバーとルーティング
└── client/              # Client SDK / クライアントSDK
```

## Usage / 使い方

### Start a cluster / クラスタの起動

```bash
# Terminal 1 — start the first node / 最初のノードを起動
go run main.go -addr :7000

# Terminal 2 — join a second node / 2番目のノードを参加
go run main.go -addr :7001 -join :7000

# Terminal 3 — join a third node / 3番目のノードを参加
go run main.go -addr :7002 -join :7000
```

### Build and run / ビルドと実行

```bash
go build -o distributed-cache
./distributed-cache -addr :7000
```

### Run tests / テストの実行

```bash
go test ./cache/ -race -v
go test ./consistent/ -v
```

## Architecture / アーキテクチャ

1. **Cache Layer / キャッシュ層** — Each node has a local in-memory store protected by `sync.RWMutex`. Items support TTL, and a background goroutine periodically evicts expired entries.

   各ノードは`sync.RWMutex`で保護されたローカルインメモリストアを持つ。TTL付きアイテムをサポートし、バックグラウンドgoroutineが定期的に期限切れエントリを削除する。

2. **Consistent Hashing / コンシステントハッシュ** — Keys are mapped to nodes using a hash ring with virtual nodes. This ensures that adding/removing a node only remaps ~1/N of the keys.

   仮想ノード付きハッシュリングでキーをノードにマッピング。ノードの追加・削除時に再マッピングされるキーは約1/Nのみ。

3. **Protocol / プロトコル** — Requests and responses are serialized with `encoding/gob` and sent over raw TCP connections.

   リクエストとレスポンスを`encoding/gob`でシリアライズし、TCP接続で送受信。

4. **Server / サーバー** — Accepts TCP connections, decodes requests, and routes them. If the current node owns the key, it handles locally; otherwise, it proxies to the correct node.

   TCP接続を受け付け、リクエストをデコードしてルーティング。自ノードが担当するキーはローカルで処理し、それ以外は適切なノードにプロキシ。

5. **Client SDK / クライアントSDK** — Application code uses the client to connect to any node in the cluster. The node handles routing transparently.

   アプリケーションはクライアントでクラスタ内の任意のノードに接続。ルーティングはノードが透過的に処理。

6. **Discovery / ディスカバリ** — Nodes register themselves and send heartbeats. A background goroutine marks unresponsive nodes as suspect or dead.

   各ノードが自身を登録しハートビートを送信。バックグラウンドgoroutineが無応答ノードをsuspectまたはdeadとしてマーク。
