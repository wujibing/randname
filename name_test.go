package randname

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := Rand()
		u.generateTime = 1
		u.GenerateUserName()
		u.GeneratePassword()
		if len(u.Username) > 17 {
			t.Fatal(u.Username)
		} else {
			u.Write(os.Stdout)
		}
	}
}

//转拼音
func TestRand2(t *testing.T) {
	tones := [][]rune{
		{'ā', 'ē', 'ī', 'ō', 'ū', 'ǖ', 'Ā', 'Ē', 'Ī', 'Ō', 'Ū', 'Ǖ'},
		{'á', 'é', 'í', 'ó', 'ú', 'ǘ', 'Á', 'É', 'Í', 'Ó', 'Ú', 'Ǘ'},
		{'ǎ', 'ě', 'ǐ', 'ǒ', 'ǔ', 'ǚ', 'Ǎ', 'Ě', 'Ǐ', 'Ǒ', 'Ǔ', 'Ǚ'},
		{'à', 'è', 'ì', 'ò', 'ù', 'ǜ', 'À', 'È', 'Ì', 'Ò', 'Ù', 'Ǜ'},
	}
	neutrals := []rune{'a', 'e', 'i', 'o', 'u', 'v', 'A', 'E', 'I', 'O', 'U', 'V'}
	fp, err := os.Open("C:\\Users\\Administrator\\Desktop\\pinyin.txt")
	if err != nil {
		return
	}
	tonesMap := make(map[rune]rune)
	for _, runes := range tones {
		for j, tone := range runes {
			tonesMap[tone] = neutrals[j]
		}
	}
	defer fp.Close()
	dict := make(map[rune]string)
	rd := bufio.NewReader(fp)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Fatal(err)
			}
		}
		fields := strings.Split(string(line), "=>")

		if len(fields) != 2 {
			continue
		}
		index, err := strconv.ParseInt(fields[0], 16, 32)
		if err != nil || index < 1 {
			continue
		}
		runes := []rune(fields[1])
		for k, v := range runes {
			if tonesMap[v] != 0 {
				runes[k] = tonesMap[v]
			}
		}
		dict[rune(index)] = string(runes)
	}

	dict2 := make(map[rune]string)
	for _, v := range names {
		for _, x := range v {
			b := dict[x]
			if len(b) == 0 {
				continue
			} else {
				dict2[x] = b
			}
		}
	}
	for key, value := range dict2 {
		fmt.Printf(`%d:"%s",`, key, value)
	}
}
