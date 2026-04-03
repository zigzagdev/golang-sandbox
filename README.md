# go-sandbox

PHPやRubyなどの動的型付け言語との比較を通じて、Goの書き方の違いを自分なりにまとめる場所。

---

## ディレクトリ構成
```
go-sandbox/
├── README.md
├── 01_variables/
│   └── variables.go       # 変数宣言・型・:= vs =
├── 02_functions/
│   └── functions.go       # 関数・複数戻り値・named return
├── 03_types/
│   └── types.go           # struct・interface・型アサーション
├── 04_errors/
│   └── errors.go          # エラーハンドリング・errのイディオム
└── 05_concurrency/
    └── goroutine.go       # goroutine・channel・select
```

---
## 参考リポジトリ

今日読んだOSSコードで実例を確認する。

| リポジトリ | 用途 |
|---|---|
| [golang-migrate](https://github.com/golang-migrate/migrate) | DBマイグレーション。実際のGoコードの書き方の参考 |