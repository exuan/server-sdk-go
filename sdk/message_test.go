/*
 * @Descripttion:
 * @version:
 * @Author: ran.ding
 * @Date: 2019-09-02 18:29:55
 * @LastEditors: ran.ding
 * @LastEditTime: 2019-09-10 11:37:27
 */
package sdk

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestRongCloud_MessageExpansionDel(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if err := rc.MessageExpansionDel("C16R-VBGG-1IE5-SD0C",
		"u01",
		"3",
		"testExp0309",
		"[\"key1\",\"key2\"]",
		1,
	); err != nil {
		t.Error(err)
		return
	}
	t.Log("do UGMessageGet suc")
}

func TestRongCloud_MessageExpansionSet(t *testing.T) {
	data, err := json.Marshal(map[string]string{"type": "3"})
	if err != nil {
		t.Log("marshal err", err)
		return
	}
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if err := rc.MessageExpansionSet("C16R-VBGG-1IE5-SD0C",
		"u01",
		"3",
		"testExp0309",
		string(data),
		1,
	); err != nil {
		t.Error(err)
		return
	}
	t.Log("do UGMessageGet suc")
}

func TestRongCloud_UGMessageGetObj(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UGMessageGetObj("target_001", []UGMessageData{
		{
			MsgUid:     "C16R-VBGG-1IE5-SD0C",
			BusChannel: "001",
		},
	}); err != nil {
		t.Error(err)
		return
	} else {
		t.Logf("do UGMessageGet suc :%+v", res)
	}
}

func TestRongCloud_UGMessageGet(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	if res, err := rc.UGMessageGet("target_001", []UGMessageData{
		{
			MsgUid:     "C16R-VBGG-1IE5-SD0C",
			BusChannel: "001",
		},
	}); err != nil {
		t.Error(err)
		return
	} else {
		t.Log("do UGMessageGet suc", string(res))
	}

}

func TestRongCloud_UGMessageModify(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	result, err := rc.UGMessagePublish("aa", "RC:TxtMsg", "{\"content\":\"1234455667788-0309-1-test\"}",
		"", "", "1", "", "0", "0", "", "{\"key1\":\"key1\"}",
		false, false, &PushExt{
			Title:                "you have a new message.",
			TemplateId:           "123456",
			ForceShowPushContent: 0,
			PushConfigs: []PushConfig{
				{
					HW: &HWAndroidPush{ChannelId: "NotificationKanong"},
				},
				{
					MI: &MIAndroidPush{ChannelId: "rongcloud_kanong"},
				},
				{
					OPPO: &OPPOAndroidPush{ChannelId: "rc_notification_id"},
				},
				{
					VIVO: &VIVOAndroidPush{Classification: "0"},
				},
				{
					APNs: &APNsPushConfig{ThreadID: "1", CollapseID: "1"},
				},
			},
		}, "testExp0309")
	if err != nil {
		t.Errorf("ug message send err:%v", err)
		return
	}
	t.Log(result)
	t.Log("ug message send suc")
	time.Sleep(1 * time.Second)
	// note : msgUID is obtained through the Post-messaging Callback, for details: https://doc.rongcloud.cn/imserver/server/v1/message/sync
	if res, err := rc.UGMessageModify("testExp0309", "aa", "C1PL-LJQR-0U1B-ADFN", "Hello", UgMessageExtension{
		BusChannel: "",
		MsgRandom:  0,
	}); err != nil {
		t.Errorf("UGMessageModify request err:%v", err)
		return
	} else {
		t.Log("UGMessageModify suc", string(res))
	}
}

func TestMessageBroadcastRecall(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	content := BroadcastRecallContent{
		MessageId:        "BC52-ESJ0-022O-001H",
		ConversationType: 6,
		IsAdmin:          false,
		IsDelete:         false,
	}

	result, err := rc.MessageBroadcastRecall("123", "RC:RcCmd", content)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		t.Log("PASS")
		t.Log(result)
	}
}

func TestChatRoomRecall(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	if err := rc.ChatRoomRecall("fDR2cVpxxR5zSMUNh3yAwh", "MersNRhaKwJkRV9mJR5JXY", "5FGT-7VA9-G4DD-4V5P", 1507778882124); err != nil {
		t.Errorf("error: %v", err)
	} else {
		t.Log("Pass")
	}
}

func TestSystemRecall(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	err := rc.SystemRecall("fDR2cVpxxR5zSMUNh3yAwh",
		"MersNRhaKwJkRV9mJR5JXY",
		"5FGT-7VA9-G4DD-4V5P",
		1507778882124,
		WithIsAdmin(1),
		WithIsDelete(1),
	)

	if err != nil {
		t.Errorf("error: %v", err)
	} else {
		t.Log("Pass")
	}
}

