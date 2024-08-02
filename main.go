package handler

import (
    "MengCDN/model"
    "MengCDN/router"
    "net/http"
)

// Handler 是 Vercel 函数处理器
func Handler(w http.ResponseWriter, r *http.Request) {
    // 初始化数据库
    model.InitDB()

    // 初始化路由
    r := router.InitRouter()

    // 使用路由处理请求
    r.ServeHTTP(w, r)
}