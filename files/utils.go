package files

import "encoding/base64"

func DecodeB64(str string) (string, error) {
	base64Text, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		//logger.ErrorLog("DecodeB64", err)
		return "", err
	}
	return string(base64Text), nil
}
