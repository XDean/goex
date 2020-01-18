package xconfig

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var key = "123456"

func TestGenerate(t *testing.T) {
	fmt.Println(EncryptString("a87fa5e7ca994364bd55b13c4f2f468f", "XDe@na8866999"))
	key, err := Decrypt("ENC~UYghJPCbXTcbjzh+Y3kHenQ7+Nz6q10aN9yX53QI4LUnE9m33rthwHNsx6I+q+BNO8F0IBfTCKsAQkhV", "XDe@na8866999")
	fmt.Println(key, err)
}

func TestDecrypt(t *testing.T) {
	encrypted := Encrypt([]byte("abc"), key)
	decrypted, err := Decrypt(encrypted, key)
	assert.NoError(t, err)
	assert.Equal(t, "abc", string(decrypted))
}

func TestDecode(t *testing.T) {
	type Bean struct {
		A, B string
		C    int
		D    map[int]string
	}

	//testDecode(t, "a", EncryptString("a", key))
	testDecode(t, Bean{A: "aaa",
		B: "bbb",
		C: 1,
	}, &Bean{A: EncryptString("aaa", key),
		B: EncryptString("bbb", key),
		C: 1,
	})
	testDecode(t, []string{"a", "b", "c"}, &[]string{EncryptString("a", key), EncryptString("b", key), "c"})
}

func testDecode(t *testing.T, expected, origin interface{}) {
	err := Decode(&origin, key)
	assert.NoError(t, err)
	assert.Equal(t, expected, reflect.ValueOf(origin).Elem().Interface())
}
