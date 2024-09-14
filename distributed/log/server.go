package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

// 对请求进行log
func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(data)
}

// 创建自定义logger 将log写入到filePath， "go"日志前缀
func Run(filePath string) {
	log = stlog.New(fileLog(filePath), "[server-go]: ", stlog.LstdFlags)
}

// 针对"/log"路径的日志请求服务做处理
func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		//处理请求
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			//写入log
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

	})
}

func write(message string) {
	log.Printf("%v \n", message)
}
