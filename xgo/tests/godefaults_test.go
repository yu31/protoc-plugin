package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yu31/protoc-plugin/xgo/tests/godefaultstest"
)

func Test_GoDefaults_LiteralMessage1_1(t *testing.T) {
	var msg *godefaultstest.LiteralMessage1
	require.NotPanics(t, msg.SetDefaults)
	require.Nil(t, msg)

	msg = &godefaultstest.LiteralMessage1{}
	require.NotPanics(t, msg.SetDefaults)

	require.Equal(t, msg.TString1, `ts1`)
	require.Equal(t, msg.TString2, ``)
	require.Equal(t, msg.TString3, `""`)
	require.Equal(t, msg.TString4, `"`)
	require.Equal(t, msg.TString5, `"ts5"`)
	require.Equal(t, msg.TString6, `"ts"6"`)
	require.Equal(t, msg.TString7, `"ts"7"`)
	require.Equal(t, msg.TString8, `[ts8]`)
	require.Equal(t, msg.TString9, `{ts9}`)
	require.Equal(t, msg.TString10, ``)

	require.Equal(t, msg.TInt32, int32(1))
	require.Equal(t, msg.TInt64, int64(2))
	require.Equal(t, msg.TUint32, uint32(3))
	require.Equal(t, msg.TUint64, uint64(4))
	require.Equal(t, msg.TSint32, int32(5))
	require.Equal(t, msg.TSint64, int64(6))
	require.Equal(t, msg.TSfixed32, int32(7))
	require.Equal(t, msg.TSfixed64, int64(8))
	require.Equal(t, msg.TFixed32, uint32(9))
	require.Equal(t, msg.TFixed64, uint64(10))

	require.Equal(t, msg.TFloat, float32(11.11))
	require.Equal(t, msg.TDouble, float64(12.12))
	require.Equal(t, msg.TBool, true)

	require.Nil(t, msg.TBytes1)
	require.Nil(t, msg.TBytes2)

	require.Equal(t, msg.TEnum1, godefaultstest.Enum1_April)
	require.Equal(t, msg.TEnum2, godefaultstest.Enum1_February)

	require.NotNil(t, msg.TConfig1)
	require.Equal(t, msg.TConfig1.Ip, "127.0.0.1")
	require.Equal(t, msg.TConfig1.Port, int32(80))

	require.Nil(t, msg.TConfig2)
}

func Test_GoDefaults_LiteralMessage1_2(t *testing.T) {
	msg := &godefaultstest.LiteralMessage1{
		TString1:  "ts111",
		TString2:  "ts112",
		TString3:  "ts113",
		TString4:  "ts114",
		TString5:  "ts115",
		TString6:  "ts116",
		TString7:  "ts117",
		TString8:  "ts118",
		TString9:  "ts119",
		TString10: "ts120",
		TInt32:    100,
		TInt64:    101,
		TUint32:   102,
		TUint64:   103,
		TSint32:   104,
		TSint64:   105,
		TSfixed32: 106,
		TSfixed64: 107,
		TFixed32:  108,
		TFixed64:  109,
		TFloat:    110,
		TDouble:   111,
		TBool:     true,
		TBytes1:   []byte(`hello b1`),
		TBytes2:   []byte(`hello b2`),
		TEnum1:    godefaultstest.Enum1_March,
		TEnum2:    godefaultstest.Enum1_June,
		TConfig1: &godefaultstest.Config{
			Ip:   "192.168.1.1",
			Port: 8080,
		},
		TConfig2: &godefaultstest.Config{
			Ip:   "192.168.1.2",
			Port: 8081,
		},
	}
	require.NotPanics(t, msg.SetDefaults)

	require.Equal(t, msg.TString1, `ts111`)
	require.Equal(t, msg.TString2, `ts112`)
	require.Equal(t, msg.TString3, `ts113`)
	require.Equal(t, msg.TString4, `ts114`)
	require.Equal(t, msg.TString5, `ts115`)
	require.Equal(t, msg.TString6, `ts116`)
	require.Equal(t, msg.TString7, `ts117`)
	require.Equal(t, msg.TString8, `ts118`)
	require.Equal(t, msg.TString9, `ts119`)
	require.Equal(t, msg.TString10, `ts120`)

	require.Equal(t, msg.TInt32, int32(100))
	require.Equal(t, msg.TInt64, int64(101))
	require.Equal(t, msg.TUint32, uint32(102))
	require.Equal(t, msg.TUint64, uint64(103))
	require.Equal(t, msg.TSint32, int32(104))
	require.Equal(t, msg.TSint64, int64(105))
	require.Equal(t, msg.TSfixed32, int32(106))
	require.Equal(t, msg.TSfixed64, int64(107))
	require.Equal(t, msg.TFixed32, uint32(108))
	require.Equal(t, msg.TFixed64, uint64(109))

	require.Equal(t, msg.TFloat, float32(110))
	require.Equal(t, msg.TDouble, float64(111))
	require.Equal(t, msg.TBool, true)

	require.Equal(t, msg.TBytes1, []byte(`hello b1`))
	require.Equal(t, msg.TBytes2, []byte(`hello b2`))

	require.Equal(t, msg.TEnum1, godefaultstest.Enum1_March)
	require.Equal(t, msg.TEnum2, godefaultstest.Enum1_June)

	require.NotNil(t, msg.TConfig1)
	require.Equal(t, msg.TConfig1.Ip, "192.168.1.1")
	require.Equal(t, msg.TConfig1.Port, int32(8080))

	require.NotNil(t, msg.TConfig2)
	require.Equal(t, msg.TConfig2.Ip, "192.168.1.2")
	require.Equal(t, msg.TConfig2.Port, int32(8081))
}

