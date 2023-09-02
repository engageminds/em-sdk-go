package em

import (
	"fmt"
	"regexp"
	"strings"
)

type EventRequest struct {
	Ts         int64   `json:"ts"`                // Client time in milliseconds
	Rid        string  `json:"rid,omitempty"`     // request id
	Fit        int64   `json:"fit"`               // Client first install App time, in seconds
	Flt        int64   `json:"flt"`               // The first time the client opens the App, in seconds (uid file creation time)
	Zo         int16   `json:"zo"`                // Time zone offset in minutes. For example, Beijing time zone Zo = 480
	Tz         string  `json:"tz"`                // Time zone name
	Session    string  `json:"session,omitempty"` // Session ID, UUID generated when the app is initialized
	Ssid       string  `json:"ssid,omitempty"`    // Standardized Serviceld, generated by em-server
	Gaid       UUID    `json:"gaid,omitempty"`    // Andriod GAID
	Idfa       UUID    `json:"idfa,omitempty"`    // iOS IDFA
	Idfv       UUID    `json:"idfv,omitempty"`    // iOS IDFV
	Oaid       string  `json:"oaid,omitempty"`    // OAID
	Dtype      uint8   `json:"dtype"`             // device_type, 0:unknown,1:phone,2:tablet,3:tv
	Lang       string  `json:"lang,omitempty"`    // language code
	Jb         uint8   `json:"jb"`                // jailbreak status, 0: normal, 1: jailbreak, no transmission during normal
	Bundle     string  `json:"bundle,omitempty"`  // The current app package name
	Make       string  `json:"make,omitempty"`    // device maker
	Brand      string  `json:"brand,omitempty"`   // device brand
	Model      string  `json:"model,omitempty"`   // device model
	Os         Os      `json:"os"`                // 操作系统, 0:iOS,1:Android,2:HarmonyOS,3:Mac,4:Windows,5:Linux
	Osv        string  `json:"osv,omitempty"`     // Os version
	Appk       string  `json:"appk,omitempty"`    // pubApp key
	Appv       string  `json:"appv,omitempty"`    // app version
	Sdk        uint8   `json:"sdk"`               // 当前SDK类型, 0:iOS,1:Android,2:JS,3:ServerJava,4:ServerGo,5:ServerPython
	Sdkv       string  `json:"sdkv,omitempty"`    // sdk version
	Width      uint32  `json:"width,omitempty"`   // screen Width
	Height     uint32  `json:"height,omitempty"`  // screen Height
	Contype    uint8   `json:"contype,omitempty"` // ConnectionType
	Carrier    string  `json:"carrier,omitempty"` // 运营商名称, NetworkOperatorName
	Mccmnc     string  `json:"mccmnc,omitempty"`  // 运营商mcc+mnc, NetworkOperator
	Gcy        string  `json:"gcy,omitempty"`     // telephony network country iso
	Sco        Sco     `json:"sco"`               // ScreenOrientation, 0:unknown,1:portrait,2:landscape
	Adtk       uint8   `json:"adtk"`              // adTrackingEnable, 0:No,1:Yes
	Ntf        uint8   `json:"ntf"`               // notificationsEnabled, 0:No,1:Yes
	Gp         uint8   `json:"gp"`                // google_play_services, 0:No,1:Yes
	BasicProps DataMap `json:"bps,omitempty"`     // Custom basic values
	Ip         string  `json:"ip,omitempty"`      // Client IP
	Ua         string  `json:"ua,omitempty"`

	Events []*Event `json:"events,omitempty"`
}

func (r *EventRequest) AddEvent(e *Event) {
	if r.Events == nil {
		r.Events = make([]*Event, 0, 2)
	}
	r.Events = append(r.Events, e)
}

type EventResponse struct {
	Code   int      `json:"code,omitempty"`
	Msg    string   `json:"msg,omitempty"`
	Events []*Event `json:"events,omitempty"`
}

type Event struct {
	Ts    int64       `json:"ts,omitempty"`    // Client time in milliseconds
	Cdid  string      `json:"cdid,omitempty"`  // media user id
	Eid   string      `json:"eid,omitempty"`   // event ID
	Props DataMap     `json:"props,omitempty"` // event values
	Err   []*EventErr `json:"err,omitempty"`   // 数据校验错误信息, 返回值
}

type EventErr struct {
	Type  string `json:"type,omitempty"`
	Prop  string `json:"prop,omitempty"`
	Value any    `json:"value,omitempty"`
}

type DataMap map[string]any

func (m *DataMap) GetStr(k string) string {
	v, ok := (*m)[k]
	if !ok {
		return ""
	}
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return s
}

func (m *DataMap) TrimKeys() {
	for k, v := range *m {
		nk := strings.TrimSpace(k)
		if nk != k {
			delete(*m, k)
			(*m)[nk] = v
		}
	}
}

var uuidReg = regexp.MustCompile("^\\w{8}-\\w{4}-4\\w{3}-[abAB89]\\w{3}-\\w{12}$")

type UUID string

func (x *UUID) UnmarshalJSON(v []byte) error {
	if len(v) < 2 { // "" double quote
		return nil
	}
	s := v[1 : len(v)-1]
	if !uuidReg.Match(s) {
		return fmt.Errorf("InvalidUUID:%s", string(s))
	}
	*x = UUID(s)
	return nil
}

type Os uint8

//goland:noinspection ALL
const (
	OsIOS Os = iota
	OsAndroid
	OsHarmony
	OsMac
	OsWindows
	OsLinux
)

// Sco ScreenOrientation
type Sco uint8

//goland:noinspection ALL
const (
	ScoUnknown Sco = iota
	ScoPortrait
	ScoLandscape
)