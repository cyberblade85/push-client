package main

import (
	"encoding/json"

	webpush "github.com/SherClockHolmes/webpush-go"
)

const (
	subscription    = `{"endpoint":"https://fcm.googleapis.com/fcm/send/ff0MeYl9SbU:APA91bHNzDtSXrEU8sTl4upQENMvgJGiLQ2Mxy9yMwbkPH4UhpXQcWvpL4QqNXOdKzr2x9nkW_f6Qn4jol3HpnfRh5HTPsjZZQJI6Iv1h4rcpNbbXhEwl0Q5ZeWERuUxQzxSsmWU3sVs","expirationTime":null,"keys":{"p256dh":"BIuxiJOuIRR5rQ62XWaRph3WVVMQIvS7dpl3HKdOQy1FfffMueTDMRZ6Fh1CDNRcyqmTFbqBXcr_WJDHvXN-P0I","auth":"i5HKXBHSL5XlEDfRnkjliw"}}`
	vapidPublicKey  = "BHF4CqKMR1IceHhUOomb1Yr3ZLgtC7zPRttt8nDMC7Rzm-ZQvKwlUMv0NoWcJh1xSFEk8hsuBBYhgiKJO7KcsRQ"
	vapidPrivateKey = "VZiosZzpVr2PUwasgUJThYlnERfpsW4n5caV1-pLV2E"
)
func main() {
	// Decode subscription
	s := &webpush.Subscription{}
	// s.Endpoint = "https://fcm.googleapis.com/fcm/send/dfkp7ohUn14:APA91bEaiBA0AJswY04oYym64ZanSuTtOuf2Y5QhAWxTec-k67fA7692SM78iCkrymsI1B9_zRsdEWGfBSE3Xjb_JEZZJQyzujg6jYT8OK6bZTrHxKJr3qomDkF14VToKVXCAyuXb5WN"
	// s.Keys.Auth= "tXdYpvz5xi6blRtR5jowVQ"
	// s.Keys.P256dh= "BLpf9QB-2mPCcJr93K7fG6c9vXHFFwbna73gyUXbExW7kxicNWn0Yy_AX7upzS5zwmzrMoUbvzBTpdOVW-hrn20"
	
	json.Unmarshal([]byte(subscription), s)
  //  fmt.Println(s.Endpoint)
	

	// Send Notification
	resp, err := webpush.SendNotification([]byte("amcik olsana"), s, &webpush.Options{
		Subscriber:      "edip.okten@hotmail.com", // Do not include "mailto:"
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		//TTL:             30,
	})
	if err != nil {
		println(err)
	}
	println(resp)
	defer resp.Body.Close()
}
