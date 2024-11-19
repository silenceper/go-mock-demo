package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// 从环境变量中读取值
	version := os.Getenv("VERSION")
	app := os.Getenv("APP")
	port := os.Getenv("PORT")
	upstreamURL := os.Getenv("UPSTREAM_URL")

	// 如果没有设置端口，使用默认值
	if port == "" {
		port = "9090"
	}
	// 设置处理函数
	http.HandleFunc("/mock", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s\n", r.Host)
		// 构建当前服务的信息
		currentServiceInfo := fmt.Sprintf("%s(version: %s, ip: %s)", app, version, r.Host)

		// 如果存在上游服务，进行调用
		if upstreamURL != "" {
			resp, err := http.Get(upstreamURL)
			if err != nil {
				log.Printf("Failed to call upstream service: %v", err)
				http.Error(w, "Failed to call upstream service", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()
			upstreamInfo, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Failed to read upstream response: %v", err)
				http.Error(w, "Failed to read upstream response", http.StatusInternalServerError)
				return
			}

			// 输出当前服务和上游服务的信息
			fmt.Fprintf(w, "%s -> %s", currentServiceInfo, upstreamInfo)
		} else {
			// 只输出当前服务的信息
			fmt.Fprint(w, currentServiceInfo)
		}
	})

	// 启动服务器
	log.Printf("Starting server on :%s , app: %s, version:%s \n", port, app, version)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
