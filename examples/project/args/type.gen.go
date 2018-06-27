// Code generated by 'micro gen' command.
// DO NOT EDIT!

package args

import (
	"github.com/henrylee2cn/teleport/codec"
	"github.com/xiaoenai/tp-micro/model/mongo"
)

// EmptyStruct alias of type struct {}
type EmptyStruct = codec.PbEmpty

// HomeResult home result
type HomeResult struct {
	Content string `json:"content"` // text
}

// DivideArg divide api arg
type DivideArg struct {
	// dividend
	A float64 `json:"a"`
	// divisor
	B float64 `param:"<range: 0.01:100000>" json:"b"`
}

// DivideResult divide api result
type DivideResult struct {
	// quotient
	C float64 `json:"c"`
}

// StatArg stat handler arg
type StatArg struct {
	Ts int64 `json:"ts"` // timestamps
}

// User user info
type User struct {
	Id        int64  `key:"pri" json:"id"`
	Name      string `key:"uni" json:"name"`
	Age       int32  `json:"age"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
	DeletedTs int64  `json:"deleted_ts"`
}

// Log comment...
type Log struct {
	Id        int64  `json:"id" key:"pri"`
	Text      string `json:"text"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
	DeletedTs int64  `json:"deleted_ts"`
}

// Device comment...
type Device struct {
	UUID      string `key:"pri" json:"uuid"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
	DeletedTs int64  `json:"deleted_ts"`
}

// Meta comment...
type Meta struct {
	Id        mongo.ObjectId `json:"_id" bson:"_id" key:"pri"`
	Hobby     []string       `json:"hobby" bson:"hobby"`
	Tags      []string       `json:"tags" bson:"tags"`
	UpdatedAt int64          `json:"updated_at" bson:"updated_at"`
	CreatedAt int64          `json:"created_at" bson:"created_at"`
	DeletedTs int64          `json:"deleted_ts" bson:"deleted_ts"`
}
