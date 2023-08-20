package tools

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestEncryptDecrypt(t *testing.T) {
	cases := []int64{1, 2, 114514, 1919810}
	fmt.Println("Test cases:", cases)
	for _, uid := range cases {
		token, err := GenerateToken(uid)
		if err != nil {
			t.Errorf("Encrypt failed")
		}
		claims, err := ParseToken(token)
		if err != nil || claims.Id != uid {
			t.Errorf("Decrypt failed")
		}
	}

	outOfDateTime = time.Microsecond
	for _, uid := range cases {
		token, err := GenerateToken(uid)
		if err != nil {
			t.Errorf("Encrypt failed")
		}
		time.Sleep(time.Second)
		claims, err := ParseToken(token)
		fmt.Println(claims, err)
		if err == nil {
			t.Errorf("out of date but decrypt successfully!")
		}
	}
}
