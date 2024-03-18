package logic

import (
	"Tmage/controller/status"
	"Tmage/util"
	"errors"
	"strings"
)

func GetFilesByTags(tag string) (filenames []string, err error) {
	tags := strings.Split(tag, " ")
	var sendTags []string
	for _, t := range tags {
		if t == "" || t == " " {
			continue
		}
		sendTags = append(sendTags, t)
	}
	if len(sendTags) == 0 {
		return []string{}, errors.New(status.StatusNoTag.Msg())
	}

	info, filenames := util.GetFiles(sendTags)
	if info != status.StatusSuccess {
		return []string{}, errors.New(info.Msg())
	}
	return
}
