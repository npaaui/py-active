package wxhelper

import (
	"testing"
	"encoding/xml"
	"fmt"
)


// EncryptMsg 加密报文
func TestEncryptMsg(t *testing.T) {
	appId := "wxfabf18ec7ccd2d1a"
	msg := `<xml><ToUserName><![CDATA[gh_274da2028f77]]></ToUserName>
<FromUserName><![CDATA[ozmLcjnM7vnrXmb3DimFLi0EOiY8]]></FromUserName>
<CreateTime>1448604897</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[sts]]></Content>
<MsgId>6221710657841833060</MsgId>
</xml>`
	aesKeyStr := "0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQEHOAlIfMhQUV="
	// AES CBC 加密报文
	b64Enc, err := EncryptMsg(msg, aesKeyStr, appId)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%s", b64Enc)
}

func TestDecryptMsg(t *testing.T) {

	appId := "wxfabf18ec7ccd2d1a"
	aesKeyStr := "0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQEHOAlIfMhQUV"
	b64Enc := "Z8JufHXESFt4chL0Q6vusyowhizt4mpo9Zn3DkyomP7vVhFKi3ICTa1yCOs2XjSl1BaDkKUWl0lQf7psDRwJtP+YD/I6l+TCw0DrRQQyOY9Lf/4FKQ9cpBN+TyhZErDtDJN2E6Euw8VjtV0FmSqH3dGj4sPmWmEiRLldM0luY1WjW1tKGGB2x5vWwFC4piADCw5v9uPYvRk3gZCeknPHmCkCg8ERhi89J7yUuALHwheCo38+4WdQ+YCVVoj7vzZypRiytdwWxvga8OmOk3H99WJdcKQxO7UsgKtpdV/m4rhl3S+iA0HvSOXgQd3v+lAvS8eXsejFUQj92hUP+tV1wKxdg0jK1vxT1Mww0O77N5hIA38atfMMSo8IjVV+HleLbFZ3ByCiyNxrrGDh8ljqFNyVwcJJcz9ZZAnu3XOf+BQ="
	//aesKey, _ := base64.StdEncoding.DecodeString(aesKeyStr)
	// AES CBC 解密报文
	src, err := DecryptMsg(b64Enc, aesKeyStr, appId)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("%s", src)

	//数据转换
	textRequestBody := &Message{}
	xml.Unmarshal(src, textRequestBody)
	fmt.Println("textRequestBody MsgType:", textRequestBody.MsgType)
	switch textRequestBody.MsgType {
	case MsgTypeEvent:
		//订阅事件
		if textRequestBody.Event == "subscribe" {
			//更新用户订阅状态
			fmt.Println("subscribe")
		}
		if textRequestBody.Event == "unsubscribe" {
			//更新用户订阅状态
			fmt.Println("unsubscribe")
		}
		break
	case MsgTypeText:
		fmt.Println(textRequestBody.Content)
		break
	default:
		break
	}
	return
}
