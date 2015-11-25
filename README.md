# 概要

GoのWeb Application FrameworkであるRevelを使って開発した簡単な検索アプリケーションです。


**環境**

* Go 1.4.2
* Revel 0.12.0


## 事前準備

Revelのインストール

```
> go get -u -v github.com/revel/revel
> go get -u -v github.com/revel/cmd/revel
```


## アプリケーションの実行

アプリケーションのルートディレクトリで下記のrevel runコマンドを実行します。

```
> revel.exe run rhw
```

アプリケーションが起動したら下記のURLにアクセスするとページを見ることができます。
アプリケーションの停止はコンソールからCtrl + Cで行います。

http://localhost:9000


## アプリケーションのビルド

revel buildコマンドはアプリケーションの実行可能なバイナリファイルを生成します。(Windowsだとexeファイル)
ビルドをする前にビルドで使用するテンポラリーフォルダが必要です。この例ではC:/tmpに作成しました。

```
> revel.exe build rhw "c:/tmp/rhw"
```

ビルドが成功するとテンポラリーフォルダに下記のファイルが生成されます。

```
> dir /b
run.bat
run.sh
src
rhw.exe
```

run.batファイルを実行するとwebアプリケーションが起動します。

```
> run.bat
```
