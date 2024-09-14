package service

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

// 管理注册和开始服务的入口
func Start(ctx context.Context, reg registry.Registration, host, port string,
	registerHandleFunc func()) (context.Context, error) {
	registerHandleFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		//启动http服务器错误
		log.Panicln(srv.ListenAndServe())

		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}

		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s)

		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}

		//手动停止之前肯定是先取消注册 关闭服务
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
