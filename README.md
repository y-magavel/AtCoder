# AtCoder

## Docker環境構築手順

infraディレクトリに移動しDockerイメージをビルドする
```shell
docker build -t procon-gardener .
```

プロジェクトディレクトリに戻り、Dockerコンテナを起動する

※コンテナ内の時間をちゃんと合わせないとコミット時間が提出時間とズレる
```shell
docker run -dit --name procon-gardener -v $PWD:/go/src/app/ -v /usr/share/zoneinfo/Asia/Tokyo:/etc/localtime procon-gardener
```

起動したコンテナの中に入る
```shell
docker exec -it procon-gardener /bin/sh
```

コンテナの中でprocon-gardenerの初期化を実行
```shell
procon-gardener init
```

コンテナの中でprocon-gardenerの設定ファイルを編集する

※ Alpine Linuxにはvimは入ってないのでviを使う
```shell
vi /root/.procon-gardener/config.json
```

※設定内容については[AtCoderの提出を取得してGitHubの芝を生やすコマンドラインツールを作った](https://qiita.com/togatoga/items/3e8fd0042dc8be702201)を参照する

※config.jsonのrepository_pathには`/go/src/app`を設定する