// Reffered from: https://github.com/DakotaNelson/sneaky-creeper/blob/master/sneakers/base94.py
package base94

import (
	"fmt"
	"strings"
)

type Base94 struct {
	Alphabet []string
	Base     int
}

func New() *Base94 {
	elements := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", "")
	return &Base94{
		Alphabet: elements,
		Base:     len(elements),
	}
}

func (base94 *Base94) Encode(num int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("%s", "only positive numbers can be encoded")
	}

	if num == 0 {
		return base94.Alphabet[0], nil
	}

	var holder string

	for num != 0 {
		holder = (base94.Alphabet[num%base94.Base]) + holder
		num = num / base94.Base
	}

	return holder, nil
}

func memorizedElementIndex(slice []string) func(string) int {
	cache := make(map[string]int)
	return func(value string) int {

		if val, found := cache[value]; found {
			return val
		}

		for i := 0; i < len(slice); i++ {
			if slice[i] == value {
				cache[value] = i
				return i
			}
		}

		return -1
	}
}

func (base94 *Base94) Decode(value string) int {
	res := 0
	mult := 1

	elementIndex := memorizedElementIndex(base94.Alphabet)

	for i := len(value) - 1; i > -1; i-- {
		res += mult * elementIndex(string(value[i]))
		mult *= base94.Base
	}

	return res
}
