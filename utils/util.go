package utils

import uuid2 "github.com/nu7hatch/gouuid"

func GetUUI() string {
	uuid, err := uuid2.NewV4()
	if err != nil {
		return ""
	}
	return uuid.String()
}
