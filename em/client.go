package em

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

const (
	VERSION = "1.0"
	UA      = "em-sdk-go/" + VERSION
)

type Config struct {
	ServerUrl  string
	Appk       string
	HttpClient *http.Client
}

type Client struct {
	cfg *Config
}

func NewClient(c *Config) (*Client, error) {
	if c == nil {
		return nil, errors.New("NilConfig")
	}
	if c.Appk == "" {
		return nil, errors.New("AppkRequired")
	}
	if c.ServerUrl == "" {
		c.ServerUrl = "https://a.engageminds.ai"
	}
	if c.HttpClient == nil {
		c.HttpClient = http.DefaultClient
	}
	return &Client{cfg: c}, nil
}

func (c *Client) Track(r *EventRequest) (*EventResponse, error) {
	return c.TrackBatch([]*EventRequest{r})
}

func (c *Client) TrackBatch(rs []*EventRequest) (*EventResponse, error) {
	if len(rs) == 0 {
		return nil, errors.New("EmptyEventRequests")
	}
	for _, r := range rs {
		if r.Appk == "" {
			r.Appk = c.cfg.Appk
		}
		r.Sdk = 4 // 4:ServerGo
		r.Sdkv = VERSION
	}
	bs, err := json.Marshal(rs)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, int(float32(len(bs))*0.7)))
	out, err := flate.NewWriter(buf, flate.DefaultCompression)
	if err != nil {
		return nil, err
	}
	if _, err := out.Write(bs); err != nil {
		return nil, err
	}
	if err := out.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, c.cfg.ServerUrl+"/s2s/es", buf)
	req.Header.Set("Content-Encoding", "deflate")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", UA)

	res, err := c.cfg.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 || !strings.HasPrefix(res.Header.Get("Content-Type"), "application/json") {
		rv := &EventResponse{
			Code: res.StatusCode,
			Msg:  string(body),
		}
		return rv, nil
	}

	rv := &EventResponse{}
	err = json.Unmarshal(body, rv)
	return rv, err
}
