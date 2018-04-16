package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func containsNode(arr []XNode, node XNode) bool {
	for _, a := range arr {
		if a.URL == node.URL && a.Address == node.Address {
			return true
		}
	}
	return false
}

// FormatNodeURL formats the url for nodeID
func FormatNodeURL(nodeID string) string {
	return fmt.Sprintf("localhost:%s", nodeID)
}

func doEvery(d time.Duration, f func()) {
	for x := range time.Tick(d) {
		fmt.Printf("%s: Scheduled mining\n", x.Format(time.RFC3339))
		f()
	}
}

func generateRandomString() string {
	randData := make([]byte, 20)
	_, err := rand.Read(randData)
	if err != nil {
		log.Panic(err)
	}

	return fmt.Sprintf("%x", randData)
}

func getPubkeyHashFromAddress(address []byte) []byte {
	pubKeyHash := Base58Decode(address)
	return pubKeyHash[1 : len(pubKeyHash)-4]
}