func Test_GoDefaults_OptionalMessage1_1(t *testing.T) {
	var msg *godefaultstest.OptionalMessage1
	require.NotPanics(t, msg.SetDefaults)
	require.Nil(t, msg)

	msg = &godefaultstest.OptionalMessage1{}
	require.NotPanics(t, msg.SetDefaults)

	require.Equal(t, *msg.TString1, `ts1`)
	require.Equal(t, *msg.TString2, ``)
	require.Equal(t, *msg.TString3, `""`)
	require.Equal(t, *msg.TString4, `"`)
	require.Equal(t, *msg.TString5, `"ts5"`)
	require.Equal(t, *msg.TString6, `"ts"6"`)
	require.Equal(t, *msg.TString7, `"ts"7"`)
	require.Equal(t, *msg.TString8, `[ts8]`)
	require.Equal(t, *msg.TString9, `{ts9}`)
	require.Nil(t, msg.TString10)

	require.Equal(t, *msg.TInt32, int32(0))
	require.Equal(t, *msg.TInt64, int64(2))
	require.Equal(t, *msg.TUint32, uint32(3))
	require.Equal(t, *msg.TUint64, uint64(4))
	require.Equal(t, *msg.TSint32, int32(5))
	require.Equal(t, *msg.TSint64, int64(6))
	require.Equal(t, *msg.TSfixed32, int32(7))
	require.Equal(t, *msg.TSfixed64, int64(8))
	require.Equal(t, *msg.TFixed32, uint32(9))
	require.Equal(t, *msg.TFixed64, uint64(10))

	require.Equal(t, *msg.TFloat, float32(11.11))
	require.Equal(t, *msg.TDouble, float64(12.12))
	require.Equal(t, *msg.TBool, true)

	require.Nil(t, msg.TBytes1)
	require.Nil(t, msg.TBytes2)

	require.Equal(t, *msg.TEnum1, godefaultstest.Enum1_February)
	require.Equal(t, *msg.TEnum2, godefaultstest.Enum1_May)

	require.NotNil(t, msg.TConfig1)
	require.Equal(t, msg.TConfig1.Ip, "127.0.0.1")
	require.Equal(t, msg.TConfig1.Port, int32(80))

	require.Nil(t, msg.TConfig2)
}

