package jose

import (
	"encoding/json"
	"testing"

	"github.com/icrowley/fake"
	assertion "github.com/stretchr/testify/assert"
)

type claimTestCase struct {
	title  string
	arg    interface{}
	expect map[string]interface{}
}

const (
	OTHER  = 0
	MAP    = 1
	STRUCT = 2
	ERROR  = 3

	ExportA  = "A"
	ExportB  = "B"
	Unexport = "unexport"
)

func buildClaimTestCase(kind int) *claimTestCase {
	a := fake.UserName()
	b := fake.Day()
	unexport := fake.UserName()
	tc := new(claimTestCase)

	switch kind {
	case MAP:
		tc.title = "MapTest"
		tc.arg = map[string]interface{}{
			ExportA:  a,
			ExportB:  b,
			Unexport: unexport,
		}
		tc.expect = map[string]interface{}{ExportA: a, ExportB: b, Unexport: unexport}
	case STRUCT:
		tc.title = "StructTest"
		tc.arg = struct {
			A        string
			B        int
			unexport string
		}{a, b, unexport}
		data, _ := json.Marshal(tc.arg)
		tc.expect = map[string]interface{}{Data: string(data)}
	case OTHER:
		tc.title = "OtherTest"
		tc.arg = "LOLOL"
		tc.expect = map[string]interface{}{Data: tc.arg}
	case ERROR:
		tc.title = "ErrorTest"
	}

	return tc
}

func TestInitClaimsByOther(t *testing.T) {
	// 準備
	assert := assertion.New(t)
	testCase := buildClaimTestCase(OTHER)

	// 執行
	claims, err := initClaims(testCase.arg)

	// 斷言
	assert.Nil(err)
	expect := testCase.expect
	assert.Equal(expect[Data], claims[Data])
}

func TestInitClaimsByStruct(t *testing.T) {
	// 準備
	assert := assertion.New(t)
	testCase := buildClaimTestCase(STRUCT)

	// 執行
	claims, err := initClaims(testCase.arg)

	// 斷言
	assert.Nil(err)
	expect := testCase.expect
	assert.Equal(expect[Data], claims[Data])
}

func TestInitClaimsByMap(t *testing.T) {
	// 準備
	assert := assertion.New(t)
	testCase := buildClaimTestCase(MAP)

	// 執行
	claims, err := initClaims(testCase.arg)

	// 斷言
	assert.Nil(err)
	expect := testCase.expect
	assert.Equal(expect[ExportA], claims[ExportA])
	assert.Equal(expect[ExportB], claims[ExportB])
	assert.Equal(expect[Unexport], claims[Unexport]) // map 無論如何都會全輸出
}

func TestInitClaimsByError(t *testing.T) {
	// 準備
	assert := assertion.New(t)
	testCase := buildClaimTestCase(ERROR)

	// 執行
	claims, err := initClaims(testCase.arg)

	// 斷言
	assert.Error(err)
	assert.Nil(claims)
}
