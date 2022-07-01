package sensor

import sdk "github.com/sensorsdata/sa-sdk-go"

func init() {
	// 从神策分析配置页面中获取数据接收的 URL
	SA_SERVER_URL := "YOUR_SERVER_URL"

	// 初始化一个 Consumer，用于数据发送
	// DefaultConsumer 是同步发送数据，因此不要在任何线上的服务中使用此 Consumer
	consumer, err := sdk.InitDefaultConsumer(SA_SERVER_URL, 10000)
	if err != nil {
		panic(err)
	}
	//...
	// 使用 Consumer 来构造 SensorsAnalytics 对象
	sa := sdk.InitSensorsAnalytics(consumer, "default", false)
	defer sa.Close()

	properties := map[string]interface{}{
		"price": 12,
		"name":  "apple",
	}

	// 记录用户事件
	err = sa.Track("ABCDEFG1234567", "ViewProduct", properties, false)
	if err != nil {
		panic(err)
	}
}