func Test_GoDefaults_OptionalMessage1_2(t *testing.T) {
	s := "ok"
	i32 := int32(10000)
	i64 := int64(10001)
	u32 := uint32(10002)
	u64 := uint64(10003)
	f32 := float32(100.11)
	f64 := float64(100.12)
	tb := true
	en := godefaultstest.Enum1_March

	msg := &godefaultstest.OptionalMessage1{
		TString1:  &s,
		TString2:  &s,
		TString3:  &s,
		TString4:  &s,
		TString5:  &s,
		TString6:  &s,
		TString7:  &s,
		TString8:  &s,
		TString9:  &s,
		TString10: &s,
		TInt32:    &i32,
		TInt64:    &i64,
		TUint32:   &u32,
		TUint64:   &u64,
		TSint32:   &i32,
		TSint64:   &i64,
		TSfixed32: &i32,
		TSfixed64: &i64,
		TFixed32:  &u32,
		TFixed64:  &u64,
		TFloat:    &f32,
		TDouble:   &f64,
		TBool:     &tb,
		TBytes1:   []byte(`HB1`),
		TBytes2:   []byte(`HB2`),
		TEnum1:    &en,
		TEnum2:    &en,
		TConfig1: &godefaultstest.Config{
			Ip:   "192.168.1.1",
			Port: 8080,
		},
		TConfig2: &godefaultstest.Config{
			Ip:   "192.168.1.2",
			Port: 8081,
		},
	}
	require.NotPanics(t, msg.SetDefaults)

	require.Equal(t, *msg.TString1, s)
	require.Equal(t, *msg.TString2, s)
	require.Equal(t, *msg.TString3, s)
	require.Equal(t, *msg.TString4, s)
	require.Equal(t, *msg.TString5, s)
	require.Equal(t, *msg.TString6, s)
	require.Equal(t, *msg.TString7, s)
	require.Equal(t, *msg.TString8, s)
	require.Equal(t, *msg.TString9, s)
	require.Equal(t, *msg.TString10, s)

	require.Equal(t, *msg.TInt32, int32(10000))
	require.Equal(t, *msg.TInt64, int64(10001))
	require.Equal(t, *msg.TUint32, uint32(10002))
	require.Equal(t, *msg.TUint64, uint64(10003))
	require.Equal(t, *msg.TSint32, int32(10000))
	require.Equal(t, *msg.TSint64, int64(10001))
	require.Equal(t, *msg.TSfixed32, int32(10000))
	require.Equal(t, *msg.TSfixed64, int64(10001))
	require.Equal(t, *msg.TFixed32, uint32(10002))
	require.Equal(t, *msg.TFixed64, uint64(10003))

	require.Equal(t, *msg.TFloat, float32(100.11))
	require.Equal(t, *msg.TDouble, float64(100.12))
	require.Equal(t, *msg.TBool, true)

	require.Equal(t, msg.TBytes1, []byte(`HB1`))
	require.Equal(t, msg.TBytes2, []byte(`HB2`))

	require.Equal(t, *msg.TEnum1, godefaultstest.Enum1_March)
	require.Equal(t, *msg.TEnum2, godefaultstest.Enum1_March)

	require.NotNil(t, msg.TConfig1)
	require.Equal(t, msg.TConfig1.Ip, "192.168.1.1")
	require.Equal(t, msg.TConfig1.Port, int32(8080))

	require.NotNil(t, msg.TConfig2)
	require.Equal(t, msg.TConfig2.Ip, "192.168.1.2")
	require.Equal(t, msg.TConfig2.Port, int32(8081))
}

func Test_GoDefaults_ListMessage1_1(t *testing.T) {
	var msg *godefaultstest.ListMessage1
	require.NotPanics(t, msg.SetDefaults)
	require.Nil(t, msg)

	msg = &godefaultstest.ListMessage1{}
	require.NotPanics(t, msg.SetDefaults)

	require.Equal(t, msg.ArrayString1, []string{"s1", "s2, s4", "s3", ""})
	require.Nil(t, msg.ArrayString2)
	require.Nil(t, msg.ArrayString3)

	require.Equal(t, msg.ArrayDouble, []float64{1.1, 1.2, 1.3})
	require.Equal(t, msg.ArrayFloat, []float32{2.1, 2.2, 2.3})

	require.Equal(t, msg.ArrayInt32, []int32{10, 11, 12})
	require.Equal(t, msg.ArrayInt64, []int64{20, 21, 22})
	require.Equal(t, msg.ArrayUint32, []uint32{30, 31, 32})
	require.Equal(t, msg.ArrayUint64, []uint64{40, 41, 42})
	require.Equal(t, msg.ArraySint32, []int32{50, 51, 52})
	require.Equal(t, msg.ArraySint64, []int64{60, 61, 62})
	require.Equal(t, msg.ArraySfixed32, []int32{70, 71, 72})
	require.Equal(t, msg.ArraySfixed64, []int64{80, 81, 82})
	require.Equal(t, msg.ArrayFixed32, []uint32{90, 91, 92})
	require.Equal(t, msg.ArrayFixed64, []uint64{100, 101, 102})

	require.Equal(t, msg.ArrayBool, []bool{true, false, true})

	require.Nil(t, msg.ArrayBytes1)
	require.Nil(t, msg.ArrayBytes2)

	require.Equal(t, msg.ArrayEnum1, []godefaultstest.Enum1{godefaultstest.Enum1_January, godefaultstest.Enum1_February, godefaultstest.Enum1_March})
	require.Equal(t, msg.ArrayEnum2, []godefaultstest.Enum1{godefaultstest.Enum1_April, godefaultstest.Enum1_May, godefaultstest.Enum1_June})

	require.Nil(t, msg.ArrayConfig1)
	require.Nil(t, msg.ArrayConfig2)
}

