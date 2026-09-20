package sdk

import (
	"encoding/json"
	"os"
	"testing"
)

type selfExtras struct {
	ID int `json:"id"`
}

// ToJSON implements the Extras Interface
func (s selfExtras) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

func TestAndroidPushMarshalVendorConfigs(t *testing.T) {
	android := AndroidPush{
		Alert: "this is a push",
		Honor: &HonorAndroidPush{
			Importance: "NORMAL",
			Image:      "https://example.com/honor.png",
		},
		HW: &HWAndroidPush{
			ChannelId:  "hw-channel",
			Importance: "NORMAL",
			Image:      "https://example.com/hw.png",
			Category:   "IM",
		},
		MI: &MIAndroidPush{
			ChannelId:    "mi-channel",
			LargeIconUri: "https://example.com/mi.png",
			TemplateId:   "2001",
			TemplateParam: map[string]string{
				"keywords1": "alice",
			},
		},
		OPPO: &OPPOAndroidPush{
			ChannelId:   "oppo-channel",
			Category:    "IM",
			NotifyLevel: 2,
		},
		VIVO: &VIVOAndroidPush{
			Classification: "1",
			Category:       "IM",
		},
		FCM: &FCMAndroidPush{
			ChannelId:   "fcm-channel",
			CollapseKey: "chat",
			ImageURL:    "https://example.com/fcm.png",
		},
		Meizu: &MeizuAndroidPush{
			NoticeMsgType: 1,
		},
	}

	body, err := json.Marshal(android)
	if err != nil {
		t.Fatalf("marshal android push: %v", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal android push: %v", err)
	}

	assertPushField(t, payload, "honor", "importance", "NORMAL")
	assertPushField(t, payload, "honor", "image", "https://example.com/honor.png")
	assertPushField(t, payload, "hw", "channelId", "hw-channel")
	assertPushField(t, payload, "hw", "importance", "NORMAL")
	assertPushField(t, payload, "hw", "image", "https://example.com/hw.png")
	assertPushField(t, payload, "hw", "category", "IM")
	assertPushField(t, payload, "mi", "channelId", "mi-channel")
	assertPushField(t, payload, "mi", "large_icon_uri", "https://example.com/mi.png")
	assertPushField(t, payload, "mi", "templateId", "2001")
	assertPushField(t, payload, "oppo", "channelId", "oppo-channel")
	assertPushField(t, payload, "oppo", "category", "IM")
	assertPushNumberField(t, payload, "oppo", "notify_level", 2)
	assertPushField(t, payload, "vivo", "classification", "1")
	assertPushField(t, payload, "vivo", "category", "IM")
	assertPushField(t, payload, "fcm", "channelId", "fcm-channel")
	assertPushField(t, payload, "fcm", "collapse_key", "chat")
	assertPushField(t, payload, "fcm", "imageUrl", "https://example.com/fcm.png")
	assertPushNumberField(t, payload, "meizu", "noticeMsgType", 1)

	for _, key := range []string{"channelId", "importance", "image", "large_icon_uri", "classification"} {
		if _, ok := payload[key]; ok {
			t.Fatalf("%s = %#v, want vendor field nested", key, payload[key])
		}
	}
}

func TestPushNotificationMarshalHarmonyOS(t *testing.T) {
	notification := PushNotification{
		PushContent: "this is a push",
		HarmonyOS: &HarmonyOSPush{
			Alert: "harmony push",
			OHOS: &OHOSPush{
				Category: "IM",
				Image:    "https://example.com/ohos.png",
			},
		},
	}

	body, err := json.Marshal(notification)
	if err != nil {
		t.Fatalf("marshal push notification: %v", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal push notification: %v", err)
	}
	harmonyOS, ok := payload["harmonyOS"].(map[string]interface{})
	if !ok {
		t.Fatalf("harmonyOS = %#v, want object", payload["harmonyOS"])
	}
	if got := harmonyOS["alert"]; got != "harmony push" {
		t.Fatalf("harmonyOS.alert = %v, want %q", got, "harmony push")
	}
	assertPushField(t, harmonyOS, "ohos", "category", "IM")
	assertPushField(t, harmonyOS, "ohos", "image", "https://example.com/ohos.png")
}

func TestPushNotificationMarshalAndroidConfig(t *testing.T) {
	notification := PushNotification{
		PushContent: "this is a push",
		Android: &AndroidPush{
			VIVO: &VIVOAndroidPush{
				Classification: "1",
				Category:       "IM",
			},
		},
	}

	body, err := json.Marshal(notification)
	if err != nil {
		t.Fatalf("marshal push notification: %v", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal push notification: %v", err)
	}
	android, ok := payload["android"].(map[string]interface{})
	if !ok {
		t.Fatalf("android = %#v, want object", payload["android"])
	}
	assertPushField(t, android, "vivo", "classification", "1")
	assertPushField(t, android, "vivo", "category", "IM")
}

func TestPushExtMarshalXiaomiTemplateFields(t *testing.T) {
	pushExt := PushExt{
		PushConfigs: []PushConfig{
			{
				MI: &MIAndroidPush{
					ChannelId:  "mi-channel",
					TemplateId: "2001",
					TemplateParam: map[string]string{
						"keywords1": "alice",
					},
				},
			},
		},
	}

	body, err := json.Marshal(pushExt)
	if err != nil {
		t.Fatalf("marshal push ext: %v", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal push ext: %v", err)
	}
	configs, ok := payload["pushConfigs"].([]interface{})
	if !ok || len(configs) != 1 {
		t.Fatalf("pushConfigs = %#v, want one config", payload["pushConfigs"])
	}
	config, ok := configs[0].(map[string]interface{})
	if !ok {
		t.Fatalf("pushConfigs[0] = %#v, want object", configs[0])
	}
	assertPushField(t, config, "MI", "channelId", "mi-channel")
	assertPushField(t, config, "MI", "templateId", "2001")
	mi, ok := config["MI"].(map[string]interface{})
	if !ok {
		t.Fatalf("MI = %#v, want object", config["MI"])
	}
	templateParam, ok := mi["templateParam"].(map[string]interface{})
	if !ok {
		t.Fatalf("MI.templateParam = %#v, want object", mi["templateParam"])
	}
	if got := templateParam["keywords1"]; got != "alice" {
		t.Fatalf("MI.templateParam.keywords1 = %v, want %q", got, "alice")
	}
}

func TestPushExtMarshalAllVendorConfigs(t *testing.T) {
	pushExt := PushExt{
		PushConfigs: []PushConfig{{
			HONOR: &HonorAndroidPush{Importance: "NORMAL"},
			FCM:   &FCMAndroidPush{CollapseKey: "chat"},
			OHOS:  &OHOSPush{Category: "IM"},
			MEIZU: &MeizuAndroidPush{NoticeMsgType: 1},
			APNs: &APNsPushConfig{
				ThreadID:          "thread-1",
				CollapseID:        "collapse-1",
				RichMediaURI:      "https://example.com/image.png",
				InterruptionLevel: "time-sensitive",
			},
		}},
	}

	body, err := json.Marshal(pushExt)
	if err != nil {
		t.Fatalf("marshal push ext: %v", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal push ext: %v", err)
	}
	configs, ok := payload["pushConfigs"].([]interface{})
	if !ok || len(configs) != 1 {
		t.Fatalf("pushConfigs = %#v, want one config", payload["pushConfigs"])
	}
	config, ok := configs[0].(map[string]interface{})
	if !ok {
		t.Fatalf("pushConfigs[0] = %#v, want object", configs[0])
	}

	assertPushField(t, config, "HONOR", "importance", "NORMAL")
	assertPushField(t, config, "FCM", "collapse_key", "chat")
	assertPushField(t, config, "OHOS", "category", "IM")
	assertPushNumberField(t, config, "MEIZU", "noticeMsgType", 1)
	assertPushField(t, config, "APNs", "thread-id", "thread-1")
	assertPushField(t, config, "APNs", "apns-collapse-id", "collapse-1")
	assertPushField(t, config, "APNs", "richMediaUri", "https://example.com/image.png")
	assertPushField(t, config, "APNs", "interruption-level", "time-sensitive")
}

func TestWithMsgPushExtObject(t *testing.T) {
	option, err := WithMsgPushExtObject(&PushExt{
		PushConfigs: []PushConfig{
			{
				MI: &MIAndroidPush{
					TemplateId: "2001",
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("create push ext option: %v", err)
	}

	options := modifyMsgOptions([]MsgOption{option})
	if options.pushExt == "" {
		t.Fatal("pushExt is empty")
	}

	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(options.pushExt), &payload); err != nil {
		t.Fatalf("unmarshal pushExt option: %v", err)
	}
	configs, ok := payload["pushConfigs"].([]interface{})
	if !ok || len(configs) != 1 {
		t.Fatalf("pushConfigs = %#v, want one config", payload["pushConfigs"])
	}
}

func assertPushField(t *testing.T, payload map[string]interface{}, vendor, field, want string) {
	t.Helper()

	vendorConfig, ok := payload[vendor].(map[string]interface{})
	if !ok {
		t.Fatalf("%s = %#v, want object", vendor, payload[vendor])
	}
	if got := vendorConfig[field]; got != want {
		t.Fatalf("%s.%s = %v, want %q", vendor, field, got, want)
	}
}

func assertPushNumberField(t *testing.T, payload map[string]interface{}, vendor, field string, want float64) {
	t.Helper()

	vendorConfig, ok := payload[vendor].(map[string]interface{})
	if !ok {
		t.Fatalf("%s = %#v, want object", vendor, payload[vendor])
	}
	if got := vendorConfig[field]; got != want {
		t.Fatalf("%s.%s = %v, want %v", vendor, field, got, want)
	}
}

func TestRongCloud_PushCustomObj(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	res, err := rc.PushCustomObj(PushCustomData{
		Platform: []string{"ios", "android"},
		Audience: struct {
			Tag      []string `json:"tag"`
			TagOr    []string `json:"tag_or"`
			Packages string   `json:"packageName"`
			TagItems []struct {
				Tags          []string `json:"tags"`
				IsNot         bool     `json:"isNot"`
				TagsOperator  string   `json:"tagsOperator"`
				ItemsOperator string   `json:"itemsOperator"`
			} `json:"tagItems,omitempty"`
			IsToAll bool `json:"is_to_all"`
		}{
			Tag:   []string{"female", "young"},
			TagOr: []string{"Beijing", "Shanghai"},
			TagItems: []struct {
				Tags          []string `json:"tags"`
				IsNot         bool     `json:"isNot"`
				TagsOperator  string   `json:"tagsOperator"`
				ItemsOperator string   `json:"itemsOperator"`
			}{
				{
					Tags:          []string{"guangdong", "hunan"},
					IsNot:         false,
					TagsOperator:  "OR",
					ItemsOperator: "OR",
				},
				{
					Tags:          []string{"20200408"},
					IsNot:         true,
					TagsOperator:  "OR",
					ItemsOperator: "AND",
				},
				{
					Tags:          []string{"male", "female"},
					IsNot:         false,
					TagsOperator:  "OR",
					ItemsOperator: "AND",
				},
			},
			IsToAll: false,
		},
		Notification: struct {
			Title string `json:"title"`
			Alert string `json:"alert"`
			Ios   struct {
				Title            string      `json:"title,omitempty"`
				ContentAvailable int         `json:"contentAvailable"`
				Badge            int         `json:"badge,omitempty"`
				ThreadId         string      `json:"thread-id"`
				ApnsCollapseId   string      `json:"apns-collapse-id"`
				Category         string      `json:"category,omitempty"`
				RichMediaUri     string      `json:"richMediaUri,omitempty"`
				Extras           interface{} `json:"extras"`
			} `json:"ios"`
			Android struct {
				Hw struct {
					ChannelId  string `json:"channelId"`
					Importance string `json:"importance"`
					Image      string `json:"image"`
				} `json:"hw"`
				Mi struct {
					ChannelId     string            `json:"channelId"`
					LargeIconUri  string            `json:"large_icon_uri"`
					TemplateId    string            `json:"templateId,omitempty"`
					TemplateParam map[string]string `json:"templateParam,omitempty"`
				} `json:"mi"`
				Oppo struct {
					ChannelId string `json:"channelId"`
				} `json:"oppo"`
				Vivo struct {
					Classification string `json:"classification"`
				} `json:"vivo"`
				Extras struct {
					Id   string `json:"id"`
					Name string `json:"name"`
				} `json:"extras"`
			} `json:"android"`
		}{
			Title: "标题",
			Alert: "this is a push",
			Ios: struct {
				Title            string      `json:"title,omitempty"`
				ContentAvailable int         `json:"contentAvailable"`
				Badge            int         `json:"badge,omitempty"`
				ThreadId         string      `json:"thread-id"`
				ApnsCollapseId   string      `json:"apns-collapse-id"`
				Category         string      `json:"category,omitempty"`
				RichMediaUri     string      `json:"richMediaUri,omitempty"`
				Extras           interface{} `json:"extras"`
			}{
				ThreadId:       "223",
				ApnsCollapseId: "111",
				Extras: struct {
					Id   string `json:"id"`
					Name string `json:"name"`
				}{
					Id:   "1",
					Name: "2",
				},
			},
			Android: struct {
				Hw struct {
					ChannelId  string `json:"channelId"`
					Importance string `json:"importance"`
					Image      string `json:"image"`
				} `json:"hw"`
				Mi struct {
					ChannelId     string            `json:"channelId"`
					LargeIconUri  string            `json:"large_icon_uri"`
					TemplateId    string            `json:"templateId,omitempty"`
					TemplateParam map[string]string `json:"templateParam,omitempty"`
				} `json:"mi"`
				Oppo struct {
					ChannelId string `json:"channelId"`
				} `json:"oppo"`
				Vivo struct {
					Classification string `json:"classification"`
				} `json:"vivo"`
				Extras struct {
					Id   string `json:"id"`
					Name string `json:"name"`
				} `json:"extras"`
			}{
				Hw: struct {
					ChannelId  string `json:"channelId"`
					Importance string `json:"importance"`
					Image      string `json:"image"`
				}{
					ChannelId:  "NotificationKanong",
					Importance: "NORMAL",
					Image:      "https://example.com/image.png",
				},
				Mi: struct {
					ChannelId     string            `json:"channelId"`
					LargeIconUri  string            `json:"large_icon_uri"`
					TemplateId    string            `json:"templateId,omitempty"`
					TemplateParam map[string]string `json:"templateParam,omitempty"`
				}{
					ChannelId:    "rongcloud_kanong",
					LargeIconUri: "https://example.com/image.png",
				},
				Oppo: struct {
					ChannelId string `json:"channelId"`
				}{
					ChannelId: "rc_notification_id",
				},
				Vivo: struct {
					Classification string `json:"classification"`
				}{
					Classification: "0",
				},
				Extras: struct {
					Id   string `json:"id"`
					Name string `json:"name"`
				}{
					Id:   "1",
					Name: "2",
				},
			},
		},
	})
	if err != nil {
		t.Errorf("push custom err:%v", err)
		return
	}
	t.Log("push suc res is:", res)
}

func TestRongCloud_PushCustomResObj(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	str := `{
	"platform": ["ios", "android"],
	"audience": {
		"tag": [
			"Female",
			"Young"
		],
		"tag_or": [
			"Beijing",
			"Shanghai"
		],
		"tagItems": [{
				"tags": [
					"guangdong",
					"hunan"
				],
				"isNot": false,
				"tagsOperator": "OR",
				"itemsOperator": "OR"
			},
			{
				"tags": [
					"20200408"
				],
				"isNot": true,
				"tagsOperator": "OR",
				"itemsOperator": "AND"
			},
			{
				"tags": [
					"male",
					"female"
				],
				"isNot": false,
				"tagsOperator": "OR",
				"itemsOperator": "OR"
			}
		],
		"userid": ["123","456"],
		"is_to_all": false
	},
	"notification": {
		"title": "Title",
		"alert": "this is a push",
		"ios": {
			"thread-id": "223",
			"apns-collapse-id": "111",
			"extras": {
				"id": "1",
				"name": "2"
			}
		},
		"android": {
			"hw": {
				"channelId": "NotificationKanong",
				"importance": "NORMAL",
				"image": "https://example.com/image.png"
			},
			"mi": {
				"channelId": "rongcloud_kanong",
				"large_icon_uri": "https://example.com/image.png"
			},
			"oppo": {
				"channelId": "rc_notification_id"
			},
			"vivo": {
				"classification": "0"
			},
			"extras": {
				"id": "1",
				"name": "2"
			}
		}
	}
}`
	res, err := rc.PushCustomResObj([]byte(str))
	if err != nil {
		t.Errorf("push custom err:%v", err)
		return
	}
	t.Log("push suc res is:", res)
}

func TestRongCloud_PushCustom(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	str := `{
  "platform":["ios","android"],
  "audience":{
    "tag":["Female","Young"],
    "tag_or":["Beijing","Shanghai"],
	 "tagItems":[
				  {
					  "tags":[
						  "guangdong",
						  "hunan"
					  ],
					  "isNot":false,
					  "tagsOperator":"OR",
					  "itemsOperator":"OR"
				  },
				  {
					  "tags":[
						  "20200408"
					  ],
					  "isNot":true,
					  "tagsOperator":"OR",
					  "itemsOperator":"AND"
				  },
				  {
					  "tags":[
						  "male",
						  "female"
					  ],
					  "isNot":false,
					  "tagsOperator":"OR",
					  "itemsOperator":"OR"
				  }
			  ],
		  "userid":[
			  "123",
			  "456"
		 ],
	  "is_to_all":false
	},
  "notification":{
    "title":"标题",
    "alert":"this is a push",
    "ios":
      {
        "thread-id":"223",
        "apns-collapse-id":"111",
        "extras": {"id": "1","name": "2"}
      },
    "android": {
        "hw":{
            "channelId":"NotificationKanong",
            "importance": "NORMAL",
            "image":"https://example.com/image.png"
        },
        "mi":{
            "channelId":"rongcloud_kanong",
            "large_icon_uri":"https://example.com/image.png"
        },
        "oppo":{
            "channelId":"rc_notification_id"
        },
        "vivo":{
            "classification":"0"
        },
        "extras": {"id": "1","name": "2"}
      }
  }
}`
	res, err := rc.PushCustom([]byte(str))
	if err != nil {
		t.Errorf("push custom err:%v", err)
		return
	}
	t.Log("push suc res is:", res)
}

func TestRongCloud_PushSend(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	push := Push{
		PlatForm: []PlatForm{
			IOSPlatForm,
			AndroidPlatForm,
		},
		Audience: Audience{
			IsToAll: true,
		},
		Notification: Notification{
			Alert: "this is a push",
			IOS: IOSPush{
				Title: "iOS platform display title",
				Alert: "iOS platform display content",
				Extras: selfExtras{
					ID: 1,
				},
			},
			Android: AndroidPush{
				Alert: "Android platform display content",
				Extras: selfExtras{
					ID: 1,
				},
			},
		},
	}
	p, err := rc.PushSend(push)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(p.Code)
		t.Log(p.ID)
	}

	msg := TXTMsg{
		Content: "hello",
		Extra:   "helloExtra",
	}
	msgr, err := msg.ToString()
	if err != nil {
		t.Fatal(err)
	}
	broadcast := Broadcast{
		PlatForm: []PlatForm{
			IOSPlatForm,
			AndroidPlatForm,
		},
		FromUserID: "u01",
		Audience: Audience{
			IsToAll: true,
		},
		Message: Message{
			Content:              msgr,
			ObjectName:           "RC:TxtMsg",
			DisableUpdateLastMsg: true,
		},
	}
	p, err = rc.PushSend(broadcast)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(p.Code)
		t.Log(p.ID)
	}
}
