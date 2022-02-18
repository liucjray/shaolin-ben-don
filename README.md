# 少林便當

一個提醒 dinbendon.net 服務有可訂購項目的 Telegram 機器人。

## 編譯

1. 安裝 Go (https://go.dev/doc/install)

2. 需要安裝 GCC

   - Windows

     - 安裝 tdm-gcc (x64) (https://jmeubank.github.io/tdm-gcc/download/)

   - Ubuntu

     - 安裝 build-essential
       ```shell
       $ sudo apt install build-essential
       ```

3. 抓下專案
   ```shell
   $ git clone https://github.com/wolftotem4/shaolin-ben-don.git
   ```

4. 切換目錄
   ```shell
   $ cd shaolin-ben-don
   ```

5. 復製 `.env.example` 至 `.env`

6. 編輯 `.env`

   ```
   ACCOUNT=(帳號)
   PASSWORD=(密碼)
   APP_DEBUG=false
   
   TELEGRAM_TOKEN=(申請的 Telegram token)
   ```

7. 編譯

   ```shell
   $ go build ./cmd/bot
   ```

8. 執行
   ```shell
   $ ./bot
   ```