func Test_GoDefaults_ListMessage1_2(t *testing.T) {
	msg := &godefaultstest.ListMessage1{
		ArrayString1:  []string{},
		ArrayString2:  []string{"s1"},
		ArrayString3:  nil,
		ArrayDouble:   []float64{},
		ArrayFloat:    []float32{1.111},
		ArrayInt32:    []int32{10},
		ArrayInt64:    []int64{20},
		ArrayUint32:   []uint32{30},
		ArrayUint64:   []uint64{40},
		ArraySint32:   []int32{50},
		ArraySint64:   []int64{60},
		ArraySfixed32: []int32{70},
		ArraySfixed64: []int64{80},
		ArrayFixed32:  []uint32{90},
		ArrayFixed64:  []uint64{100},
		ArrayBool:     []bool{true},
		ArrayBytes1:   [][]byte{[]byte(`b1`)},
		ArrayBytes2:   [][]byte{[]byte(`b2`)},
		ArrayEnum1:    []godefaultstest.Enum1{godefaultstest.Enum1_January},
		ArrayEnum2:    []godefaultstest.Enum1{godefaultstest.Enum1_April, godefaultstest.Enum1_May},
		ArrayConfig1:  nil,
		ArrayConfig2:  nil,
	}
	require.NotPanics(t, msg.SetDefaults)

	require.Equal(t, msg.ArrayString1, []string{})
	require.Equal(t, msg.ArrayString2, []string{"s1"})
	require.Nil(t, msg.ArrayString3)

	require.Equal(t, msg.ArrayDouble, []float64{})
	require.Equal(t, msg.ArrayFloat, []float32{1.111})

	require.Equal(t, msg.ArrayInt32, []int32{10})
	require.Equal(t, msg.ArrayInt64, []int64{20})
	require.Equal(t, msg.ArrayUint32, []uint32{30})
	require.Equal(t, msg.ArrayUint64, []uint64{40})
	require.Equal(t, msg.ArraySint32, []int32{50})
	require.Equal(t, msg.ArraySint64, []int64{60})
	require.Equal(t, msg.ArraySfixed32, []int32{70})
	require.Equal(t, msg.ArraySfixed64, []int64{80})
	require.Equal(t, msg.ArrayFixed32, []uint32{90})
	require.Equal(t, msg.ArrayFixed64, []uint64{100})

	require.Equal(t, msg.ArrayBool, []bool{true})

	require.Equal(t, msg.ArrayBytes1, [][]byte{[]byte(`b1`)})
	require.Equal(t, msg.ArrayBytes2, [][]byte{[]byte(`b2`)})

	require.Equal(t, msg.ArrayEnum1, []godefaultstest.Enum1{godefaultstest.Enum1_January})
	require.Equal(t, msg.ArrayEnum2, []godefaultstest.Enum1{godefaultstest.Enum1_April, godefaultstest.Enum1_May})

	require.Nil(t, msg.ArrayConfig1)
	require.Nil(t, msg.ArrayConfig2)
}

