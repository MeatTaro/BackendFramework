學習目標:  
    1. 建立程序入口main.go
        - 設置go module
        - net/http 設置Server處理程序，定義Server參數
        - 監聽 Addr 並呼叫Server處理requests 
        - 定義資料庫資訊，並開啟資料庫連結

    2. 資料庫連接 (PostgresSQL)
        - 安裝PostGresSQL
        - 操作 pgAdmin

    3. 練習MVC架構
        controller 控制器 — 事件層
            Controller 是整個運作過程中的核心，掌握與瀏覽器之間的互動行為。
            - net/http 收發Request 與 Response
            - 設置事件func與HTTP Method，進而觸發邏輯層來調度資料庫的資料
            - html/template 傳遞資料給View來產生樣板，回傳至客戶端瀏覽器

        model 模型 — 邏輯層
            建立資料模型，「新增、修改、刪除、瀏覽」程式邏輯來管理資料庫。
            - 建立 Struct 匹配資料庫結構
            - 使用 database/sql介面 建立操作資料庫的程式邏輯

        view 視圖 — 表現層
            View 主要管理第一線與使用者互動的介面，也就是 HTML 樣板，若是動態網站，他則會依照 Model 取出的資料內容，動態呈現使用者需要的網頁內容。


結論:
    似乎是因為id type的問題導致葉面無法正常顯示(postgressql無字曾id)，因此打算引入uuid

參考資料: 
    - https://medium.com/pierceshih/%E7%AD%86%E8%A8%98-%E4%BD%95%E8%AC%82-mvc-%E8%BB%9F%E9%AB%94%E8%A8%AD%E8%A8%88%E6%A8%A1%E5%BC%8F-af1ff10901e6

    
    



    
