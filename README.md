# webmaster_on_golang
webmasters API sample<br><br>

## 使い方<br>
[Google API Console - Google Cloud Console](https://console.cloud.google.com/apis)でAnalytics APIあたりを有効にし、サービスアカウントキーを取得し、secret.jsonにリネームした後、後述するga2cwディレクトリに配置する。<br>
この内容に関しては下記の記事に詳しい説明があります。<br>
[【最新版】GoogleAnalytics API 認証設定について画像多めで説明します。 | 東京上野のWeb制作会社LIG](https://liginc.co.jp/356517)<br><br>
設定終了後は下記のコマンドを参考に動作するかお試しください。<br><br>
```
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
go get google.golang.org/api/webmasters/v3
cd サンプルコードのクローンを配置したいディレクトリ
git clone https://github.com/inunosinsi/webmaster_on_golang.git
cd webmaster_on_golang
GOOGLE_APPLICATION_CREDENTIALS=secret.json go run main.go
```
