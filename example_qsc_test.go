// +build quantum

package kp_test

import (
	"context"
	"fmt"

	kp "github.com/IBM/keyprotect-go-client"
)

func ExampleQSCClient_CreateRootKey() {
	qscConfig := kp.ClientQSCConfig{
		AlgorithmID: kp.KP_QSC_ALGO_KYBER768,
	}
	client, _ := kp.NewWithLogger(
		kp.ClientConfig{
			BaseURL:    kp.DefaultBaseQSCURL,
			APIKey:     "notARealApiKey",
			InstanceID: "a6493c3a-5b29-4ac3-9eaa-deadbeef3bfd",
		},
		kp.DefaultTransport(),
		nil,
		kp.WithQSC(qscConfig),
	)
	ctx := context.Background()

	rootkey, err := client.CreateRootKey(ctx, "mynewrootkey", nil)
	if err != nil {
		fmt.Println("Error while creating root key: ", err)
	} else {
		fmt.Println("New root key created: ", *rootkey)
	}
}

func ExampleQSCClient_WrapCreateDEK() {
	qscConfig := kp.ClientQSCConfig{
		AlgorithmID: kp.KP_QSC_ALGO_KYBER768,
	}
	client, _ := kp.NewWithLogger(
		kp.ClientConfig{
			BaseURL:    kp.DefaultBaseQSCURL,
			APIKey:     "notARealApiKey",
			InstanceID: "a6493c3a-5b29-4ac3-9eaa-deadbeef3bfd",
		},
		kp.DefaultTransport(),
		nil,
		kp.WithQSC(qscConfig),
	)

	keyId := "1234abcd-abcd-asdf-9eaa-deadbeefabcd"
	aad := []string{
		"AAD can be pretty much any string value.",
		"This entire array of strings is the AAD.",
		"It has to be the same on wrap and unwrap, however",
		"This can be useful, if the DEK should be bound to an application name",
		"or possibly a hostname, IP address, or even email address.",
		"For example",
		"appname=golang-examples;",
		"It is not secret though, so don't put anything sensitive here",
	}

	ctx := context.Background()

	dek, wrappedDek, err := client.WrapCreateDEK(ctx, keyId, &aad)
	if err != nil {
		fmt.Println("Error while creating a DEK: ", err)
	} else {
		fmt.Println("Created new random DEK")
	}

	if len(dek) != 32 {
		fmt.Println("DEK length was not 32 bytes (not a 256 bit key)")
	}

	if len(wrappedDek) > 0 {
		fmt.Printf("Your WDEK is: %v\n", string(wrappedDek))
	}

	// dek is your plaintext DEK, use it for encrypt/decrypt and throw it away
	// wrappedDek is your WDEK, keep this and pass it to Unwrap to get back your DEK when you need it again
}

func ExampleQSCClient_UnwrapV2() {
	qscConfig := kp.ClientQSCConfig{
		AlgorithmID: kp.KP_QSC_ALGO_KYBER768,
	}
	client, _ := kp.NewWithLogger(
		kp.ClientConfig{
			BaseURL:    kp.DefaultBaseQSCURL,
			APIKey:     "notARealApiKey",
			InstanceID: "a6493c3a-5b29-4ac3-9eaa-deadbeef3bfd",
		},
		kp.DefaultTransport(),
		nil,
		kp.WithQSC(qscConfig),
	)

	keyId := "1234abcd-abcd-asdf-9eaa-deadbeefabcd"
	wrappedDek := []byte("dGhpcyBpc24ndCBhIHJlYWwgcGF5bG9hZAo=")
	aad := []string{
		"AAD can be pretty much any string value.",
		"This entire array of strings is the AAD.",
		"It has to be the same on wrap and unwrap, however",
		"This can be useful, if the DEK should be bound to an application name",
		"or possibly a hostname, IP address, or even email address.",
		"For example",
		"appname=golang-examples;",
		"It is not secret though, so don't put anything sensitive here",
	}

	ctx := context.Background()

	dek, rewrapped, err := client.UnwrapV2(ctx, keyId, wrappedDek, &aad)
	if err != nil {
		fmt.Println("Error while unwrapping DEK: ", err)
	} else {
		fmt.Println("Unwrapped key successfully")
	}

	if len(dek) != 32 {
		fmt.Println("DEK length was not 32 bytes (not a 256 bit key)")
	}

	// dek is your plaintext DEK, use it for encrypt/decrypt then throw it away
	// rewrapped is POSSIBLY a new WDEK, if it is not empty, store that and use it on next Unwrap

	if len(rewrapped) > 0 {
		fmt.Printf("Your DEK was rewrapped with a new key version. Your new WDEK is %v\n", rewrapped)

		// store new WDEK
		wrappedDek = rewrapped
	}

}

func ExampleQSCClient_CreateStandardKey() {
	qscConfig := kp.ClientQSCConfig{
		AlgorithmID: kp.KP_QSC_ALGO_KYBER768,
	}
	client, _ := kp.NewWithLogger(
		kp.ClientConfig{
			BaseURL:    kp.DefaultBaseQSCURL,
			APIKey:     "notARealApiKey",
			InstanceID: "a6493c3a-5b29-4ac3-9eaa-deadbeef3bfd",
		},
		kp.DefaultTransport(),
		nil,
		kp.WithQSC(qscConfig),
	)

	rootkey, err := client.CreateStandardKey(context.Background(), "mynewstandardkey", nil)
	if err != nil {
		fmt.Println("Error while creating standard key: ", err)
	} else {
		fmt.Println("New standard key created: ", *rootkey)
	}
}

func ExampleQSCClient_GetKey() {
	qscConfig := kp.ClientQSCConfig{
		AlgorithmID: kp.KP_QSC_ALGO_KYBER768,
	}
	client, _ := kp.NewWithLogger(
		kp.ClientConfig{
			BaseURL:    kp.DefaultBaseQSCURL,
			APIKey:     "notARealApiKey",
			InstanceID: "a6493c3a-5b29-4ac3-9eaa-deadbeef3bfd",
		},
		kp.DefaultTransport(),
		nil,
		kp.WithQSC(qscConfig),
	)
	keyId := "1234abcd-abcd-asdf-9eaa-deadbeefabcd"

	fmt.Println("Getting key")
	key, err := client.GetKey(context.Background(), keyId)
	if err != nil {
		fmt.Println("Get Key failed with error: ", err)
	} else {
		fmt.Printf("Key: %v\n", *key)
	}
}

func ExampleQSCClient_DeleteKey() {
	qscConfig := kp.ClientQSCConfig{
		AlgorithmID: kp.KP_QSC_ALGO_KYBER768,
	}
	client, _ := kp.NewWithLogger(
		kp.ClientConfig{
			BaseURL:    kp.DefaultBaseQSCURL,
			APIKey:     "notARealApiKey",
			InstanceID: "a6493c3a-5b29-4ac3-9eaa-deadbeef3bfd",
		},
		kp.DefaultTransport(),
		nil,
		kp.WithQSC(qscConfig),
	)
	keyId := "1234abcd-abcd-asdf-9eaa-deadbeefabcd"

	fmt.Println("Deleting standard key")
	delKey, err := client.DeleteKey(context.Background(), keyId, kp.ReturnRepresentation)
	if err != nil {
		fmt.Println("Error while deleting: ", err)
	} else {
		fmt.Println("Deleted key: ", delKey)
	}
}