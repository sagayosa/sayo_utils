package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io"
	"math/rand"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sagayosa/sayo_utils/constant"
	sayoerror "github.com/sagayosa/sayo_utils/sayo_error"
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

func JSONPersistence(filePath string, source interface{}) error {
	bts, err := json.Marshal(source)
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, bts, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, prettyJSON.Bytes(), 0777)
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
	// listener, err := net.Listen("tcp", ":0")
	// if err != nil {
	// 	return 0, err
	// }
	// defer listener.Close()
	// addr := listener.Addr().(*net.TCPAddr)
	// return addr.Port, nil

	f := func() (int, error) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomNumber := r.Intn(50001) + 10000
		listener, err := net.Listen("tcp", StringPlus(":", strconv.Itoa(randomNumber)))
		if err != nil {
			return 0, err
		}
		defer listener.Close()
		return randomNumber, nil
	}

	for i := 0; i < constant.GetAvailablePortRandomTimes; i++ {
		number, err := f()
		if err != nil {
			continue
		}

		return number, nil
	}

	return 0, sayoerror.ErrGetAvailablePortTimesLimited
}

// Copy the field with the same name and same type in source and dest from source to dest and return dest itself
func FillSameField(source interface{}, dest interface{}) interface{} {
	valSource := reflect.ValueOf(source)
	valDest := reflect.ValueOf(dest)

	if valSource.Kind() == reflect.Ptr {
		valSource = valSource.Elem()
	}
	if valDest.Kind() == reflect.Ptr {
		valDest = valDest.Elem()
	}

	for i := 0; i < valSource.NumField(); i++ {
		fieldSource := valSource.Field(i)
		fieldDest := valDest.FieldByName(valSource.Type().Field(i).Name)

		if fieldDest.IsValid() && fieldDest.CanSet() && fieldSource.Kind() == fieldDest.Kind() {
			fieldDest.Set(fieldSource)
		}
	}

	return dest
}

func SplitIPInfo(info string) (string, string, error) {
	return net.SplitHostPort(info)
}

func Debounce(fn func(), delay time.Duration) func() {
	var mutex sync.Mutex
	var timer *time.Timer

	return func() {
		mutex.Lock()
		defer mutex.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(delay, func() {
			mutex.Lock()
			fn()
			mutex.Unlock()
		})
	}
}

func GenerateURL(addr string, url string) string {
	return StringPlus("http://", addr, "/", url)
}
