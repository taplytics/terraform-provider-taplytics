package taplytics_tf

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("c4863f6e3fabbe05b530139069573963a8b3a45f")
	featureFlags, err := client.getFeatureFlags("test")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(featureFlags)

	bucketing, err := client.getUserBucketing("test")
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(bucketing)
}