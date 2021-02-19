slack-haiku-reactor
===

![利用イメージ](./docs/image.png "利用イメージ")

投稿されたメッセージ内に存在する俳句に反応し、 Emoji でリアクションする Slack Bot です。

## Requirement

- Go 1.14+
- Slack

## Setup

Slack のソケットモードを利用しています。

環境変数にて、 Slack の Token が 2 種類、必要です。

- `SLACK_BOT_TOKEN`: `xoxb-` から始まる Bot トークン
- `SLACK_APP_TOKEN`: `xapp-` から始まる App トークン

`.env` ファイルが利用できます。 [.env.sample](/.env.sample) をコピーしてご利用ください。

デフォルトでは Emoji `:575:` をリアクションします。デフォルトでは存在しない絵文字なのでカスタム絵文字で追加してください。
環境変数 `REACT_EMOJI_FOR_HAIKU` で任意の Emoji にすることも可能です。

## How to install in your workspace

(2021/2/20 時点での情報です)

1. [ここ](https://api.slack.com/apps?new_app=1) からSlack App をワークスペースに作成
2. 左メニュー "Socket Mode" へ、"Enable Socket Mode" をオンに
   "Token Name" は適当につけて Scope はそのままで Generate
   発行される Token をコピー (これが `SLACK_APP_TOKEN`)
3. 左メニュー "Event Subscriptions" へ、 "Enable Events" をオンに
   Subscribe to bot events 内で "Add Bot User Event" を押し、 "message.channels" を追加
4. 左メニュー "OAuth & Permissions" へ
   Scopes 内で "Add an OAuth Scope" を押し、 "reactions:write" を追加
5. 左メニュー "App Home" へ、 App の見た目を整える
   主に "App Display Name" など
6. 左メニュー "Install App" へ、 ワークスペースにインストール
7. 左メニュー "Install App" へ、 もう一度
   "Bot User OAuth Access Token" をコピー (これが `SLACK_BOT_TOKEN`)
