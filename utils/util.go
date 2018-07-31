package utils

import uuid2 "github.com/nu7hatch/gouuid"

/**
 *
 * :=  created by:  Shuza
 * :=  create date:  23-Apr-18
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 *
 **/

func GetUUI() string {
	uuid, err := uuid2.NewV4()
	if err != nil {
		return ""
	}
	return uuid.String()
}
