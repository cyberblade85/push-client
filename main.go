package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	webpush "github.com/SherClockHolmes/webpush-go"
)

const (
	subscription    = `{"endpoint":"https://fcm.googleapis.com/fcm/send/fMUtDWbXoMw:APA91bGL5HsNCxu9QEDB8mefnipF74twF4Bq_9wC4e2lrZjUGPPDoxWCKM_nLU9rng9I8xTE6JbAfGkGNji3Se-d0G-aYsq0RzvYVctBuVm2JUjCWMPptFD_gAx8NKJcBYb4mBm2Y6y4","expirationTime":null,"keys":{"p256dh":"BGm0m1O8aet-Wrwzy9j08T4cplxhXas8nVW8pDkrBEMJuOqG_55S1lSF4HBjJRwupUbCLiFt6SvRzuxQ2XmWpBo","auth":"qzHEnU4KtIqRnuCyZWatCw"}}`
	vapidPublicKey  = "BIZoty0iCIkAuR4rZtLEaixmbIL3-Xwg3NO8BkX6-vrcIX0UHdZ8Ux9JojPhA8AutoUNANGPHxayJiv4IKcM8po"
	vapidPrivateKey = "ozhrL0H8JzaI1zZZSrag4Eu-1_MB_CiT3cuuSErAS_I"
)
func main() {
	// Decode subscription
	s := &webpush.Subscription{}
	// s.Endpoint = "https://fcm.googleapis.com/fcm/send/dfkp7ohUn14:APA91bEaiBA0AJswY04oYym64ZanSuTtOuf2Y5QhAWxTec-k67fA7692SM78iCkrymsI1B9_zRsdEWGfBSE3Xjb_JEZZJQyzujg6jYT8OK6bZTrHxKJr3qomDkF14VToKVXCAyuXb5WN"
	// s.Keys.Auth= "tXdYpvz5xi6blRtR5jowVQ"
	// s.Keys.P256dh= "BLpf9QB-2mPCcJr93K7fG6c9vXHFFwbna73gyUXbExW7kxicNWn0Yy_AX7upzS5zwmzrMoUbvzBTpdOVW-hrn20"
	
	json.Unmarshal([]byte(subscription), s)
	

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
	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
	defer resp.Body.Close()
}
