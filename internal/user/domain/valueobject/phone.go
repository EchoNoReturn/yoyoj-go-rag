package valueobject


import "errors"

const valid_phone_length = 10
const valid_area_code_length = 3

type Phone struct {
	Number string
	AreaCode string
}

func NewPhone(areaCode, number string) (Phone, error) {
	// 区号不是+开头自动补全
	if len(areaCode) > 0 && areaCode[0] != '+' {
		areaCode = "+" + areaCode
	}

	// 简单校验长度
	if len(areaCode) != valid_area_code_length+1 || len(number) != valid_phone_length {
		return Phone{}, errors.New("invalid phone number")
	}
	return Phone{
		AreaCode: areaCode,
		Number: number,
	}, nil
}

func (p Phone) FullNumber() string {
	return p.AreaCode + p.Number
}