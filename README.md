# EngageMinds SDK GO

1. 添加 sdk

```bash
go get -u github.com/engageminds/em-sdk-go
```

2. 使用

```go
package main

import (
	"github.com/engageminds/em-sdk-go/em"
	"time"
)

func main() {
	// emCli 全局一个实例
	emCli, err := em.NewClient(&em.Config{
		Appk: "{ReplaceWithYourAppk}",
	})

	// 发送事件
	req := &em.EventRequest{
		Ts:     time.Now().UnixMilli(),
		Gaid:   "54287c36-dbd4-4e10-8b49-30541c517113",
		Make:   "Huawei",
		Brand:  "Huawei",
		Model:  "mate60 pro",
		Os:     em.OsAndroid,
		Osv:    "12",
		Bundle: "com.StoryEmApp001.game",
		Ip:     "1.1.1.1",
	}
	req.AddEvent(&em.Event{
		Ts:   time.Now().UnixMilli(),
		Cdid: "storyEventTest304",
		Eid:  "storyEventTest304",
		Props: em.DataMap{
			"EPstoryEventTest304": "aaaa",
			"UPstoryEventTest304": 1234,
		},
	})
	res, err := emCli.track(req)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debugf("%+v", res)
}

```
