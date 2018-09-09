package sdk_test

import (
	"encoding/json"
	"fmt"

	tp "github.com/henrylee2cn/teleport"

	"github.com/xiaoenai/tp-micro/examples/project/sdk"
)

func toJsonBytes(i interface{}) []byte {
	b, _ := json.MarshalIndent(i, "", "  ")
	return b
}

func ExampleStat() {
	rerr := sdk.Stat(nil, &sdk.StatArg{})
	if rerr != nil {
		tp.Errorf("Stat: rerr: %s", toJsonBytes(rerr))
	}
	fmt.Printf("")
	// Output:
}

func ExampleHome() {
	result, rerr := sdk.Home(nil, &sdk.EmptyStruct{})
	if rerr != nil {
		tp.Errorf("Home: rerr: %s", toJsonBytes(rerr))
	} else {
		tp.Infof("Home: result: %s", toJsonBytes(result))
	}
	fmt.Printf("")
	// Output:
}

func ExampleMath_Divide() {
	result, rerr := sdk.Math_Divide(nil, &sdk.DivideArg{})
	if rerr != nil {
		tp.Errorf("Math_Divide: rerr: %s", toJsonBytes(rerr))
	} else {
		tp.Infof("Math_Divide: result: %s", toJsonBytes(result))
	}
	fmt.Printf("")
	// Output:
}
