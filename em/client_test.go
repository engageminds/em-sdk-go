package em

import (
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	emCli, err := NewClient(&Config{
		Appk:      "storyemapp001",
		ServerUrl: "https://a.im5v.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	req := &EventRequest{
		Ts:     time.Now().UnixMilli(),
		Gaid:   "54287c36-dbd4-4e10-8b49-30541c517113",
		Make:   "Huawei",
		Brand:  "Huawei",
		Model:  "mate60 pro",
		Os:     OsAndroid,
		Osv:    "12",
		Bundle: "com.StoryEmApp001.game",
		Ip:     "1.1.1.1",
	}
	req.AddEvent(&Event{
		Ts:   time.Now().UnixMilli(),
		Cdid: "storyEventTest304",
		Eid:  "storyEventTest304",
		Props: DataMap{
			"EPstoryEventTest304": "aaaa",
			"UPstoryEventTest304": 1234,
		},
	})
	res, err := emCli.Track(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", res)
}
