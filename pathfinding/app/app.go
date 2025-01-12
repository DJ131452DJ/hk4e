package app

import (
	"context"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hk4e/common/config"
	"hk4e/common/mq"
	"hk4e/common/rpc"
	"hk4e/node/api"
	"hk4e/pathfinding/handle"
	"hk4e/pkg/logger"
)

var APPID string

func Run(ctx context.Context, configFile string) error {
	config.InitConfig(configFile)

	// natsrpc client
	client, err := rpc.NewClient()
	if err != nil {
		return err
	}

	// 注册到节点服务器
	rsp, err := client.Discovery.RegisterServer(context.TODO(), &api.RegisterServerReq{
		ServerType: api.PATHFINDING,
	})
	if err != nil {
		return err
	}
	APPID = rsp.GetAppId()
	go func() {
		ticker := time.NewTicker(time.Second * 15)
		for {
			<-ticker.C
			_, err := client.Discovery.KeepaliveServer(context.TODO(), &api.KeepaliveServerReq{
				ServerType: api.PATHFINDING,
				AppId:      APPID,
			})
			if err != nil {
				logger.Error("keepalive error: %v", err)
			}
		}
	}()
	defer func() {
		_, _ = client.Discovery.CancelServer(context.TODO(), &api.CancelServerReq{
			ServerType: api.PATHFINDING,
			AppId:      APPID,
		})
	}()

	logger.InitLogger("pathfinding_" + APPID)
	logger.Warn("pathfinding start, appid: %v", APPID)
	defer func() {
		logger.CloseLogger()
	}()

	messageQueue := mq.NewMessageQueue(api.PATHFINDING, APPID, client)
	defer messageQueue.Close()

	_ = handle.NewHandle(messageQueue)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-ctx.Done():
			return nil
		case s := <-c:
			logger.Warn("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				logger.Warn("pathfinding exit, appid: %v", APPID)
				return nil
			case syscall.SIGHUP:
			default:
				return nil
			}
		}
	}
}
