package ip

import (
	"testing"
)

func TestLocalIpv4Addrs(t *testing.T) {
	ips, err := LocalIpv4Addrs()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", ips)
	if len(ips) == 0 {
		t.Error("not found")
	}

	// loopback is excluded
	for _, ip := range ips {
		if ip == "127.0.0.1" {
			t.Errorf("loopback not excluded")
		}
	}
}
