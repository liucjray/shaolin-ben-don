# 少林便當

一個提醒 dinbendon.net 服務有可訂購項目的 Telegram 機器人。

## 編譯

1. 安裝 Go (https://go.dev/doc/install)

2. 抓下專案
   ```shell
   $ git clone https://github.com/wolftotem4/shaolin-ben-don.git
   ```

3. 切換目錄
   ```shell
   $ cd shaolin-ben-don
   ```

4. 復製 `.env.example` 至 `.env`

5. 編輯 `.env`

   ```
   ACCOUNT=(帳號)
   PASSWORD=(密碼)
   APP_DEBUG=false
   
   TELEGRAM_TOKEN=(申請的 Telegram token)
   ```

6. 編譯

   ```shell
   $ go build ./cmd/bot
   ```

7. 執行
   ```shell
   $ ./bot
   ```