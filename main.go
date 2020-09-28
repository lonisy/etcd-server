package main

import (
	"encoding/base64"
	"etcd-server/app"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/etc", getConfig)
	server := &http.Server{
		Addr:         ":8092",
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	log.Fatal(server.ListenAndServe())
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	app_key := r.URL.Query().Get("app_key")
	if app_key != "" {
		secret_key, _ := app.RedisClient.HGet(fmt.Sprintf("etc:app:%s", app_key), "secret_key").Result()
		if secret_key != "" {
			config_text, _ := app.RedisClient.Get(fmt.Sprintf("etc:%s", app_key)).Result()
			if config_text != "" {
				text := []byte(config_text)
				key := []byte(secret_key) // 必须是16位
				x1, err := app.Aes.EncryptAES(text, key)
				if err != nil {
					log.Fatalln(err)
				}
				w.Write([]byte(base64.StdEncoding.EncodeToString(x1)))
			}else{
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Connection", "close")
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Connection", "close")
		}
	}
}
