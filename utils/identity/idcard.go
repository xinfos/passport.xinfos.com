package identity

import (
	"errors"
	"strconv"
	"time"
)

type IDCard struct {
	Province string
	City     string
	District string
	Address  string
	Gender   int
	Age      int
	Birthday string
}

func IsValidCitizenNo18(citizenNo18 string) (*IDCard, error) {
	nLen := len(citizenNo18)
	if nLen != 18 {
		return nil, errors.New("长度不合法")
	}

	idCard := &IDCard{}

	idCard.Province = province(citizenNo18)
	if idCard.Province == "" {
		return nil, errors.New("身份证所在省份不合法")
	}

	idCard.City = city(citizenNo18)
	if idCard.City == "" {
		return nil, errors.New("身份证所在城市长度不合法")
	}

	idCard.District = district(citizenNo18)
	if idCard.District == "" {
		return nil, errors.New("身份证所在地区不合法")
	}

	idCard.Address = address(citizenNo18)
	if idCard.Address == "" {
		return nil, errors.New("身份证地址不合法")
	}

	idCard.Gender = gender(citizenNo18)
	if idCard.Gender < 0 || idCard.Gender > 3 {
		return nil, errors.New("性别不合法")
	}

	idCard.Age = age(citizenNo18)
	if idCard.Age <= 0 {
		return nil, errors.New("年龄不合法")
	}

	idCard.Birthday = birthday(citizenNo18)
	if idCard.Birthday == "" {
		return nil, errors.New("生日不合法")
	}

	return idCard, nil
}

func province(id string) string {
	return idDataMap[id[:2]+"0000"]
}

func city(id string) string {
	return idDataMap[id[:4]+"00"]
}

func district(id string) string {
	return idDataMap[id[:6]]
}

func address(id string) string {
	return province(id) + city(id) + district(id)
}

func gender(id string) int {
	val, _ := strconv.Atoi(id[16:17])
	if val%2 != 0 {
		return 1
	}
	return 2
}

func birthday(id string) string {
	return id[6:10] + "-" + id[10:12] + "-" + id[12:14]
}

func age(id string) int {
	nowTime := time.Now().Unix()
	formatTime, err := time.Parse("2006-01-02", birthday(id))
	if err != nil {
		return 0
	}
	olaTime := int(formatTime.Unix())
	return (int(nowTime) - olaTime) / 31536000
}
