package randname

import (
	"fmt"
	"io"
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
		Username     string
		Password     string
		generateTime uint8
	}
)

//写入文件
func (u *User) Write(writer io.Writer) {
	fmt.Fprintf(writer, "%s\t%s\t%s\t%d\t%d\t%d\r\n", u.Username, u.Password, u.Name, u.Year, u.Month, u.Day)
}
func (u *User) GenerateUserName() {
	if u.generateTime == 0 {
		u.generateTime = 1
		u.Username = fmt.Sprintf("%s%d", u.Pinyin, u.Year)
	} else {
		size := 17 - len(u.Pinyin) - 4
		if size < 4 {
			u.Username = fmt.Sprintf("%s%s", u.Pinyin, RandomBytes(2, 4, false))
		} else {
			u.Username = fmt.Sprintf("%s%s%d", u.Pinyin, RandomBytes(2, 4, false), u.Year)
		}
	}
}

func (u *User) GeneratePassword() {
	u.Password = RandomBytes(8, 10, true)
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
