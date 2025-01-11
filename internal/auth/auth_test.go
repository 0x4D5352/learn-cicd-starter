package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://www.example.com", strings.NewReader("hello"))
	empty, err := GetAPIKey(req.Header)
	if empty != "" || err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected failure, got str %s and err %v", empty, err)
	}
	req.Header.Add("Authorization", "onlyone")
	short, err := GetAPIKey(req.Header)
	if short != "" || err == nil {
		t.Fatalf("expected failure, got str %s and err %v", short, err)
	}
	req.Header.Set("Authorization", "Bearer Token Too Long")
	long, err := GetAPIKey(req.Header)
	if long != "" || err == nil {
		t.Fatalf("expected failure, got str %s and err %v", long, err)
	}
	apiToken := "foobar1234567890"
	req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", apiToken))
	token, err := GetAPIKey(req.Header)
	if token != "foobar1234567890" || err != nil {
		t.Fatalf("expected token 'foobar1234567890', got str %s and err %v", token, err)
	}
}