func Test_GoDefaults_MapMessage1_1(t *testing.T) {
	var msg *godefaultstest.MapMessage1
	require.NotPanics(t, msg.SetDefaults)
	require.Nil(t, msg)

	msg = &godefaultstest.MapMessage1{}
	require.NotPanics(t, msg.SetDefaults)

	require.Equal(t, msg.MapStringString1, map[string]string{"k11": "v11", "k12": "v12"})
	require.Equal(t, msg.MapStringString2, map[string]string{"": ""})
	require.Nil(t, msg.MapStringString3)
	require.Nil(t, msg.MapStringString4)

	require.Equal(t, msg.MapInt32Double, map[int32]float64{11: 10.2, 10: 10.1})
	require.Equal(t, msg.MapInt32Float, map[int32]float32{20: 20.1, 21: 20.2})
	require.Equal(t, msg.MapInt32Int32, map[int32]int32{30: 1, 31: 11})
	require.Equal(t, msg.MapInt32Int64, map[int32]int64{40: 2, 41: 12})
	require.Equal(t, msg.MapInt32Uint32, map[int32]uint32{50: 3, 51: 13})
	require.Equal(t, msg.MapInt32Uint64, map[int32]uint64{60: 4, 61: 14})
	require.Equal(t, msg.MapInt32Sint32, map[int32]int32{70: 5, 71: 15})
	require.Equal(t, msg.MapInt32Sint64, map[int32]int64{80: 6, 81: 16})
	require.Equal(t, msg.MapInt32Sfixed32, map[int32]int32{90: 7, 91: 17})
	require.Equal(t, msg.MapInt32Sfixed64, map[int32]int64{100: 8, 101: 18})
	require.Equal(t, msg.MapInt32Fixed32, map[int32]uint32{110: 9, 111: 19})
	require.Equal(t, msg.MapInt32Fixed64, map[int32]uint64{120: 10, 121: 20})
	require.Equal(t, msg.MapInt32Bool, map[int32]bool{131: false, 130: true})
	require.Equal(t, msg.MapInt32String, map[int32]string{141: "v2", 140: "v1"})
	require.Equal(t, msg.MapInt32Enum1, map[int32]godefaultstest.Enum1{160: 0, 161: 1})
	require.Equal(t, msg.MapInt32Enum2, map[int32]godefaultstest.Enum1{170: 3, 171: 4})

	require.Nil(t, msg.MapInt32Bytes)
	require.Nil(t, msg.MapInt32Config)

	require.Equal(t, msg.MapInt64Int32, map[int64]int32{200: 100, 201: 101})
	require.Equal(t, msg.MapUint32Int32, map[uint32]int32{210: 110, 211: 111})
	require.Equal(t, msg.MapUint64Int32, map[uint64]int32{220: 120, 221: 121})
	require.Equal(t, msg.MapSint32Int32, map[int32]int32{230: 130, 231: 131})
	require.Equal(t, msg.MapSint64Int32, map[int64]int32{240: 140, 241: 141})
	require.Equal(t, msg.MapFixed32Int32, map[uint32]int32{250: 150, 251: 151})
	require.Equal(t, msg.MapFixed64Int32, map[uint64]int32{260: 160, 261: 161})
	require.Equal(t, msg.MapSfixed32Int32, map[int32]int32{271: 171, 270: 170})
	require.Equal(t, msg.MapSfixed64Int32, map[int64]int32{281: 181, 280: 180})
	require.Equal(t, msg.MapStringInt32, map[string]int32{"k2": 1001, "k1": 1000})
}

