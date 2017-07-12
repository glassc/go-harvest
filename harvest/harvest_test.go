package harvest

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil, "myaccount")

	if actual, expected := c.baseURL.String(), "https://myaccount.harvestapp.com"; actual != expected {
		t.Errorf("NewClient baseURL is %v, expected %v", actual, expected)
	}
}