func TestMessage_PrivateSend(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	msg := TXTMsg{
		Content: "hello",
		Extra:   "helloExtra",
	}

	result, err := rc.PrivateSend(
		"7Szq13MKRVortoknTAk7W8",
		[]string{"4kIvGJmETlYqDoVFgWdYdM"},
		"RC:TxtMsg",
		&msg,
		"",
		"",
		1,
		0,
		1,
		0,
		0,
	)
	t.Log(err)
	t.Log(result)
}

func TestMessage_PrivateSend_SightMsg(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
		WithRongCloudURI("https://api.rong-api.com"),
	)

	msg := SightMsg{
		Content: "hello",
		Extra:   "helloExtra",
		User: MsgUserInfo{
			ID:       "4kIvGJmETlYqDoVFgWdYdM",
			Name:     "ming",
			Portrait: "http://www.rongcloud.cn/images/logo.png",
		},
		SightURL: "http://www.rongcloud.cn/xxx.mp4",
		Duration: 10,
		Name:     "hello",
	}

	resp, err := rc.PrivateSend(
		"7Szq13MKRVortoknTAk7W8",
		[]string{"4kIvGJmETlYqDoVFgWdYdM"},
		"RC:SightMsg",
		&msg,
		"",
		"",
		1,
		0,
		1,
		0,
		0,
	)
	t.Log(resp)
	t.Log(err)
}

func TestMessage_PrivateSendOptions(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	msg := TXTMsg{
		Content: "hello",
		Extra:   "helloExtra",
	}

	result, err := rc.PrivateSend(
		"7Szq13MKRVortoknTAk7W8",
		[]string{"4kIvGJmETlYqDoVFgWdYdM"},
		"RC:TxtMsg",
		&msg,
		"",
		"",
		1,
		0,
		1,
		0,
		0,
		WithMsgDisablePush(true),
		WithMsgPushExt(""),
		WithMsgBusChannel("bus"),
	)
	t.Log(err)
	t.Log(result)
}

func TestMessage_PrivateRecall(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	err := rc.PrivateRecall(
		"7Szq13MKRVortoknTAk7W8",
		"4kIvGJmETlYqDoVFgWdYdM",
		"B7CE-U880-31M6-D3EE",
		1543566558208,
	)
	t.Log(err)
}

func TestMessage_PrivateSendTemplate(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	tpl1 := TemplateMsgContent{
		TargetID: "4kIvGJmETlYqDoVFgWdYdM",
		Data: map[string]string{
			"{name}":  "Xiao Ming",
			"{score}": "90",
		},
		PushContent: "{name} your score is out",
	}

	tpl2 := TemplateMsgContent{
		TargetID: "GvYBoFJQTggripS_qoiVaA",
		Data: map[string]string{
			"{name}":  "Xiao Hong",
			"{score}": "95",
		},
		PushContent: "{name}, your score is out",
	}

	msg := TXTMsg{
		Content: "{name}, Chinese score {score} points",
		Extra:   "helloExtra",
	}

	var tpl []TemplateMsgContent
	tpl = append(tpl, tpl1)
	tpl = append(tpl, tpl2)
	result, err := rc.PrivateSendTemplate(
		"7Szq13MKRVortoknTAk7W8",
		"RC:TxtMsg",
		msg,
		tpl)
	t.Log(err)
	t.Log(result)
}

func TestRongCloud_GroupSend(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	msg := TXTMsg{
		Content: "hello",
		Extra:   "helloExtra",
	}

	result, err := rc.GroupSend(
		"7Szq13MKRVortoknTAk7W8",
		[]string{"CFtiYbXNQNYtSr7rzUfHco"},
		[]string{},
		"RC:TxtMsg",
		&msg,
		"",
		"",
		1,
		0,
	)
	t.Log(err)
	t.Log(result)
}

func TestRongCloud_PrivateRecall(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	err := rc.GroupRecall(
		"7Szq13MKRVortoknTAk7W8",
		"CFtiYbXNQNYtSr7rzUfHco",
		"B7CE-U880-31M6-D3EE",
		1543566558208,
	)

	t.Log(err)
}

func TestRongCloud_GroupSendMention(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	msg := MentionMsgContent{
		Content:       "@user_2 hello",
		MentionedInfo: MentionedInfo{Type: 2, UserIDs: []string{"4kIvGJmETlYqDoVFgWdYdM"}, PushContent: "Someone mentioned you"},
	}
	result, err := rc.GroupSendMention(
		"7Szq13MKRVortoknTAk7W8",
		[]string{"cYgiKZzRSUsrfrx6C3u_GI"},
		"RC:TxtMsg",
		msg,
		"",
		"",
		1,
		0,
		1,
		0,
	)
	t.Log(err)
	t.Log(result)
}

