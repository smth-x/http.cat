package httpRequest

import (
	"fmt"
	"http.cat/src/const"
	ut "http.cat/src/utils"
	"io"
	"net/http"
	"os"
	"strconv"
)

func GetRequest(fileName string, httpCode int) error {
	path := "https://http.cat/"
	path += strconv.Itoa(httpCode)

	res, err := http.Get(path)
	if !ut.IsErrorNil(err) {
		return err
	}
	defer res.Body.Close()

	file, err := os.Create(fileName)
	if !ut.IsErrorNil(err) {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if !ut.IsErrorNil(err) {
		return err
	}
	return nil
}
func DeleteImg() {
	err := os.Remove(_const.FileName)
	if !ut.IsErrorNil(err) {
		fmt.Println(err)
		return
	}
}
