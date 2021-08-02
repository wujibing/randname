package randname

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	User struct {
		Name         string
		Pinyin       string
		Year         uint16
		Month        uint8
		Day          uint8
		generateTime uint8
	}
)

func (u *User) GenerateUserName() string {
	if u.generateTime == 0 {
		u.generateTime = 1
		return fmt.Sprintf("%s%d", u.Pinyin, u.Year)
	} else {
		return fmt.Sprintf("%s%d%d", u.Pinyin, u.Year, rand.Intn(1000))
	}
}

func Rand() *User {
	rand.Seed(time.Now().UnixNano() + rand.Int63())
	var size int = 2
	if rand.Intn(100) < 60 {
		size = 3
	}
	runes := make([]rune, size)
	py := ""
	for i := 0; i < size; i++ {
		runes[i] = names[i].Get()
		py += pinyin[runes[i]]
	}
	u := &User{
		Name:   string(runes),
		Pinyin: py,
		Year:   uint16(1980 + rand.Intn(20)),
		Month:  uint8(rand.Intn(12) + 1),
	}

	if u.Month == 2 {
		u.Day = uint8(rand.Intn(26) + 1)
	} else {
		u.Day = uint8(rand.Intn(29) + 1)
	}
	return u
}