func TestRongCloud_ChatRoomSend(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	msg := TXTMsg{
		Content: "hello",
		Extra:   "helloExtra",
	}

	result, err := rc.ChatRoomSend(
		"7Szq13MKRVortoknTAk7W8",
		[]string{"4kIvGJmETlYqDoVFgWdYdM"},
		"RC:TxtMsg",
		&msg, 0, 0,
	)
	t.Log(err)
	t.Log(result)
}

func TestRongCloud_ChatroomBroadcast(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	msg := TXTMsg{
		Content: "hello",
		Extra:   "helloExtra",
	}

	err := rc.ChatRoomBroadcast(
		"7Szq13MKRVortoknTAk7W8",
		"RC:TxtMsg",
		&msg,
	)
	t.Log(err)
}

func TestRongCloud_OnlineBroadcast(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	result, err := rc.OnlineBroadcast(
		"someone",
		"RC:TxtMsg",
		"hello everyone",
	)
	t.Log(result)
	t.Log(err)
}

func TestRongCloud_SystemSend(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	msg := TXTMsg{
		Content: "hello",
		Extra:   "helloExtra",
	}

	result, err := rc.SystemSend(
		"7Szq13MKRVortoknTAk7W8",
		[]string{"4kIvGJmETlYqDoVFgWdYdM"},
		"RC:TxtMsg",
		&msg,
		"",
		"",
		0,
		1,
	)
	t.Log(err)
	t.Log(result)
}

func TestRongCloud_SystemBroadcast(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	msg := TXTMsg{
		Content: "hello",
		Extra:   "helloExtra",
	}

	result, err := rc.SystemBroadcast(
		"7Szq13MKRVortoknTAk7W8",
		"RC:TxtMsg",
		&msg,
	)
	t.Log(err)
	t.Log(result)
}

func TestRongCloud_SystemBroadcastOption(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	msg := TXTMsg{
		Content: "hello",
		Extra:   "helloExtra",
	}

	result, err := rc.SystemBroadcast(
		"7Szq13MKRVortoknTAk7W8",
		"RC:TxtMsg",
		&msg,
		WithMsgPushContent("thisisapush"),
	)
	t.Log(err)
	t.Log(result)
}

func TestRongCloud_SystemSendTemplate(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	tpl1 := TemplateMsgContent{
		TargetID: "4kIvGJmETlYqDoVFgWdYdM",
		Data: map[string]string{
			"{name}":  "Xiao Ming",
			"{score}": "90",
		},
		PushContent: "{name} your score is out",
	}

	tpl2 := TemplateMsgContent{
		TargetID: "GvYBoFJQTggripS_qoiVaA",
		Data: map[string]string{
			"{name}":  "Xiao Hong",
			"{score}": "95",
		},
		PushContent: "{name} your score is out",
	}

	msg := TXTMsg{
		Content: "{name}, your Chinese score is {score} points",
		Extra:   "helloExtra",
	}

	var tpl []TemplateMsgContent
	tpl = append(tpl, tpl1)
	tpl = append(tpl, tpl2)
	result, err := rc.SystemSendTemplate(
		"7Szq13MKRVortoknTAk7W8",
		"RC:TxtMsg",
		msg,
		tpl)
	t.Log(err)
	t.Log(result)
}

func TestRongCloud_HistoryGet(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	history, err := rc.HistoryGet(
		"2018030210",
	)
	t.Log(err)
	t.Log(history)
}

func TestRongCloud_HistoryRemove(t *testing.T) {

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	err := rc.HistoryRemove(
		"2018030210",
	)
	t.Log(err)
}

func TestRongCloud_GetPrivateHistoryMessage(t *testing.T) {
	// rc := NewRongCloud(
	// 	os.Getenv("APP_KEY"),
	// 	os.Getenv("APP_SECRET"),
	// 	REGION_BJ,
	// 	WithRongCloudURI("http://xxx"),
	// )

	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	model := QueryHistoryMessageModel{
		UserID:       "uid1",
		TargetID:     "tid1",
		StartTime:    1743584071876,
		EndTime:      1743584071077,
		PageSize:     10,
		IncludeStart: true,
	}
	resp, err := rc.GetPrivateHistoryMessage(model)
	if err != nil {
		t.Errorf("GetPrivateHistoryMessage failed: %v", err)
		return
	}

	jsonResp, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		t.Errorf("JSON marshal error: %v", err)
		return
	}
	t.Logf("GetPrivateHistoryMessage response: \n%s", string(jsonResp))
}

func TestRongCloud_GetGroupHistoryMessage(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	model := QueryHistoryMessageModel{
		UserID:       "gid1",
		TargetID:     "tid01",
		StartTime:    1743584071876,
		EndTime:      1743584071077,
		PageSize:     10,
		IncludeStart: true,
	}
	resp, err := rc.GetGroupHistoryMessage(model)
	if err != nil {
		t.Errorf("GetGroupHistoryMessage failed: %v", err)
		return
	}

	jsonResp, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		t.Errorf("JSON marshal error: %v", err)
		return
	}
	t.Logf("GetGroupHistoryMessage response: \n%s", string(jsonResp))
}

