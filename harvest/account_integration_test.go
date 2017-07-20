// +build integration

package harvest_test

import (
	"testing"
)


func TestWhoAmI(t *testing.T) {
	
	_, err := client.Account.WhoAmI()

	if err != nil {
		t.Fatalf("Account.WhoAmI returned %v", err)
		return
	}
}