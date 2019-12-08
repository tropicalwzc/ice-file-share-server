package main

import (
    "fmt"
    "net"
	"net/http"
    "github.com/gorilla/websocket"
    "log"
    "os"
    "strings"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024000,
    WriteBufferSize: 1024000,
}

//调用os.MkdirAll递归创建文件夹
func createFile(filePath string)  error  {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath,os.ModePerm)
		return err
	}
	return nil
}
 
// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
    return true
}

func localIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}

func main() {
    createFile("ice_share_folder")
    ips, _ :=localIPv4s()
    ipall := len(ips)
    for i:=0;i<ipall;i++{
        ipstr:=strings.Split(ips[i],".")
        if len(ipstr)!=4{ 
            continue
        }
        if ipstr[1]=="254"{
            continue
        }
        if ipstr[3]=="1"{
            continue
        }
        fmt.Println("本机ip地址")
        fmt.Println(ips[i]+":7777")
        fmt.Println("需要允许在专用和公用网络通信才能确保连接成功")
        fmt.Println("在浏览器打开"+ips[i]+":7777\n即可与其他位于同一个wifi下的电脑共享位于\nice_share_folder中的所有文件:")
        fmt.Println()
    }
    //http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("C:\\sharefolder"))))
    http.Handle("/", http.FileServer(http.Dir("ice_share_folder")))
	/*
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

        for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }

            // Print the message to the console
            fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

            // Write message back to browser
            if err = conn.WriteMessage(msgType, msg); err != nil {
                return
            }
        }
    })
	
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "/Users/tropical_fish/go/net_test/websocket.html")
	})
	http.ListenAndServe(":8080", nil)
	*/
	
	log.Fatal(http.ListenAndServe(":7777", nil))
}