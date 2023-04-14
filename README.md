# backend-proto-playground

[https://buf.build/docs/tutorials/getting-started-with-buf-cli/#configure-a-bufgenyaml](buf.build)

## 作り方

```bash
# stubの用意
$ cd proto
$ buf mod update #　何かモジュールなど追加した場合
$ buf build

# stubの作成
$ cd ./
$ buf generate proto
```
