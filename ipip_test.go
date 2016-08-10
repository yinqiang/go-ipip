package ipip

import (
	"fmt"
	"testing"
)

func Test_Load(t *testing.T) {
	p := NewIpip()
	if err := p.Load("17monipdb.dat"); err != nil {
		t.Fatal(err)
	}
}

func Test_Find(t *testing.T) {
	p := NewIpip()
	if err := p.Load("17monipdb.dat"); err != nil {
		t.Fatal(err)
	}
	if _, err := p.Find("110.172.245.98"); err != nil {
		t.Fatal(err)
	}
	if _, err := p.Find("google.com"); err != nil && err != ErrInvalidIp {
		t.Fatal(err)
	}
	if _, err := p.Find("110.172.245"); err != nil && err != ErrInvalidIp {
		t.Fatal(err)
	}
	if _, err := p.Find("aaa"); err != nil && err != ErrInvalidIp {
		t.Fatal(err)
	}
}
