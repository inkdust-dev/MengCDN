package handler

import (
    "MengCDN/internal/model"
    "MengCDN/internal/router"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // 初始化数据库和路由
    model.InitDB()
    r := router.InitRouter()

    // 使用路由处理请求
    r.ServeHTTP(w, r)
}
