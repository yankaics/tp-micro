// Code generated by 'micro gen' command.
// The temporary code used to ensure successful compilation!
// When the project is completed, it should be removed!

package logic

import (
	tp "github.com/henrylee2cn/teleport"

	"github.com/xiaoenai/tp-micro/examples/project/args"
	// "github.com/xiaoenai/tp-micro/examples/project/logic/model"
	// "github.com/xiaoenai/tp-micro/examples/project/rerrs"
)

// Stat handler
func Stat(ctx tp.PushCtx, arg *args.StatArg) *tp.Rerror {
	return nil
}

// Home handler
func Home(ctx tp.PullCtx, arg *args.EmptyStruct) (*args.HomeResult, *tp.Rerror) {
	return new(args.HomeResult), nil
}

// Divide handler
func Math_Divide(ctx tp.PullCtx, arg *args.DivideArg) (*args.DivideResult, *tp.Rerror) {
	return new(args.DivideResult), nil
}
