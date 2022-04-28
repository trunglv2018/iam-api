package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

func validateVariadicArgs(expectedLen int, args ...interface{}) error {
	if len(args) != expectedLen {
		return fmt.Errorf("Expected %d arguments, but got %d", expectedLen, len(args))
	}

	for _, p := range args {
		_, ok := p.(string)
		if !ok {
			return errors.New("Argument must be a string")
		}
	}

	return nil
}

func KeyNotMatch2Func(args ...interface{}) (interface{}, error) {
	if err := validateVariadicArgs(2, args...); err != nil {
		return false, fmt.Errorf("%s: %s", "keyMatch2", err)
	}
	name1 := args[0].(string)
	name2 := args[1].(string)
	return bool(!util.KeyMatch2(name1, name2)), nil
}

func NotContainCustoCDFunc(args ...interface{}) (interface{}, error) {
	if err := validateVariadicArgs(4, args...); err != nil {
		return false, fmt.Errorf("%s: %s", "NotContainCustoCDFunc", err)
	}
	robj := args[0].(string)
	pobj := args[1].(string)
	paramName := args[2].(string)
	comapareValue := args[3].(string)
	var result = util.KeyGet2(robj, pobj, paramName)
	fmt.Println(result, comapareValue)
	if result == "" {
		return true, nil
	}
	return result == comapareValue, nil
}

func NewEnforceContext(suffix string) casbin.EnforceContext {
	return casbin.EnforceContext{
		RType: "r" + suffix,
		PType: "p" + suffix,
		EType: "e" + suffix,
		MType: "m" + suffix,
	}
}

// key1: regex
// key2: url
// key3: custoCD
func KeyContain(key1, key2, key3 string) bool {
	re := regexp.MustCompile(key1)
	match := re.FindStringSubmatch(key2)
	fmt.Println(key1, key2, key3, match)

	if len(match) == 0 {
		return false
	}
	return strings.ReplaceAll(match[0], "/", "") == key3
}
func KeyContainFunc(args ...interface{}) (interface{}, error) {
	if err := validateVariadicArgs(3, args...); err != nil {
		return false, fmt.Errorf("%s: %s", "keyMatch2", err)
	}
	name1 := args[0].(string)
	name2 := args[1].(string)
	name3 := args[2].(string)
	return bool(KeyContain(name1, name2, name3)), nil
}

func CustomRegexMatch(key1 string, key2 string) bool {
	match, _ := regexp.MatchString(key1, key2)
	return match
}

func CustomRegexMatchFunc(args ...interface{}) (interface{}, error) {
	if err := validateVariadicArgs(2, args...); err != nil {
		return false, fmt.Errorf("%s: %s", "keyMatch2", err)
	}
	name1 := args[0].(string)
	name2 := args[1].(string)
	return bool(CustomRegexMatch(name1, name2)), nil
}
