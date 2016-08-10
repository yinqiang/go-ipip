package ipip

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"net"
)

var (
	ErrInvalidIp  = errors.New("invalid ip")
	ErrIpNotFound = errors.New("ip not found")

	field_drt = []byte("\t")
)

const (
	na = "N/A"
)

type LocationInfo struct {
	Country string
	Region  string
	City    string
	Isp     string
}

func newLocationInfo(b []byte) *LocationInfo {
	info := &LocationInfo{
		Country: na,
		Region:  na,
		City:    na,
		Isp:     na,
	}

	fields := bytes.Split(b, field_drt)

	switch len(fields) {
	case 4:
		// free version
		info.Country = string(fields[0])
		info.Region = string(fields[1])
		info.City = string(fields[2])

	case 5:
		// pay version
		info.Country = string(fields[0])
		info.Region = string(fields[1])
		info.City = string(fields[2])
		info.Isp = string(fields[4])

	default:
		panic("unknow ip info:" + string(b))
	}

	return info
}

type Ipip struct {
	offset uint32
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

func (p *Ipip) Find(ipstr string) (*LocationInfo, error) {
	ip := net.ParseIP(ipstr).To4()
	if ip == nil {
		return nil, ErrInvalidIp
	}

	tmp_offset := uint32(ip[0]) * 4
	start := binary.LittleEndian.Uint32(p.index[tmp_offset : tmp_offset+4])

	nip := binary.BigEndian.Uint32(ip)
	var index_offset uint32 = 0
	var index_length uint32 = 0
	var max_comp_len uint32 = p.offset - 1028
	start = start*8 + 1024

	for start < max_comp_len {
		n := binary.BigEndian.Uint32(p.index[start : start+4])
		if n >= nip {
			tmp_index := []byte{0, 0, 0, 0}
			copy(tmp_index, p.index[start+4:start+7])
			index_offset = binary.LittleEndian.Uint32(tmp_index)
			index_length = uint32(p.index[start+7])
			break
		}
		start += 8
	}

	if index_offset == 0 {
		return nil, ErrIpNotFound
	}

	res_offset := p.offset + index_offset - 1024
	return newLocationInfo(p.binary[res_offset : res_offset+index_length]), nil
}
