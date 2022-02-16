# 少林便當

## 執行方式

1. 安裝 Go (https://go.dev/doc/install)

2. 安裝 tdm-gcc (x64) (https://jmeubank.github.io/tdm-gcc/download/)

3.  抓下專案
    ```shell
    git clone https://github.com/wolftotem4/shaolin-ben-don.git
    ```
    
4.  切換目錄
    ```shell
    cd shaolin-ben-don
    ```
    
4.  復製 `.env.example` 至 `.env`
    
4.  編輯 `.env`
    
    ```
    ACCOUNT=(帳號)
    PASSWORD=(密碼)
    APP_DEBUG=false
    
    TELEGRAM_TOKEN=(申請的 Telegram token)
    ```
    
7. 執行

   ```shell
   go run ./cmd/bot
   ```

   