func Test_GoDefaults_MapMessage1_2(t *testing.T) {
	msg := &godefaultstest.MapMessage1{
		MapStringString1: map[string]string{},
		MapStringString2: map[string]string{"k11": "v11"},
		MapStringString3: map[string]string{"k11": "v11", "k12": "v12"},
		MapStringString4: map[string]string{"k11": "v11", "k12": "v12"},
		MapInt32Double:   map[int32]float64{11: 10.2},
		MapInt32Float:    map[int32]float32{20: 20.1},
		MapInt32Int32:    map[int32]int32{},
		MapInt32Int64:    map[int32]int64{40: 2},
		MapInt32Uint32:   map[int32]uint32{50: 3},
		MapInt32Uint64:   map[int32]uint64{60: 4},
		MapInt32Sint32:   map[int32]int32{70: 5},
		MapInt32Sint64:   map[int32]int64{80: 6},
		MapInt32Sfixed32: map[int32]int32{90: 7},
		MapInt32Sfixed64: map[int32]int64{100: 8},
		MapInt32Fixed32:  map[int32]uint32{110: 9},
		MapInt32Fixed64:  map[int32]uint64{120: 10},
		MapInt32Bool:     map[int32]bool{131: false},
		MapInt32String:   map[int32]string{},
		MapInt32Bytes:    map[int32][]byte{},
		MapInt32Enum1:    map[int32]godefaultstest.Enum1{},
		MapInt32Enum2:    map[int32]godefaultstest.Enum1{160: 0},
		MapInt32Config:   map[int32]*godefaultstest.Config{},
		MapInt64Int32:    map[int64]int32{200: 100},
		MapUint32Int32:   map[uint32]int32{210: 110},
		MapUint64Int32:   map[uint64]int32{220: 120},
		MapSint32Int32:   map[int32]int32{230: 130},
		MapSint64Int32:   map[int64]int32{240: 140},
		MapFixed32Int32:  map[uint32]int32{250: 150},
		MapFixed64Int32:  map[uint64]int32{260: 160},
		MapSfixed32Int32: map[int32]int32{271: 171},
		MapSfixed64Int32: map[int64]int32{281: 181},
		MapStringInt32:   map[string]int32{"k2": 1001},
	}
	require.NotPanics(t, msg.SetDefaults)

	require.Equal(t, msg.MapStringString1, map[string]string{})
	require.Equal(t, msg.MapStringString2, map[string]string{"k11": "v11"})
	require.Equal(t, msg.MapStringString3, map[string]string{"k11": "v11", "k12": "v12"})
	require.Equal(t, msg.MapStringString4, map[string]string{"k11": "v11", "k12": "v12"})

	require.Equal(t, msg.MapInt32Double, map[int32]float64{11: 10.2})
	require.Equal(t, msg.MapInt32Float, map[int32]float32{20: 20.1})
	require.Equal(t, msg.MapInt32Int32, map[int32]int32{})
	require.Equal(t, msg.MapInt32Int64, map[int32]int64{40: 2})
	require.Equal(t, msg.MapInt32Uint32, map[int32]uint32{50: 3})
	require.Equal(t, msg.MapInt32Uint64, map[int32]uint64{60: 4})
	require.Equal(t, msg.MapInt32Sint32, map[int32]int32{70: 5})
	require.Equal(t, msg.MapInt32Sint64, map[int32]int64{80: 6})
	require.Equal(t, msg.MapInt32Sfixed32, map[int32]int32{90: 7})
	require.Equal(t, msg.MapInt32Sfixed64, map[int32]int64{100: 8})
	require.Equal(t, msg.MapInt32Fixed32, map[int32]uint32{110: 9})
	require.Equal(t, msg.MapInt32Fixed64, map[int32]uint64{120: 10})
	require.Equal(t, msg.MapInt32Bool, map[int32]bool{131: false})
	require.Equal(t, msg.MapInt32String, map[int32]string{})
	require.Equal(t, msg.MapInt32Enum1, map[int32]godefaultstest.Enum1{})
	require.Equal(t, msg.MapInt32Enum2, map[int32]godefaultstest.Enum1{160: 0})

	require.Equal(t, msg.MapInt32Bytes, map[int32][]byte{})
	require.Equal(t, msg.MapInt32Config, map[int32]*godefaultstest.Config{})

	require.Equal(t, msg.MapInt64Int32, map[int64]int32{200: 100})
	require.Equal(t, msg.MapUint32Int32, map[uint32]int32{210: 110})
	require.Equal(t, msg.MapUint64Int32, map[uint64]int32{220: 120})
	require.Equal(t, msg.MapSint32Int32, map[int32]int32{230: 130})
	require.Equal(t, msg.MapSint64Int32, map[int64]int32{240: 140})
	require.Equal(t, msg.MapFixed32Int32, map[uint32]int32{250: 150})
	require.Equal(t, msg.MapFixed64Int32, map[uint64]int32{260: 160})
	require.Equal(t, msg.MapSfixed32Int32, map[int32]int32{271: 171})
	require.Equal(t, msg.MapSfixed64Int32, map[int64]int32{281: 181})
	require.Equal(t, msg.MapStringInt32, map[string]int32{"k2": 1001})
}

func Test_GoDefaults_OneofMessag1_1(t *testing.T) {
	var msg *godefaultstest.OneofMessag1
	require.NotPanics(t, msg.SetDefaults)
	require.Nil(t, msg)

	msg = &godefaultstest.OneofMessag1{}
	require.NotPanics(t, msg.SetDefaults)

	require.NotNil(t, msg.OneofTyp1)
	v, ok := msg.OneofTyp1.(*godefaultstest.OneofMessag1_Oneof1Double)
	require.True(t, ok)
	require.Equal(t, v.Oneof1Double, float64(1.1))

	require.Nil(t, msg.OneofTyp2)
}

func Test_GoDefaults_OneofMessag1_2(t *testing.T) {
	msg := &godefaultstest.OneofMessag1{
		OneofTyp1: &godefaultstest.OneofMessag1_Oneof1Sint32{},
		OneofTyp2: &godefaultstest.OneofMessag1_Oneof2String1{},
	}
	msg.SetDefaults()

	v1, ok := msg.OneofTyp1.(*godefaultstest.OneofMessag1_Oneof1Sint32)
	require.True(t, ok)
	require.Equal(t, v1.Oneof1Sint32, int32(5))

	v2, ok := msg.OneofTyp2.(*godefaultstest.OneofMessag1_Oneof2String1)
	require.True(t, ok)
	require.Equal(t, v2.Oneof2String1, "ts1")
}
