# try-go-clean-arch

Goでクリーンアーキテクチャを試す

# Architecture
Clearn Archtecture \
大きく４つのレイヤーに分けた


- domain
- usecase
- adapter (interface)
- infrastructure

# API

APIのサンプルとして以下のエンドポイントを用意

GET
- api/v1/courses
- api/v1/course/:id

# Get Started

- Dockerビルド

```=sh
make docker
```

- APIサーバおよびDBサーバの起動

```=sh
make run
```

- 終了

```=sh
make stop
```

# Ref

- [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [bxcodex/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [pospomeのサーバサイドアーキテクチャ](https://booth.pm/ja/items/1045782)
- [API サーバーを Clean Architecture で構築する](https://tech-blog.optim.co.jp/entry/2019/01/29/173000)
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)