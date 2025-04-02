package binance

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"pledge-backend/log"
	"time"
)

var PlgrPriceChan = make(chan string, 10)
var Con *websocket.Conn

// 定义消息结构体
type BinanceMessage struct {
	Stream string `json:"stream"`
	Data   struct {
		E  string `json:"e"`
		E1 int64  `json:"E"`
		S  string `json:"s"`
		T  int64  `json:"t"`
		P  string `json:"p"`
		Q  string `json:"q"`
		T1 int64  `json:"T"`
		M  bool   `json:"m"`
		M1 bool   `json:"M"`
	} `json:"data"`
}

func GetExchangePrice() {
	// 币安的WebSocket地址
	binanceUrl := "wss://data-stream.binance.vision/stream"
	// 建立WebSocket连接
	c, _, err := websocket.DefaultDialer.Dial(binanceUrl, nil)
	Con = c
	if err != nil {
		log.Logger.Sugar().Fatal("Failed to connect:", err)
	}
	defer Con.Close()
	go func(c *websocket.Conn) {
		for {
			time.Sleep(10 * time.Second)
			err := c.WriteControl(websocket.PongMessage, []byte{}, time.Now())
			if err != nil {
				log.Logger.Sugar().Warn("发送PONG帧失败", err)
				return
			}
			log.Logger.Sugar().Info("发送PONG帧")
		}
	}(c)
	// 循环接收消息
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Logger.Sugar().Error("Failed to read message:", err)
			return
		}

		var binanceMsg BinanceMessage
		// 解析JSON消息
		err = json.Unmarshal(message, &binanceMsg)
		if err != nil {
			log.Logger.Sugar().Error("Failed to parse message:", err)
			continue
		}
		PlgrPriceChan <- binanceMsg.Data.P
		// 打印交易对符号和价格
		log.Logger.Sugar().Info("Symbol: ", binanceMsg.Data.S, ", Price:", binanceMsg.Data.P)
		time.Sleep(1 * time.Second)
	}

}
