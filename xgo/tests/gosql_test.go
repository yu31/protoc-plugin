package tests

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/yu31/protoc-plugin/xgo/tests/gosqltest"

	gogoproto "github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/stretchr/testify/require"
)

// Test the code generated correct.
func Test_GoSQL_Generated(t *testing.T) {
	var ok bool
	var datas []interface{}

	datas = append(datas, &gosqltest.User1{})
	datas = append(datas, &gosqltest.User1_Meta1{})
	datas = append(datas, &gosqltest.User2{})
	datas = append(datas, &gosqltest.User3{})
	datas = append(datas, &gosqltest.User4{})
	for _, data := range datas {
		_, ok = data.(sql.Scanner)
		require.True(t, ok, fmt.Sprintf("raw: %v", data))
		_, ok = data.(driver.Valuer)
		require.True(t, ok, fmt.Sprintf("raw: %v", data))
	}

	// No sets.
	datas = datas[0:0]
	datas = append(datas, &gosqltest.User1_Meta1_Meta2{})
	datas = append(datas, &gosqltest.User5{})
	datas = append(datas, &gosqltest.User6{})
	datas = append(datas, &gosqltest.User7{})
	for _, data := range datas {
		_, ok = data.(sql.Scanner)
		require.False(t, ok, fmt.Sprintf("raw: %v", data))
		_, ok = data.(driver.Valuer)
		require.False(t, ok, fmt.Sprintf("raw: %v", data))
	}
}

func Test_GoSQL_Serialize_JSON(t *testing.T) {
	user := gosqltest.User1{
		Id: "xxx-1",
		Meta1: &gosqltest.User1_Meta1{
			Age:   999,
			Meta1: &gosqltest.User1_Meta1_Meta2{Sex: 1},
		},
	}

	// Test marshal.
	b1, err := user.Value()
	require.Nil(t, err)
	b2, err := json.Marshal(&user)
	require.Nil(t, err)
	require.Equal(t, b1, b2)

	// Test unmarshal.
	userA := new(gosqltest.User1)
	userB := new(gosqltest.User1)
	err = userA.Scan(b1)
	require.Nil(t, err)
	err = json.Unmarshal(b2, userB)
	require.Nil(t, err)
	require.Equal(t, userA, userB)
	require.True(t, reflect.DeepEqual(userA, userB))
}

func Test_GoSQL_Serialize_ProtoJSON(t *testing.T) {
	user := gosqltest.User2{Id: "xxx-2"}

	// Test marshal.
	b1, err := user.Value()
	require.Nil(t, err)
	b2, err := protojson.Marshal(&user)
	require.Nil(t, err)
	require.NotEqual(t, b1, b2)

	_marshal := protojson.MarshalOptions{
		Multiline:       false,
		Indent:          " ",
		AllowPartial:    false,
		UseProtoNames:   false,
		UseEnumNumbers:  true,
		EmitUnpopulated: true,
	}

	b3, err := _marshal.Marshal(&user)
	require.Nil(t, err)
	require.Equal(t, b1, b3)

	// Test unmarshal.
	userA := new(gosqltest.User2)
	userB := new(gosqltest.User2)
	userC := new(gosqltest.User2)
	err = userA.Scan(b1)
	require.Nil(t, err)
	err = protojson.Unmarshal(b2, userB)
	require.Nil(t, err)
	require.Equal(t, userA, userB)
	require.True(t, reflect.DeepEqual(userA, userB))

	_unmarshal := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	err = _unmarshal.Unmarshal(b3, userC)
	require.Nil(t, err)
	require.Equal(t, userA, userC)
	require.True(t, reflect.DeepEqual(userA, userC))
}

func Test_GoSQL_Serialize_Proto(t *testing.T) {
	user := gosqltest.User3{Id: "xxx-3"}

	// Test marshal.
	b1, err := user.Value()
	require.Nil(t, err)
	b2, err := proto.Marshal(&user)
	require.Nil(t, err)
	require.Equal(t, b1, b2)

	// Test unmarshal.
	userA := new(gosqltest.User3)
	userB := new(gosqltest.User3)
	err = userA.Scan(b1)
	require.Nil(t, err)
	err = proto.Unmarshal(b2, userB)
	require.Nil(t, err)
	require.Equal(t, userA, userB)
	require.True(t, reflect.DeepEqual(userA, userB))
}

func Test_GoSQL_Serialize_GoGoProto(t *testing.T) {
	// Test marshal.
	user := gosqltest.User4{Id: "xxx-4"}
	b1, err := user.Value()
	require.Nil(t, err)
	b2, err := gogoproto.Marshal(&user)
	require.Nil(t, err)
	require.Equal(t, b1, b2)
}
