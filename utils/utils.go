package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func StringPlus(segments ...string) string {
	var builder strings.Builder

	for _, seg := range segments {
		builder.WriteString(seg)
	}

	return builder.String()
}

func SHA256(filePath string) (res string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}

	return string(hash.Sum(nil)), nil
}

func JSON(filePath string, dst interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bts, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bts, dst); err != nil {
		return err
	}

	return nil
}

func Post(URL string, data interface{}) (code int, body []byte, err error) {
	bts, err := json.Marshal(data)
	if err != nil {
		return
	}

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(bts))
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return resp.StatusCode, body, nil
}

func Get(URL string, data map[string]interface{}) (code int, body []byte, err error) {
	params := url.Values{}
	for k, v := range data {
		params.Add(k, fmt.Sprintf("%v", v))
	}

	reqURL, err := url.Parse(URL)
	if err != nil {
		return
	}
	reqURL.RawQuery = params.Encode()

	resp, err := http.Get(reqURL.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return resp.StatusCode, body, nil
}

func UnMarshalUnknownAny(source interface{}, dest interface{}) error {
	bts, err := json.Marshal(source)
	if err != nil {
		return err
	}

	return json.Unmarshal(bts, dest)
}

func ChangeRoutineWorkDir(workDir string) error {
	return os.Chdir(workDir)
}

// could there be a better way?
func GetAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}
