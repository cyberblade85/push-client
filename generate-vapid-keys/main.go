package main

import (
	webpush "github.com/SherClockHolmes/webpush-go"
)



func main() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		// TODO: Handle error
	}

	println("Vapid Public key: ", publicKey)
	println("Vapid Private key: ", privateKey)

}
