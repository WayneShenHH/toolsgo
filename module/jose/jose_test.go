package jose

import (
	"testing"

	"github.com/icrowley/fake"
	assertion "github.com/stretchr/testify/assert"
)

type testCase struct {
	title  string
	args   interface{}
	expect interface{}
}

type tester struct {
	A        string
	B        int
	unexport string
}

func buildJWTTestCase(kind int) *testCase {
	tc := new(testCase)
	a := fake.UserName()
	b := fake.Day()
	unexport := fake.UserName()

	switch kind {
	case STRUCT:
		tc.title = "StructTest"
		tc.args = tester{a, b, unexport}
		tc.expect = tester{A: a, B: b}
	case MAP:
		tc.title = "MapTest"
		tc.args = map[string]interface{}{ExportA: a, ExportB: b, Unexport: unexport}
		tc.expect = map[string]interface{}{ExportA: a, ExportB: float64(b), Unexport: unexport}
	case OTHER:
		tc.title = "OtherTest"
		tc.args = []interface{}{a, b, struct{}{}}
		tc.expect = map[string]interface{}{Data: []interface{}{a, float64(b), map[string]interface{}{}}}
	case ERROR:
		tc.title = "ErrorTest"
		tc.args = nil
	}

	return tc
}

func TestStructJWT(t *testing.T) {
	// 準備
	assert := assertion.New(t)
	tc := buildJWTTestCase(STRUCT)
	temp := new(tester)

	// 執行
	jwt, err1 := GenerateJWT(tc.args)
	err2 := Bind(jwt, temp)

	// 斷言
	assert.Nil(err1)
	assert.Nil(err2)

	expect := tc.expect.(tester)
	assert.Equal(expect.A, temp.A)
	assert.Equal(expect.B, temp.B)
	assert.Equal(expect.unexport, temp.unexport)
}

func TestMapJWT(t *testing.T) {
	// 準備
	assert := assertion.New(t)
	tc := buildJWTTestCase(MAP)

	// 執行
	jwt, err1 := GenerateJWT(tc.args)
	actual, err2 := ParseJWT(jwt)

	// 斷言
	assert.Nil(err1)
	assert.Nil(err2)

	expect := tc.expect.(map[string]interface{})
	assert.Equal(expect[ExportA], actual[ExportA])
	assert.Equal(expect[ExportB], actual[ExportB])
	assert.Equal(expect[Unexport], actual[Unexport])
}

func TestOtherJWT(t *testing.T) {
	// 準備
	assert := assertion.New(t)
	tc := buildJWTTestCase(OTHER)

	// 執行
	jwt, err1 := GenerateJWT(tc.args)
	actual, err2 := ParseJWT(jwt)

	// 斷言
	assert.Nil(err1)
	assert.Nil(err2)

	expect := tc.expect.(map[string]interface{})
	assert.Equal(expect[Data], actual[Data])
}

func TestErrorJWT(t *testing.T) {
	// 準備
	assert := assertion.New(t)
	tc := buildJWTTestCase(ERROR)

	// 執行
	jwt, err1 := GenerateJWT(tc.args)
	_, err2 := ParseJWT(jwt)

	// 斷言
	assert.Error(err1)
	assert.Error(err2)
}
