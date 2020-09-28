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

	//http.HandleFunc("/api/etc", func(w http.ResponseWriter, r *http.Request) {
	//	query := r.URL.Query()
	//	if len(query["app_key"]) >= 1 {
	//		app_key := query["app_key"][0]
	//		fmt.Println(app_key)
	//		rdb := app.RedisClient
	//		val, err := rdb.Get(fmt.Sprintf("etc:%s", app_key)).Result()
	//		if err != nil {
	//			//panic(err)
	//			fmt.Println(err)
	//		}
	//		if val != "" {
	//			//d := []byte(val)
	//			//key := []byte("hgfedcba87654322") // 必须是16位
	//		}
	//		fmt.Println("key", val)
	//	} else {
	//		w.WriteHeader(404)
	//	}
	//	//
	//	//if app_key == "" {
	//	//	w.WriteHeader(404)
	//	//}
	//	//fmt.Println(app_key)
	//	// abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789
	//
	//	//fmt.Printf("%x\n", "A")
	//	//fmt.Printf("%x\n", "B")
	//	//fmt.Printf("%x\n", "C")
	//	//
	//	//
	//	////fmt.Printf("%s\n", 	string([]byte("0034")))
	//	//
	//	//fmt.Printf("%x\n", "a")
	//	//fmt.Printf("%x\n", "b")
	//	//fmt.Printf("%x\n", "x")
	//	//fmt.Printf("%x\n", "y")
	//	//fmt.Printf("%x\n", "z")
	//	////fmt.Printf("%c", temp)		//fmt.Printf("%x\n", "c")
	//	//var s byte = '?'
	//	//fmt.Println(s) //63
	//	////输出 2/8/10 进制格式
	//	//fmt.Printf("%b,%o,%d\n", s, s, s) // 111111,77,63
	//	//// 以16进制输出字符串
	//	//fmt.Printf("%x\n", "hex this")
	//	//// 输出数值所表示的 Unicode 字符
	//	//fmt.Printf("%c\n", 63)
	//	////输出数值所表示的 Unicode 字符（带单引号）。对于无法显示的字符，将输出其转义字符。
	//	//fmt.Printf("%q\n", 63)
	//	////输出 Unicode 码点（例如 U+1234，等同于字符串 "U+%04X" 的显示结果）
	//	//fmt.Printf("%U\n", 63)
	//	//sourceContent := "{\"abcdloaf1231sdf2ABCDEFG你好\":\"2823asf,231\"}"
	//	//a := []rune(sourceContent)
	//	//for k, v := range []rune(sourceContent) {
	//	//	fmt.Printf("%T", v)
	//	//	//fmt.Println("k2=", k, "v2=", v)
	//	//	a[k] = v + 1
	//	//}
	//	//for k, v := range a {
	//	//	fmt.Println("k2=", k, "v2=", v)
	//	//	a[k] = v+5
	//	//}
	//	//for k, v := range a {
	//	//	fmt.Println("k2=", k, "v2=", v)
	//	//	a[k] = v-10
	//	//}
	//	//fmt.Println(sourceContent)
	//	//fmt.Println(string(a))
	//	//fmt.Sprintf("%T", a)
	//	//return
	//	//for i := 'a'; i <= 'z'; i++ {
	//	//	fmt.Printf("%c", i)
	//	//}
	//	//for i := 'a'; i <= 'z'; i++ {
	//	//	fmt.Printf("%s", strings.ToUpper(string(i)))
	//	//}
	//	//for i := 0; i <= 9; i++ {
	//	//	fmt.Printf("%d", i)
	//	//}
	//
	//	d := []byte("{\"name\":123123123123}")
	//	key := []byte("hgfedcba87654322") // 必须是16位
	//
	//	fmt.Println("加密前:", string(d))
	//	x1, err := encryptAES(d, key)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	fmt.Println("加密后:", string(x1))
	//	//w.Write([]byte(string(x1)))
	//	w.Write([]byte(base64.StdEncoding.EncodeToString(x1)))
	//	x2, err := decryptAES(x1, key)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	fmt.Println("解密后:", string(x2))
	//})
	//http.ListenAndServe("127.0.0.1:8080", nil)
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	app_key := r.URL.Query().Get("app_key")
	if app_key != "" {
		secret_key, _ := app.RedisClient.HGet(fmt.Sprintf("etc:app:%s", app_key), "secret_key").Result()
		if secret_key != "" {
			fmt.Println(secret_key)
			config_text, _ := app.RedisClient.Get(fmt.Sprintf("etc:%s", app_key)).Result()
			if config_text != "" {
				//w.Write([]byte(config_text))

				text := []byte(config_text)
				key := []byte(secret_key) // 必须是16位

				//fmt.Println("加密前:", string(text))
				x1, err := app.Aes.EncryptAES(text, key)
				if err != nil {
					log.Fatalln(err)
				}
				//fmt.Println("加密后:", string(x1))
				//w.Write([]byte(string(x1)))
				w.Write([]byte(base64.StdEncoding.EncodeToString(x1)))
			}
		}
	}
}