func TestRongCloud_GetUltraGroupHistoryMessage(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	model := QueryHistoryMessageModel{
		UserID:       "uid01",
		TargetID:     "ugid01",
		BusChannel:   "bus1",
		StartTime:    1743584071876,
		EndTime:      1743584071077,
		PageSize:     10,
		IncludeStart: true,
	}
	resp, err := rc.GetUltraGroupHistoryMessage(model)
	if err != nil {
		t.Errorf("GetUltraGroupHistoryMessage failed: %v", err)
		return
	}

	jsonResp, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		t.Errorf("JSON marshal error: %v", err)
		return
	}
	t.Logf("GetUltraGroupHistoryMessage response: \n%s", string(jsonResp))
}

func TestRongCloud_GetChatroomHistoryMessage(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	model := QueryHistoryMessageModel{
		UserID:       "uid1",
		TargetID:     "chId01",
		StartTime:    1743584071876,
		EndTime:      1743584071077,
		PageSize:     10,
		IncludeStart: true,
	}
	resp, err := rc.GetChatroomHistoryMessage(model)
	if err != nil {
		t.Errorf("GetChatroomHistoryMessage failed: %v", err)
		return
	}

	jsonResp, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		t.Errorf("JSON marshal error: %v", err)
		return
	}
	t.Logf("GetChatroomHistoryMessage response: \n%s", string(jsonResp))
}

func TestRongCloud_ConversationMessageHistoryClean(t *testing.T) {
	rc := NewRongCloud(
		"n19jmcy59f1q9",
		"CuhqdZMeuLsKj",
		REGION_BJ,
	)

	// 最小调用校验：群组会话(3)，示例参数
	err := rc.ConversationMessageHistoryClean("3", "uid_test", "gid_test", WithCleanMsgTimestamp(1566281295943))
	if err != nil {
		t.Logf("ConversationMessageHistoryClean returned error (acceptable in unit env): %v", err)
	} else {
		t.Log("ConversationMessageHistoryClean ok")
	}
}

func TestRongCloud_PrivateStreamSend(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	// First packet
	msg := PrivateStreamMessage{
		FromUserID: "user001",
		ToUserID:   "user002",
		ObjectName: "RC:StreamMsg",
		Content: StreamMessage{
			Content:  "Hello",
			Seq:      1,
			Complete: false,
			Type:     "text",
		},
	}

	result, err := rc.PrivateStreamSend(msg)
	if err != nil {
		t.Logf("PrivateStreamSend first packet error: %v", err)
	} else {
		t.Logf("PrivateStreamSend first: code=%d, messageUID=%s", result.Code, result.MessageUID)
	}

	// Continuation packet
	msg.Content.Content = "World"
	msg.Content.Seq = 2
	msg.MessageUID = result.MessageUID
	msg.Content.MessageUID = result.MessageUID

	result, err = rc.PrivateStreamSend(msg)
	if err != nil {
		t.Logf("PrivateStreamSend cont packet error: %v", err)
	} else {
		t.Logf("PrivateStreamSend cont: code=%d, messageUID=%s", result.Code, result.MessageUID)
	}

	// End packet
	msg.Content.Content = "!"
	msg.Content.Seq = 3
	msg.Content.Complete = true

	result, err = rc.PrivateStreamSend(msg)
	if err != nil {
		t.Logf("PrivateStreamSend end packet error: %v", err)
	} else {
		t.Logf("PrivateStreamSend end: code=%d, messageUID=%s", result.Code, result.MessageUID)
	}
}

func TestRongCloud_GroupStreamSend(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)

	isMentioned := 1
	msg := GroupStreamMessage{
		FromUserID: "user001",
		ToGroupID:  "group001",
		ObjectName: "RC:StreamMsg",
		Content: StreamMessage{
			Content:  "Group stream content",
			Seq:      1,
			Complete: false,
			Type:     "text",
			// mentionedInfo is only valid on the first stream in group scenarios,
			// and requires IsMentioned = 1 to take effect.
			MentionedInfo: &MentionedInfo{
				Type:        2,
				UserIDs:     []string{"user002"},
				PushContent: "Someone mentioned you",
			},
		},
		ToUserIDs:   []string{"user002", "user003"},
		IsMentioned: &isMentioned,
	}

	result, err := rc.GroupStreamSend(msg)
	if err != nil {
		t.Logf("GroupStreamSend error: %v", err)
	} else {
		t.Logf("GroupStreamSend: code=%d, messageUID=%s", result.Code, result.MessageUID)
	}
}
