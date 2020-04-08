package base

import (
	"errors"
	"regexp"

	"github.com/spf13/cast"
)

// TODO: if merge bus plugin & net plugin, this file can be moved to plugin/netPlugin dir

const (
	ProtoTypeUnknown ProtoType = "unknown"
	ProtoTypeTcp     ProtoType = "tcp"
	ProtoTypeUdp     ProtoType = "udp"
	ProtoTypeHttp    ProtoType = "http"
	ProtoTypeHttps   ProtoType = "https"
	ProtoTypeWs      ProtoType = "ws"
	ProtoTypeWss     ProtoType = "wss"
)

func ProtoTypeToStr(t ProtoType) string {
	return cast.ToString(t)
}

func StrToProtoType(t string) ProtoType {
	switch t {
	case "tcp":
		return ProtoTypeTcp
	case "udp":
		return ProtoTypeUdp
	case "http":
		return ProtoTypeHttp
	case "https":
		return ProtoTypeHttps
	case "ws":
		return ProtoTypeWs
	case "wss":
		return ProtoTypeWss
	default:
		return ProtoTypeUnknown
	}
}

type Endpoint struct {
	isIpv6 bool
	ext    struct {
		Proto ProtoType
		Ip    string
		Port  uint16
		Path  string
	}
}

func NewFromString(url string) (*Endpoint, error) {
	if url == "" {
		return nil, errors.New("AFEndpoint url is empty")
	}

	r, err := regexp.Compile("((.*)://)?([^:/]+)(:(\\d+))?(/.*)?")
	if err != nil {
		return nil, err
	}

	if r.MatchString(url) {
		return nil, errors.New("unmatched url ` " + url + " `")
	}

	strArr := r.FindStringSubmatch(url)

	port, err := cast.ToUint16E(strArr[5])
	if err != nil {
		return nil, err
	}

	ep := &Endpoint{
		isIpv6: false,
		ext: struct {
			Proto ProtoType
			Ip    string
			Port  uint16
			Path  string
		}{
			Proto: StrToProtoType(strArr[2]),
			Ip:    strArr[3],
			Port:  port,
			Path:  strArr[6],
		},
	}
	return ep, nil
}

func (a *Endpoint) ToString() string {
	var url string
	if a.ext.Proto != ProtoTypeUnknown {
		url += string(a.ext.Proto)
	}

	url += a.GetIP() + ":" + cast.ToString(a.GetPort()) + a.GetPath()

	return url
}

//******* GET & SET ********
func (a *Endpoint) Proto() ProtoType {
	return a.ext.Proto
}

func (a *Endpoint) SetProto(proto ProtoType) {
	a.ext.Proto = proto
}

func (a *Endpoint) GetIP() string {
	return a.ext.Ip
}

func (a *Endpoint) SetIP(ip string) {
	a.ext.Ip = ip
}

func (a *Endpoint) GetPath() string {
	return a.ext.Path
}

func (a *Endpoint) SetPath(path string) {
	a.ext.Path = path
}

func (a *Endpoint) GetPort() uint16 {
	return a.ext.Port
}

func (a *Endpoint) SetPort(port uint16) {
	a.ext.Port = port
}

func (a *Endpoint) IsV6() bool {
	return a.isIpv6
}

func (a *Endpoint) SetIsV6(v6 bool) {
	a.isIpv6 = v6
}
