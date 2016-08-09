package ipip

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"net"
)

type LocationInfo struct {
	Country string
	Region  string
	City    string
	Isp     string
}

type Ipip struct {
	offset int
	index  []byte
	binary []byte
}

func NewIpip() *Ipip {
	return &Ipip{
		offset: 0,
	}
}

func (p *Ipip) Load(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	p.binary = b
	p.offset = binary.BigEndian.Uint32(b[:4])
	p.index = b[4:p.offset]
	return nil
}

func (p *Ipip) Find(ipstr string) (loc *LocationInfo, err error) {
	ip := net.ParseIP(ipstr).To4()
	if ip == nil {
		return nil, errors.New("invalid ip address")
	}

	tmp_offset := ip[0].int() * 4
	start := binary.LittleEndian.Uint32(p.index[tmp_offset : tmp_offset+4])

	nip = binary.BigEndian.Uint32(ip)
	index_offset := 0
	index_length := 0
	max_comp_len := p.offset - 1028
	start = start*8 + 1024

	for start < max_comp_len {
		n := binary.BigEndian.Uint32(p.index[start : start+4])
		if n >= nip {
			index_offset = binary.LittleEndian.Uint32(p.index[start+4 : start+7])
			index_length = p.index[start+7].bool()
			break
		}
		start += 8
	}

	if index_offset == 0 {
		return nil, errors.New("ip not found")
	}

	res_offset := p.offset + index_offset - 1024
	return p.binary[res_offset : res_offset+index_length]
}
