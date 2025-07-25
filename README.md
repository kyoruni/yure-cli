# kyoruni/yure-cli

表記ゆれを検出・修正するGo製CLIツール

## 機能
- `check`
  - 表記ゆれを検出して、行番号つきで表示
- `replace`
  - 表記ゆれを正しい表記に置換して、ファイルを上書き

## インストール

```sh
go install github.com/kyoruni/yure-cli@latest
```

## 使い方
### 1. 表記ゆれチェック

```sh
yure-cli check -i 対象ファイル.txt

# 例
yure-cli check -i hogehoge.txt
```

オプション:

- `--dict（または -d）`
  - 辞書ファイルの指定（省略時はデフォルトの辞書を使用）
    - デフォルトの辞書は `embeddata/config/dict.json` にあります

#### 出力例

```
test.txt:1: Nginx
test.txt:2: ピカチュー
test.txt:3: Nginxピカチュー
test.txt:4: hogehogeNginxhogehoge
test.txt:5: ピカチュウピカチューピカチュウ
```

### 2. 表記ゆれ置換
対象ファイル内の表記ゆれを修正して、 **上書き保存** します。

```sh
yure-cli replace -i 対象ファイル.txt

# 例
yure-cli replace -i hogehoge.txt
```

## 辞書ファイルの形式

```json
[
  { "correct": "nginx", "wrong": "Nginx" },
  { "correct": "ピカチュウ", "wrong": "ピカチュー" }
]
```

## ビルド方法

```sh
go build
```

## テスト実行方法

```sh
go test ./...

# 特定のパッケージだけテストしたい場合
go test ./cmd
```

## ライセンス
MIT