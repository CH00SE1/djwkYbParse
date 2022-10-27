package service

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

/**
 * @title apiMedicine
 * @author CH00SE1
 * @date 2022-10-27 16:38:49
 */

func RequestMedicineInfo(url, text, method string, t any) ([]byte, error) {
	payload := strings.NewReader(text)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("response error:", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return body, err
}
