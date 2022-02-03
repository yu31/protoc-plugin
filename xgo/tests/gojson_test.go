package tests

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/yu31/protoc-plugin/xgo/tests/gojsonexternal"
	"github.com/yu31/protoc-plugin/xgo/tests/gojsontest"
	"google.golang.org/protobuf/encoding/protojson"
)

var pMarshal = &protojson.MarshalOptions{
	Multiline:       true,
	Indent:          "",
	AllowPartial:    true,
	UseProtoNames:   true,
	UseEnumNumbers:  true,
	EmitUnpopulated: true,
}

type Model2Type gojsontest.Model2

type Model3Type gojsontest.Model3

var model2 = gojsontest.Model2{
	TypeDouble1:  64.11,
	TypeDouble2:  64.22,
	TypeDouble3:  64.33,
	TypeDouble4:  64.44,
	TypeDouble5:  64.55,
	TypeFloat:    32.11,
	TypeInt32:    321,
	TypeInt64:    641,
	TypeUint32:   322,
	TypeUint64:   642,
	TypeSint32:   323,
	TypeSint64:   643,
	TypeFixed32:  324,
	TypeFixed64:  644,
	TypeSfixed32: 325,
	TypeSfixed64: 646,
	TypeBool1:    true,
	TypeBool2:    true,
	TypeString1:  "TypeString1",
	TypeString2:  "TypeString2",
	TypeString3:  "TypeString3",
	TypeString4:  "TypeString4",
	TypeString5:  "TypeString5",
	TypeBytes:    []byte("TypeBytes"),
	TypeEmbedMessage: &gojsontest.Model2_EmbedMessage1{
		Age1: "201",
		Age2: "202",
		Age3: "203",
	},
	TypeStandMessage: &gojsontest.StandMessage1{
		Name1: "q21",
		Name2: "q22",
		Name3: "q23",
	},
	TypeEmbedEnum:    gojsontest.Model2_august,
	TypeStandEnum:    gojsontest.StandEnum1_February,
	TypeExternalEnum: gojsonexternal.ExternalEnum1_Friday,
	TypeExternalMessage: &gojsonexternal.ExternalMessage1{
		Ip1: "127.0.1.1",
		Ip2: "127.0.1.2",
		Ip3: "127.0.1.3",
	},
	ArrayDouble:   []float64{64.11, 64.12, 64.13, 64.15, 64.15},
	ArrayFloat:    []float32{32.11, 32.12, 32.13, 32.15, 32.15},
	ArrayInt32:    []int32{3211, 3212, 3213, 3214, 3215},
	ArrayInt64:    []int64{6411, 6412, 6413, 6414, 6415},
	ArrayUint32:   []uint32{3221, 3222, 3223, 3224, 3225},
	ArrayUint64:   []uint64{6421, 6422, 6423, 6424, 6425},
	ArraySint32:   []int32{3231, 3232, 3233, 3234, 3235},
	ArraySint64:   []int64{6431, 6432, 6433, 6434, 6435},
	ArrayFixed32:  []uint32{3241, 3242, 3243, 3244, 3245},
	ArrayFixed64:  []uint64{6441, 6442, 6443, 6444, 6445},
	ArraySfixed32: []int32{3251, 3252, 3253, 3254, 3255},
	ArraySfixed64: []int64{6451, 6452, 6453, 6454, 6455},
	ArrayBool:     []bool{true, false, true, false, true},
	ArrayString:   []string{"s_arr_1", "s_arr_2", "s_arr_3", "s_arr_4", "s_arr_5"},
	ArrayBytes: [][]byte{
		[]byte("b_arr_1"), []byte("b_arr_2"), []byte("b_arr_3"), []byte("b_arr_4"), []byte("b_arr_5"),
	},
	ArrayEmbedMessage: []*gojsontest.Model2_EmbedMessage1{
		{Age1: "311", Age2: "312", Age3: "313"},
	},
	ArrayStandMessage: []*gojsontest.StandMessage1{
		{Name1: "q11", Name2: "q12", Name3: "q13"},
	},
	ArrayExternalMessage: []*gojsonexternal.ExternalMessage1{
		{Ip1: "127.0.1.1", Ip2: "127.0.1.2", Ip3: "127.0.1.3"},
	},
	ArrayEmbedEnum: []gojsontest.Model2_EmbedEnum1{
		gojsontest.Model2_july,
		gojsontest.Model2_august,
		gojsontest.Model2_september,
		gojsontest.Model2_october,
		gojsontest.Model2_november,
		gojsontest.Model2_december,
	},
	ArrayStandEnum: []gojsontest.StandEnum1{
		gojsontest.StandEnum1_January,
		gojsontest.StandEnum1_February,
		gojsontest.StandEnum1_March,
		gojsontest.StandEnum1_April,
		gojsontest.StandEnum1_May,
		gojsontest.StandEnum1_June,
	},
	ArrayExternalEnum: []gojsonexternal.ExternalEnum1{
		gojsonexternal.ExternalEnum1_Monday,
		gojsonexternal.ExternalEnum1_Tuesday,
		gojsonexternal.ExternalEnum1_Wednesday,
		gojsonexternal.ExternalEnum1_Thursday,
		gojsonexternal.ExternalEnum1_Friday,
		gojsonexternal.ExternalEnum1_Saturday,
		gojsonexternal.ExternalEnum1_Sunday,
	},
	MapInt32Double: map[int32]float64{
		3211: 64.11,
	},
	MapInt32Float: map[int32]float32{
		3221: 32.11,
	},
	MapInt32Int32: map[int32]int32{
		3231: 31,
	},
	MapInt32Int64: map[int32]int64{
		3241: 41,
	},
	MapInt32Uint32: map[int32]uint32{
		3251: 51,
	},
	MapInt32Uint64: map[int32]uint64{
		3261: 61,
	},
	MapInt32Sint32: map[int32]int32{
		3271: 71,
	},
	MapInt32Sint64: map[int32]int64{
		3281: 81,
	},
	MapInt32Fixed32: map[int32]uint32{
		3291: 91,
	},
	MapInt32Fixed64: map[int32]uint64{
		32101: 101,
	},
	MapInt32Sfixed32: map[int32]int32{
		32111: 111,
	},
	MapInt32Sfixed64: map[int32]int64{
		32121: 121,
	},
	MapInt32Bool: map[int32]bool{
		32131: true,
	},
	MapInt32String: map[int32]string{
		32141: "mvs11",
	},
	MapInt32Bytes: map[int32][]byte{
		32151: []byte("mvb11"),
	},
	MapInt32EmbedMessage: map[int32]*gojsontest.Model2_EmbedMessage1{
		32161: {Age1: "311", Age2: "312", Age3: "313"},
	},
	MapInt32StandMessage: map[int32]*gojsontest.StandMessage1{
		32171: {Name1: "q11", Name2: "q12", Name3: "q13"},
	},
	MapInt32EmbedEnum: map[int32]gojsontest.Model2_EmbedEnum1{
		32181: gojsontest.Model2_july,
	},
	MapInt32StandEnum: map[int32]gojsontest.StandEnum1{
		32191: gojsontest.StandEnum1_January,
	},
	MapInt64Int32: map[int64]int32{
		6411: 11,
	},
	MapUint32Int32: map[uint32]int32{
		3221: 21,
	},
	MapUint64Int32: map[uint64]int32{
		6431: 31,
	},
	MapSint32Int32: map[int32]int32{
		3241: 41,
	},
	MapSint64Int32: map[int64]int32{
		6451: 51,
	},
	MapFixed32Int32: map[uint32]int32{
		3261: 61,
	},
	MapFixed64Int32: map[uint64]int32{
		6471: 71,
	},
	MapSfixed32Int32: map[int32]int32{
		3281: 81,
	},
	MapSfixed64Int32: map[int64]int32{
		6491: 91,
	},
	MapStringInt32: map[string]int32{
		"mks11": 100,
	},
	MapStringString: map[string]string{
		"mks21": "mvs21",
	},
	MapStringEmbedMessage: map[string]*gojsontest.Model2_EmbedMessage1{
		"mks31": {Age1: "411", Age2: "412", Age3: "413"},
	},
	MapStringStandMessage: map[string]*gojsontest.StandMessage1{
		"mks41": {Name1: "q111", Name2: "q112", Name3: "q113"},
	},
	MapStringExternalMessage: map[string]*gojsonexternal.ExternalMessage1{
		"mks51": {Ip1: "127.0.1.1", Ip2: "127.0.1.2", Ip3: "127.0.1.3"},
	},
	MapStringEmbedEnum: map[string]gojsontest.Model2_EmbedEnum1{
		"mks61": gojsontest.Model2_july,
	},
	MapStringStandEnum: map[string]gojsontest.StandEnum1{
		"mks71": gojsontest.StandEnum1_January,
	},
	MapStringExternalEnum: map[string]gojsonexternal.ExternalEnum1{
		"mks81": gojsonexternal.ExternalEnum1_Monday,
	},
}

var model2Type = Model2Type{
	TypeDouble1:  64.11,
	TypeDouble2:  64.22,
	TypeDouble3:  64.33,
	TypeDouble4:  64.44,
	TypeDouble5:  64.55,
	TypeFloat:    32.11,
	TypeInt32:    321,
	TypeInt64:    641,
	TypeUint32:   322,
	TypeUint64:   642,
	TypeSint32:   323,
	TypeSint64:   643,
	TypeFixed32:  324,
	TypeFixed64:  644,
	TypeSfixed32: 325,
	TypeSfixed64: 646,
	TypeBool1:    true,
	TypeBool2:    true,
	TypeString1:  "TypeString1",
	TypeString2:  "TypeString2",
	TypeString3:  "TypeString3",
	TypeString4:  "TypeString4",
	TypeString5:  "TypeString5",
	TypeBytes:    []byte("TypeBytes"),
	TypeEmbedMessage: &gojsontest.Model2_EmbedMessage1{
		Age1: "201",
		Age2: "202",
		Age3: "203",
	},
	TypeStandMessage: &gojsontest.StandMessage1{
		Name1: "q21",
		Name2: "q22",
		Name3: "q23",
	},
	TypeEmbedEnum:    gojsontest.Model2_august,
	TypeStandEnum:    gojsontest.StandEnum1_February,
	TypeExternalEnum: gojsonexternal.ExternalEnum1_Friday,
	TypeExternalMessage: &gojsonexternal.ExternalMessage1{
		Ip1: "127.0.1.1",
		Ip2: "127.0.1.2",
		Ip3: "127.0.1.3",
	},
	ArrayDouble:   []float64{64.11, 64.12, 64.13, 64.15, 64.15},
	ArrayFloat:    []float32{32.11, 32.12, 32.13, 32.15, 32.15},
	ArrayInt32:    []int32{3211, 3212, 3213, 3214, 3215},
	ArrayInt64:    []int64{6411, 6412, 6413, 6414, 6415},
	ArrayUint32:   []uint32{3221, 3222, 3223, 3224, 3225},
	ArrayUint64:   []uint64{6421, 6422, 6423, 6424, 6425},
	ArraySint32:   []int32{3231, 3232, 3233, 3234, 3235},
	ArraySint64:   []int64{6431, 6432, 6433, 6434, 6435},
	ArrayFixed32:  []uint32{3241, 3242, 3243, 3244, 3245},
	ArrayFixed64:  []uint64{6441, 6442, 6443, 6444, 6445},
	ArraySfixed32: []int32{3251, 3252, 3253, 3254, 3255},
	ArraySfixed64: []int64{6451, 6452, 6453, 6454, 6455},
	ArrayBool:     []bool{true, false, true, false, true},
	ArrayString:   []string{"s_arr_1", "s_arr_2", "s_arr_3", "s_arr_4", "s_arr_5"},
	ArrayBytes: [][]byte{
		[]byte("b_arr_1"), []byte("b_arr_2"), []byte("b_arr_3"), []byte("b_arr_4"), []byte("b_arr_5"),
	},
	ArrayEmbedMessage: []*gojsontest.Model2_EmbedMessage1{
		{Age1: "311", Age2: "312", Age3: "313"},
	},
	ArrayStandMessage: []*gojsontest.StandMessage1{
		{Name1: "q11", Name2: "q12", Name3: "q13"},
	},
	ArrayExternalMessage: []*gojsonexternal.ExternalMessage1{
		{Ip1: "127.0.1.1", Ip2: "127.0.1.2", Ip3: "127.0.1.3"},
	},
	ArrayEmbedEnum: []gojsontest.Model2_EmbedEnum1{
		gojsontest.Model2_july,
		gojsontest.Model2_august,
		gojsontest.Model2_september,
		gojsontest.Model2_october,
		gojsontest.Model2_november,
		gojsontest.Model2_december,
	},
	ArrayStandEnum: []gojsontest.StandEnum1{
		gojsontest.StandEnum1_January,
		gojsontest.StandEnum1_February,
		gojsontest.StandEnum1_March,
		gojsontest.StandEnum1_April,
		gojsontest.StandEnum1_May,
		gojsontest.StandEnum1_June,
	},
	ArrayExternalEnum: []gojsonexternal.ExternalEnum1{
		gojsonexternal.ExternalEnum1_Monday,
		gojsonexternal.ExternalEnum1_Tuesday,
		gojsonexternal.ExternalEnum1_Wednesday,
		gojsonexternal.ExternalEnum1_Thursday,
		gojsonexternal.ExternalEnum1_Friday,
		gojsonexternal.ExternalEnum1_Saturday,
		gojsonexternal.ExternalEnum1_Sunday,
	},
	MapInt32Double: map[int32]float64{
		3211: 64.11,
	},
	MapInt32Float: map[int32]float32{
		3221: 32.11,
	},
	MapInt32Int32: map[int32]int32{
		3231: 31,
	},
	MapInt32Int64: map[int32]int64{
		3241: 41,
	},
	MapInt32Uint32: map[int32]uint32{
		3251: 51,
	},
	MapInt32Uint64: map[int32]uint64{
		3261: 61,
	},
	MapInt32Sint32: map[int32]int32{
		3271: 71,
	},
	MapInt32Sint64: map[int32]int64{
		3281: 81,
	},
	MapInt32Fixed32: map[int32]uint32{
		3291: 91,
	},
	MapInt32Fixed64: map[int32]uint64{
		32101: 101,
	},
	MapInt32Sfixed32: map[int32]int32{
		32111: 111,
	},
	MapInt32Sfixed64: map[int32]int64{
		32121: 121,
	},
	MapInt32Bool: map[int32]bool{
		32131: true,
	},
	MapInt32String: map[int32]string{
		32141: "mvs11",
	},
	MapInt32Bytes: map[int32][]byte{
		32151: []byte("mvb11"),
	},
	MapInt32EmbedMessage: map[int32]*gojsontest.Model2_EmbedMessage1{
		32161: {Age1: "311", Age2: "312", Age3: "313"},
	},
	MapInt32StandMessage: map[int32]*gojsontest.StandMessage1{
		32171: {Name1: "q11", Name2: "q12", Name3: "q13"},
	},
	MapInt32EmbedEnum: map[int32]gojsontest.Model2_EmbedEnum1{
		32181: gojsontest.Model2_july,
	},
	MapInt32StandEnum: map[int32]gojsontest.StandEnum1{
		32191: gojsontest.StandEnum1_January,
	},
	MapInt64Int32: map[int64]int32{
		6411: 11,
	},
	MapUint32Int32: map[uint32]int32{
		3221: 21,
	},
	MapUint64Int32: map[uint64]int32{
		6431: 31,
	},
	MapSint32Int32: map[int32]int32{
		3241: 41,
	},
	MapSint64Int32: map[int64]int32{
		6451: 51,
	},
	MapFixed32Int32: map[uint32]int32{
		3261: 61,
	},
	MapFixed64Int32: map[uint64]int32{
		6471: 71,
	},
	MapSfixed32Int32: map[int32]int32{
		3281: 81,
	},
	MapSfixed64Int32: map[int64]int32{
		6491: 91,
	},
	MapStringInt32: map[string]int32{
		"mks11": 100,
	},
	MapStringString: map[string]string{
		"mks21": "mvs21",
	},
	MapStringEmbedMessage: map[string]*gojsontest.Model2_EmbedMessage1{
		"mks31": {Age1: "411", Age2: "412", Age3: "413"},
	},
	MapStringStandMessage: map[string]*gojsontest.StandMessage1{
		"mks41": {Name1: "q111", Name2: "q112", Name3: "q113"},
	},
	MapStringExternalMessage: map[string]*gojsonexternal.ExternalMessage1{
		"mks51": {Ip1: "127.0.1.1", Ip2: "127.0.1.2", Ip3: "127.0.1.3"},
	},
	MapStringEmbedEnum: map[string]gojsontest.Model2_EmbedEnum1{
		"mks61": gojsontest.Model2_july,
	},
	MapStringStandEnum: map[string]gojsontest.StandEnum1{
		"mks71": gojsontest.StandEnum1_January,
	},
	MapStringExternalEnum: map[string]gojsonexternal.ExternalEnum1{
		"mks81": gojsonexternal.ExternalEnum1_Monday,
	},
}

var model3 = gojsontest.Model3{
	TString1:  "Hello 1",
	TString2:  "Hello 2",
	TString3:  "Hello 3",
	TString4:  "Hello 4",
	TString5:  "Hello 5",
	TString6:  "Hello 6",
	TString7:  "Hello 7",
	TString8:  "Hello 8",
	TString9:  "Hello 9",
	TString10: "Hello 10",
	TInt32:    1,
	TInt64:    2,
	TUint32:   3,
	TUint64:   4,
	TSint32:   5,
	TSint64:   6,
	TSfixed32: 7,
	TSfixed64: 8,
	TFixed32:  9,
	TFixed64:  10,
	TFloat:    11,
	TDouble:   12,
	TBool:     true,
}

func Test_GoJSON_Serialize1(t *testing.T) {
	m1 := &gojsontest.Model1{
		OneofType1:   &gojsontest.Model1_Oneof1Double{Oneof1Double: 1.1},
		OneofType2:   &gojsontest.Model1_Oneof2Float{Oneof2Float: 2.1},
		OneofType3:   &gojsontest.Model1_Oneof3Int32{Oneof3Int32: 3},
		Oneof_Type4:  &gojsontest.Model1_Oneof4Int64{Oneof4Int64: 4},
		Oneof_Type5:  &gojsontest.Model1_Oneof5Uint32{Oneof5Uint32: 5},
		OneofType6:   &gojsontest.Model1_Oneof6Uint64{Oneof6Uint64: 6},
		OneofType7:   &gojsontest.Model1_Oneof7Sint32{Oneof7Sint32: 7},
		Oneof_Type8:  &gojsontest.Model1_Oneof8Sint64{Oneof8Sint64: 8},
		Oneof_Type9:  &gojsontest.Model1_Oneof9Fixed32{Oneof9Fixed32: 9},
		Oneof_Type10: &gojsontest.Model1_Oneof10Fixed64{Oneof10Fixed64: 10},
		Oneof_Type11: &gojsontest.Model1_Oneof11Sfixed32{Oneof11Sfixed32: 11},
		Oneof_Type12: &gojsontest.Model1_Oneof12Sfixed64{Oneof12Sfixed64: 12},
		Oneof_Type13: &gojsontest.Model1_Oneof13Bool{Oneof13Bool: true},
		Oneof_Type14: &gojsontest.Model1_Oneof14String{Oneof14String: "Model1_Oneof14String"},
		Oneof_Type15: &gojsontest.Model1_Oneof15Bytes{Oneof15Bytes: []byte("Model1_Oneof15Bytes")},
		Oneof_Type16: &gojsontest.Model1_Oneof16EmbedMessage{Oneof16EmbedMessage: &gojsontest.Model1_EmbedMessage1{
			Age1: "101",
			Age2: "102",
			Age3: "103",
		}},
		Oneof_Type17: &gojsontest.Model1_Oneof17StandMessage{Oneof17StandMessage: &gojsontest.StandMessage1{
			Name1: "q11",
			Name2: "q12",
			Name3: "q13",
		}},
		Oneof_Type18: &gojsontest.Model1_Oneof18ExternalMessage{Oneof18ExternalMessage: &gojsonexternal.ExternalMessage1{
			Ip1: "127.0.0.1",
			Ip2: "127.0.0.2",
			Ip3: "127.0.0.3",
		}},
		Oneof_Type19:     &gojsontest.Model1_Oneof19EmbedEnum{Oneof19EmbedEnum: gojsontest.Model1_august},
		Oneof_Type20:     &gojsontest.Model1_Oneof20StandEnum{Oneof20StandEnum: gojsontest.StandEnum1_January},
		Oneof_Type21:     &gojsontest.Model1_Oneof21ExternalEnum{Oneof21ExternalEnum: gojsonexternal.ExternalEnum1_Friday},
		Oneof_Type22Null: &gojsontest.Model1_Oneof22ExternalMessage{Oneof22ExternalMessage: nil},
		Oneof_Type23Null: nil,
		TypeDouble1:      64.11,
		TypeDouble2:      64.22,
		TypeDouble3:      64.33,
		TypeDouble4:      64.44,
		Type_Double5:     64.55,
		TypeFloat:        32.11,
		TypeInt32:        321,
		TypeInt64:        641,
		TypeUint32:       322,
		TypeUint64:       642,
		TypeSint32:       323,
		TypeSint64:       643,
		TypeFixed32:      324,
		TypeFixed64:      644,
		TypeSfixed32:     325,
		TypeSfixed64:     646,
		TypeBool1:        true,
		TypeBool2:        false,
		TypeString1:      "TypeString1",
		TypeString2:      "TypeString2",
		TypeString3:      "TypeString3",
		TypeString4:      "TypeString4",
		TypeString5:      "TypeString5",
		TypeBytes:        []byte("TypeBytes"),
		TypeEmbedMessage: &gojsontest.Model1_EmbedMessage1{
			Age1: "201",
			Age2: "202",
			Age3: "203",
		},
		TypeStandMessage: &gojsontest.StandMessage1{
			Name1: "q21",
			Name2: "q22",
			Name3: "q23",
		},
		TypeEmbedEnum:    gojsontest.Model1_july,
		TypeStandEnum:    gojsontest.StandEnum1_January,
		TypeExternalEnum: gojsonexternal.ExternalEnum1_Friday,
		TypeExternalMessage: &gojsonexternal.ExternalMessage1{
			Ip1: "127.0.1.1",
			Ip2: "127.0.1.2",
			Ip3: "127.0.1.3",
		},
		TypeBytesNull:           nil,
		TypeEmbedMessageNull:    nil,
		TypeStandMessageNull:    nil,
		TypeExternalMessageNull: nil,
		ArrayDouble:             []float64{64.11, 64.12, 64.13, 64.15, 64.15},
		ArrayFloat:              []float32{32.11, 32.12, 32.13, 32.15, 32.15},
		ArrayInt32:              []int32{3211, 3212, 3213, 3214, 3215},
		ArrayInt64:              []int64{6411, 6412, 6413, 6414, 6415},
		ArrayUint32:             []uint32{3221, 3222, 3223, 3224, 3225},
		ArrayUint64:             []uint64{6421, 6422, 6423, 6424, 6425},
		ArraySint32:             []int32{3231, 3232, 3233, 3234, 3235},
		ArraySint64:             []int64{6431, 6432, 6433, 6434, 6435},
		ArrayFixed32:            []uint32{3241, 3242, 3243, 3244, 3245},
		ArrayFixed64:            []uint64{6441, 6442, 6443, 6444, 6445},
		ArraySfixed32:           []int32{3251, 3252, 3253, 3254, 3255},
		ArraySfixed64:           []int64{6451, 6452, 6453, 6454, 6455},
		ArrayBool:               []bool{true, false, true, false, true},
		ArrayString:             []string{"s_arr_1", "s_arr_2", "s_arr_3", "s_arr_4", "s_arr_5"},
		ArrayBytes:              [][]byte{[]byte("b_arr_1"), []byte("b_arr_2"), []byte("b_arr_3"), []byte("b_arr_4"), []byte("b_arr_5")},
		ArrayEmbedMessage: []*gojsontest.Model1_EmbedMessage1{
			{Age1: "311", Age2: "312", Age3: "313"},
		},
		ArrayStandMessage: []*gojsontest.StandMessage1{
			{Name1: "q11", Name2: "q12", Name3: "q13"},
		},
		ArrayExternalMessage: []*gojsonexternal.ExternalMessage1{
			{Ip1: "127.0.1.1", Ip2: "127.0.1.2", Ip3: "127.0.1.3"},
		},
		ArrayEmbedEnum: []gojsontest.Model1_EmbedEnum1{
			gojsontest.Model1_july,
			gojsontest.Model1_august,
			gojsontest.Model1_september,
			gojsontest.Model1_october,
			gojsontest.Model1_november,
			gojsontest.Model1_december,
		},
		ArrayStandEnum: []gojsontest.StandEnum1{
			gojsontest.StandEnum1_January,
			gojsontest.StandEnum1_February,
			gojsontest.StandEnum1_March,
			gojsontest.StandEnum1_April,
			gojsontest.StandEnum1_May,
			gojsontest.StandEnum1_June,
		},
		ArrayExternalEnum: []gojsonexternal.ExternalEnum1{
			gojsonexternal.ExternalEnum1_Monday,
			gojsonexternal.ExternalEnum1_Tuesday,
			gojsonexternal.ExternalEnum1_Wednesday,
			gojsonexternal.ExternalEnum1_Thursday,
			gojsonexternal.ExternalEnum1_Friday,
			gojsonexternal.ExternalEnum1_Saturday,
			gojsonexternal.ExternalEnum1_Sunday,
		},
		ArrayStandEnumNull: nil,
		MapInt32Double:     map[int32]float64{3211: 64.11, 3212: 64.12, 3213: 64.13, 3214: 64.14, 3215: 64.15},
		MapInt32Float:      map[int32]float32{3221: 32.11, 3222: 32.12, 3223: 32.13, 3224: 32.14, 3225: 32.15},
		MapInt32Int32:      map[int32]int32{3231: 31, 3242: 32, 3233: 33, 3234: 34, 3235: 35},
		MapInt32Int64:      map[int32]int64{3241: 41, 3242: 42, 3243: 43, 3244: 44, 3245: 45},
		MapInt32Uint32:     map[int32]uint32{3251: 51, 3252: 52, 3253: 53, 3254: 54, 3255: 55},
		MapInt32Uint64:     map[int32]uint64{3261: 61, 3262: 62, 3263: 63, 3264: 64, 3265: 65},
		MapInt32Sint32:     map[int32]int32{3271: 71, 3272: 72, 3273: 73, 3274: 74, 3275: 75},
		MapInt32Sint64:     map[int32]int64{3281: 81, 3282: 82, 3283: 83, 3284: 84, 3285: 85},
		MapInt32Fixed32:    map[int32]uint32{3291: 91, 3292: 92, 3293: 93, 3294: 94, 3295: 95},
		MapInt32Fixed64:    map[int32]uint64{32101: 101, 32102: 102, 32103: 103, 32104: 104, 32105: 105},
		MapInt32Sfixed32:   map[int32]int32{32111: 111, 32112: 112, 32113: 113, 32114: 114, 32115: 115},
		MapInt32Sfixed64:   map[int32]int64{32121: 121, 32122: 122, 32123: 123, 32124: 124, 32125: 125},
		MapInt32Bool:       map[int32]bool{32131: true, 32132: false, 32133: true, 32134: false, 32135: true},
		MapInt32String:     map[int32]string{32141: "mvs11", 32142: "mvs12", 32143: "mvs13", 32144: "mvs14", 32145: "mvs15"},
		MapInt32Bytes:      map[int32][]byte{32151: []byte("mvb11"), 32152: []byte("mvb12"), 32153: []byte("mvb13"), 32154: []byte("mvb14"), 32155: []byte("mvb15")},
		MapInt32EmbedMessage: map[int32]*gojsontest.Model1_EmbedMessage1{
			32161: {Age1: "311", Age2: "312", Age3: "313"},
			32162: {Age1: "321", Age2: "322", Age3: "323"},
			32163: {Age1: "331", Age2: "332", Age3: "323"},
			32164: {Age1: "341", Age2: "342", Age3: "323"},
			32165: {Age1: "351", Age2: "352", Age3: "323"},
		},
		MapInt32StandMessage: map[int32]*gojsontest.StandMessage1{
			32171: {Name1: "q11", Name2: "q12", Name3: "q13"},
			32172: {Name1: "q21", Name2: "q22", Name3: "q23"},
			32173: {Name1: "q31", Name2: "q32", Name3: "q33"},
			32174: {Name1: "q41", Name2: "q42", Name3: "q43"},
			32175: {Name1: "q51", Name2: "q52", Name3: "q53"},
		},
		MapInt32EmbedEnum: map[int32]gojsontest.Model1_EmbedEnum1{
			32181: gojsontest.Model1_july,
			32182: gojsontest.Model1_august,
			32183: gojsontest.Model1_september,
			32184: gojsontest.Model1_october,
			32185: gojsontest.Model1_november,
			32186: gojsontest.Model1_december,
		},
		MapInt32StandEnum: map[int32]gojsontest.StandEnum1{
			32191: gojsontest.StandEnum1_January,
			32192: gojsontest.StandEnum1_February,
			32193: gojsontest.StandEnum1_March,
			32194: gojsontest.StandEnum1_April,
			32195: gojsontest.StandEnum1_May,
			32196: gojsontest.StandEnum1_June,
		},
		MapInt64Int32:      map[int64]int32{6411: 11, 6412: 12, 6413: 13, 6414: 14, 6415: 15},
		MapUint32Int32:     map[uint32]int32{3221: 21, 3222: 22, 3223: 23, 3224: 24, 3225: 25},
		MapUint64Int32:     map[uint64]int32{6431: 31, 6432: 32, 6433: 33, 6434: 34, 6435: 35},
		MapSint32Int32:     map[int32]int32{3241: 41, 3242: 42, 3243: 43, 3244: 44, 3245: 45},
		MapSint64Int32:     map[int64]int32{6451: 51, 6452: 52, 6453: 53, 6454: 54, 6455: 55},
		MapFixed32Int32:    map[uint32]int32{3261: 61, 3262: 62, 3263: 63, 3264: 64, 3265: 65},
		MapFixed64Int32:    map[uint64]int32{6471: 71, 6472: 72, 6473: 73, 6474: 74, 6475: 75},
		MapSfixed32Int32:   map[int32]int32{3281: 81, 3282: 82, 3283: 83, 3284: 84, 3285: 85},
		MapSfixed64Int32:   map[int64]int32{6491: 91, 6492: 92, 6493: 93, 6494: 94, 6495: 95},
		MapStringInt32:     map[string]int32{"mks11": 100, "mks12": 101, "mks13": 102, "mks14": 103, "mks15": 104},
		MapStringInt32Null: nil,
		MapStringString:    map[string]string{"mks21": "mvs21", "mks22": "mvs22", "mks23": "mvs23", "mks24": "mvs24", "mks25": "mvs25"},
		MapStringEmbedMessage: map[string]*gojsontest.Model1_EmbedMessage1{
			"mks31": {Age1: "411", Age2: "412", Age3: "413"},
			"mks32": {Age1: "421", Age2: "422", Age3: "433"},
			"mks33": {Age1: "431", Age2: "432", Age3: "433"},
			"mks34": {Age1: "441", Age2: "442", Age3: "433"},
			"mks35": {Age1: "451", Age2: "452", Age3: "433"},
		},
		MapStringStandMessage: map[string]*gojsontest.StandMessage1{
			"mks41": {Name1: "q111", Name2: "q112", Name3: "q113"},
			"mks42": {Name1: "q121", Name2: "q122", Name3: "q123"},
			"mks43": {Name1: "q131", Name2: "q132", Name3: "q133"},
			"mks44": {Name1: "q141", Name2: "q142", Name3: "q143"},
			"mks45": {Name1: "q151", Name2: "q152", Name3: "q153"},
			"mks46": {Name1: "q161", Name2: "q162", Name3: "q163"},
		},
		MapStringExternalMessage: map[string]*gojsonexternal.ExternalMessage1{
			"mks51": {Ip1: "127.0.1.1", Ip2: "127.0.1.2", Ip3: "127.0.1.3"},
			"mks52": {Ip1: "127.0.2.1", Ip2: "127.0.2.2", Ip3: "127.0.2.3"},
			"mks53": {Ip1: "127.0.3.1", Ip2: "127.0.3.2", Ip3: "127.0.3.3"},
			"mks54": {Ip1: "127.0.4.1", Ip2: "127.0.4.2", Ip3: "127.0.4.3"},
			"mks55": {Ip1: "127.0.5.1", Ip2: "127.0.5.2", Ip3: "127.0.5.3"},
			"mks56": {Ip1: "127.0.6.1", Ip2: "127.0.6.2", Ip3: "127.0.6.3"},
			"mks57": {Ip1: "127.0.7.1", Ip2: "127.0.7.2", Ip3: "127.0.7.3"},
		},
		MapStringEmbedEnum: map[string]gojsontest.Model1_EmbedEnum1{
			"mks61": gojsontest.Model1_july,
			"mks62": gojsontest.Model1_august,
			"mks63": gojsontest.Model1_september,
			"mks64": gojsontest.Model1_october,
			"mks65": gojsontest.Model1_november,
			"mks66": gojsontest.Model1_december,
		},
		MapStringStandEnum: map[string]gojsontest.StandEnum1{
			"mks71": gojsontest.StandEnum1_January,
			"mks72": gojsontest.StandEnum1_February,
			"mks73": gojsontest.StandEnum1_March,
			"mks74": gojsontest.StandEnum1_April,
			"mks75": gojsontest.StandEnum1_May,
			"mks76": gojsontest.StandEnum1_June,
		},
		MapStringExternalEnum: map[string]gojsonexternal.ExternalEnum1{
			"mks81": gojsonexternal.ExternalEnum1_Monday,
			"mks82": gojsonexternal.ExternalEnum1_Tuesday,
			"mks83": gojsonexternal.ExternalEnum1_Wednesday,
			"mks84": gojsonexternal.ExternalEnum1_Thursday,
			"mks85": gojsonexternal.ExternalEnum1_Friday,
			"mks86": gojsonexternal.ExternalEnum1_Saturday,
			"mks87": gojsonexternal.ExternalEnum1_Sunday,
		},
	}

	b1, err := m1.MarshalJSON()
	require.Nil(t, err)
	_ = b1

	//fmt.Println(string(b1))

	m2 := new(gojsontest.Model1)
	err = m2.UnmarshalJSON(b1)
	require.Nil(t, err)

	bb, _ := pMarshal.Marshal(m2)
	_ = bb

	fmt.Println("---------------------------------")
	//fmt.Println(string(bb))
}

func Test_GoJSON_Serialize2(t *testing.T) {
	var m11 interface{}
	var m12 interface{}

	m11 = &model2
	m12 = &model2Type

	_, ok := m11.(json.Marshaler)
	require.True(t, ok)
	_, ok = m11.(json.Unmarshaler)
	require.True(t, ok)

	_, ok = m12.(json.Marshaler)
	require.False(t, ok)
	_, ok = m12.(json.Unmarshaler)
	require.False(t, ok)

	//b1, err := json.Marshal(m1)
	b1, err := m11.(json.Marshaler).MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	b2, err := json.Marshal(m12)
	require.Nil(t, err)

	require.Equal(t, b1, b2)

	m21 := &gojsontest.Model2{}
	m22 := &Model2Type{}

	err = m21.UnmarshalJSON(b2)
	require.Nil(t, err)
	err = json.Unmarshal(b1, m22)
	require.Nil(t, err)
	require.Equal(t, m11, m21)
	require.Equal(t, m12, m22)
}

func Test_GoJSON_Serialize3(t *testing.T) {
	b, err := model3.MarshalJSON()
	require.Nil(t, err)
	_ = b
	//fmt.Println(string(b))
}

func Test_GoJSON_NameStyleTextName(t *testing.T) {
	data1 := &gojsontest.NameStyleTextName{
		NameStyle1:   1,
		Names_Style2: 2,
		Name_Style3:  3,
		NameStyle4:   4,
		Namestyle5:   5,
		NameStyle6:   6,
		NameStyle7:   7,
		Namestyle8:   8,
		DataType1:    &gojsontest.NameStyleTextName_Float1{Float1: "float32"},
		Data_Type2:   &gojsontest.NameStyleTextName_Float2{Float2: "float64"},
		Data_Type3:   &gojsontest.NameStyleTextName_Integer3{Integer3: "int32"},
		DataType4:    &gojsontest.NameStyleTextName_Integer4{Integer4: "int64"},
		Datatype5:    &gojsontest.NameStyleTextName_Integer5{Integer5: "uint32"},
		DataType6:    &gojsontest.NameStyleTextName_Integer6{Integer6: "uint64"},
		DataType7:    &gojsontest.NameStyleTextName_Integer7{Integer7: "int"},
		Datatype8:    &gojsontest.NameStyleTextName_Integer8{Integer8: "uint32"},
	}

	expected := []byte(`{"name_style1":1,"names_Style2":2,"Name_Style3":3,"Name_style4":4,"namestyle5":5,"nameStyle6":6,"NameStyle7":7,"Namestyle8":8,"data_type1":{"float1":"float32"},"data_Type2":{"float2":"float64"},"Data_Type3":{"integer3":"int32"},"Data_type4":{"integer4":"int64"},"datatype5":{"integer5":"uint32"},"dataType6":{"integer6":"uint64"},"DataType7":{"integer7":"int"},"Datatype8":{"integer8":"uint32"}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.NameStyleTextName{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)
}

func Test_GoJSON_NameStyleGoName(t *testing.T) {
	data1 := &gojsontest.NameStyleGoName{
		NameStyle1:   1,
		Names_Style2: 2,
		Name_Style3:  3,
		NameStyle4:   4,
		Namestyle5:   5,
		NameStyle6:   6,
		NameStyle7:   7,
		Namestyle8:   8,
		DataType1:    &gojsontest.NameStyleGoName_Float1{Float1: "float32"},
		Data_Type2:   &gojsontest.NameStyleGoName_Float2{Float2: "float64"},
		Data_Type3:   &gojsontest.NameStyleGoName_Integer3{Integer3: "int32"},
		DataType4:    &gojsontest.NameStyleGoName_Integer4{Integer4: "int64"},
		Datatype5:    &gojsontest.NameStyleGoName_Integer5{Integer5: "uint32"},
		DataType6:    &gojsontest.NameStyleGoName_Integer6{Integer6: "uint64"},
		DataType7:    &gojsontest.NameStyleGoName_Integer7{Integer7: "int"},
		Datatype8:    &gojsontest.NameStyleGoName_Integer8{Integer8: "uint32"},
	}

	expected := []byte(`{"NameStyle1":1,"Names_Style2":2,"Name_Style3":3,"NameStyle4":4,"Namestyle5":5,"NameStyle6":6,"NameStyle7":7,"Namestyle8":8,"DataType1":{"Float1":"float32"},"Data_Type2":{"Float2":"float64"},"Data_Type3":{"Integer3":"int32"},"DataType4":{"Integer4":"int64"},"Datatype5":{"Integer5":"uint32"},"DataType6":{"Integer6":"uint64"},"DataType7":{"Integer7":"int"},"Datatype8":{"Integer8":"uint32"}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.NameStyleGoName{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)
}

func Test_GoJSON_NameStyleJSONName(t *testing.T) {
	data1 := &gojsontest.NameStyleJSONName{
		NameStyle1:   1,
		Names_Style2: 2,
		Name_Style3:  3,
		NameStyle4:   4,
		Namestyle5:   5,
		NameStyle6:   6,
		NameStyle7:   7,
		Namestyle8:   8,
		DataType1:    &gojsontest.NameStyleJSONName_Float1{Float1: "float32"},
		Data_Type2:   &gojsontest.NameStyleJSONName_Float2{Float2: "float64"},
		Data_Type3:   &gojsontest.NameStyleJSONName_Integer3{Integer3: "int32"},
		DataType4:    &gojsontest.NameStyleJSONName_Integer4{Integer4: "int64"},
		Datatype5:    &gojsontest.NameStyleJSONName_Integer5{Integer5: "uint32"},
		DataType6:    &gojsontest.NameStyleJSONName_Integer6{Integer6: "uint64"},
		DataType7:    &gojsontest.NameStyleJSONName_Integer7{Integer7: "int"},
		Datatype8:    &gojsontest.NameStyleJSONName_Integer8{Integer8: "uint32"},
	}

	expected := []byte(`{"nameStyle1":1,"namesStyle2":2,"NameStyle3":3,"NameStyle4":4,"namestyle5":5,"nameStyle6":6,"NameStyle7":7,"Namestyle8":8,"data_type1":{"float1":"float32"},"data_Type2":{"float2":"float64"},"Data_Type3":{"integer3":"int32"},"Data_type4":{"integer4":"int64"},"datatype5":{"integer5":"uint32"},"dataType6":{"integer6":"uint64"},"DataType7":{"integer7":"int"},"Datatype8":{"integer8":"uint32"}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.NameStyleJSONName{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)
}

func Test_GoJSON_FieldCustomName(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		data1 := &gojsontest.FieldCustomName{
			TString:          "Hello String",
			TInt32:           11,
			TInt64:           12,
			TUint32:          13,
			TUint64:          14,
			TSint32:          15,
			TSint64:          16,
			TSfixed32:        17,
			TSfixed64:        18,
			TFixed32:         19,
			TFixed64:         20,
			TFloat:           21.21,
			TDouble:          21.22,
			TBool:            true,
			TEnum1:           gojsontest.FieldCustomName_stopped,
			TEnum2:           gojsontest.FieldCustomName_running,
			TBytes:           []byte("Hello Bytes"),
			TAliases:         &gojsontest.FieldCustomName_Aliases{},
			TConfig:          &gojsontest.FieldCustomName_Config{Ip: "192.168.1.1", Port: 8081},
			ArrayDouble:      []float64{1.1, 1.2, 1.3},
			ArrayFloat:       []float32{2.2, 2.2, 2.3},
			ArrayInt32:       []int32{1, 2, 3},
			ArrayInt64:       []int64{1, 2, 3},
			ArrayUint32:      []uint32{1, 2, 3},
			ArrayUint64:      []uint64{1, 2, 3},
			ArraySint32:      []int32{1, 2, 3},
			ArraySint64:      []int64{1, 2, 3},
			ArraySfixed32:    []int32{1, 2, 3},
			ArraySfixed64:    []int64{1, 2, 3},
			ArrayFixed32:     []uint32{1, 2, 3},
			ArrayFixed64:     []uint64{1, 2, 3},
			ArrayBool:        []bool{false, true, false},
			ArrayString:      []string{"s1", "s1", "s1"},
			ArrayBytes:       [][]byte{[]byte("b1"), []byte("b2")},
			ArrayEnum1:       []gojsontest.FieldCustomName_Enum{gojsontest.FieldCustomName_running, gojsontest.FieldCustomName_stopped},
			ArrayEnum2:       []gojsontest.FieldCustomName_Enum{gojsontest.FieldCustomName_stopped, gojsontest.FieldCustomName_running},
			ArrayAliases:     []*gojsontest.FieldCustomName_Aliases{nil, {}, nil, {}, nil},
			ArrayConfig:      []*gojsontest.FieldCustomName_Config{{Ip: "10.1", Port: 80}, nil, {Ip: "10.1", Port: 80}, nil},
			MapInt32Double:   map[int32]float64{1: 1},
			MapInt32Float:    map[int32]float32{1: 1},
			MapInt32Int32:    map[int32]int32{1: 1},
			MapInt32Int64:    map[int32]int64{1: 1},
			MapInt32Uint32:   map[int32]uint32{1: 1},
			MapInt32Uint64:   map[int32]uint64{1: 1},
			MapInt32Sint32:   map[int32]int32{1: 1},
			MapInt32Sint64:   map[int32]int64{1: 1},
			MapInt32Sfixed32: map[int32]int32{1: 1},
			MapInt32Sfixed64: map[int32]int64{1: 1},
			MapInt32Fixed32:  map[int32]uint32{1: 1},
			MapInt32Fixed64:  map[int32]uint64{1: 1},
			MapInt32Bool:     map[int32]bool{1: true},
			MapInt32String:   map[int32]string{1: "s1"},
			MapInt32Bytes:    map[int32][]byte{1: []byte("b1")},
			MapInt32Enum1:    map[int32]gojsontest.FieldCustomName_Enum{1: gojsontest.FieldCustomName_running},
			MapInt32Enum2:    map[int32]gojsontest.FieldCustomName_Enum{1: gojsontest.FieldCustomName_stopped},
			MapInt32Aliases:  map[int32]*gojsontest.FieldCustomName_Aliases{1: nil},
			MapInt32Config:   map[int32]*gojsontest.FieldCustomName_Config{1: {Ip: "10.1", Port: 80}},
			MapInt64Int32:    map[int64]int32{1: 1},
			MapUint32Int32:   map[uint32]int32{1: 1},
			MapUint64Int32:   map[uint64]int32{1: 1},
			MapSint32Int32:   map[int32]int32{1: 1},
			MapSint64Int32:   map[int64]int32{1: 1},
			MapFixed32Int32:  map[uint32]int32{1: 1},
			MapFixed64Int32:  map[uint64]int32{1: 1},
			MapSfixed32Int32: map[int32]int32{1: 1},
			MapSfixed64Int32: map[int64]int32{1: 1},
			MapStringInt32:   map[string]int32{"k1": 1},
			DataType1:        &gojsontest.FieldCustomName_One1TString{One1TString: "ss1"},
			DataType2:        &gojsontest.FieldCustomName_One2TString{One2TString: "ss2"},
		}

		expected := []byte(`{"ts":"Hello String","ti32":11,"ti64":12,"tu32":13,"tu64":14,"tsi32":15,"tsi64":16,"tsf32":17,"tsf64":18,"tfi32":19,"tfi64":20,"tfl":21.21,"tdl":21.22,"tbl":true,"te1":1,"te2":"running","tbs":"SGVsbG8gQnl0ZXM=","ta":{},"tc":{"cf":"192.168.1.1","cp":8081},"adl":[1.1,1.2,1.3],"afl":[2.2,2.2,2.3],"ai32":[1,2,3],"ai64":[1,2,3],"au32":[1,2,3],"au64":[1,2,3],"asi32":[1,2,3],"asi64":[1,2,3],"asf32":[1,2,3],"asf64":[1,2,3],"afi32":[1,2,3],"afi64":[1,2,3],"abl":[false,true,false],"as":["s1","s1","s1"],"abs":["YjE=","YjI="],"ae1":[0,1],"ae2":["stopped","running"],"aa":[null,{},null,{},null],"ac":[{"cf":"10.1","cp":80},null,{"cf":"10.1","cp":80},null],"m32dl":{"1":1},"m32fl":{"1":1},"m32i32":{"1":1},"m32i64":{"1":1},"m32u32":{"1":1},"m32u64":{"1":1},"m32si32":{"1":1},"m32si64":{"1":1},"m32sf32":{"1":1},"m32sf64":{"1":1},"m32fi32":{"1":1},"m32fi64":{"1":1},"m32bl":{"1":true},"m32s":{"1":"s1"},"m32b":{"1":"YjE="},"m32e1":{"1":0},"m32e2":{"1":"stopped"},"m32a":{"1":null},"m32c":{"1":{"cf":"10.1","cp":80}},"mi64i32":{"1":1},"mu32i32":{"1":1},"mu64i32":{"1":1},"ms32i32":{"1":1},"ms64i32":{"1":1},"mf32i32":{"1":1},"mf64i32":{"1":1},"msf32i32":{"1":1},"msf64i32":{"1":1},"msi32":{"k1":1},"dt1":{"o1ts":"ss1"},"o2ts":"ss2"}`)

		b1, err := data1.MarshalJSON()
		require.Nil(t, err)
		//fmt.Println(string(b1))
		require.Equal(t, string(expected), string(b1))

		data2 := &gojsontest.FieldCustomName{}
		err = data2.UnmarshalJSON(expected)
		require.Nil(t, err)
		require.Equal(t, data1, data2)
	})

	t.Run("Case2", func(t *testing.T) {
		type CaseDesc struct {
			Name     string
			Data     *gojsontest.FieldCustomName
			Expected []byte // expected error message.
		}

		cases := []*CaseDesc{
			{"one1_t_string", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TString{One1TString: "ss1"}, DataType2: &gojsontest.FieldCustomName_One2TString{One2TString: "ss2"}}, []byte(`{"te2":"running","dt1":{"o1ts":"ss1"},"o2ts":"ss2"}`)},
			{"one1_t_int32", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TInt32{One1TInt32: 11}, DataType2: &gojsontest.FieldCustomName_One2TInt32{One2TInt32: 22}}, []byte(`{"te2":"running","dt1":{"o1i32":11},"o2i32":22}`)},
			{"one1_t_int64", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TInt64{One1TInt64: 11}, DataType2: &gojsontest.FieldCustomName_One2TInt64{One2TInt64: 22}}, []byte(`{"te2":"running","dt1":{"o1i64":11},"o2i64":22}`)},
			{"one1_t_uint32", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TUint32{One1TUint32: 11}, DataType2: &gojsontest.FieldCustomName_One2TUint32{One2TUint32: 22}}, []byte(`{"te2":"running","dt1":{"o1u32":11},"o2u32":22}`)},
			{"one1_t_uint64", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TUint64{One1TUint64: 11}, DataType2: &gojsontest.FieldCustomName_One2TUint64{One2TUint64: 22}}, []byte(`{"te2":"running","dt1":{"o1u64":11},"o2u64":22}`)},
			{"one1_t_sint32", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TSint32{One1TSint32: 11}, DataType2: &gojsontest.FieldCustomName_One2TSint32{One2TSint32: 22}}, []byte(`{"te2":"running","dt1":{"o1si32":11},"o2si32":22}`)},
			{"one1_t_sint64", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TSint64{One1TSint64: 11}, DataType2: &gojsontest.FieldCustomName_One2TSfixed64{One2TSfixed64: 22}}, []byte(`{"te2":"running","dt1":{"o1si64":11},"o2sf64":22}`)},
			{"one1_t_sfixed32", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TSfixed32{One1TSfixed32: 11}, DataType2: &gojsontest.FieldCustomName_One2TSint32{One2TSint32: 22}}, []byte(`{"te2":"running","dt1":{"o1sf32":11},"o2si32":22}`)},
			{"one1_t_sfixed64", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TSfixed64{One1TSfixed64: 11}, DataType2: &gojsontest.FieldCustomName_One2TSint64{One2TSint64: 22}}, []byte(`{"te2":"running","dt1":{"o1sf64":11},"o2si64":22}`)},
			{"one1_t_fixed32", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TFixed32{One1TFixed32: 11}, DataType2: &gojsontest.FieldCustomName_One2TFixed32{One2TFixed32: 22}}, []byte(`{"te2":"running","dt1":{"o1fi32":11},"o2fi32":22}`)},
			{"one1_t_fixed64", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TFixed64{One1TFixed64: 11}, DataType2: &gojsontest.FieldCustomName_One2TFixed64{One2TFixed64: 22}}, []byte(`{"te2":"running","dt1":{"o1fi64":11},"o2fi64":22}`)},
			{"one1_t_float", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TFloat{One1TFloat: 11}, DataType2: &gojsontest.FieldCustomName_One2TFloat{One2TFloat: 22}}, []byte(`{"te2":"running","dt1":{"o1tf":11},"o2tf":22}`)},
			{"one1_t_double", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TDouble{One1TDouble: 11}, DataType2: &gojsontest.FieldCustomName_One2TDouble{One2TDouble: 22}}, []byte(`{"te2":"running","dt1":{"o1df":11},"o2df":22}`)},
			{"one1_t_bool", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TBool{One1TBool: true}, DataType2: &gojsontest.FieldCustomName_One2TBool{One2TBool: true}}, []byte(`{"te2":"running","dt1":{"o1bl":true},"o2bl":true}`)},
			{"one1_t_enum1", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TEnum1{One1TEnum1: gojsontest.FieldCustomName_stopped}, DataType2: &gojsontest.FieldCustomName_One2TEnum1{One2TEnum1: gojsontest.FieldCustomName_running}}, []byte(`{"te2":"running","dt1":{"o1e1":1}}`)},
			{"one1_t_enum2", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TEnum2{One1TEnum2: gojsontest.FieldCustomName_running}, DataType2: &gojsontest.FieldCustomName_One2TEnum2{One2TEnum2: gojsontest.FieldCustomName_stopped}}, []byte(`{"te2":"running","dt1":{"o1e2":"running"},"o2e2":"stopped"}`)},
			{"one1_t_bytes", &gojsontest.FieldCustomName{DataType1: &gojsontest.FieldCustomName_One1TBytes{One1TBytes: []byte(`b1`)}, DataType2: &gojsontest.FieldCustomName_One2TBytes{One2TBytes: []byte(`b2`)}}, []byte(`{"te2":"running","dt1":{"o1tb":"YjE="},"o2tb":"YjI="}`)},
			{"one1_t_aliases", &gojsontest.FieldCustomName{DataType1: nil, DataType2: &gojsontest.FieldCustomName_One2TAliases{One2TAliases: &gojsontest.FieldCustomName_Aliases{}}}, []byte(`{"te2":"running","o2ta":{}}`)},
			{"one1_t_config", &gojsontest.FieldCustomName{DataType1: nil, DataType2: &gojsontest.FieldCustomName_One2TConfig{One2TConfig: &gojsontest.FieldCustomName_Config{Ip: "10.1.1.1", Port: 90}}}, []byte(`{"te2":"running","o2tc":{"cf":"10.1.1.1","cp":90}}`)},
		}

		for _, c := range cases {
			data1 := c.Data

			b1, err := data1.MarshalJSON()
			require.Nil(t, err, c.Name)
			//fmt.Println(string(b1))
			require.Equal(t, string(c.Expected), string(b1), c.Name)
			//
			if c.Name == "one1_t_enum1" {
				continue
			}
			data2 := &gojsontest.FieldCustomName{}
			err = data2.UnmarshalJSON(c.Expected)
			require.Nil(t, err, c.Name)
			require.Equal(t, data1, data2, c.Name)
		}
	})
}

func Test_GoJSON_OneofHide1(t *testing.T) {
	t.Run("TestMarshal", func(t *testing.T) {
		data1 := &gojsontest.OneofHide1{
			DataType1: &gojsontest.OneofHide1_One1String1{One1String1: "s1"},
			DataType2: &gojsontest.OneofHide1_One2String2{One2String2: "s2"},
		}

		expected := []byte(`{"one1_string1":"s1","one2_string2":"s2"}`)

		b1, err := data1.MarshalJSON()
		require.Nil(t, err)
		//fmt.Println(string(b1))
		require.Equal(t, string(expected), string(b1))

		data2 := &gojsontest.OneofHide1{}
		err = data2.UnmarshalJSON(expected)
		require.Nil(t, err)
		require.Equal(t, data1, data2)
		require.NotNil(t, data2.DataType1)
		require.NotNil(t, data2.DataType2)

		x1, ok := data2.DataType1.(*gojsontest.OneofHide1_One1String1)
		require.True(t, ok)
		require.Equal(t, "s1", x1.One1String1)

		x2, ok := data2.DataType2.(*gojsontest.OneofHide1_One2String2)
		require.True(t, ok)
		require.Equal(t, "s2", x2.One2String2)
	})

	t.Run("TestUnmarshal", func(t *testing.T) {
		{
			b := []byte(`{"one1_string1": "s1", "one2_string1": "s2"}`)
			data := &gojsontest.OneofHide1{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.NotNil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x1, ok := data.DataType1.(*gojsontest.OneofHide1_One1String1)
			require.True(t, ok)
			require.Equal(t, "s1", x1.One1String1)

			x2, ok := data.DataType2.(*gojsontest.OneofHide1_One2String1)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String1)
		}
		{
			b := []byte(`{"one1_string2": "s1", "one2_string2": "s2"}`)
			data := &gojsontest.OneofHide1{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.NotNil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x1, ok := data.DataType1.(*gojsontest.OneofHide1_One1String2)
			require.True(t, ok)
			require.Equal(t, "s1", x1.One1String2)

			x2, ok := data.DataType2.(*gojsontest.OneofHide1_One2String2)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String2)
		}
		{
			b := []byte(`{"data_typ1": {"one1_string2": "s1"}, "one2_string2": "s2"}`)
			data := &gojsontest.OneofHide1{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.Nil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x2, ok := data.DataType2.(*gojsontest.OneofHide1_One2String2)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String2)
		}
	})
}

func Test_GoJSON_OneofHide2(t *testing.T) {
	t.Run("TestMarshal", func(t *testing.T) {
		data1 := &gojsontest.OneofHide2{
			DataType1: &gojsontest.OneofHide2_One1String1{One1String1: "s1"},
			DataType2: &gojsontest.OneofHide2_One2String2{One2String2: "s2"},
		}

		expected := []byte(`{"one1_string1":"s1","data_type2":{"one2_string2":"s2"}}`)

		b1, err := data1.MarshalJSON()
		require.Nil(t, err)
		//fmt.Println(string(b1))
		require.Equal(t, string(expected), string(b1))

		data2 := &gojsontest.OneofHide2{}
		err = data2.UnmarshalJSON(expected)
		require.Nil(t, err)
		require.Equal(t, data1, data2)
		require.NotNil(t, data2.DataType1)
		require.NotNil(t, data2.DataType2)

		x1, ok := data2.DataType1.(*gojsontest.OneofHide2_One1String1)
		require.True(t, ok)
		require.Equal(t, "s1", x1.One1String1)

		x2, ok := data2.DataType2.(*gojsontest.OneofHide2_One2String2)
		require.True(t, ok)
		require.Equal(t, "s2", x2.One2String2)
	})

	t.Run("TestUnmarshal", func(t *testing.T) {
		{
			b := []byte(`{"one1_string1": "s1", "data_type2": {"one2_string1": "s2"}}`)
			data := &gojsontest.OneofHide2{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.NotNil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x1, ok := data.DataType1.(*gojsontest.OneofHide2_One1String1)
			require.True(t, ok)
			require.Equal(t, "s1", x1.One1String1)

			x2, ok := data.DataType2.(*gojsontest.OneofHide2_One2String1)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String1)
		}
		{
			b := []byte(`{"one1_string2": "s1", "data_type2": {"one2_string2": "s2"}}`)
			data := &gojsontest.OneofHide2{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.NotNil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x1, ok := data.DataType1.(*gojsontest.OneofHide2_One1String2)
			require.True(t, ok)
			require.Equal(t, "s1", x1.One1String2)

			x2, ok := data.DataType2.(*gojsontest.OneofHide2_One2String2)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String2)
		}
		{
			b := []byte(`{"data_typ1": {"one1_string2": "s1"}, "data_type2": {"one2_string2": "s2"}}`)
			data := &gojsontest.OneofHide2{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.Nil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x2, ok := data.DataType2.(*gojsontest.OneofHide2_One2String2)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String2)
		}
	})
}

func Test_GoJSON_OneofHide3(t *testing.T) {
	t.Run("TestMarshal", func(t *testing.T) {
		data1 := &gojsontest.OneofHide3{
			DataType1: &gojsontest.OneofHide3_One1String1{One1String1: "s1"},
			DataType2: &gojsontest.OneofHide3_One2String2{One2String2: "s2"},
		}

		expected := []byte(`{"one1_string1":"s1","data_type2":{"one2_string2":"s2"}}`)

		b1, err := data1.MarshalJSON()
		require.Nil(t, err)
		//fmt.Println(string(b1))
		require.Equal(t, string(expected), string(b1))

		data2 := &gojsontest.OneofHide3{}
		err = data2.UnmarshalJSON(expected)
		require.Nil(t, err)
		require.NotNil(t, data2.DataType1)
		require.NotNil(t, data2.DataType2)

		x1, ok := data2.DataType1.(*gojsontest.OneofHide3_One1String1)
		require.True(t, ok)
		require.Equal(t, "s1", x1.One1String1)

		x2, ok := data2.DataType2.(*gojsontest.OneofHide3_One2String2)
		require.True(t, ok)
		require.Equal(t, "s2", x2.One2String2)
	})

	t.Run("TestUnmarshal", func(t *testing.T) {
		{
			b := []byte(`{ "one1_string1": "s1", "data_type2": {"one2_string1": "s2"} }`)
			data := &gojsontest.OneofHide3{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.NotNil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x1, ok := data.DataType1.(*gojsontest.OneofHide3_One1String1)
			require.True(t, ok)
			require.Equal(t, "s1", x1.One1String1)

			x2, ok := data.DataType2.(*gojsontest.OneofHide3_One2String1)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String1)
		}
		{
			b := []byte(`{"one1_string2": "s1", "data_type2": {"one2_string2": "s2"} }`)
			data := &gojsontest.OneofHide3{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.NotNil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x1, ok := data.DataType1.(*gojsontest.OneofHide3_One1String2)
			require.True(t, ok)
			require.Equal(t, "s1", x1.One1String2)

			x2, ok := data.DataType2.(*gojsontest.OneofHide3_One2String2)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String2)
		}
		{
			b := []byte(`{"data_typ1": {"one1_string2": "s1"}, "one2_string2": "s2"}`)
			data := &gojsontest.OneofHide3{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.Nil(t, data.DataType1)
			require.Nil(t, data.DataType2)
		}
	})
}

func Test_GoJSON_OneofHide4(t *testing.T) {
	t.Run("TestMarshal", func(t *testing.T) {
		data1 := &gojsontest.OneofHide4{
			DataType1: &gojsontest.OneofHide4_One1String1{One1String1: "s1"},
			DataType2: &gojsontest.OneofHide4_One2String2{One2String2: "s2"},
		}

		expected := []byte(`{"data_type1":{"one1_string1":"s1"},"data_type2":{"one2_string2":"s2"}}`)

		b1, err := data1.MarshalJSON()
		require.Nil(t, err)
		//fmt.Println(string(b1))
		require.Equal(t, string(expected), string(b1))

		data2 := &gojsontest.OneofHide4{}
		err = data2.UnmarshalJSON(expected)
		require.Nil(t, err)
		require.NotNil(t, data2.DataType1)
		require.NotNil(t, data2.DataType2)

		x1, ok := data2.DataType1.(*gojsontest.OneofHide4_One1String1)
		require.True(t, ok)
		require.Equal(t, "s1", x1.One1String1)

		x2, ok := data2.DataType2.(*gojsontest.OneofHide4_One2String2)
		require.True(t, ok)
		require.Equal(t, "s2", x2.One2String2)
	})

	t.Run("TestUnmarshal", func(t *testing.T) {
		{
			b := []byte(`{ "data_type1":{"one1_string1": "s1"}, "data_type2": {"one2_string1": "s2"} }`)
			data := &gojsontest.OneofHide4{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.NotNil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x1, ok := data.DataType1.(*gojsontest.OneofHide4_One1String1)
			require.True(t, ok)
			require.Equal(t, "s1", x1.One1String1)

			x2, ok := data.DataType2.(*gojsontest.OneofHide4_One2String1)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String1)
		}
		{
			b := []byte(`{"data_type1":{"one1_string2": "s1"}, "data_type2": {"one2_string2": "s2"} }`)
			data := &gojsontest.OneofHide4{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.NotNil(t, data.DataType1)
			require.NotNil(t, data.DataType2)

			x1, ok := data.DataType1.(*gojsontest.OneofHide4_One1String2)
			require.True(t, ok)
			require.Equal(t, "s1", x1.One1String2)

			x2, ok := data.DataType2.(*gojsontest.OneofHide4_One2String2)
			require.True(t, ok)
			require.Equal(t, "s2", x2.One2String2)
		}
		{
			b := []byte(`{"one1_string2": "s1", "one2_string2": "s2"}`)
			data := &gojsontest.OneofHide4{}
			err := data.UnmarshalJSON(b)
			require.Nil(t, err)
			require.Nil(t, data.DataType1)
			require.Nil(t, data.DataType2)
		}
	})
}

func Test_GoJSON_FieldOmitempty1(t *testing.T) {
	data1 := &gojsontest.FieldOmitempty1{
		TString1:  "1",
		TString2:  "",
		DataType1: &gojsontest.FieldOmitempty1_One1Int32{One1Int32: 1},
		DataType2: &gojsontest.FieldOmitempty1_One2Int64{One2Int64: 0},
		DataType3: nil,
	}

	expected := []byte(`{"t_string1":"1","data_type1":{"one1_int32":1}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.FieldOmitempty1{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	data1.DataType2 = nil
	require.Equal(t, data1, data2)
}

func Test_GoJSON_FieldOmitempty2(t *testing.T) {
	data1 := &gojsontest.FieldOmitempty2{
		TString1:  "1",
		TString2:  "",
		TString3:  "3",
		TString4:  "",
		DataType1: &gojsontest.FieldOmitempty2_One1Int32{One1Int32: 1},
		DataType2: &gojsontest.FieldOmitempty2_One2Int64{One2Int64: 0},
		DataType3: nil,
		DataType4: &gojsontest.FieldOmitempty2_One4Uint64{One4Uint64: 0},
		DataType5: nil,
		DataType6: &gojsontest.FieldOmitempty2_One6Sint32{One6Sint32: 0},
		DataType7: &gojsontest.FieldOmitempty2_One7Bool2{One7Bool2: false},
	}

	expected := []byte(`{"t_string1":"1","t_string3":"3","t_string4":"","data_type1":{"one1_int32":1},"data_type5":null,"data_type6":{"one6_sint32":0}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.FieldOmitempty2{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	data1.DataType2 = nil
	data1.DataType4 = nil
	data1.DataType7 = nil
	require.Equal(t, data1, data2)
}

func Test_GoJSON_FieldOmitempty3(t *testing.T) {
	data1 := &gojsontest.FieldOmitempty3{
		TString1:  "1",
		TString2:  "",
		TString3:  "3",
		TString4:  "",
		TString5:  "",
		DataType1: &gojsontest.FieldOmitempty3_One1Int32{One1Int32: 32},
		DataType2: nil,
		DataType3: &gojsontest.FieldOmitempty3_One3Uint32{One3Uint32: 0},
		DataType4: nil,
		DataType5: &gojsontest.FieldOmitempty3_One5String1{One5String1: ""},
		DataType6: &gojsontest.FieldOmitempty3_One6Sint32{One6Sint32: 32},
		DataType7: &gojsontest.FieldOmitempty3_One7Bool1{One7Bool1: false},
		DataType8: &gojsontest.FieldOmitempty3_One8Bool1{One8Bool1: false},
	}

	expected := []byte(`{"t_string1":"1","t_string2":"","t_string3":"3","t_string5":"","dt1":{"one1_int32":32},"dt3":{"one3_uint32":0},"dt4":null,"dt6":{"one6_sint32":32},"dt8":{"one8_bool1":false}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.FieldOmitempty3{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	data1.DataType5 = nil
	data1.DataType7 = nil
	require.Equal(t, data1, data2)
}

func Test_GoJSON_FieldOmitempty4(t *testing.T) {
	data1 := &gojsontest.FieldOmitempty4{
		TString1:  "1",
		TString2:  "",
		TString3:  "3",
		TString4:  "",
		DataType1: &gojsontest.FieldOmitempty4_One1Int32{One1Int32: 1},
		DataType2: &gojsontest.FieldOmitempty4_One2Int64{One2Int64: 0},
		DataType3: nil,
		DataType4: nil,
		DataType5: nil,
		DataType6: &gojsontest.FieldOmitempty4_One6Sint32{One6Sint32: 0},
		DataType7: &gojsontest.FieldOmitempty4_One7Bool2{One7Bool2: false},
	}

	expected := []byte(`{"t_string1":"1","t_string2":"","t_string3":"3","data_type1":{"one1_int32":1},"data_type2":{"one2_int64":0},"data_type3":null,"data_type7":{"one7_bool2":false}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.FieldOmitempty4{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)

	data1.DataType6 = nil

	require.Equal(t, data1, data2)
}

func Test_GoJSON_FieldIgnore1(t *testing.T) {
	data1 := &gojsontest.FieldIgnore1{
		TString1:  "",
		TString2:  "",
		DataType1: nil,
		DataType2: nil,
	}
	_, ok1 := interface{}(data1).(json.Marshaler)
	require.False(t, ok1)
	_, ok2 := (interface{})(data1).(json.Unmarshaler)
	require.False(t, ok2)
}

func Test_GoJSON_FieldIgnore2(t *testing.T) {
	data1 := &gojsontest.FieldIgnore2{
		NameStyle1:   1,
		Names_Style2: 2,
		Name_Style3:  3,
		NameStyle4:   4,
		NameStyle5:   5,
		DataType1:    &gojsontest.FieldIgnore2_Integer1{Integer1: "int32"},
		DataType2:    nil,
		DataType3:    &gojsontest.FieldIgnore2_Integer3{Integer3: "int64"},
		DataType4:    &gojsontest.FieldIgnore2_Integer4{Integer4: ""},
		DataType5:    &gojsontest.FieldIgnore2_Integer5{Integer5: "uint32"},
		DataType6:    &gojsontest.FieldIgnore2_Integer6{Integer6: ""},
		DataType7:    &gojsontest.FieldIgnore2_Float7{Float7: "float32"},
		DataType8:    &gojsontest.FieldIgnore2_Float8{Float8: "float64"},
	}

	expected := []byte(`{"ns1":1,"ns3":3,"ns4":4,"dt3":{"i3":"int64"},"dt4":{"i4":""},"dt7":{"f7":"float32"},"dt8":{"f8":"float64"}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.FieldIgnore2{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.NotEqual(t, data1, data2)

	data1.Names_Style2 = 0
	data1.NameStyle5 = 0
	data1.DataType1 = nil
	data1.DataType2 = nil
	data1.DataType5 = nil
	data1.DataType6 = nil

	require.Equal(t, data1, data2)
}

func Test_GoJSON_FieldDisableUnknown(t *testing.T) {
	data1 := &gojsontest.FieldDisallowUnknown{NameStyle1: 1}

	expected := []byte(`{"ns1":1,"oneof1":null}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.FieldDisallowUnknown{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)

	seed3 := []byte(`{"ns1":1, "ns2": 2}`)
	data3 := &gojsontest.FieldDisallowUnknown{}
	err = data3.UnmarshalJSON(seed3)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), `json: unknown field "ns2"`)

	seed4 := []byte(`{"ns1":1, "oneof1": {"ts3": "11"}}`)
	data4 := &gojsontest.FieldDisallowUnknown{}
	err = data4.UnmarshalJSON(seed4)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), `json: unknown oneof field "ts3"`)

	seed5 := []byte(`{"ns1":1, "ti3": 2}`)
	data5 := &gojsontest.FieldDisallowUnknown{}
	err = data5.UnmarshalJSON(seed5)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), `json: unknown field "ti3"`)
}

func Test_GoJSON_FieldAllowUnknown(t *testing.T) {
	data1 := &gojsontest.FieldAllowUnknown{NameStyle1: 1}

	expected := []byte(`{"ns1":1,"oneof1":null}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.FieldAllowUnknown{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)

	seed3 := []byte(`{"ns1":1, "ns2": 2}`)
	data3 := &gojsontest.FieldAllowUnknown{}
	err = data3.UnmarshalJSON(seed3)
	require.Nil(t, err)
	require.Equal(t, data2, data3)

	seed4 := []byte(`{"ns1":1, "oneof1": {"ts3": "11"}}`)
	data4 := &gojsontest.FieldAllowUnknown{}
	err = data4.UnmarshalJSON(seed4)
	require.Nil(t, err)
	require.Equal(t, data2, data4)

	seed5 := []byte(`{"ns1":1, "ti3": 2}`)
	data5 := &gojsontest.FieldAllowUnknown{}
	err = data5.UnmarshalJSON(seed5)
	require.Nil(t, err)
	require.Equal(t, data2, data5)
}

func Test_GoJSON_EnumUseString1(t *testing.T) {
	data1 := &gojsontest.EnumUseString1{
		TStatus1: gojsontest.EnumUseString1_disabled,
		TStatus2: gojsontest.EnumUseString1_failed,
		AStatus1: []gojsontest.EnumUseString1_Status1{gojsontest.EnumUseString1_disabled, gojsontest.EnumUseString1_enabled},
		AStatus2: []gojsontest.EnumUseString1_Status2{gojsontest.EnumUseString1_failed, gojsontest.EnumUseString1_success},
		AStatus3: nil,
		MStatus1: map[string]gojsontest.EnumUseString1_Status1{"s1": gojsontest.EnumUseString1_disabled},
		MStatus2: map[string]gojsontest.EnumUseString1_Status2{"s1": gojsontest.EnumUseString1_failed},
		MStatus3: nil,
	}

	expected := []byte(`{"t_status1":"disabled","t_status2":1,"a_status1":["disabled","enabled"],"a_status2":[1,0],"a_status3":null,"m_status1":{"s1":1},"m_status2":{"s1":1},"m_status3":null}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.EnumUseString1{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)

	require.Equal(t, data2.TStatus1, gojsontest.EnumUseString1_disabled)
	require.Equal(t, data2.TStatus2, gojsontest.EnumUseString1_failed)

	require.Equal(t, data2.AStatus1[0], gojsontest.EnumUseString1_disabled)
	require.Equal(t, data2.AStatus1[1], gojsontest.EnumUseString1_enabled)
	require.Equal(t, data2.AStatus2[0], gojsontest.EnumUseString1_failed)
	require.Equal(t, data2.AStatus2[1], gojsontest.EnumUseString1_success)
	require.Nil(t, data2.AStatus3)

	require.Equal(t, data2.MStatus1["s1"], gojsontest.EnumUseString1_disabled)
	require.Equal(t, data2.MStatus2["s1"], gojsontest.EnumUseString1_failed)
	require.Nil(t, data2.MStatus3)
}

func Test_GoJSON_EnumUseString2(t *testing.T) {
	data1 := &gojsontest.EnumUseString2{
		TStatus1: gojsontest.EnumUseString2_disabled,
		TStatus2: gojsontest.EnumUseString2_failed,
		AStatus1: []gojsontest.EnumUseString2_Status1{gojsontest.EnumUseString2_disabled, gojsontest.EnumUseString2_enabled},
		AStatus2: []gojsontest.EnumUseString2_Status2{gojsontest.EnumUseString2_failed, gojsontest.EnumUseString2_success},
		AStatus3: nil,
		MStatus1: map[string]gojsontest.EnumUseString2_Status1{"s1": gojsontest.EnumUseString2_disabled},
		MStatus2: map[string]gojsontest.EnumUseString2_Status2{"s1": gojsontest.EnumUseString2_failed},
		MStatus3: nil,
	}

	expected := []byte(`{"t_status1":"disabled","t_status2":"failed","a_status1":["disabled","enabled"],"a_status2":["failed","success"],"a_status3":null,"m_status1":{"s1":"disabled"},"m_status2":{"s1":"failed"},"m_status3":null}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.EnumUseString2{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)

	require.Equal(t, data2.TStatus1, gojsontest.EnumUseString2_disabled)
	require.Equal(t, data2.TStatus2, gojsontest.EnumUseString2_failed)

	require.Equal(t, data2.AStatus1[0], gojsontest.EnumUseString2_disabled)
	require.Equal(t, data2.AStatus1[1], gojsontest.EnumUseString2_enabled)
	require.Equal(t, data2.AStatus2[0], gojsontest.EnumUseString2_failed)
	require.Equal(t, data2.AStatus2[1], gojsontest.EnumUseString2_success)
	require.Nil(t, data2.AStatus3)

	require.Equal(t, data2.MStatus1["s1"], gojsontest.EnumUseString2_disabled)
	require.Equal(t, data2.MStatus2["s1"], gojsontest.EnumUseString2_failed)
	require.Nil(t, data2.MStatus3)
}

func Test_GoJSON_EnumUseString3(t *testing.T) {
	data1 := &gojsontest.EnumUseString3{
		TStatus1: gojsontest.EnumUseString3_disabled,
		TStatus2: gojsontest.EnumUseString3_failed,
		AStatus1: []gojsontest.EnumUseString3_Status1{gojsontest.EnumUseString3_disabled, gojsontest.EnumUseString3_enabled},
		AStatus2: []gojsontest.EnumUseString3_Status2{gojsontest.EnumUseString3_failed, gojsontest.EnumUseString3_success},
		AStatus3: nil,
		MStatus1: map[string]gojsontest.EnumUseString3_Status1{"s1": gojsontest.EnumUseString3_disabled},
		MStatus2: map[string]gojsontest.EnumUseString3_Status2{"s1": gojsontest.EnumUseString3_failed},
		MStatus3: nil,
	}

	expected := []byte(`{"t_status1":1,"t_status2":1,"a_status1":[1,0],"a_status2":[1,0],"a_status3":null,"m_status1":{"s1":1},"m_status2":{"s1":1},"m_status3":null}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.EnumUseString3{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)

	require.Equal(t, data2.TStatus1, gojsontest.EnumUseString3_disabled)
	require.Equal(t, data2.TStatus2, gojsontest.EnumUseString3_failed)

	require.Equal(t, data2.AStatus1[0], gojsontest.EnumUseString3_disabled)
	require.Equal(t, data2.AStatus1[1], gojsontest.EnumUseString3_enabled)
	require.Equal(t, data2.AStatus2[0], gojsontest.EnumUseString3_failed)
	require.Equal(t, data2.AStatus2[1], gojsontest.EnumUseString3_success)
	require.Nil(t, data2.AStatus3)

	require.Equal(t, data2.MStatus1["s1"], gojsontest.EnumUseString3_disabled)
	require.Equal(t, data2.MStatus2["s1"], gojsontest.EnumUseString3_failed)
	require.Nil(t, data2.MStatus3)
}

func Test_GoJSON_EnumUseString4(t *testing.T) {
	data1 := &gojsontest.EnumUseString4{
		TStatus1: gojsontest.EnumUseString4_disabled,
		TStatus2: gojsontest.EnumUseString4_enabled,
		AStatus1: []gojsontest.EnumUseString4_Status{gojsontest.EnumUseString4_disabled, gojsontest.EnumUseString4_enabled},
		AStatus2: []gojsontest.EnumUseString4_Status{gojsontest.EnumUseString4_disabled, gojsontest.EnumUseString4_enabled},
		AStatus3: nil,
		MStatus1: map[string]gojsontest.EnumUseString4_Status{"s1": gojsontest.EnumUseString4_disabled},
		MStatus2: map[string]gojsontest.EnumUseString4_Status{"s1": gojsontest.EnumUseString4_enabled},
		MStatus3: nil,
	}

	expected := []byte(`{"t_status1":"disabled","t_status2":1,"a_status1":["disabled","enabled"],"a_status2":[2,1],"a_status3":null,"m_status1":{"s1":"disabled"},"m_status2":{"s1":1},"m_status3":null}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.EnumUseString4{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)

	require.Equal(t, data2.TStatus1, gojsontest.EnumUseString4_disabled)
	require.Equal(t, data2.TStatus2, gojsontest.EnumUseString4_enabled)

	require.Equal(t, data2.AStatus1[0], gojsontest.EnumUseString4_disabled)
	require.Equal(t, data2.AStatus1[1], gojsontest.EnumUseString4_enabled)
	require.Equal(t, data2.AStatus2[0], gojsontest.EnumUseString4_disabled)
	require.Equal(t, data2.AStatus2[1], gojsontest.EnumUseString4_enabled)
	require.Nil(t, data2.AStatus3)

	require.Equal(t, data2.MStatus1["s1"], gojsontest.EnumUseString4_disabled)
	require.Equal(t, data2.MStatus2["s1"], gojsontest.EnumUseString4_enabled)
	require.Nil(t, data2.MStatus3)
}

func Test_GoJSON_EnumUseString5(t *testing.T) {
	data1 := &gojsontest.EnumUseString5{
		TStatus: gojsontest.EnumUseString5_disabled,
		AStatus: []gojsontest.EnumUseString5_Status{gojsontest.EnumUseString5_disabled, gojsontest.EnumUseString5_enabled},
		MStatus: map[string]gojsontest.EnumUseString5_Status{
			"s1": gojsontest.EnumUseString5_disabled,
		},
	}

	expected := []byte(`{"t_status":"disabled","a_status":["disabled","enabled"],"m_status":{"s1":"disabled"}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, string(expected), string(b1))

	data2 := &gojsontest.EnumUseString5{}
	err = data2.UnmarshalJSON(expected)
	require.Nil(t, err)
	require.Equal(t, data1, data2)
	require.Equal(t, data2.TStatus, gojsontest.EnumUseString5_disabled)
	require.Equal(t, data2.AStatus[0], gojsontest.EnumUseString5_disabled)
	require.Equal(t, data2.AStatus[1], gojsontest.EnumUseString5_enabled)
	require.Equal(t, data2.MStatus["s1"], gojsontest.EnumUseString5_disabled)
}

func Test_GoJSON_SerializeBytes1(t *testing.T) {
	type SerializeBytes struct {
		Bytes1      []byte            `json:"bytes1,omitempty"`
		Bytes2      []byte            `json:"bytes2,omitempty"`
		Bytes3      []byte            `json:"bytes3,omitempty"`
		ArrayBytes1 [][]byte          `json:"array_bytes1,omitempty"`
		ArrayBytes2 [][]byte          `json:"array_bytes2,omitempty"`
		ArrayBytes3 [][]byte          `json:"array_bytes3,omitempty"`
		MapBytes1   map[string][]byte `json:"map_bytes1,omitempty"`
		MapBytes2   map[string][]byte `json:"map_bytes2,omitempty"`
		MapBytes3   map[string][]byte `json:"map_bytes3,omitempty"`
		MapBytes4   map[string][]byte `json:"map_bytes4,omitempty"`
	}

	data1 := &gojsontest.SerializeBytes1{
		Bytes1: []byte("Bytes1"),
		Bytes2: []byte(""),
		Bytes3: nil,
		ArrayBytes1: [][]byte{
			[]byte("ArrayBytes11"),
			[]byte("ArrayBytes12"),
			[]byte(""),
		},
		ArrayBytes2: [][]byte{},
		ArrayBytes3: nil,
		MapBytes1: map[string][]byte{
			"key1": []byte("value1"),
		},
		MapBytes2: map[string][]byte{
			"key2": []byte(""),
		},
		MapBytes3: make(map[string][]byte),
		MapBytes4: nil,
	}

	data2 := &SerializeBytes{
		Bytes1: []byte("Bytes1"),
		Bytes2: []byte(""),
		Bytes3: nil,
		ArrayBytes1: [][]byte{
			[]byte("ArrayBytes11"),
			[]byte("ArrayBytes12"),
			[]byte(""),
		},
		ArrayBytes2: [][]byte{},
		ArrayBytes3: nil,
		MapBytes1: map[string][]byte{
			"key1": []byte("value1"),
		},
		MapBytes2: map[string][]byte{
			"key2": []byte(""),
		},
		MapBytes3: make(map[string][]byte),
		MapBytes4: nil,
	}

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)

	b2, err := json.Marshal(data2)
	require.Nil(t, err)

	//fmt.Println(string(b1))
	//fmt.Println(string(b2))

	require.Equal(t, b1, b2)

	data3 := &gojsontest.SerializeBytes1{}
	data4 := &SerializeBytes{}

	err = data3.UnmarshalJSON(b2)
	require.Nil(t, err)
	err = json.Unmarshal(b1, data4)
	require.Nil(t, err)

	require.NotEqual(t, data1, data3)
	data1.Bytes2 = nil
	data1.ArrayBytes2 = nil
	data1.MapBytes3 = nil
	require.Equal(t, data1, data3)

	require.NotEqual(t, data2, data4)
	data2.Bytes2 = nil
	data2.ArrayBytes2 = nil
	data2.MapBytes3 = nil
	require.Equal(t, data2, data4)
}

func Test_GoJSON_SerializeBytes2(t *testing.T) {
	type SerializeBytes struct {
		Bytes1      []byte            `json:"bytes1"`
		Bytes2      []byte            `json:"bytes2"`
		Bytes3      []byte            `json:"bytes3"`
		ArrayBytes1 [][]byte          `json:"array_bytes1"`
		ArrayBytes2 [][]byte          `json:"array_bytes2"`
		ArrayBytes3 [][]byte          `json:"array_bytes3"`
		MapBytes1   map[string][]byte `json:"map_bytes1"`
		MapBytes2   map[string][]byte `json:"map_bytes2"`
		MapBytes3   map[string][]byte `json:"map_bytes3"`
		MapBytes4   map[string][]byte `json:"map_bytes4"`
	}

	data1 := &gojsontest.SerializeBytes2{
		Bytes1: []byte("Bytes1"),
		Bytes2: []byte(""),
		Bytes3: nil,
		ArrayBytes1: [][]byte{
			[]byte("ArrayBytes11"),
			[]byte("ArrayBytes12"),
			[]byte(""),
		},
		ArrayBytes2: [][]byte{},
		ArrayBytes3: nil,
		MapBytes1: map[string][]byte{
			"key1": []byte("value1"),
		},
		MapBytes2: map[string][]byte{
			"key2": []byte(""),
		},
		MapBytes3: make(map[string][]byte),
		MapBytes4: nil,
	}

	data2 := &SerializeBytes{
		Bytes1: []byte("Bytes1"),
		Bytes2: []byte(""),
		Bytes3: nil,
		ArrayBytes1: [][]byte{
			[]byte("ArrayBytes11"),
			[]byte("ArrayBytes12"),
			[]byte(""),
		},
		ArrayBytes2: [][]byte{},
		ArrayBytes3: nil,
		MapBytes1: map[string][]byte{
			"key1": []byte("value1"),
		},
		MapBytes2: map[string][]byte{
			"key2": []byte(""),
		},
		MapBytes3: make(map[string][]byte),
		MapBytes4: nil,
	}

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)

	b2, err := json.Marshal(data2)
	require.Nil(t, err)

	//fmt.Println(string(b1))
	//fmt.Println(string(b2))

	require.Equal(t, b1, b2)

	data3 := &gojsontest.SerializeBytes2{}
	data4 := &SerializeBytes{}

	err = data3.UnmarshalJSON(b2)
	require.Nil(t, err)
	err = json.Unmarshal(b1, data4)
	require.Nil(t, err)

	require.Equal(t, data1, data3)
	require.Equal(t, data2, data4)
}

func Test_GoJSON_SerializeOmitempty1(t *testing.T) {
	type SerializeOmitempty1 struct {
		String1       string                                      `json:"string1,omitempty"`
		String2       string                                      `json:"string2,omitempty"`
		Bytes1        []byte                                      `json:"bytes1,omitempty"`
		Bytes2        []byte                                      `json:"bytes2,omitempty"`
		Bytes3        []byte                                      `json:"bytes3,omitempty"`
		ArrayString1  []string                                    `json:"array_string1,omitempty"`
		ArrayString2  []string                                    `json:"array_string2,omitempty"`
		ArrayString3  []string                                    `json:"array_string3,omitempty"`
		ArrayMessage1 []*gojsonexternal.ExternalMessage1          `json:"array_message1,omitempty"`
		ArrayMessage2 []*gojsonexternal.ExternalMessage1          `json:"array_message2,omitempty"`
		ArrayMessage3 []*gojsonexternal.ExternalMessage1          `json:"array_message3,omitempty"`
		ArrayEnum1    []gojsonexternal.ExternalEnum1              `json:"array_enum1,omitempty"`
		ArrayEnum2    []gojsonexternal.ExternalEnum1              `json:"array_enum2,omitempty"`
		ArrayEnum3    []gojsonexternal.ExternalEnum1              `json:"array_enum3,omitempty"`
		MapString1    map[string]string                           `json:"map_string1,omitempty"`
		MapString2    map[string]string                           `json:"map_string2,omitempty"`
		MapString3    map[string]string                           `json:"map_string3,omitempty"`
		MapMessage1   map[string]*gojsonexternal.ExternalMessage1 `json:"map_message1,omitempty"`
		MapMessage2   map[string]*gojsonexternal.ExternalMessage1 `json:"map_message2,omitempty"`
		MapMessage3   map[string]*gojsonexternal.ExternalMessage1 `json:"map_message3,omitempty"`
		MapEnum1      map[string]gojsonexternal.ExternalEnum1     `json:"map_enum1,omitempty"`
		MapEnum2      map[string]gojsonexternal.ExternalEnum1     `json:"map_enum2,omitempty"`
		MapEnum3      map[string]gojsonexternal.ExternalEnum1     `json:"map_enum3,omitempty"`
	}

	data1 := &gojsontest.SerializeOmitempty1{
		String1:      "String1",
		String2:      "",
		Bytes1:       []byte("Bytes1"),
		Bytes2:       []byte(""),
		Bytes3:       nil,
		ArrayString1: []string{"s1", "s2", ""},
		ArrayString2: []string{},
		ArrayString3: nil,
		ArrayMessage1: []*gojsonexternal.ExternalMessage1{
			{Ip1: "10.10", Ip2: "", Ip3: "10.12"},
			{Ip1: "10.10", Ip2: "10.11", Ip3: "10.12"},
			{Ip1: "10.10", Ip2: "", Ip3: ""},
		},
		ArrayMessage2: []*gojsonexternal.ExternalMessage1{},
		ArrayMessage3: nil,
		ArrayEnum1: []gojsonexternal.ExternalEnum1{
			gojsonexternal.ExternalEnum1_Monday,
			gojsonexternal.ExternalEnum1_Tuesday,
			gojsonexternal.ExternalEnum1_Wednesday,
			gojsonexternal.ExternalEnum1_Thursday,
			gojsonexternal.ExternalEnum1_Friday,
			gojsonexternal.ExternalEnum1_Saturday,
			gojsonexternal.ExternalEnum1_Sunday,
		},
		ArrayEnum2: []gojsonexternal.ExternalEnum1{},
		ArrayEnum3: nil,
		MapString1: map[string]string{"k1": "v1"},
		MapString2: make(map[string]string),
		MapString3: nil,
		MapMessage1: map[string]*gojsonexternal.ExternalMessage1{
			"k1": {Ip1: "10.10", Ip2: "10.11", Ip3: "10.12"},
		},
		MapMessage2: make(map[string]*gojsonexternal.ExternalMessage1),
		MapMessage3: nil,
		MapEnum1: map[string]gojsonexternal.ExternalEnum1{
			"m1": gojsonexternal.ExternalEnum1_Monday,
		},
		MapEnum2: make(map[string]gojsonexternal.ExternalEnum1),
		MapEnum3: nil,
	}

	data2 := &SerializeOmitempty1{
		String1:      "String1",
		String2:      "",
		Bytes1:       []byte("Bytes1"),
		Bytes2:       []byte(""),
		Bytes3:       nil,
		ArrayString1: []string{"s1", "s2", ""},
		ArrayString2: []string{},
		ArrayString3: nil,
		ArrayMessage1: []*gojsonexternal.ExternalMessage1{
			{Ip1: "10.10", Ip2: "", Ip3: "10.12"},
			{Ip1: "10.10", Ip2: "10.11", Ip3: "10.12"},
			{Ip1: "10.10", Ip2: "", Ip3: ""},
		},
		ArrayMessage2: []*gojsonexternal.ExternalMessage1{},
		ArrayMessage3: nil,
		ArrayEnum1: []gojsonexternal.ExternalEnum1{
			gojsonexternal.ExternalEnum1_Monday,
			gojsonexternal.ExternalEnum1_Tuesday,
			gojsonexternal.ExternalEnum1_Wednesday,
			gojsonexternal.ExternalEnum1_Thursday,
			gojsonexternal.ExternalEnum1_Friday,
			gojsonexternal.ExternalEnum1_Saturday,
			gojsonexternal.ExternalEnum1_Sunday,
		},
		ArrayEnum2: []gojsonexternal.ExternalEnum1{},
		ArrayEnum3: nil,
		MapString1: map[string]string{"k1": "v1"},
		MapString2: make(map[string]string),
		MapString3: nil,
		MapMessage1: map[string]*gojsonexternal.ExternalMessage1{
			"k1": {Ip1: "10.10", Ip2: "10.11", Ip3: "10.12"},
		},
		MapMessage2: make(map[string]*gojsonexternal.ExternalMessage1),
		MapMessage3: nil,
		MapEnum1: map[string]gojsonexternal.ExternalEnum1{
			"m1": gojsonexternal.ExternalEnum1_Monday,
		},
		MapEnum2: make(map[string]gojsonexternal.ExternalEnum1),
		MapEnum3: nil,
	}

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)

	b2, err := json.Marshal(data2)
	require.Nil(t, err)

	//fmt.Println(string(b1))
	//fmt.Println(string(b2))

	require.Equal(t, b1, b2)

	data3 := &gojsontest.SerializeOmitempty1{}
	data4 := &SerializeOmitempty1{}

	err = data3.UnmarshalJSON(b2)
	require.Nil(t, err)
	err = json.Unmarshal(b1, data4)
	require.Nil(t, err)

	require.NotEqual(t, data1, data3)
	data1.Bytes2 = nil
	data1.Bytes3 = nil
	data1.ArrayString2 = nil
	data1.ArrayString3 = nil
	data1.ArrayMessage2 = nil
	data1.ArrayMessage3 = nil
	data1.ArrayEnum2 = nil
	data1.ArrayEnum3 = nil
	data1.MapString2 = nil
	data1.MapString3 = nil
	data1.MapMessage2 = nil
	data1.MapMessage3 = nil
	data1.MapEnum2 = nil
	data1.MapEnum3 = nil
	require.Equal(t, data1, data3)

	require.NotEqual(t, data2, data4)
	data2.Bytes2 = nil
	data2.Bytes3 = nil
	data2.ArrayString2 = nil
	data2.ArrayString3 = nil
	data2.ArrayMessage2 = nil
	data2.ArrayMessage3 = nil
	data2.ArrayEnum2 = nil
	data2.ArrayEnum3 = nil
	data2.MapString2 = nil
	data2.MapString3 = nil
	data2.MapMessage2 = nil
	data2.MapMessage3 = nil
	data2.MapEnum2 = nil
	data2.MapEnum3 = nil
	require.Equal(t, data2, data4)
}

func Test_GoJSON_SerializeOmitempty2(t *testing.T) {
	type SerializeOmitempty2 struct {
		String1       string                                      `json:"string1"`
		String2       string                                      `json:"string2"`
		Bytes1        []byte                                      `json:"bytes1"`
		Bytes2        []byte                                      `json:"bytes2"`
		Bytes3        []byte                                      `json:"bytes3"`
		ArrayString1  []string                                    `json:"array_string1"`
		ArrayString2  []string                                    `json:"array_string2"`
		ArrayString3  []string                                    `json:"array_string3"`
		ArrayMessage1 []*gojsonexternal.ExternalMessage1          `json:"array_message1"`
		ArrayMessage2 []*gojsonexternal.ExternalMessage1          `json:"array_message2"`
		ArrayMessage3 []*gojsonexternal.ExternalMessage1          `json:"array_message3"`
		ArrayEnum1    []gojsonexternal.ExternalEnum1              `json:"array_enum1"`
		ArrayEnum2    []gojsonexternal.ExternalEnum1              `json:"array_enum2"`
		ArrayEnum3    []gojsonexternal.ExternalEnum1              `json:"array_enum3"`
		MapString1    map[string]string                           `json:"map_string1"`
		MapString2    map[string]string                           `json:"map_string2"`
		MapString3    map[string]string                           `json:"map_string3"`
		MapMessage1   map[string]*gojsonexternal.ExternalMessage1 `json:"map_message1"`
		MapMessage2   map[string]*gojsonexternal.ExternalMessage1 `json:"map_message2"`
		MapMessage3   map[string]*gojsonexternal.ExternalMessage1 `json:"map_message3"`
		MapEnum1      map[string]gojsonexternal.ExternalEnum1     `json:"map_enum1"`
		MapEnum2      map[string]gojsonexternal.ExternalEnum1     `json:"map_enum2"`
		MapEnum3      map[string]gojsonexternal.ExternalEnum1     `json:"map_enum3"`
	}

	data1 := &gojsontest.SerializeOmitempty2{
		String1:      "String1",
		String2:      "",
		Bytes1:       []byte("Bytes1"),
		Bytes2:       []byte(""),
		Bytes3:       nil,
		ArrayString1: []string{"s1", "s2", ""},
		ArrayString2: []string{},
		ArrayString3: nil,
		ArrayMessage1: []*gojsonexternal.ExternalMessage1{
			{Ip1: "10.10", Ip2: "", Ip3: "10.12"},
			{Ip1: "10.10", Ip2: "10.11", Ip3: "10.12"},
			{Ip1: "10.10", Ip2: "", Ip3: ""},
		},
		ArrayMessage2: []*gojsonexternal.ExternalMessage1{},
		ArrayMessage3: nil,
		ArrayEnum1: []gojsonexternal.ExternalEnum1{
			gojsonexternal.ExternalEnum1_Monday,
			gojsonexternal.ExternalEnum1_Tuesday,
			gojsonexternal.ExternalEnum1_Wednesday,
			gojsonexternal.ExternalEnum1_Thursday,
			gojsonexternal.ExternalEnum1_Friday,
			gojsonexternal.ExternalEnum1_Saturday,
			gojsonexternal.ExternalEnum1_Sunday,
		},
		ArrayEnum2: []gojsonexternal.ExternalEnum1{},
		ArrayEnum3: nil,
		MapString1: map[string]string{"k1": "v1"},
		MapString2: make(map[string]string),
		MapString3: nil,
		MapMessage1: map[string]*gojsonexternal.ExternalMessage1{
			"k1": {Ip1: "10.10", Ip2: "10.11", Ip3: "10.12"},
		},
		MapMessage2: make(map[string]*gojsonexternal.ExternalMessage1),
		MapMessage3: nil,
		MapEnum1: map[string]gojsonexternal.ExternalEnum1{
			"m1": gojsonexternal.ExternalEnum1_Monday,
		},
		MapEnum2: make(map[string]gojsonexternal.ExternalEnum1),
		MapEnum3: nil,
	}

	data2 := &SerializeOmitempty2{
		String1:      "String1",
		String2:      "",
		Bytes1:       []byte("Bytes1"),
		Bytes2:       []byte(""),
		Bytes3:       nil,
		ArrayString1: []string{"s1", "s2", ""},
		ArrayString2: []string{},
		ArrayString3: nil,
		ArrayMessage1: []*gojsonexternal.ExternalMessage1{
			{Ip1: "10.10", Ip2: "", Ip3: "10.12"},
			{Ip1: "10.10", Ip2: "10.11", Ip3: "10.12"},
			{Ip1: "10.10", Ip2: "", Ip3: ""},
		},
		ArrayMessage2: []*gojsonexternal.ExternalMessage1{},
		ArrayMessage3: nil,
		ArrayEnum1: []gojsonexternal.ExternalEnum1{
			gojsonexternal.ExternalEnum1_Monday,
			gojsonexternal.ExternalEnum1_Tuesday,
			gojsonexternal.ExternalEnum1_Wednesday,
			gojsonexternal.ExternalEnum1_Thursday,
			gojsonexternal.ExternalEnum1_Friday,
			gojsonexternal.ExternalEnum1_Saturday,
			gojsonexternal.ExternalEnum1_Sunday,
		},
		ArrayEnum2: []gojsonexternal.ExternalEnum1{},
		ArrayEnum3: nil,
		MapString1: map[string]string{"k1": "v1"},
		MapString2: make(map[string]string),
		MapString3: nil,
		MapMessage1: map[string]*gojsonexternal.ExternalMessage1{
			"k1": {Ip1: "10.10", Ip2: "10.11", Ip3: "10.12"},
		},
		MapMessage2: make(map[string]*gojsonexternal.ExternalMessage1),
		MapMessage3: nil,
		MapEnum1: map[string]gojsonexternal.ExternalEnum1{
			"m1": gojsonexternal.ExternalEnum1_Monday,
		},
		MapEnum2: make(map[string]gojsonexternal.ExternalEnum1),
		MapEnum3: nil,
	}

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)

	b2, err := json.Marshal(data2)
	require.Nil(t, err)

	//fmt.Println(string(b1))
	//fmt.Println(string(b2))

	require.Equal(t, b1, b2)

	data3 := &gojsontest.SerializeOmitempty2{}
	data4 := &SerializeOmitempty2{}

	err = data3.UnmarshalJSON(b2)
	require.Nil(t, err)
	err = json.Unmarshal(b1, data4)
	require.Nil(t, err)

	require.Equal(t, data1, data3)
	require.Equal(t, data2, data4)
}

func Test_GoJSON_MarshalNull(t *testing.T) {
	var model *gojsontest.Model1

	b, err := model.MarshalJSON()
	require.Nil(t, err)

	require.Equal(t, []byte("null"), b)
}

func Test_GoJSON_UnmarshalNull(t *testing.T) {
	model := new(gojsontest.Model1)
	err := model.UnmarshalJSON([]byte("null"))
	require.Nil(t, err)
}

// Expected error.
func Test_GoJSON_UnmarshalStructNil(t *testing.T) {
	var model *gojsontest.Model1

	err := model.UnmarshalJSON([]byte("null"))
	require.Equal(t, "json: Unmarshal: tests/gojsontest.(*Model1) is nil", err.Error())
	require.Nil(t, model)
}

func Test_GoJSON_UnmarshalData_CheckCorrect(t *testing.T) {
	type UnmarshalData gojsontest.UnmarshalData

	type CaseDesc struct {
		Name     string
		B        []byte
		Expected interface{}        // Expected value.
		Actual1  func() interface{} // Actual value.
		Actual2  func() interface{} // Actual value.
	}

	b1 := []byte("Hello Bytes 1")
	b2 := []byte("Hello Bytes 2")

	bb1 := make([]byte, base64.StdEncoding.EncodedLen(len(b1)))
	base64.StdEncoding.Encode(bb1, b1)

	bb2 := make([]byte, base64.StdEncoding.EncodedLen(len(b2)))
	base64.StdEncoding.Encode(bb2, b2)

	data1 := &gojsontest.UnmarshalData{}
	data2 := &UnmarshalData{}

	{
		_, ok := (interface{}(data1)).(json.Marshaler)
		require.True(t, ok)
		_, ok = (interface{}(data2)).(json.Marshaler)
		require.False(t, ok)

		// Test empty
		b := []byte(`{ }`)
		err := data1.UnmarshalJSON(b)
		require.Nil(t, err)
		err = json.Unmarshal(b, data2)
		require.Nil(t, err)

		// Test null
		b = []byte(`null`)
		err = data1.UnmarshalJSON(b)
		require.Nil(t, err)
		err = json.Unmarshal(b, data2)
		require.Nil(t, err)
	}

	// {"Empty", []byte(`{}`), nil, func() interface{} { return nil }, func() interface{} { return nil }},

	literalCases := []*CaseDesc{
		{"t_string 1", []byte(`{"t_string": null}`), "", func() interface{} { return data1.TString }, func() interface{} { return data2.TString }},
		{"t_string 2", []byte(`{"t_string": "Hello World"}`), "Hello World", func() interface{} { return data1.TString }, func() interface{} { return data2.TString }},
		{"t_string 4", []byte(`{"t_string": "Hello C"}`), "Hello C", func() interface{} { return data1.TString }, func() interface{} { return data2.TString }},
		{"t_string 5", []byte(`{}`), "Hello C", func() interface{} { return data1.TString }, func() interface{} { return data2.TString }},
		{"t_string 6", []byte(`{"t_string": ""}`), "", func() interface{} { return data1.TString }, func() interface{} { return data2.TString }},
		{"t_string 7", []byte(`{"t_string": "Hello String"}`), "Hello String", func() interface{} { return data1.TString }, func() interface{} { return data2.TString }},

		{"t_int32 1", []byte(`{"t_int32": 1101}`), int32(1101), func() interface{} { return data1.TInt32 }, func() interface{} { return data2.TInt32 }},
		{"t_int32 2", []byte(`{"t_int32": 0}`), int32(0), func() interface{} { return data1.TInt32 }, func() interface{} { return data2.TInt32 }},
		{"t_int32 3", []byte(`{"t_int32": 1101}`), int32(1101), func() interface{} { return data1.TInt32 }, func() interface{} { return data2.TInt32 }},
		{"t_int32 4", []byte(`{}`), int32(1101), func() interface{} { return data1.TInt32 }, func() interface{} { return data2.TInt32 }},

		{"t_int64 1", []byte(`{"t_int64": 1102}`), int64(1102), func() interface{} { return data1.TInt64 }, func() interface{} { return data2.TInt64 }},
		{"t_int64 2", []byte(`{"t_int64": 0}`), int64(0), func() interface{} { return data1.TInt64 }, func() interface{} { return data2.TInt64 }},
		{"t_int64 3", []byte(`{"t_int64": 1102}`), int64(1102), func() interface{} { return data1.TInt64 }, func() interface{} { return data2.TInt64 }},
		{"t_int64 4", []byte(`{}`), int64(1102), func() interface{} { return data1.TInt64 }, func() interface{} { return data2.TInt64 }},

		{"t_uint32 1", []byte(`{"t_uint32": 1103}`), uint32(1103), func() interface{} { return data1.TUint32 }, func() interface{} { return data2.TUint32 }},
		{"t_uint32 2", []byte(`{"t_uint32": 0}`), uint32(0), func() interface{} { return data1.TUint32 }, func() interface{} { return data2.TUint32 }},
		{"t_uint32 3", []byte(`{"t_uint32": 1103}`), uint32(1103), func() interface{} { return data1.TUint32 }, func() interface{} { return data2.TUint32 }},
		{"t_uint32 4", []byte(`{}`), uint32(1103), func() interface{} { return data1.TUint32 }, func() interface{} { return data2.TUint32 }},

		{"t_uint64 1", []byte(`{"t_uint64": 1104}`), uint64(1104), func() interface{} { return data1.TUint64 }, func() interface{} { return data2.TUint64 }},
		{"t_uint64 2", []byte(`{"t_uint64": 0}`), uint64(0), func() interface{} { return data1.TUint64 }, func() interface{} { return data2.TUint64 }},
		{"t_uint64 3", []byte(`{"t_uint64": 1104}`), uint64(1104), func() interface{} { return data1.TUint64 }, func() interface{} { return data2.TUint64 }},
		{"t_uint64 4", []byte(`{}`), uint64(1104), func() interface{} { return data1.TUint64 }, func() interface{} { return data2.TUint64 }},

		{"t_sint32 1", []byte(`{"t_sint32": 1101}`), int32(1101), func() interface{} { return data1.TSint32 }, func() interface{} { return data2.TSint32 }},
		{"t_sint32 2", []byte(`{"t_sint32": 0}`), int32(0), func() interface{} { return data1.TSint32 }, func() interface{} { return data2.TSint32 }},
		{"t_sint32 3", []byte(`{"t_sint32": 1101}`), int32(1101), func() interface{} { return data1.TSint32 }, func() interface{} { return data2.TSint32 }},
		{"t_sint32 4", []byte(`{}`), int32(1101), func() interface{} { return data1.TSint32 }, func() interface{} { return data2.TSint32 }},

		{"t_sint64 1", []byte(`{"t_sint64": 1102}`), int64(1102), func() interface{} { return data1.TSint64 }, func() interface{} { return data2.TSint64 }},
		{"t_sint64 2", []byte(`{"t_sint64": 0}`), int64(0), func() interface{} { return data1.TSint64 }, func() interface{} { return data2.TSint64 }},
		{"t_sint64 3", []byte(`{"t_sint64": 1102}`), int64(1102), func() interface{} { return data1.TSint64 }, func() interface{} { return data2.TSint64 }},
		{"t_sint64 4", []byte(`{}`), int64(1102), func() interface{} { return data1.TSint64 }, func() interface{} { return data2.TSint64 }},

		{"t_sfixed32 1", []byte(`{"t_sfixed32": 1105}`), int32(1105), func() interface{} { return data1.TSfixed32 }, func() interface{} { return data2.TSfixed32 }},
		{"t_sfixed32 2", []byte(`{"t_sfixed32": 0}`), int32(0), func() interface{} { return data1.TSfixed32 }, func() interface{} { return data2.TSfixed32 }},
		{"t_sfixed32 3", []byte(`{"t_sfixed32": 1105}`), int32(1105), func() interface{} { return data1.TSfixed32 }, func() interface{} { return data2.TSfixed32 }},
		{"t_sfixed32 4", []byte(`{}`), int32(1105), func() interface{} { return data1.TSfixed32 }, func() interface{} { return data2.TSfixed32 }},

		{"t_sfixed64 1", []byte(`{"t_sfixed64": 1106}`), int64(1106), func() interface{} { return data1.TSfixed64 }, func() interface{} { return data2.TSfixed64 }},
		{"t_sfixed64 2", []byte(`{"t_sfixed64": 0}`), int64(0), func() interface{} { return data1.TSfixed64 }, func() interface{} { return data2.TSfixed64 }},
		{"t_sfixed64 3", []byte(`{"t_sfixed64": 1106}`), int64(1106), func() interface{} { return data1.TSfixed64 }, func() interface{} { return data2.TSfixed64 }},
		{"t_sfixed64 4", []byte(`{}`), int64(1106), func() interface{} { return data1.TSfixed64 }, func() interface{} { return data2.TSfixed64 }},

		{"t_fixed32 1", []byte(`{"t_fixed32": 1107}`), uint32(1107), func() interface{} { return data1.TFixed32 }, func() interface{} { return data2.TFixed32 }},
		{"t_fixed32 2", []byte(`{"t_fixed32": 0}`), uint32(0), func() interface{} { return data1.TFixed32 }, func() interface{} { return data2.TFixed32 }},
		{"t_fixed32 3", []byte(`{"t_fixed32": 1107}`), uint32(1107), func() interface{} { return data1.TFixed32 }, func() interface{} { return data2.TFixed32 }},
		{"t_fixed32 4", []byte(`{}`), uint32(1107), func() interface{} { return data1.TFixed32 }, func() interface{} { return data2.TFixed32 }},

		{"t_fixed64 1", []byte(`{"t_fixed64": 1108}`), uint64(1108), func() interface{} { return data1.TFixed64 }, func() interface{} { return data2.TFixed64 }},
		{"t_fixed64 2", []byte(`{"t_fixed64": 0}`), uint64(0), func() interface{} { return data1.TFixed64 }, func() interface{} { return data2.TFixed64 }},
		{"t_fixed64 3", []byte(`{"t_fixed64": 1108}`), uint64(1108), func() interface{} { return data1.TFixed64 }, func() interface{} { return data2.TFixed64 }},
		{"t_fixed64 4", []byte(`{}`), uint64(1108), func() interface{} { return data1.TFixed64 }, func() interface{} { return data2.TFixed64 }},

		{"t_float 1", []byte(`{"t_float": 1109.1}`), float32(1109.1), func() interface{} { return data1.TFloat }, func() interface{} { return data2.TFloat }},
		{"t_float 2", []byte(`{"t_float": 0}`), float32(0), func() interface{} { return data1.TFloat }, func() interface{} { return data2.TFloat }},
		{"t_float 3", []byte(`{"t_float": 1109.1}`), float32(1109.1), func() interface{} { return data1.TFloat }, func() interface{} { return data2.TFloat }},
		{"t_float 4", []byte(`{}`), float32(1109.1), func() interface{} { return data1.TFloat }, func() interface{} { return data2.TFloat }},

		{"t_double 1", []byte(`{"t_double": 1110.1}`), float64(1110.1), func() interface{} { return data1.TDouble }, func() interface{} { return data2.TDouble }},
		{"t_double 2", []byte(`{"t_double": 0}`), float64(0), func() interface{} { return data1.TDouble }, func() interface{} { return data2.TDouble }},
		{"t_double 3", []byte(`{"t_double": 1110.1}`), float64(1110.1), func() interface{} { return data1.TDouble }, func() interface{} { return data2.TDouble }},
		{"t_double 4", []byte(`{}`), float64(1110.1), func() interface{} { return data1.TDouble }, func() interface{} { return data2.TDouble }},

		{"t_bool 1", []byte(`{"t_bool": true}`), true, func() interface{} { return data1.TBool }, func() interface{} { return data2.TBool }},
		{"t_bool 2", []byte(`{"t_bool": false}`), false, func() interface{} { return data1.TBool }, func() interface{} { return data2.TBool }},
		{"t_bool 3", []byte(`{"t_bool": true}`), true, func() interface{} { return data1.TBool }, func() interface{} { return data2.TBool }},
		{"t_bool 4", []byte(`{}`), true, func() interface{} { return data1.TBool }, func() interface{} { return data2.TBool }},

		{"t_bytes 1", []byte(`{}`), []byte(nil), func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},
		{"t_bytes 2", []byte(`{"t_bytes": null}`), []byte(nil), func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},
		{"t_bytes 3", []byte(`{"t_bytes": ""}`), []byte(""), func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},
		{"t_bytes 4", []byte(fmt.Sprintf(`{"t_bytes": "%s"}`, bb1)), b1, func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},
		{"t_bytes 5", []byte(fmt.Sprintf(`{"t_bytes": "%s"}`, bb2)), b2, func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},
		{"t_bytes 6", []byte(fmt.Sprintf(`{"t_bytes": "%s"}`, bb1)), b1, func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},
		{"t_bytes 7", []byte(`{}`), b1, func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},
		{"t_bytes 8", []byte(`{"t_bytes": null}`), []byte(nil), func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},
		{"t_bytes 9", []byte(`{"t_bytes": ""}`), []byte(""), func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},
		{"t_bytes 10", []byte(fmt.Sprintf(`{"t_bytes": "%s"}`, bb1)), b1, func() interface{} { return data1.TBytes }, func() interface{} { return data2.TBytes }},

		{"t_enum1 1", []byte(`{"t_enum1": 1}`), gojsontest.UnmarshalData_stopped, func() interface{} { return data1.TEnum1 }, func() interface{} { return data2.TEnum1 }},
		{"t_enum1 2", []byte(`{"t_enum1": 0}`), gojsontest.UnmarshalData_running, func() interface{} { return data1.TEnum1 }, func() interface{} { return data2.TEnum1 }},
		{"t_enum1 3", []byte(`{"t_enum1": 1}`), gojsontest.UnmarshalData_stopped, func() interface{} { return data1.TEnum1 }, func() interface{} { return data2.TEnum1 }},
		{"t_enum1 4", []byte(`{}`), gojsontest.UnmarshalData_stopped, func() interface{} { return data1.TEnum1 }, func() interface{} { return data2.TEnum1 }},

		{"t_aliases 1", []byte(`{}`), (*gojsontest.UnmarshalData_Aliases)(nil), func() interface{} { return data1.TAliases }, func() interface{} { return data2.TAliases }},
		{"t_aliases 2", []byte(`{"t_aliases": null}`), (*gojsontest.UnmarshalData_Aliases)(nil), func() interface{} { return data1.TAliases }, func() interface{} { return data2.TAliases }},
		{"t_aliases 3", []byte(`{"t_aliases": {"k1": "v1"}}`), &gojsontest.UnmarshalData_Aliases{}, func() interface{} { return data1.TAliases }, func() interface{} { return data2.TAliases }},
		{"t_aliases 4", []byte(`{"t_aliases": {}}`), &gojsontest.UnmarshalData_Aliases{}, func() interface{} { return data1.TAliases }, func() interface{} { return data2.TAliases }},
		{"t_aliases 5", []byte(`{}`), &gojsontest.UnmarshalData_Aliases{}, func() interface{} { return data1.TAliases }, func() interface{} { return data2.TAliases }},
		{"t_aliases 6", []byte(`{"t_aliases": null}`), (*gojsontest.UnmarshalData_Aliases)(nil), func() interface{} { return data1.TAliases }, func() interface{} { return data2.TAliases }},
		{"t_aliases 7", []byte(`{}`), (*gojsontest.UnmarshalData_Aliases)(nil), func() interface{} { return data1.TAliases }, func() interface{} { return data2.TAliases }},
		{"t_aliases 8", []byte(`{"t_aliases": {}}`), &gojsontest.UnmarshalData_Aliases{}, func() interface{} { return data1.TAliases }, func() interface{} { return data2.TAliases }},

		{"t_config 1", []byte(`{}`), (*gojsontest.UnmarshalData_Config)(nil), func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
		{"t_config 2", []byte(`{"t_config": null}`), (*gojsontest.UnmarshalData_Config)(nil), func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
		{"t_config 3", []byte(`{"t_config": {"ip": "192.168.1.1", "port": 8080}}`), &gojsontest.UnmarshalData_Config{Ip: "192.168.1.1", Port: 8080}, func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
		{"t_config 4", []byte(`{"t_config": null}`), (*gojsontest.UnmarshalData_Config)(nil), func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
		{"t_config 5", []byte(`{"t_config": {}}`), &gojsontest.UnmarshalData_Config{}, func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
		{"t_config 6", []byte(`{"t_config": {"ip": "192.168.1.1", "port": 8080}}`), &gojsontest.UnmarshalData_Config{Ip: "192.168.1.1", Port: 8080}, func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
		{"t_config 7", []byte(`{}`), &gojsontest.UnmarshalData_Config{Ip: "192.168.1.1", Port: 8080}, func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
		{"t_config 8", []byte(`{"t_config": {"ip": "172.10.1.1"}}`), &gojsontest.UnmarshalData_Config{Ip: "172.10.1.1", Port: 8080}, func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
		{"t_config 9", []byte(`{"t_config": {"xxx": "172.10.1.1"}}`), &gojsontest.UnmarshalData_Config{Ip: "172.10.1.1", Port: 8080}, func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
		{"t_config 10", []byte(`{"t_config": {"ip": "", "port": 0}}`), &gojsontest.UnmarshalData_Config{Ip: "", Port: 0}, func() interface{} { return data1.TConfig }, func() interface{} { return data2.TConfig }},
	}

	arrayCases := []*CaseDesc{
		{"array_double 1", []byte(`{"array_double": null}`), []float64(nil), func() interface{} { return data1.ArrayDouble }, func() interface{} { return data2.ArrayDouble }},
		{"array_double 2", []byte(`{"array_double": []}`), []float64{}, func() interface{} { return data1.ArrayDouble }, func() interface{} { return data2.ArrayDouble }},
		{"array_double 3", []byte(`{"array_double": [1.1, 1.2, 1.3]}`), []float64{1.1, 1.2, 1.3}, func() interface{} { return data1.ArrayDouble }, func() interface{} { return data2.ArrayDouble }},
		{"array_double 4", []byte(`{"array_double": [2.1, 2.2, 2.3]}`), []float64{2.1, 2.2, 2.3}, func() interface{} { return data1.ArrayDouble }, func() interface{} { return data2.ArrayDouble }},
		{"array_double 5", []byte(`{"array_double": [3.3]}`), []float64{3.3}, func() interface{} { return data1.ArrayDouble }, func() interface{} { return data2.ArrayDouble }},
		{"array_double 6", []byte(`{"array_double": [3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 3.10]}`), []float64{3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 3.10}, func() interface{} { return data1.ArrayDouble }, func() interface{} { return data2.ArrayDouble }},
		{"array_double 7", []byte(`{"array_double": []}`), []float64{}, func() interface{} { return data1.ArrayDouble }, func() interface{} { return data2.ArrayDouble }},
		{"array_double 8", []byte(`{"array_double": null}`), []float64(nil), func() interface{} { return data1.ArrayDouble }, func() interface{} { return data2.ArrayDouble }},

		{"array_float 1", []byte(`{"array_float": null}`), []float32(nil), func() interface{} { return data1.ArrayFloat }, func() interface{} { return data2.ArrayFloat }},
		{"array_float 2", []byte(`{"array_float": []}`), []float32{}, func() interface{} { return data1.ArrayFloat }, func() interface{} { return data2.ArrayFloat }},
		{"array_float 3", []byte(`{"array_float": [1.1, 1.2, 1.3]}`), []float32{1.1, 1.2, 1.3}, func() interface{} { return data1.ArrayFloat }, func() interface{} { return data2.ArrayFloat }},
		{"array_float 4", []byte(`{"array_float": [2.1, 2.2, 2.3]}`), []float32{2.1, 2.2, 2.3}, func() interface{} { return data1.ArrayFloat }, func() interface{} { return data2.ArrayFloat }},
		{"array_float 5", []byte(`{"array_float": [3.3]}`), []float32{3.3}, func() interface{} { return data1.ArrayFloat }, func() interface{} { return data2.ArrayFloat }},
		{"array_float 6", []byte(`{"array_float": [3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 3.10]}`), []float32{3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 3.10}, func() interface{} { return data1.ArrayFloat }, func() interface{} { return data2.ArrayFloat }},
		{"array_float 7", []byte(`{"array_float": []}`), []float32{}, func() interface{} { return data1.ArrayFloat }, func() interface{} { return data2.ArrayFloat }},
		{"array_float 8", []byte(`{"array_float": null}`), []float32(nil), func() interface{} { return data1.ArrayFloat }, func() interface{} { return data2.ArrayFloat }},

		{"array_int32 1", []byte(`{"array_int32": null}`), []int32(nil), func() interface{} { return data1.ArrayInt32 }, func() interface{} { return data2.ArrayInt32 }},
		{"array_int32 2", []byte(`{"array_int32": []}`), []int32{}, func() interface{} { return data1.ArrayInt32 }, func() interface{} { return data2.ArrayInt32 }},
		{"array_int32 3", []byte(`{"array_int32": [11, 12, 13]}`), []int32{11, 12, 13}, func() interface{} { return data1.ArrayInt32 }, func() interface{} { return data2.ArrayInt32 }},
		{"array_int32 4", []byte(`{"array_int32": [21, 22, 23]}`), []int32{21, 22, 23}, func() interface{} { return data1.ArrayInt32 }, func() interface{} { return data2.ArrayInt32 }},
		{"array_int32 5", []byte(`{"array_int32": [33]}`), []int32{33}, func() interface{} { return data1.ArrayInt32 }, func() interface{} { return data2.ArrayInt32 }},
		{"array_int32 6", []byte(`{"array_int32": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArrayInt32 }, func() interface{} { return data2.ArrayInt32 }},
		{"array_int32 7", []byte(`{"array_int32": []}`), []int32{}, func() interface{} { return data1.ArrayInt32 }, func() interface{} { return data2.ArrayInt32 }},
		{"array_int32 8", []byte(`{"array_int32": null}`), []int32(nil), func() interface{} { return data1.ArrayInt32 }, func() interface{} { return data2.ArrayInt32 }},

		{"array_int64 1", []byte(`{"array_int64": null}`), []int64(nil), func() interface{} { return data1.ArrayInt64 }, func() interface{} { return data2.ArrayInt64 }},
		{"array_int64 2", []byte(`{"array_int64": []}`), []int64{}, func() interface{} { return data1.ArrayInt64 }, func() interface{} { return data2.ArrayInt64 }},
		{"array_int64 3", []byte(`{"array_int64": [11, 12, 13]}`), []int64{11, 12, 13}, func() interface{} { return data1.ArrayInt64 }, func() interface{} { return data2.ArrayInt64 }},
		{"array_int64 4", []byte(`{"array_int64": [21, 22, 23]}`), []int64{21, 22, 23}, func() interface{} { return data1.ArrayInt64 }, func() interface{} { return data2.ArrayInt64 }},
		{"array_int64 5", []byte(`{"array_int64": [33]}`), []int64{33}, func() interface{} { return data1.ArrayInt64 }, func() interface{} { return data2.ArrayInt64 }},
		{"array_int64 6", []byte(`{"array_int64": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArrayInt64 }, func() interface{} { return data2.ArrayInt64 }},
		{"array_int64 7", []byte(`{"array_int64": []}`), []int64{}, func() interface{} { return data1.ArrayInt64 }, func() interface{} { return data2.ArrayInt64 }},
		{"array_int64 8", []byte(`{"array_int64": null}`), []int64(nil), func() interface{} { return data1.ArrayInt64 }, func() interface{} { return data2.ArrayInt64 }},

		{"array_uint32 1", []byte(`{"array_uint32": null}`), []uint32(nil), func() interface{} { return data1.ArrayUint32 }, func() interface{} { return data2.ArrayUint32 }},
		{"array_uint32 2", []byte(`{"array_uint32": []}`), []uint32{}, func() interface{} { return data1.ArrayUint32 }, func() interface{} { return data2.ArrayUint32 }},
		{"array_uint32 3", []byte(`{"array_uint32": [11, 12, 13]}`), []uint32{11, 12, 13}, func() interface{} { return data1.ArrayUint32 }, func() interface{} { return data2.ArrayUint32 }},
		{"array_uint32 4", []byte(`{"array_uint32": [21, 22, 23]}`), []uint32{21, 22, 23}, func() interface{} { return data1.ArrayUint32 }, func() interface{} { return data2.ArrayUint32 }},
		{"array_uint32 5", []byte(`{"array_uint32": [33]}`), []uint32{33}, func() interface{} { return data1.ArrayUint32 }, func() interface{} { return data2.ArrayUint32 }},
		{"array_uint32 6", []byte(`{"array_uint32": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArrayUint32 }, func() interface{} { return data2.ArrayUint32 }},
		{"array_uint32 7", []byte(`{"array_uint32": []}`), []uint32{}, func() interface{} { return data1.ArrayUint32 }, func() interface{} { return data2.ArrayUint32 }},
		{"array_uint32 8", []byte(`{"array_uint32": null}`), []uint32(nil), func() interface{} { return data1.ArrayUint32 }, func() interface{} { return data2.ArrayUint32 }},

		{"array_uint64 1", []byte(`{"array_uint64": null}`), []uint64(nil), func() interface{} { return data1.ArrayUint64 }, func() interface{} { return data2.ArrayUint64 }},
		{"array_uint64 2", []byte(`{"array_uint64": []}`), []uint64{}, func() interface{} { return data1.ArrayUint64 }, func() interface{} { return data2.ArrayUint64 }},
		{"array_uint64 3", []byte(`{"array_uint64": [11, 12, 13]}`), []uint64{11, 12, 13}, func() interface{} { return data1.ArrayUint64 }, func() interface{} { return data2.ArrayUint64 }},
		{"array_uint64 4", []byte(`{"array_uint64": [21, 22, 23]}`), []uint64{21, 22, 23}, func() interface{} { return data1.ArrayUint64 }, func() interface{} { return data2.ArrayUint64 }},
		{"array_uint64 5", []byte(`{"array_uint64": [33]}`), []uint64{33}, func() interface{} { return data1.ArrayUint64 }, func() interface{} { return data2.ArrayUint64 }},
		{"array_uint64 6", []byte(`{"array_uint64": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArrayUint64 }, func() interface{} { return data2.ArrayUint64 }},
		{"array_uint64 7", []byte(`{"array_uint64": []}`), []uint64{}, func() interface{} { return data1.ArrayUint64 }, func() interface{} { return data2.ArrayUint64 }},
		{"array_uint64 8", []byte(`{"array_uint64": null}`), []uint64(nil), func() interface{} { return data1.ArrayUint64 }, func() interface{} { return data2.ArrayUint64 }},

		{"array_sint32 1", []byte(`{"array_sint32": null}`), []int32(nil), func() interface{} { return data1.ArraySint32 }, func() interface{} { return data2.ArraySint32 }},
		{"array_sint32 2", []byte(`{"array_sint32": []}`), []int32{}, func() interface{} { return data1.ArraySint32 }, func() interface{} { return data2.ArraySint32 }},
		{"array_sint32 3", []byte(`{"array_sint32": [11, 12, 13]}`), []int32{11, 12, 13}, func() interface{} { return data1.ArraySint32 }, func() interface{} { return data2.ArraySint32 }},
		{"array_sint32 4", []byte(`{"array_sint32": [21, 22, 23]}`), []int32{21, 22, 23}, func() interface{} { return data1.ArraySint32 }, func() interface{} { return data2.ArraySint32 }},
		{"array_sint32 5", []byte(`{"array_sint32": [33]}`), []int32{33}, func() interface{} { return data1.ArraySint32 }, func() interface{} { return data2.ArraySint32 }},
		{"array_sint32 6", []byte(`{"array_sint32": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArraySint32 }, func() interface{} { return data2.ArraySint32 }},
		{"array_sint32 7", []byte(`{"array_sint32": []}`), []int32{}, func() interface{} { return data1.ArraySint32 }, func() interface{} { return data2.ArraySint32 }},
		{"array_sint32 8", []byte(`{"array_sint32": null}`), []int32(nil), func() interface{} { return data1.ArraySint32 }, func() interface{} { return data2.ArraySint32 }},

		{"array_sint64 1", []byte(`{"array_sint64": null}`), []int64(nil), func() interface{} { return data1.ArraySint64 }, func() interface{} { return data2.ArraySint64 }},
		{"array_sint64 2", []byte(`{"array_sint64": []}`), []int64{}, func() interface{} { return data1.ArraySint64 }, func() interface{} { return data2.ArraySint64 }},
		{"array_sint64 3", []byte(`{"array_sint64": [11, 12, 13]}`), []int64{11, 12, 13}, func() interface{} { return data1.ArraySint64 }, func() interface{} { return data2.ArraySint64 }},
		{"array_sint64 4", []byte(`{"array_sint64": [21, 22, 23]}`), []int64{21, 22, 23}, func() interface{} { return data1.ArraySint64 }, func() interface{} { return data2.ArraySint64 }},
		{"array_sint64 5", []byte(`{"array_sint64": [33]}`), []int64{33}, func() interface{} { return data1.ArraySint64 }, func() interface{} { return data2.ArraySint64 }},
		{"array_sint64 6", []byte(`{"array_sint64": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArraySint64 }, func() interface{} { return data2.ArraySint64 }},
		{"array_sint64 7", []byte(`{"array_sint64": []}`), []int64{}, func() interface{} { return data1.ArraySint64 }, func() interface{} { return data2.ArraySint64 }},
		{"array_sint64 8", []byte(`{"array_sint64": null}`), []int64(nil), func() interface{} { return data1.ArraySint64 }, func() interface{} { return data2.ArraySint64 }},

		{"array_sfixed32 1", []byte(`{"array_sfixed32": null}`), []int32(nil), func() interface{} { return data1.ArraySfixed32 }, func() interface{} { return data2.ArraySfixed32 }},
		{"array_sfixed32 2", []byte(`{"array_sfixed32": []}`), []int32{}, func() interface{} { return data1.ArraySfixed32 }, func() interface{} { return data2.ArraySfixed32 }},
		{"array_sfixed32 3", []byte(`{"array_sfixed32": [11, 12, 13]}`), []int32{11, 12, 13}, func() interface{} { return data1.ArraySfixed32 }, func() interface{} { return data2.ArraySfixed32 }},
		{"array_sfixed32 4", []byte(`{"array_sfixed32": [21, 22, 23]}`), []int32{21, 22, 23}, func() interface{} { return data1.ArraySfixed32 }, func() interface{} { return data2.ArraySfixed32 }},
		{"array_sfixed32 5", []byte(`{"array_sfixed32": [33]}`), []int32{33}, func() interface{} { return data1.ArraySfixed32 }, func() interface{} { return data2.ArraySfixed32 }},
		{"array_sfixed32 6", []byte(`{"array_sfixed32": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArraySfixed32 }, func() interface{} { return data2.ArraySfixed32 }},
		{"array_sfixed32 7", []byte(`{"array_sfixed32": []}`), []int32{}, func() interface{} { return data1.ArraySfixed32 }, func() interface{} { return data2.ArraySfixed32 }},
		{"array_sfixed32 8", []byte(`{"array_sfixed32": null}`), []int32(nil), func() interface{} { return data1.ArraySfixed32 }, func() interface{} { return data2.ArraySfixed32 }},

		{"array_sfixed64 1", []byte(`{"array_sfixed64": null}`), []int64(nil), func() interface{} { return data1.ArraySfixed64 }, func() interface{} { return data2.ArraySfixed64 }},
		{"array_sfixed64 2", []byte(`{"array_sfixed64": []}`), []int64{}, func() interface{} { return data1.ArraySfixed64 }, func() interface{} { return data2.ArraySfixed64 }},
		{"array_sfixed64 3", []byte(`{"array_sfixed64": [11, 12, 13]}`), []int64{11, 12, 13}, func() interface{} { return data1.ArraySfixed64 }, func() interface{} { return data2.ArraySfixed64 }},
		{"array_sfixed64 4", []byte(`{"array_sfixed64": [21, 22, 23]}`), []int64{21, 22, 23}, func() interface{} { return data1.ArraySfixed64 }, func() interface{} { return data2.ArraySfixed64 }},
		{"array_sfixed64 5", []byte(`{"array_sfixed64": [33]}`), []int64{33}, func() interface{} { return data1.ArraySfixed64 }, func() interface{} { return data2.ArraySfixed64 }},
		{"array_sfixed64 6", []byte(`{"array_sfixed64": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArraySfixed64 }, func() interface{} { return data2.ArraySfixed64 }},
		{"array_sfixed64 7", []byte(`{"array_sfixed64": []}`), []int64{}, func() interface{} { return data1.ArraySfixed64 }, func() interface{} { return data2.ArraySfixed64 }},
		{"array_sfixed64 8", []byte(`{"array_sfixed64": null}`), []int64(nil), func() interface{} { return data1.ArraySfixed64 }, func() interface{} { return data2.ArraySfixed64 }},

		{"array_fixed32 1", []byte(`{"array_fixed32": null}`), []uint32(nil), func() interface{} { return data1.ArrayFixed32 }, func() interface{} { return data2.ArrayFixed32 }},
		{"array_fixed32 2", []byte(`{"array_fixed32": []}`), []uint32{}, func() interface{} { return data1.ArrayFixed32 }, func() interface{} { return data2.ArrayFixed32 }},
		{"array_fixed32 3", []byte(`{"array_fixed32": [11, 12, 13]}`), []uint32{11, 12, 13}, func() interface{} { return data1.ArrayFixed32 }, func() interface{} { return data2.ArrayFixed32 }},
		{"array_fixed32 4", []byte(`{"array_fixed32": [21, 22, 23]}`), []uint32{21, 22, 23}, func() interface{} { return data1.ArrayFixed32 }, func() interface{} { return data2.ArrayFixed32 }},
		{"array_fixed32 5", []byte(`{"array_fixed32": [33]}`), []uint32{33}, func() interface{} { return data1.ArrayFixed32 }, func() interface{} { return data2.ArrayFixed32 }},
		{"array_fixed32 6", []byte(`{"array_fixed32": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArrayFixed32 }, func() interface{} { return data2.ArrayFixed32 }},
		{"array_fixed32 7", []byte(`{"array_fixed32": []}`), []uint32{}, func() interface{} { return data1.ArrayFixed32 }, func() interface{} { return data2.ArrayFixed32 }},
		{"array_fixed32 8", []byte(`{"array_fixed32": null}`), []uint32(nil), func() interface{} { return data1.ArrayFixed32 }, func() interface{} { return data2.ArrayFixed32 }},

		{"array_fixed64 1", []byte(`{"array_fixed64": null}`), []uint64(nil), func() interface{} { return data1.ArrayFixed64 }, func() interface{} { return data2.ArrayFixed64 }},
		{"array_fixed64 2", []byte(`{"array_fixed64": []}`), []uint64{}, func() interface{} { return data1.ArrayFixed64 }, func() interface{} { return data2.ArrayFixed64 }},
		{"array_fixed64 3", []byte(`{"array_fixed64": [11, 12, 13]}`), []uint64{11, 12, 13}, func() interface{} { return data1.ArrayFixed64 }, func() interface{} { return data2.ArrayFixed64 }},
		{"array_fixed64 4", []byte(`{"array_fixed64": [21, 22, 23]}`), []uint64{21, 22, 23}, func() interface{} { return data1.ArrayFixed64 }, func() interface{} { return data2.ArrayFixed64 }},
		{"array_fixed64 5", []byte(`{"array_fixed64": [33]}`), []uint64{33}, func() interface{} { return data1.ArrayFixed64 }, func() interface{} { return data2.ArrayFixed64 }},
		{"array_fixed64 6", []byte(`{"array_fixed64": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}`), []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func() interface{} { return data1.ArrayFixed64 }, func() interface{} { return data2.ArrayFixed64 }},
		{"array_fixed64 7", []byte(`{"array_fixed64": []}`), []uint64{}, func() interface{} { return data1.ArrayFixed64 }, func() interface{} { return data2.ArrayFixed64 }},
		{"array_fixed64 8", []byte(`{"array_fixed64": null}`), []uint64(nil), func() interface{} { return data1.ArrayFixed64 }, func() interface{} { return data2.ArrayFixed64 }},

		{"array_bool 1", []byte(`{"array_bool": null}`), []bool(nil), func() interface{} { return data1.ArrayBool }, func() interface{} { return data2.ArrayBool }},
		{"array_bool 2", []byte(`{"array_bool": []}`), []bool{}, func() interface{} { return data1.ArrayBool }, func() interface{} { return data2.ArrayBool }},
		{"array_bool 3", []byte(`{"array_bool": [true, false, true]}`), []bool{true, false, true}, func() interface{} { return data1.ArrayBool }, func() interface{} { return data2.ArrayBool }},
		{"array_bool 4", []byte(`{"array_bool": [false, true, false]}`), []bool{false, true, false}, func() interface{} { return data1.ArrayBool }, func() interface{} { return data2.ArrayBool }},
		{"array_bool 5", []byte(`{"array_bool": [true]}`), []bool{true}, func() interface{} { return data1.ArrayBool }, func() interface{} { return data2.ArrayBool }},
		{"array_bool 6", []byte(`{"array_bool": [true, true, true, false, true, true, true, true, false, true]}`), []bool{true, true, true, false, true, true, true, true, false, true}, func() interface{} { return data1.ArrayBool }, func() interface{} { return data2.ArrayBool }},
		{"array_bool 7", []byte(`{"array_bool": []}`), []bool{}, func() interface{} { return data1.ArrayBool }, func() interface{} { return data2.ArrayBool }},
		{"array_bool 8", []byte(`{"array_bool": null}`), []bool(nil), func() interface{} { return data1.ArrayBool }, func() interface{} { return data2.ArrayBool }},

		{"array_string 1", []byte(`{"array_string": null}`), []string(nil), func() interface{} { return data1.ArrayString }, func() interface{} { return data2.ArrayString }},
		{"array_string 2", []byte(`{"array_string": []}`), []string{}, func() interface{} { return data1.ArrayString }, func() interface{} { return data2.ArrayString }},
		{"array_string 3", []byte(`{"array_string": ["s1", "s2", null, "s3", ""]}`), []string{"s1", "s2", "", "s3", ""}, func() interface{} { return data1.ArrayString }, func() interface{} { return data2.ArrayString }},
		// {"array_string 4", []byte(`{"array_string": ["", null, "s4"]}`), []string{"", "s2", "s4"}, func() interface{} { return data1.ArrayString }, func() interface{} { return data2.ArrayString }},
		{"array_string 5", []byte(`{"array_string": ["", "s6", "s4"]}`), []string{"", "s6", "s4"}, func() interface{} { return data1.ArrayString }, func() interface{} { return data2.ArrayString }},
		{"array_string 6", []byte(`{"array_string": ["", "s1"]}`), []string{"", "s1"}, func() interface{} { return data1.ArrayString }, func() interface{} { return data2.ArrayString }},
		{"array_string 7", []byte(`{"array_string": ["s1", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", ""]}`), []string{"s1", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", ""}, func() interface{} { return data1.ArrayString }, func() interface{} { return data2.ArrayString }},
		{"array_string 8", []byte(`{"array_string": []}`), []string{}, func() interface{} { return data1.ArrayString }, func() interface{} { return data2.ArrayString }},
		{"array_string 9", []byte(`{"array_string": null}`), []string(nil), func() interface{} { return data1.ArrayString }, func() interface{} { return data2.ArrayString }},

		{"array_bytes 1", []byte(`{"array_bytes": null}`), [][]uint8(nil), func() interface{} { return data1.ArrayBytes }, func() interface{} { return data2.ArrayBytes }},
		{"array_bytes 2", []byte(`{"array_bytes": []}`), [][]byte{}, func() interface{} { return data1.ArrayBytes }, func() interface{} { return data2.ArrayBytes }},
		{"array_bytes 3", []byte(fmt.Sprintf(`{"array_bytes": ["%s", null, "%s", ""]}`, bb1, bb2)), [][]byte{b1, nil, b2, []byte("")}, func() interface{} { return data1.ArrayBytes }, func() interface{} { return data2.ArrayBytes }},
		{"array_bytes 4", []byte(fmt.Sprintf(`{"array_bytes": [null, "%s","%s", ""]}`, bb1, bb2)), [][]byte{nil, b1, b2, []byte("")}, func() interface{} { return data1.ArrayBytes }, func() interface{} { return data2.ArrayBytes }},
		{"array_bytes 5", []byte(fmt.Sprintf(`{"array_bytes": ["", "%s", null]}`, bb1)), [][]byte{[]byte(""), b1, nil}, func() interface{} { return data1.ArrayBytes }, func() interface{} { return data2.ArrayBytes }},
		{"array_bytes 6", []byte(fmt.Sprintf(`{"array_bytes": [null, "%s"]}`, bb1)), [][]byte{nil, b1}, func() interface{} { return data1.ArrayBytes }, func() interface{} { return data2.ArrayBytes }},
		{"array_bytes 7", []byte(fmt.Sprintf(`{"array_bytes": ["%s", "%s", null, "%s", "%s", null, "%s", "%s", null, null]}`, bb1, bb2, bb2, bb1, bb1, bb2)), [][]byte{b1, b2, nil, b2, b1, nil, b1, b2, nil, nil}, func() interface{} { return data1.ArrayBytes }, func() interface{} { return data2.ArrayBytes }},
		{"array_bytes 8", []byte(`{"array_bytes": []}`), [][]byte{}, func() interface{} { return data1.ArrayBytes }, func() interface{} { return data2.ArrayBytes }},
		{"array_bytes 9", []byte(`{"array_bytes": null}`), [][]uint8(nil), func() interface{} { return data1.ArrayBytes }, func() interface{} { return data2.ArrayBytes }},

		{"array_enum1 1", []byte(`{"array_enum1": null}`), []gojsontest.UnmarshalData_Enum(nil), func() interface{} { return data1.ArrayEnum1 }, func() interface{} { return data2.ArrayEnum1 }},
		{"array_enum1 2", []byte(`{"array_enum1": []}`), []gojsontest.UnmarshalData_Enum{}, func() interface{} { return data1.ArrayEnum1 }, func() interface{} { return data2.ArrayEnum1 }},
		{"array_enum1 3", []byte(`{"array_enum1": [1, 0, 1]}`), []gojsontest.UnmarshalData_Enum{gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_running, gojsontest.UnmarshalData_stopped}, func() interface{} { return data1.ArrayEnum1 }, func() interface{} { return data2.ArrayEnum1 }},
		{"array_enum1 4", []byte(`{"array_enum1": [0, 1, 0]}`), []gojsontest.UnmarshalData_Enum{gojsontest.UnmarshalData_running, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_running}, func() interface{} { return data1.ArrayEnum1 }, func() interface{} { return data2.ArrayEnum1 }},
		{"array_enum1 5", []byte(`{"array_enum1": [1]}`), []gojsontest.UnmarshalData_Enum{gojsontest.UnmarshalData_stopped}, func() interface{} { return data1.ArrayEnum1 }, func() interface{} { return data2.ArrayEnum1 }},
		{"array_enum1 5", []byte(`{"array_enum1": [1,1,1,1,1,1,1,1,1,1]}`), []gojsontest.UnmarshalData_Enum{gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped}, func() interface{} { return data1.ArrayEnum1 }, func() interface{} { return data2.ArrayEnum1 }},
		{"array_enum1 7", []byte(`{"array_enum1": []}`), []gojsontest.UnmarshalData_Enum{}, func() interface{} { return data1.ArrayEnum1 }, func() interface{} { return data2.ArrayEnum1 }},
		{"array_enum1 8", []byte(`{"array_enum1": null}`), []gojsontest.UnmarshalData_Enum(nil), func() interface{} { return data1.ArrayEnum1 }, func() interface{} { return data2.ArrayEnum1 }},

		{"array_aliases 1", []byte(`{"array_aliases": null}`), []*gojsontest.UnmarshalData_Aliases(nil), func() interface{} { return data1.ArrayAliases }, func() interface{} { return data2.ArrayAliases }},
		{"array_aliases 2", []byte(`{"array_aliases": []}`), []*gojsontest.UnmarshalData_Aliases{}, func() interface{} { return data1.ArrayAliases }, func() interface{} { return data2.ArrayAliases }},
		{"array_aliases 3", []byte(`{"array_aliases": null}`), []*gojsontest.UnmarshalData_Aliases(nil), func() interface{} { return data1.ArrayAliases }, func() interface{} { return data2.ArrayAliases }},

		{"array_config 1", []byte(`{"array_config": null}`), []*gojsontest.UnmarshalData_Config(nil), func() interface{} { return data1.ArrayConfig }, func() interface{} { return data2.ArrayConfig }},
		{"array_config 2", []byte(`{"array_config": []}`), []*gojsontest.UnmarshalData_Config{}, func() interface{} { return data1.ArrayConfig }, func() interface{} { return data2.ArrayConfig }},
		{"array_config 3", []byte(`{"array_config": [{"ip": "192", "port": 8080}, {"ip": "193", "port": 8081}, null, {"ip": "194", "port": 8082}]}`), []*gojsontest.UnmarshalData_Config{{Ip: "192", Port: 8080}, {Ip: "193", Port: 8081}, nil, {Ip: "194", Port: 8082}}, func() interface{} { return data1.ArrayConfig }, func() interface{} { return data2.ArrayConfig }},
		{"array_config 4", []byte(`{"array_config": [{"ip": "10.10", "port": 9090}, null, {"ip": "10.20", "port": 9091}, {"ip": "10.30", "port": 9092}]}`), []*gojsontest.UnmarshalData_Config{{Ip: "10.10", Port: 9090}, nil, {Ip: "10.20", Port: 9091}, {Ip: "10.30", Port: 9092}}, func() interface{} { return data1.ArrayConfig }, func() interface{} { return data2.ArrayConfig }},
		{"array_config 5", []byte(`{"array_config": [{"ip": "10.10", "port": 9090}]}`), []*gojsontest.UnmarshalData_Config{{Ip: "10.10", Port: 9090}}, func() interface{} { return data1.ArrayConfig }, func() interface{} { return data2.ArrayConfig }},
		{"array_config 6", []byte(`{"array_config": [{"ip": "10.10", "port": 9090}, {"ip": "10.20", "port": 9091}, {"ip": "10.30", "port": 9092},{"ip": "10.40", "port": 9093}, {"ip": "10.50", "port": 9094}, {"ip": "10.60", "port": 9095}, {"ip": "10.70", "port": 9096}, {"ip": "10.80", "port": 9097}, {"ip": "10.90", "port": 9098}]}`), []*gojsontest.UnmarshalData_Config{{Ip: "10.10", Port: 9090}, {Ip: "10.20", Port: 9091}, {Ip: "10.30", Port: 9092}, {Ip: "10.40", Port: 9093}, {Ip: "10.50", Port: 9094}, {Ip: "10.60", Port: 9095}, {Ip: "10.70", Port: 9096}, {Ip: "10.80", Port: 9097}, {Ip: "10.90", Port: 9098}}, func() interface{} { return data1.ArrayConfig }, func() interface{} { return data2.ArrayConfig }},
		{"array_config 7", []byte(`{"array_config": []}`), []*gojsontest.UnmarshalData_Config{}, func() interface{} { return data1.ArrayConfig }, func() interface{} { return data2.ArrayConfig }},
		{"array_config 8", []byte(`{"array_config": null}`), []*gojsontest.UnmarshalData_Config(nil), func() interface{} { return data1.ArrayConfig }, func() interface{} { return data2.ArrayConfig }},
		{"array_config 9", []byte(`{"array_config": []}`), []*gojsontest.UnmarshalData_Config{}, func() interface{} { return data1.ArrayConfig }, func() interface{} { return data2.ArrayConfig }},
	}

	mapCases := []*CaseDesc{
		{"map_int32_int32 1", []byte(`{"map_int32_int32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapInt32Int32 }, func() interface{} { return data2.MapInt32Int32 }},
		{"map_int32_int32 2", []byte(`{"map_int32_int32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapInt32Int32 }, func() interface{} { return data2.MapInt32Int32 }},
		{"map_int32_int32 3", []byte(`{"map_int32_int32": {"1":1, "2": 2, "3": 3}}`), map[int32]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Int32 }, func() interface{} { return data2.MapInt32Int32 }},
		{"map_int32_int32 4", []byte(`{"map_int32_int32": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]int32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Int32 }, func() interface{} { return data2.MapInt32Int32 }},
		{"map_int32_int32 5", []byte(`{"map_int32_int32": {"8":8}}`), map[int32]int32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Int32 }, func() interface{} { return data2.MapInt32Int32 }},
		{"map_int32_int32 6", []byte(`{"map_int32_int32": {}}`), map[int32]int32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Int32 }, func() interface{} { return data2.MapInt32Int32 }},
		{"map_int32_int32 7", []byte(`{"map_int32_int32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapInt32Int32 }, func() interface{} { return data2.MapInt32Int32 }},
		{"map_int32_int32 8", []byte(`{"map_int32_int32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapInt32Int32 }, func() interface{} { return data2.MapInt32Int32 }},

		{"map_int32_int64 1", []byte(`{"map_int32_int64": null}`), map[int32]int64(nil), func() interface{} { return data1.MapInt32Int64 }, func() interface{} { return data2.MapInt32Int64 }},
		{"map_int32_int64 2", []byte(`{"map_int32_int64": {}}`), map[int32]int64{}, func() interface{} { return data1.MapInt32Int64 }, func() interface{} { return data2.MapInt32Int64 }},
		{"map_int32_int64 3", []byte(`{"map_int32_int64": {"1":1, "2": 2, "3": 3}}`), map[int32]int64{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Int64 }, func() interface{} { return data2.MapInt32Int64 }},
		{"map_int32_int64 4", []byte(`{"map_int32_int64": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]int64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Int64 }, func() interface{} { return data2.MapInt32Int64 }},
		{"map_int32_int64 5", []byte(`{"map_int32_int64": {"8":8}}`), map[int32]int64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Int64 }, func() interface{} { return data2.MapInt32Int64 }},
		{"map_int32_int64 6", []byte(`{"map_int32_int64": {}}`), map[int32]int64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Int64 }, func() interface{} { return data2.MapInt32Int64 }},
		{"map_int32_int64 7", []byte(`{"map_int32_int64": null}`), map[int32]int64(nil), func() interface{} { return data1.MapInt32Int64 }, func() interface{} { return data2.MapInt32Int64 }},
		{"map_int32_int64 8", []byte(`{"map_int32_int64": {}}`), map[int32]int64{}, func() interface{} { return data1.MapInt32Int64 }, func() interface{} { return data2.MapInt32Int64 }},

		{"map_int32_uint32 1", []byte(`{"map_int32_uint32": null}`), map[int32]uint32(nil), func() interface{} { return data1.MapInt32Uint32 }, func() interface{} { return data2.MapInt32Uint32 }},
		{"map_int32_uint32 2", []byte(`{"map_int32_uint32": {}}`), map[int32]uint32{}, func() interface{} { return data1.MapInt32Uint32 }, func() interface{} { return data2.MapInt32Uint32 }},
		{"map_int32_uint32 3", []byte(`{"map_int32_uint32": {"1":1, "2": 2, "3": 3}}`), map[int32]uint32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Uint32 }, func() interface{} { return data2.MapInt32Uint32 }},
		{"map_int32_uint32 4", []byte(`{"map_int32_uint32": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]uint32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Uint32 }, func() interface{} { return data2.MapInt32Uint32 }},
		{"map_int32_uint32 5", []byte(`{"map_int32_uint32": {"8":8}}`), map[int32]uint32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Uint32 }, func() interface{} { return data2.MapInt32Uint32 }},
		{"map_int32_uint32 6", []byte(`{"map_int32_uint32": {}}`), map[int32]uint32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Uint32 }, func() interface{} { return data2.MapInt32Uint32 }},
		{"map_int32_uint32 7", []byte(`{"map_int32_uint32": null}`), map[int32]uint32(nil), func() interface{} { return data1.MapInt32Uint32 }, func() interface{} { return data2.MapInt32Uint32 }},
		{"map_int32_uint32 8", []byte(`{"map_int32_uint32": {}}`), map[int32]uint32{}, func() interface{} { return data1.MapInt32Uint32 }, func() interface{} { return data2.MapInt32Uint32 }},

		{"map_int32_uint64 1", []byte(`{"map_int32_uint64": null}`), map[int32]uint64(nil), func() interface{} { return data1.MapInt32Uint64 }, func() interface{} { return data2.MapInt32Uint64 }},
		{"map_int32_uint64 2", []byte(`{"map_int32_uint64": {}}`), map[int32]uint64{}, func() interface{} { return data1.MapInt32Uint64 }, func() interface{} { return data2.MapInt32Uint64 }},
		{"map_int32_uint64 3", []byte(`{"map_int32_uint64": {"1":1, "2": 2, "3": 3}}`), map[int32]uint64{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Uint64 }, func() interface{} { return data2.MapInt32Uint64 }},
		{"map_int32_uint64 4", []byte(`{"map_int32_uint64": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]uint64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Uint64 }, func() interface{} { return data2.MapInt32Uint64 }},
		{"map_int32_uint64 5", []byte(`{"map_int32_uint64": {"8":8}}`), map[int32]uint64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Uint64 }, func() interface{} { return data2.MapInt32Uint64 }},
		{"map_int32_uint64 6", []byte(`{"map_int32_uint64": {}}`), map[int32]uint64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Uint64 }, func() interface{} { return data2.MapInt32Uint64 }},
		{"map_int32_uint64 7", []byte(`{"map_int32_uint64": null}`), map[int32]uint64(nil), func() interface{} { return data1.MapInt32Uint64 }, func() interface{} { return data2.MapInt32Uint64 }},
		{"map_int32_uint64 8", []byte(`{"map_int32_uint64": {}}`), map[int32]uint64{}, func() interface{} { return data1.MapInt32Uint64 }, func() interface{} { return data2.MapInt32Uint64 }},

		{"map_int32_sint32 1", []byte(`{"map_int32_sint32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapInt32Sint32 }, func() interface{} { return data2.MapInt32Sint32 }},
		{"map_int32_sint32 2", []byte(`{"map_int32_sint32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapInt32Sint32 }, func() interface{} { return data2.MapInt32Sint32 }},
		{"map_int32_sint32 3", []byte(`{"map_int32_sint32": {"1":1, "2": 2, "3": 3}}`), map[int32]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Sint32 }, func() interface{} { return data2.MapInt32Sint32 }},
		{"map_int32_sint32 4", []byte(`{"map_int32_sint32": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]int32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Sint32 }, func() interface{} { return data2.MapInt32Sint32 }},
		{"map_int32_sint32 5", []byte(`{"map_int32_sint32": {"8":8}}`), map[int32]int32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Sint32 }, func() interface{} { return data2.MapInt32Sint32 }},
		{"map_int32_sint32 6", []byte(`{"map_int32_sint32": {}}`), map[int32]int32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Sint32 }, func() interface{} { return data2.MapInt32Sint32 }},
		{"map_int32_sint32 7", []byte(`{"map_int32_sint32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapInt32Sint32 }, func() interface{} { return data2.MapInt32Sint32 }},
		{"map_int32_sint32 8", []byte(`{"map_int32_sint32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapInt32Sint32 }, func() interface{} { return data2.MapInt32Sint32 }},

		{"map_int32_sint64 1", []byte(`{"map_int32_sint64": null}`), map[int32]int64(nil), func() interface{} { return data1.MapInt32Sint64 }, func() interface{} { return data2.MapInt32Sint64 }},
		{"map_int32_sint64 2", []byte(`{"map_int32_sint64": {}}`), map[int32]int64{}, func() interface{} { return data1.MapInt32Sint64 }, func() interface{} { return data2.MapInt32Sint64 }},
		{"map_int32_sint64 3", []byte(`{"map_int32_sint64": {"1":1, "2": 2, "3": 3}}`), map[int32]int64{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Sint64 }, func() interface{} { return data2.MapInt32Sint64 }},
		{"map_int32_sint64 4", []byte(`{"map_int32_sint64": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]int64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Sint64 }, func() interface{} { return data2.MapInt32Sint64 }},
		{"map_int32_sint64 5", []byte(`{"map_int32_sint64": {"8":8}}`), map[int32]int64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Sint64 }, func() interface{} { return data2.MapInt32Sint64 }},
		{"map_int32_sint64 6", []byte(`{"map_int32_sint64": {}}`), map[int32]int64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Sint64 }, func() interface{} { return data2.MapInt32Sint64 }},
		{"map_int32_sint64 7", []byte(`{"map_int32_sint64": null}`), map[int32]int64(nil), func() interface{} { return data1.MapInt32Sint64 }, func() interface{} { return data2.MapInt32Sint64 }},
		{"map_int32_sint64 8", []byte(`{"map_int32_sint64": {}}`), map[int32]int64{}, func() interface{} { return data1.MapInt32Sint64 }, func() interface{} { return data2.MapInt32Sint64 }},

		{"map_int32_sfixed32 1", []byte(`{"map_int32_sfixed32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapInt32Sfixed32 }, func() interface{} { return data2.MapInt32Sfixed32 }},
		{"map_int32_sfixed32 2", []byte(`{"map_int32_sfixed32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapInt32Sfixed32 }, func() interface{} { return data2.MapInt32Sfixed32 }},
		{"map_int32_sfixed32 3", []byte(`{"map_int32_sfixed32": {"1":1, "2": 2, "3": 3}}`), map[int32]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Sfixed32 }, func() interface{} { return data2.MapInt32Sfixed32 }},
		{"map_int32_sfixed32 4", []byte(`{"map_int32_sfixed32": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]int32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Sfixed32 }, func() interface{} { return data2.MapInt32Sfixed32 }},
		{"map_int32_sfixed32 5", []byte(`{"map_int32_sfixed32": {"8":8}}`), map[int32]int32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Sfixed32 }, func() interface{} { return data2.MapInt32Sfixed32 }},
		{"map_int32_sfixed32 6", []byte(`{"map_int32_sfixed32": {}}`), map[int32]int32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Sfixed32 }, func() interface{} { return data2.MapInt32Sfixed32 }},
		{"map_int32_sfixed32 7", []byte(`{"map_int32_sfixed32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapInt32Sfixed32 }, func() interface{} { return data2.MapInt32Sfixed32 }},
		{"map_int32_sfixed32 8", []byte(`{"map_int32_sfixed32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapInt32Sfixed32 }, func() interface{} { return data2.MapInt32Sfixed32 }},

		{"map_int32_sfixed64 1", []byte(`{"map_int32_sfixed64": null}`), map[int32]int64(nil), func() interface{} { return data1.MapInt32Sfixed64 }, func() interface{} { return data2.MapInt32Sfixed64 }},
		{"map_int32_sfixed64 2", []byte(`{"map_int32_sfixed64": {}}`), map[int32]int64{}, func() interface{} { return data1.MapInt32Sfixed64 }, func() interface{} { return data2.MapInt32Sfixed64 }},
		{"map_int32_sfixed64 3", []byte(`{"map_int32_sfixed64": {"1":1, "2": 2, "3": 3}}`), map[int32]int64{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Sfixed64 }, func() interface{} { return data2.MapInt32Sfixed64 }},
		{"map_int32_sfixed64 4", []byte(`{"map_int32_sfixed64": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]int64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Sfixed64 }, func() interface{} { return data2.MapInt32Sfixed64 }},
		{"map_int32_sfixed64 5", []byte(`{"map_int32_sfixed64": {"8":8}}`), map[int32]int64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Sfixed64 }, func() interface{} { return data2.MapInt32Sfixed64 }},
		{"map_int32_sfixed64 6", []byte(`{"map_int32_sfixed64": {}}`), map[int32]int64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Sfixed64 }, func() interface{} { return data2.MapInt32Sfixed64 }},
		{"map_int32_sfixed64 7", []byte(`{"map_int32_sfixed64": null}`), map[int32]int64(nil), func() interface{} { return data1.MapInt32Sfixed64 }, func() interface{} { return data2.MapInt32Sfixed64 }},
		{"map_int32_sfixed64 8", []byte(`{"map_int32_sfixed64": {}}`), map[int32]int64{}, func() interface{} { return data1.MapInt32Sfixed64 }, func() interface{} { return data2.MapInt32Sfixed64 }},

		{"map_int32_fixed32 1", []byte(`{"map_int32_fixed32": null}`), map[int32]uint32(nil), func() interface{} { return data1.MapInt32Fixed32 }, func() interface{} { return data2.MapInt32Fixed32 }},
		{"map_int32_fixed32 2", []byte(`{"map_int32_fixed32": {}}`), map[int32]uint32{}, func() interface{} { return data1.MapInt32Fixed32 }, func() interface{} { return data2.MapInt32Fixed32 }},
		{"map_int32_fixed32 3", []byte(`{"map_int32_fixed32": {"1":1, "2": 2, "3": 3}}`), map[int32]uint32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Fixed32 }, func() interface{} { return data2.MapInt32Fixed32 }},
		{"map_int32_fixed32 4", []byte(`{"map_int32_fixed32": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]uint32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Fixed32 }, func() interface{} { return data2.MapInt32Fixed32 }},
		{"map_int32_fixed32 5", []byte(`{"map_int32_fixed32": {"8":8}}`), map[int32]uint32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Fixed32 }, func() interface{} { return data2.MapInt32Fixed32 }},
		{"map_int32_fixed32 6", []byte(`{"map_int32_fixed32": {}}`), map[int32]uint32{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Fixed32 }, func() interface{} { return data2.MapInt32Fixed32 }},
		{"map_int32_fixed32 7", []byte(`{"map_int32_fixed32": null}`), map[int32]uint32(nil), func() interface{} { return data1.MapInt32Fixed32 }, func() interface{} { return data2.MapInt32Fixed32 }},
		{"map_int32_fixed32 8", []byte(`{"map_int32_fixed32": {}}`), map[int32]uint32{}, func() interface{} { return data1.MapInt32Fixed32 }, func() interface{} { return data2.MapInt32Fixed32 }},

		{"map_int32_fixed64 1", []byte(`{"map_int32_fixed64": null}`), map[int32]uint64(nil), func() interface{} { return data1.MapInt32Fixed64 }, func() interface{} { return data2.MapInt32Fixed64 }},
		{"map_int32_fixed64 2", []byte(`{"map_int32_fixed64": {}}`), map[int32]uint64{}, func() interface{} { return data1.MapInt32Fixed64 }, func() interface{} { return data2.MapInt32Fixed64 }},
		{"map_int32_fixed64 3", []byte(`{"map_int32_fixed64": {"1":1, "2": 2, "3": 3}}`), map[int32]uint64{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt32Fixed64 }, func() interface{} { return data2.MapInt32Fixed64 }},
		{"map_int32_fixed64 4", []byte(`{"map_int32_fixed64": {"1":10, "2": 2, "3": 30, "4": 4, "5": 5, "6": 6, "7": 7}}`), map[int32]uint64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7}, func() interface{} { return data1.MapInt32Fixed64 }, func() interface{} { return data2.MapInt32Fixed64 }},
		{"map_int32_fixed64 5", []byte(`{"map_int32_fixed64": {"8":8}}`), map[int32]uint64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Fixed64 }, func() interface{} { return data2.MapInt32Fixed64 }},
		{"map_int32_fixed64 6", []byte(`{"map_int32_fixed64": {}}`), map[int32]uint64{1: 10, 2: 2, 3: 30, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8}, func() interface{} { return data1.MapInt32Fixed64 }, func() interface{} { return data2.MapInt32Fixed64 }},
		{"map_int32_fixed64 7", []byte(`{"map_int32_fixed64": null}`), map[int32]uint64(nil), func() interface{} { return data1.MapInt32Fixed64 }, func() interface{} { return data2.MapInt32Fixed64 }},
		{"map_int32_fixed64 8", []byte(`{"map_int32_fixed64": {}}`), map[int32]uint64{}, func() interface{} { return data1.MapInt32Fixed64 }, func() interface{} { return data2.MapInt32Fixed64 }},

		{"map_int32_bool 1", []byte(`{"map_int32_bool": null}`), map[int32]bool(nil), func() interface{} { return data1.MapInt32Bool }, func() interface{} { return data2.MapInt32Bool }},
		{"map_int32_bool 2", []byte(`{"map_int32_bool": {}}`), map[int32]bool{}, func() interface{} { return data1.MapInt32Bool }, func() interface{} { return data2.MapInt32Bool }},
		{"map_int32_bool 3", []byte(`{"map_int32_bool": {"1":true, "2": false, "3": true}}`), map[int32]bool{1: true, 2: false, 3: true}, func() interface{} { return data1.MapInt32Bool }, func() interface{} { return data2.MapInt32Bool }},
		{"map_int32_bool 4", []byte(`{"map_int32_bool": {"1":true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true}}`), map[int32]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true}, func() interface{} { return data1.MapInt32Bool }, func() interface{} { return data2.MapInt32Bool }},
		{"map_int32_bool 5", []byte(`{"map_int32_bool": {"8":true}}`), map[int32]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true}, func() interface{} { return data1.MapInt32Bool }, func() interface{} { return data2.MapInt32Bool }},
		{"map_int32_bool 6", []byte(`{"map_int32_bool": {}}`), map[int32]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true}, func() interface{} { return data1.MapInt32Bool }, func() interface{} { return data2.MapInt32Bool }},
		{"map_int32_bool 7", []byte(`{"map_int32_bool": null}`), map[int32]bool(nil), func() interface{} { return data1.MapInt32Bool }, func() interface{} { return data2.MapInt32Bool }},
		{"map_int32_bool 8", []byte(`{"map_int32_bool": {}}`), map[int32]bool{}, func() interface{} { return data1.MapInt32Bool }, func() interface{} { return data2.MapInt32Bool }},

		{"map_int32_string 1", []byte(`{"map_int32_string": null}`), map[int32]string(nil), func() interface{} { return data1.MapInt32String }, func() interface{} { return data2.MapInt32String }},
		{"map_int32_string 2", []byte(`{"map_int32_string": {}}`), map[int32]string{}, func() interface{} { return data1.MapInt32String }, func() interface{} { return data2.MapInt32String }},
		{"map_int32_string 3", []byte(`{"map_int32_string": {"1":"v1", "2": "", "3": null, "4": "v4"}}`), map[int32]string{1: "v1", 2: "", 3: "", 4: "v4"}, func() interface{} { return data1.MapInt32String }, func() interface{} { return data2.MapInt32String }},
		{"map_int32_string 4", []byte(`{"map_int32_string": {"1":null, "2": "v2", "5": "v5"}}`), map[int32]string{1: "", 2: "v2", 3: "", 4: "v4", 5: "v5"}, func() interface{} { return data1.MapInt32String }, func() interface{} { return data2.MapInt32String }},
		{"map_int32_string 5", []byte(`{"map_int32_string": {"10":"v10"}}`), map[int32]string{1: "", 2: "v2", 3: "", 4: "v4", 5: "v5", 10: "v10"}, func() interface{} { return data1.MapInt32String }, func() interface{} { return data2.MapInt32String }},
		{"map_int32_string 6", []byte(`{"map_int32_string": {}}`), map[int32]string{1: "", 2: "v2", 3: "", 4: "v4", 5: "v5", 10: "v10"}, func() interface{} { return data1.MapInt32String }, func() interface{} { return data2.MapInt32String }},
		{"map_int32_string 7", []byte(`{"map_int32_string": null}`), map[int32]string(nil), func() interface{} { return data1.MapInt32String }, func() interface{} { return data2.MapInt32String }},
		{"map_int32_string 8", []byte(`{"map_int32_string": {}}`), map[int32]string{}, func() interface{} { return data1.MapInt32String }, func() interface{} { return data2.MapInt32String }},

		{"map_int32_bytes 1", []byte(`{"map_int32_bytes": null}`), map[int32][]uint8(nil), func() interface{} { return data1.MapInt32Bytes }, func() interface{} { return data2.MapInt32Bytes }},
		{"map_int32_bytes 2", []byte(`{"map_int32_bytes": {}}`), map[int32][]byte{}, func() interface{} { return data1.MapInt32Bytes }, func() interface{} { return data2.MapInt32Bytes }},
		{"map_int32_bytes 3", []byte(fmt.Sprintf(`{"map_int32_bytes": {"1": "", "2": "%s", "3": null, "4": "%s"}}`, bb1, bb2)), map[int32][]byte{1: []byte(""), 2: b1, 3: nil, 4: b2}, func() interface{} { return data1.MapInt32Bytes }, func() interface{} { return data2.MapInt32Bytes }},
		{"map_int32_bytes 4", []byte(fmt.Sprintf(`{"map_int32_bytes": {"1": "%s", "2": null, "3": "", "5": "%s"}}`, bb1, bb2)), map[int32][]byte{1: b1, 2: nil, 3: []byte(""), 4: b2, 5: b2}, func() interface{} { return data1.MapInt32Bytes }, func() interface{} { return data2.MapInt32Bytes }},
		{"map_int32_bytes 5", []byte(fmt.Sprintf(`{"map_int32_bytes": {"7": "%s", "8": "%s"}}`, bb1, bb2)), map[int32][]byte{1: b1, 2: nil, 3: []byte(""), 4: b2, 5: b2, 7: b1, 8: b2}, func() interface{} { return data1.MapInt32Bytes }, func() interface{} { return data2.MapInt32Bytes }},
		{"map_int32_bytes 6", []byte(`{"map_int32_bytes": {}}`), map[int32][]byte{1: b1, 2: nil, 3: []byte(""), 4: b2, 5: b2, 7: b1, 8: b2}, func() interface{} { return data1.MapInt32Bytes }, func() interface{} { return data2.MapInt32Bytes }},
		{"map_int32_bytes 7", []byte(`{"map_int32_bytes": null}`), map[int32][]uint8(nil), func() interface{} { return data1.MapInt32Bytes }, func() interface{} { return data2.MapInt32Bytes }},
		{"map_int32_bytes 8", []byte(`{"map_int32_bytes": {}}`), map[int32][]byte{}, func() interface{} { return data1.MapInt32Bytes }, func() interface{} { return data2.MapInt32Bytes }},

		{"map_int32_enum1 1", []byte(`{"map_int32_enum1": null}`), map[int32]gojsontest.UnmarshalData_Enum(nil), func() interface{} { return data1.MapInt32Enum1 }, func() interface{} { return data2.MapInt32Enum1 }},
		{"map_int32_enum1 2", []byte(`{"map_int32_enum1": {}}`), map[int32]gojsontest.UnmarshalData_Enum{}, func() interface{} { return data1.MapInt32Enum1 }, func() interface{} { return data2.MapInt32Enum1 }},
		{"map_int32_enum1 3", []byte(`{"map_int32_enum1": {"1": 1, "2": 0, "3": 1, "4": 0}}`), map[int32]gojsontest.UnmarshalData_Enum{1: gojsontest.UnmarshalData_stopped, 2: gojsontest.UnmarshalData_running, 3: gojsontest.UnmarshalData_stopped, 4: gojsontest.UnmarshalData_running}, func() interface{} { return data1.MapInt32Enum1 }, func() interface{} { return data2.MapInt32Enum1 }},
		{"map_int32_enum1 4", []byte(`{"map_int32_enum1": {"1": 0, "2": 1, "5": 1}}`), map[int32]gojsontest.UnmarshalData_Enum{1: gojsontest.UnmarshalData_running, 2: gojsontest.UnmarshalData_stopped, 3: gojsontest.UnmarshalData_stopped, 4: gojsontest.UnmarshalData_running, 5: gojsontest.UnmarshalData_stopped}, func() interface{} { return data1.MapInt32Enum1 }, func() interface{} { return data2.MapInt32Enum1 }},
		{"map_int32_enum1 5", []byte(`{"map_int32_enum1": {"10": 0}}`), map[int32]gojsontest.UnmarshalData_Enum{1: gojsontest.UnmarshalData_running, 2: gojsontest.UnmarshalData_stopped, 3: gojsontest.UnmarshalData_stopped, 4: gojsontest.UnmarshalData_running, 5: gojsontest.UnmarshalData_stopped, 10: gojsontest.UnmarshalData_running}, func() interface{} { return data1.MapInt32Enum1 }, func() interface{} { return data2.MapInt32Enum1 }},
		{"map_int32_enum1 6", []byte(`{"map_int32_enum1": {}}`), map[int32]gojsontest.UnmarshalData_Enum{1: gojsontest.UnmarshalData_running, 2: gojsontest.UnmarshalData_stopped, 3: gojsontest.UnmarshalData_stopped, 4: gojsontest.UnmarshalData_running, 5: gojsontest.UnmarshalData_stopped, 10: gojsontest.UnmarshalData_running}, func() interface{} { return data1.MapInt32Enum1 }, func() interface{} { return data2.MapInt32Enum1 }},
		{"map_int32_enum1 7", []byte(`{"map_int32_enum1": null}`), map[int32]gojsontest.UnmarshalData_Enum(nil), func() interface{} { return data1.MapInt32Enum1 }, func() interface{} { return data2.MapInt32Enum1 }},
		{"map_int32_enum1 8", []byte(`{"map_int32_enum1": {}}`), map[int32]gojsontest.UnmarshalData_Enum{}, func() interface{} { return data1.MapInt32Enum1 }, func() interface{} { return data2.MapInt32Enum1 }},

		{"map_int32_aliases 1", []byte(`{"map_int32_aliases": null}`), map[int32]*gojsontest.UnmarshalData_Aliases(nil), func() interface{} { return data1.MapInt32Aliases }, func() interface{} { return data2.MapInt32Aliases }},
		{"map_int32_aliases 2", []byte(`{"map_int32_aliases": {}}`), map[int32]*gojsontest.UnmarshalData_Aliases{}, func() interface{} { return data1.MapInt32Aliases }, func() interface{} { return data2.MapInt32Aliases }},
		{"map_int32_aliases 3", []byte(`{"map_int32_aliases": {"1": null, "2": {}, "3": null, "4": {} }}`), map[int32]*gojsontest.UnmarshalData_Aliases{1: nil, 2: {}, 3: nil, 4: {}}, func() interface{} { return data1.MapInt32Aliases }, func() interface{} { return data2.MapInt32Aliases }},
		{"map_int32_aliases 4", []byte(`{"map_int32_aliases": {"1": {}, "2": null, "5": null }}`), map[int32]*gojsontest.UnmarshalData_Aliases{1: {}, 2: nil, 3: nil, 4: {}, 5: nil}, func() interface{} { return data1.MapInt32Aliases }, func() interface{} { return data2.MapInt32Aliases }},
		{"map_int32_aliases 5", []byte(`{"map_int32_aliases": {}}`), map[int32]*gojsontest.UnmarshalData_Aliases{1: {}, 2: nil, 3: nil, 4: {}, 5: nil}, func() interface{} { return data1.MapInt32Aliases }, func() interface{} { return data2.MapInt32Aliases }},
		{"map_int32_aliases 6", []byte(`{"map_int32_aliases": null}`), map[int32]*gojsontest.UnmarshalData_Aliases(nil), func() interface{} { return data1.MapInt32Aliases }, func() interface{} { return data2.MapInt32Aliases }},
		{"map_int32_aliases 7", []byte(`{"map_int32_aliases": {}}`), map[int32]*gojsontest.UnmarshalData_Aliases{}, func() interface{} { return data1.MapInt32Aliases }, func() interface{} { return data2.MapInt32Aliases }},

		{"map_int32_config 1", []byte(`{"map_int32_config": null}`), map[int32]*gojsontest.UnmarshalData_Config(nil), func() interface{} { return data1.MapInt32Config }, func() interface{} { return data2.MapInt32Config }},
		{"map_int32_config 2", []byte(`{"map_int32_config": {}}`), map[int32]*gojsontest.UnmarshalData_Config{}, func() interface{} { return data1.MapInt32Config }, func() interface{} { return data2.MapInt32Config }},
		{"map_int32_config 3", []byte(`{"map_int32_config": {"1": {"ip": "192", "port": 8080}, "2": null, "3": {}, "4": {"ip": "193"}, "5": {"port": 8081} }}`), map[int32]*gojsontest.UnmarshalData_Config{1: {Ip: "192", Port: 8080}, 2: nil, 3: {}, 4: {Ip: "193", Port: 0}, 5: {Ip: "", Port: 8081}}, func() interface{} { return data1.MapInt32Config }, func() interface{} { return data2.MapInt32Config }},
		{"map_int32_config 4", []byte(`{"map_int32_config": {"2": {}, "3": null, "6": {"ip": "194","port": 9091}, "7": null }}`), map[int32]*gojsontest.UnmarshalData_Config{1: {Ip: "192", Port: 8080}, 2: {}, 3: nil, 4: {Ip: "193", Port: 0}, 5: {Ip: "", Port: 8081}, 6: {Ip: "194", Port: 9091}, 7: nil}, func() interface{} { return data1.MapInt32Config }, func() interface{} { return data2.MapInt32Config }},
		{"map_int32_config 5", []byte(`{"map_int32_config": {}}`), map[int32]*gojsontest.UnmarshalData_Config{1: {Ip: "192", Port: 8080}, 2: {}, 3: nil, 4: {Ip: "193", Port: 0}, 5: {Ip: "", Port: 8081}, 6: {Ip: "194", Port: 9091}, 7: nil}, func() interface{} { return data1.MapInt32Config }, func() interface{} { return data2.MapInt32Config }},
		{"map_int32_config 6", []byte(`{"map_int32_config": null}`), map[int32]*gojsontest.UnmarshalData_Config(nil), func() interface{} { return data1.MapInt32Config }, func() interface{} { return data2.MapInt32Config }},
		{"map_int32_config 7", []byte(`{"map_int32_config": {}}`), map[int32]*gojsontest.UnmarshalData_Config{}, func() interface{} { return data1.MapInt32Config }, func() interface{} { return data2.MapInt32Config }},

		{"map_int64_int32 1", []byte(`{"map_int64_int32": null}`), map[int64]int32(nil), func() interface{} { return data1.MapInt64Int32 }, func() interface{} { return data2.MapInt64Int32 }},
		{"map_int64_int32 2", []byte(`{"map_int64_int32": {}}`), map[int64]int32{}, func() interface{} { return data1.MapInt64Int32 }, func() interface{} { return data2.MapInt64Int32 }},
		{"map_int64_int32 3", []byte(`{"map_int64_int32": {"1": 1, "2": 2, "3": 3}}`), map[int64]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapInt64Int32 }, func() interface{} { return data2.MapInt64Int32 }},
		{"map_int64_int32 4", []byte(`{"map_int64_int32": {"2": 22, "4": 4, "5":5}}`), map[int64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapInt64Int32 }, func() interface{} { return data2.MapInt64Int32 }},
		{"map_int64_int32 5", []byte(`{"map_int64_int32": {}}`), map[int64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapInt64Int32 }, func() interface{} { return data2.MapInt64Int32 }},
		{"map_int64_int32 6", []byte(`{"map_int64_int32": null}`), map[int64]int32(nil), func() interface{} { return data1.MapInt64Int32 }, func() interface{} { return data2.MapInt64Int32 }},
		{"map_int64_int32 7", []byte(`{"map_int64_int32": {}}`), map[int64]int32{}, func() interface{} { return data1.MapInt64Int32 }, func() interface{} { return data2.MapInt64Int32 }},

		{"map_uint32_int32 1", []byte(`{"map_uint32_int32": null}`), map[uint32]int32(nil), func() interface{} { return data1.MapUint32Int32 }, func() interface{} { return data2.MapUint32Int32 }},
		{"map_uint32_int32 2", []byte(`{"map_uint32_int32": {}}`), map[uint32]int32{}, func() interface{} { return data1.MapUint32Int32 }, func() interface{} { return data2.MapUint32Int32 }},
		{"map_uint32_int32 3", []byte(`{"map_uint32_int32": {"1": 1, "2": 2, "3": 3}}`), map[uint32]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapUint32Int32 }, func() interface{} { return data2.MapUint32Int32 }},
		{"map_uint32_int32 4", []byte(`{"map_uint32_int32": {"2": 22, "4": 4, "5":5}}`), map[uint32]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapUint32Int32 }, func() interface{} { return data2.MapUint32Int32 }},
		{"map_uint32_int32 5", []byte(`{"map_uint32_int32": {}}`), map[uint32]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapUint32Int32 }, func() interface{} { return data2.MapUint32Int32 }},
		{"map_uint32_int32 6", []byte(`{"map_uint32_int32": null}`), map[uint32]int32(nil), func() interface{} { return data1.MapUint32Int32 }, func() interface{} { return data2.MapUint32Int32 }},
		{"map_uint32_int32 7", []byte(`{"map_uint32_int32": {}}`), map[uint32]int32{}, func() interface{} { return data1.MapUint32Int32 }, func() interface{} { return data2.MapUint32Int32 }},

		{"map_uint64_int32 1", []byte(`{"map_uint64_int32": null}`), map[uint64]int32(nil), func() interface{} { return data1.MapUint64Int32 }, func() interface{} { return data2.MapUint64Int32 }},
		{"map_uint64_int32 2", []byte(`{"map_uint64_int32": {}}`), map[uint64]int32{}, func() interface{} { return data1.MapUint64Int32 }, func() interface{} { return data2.MapUint64Int32 }},
		{"map_uint64_int32 3", []byte(`{"map_uint64_int32": {"1": 1, "2": 2, "3": 3}}`), map[uint64]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapUint64Int32 }, func() interface{} { return data2.MapUint64Int32 }},
		{"map_uint64_int32 4", []byte(`{"map_uint64_int32": {"2": 22, "4": 4, "5":5}}`), map[uint64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapUint64Int32 }, func() interface{} { return data2.MapUint64Int32 }},
		{"map_uint64_int32 5", []byte(`{"map_uint64_int32": {}}`), map[uint64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapUint64Int32 }, func() interface{} { return data2.MapUint64Int32 }},
		{"map_uint64_int32 6", []byte(`{"map_uint64_int32": null}`), map[uint64]int32(nil), func() interface{} { return data1.MapUint64Int32 }, func() interface{} { return data2.MapUint64Int32 }},
		{"map_uint64_int32 7", []byte(`{"map_uint64_int32": {}}`), map[uint64]int32{}, func() interface{} { return data1.MapUint64Int32 }, func() interface{} { return data2.MapUint64Int32 }},

		{"map_sint32_int32 1", []byte(`{"map_sint32_int32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapSint32Int32 }, func() interface{} { return data2.MapSint32Int32 }},
		{"map_sint32_int32 2", []byte(`{"map_sint32_int32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapSint32Int32 }, func() interface{} { return data2.MapSint32Int32 }},
		{"map_sint32_int32 3", []byte(`{"map_sint32_int32": {"1": 1, "2": 2, "3": 3}}`), map[int32]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapSint32Int32 }, func() interface{} { return data2.MapSint32Int32 }},
		{"map_sint32_int32 4", []byte(`{"map_sint32_int32": {"2": 22, "4": 4, "5":5}}`), map[int32]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapSint32Int32 }, func() interface{} { return data2.MapSint32Int32 }},
		{"map_sint32_int32 5", []byte(`{"map_sint32_int32": {}}`), map[int32]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapSint32Int32 }, func() interface{} { return data2.MapSint32Int32 }},
		{"map_sint32_int32 6", []byte(`{"map_sint32_int32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapSint32Int32 }, func() interface{} { return data2.MapSint32Int32 }},
		{"map_sint32_int32 7", []byte(`{"map_sint32_int32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapSint32Int32 }, func() interface{} { return data2.MapSint32Int32 }},

		{"map_sint64_int32 1", []byte(`{"map_sint64_int32": null}`), map[int64]int32(nil), func() interface{} { return data1.MapSint64Int32 }, func() interface{} { return data2.MapSint64Int32 }},
		{"map_sint64_int32 2", []byte(`{"map_sint64_int32": {}}`), map[int64]int32{}, func() interface{} { return data1.MapSint64Int32 }, func() interface{} { return data2.MapSint64Int32 }},
		{"map_sint64_int32 3", []byte(`{"map_sint64_int32": {"1": 1, "2": 2, "3": 3}}`), map[int64]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapSint64Int32 }, func() interface{} { return data2.MapSint64Int32 }},
		{"map_sint64_int32 4", []byte(`{"map_sint64_int32": {"2": 22, "4": 4, "5":5}}`), map[int64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapSint64Int32 }, func() interface{} { return data2.MapSint64Int32 }},
		{"map_sint64_int32 5", []byte(`{"map_sint64_int32": {}}`), map[int64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapSint64Int32 }, func() interface{} { return data2.MapSint64Int32 }},
		{"map_sint64_int32 6", []byte(`{"map_sint64_int32": null}`), map[int64]int32(nil), func() interface{} { return data1.MapSint64Int32 }, func() interface{} { return data2.MapSint64Int32 }},
		{"map_sint64_int32 7", []byte(`{"map_sint64_int32": {}}`), map[int64]int32{}, func() interface{} { return data1.MapSint64Int32 }, func() interface{} { return data2.MapSint64Int32 }},

		{"map_fixed32_int32 1", []byte(`{"map_fixed32_int32": null}`), map[uint32]int32(nil), func() interface{} { return data1.MapFixed32Int32 }, func() interface{} { return data2.MapFixed32Int32 }},
		{"map_fixed32_int32 2", []byte(`{"map_fixed32_int32": {}}`), map[uint32]int32{}, func() interface{} { return data1.MapFixed32Int32 }, func() interface{} { return data2.MapFixed32Int32 }},
		{"map_fixed32_int32 3", []byte(`{"map_fixed32_int32": {"1": 1, "2": 2, "3": 3}}`), map[uint32]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapFixed32Int32 }, func() interface{} { return data2.MapFixed32Int32 }},
		{"map_fixed32_int32 4", []byte(`{"map_fixed32_int32": {"2": 22, "4": 4, "5":5}}`), map[uint32]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapFixed32Int32 }, func() interface{} { return data2.MapFixed32Int32 }},
		{"map_fixed32_int32 5", []byte(`{"map_fixed32_int32": {}}`), map[uint32]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapFixed32Int32 }, func() interface{} { return data2.MapFixed32Int32 }},
		{"map_fixed32_int32 6", []byte(`{"map_fixed32_int32": null}`), map[uint32]int32(nil), func() interface{} { return data1.MapFixed32Int32 }, func() interface{} { return data2.MapFixed32Int32 }},
		{"map_fixed32_int32 7", []byte(`{"map_fixed32_int32": {}}`), map[uint32]int32{}, func() interface{} { return data1.MapFixed32Int32 }, func() interface{} { return data2.MapFixed32Int32 }},

		{"map_fixed64_int32 1", []byte(`{"map_fixed64_int32": null}`), map[uint64]int32(nil), func() interface{} { return data1.MapFixed64Int32 }, func() interface{} { return data2.MapFixed64Int32 }},
		{"map_fixed64_int32 2", []byte(`{"map_fixed64_int32": {}}`), map[uint64]int32{}, func() interface{} { return data1.MapFixed64Int32 }, func() interface{} { return data2.MapFixed64Int32 }},
		{"map_fixed64_int32 3", []byte(`{"map_fixed64_int32": {"1": 1, "2": 2, "3": 3}}`), map[uint64]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapFixed64Int32 }, func() interface{} { return data2.MapFixed64Int32 }},
		{"map_fixed64_int32 4", []byte(`{"map_fixed64_int32": {"2": 22, "4": 4, "5":5}}`), map[uint64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapFixed64Int32 }, func() interface{} { return data2.MapFixed64Int32 }},
		{"map_fixed64_int32 5", []byte(`{"map_fixed64_int32": {}}`), map[uint64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapFixed64Int32 }, func() interface{} { return data2.MapFixed64Int32 }},
		{"map_fixed64_int32 6", []byte(`{"map_fixed64_int32": null}`), map[uint64]int32(nil), func() interface{} { return data1.MapFixed64Int32 }, func() interface{} { return data2.MapFixed64Int32 }},
		{"map_fixed64_int32 7", []byte(`{"map_fixed64_int32": {}}`), map[uint64]int32{}, func() interface{} { return data1.MapFixed64Int32 }, func() interface{} { return data2.MapFixed64Int32 }},

		{"map_sfixed32_int32 1", []byte(`{"map_sfixed32_int32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapSfixed32Int32 }, func() interface{} { return data2.MapSfixed32Int32 }},
		{"map_sfixed32_int32 2", []byte(`{"map_sfixed32_int32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapSfixed32Int32 }, func() interface{} { return data2.MapSfixed32Int32 }},
		{"map_sfixed32_int32 3", []byte(`{"map_sfixed32_int32": {"1": 1, "2": 2, "3": 3}}`), map[int32]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapSfixed32Int32 }, func() interface{} { return data2.MapSfixed32Int32 }},
		{"map_sfixed32_int32 4", []byte(`{"map_sfixed32_int32": {"2": 22, "4": 4, "5":5}}`), map[int32]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapSfixed32Int32 }, func() interface{} { return data2.MapSfixed32Int32 }},
		{"map_sfixed32_int32 5", []byte(`{"map_sfixed32_int32": {}}`), map[int32]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapSfixed32Int32 }, func() interface{} { return data2.MapSfixed32Int32 }},
		{"map_sfixed32_int32 6", []byte(`{"map_sfixed32_int32": null}`), map[int32]int32(nil), func() interface{} { return data1.MapSfixed32Int32 }, func() interface{} { return data2.MapSfixed32Int32 }},
		{"map_sfixed32_int32 7", []byte(`{"map_sfixed32_int32": {}}`), map[int32]int32{}, func() interface{} { return data1.MapSfixed32Int32 }, func() interface{} { return data2.MapSfixed32Int32 }},

		{"map_sfixed64_int32 1", []byte(`{"map_sfixed64_int32": null}`), map[int64]int32(nil), func() interface{} { return data1.MapSfixed64Int32 }, func() interface{} { return data2.MapSfixed64Int32 }},
		{"map_sfixed64_int32 2", []byte(`{"map_sfixed64_int32": {}}`), map[int64]int32{}, func() interface{} { return data1.MapSfixed64Int32 }, func() interface{} { return data2.MapSfixed64Int32 }},
		{"map_sfixed64_int32 3", []byte(`{"map_sfixed64_int32": {"1": 1, "2": 2, "3": 3}}`), map[int64]int32{1: 1, 2: 2, 3: 3}, func() interface{} { return data1.MapSfixed64Int32 }, func() interface{} { return data2.MapSfixed64Int32 }},
		{"map_sfixed64_int32 4", []byte(`{"map_sfixed64_int32": {"2": 22, "4": 4, "5":5}}`), map[int64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapSfixed64Int32 }, func() interface{} { return data2.MapSfixed64Int32 }},
		{"map_sfixed64_int32 5", []byte(`{"map_sfixed64_int32": {}}`), map[int64]int32{1: 1, 2: 22, 3: 3, 4: 4, 5: 5}, func() interface{} { return data1.MapSfixed64Int32 }, func() interface{} { return data2.MapSfixed64Int32 }},
		{"map_sfixed64_int32 6", []byte(`{"map_sfixed64_int32": null}`), map[int64]int32(nil), func() interface{} { return data1.MapSfixed64Int32 }, func() interface{} { return data2.MapSfixed64Int32 }},
		{"map_sfixed64_int32 7", []byte(`{"map_sfixed64_int32": {}}`), map[int64]int32{}, func() interface{} { return data1.MapSfixed64Int32 }, func() interface{} { return data2.MapSfixed64Int32 }},

		{"map_string_int32 1", []byte(`{"map_string_int32": null}`), map[string]int32(nil), func() interface{} { return data1.MapStringInt32 }, func() interface{} { return data2.MapStringInt32 }},
		{"map_string_int32 2", []byte(`{"map_string_int32": {}}`), map[string]int32{}, func() interface{} { return data1.MapStringInt32 }, func() interface{} { return data2.MapStringInt32 }},
		{"map_string_int32 3", []byte(`{"map_string_int32": {"k1":1, "k2": 2, "k3": 3}}`), map[string]int32{"k1": 1, "k2": 2, "k3": 3}, func() interface{} { return data1.MapStringInt32 }, func() interface{} { return data2.MapStringInt32 }},
		{"map_string_int32 4", []byte(`{"map_string_int32": {"k1": 11, "k4": 4}}`), map[string]int32{"k1": 11, "k2": 2, "k3": 3, "k4": 4}, func() interface{} { return data1.MapStringInt32 }, func() interface{} { return data2.MapStringInt32 }},
		{"map_string_int32 5", []byte(`{"map_string_int32": {}}`), map[string]int32{"k1": 11, "k2": 2, "k3": 3, "k4": 4}, func() interface{} { return data1.MapStringInt32 }, func() interface{} { return data2.MapStringInt32 }},
		{"map_string_int32 6", []byte(`{"map_string_int32": null}`), map[string]int32(nil), func() interface{} { return data1.MapStringInt32 }, func() interface{} { return data2.MapStringInt32 }},
		{"map_string_int32 7", []byte(`{"map_string_int32": {}}`), map[string]int32{}, func() interface{} { return data1.MapStringInt32 }, func() interface{} { return data2.MapStringInt32 }},
	}

	var cases []*CaseDesc

	cases = append(cases, literalCases...)
	cases = append(cases, arrayCases...)
	cases = append(cases, mapCases...)

	// standard json.
	for _, c := range cases {
		err := json.Unmarshal(c.B, data2)
		require.Nil(t, err, c.Name, err)
		require.Equal(t, c.Expected, c.Actual2(), "STD: "+c.Name)
	}
	// go json.
	for _, c := range cases {
		err := data1.UnmarshalJSON(c.B)
		require.Nil(t, err, c.Name, err)
		require.Equal(t, c.Expected, c.Actual1(), "GEN: "+c.Name)
	}

	// use enum string
	enumCases := []*CaseDesc{
		{"t_enum2 1", []byte(`{"t_enum2": "stopped"}`), gojsontest.UnmarshalData_stopped, func() interface{} { return data1.TEnum2 }, func() interface{} { return data2.TEnum2 }},
		{"t_enum2 2", []byte(`{"t_enum2": "running"}`), gojsontest.UnmarshalData_running, func() interface{} { return data1.TEnum2 }, func() interface{} { return data2.TEnum2 }},
		{"t_enum2 3", []byte(`{"t_enum2": "stopped"}`), gojsontest.UnmarshalData_stopped, func() interface{} { return data1.TEnum2 }, func() interface{} { return data2.TEnum2 }},
		{"t_enum2 4", []byte(`{}`), gojsontest.UnmarshalData_stopped, func() interface{} { return data1.TEnum2 }, func() interface{} { return data2.TEnum2 }},

		{"array_enum2 1", []byte(`{"array_enum2": null}`), []gojsontest.UnmarshalData_Enum(nil), func() interface{} { return data1.ArrayEnum2 }, func() interface{} { return data2.ArrayEnum2 }},
		{"array_enum2 2", []byte(`{"array_enum2": []}`), []gojsontest.UnmarshalData_Enum{}, func() interface{} { return data1.ArrayEnum2 }, func() interface{} { return data2.ArrayEnum2 }},
		{"array_enum2 3", []byte(`{"array_enum2": ["stopped", "running", "stopped"]}`), []gojsontest.UnmarshalData_Enum{gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_running, gojsontest.UnmarshalData_stopped}, func() interface{} { return data1.ArrayEnum2 }, func() interface{} { return data2.ArrayEnum2 }},
		{"array_enum2 4", []byte(`{"array_enum2": ["running", "stopped", "running"]}`), []gojsontest.UnmarshalData_Enum{gojsontest.UnmarshalData_running, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_running}, func() interface{} { return data1.ArrayEnum2 }, func() interface{} { return data2.ArrayEnum2 }},
		{"array_enum2 5", []byte(`{"array_enum2": ["stopped"]}`), []gojsontest.UnmarshalData_Enum{gojsontest.UnmarshalData_stopped}, func() interface{} { return data1.ArrayEnum2 }, func() interface{} { return data2.ArrayEnum2 }},
		{"array_enum2 6", []byte(`{"array_enum2": ["stopped","stopped","stopped","stopped","stopped","stopped","stopped","stopped","stopped","stopped"]}`), []gojsontest.UnmarshalData_Enum{gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped, gojsontest.UnmarshalData_stopped}, func() interface{} { return data1.ArrayEnum2 }, func() interface{} { return data2.ArrayEnum2 }},
		{"array_enum2 7", []byte(`{"array_enum2": []}`), []gojsontest.UnmarshalData_Enum{}, func() interface{} { return data1.ArrayEnum2 }, func() interface{} { return data2.ArrayEnum2 }},
		{"array_enum2 8", []byte(`{"array_enum2": null}`), []gojsontest.UnmarshalData_Enum(nil), func() interface{} { return data1.ArrayEnum2 }, func() interface{} { return data2.ArrayEnum2 }},

		{"map_int32_enum2 1", []byte(`{"map_int32_enum2": null}`), map[int32]gojsontest.UnmarshalData_Enum(nil), func() interface{} { return data1.MapInt32Enum2 }, func() interface{} { return data2.MapInt32Enum2 }},
		{"map_int32_enum2 2", []byte(`{"map_int32_enum2": {}}`), map[int32]gojsontest.UnmarshalData_Enum{}, func() interface{} { return data1.MapInt32Enum2 }, func() interface{} { return data2.MapInt32Enum2 }},
		{"map_int32_enum2 3", []byte(`{"map_int32_enum2": {"1": "stopped", "2": "running", "3": "stopped", "4": "running"}}`), map[int32]gojsontest.UnmarshalData_Enum{1: gojsontest.UnmarshalData_stopped, 2: gojsontest.UnmarshalData_running, 3: gojsontest.UnmarshalData_stopped, 4: gojsontest.UnmarshalData_running}, func() interface{} { return data1.MapInt32Enum2 }, func() interface{} { return data2.MapInt32Enum2 }},
		{"map_int32_enum2 4", []byte(`{"map_int32_enum2": {"1": "running", "2": "stopped", "5": "stopped"}}`), map[int32]gojsontest.UnmarshalData_Enum{1: gojsontest.UnmarshalData_running, 2: gojsontest.UnmarshalData_stopped, 3: gojsontest.UnmarshalData_stopped, 4: gojsontest.UnmarshalData_running, 5: gojsontest.UnmarshalData_stopped}, func() interface{} { return data1.MapInt32Enum2 }, func() interface{} { return data2.MapInt32Enum2 }},
		{"map_int32_enum2 5", []byte(`{"map_int32_enum2": {"10": "running"}}`), map[int32]gojsontest.UnmarshalData_Enum{1: gojsontest.UnmarshalData_running, 2: gojsontest.UnmarshalData_stopped, 3: gojsontest.UnmarshalData_stopped, 4: gojsontest.UnmarshalData_running, 5: gojsontest.UnmarshalData_stopped, 10: gojsontest.UnmarshalData_running}, func() interface{} { return data1.MapInt32Enum2 }, func() interface{} { return data2.MapInt32Enum2 }},
		{"map_int32_enum2 6", []byte(`{"map_int32_enum2": {}}`), map[int32]gojsontest.UnmarshalData_Enum{1: gojsontest.UnmarshalData_running, 2: gojsontest.UnmarshalData_stopped, 3: gojsontest.UnmarshalData_stopped, 4: gojsontest.UnmarshalData_running, 5: gojsontest.UnmarshalData_stopped, 10: gojsontest.UnmarshalData_running}, func() interface{} { return data1.MapInt32Enum2 }, func() interface{} { return data2.MapInt32Enum2 }},
		{"map_int32_enum2 7", []byte(`{"map_int32_enum2": null}`), map[int32]gojsontest.UnmarshalData_Enum(nil), func() interface{} { return data1.MapInt32Enum2 }, func() interface{} { return data2.MapInt32Enum2 }},
		{"map_int32_enum2 8", []byte(`{"map_int32_enum2": {}}`), map[int32]gojsontest.UnmarshalData_Enum{}, func() interface{} { return data1.MapInt32Enum2 }, func() interface{} { return data2.MapInt32Enum2 }},
	}

	// go json.
	for _, c := range enumCases {
		err := data1.UnmarshalJSON(c.B)
		require.Nil(t, err, c.Name, err)
		require.Equal(t, c.Expected, c.Actual1(), "GEN: "+c.Name)
	}
}

func Test_GoJSON_UnmarshalData_CheckDiff(t *testing.T) {
	type UnmarshalData gojsontest.UnmarshalData

	var err error

	data1 := &gojsontest.UnmarshalData{}
	data2 := &UnmarshalData{}

	// Literal String
	{
		b1 := []byte(`{"t_string": "Hello World"}`)

		err = json.Unmarshal(b1, data2)
		require.Nil(t, err)
		require.Equal(t, "Hello World", data2.TString)

		err = data1.UnmarshalJSON(b1)
		require.Nil(t, err)
		require.Equal(t, "Hello World", data1.TString)

		b2 := []byte(`{"t_string": null}`)

		err = json.Unmarshal(b2, data2)
		require.Nil(t, err)
		require.Equal(t, "Hello World", data2.TString)

		err = data1.UnmarshalJSON(b2)
		require.Nil(t, err)
		require.Equal(t, "", data1.TString)
	}

	// Array String
	{
		b1 := []byte(`{"array_string": ["s1", "s2", null, "s3", ""]}`)

		err = json.Unmarshal(b1, data2)
		require.Nil(t, err)
		require.Equal(t, []string{"s1", "s2", "", "s3", ""}, data2.ArrayString)

		err = data1.UnmarshalJSON(b1)
		require.Nil(t, err)
		require.Equal(t, []string{"s1", "s2", "", "s3", ""}, data1.ArrayString)

		b2 := []byte(`{"array_string": ["", null, "s4"]}`)

		err = json.Unmarshal(b2, data2)
		require.Nil(t, err)
		require.Equal(t, []string{"", "s2", "s4"}, data2.ArrayString)

		err = data1.UnmarshalJSON(b2)
		require.Nil(t, err)
		require.Equal(t, []string{"", "", "s4"}, data1.ArrayString)
	}
}

func Test_GoJSON_UnmarshalData_CheckBoundary(t *testing.T) {
	type UnmarshalData gojsontest.UnmarshalData

	b := []byte(`
{
	"t_string": "t_string_value",
	"t_int32": 11,
	"t_enum1": 1,
	"t_config": {
		"ip": "192.168.10.1",
		"port": 8080
	},
	"array_int32": [
		1,
		2,
		3
	],
	"array_int64": [
		1,
		2,
		3
	],
	"array_config": [{
			"ip": "172.1",
			"port": 8081
		},
		{
			"ip": "172.2",
			"port": 8082
		},
		{
			"ip": "172.3",
			"port": 8083
		}
	],
	"map_int32_int32": {
		"1": 1,
		"2": 2,
		"3": 3
	},
	"map_int32_int64": {
		"1": 1,
		"2": 2,
		"3": 3
	},
	"map_int32_config": {
		"1": {
			"ip": "10.1",
			"port": 8081
		},
		"2": {},
		"3": null,
		"4": {
			"ip": "10.4",
			"port": 8084
		}
	}
}
`)

	t.Run("standard json", func(t *testing.T) {
		tConfig := &gojsontest.UnmarshalData_Config{Ip: "172", Port: 80}
		arrayInt32 := []int32{4, 5, 7, 8}
		arrayInt64 := []int64{10, 11}

		aConfig1 := &gojsontest.UnmarshalData_Config{Ip: "10.1", Port: 1000}
		aConfig2 := &gojsontest.UnmarshalData_Config{Ip: "10.2", Port: 1001}
		arrayConfig := []*gojsontest.UnmarshalData_Config{aConfig1, aConfig2, nil, nil, nil}

		mapInt32Int32 := map[int32]int32{1: 10, 2: 20, 3: 30, 4: 4}

		mapConfig1 := &gojsontest.UnmarshalData_Config{Ip: "192.1", Port: 1001}
		mapConfig2 := &gojsontest.UnmarshalData_Config{Ip: "192.2", Port: 1002}
		mapConfig3 := &gojsontest.UnmarshalData_Config{Ip: "192.3", Port: 1003}
		mapConfig5 := &gojsontest.UnmarshalData_Config{Ip: "192.5", Port: 1005}
		mapInt32Config := map[int32]*gojsontest.UnmarshalData_Config{1: mapConfig1, 2: mapConfig2, 3: mapConfig3, 5: mapConfig5}

		data2 := &UnmarshalData{TConfig: tConfig, ArrayInt32: arrayInt32, ArrayInt64: arrayInt64, ArrayConfig: arrayConfig, MapInt32Int32: mapInt32Int32, MapInt32Config: mapInt32Config}
		_, ok := (interface{}(data2)).(json.Marshaler)
		require.False(t, ok)

		err := json.Unmarshal(b, data2)
		require.Nil(t, err)
		require.Equal(t, data2.TString, "t_string_value")
		require.Equal(t, data2.TInt32, int32(11))
		require.Equal(t, data2.TEnum1, gojsontest.UnmarshalData_stopped)
		require.Equal(t, data2.TConfig, &gojsontest.UnmarshalData_Config{Ip: "192.168.10.1", Port: 8080})
		require.Equal(t, unsafe.Pointer(data2.TConfig), unsafe.Pointer(tConfig))
		require.Equal(t, reflect.ValueOf(data2.TConfig).Pointer(), reflect.ValueOf(tConfig).Pointer())

		require.Equal(t, data2.ArrayInt32, []int32{1, 2, 3})
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data2.ArrayInt32)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&arrayInt32)).Data)
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data2.ArrayInt32)).Len, 3)
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data2.ArrayInt32)).Cap, 4)

		require.Equal(t, data2.ArrayInt64, []int64{1, 2, 3})
		require.NotEqual(t, (*reflect.SliceHeader)(unsafe.Pointer(&data2.ArrayInt64)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&arrayInt64)).Data)

		require.Equal(t, data2.ArrayConfig, []*gojsontest.UnmarshalData_Config{{Ip: "172.1", Port: 8081}, {Ip: "172.2", Port: 8082}, {Ip: "172.3", Port: 8083}})
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data2.ArrayConfig)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&arrayConfig)).Data)
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data2.ArrayConfig)).Len, 3)
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data2.ArrayConfig)).Cap, 5)
		require.Equal(t, reflect.ValueOf(data2.ArrayConfig[0]).Pointer(), reflect.ValueOf(aConfig1).Pointer())
		require.Equal(t, reflect.ValueOf(data2.ArrayConfig[1]).Pointer(), reflect.ValueOf(aConfig2).Pointer())
		require.Equal(t, unsafe.Pointer(data2.ArrayConfig[0]), unsafe.Pointer(aConfig1))
		require.Equal(t, unsafe.Pointer(data2.ArrayConfig[1]), unsafe.Pointer(aConfig2))

		require.Equal(t, data2.MapInt32Int32, map[int32]int32{1: 1, 2: 2, 3: 3, 4: 4})
		require.Equal(t, reflect.ValueOf(data2.MapInt32Int32).Pointer(), reflect.ValueOf(mapInt32Int32).Pointer())

		require.Equal(t, data2.MapInt32Int64, map[int32]int64{1: 1, 2: 2, 3: 3})

		require.Equal(t, data2.MapInt32Config, map[int32]*gojsontest.UnmarshalData_Config{1: {Ip: "10.1", Port: 8081}, 2: {Ip: "", Port: 0}, 3: nil, 4: {Ip: "10.4", Port: 8084}, 5: {Ip: "192.5", Port: 1005}})
		require.Equal(t, reflect.ValueOf(data2.MapInt32Config).Pointer(), reflect.ValueOf(mapInt32Config).Pointer())
		require.NotEqual(t, reflect.ValueOf(data2.MapInt32Config[1]).Pointer(), reflect.ValueOf(mapConfig1).Pointer()) // diff with standard json.
		require.Equal(t, reflect.ValueOf(data2.MapInt32Config[5]).Pointer(), reflect.ValueOf(mapConfig5).Pointer())
	})

	//------ ------

	t.Run("gojson", func(t *testing.T) {
		tConfig := &gojsontest.UnmarshalData_Config{Ip: "172", Port: 80}
		arrayInt32 := []int32{4, 5, 7, 8}
		arrayInt64 := []int64{10, 11}

		aConfig1 := &gojsontest.UnmarshalData_Config{Ip: "10.1", Port: 1000}
		aConfig2 := &gojsontest.UnmarshalData_Config{Ip: "10.2", Port: 1001}
		arrayConfig := []*gojsontest.UnmarshalData_Config{aConfig1, aConfig2, nil, nil, nil}

		mapInt32Int32 := map[int32]int32{1: 10, 2: 20, 3: 30, 4: 4}

		mapConfig1 := &gojsontest.UnmarshalData_Config{Ip: "192.1", Port: 1001}
		mapConfig2 := &gojsontest.UnmarshalData_Config{Ip: "192.2", Port: 1002}
		mapConfig3 := &gojsontest.UnmarshalData_Config{Ip: "192.3", Port: 1003}
		mapConfig5 := &gojsontest.UnmarshalData_Config{Ip: "192.5", Port: 1005}
		mapInt32Config := map[int32]*gojsontest.UnmarshalData_Config{1: mapConfig1, 2: mapConfig2, 3: mapConfig3, 5: mapConfig5}

		data1 := &gojsontest.UnmarshalData{TConfig: tConfig, ArrayInt32: arrayInt32, ArrayInt64: arrayInt64, ArrayConfig: arrayConfig, MapInt32Int32: mapInt32Int32, MapInt32Config: mapInt32Config}
		_, ok := (interface{}(data1)).(json.Marshaler)
		require.True(t, ok)

		err := data1.UnmarshalJSON(b)
		require.Nil(t, err)
		require.Equal(t, data1.TString, "t_string_value")
		require.Equal(t, data1.TInt32, int32(11))
		require.Equal(t, data1.TEnum1, gojsontest.UnmarshalData_stopped)
		require.Equal(t, data1.TConfig, &gojsontest.UnmarshalData_Config{Ip: "192.168.10.1", Port: 8080})
		require.Equal(t, unsafe.Pointer(data1.TConfig), unsafe.Pointer(tConfig))

		require.Equal(t, data1.ArrayInt32, []int32{1, 2, 3})
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data1.ArrayInt32)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&arrayInt32)).Data)
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data1.ArrayInt32)).Len, 3)
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data1.ArrayInt32)).Cap, 4)

		require.Equal(t, data1.ArrayInt64, []int64{1, 2, 3})
		require.NotEqual(t, (*reflect.SliceHeader)(unsafe.Pointer(&data1.ArrayInt64)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&arrayInt64)).Data)

		require.Equal(t, data1.ArrayConfig, []*gojsontest.UnmarshalData_Config{{Ip: "172.1", Port: 8081}, {Ip: "172.2", Port: 8082}, {Ip: "172.3", Port: 8083}})
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data1.ArrayConfig)).Data, (*reflect.SliceHeader)(unsafe.Pointer(&arrayConfig)).Data)
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data1.ArrayConfig)).Len, 3)
		require.Equal(t, (*reflect.SliceHeader)(unsafe.Pointer(&data1.ArrayConfig)).Cap, 5)
		require.Equal(t, reflect.ValueOf(data1.ArrayConfig[0]).Pointer(), reflect.ValueOf(aConfig1).Pointer())
		require.Equal(t, reflect.ValueOf(data1.ArrayConfig[1]).Pointer(), reflect.ValueOf(aConfig2).Pointer())
		require.Equal(t, unsafe.Pointer(data1.ArrayConfig[0]), unsafe.Pointer(aConfig1))
		require.Equal(t, unsafe.Pointer(data1.ArrayConfig[1]), unsafe.Pointer(aConfig2))

		require.Equal(t, data1.MapInt32Int32, map[int32]int32{1: 1, 2: 2, 3: 3, 4: 4})
		require.Equal(t, reflect.ValueOf(data1.MapInt32Int32).Pointer(), reflect.ValueOf(mapInt32Int32).Pointer())

		require.Equal(t, data1.MapInt32Int64, map[int32]int64{1: 1, 2: 2, 3: 3})
		// NOTICE: Diff with standard json.
		require.Equal(t, data1.MapInt32Config, map[int32]*gojsontest.UnmarshalData_Config{1: {Ip: "10.1", Port: 8081}, 2: {Ip: "192.2", Port: 1002}, 3: nil, 4: {Ip: "10.4", Port: 8084}, 5: {Ip: "192.5", Port: 1005}})
		require.Equal(t, reflect.ValueOf(data1.MapInt32Config).Pointer(), reflect.ValueOf(mapInt32Config).Pointer())
		// NOTICE: Diff with standard json.
		require.Equal(t, reflect.ValueOf(data1.MapInt32Config[1]).Pointer(), reflect.ValueOf(mapConfig1).Pointer())
		require.Equal(t, reflect.ValueOf(data1.MapInt32Config[5]).Pointer(), reflect.ValueOf(mapConfig5).Pointer())
	})
}

func Test_GoJSON_UnmarshalOneofNotHide_CheckCorrect(t *testing.T) {
	type CaseDesc struct {
		Name     string
		B        []byte
		Expected interface{}                             // Expected value.
		Actual   func() (interface{}, bool, interface{}) // return oneof type, assert ok, value.
	}

	b1 := []byte("Hello Bytes 1")
	b2 := []byte("Hello Bytes 2")

	bb1 := make([]byte, base64.StdEncoding.EncodedLen(len(b1)))
	base64.StdEncoding.Encode(bb1, b1)

	bb2 := make([]byte, base64.StdEncoding.EncodedLen(len(b2)))
	base64.StdEncoding.Encode(bb2, b2)

	//var data *gojsontest.UnmarshalOneofNotHide
	data := &gojsontest.UnmarshalOneofNotHide{}

	cases := []*CaseDesc{
		{"t_string 1", []byte(`{"type": {"t_string": null} }`), "", func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TString)
			if ok {
				return tt, ok, tt.TString
			} else {
				return nil, ok, nil
			}
		}},
		{"t_string 2", []byte(`{"type": {"t_string": "Hello C"} }`), "Hello C", func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TString)
			if ok {
				return tt, ok, tt.TString
			} else {
				return nil, ok, nil
			}
		}},
		{"t_string 3", []byte(`{"type": {"t_string": null} }`), "", func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TString)
			if ok {
				return tt, ok, tt.TString
			} else {
				return nil, ok, nil
			}
		}},

		{"t_int32 1", []byte(`{"type": {"t_int32": 11} }`), int32(11), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TInt32)
			if ok {
				return tt, ok, tt.TInt32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_int64 1", []byte(`{"type": {"t_int64": 12} }`), int64(12), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TInt64)
			if ok {
				return tt, ok, tt.TInt64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_uint32 1", []byte(`{"type": {"t_uint32": 13} }`), uint32(13), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TUint32)
			if ok {
				return tt, ok, tt.TUint32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_uint64 1", []byte(`{"type": {"t_uint64": 14} }`), uint64(14), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TUint64)
			if ok {
				return tt, ok, tt.TUint64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_sint32 1", []byte(`{"type": {"t_sint32": 15} }`), int32(15), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TSint32)
			if ok {
				return tt, ok, tt.TSint32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_sint64 1", []byte(`{"type": {"t_sint64": 16} }`), int64(16), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TSint64)
			if ok {
				return tt, ok, tt.TSint64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_sfixed32 1", []byte(`{"type": {"t_sfixed32": 17} }`), int32(17), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TSfixed32)
			if ok {
				return tt, ok, tt.TSfixed32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_sfixed64 1", []byte(`{"type": {"t_sfixed64": 18} }`), int64(18), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TSfixed64)
			if ok {
				return tt, ok, tt.TSfixed64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_fixed32 1", []byte(`{"type": {"t_fixed32": 19} }`), uint32(19), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TFixed32)
			if ok {
				return tt, ok, tt.TFixed32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_fixed64 1", []byte(`{"type": {"t_fixed64": 20} }`), uint64(20), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TFixed64)
			if ok {
				return tt, ok, tt.TFixed64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_float 1", []byte(`{"type": {"t_float": 32.32} }`), float32(32.32), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TFloat)
			if ok {
				return tt, ok, tt.TFloat
			} else {
				return nil, ok, nil
			}
		}},

		{"t_double 1", []byte(`{"type": {"t_double": 64.64} }`), float64(64.64), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TDouble)
			if ok {
				return tt, ok, tt.TDouble
			} else {
				return nil, ok, nil
			}
		}},

		{"t_bool 1", []byte(`{"type": {"t_bool": true} }`), true, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TBool)
			if ok {
				return tt, ok, tt.TBool
			} else {
				return nil, ok, nil
			}
		}},

		{"t_enum1 1", []byte(`{"type": {"t_enum1": 1} }`), gojsontest.UnmarshalOneofNotHide_stopped, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TEnum1)
			if ok {
				return tt, ok, tt.TEnum1
			} else {
				return nil, ok, nil
			}
		}},

		{"t_enum2 1", []byte(`{"type": {"t_enum2": "stopped"} }`), gojsontest.UnmarshalOneofNotHide_stopped, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TEnum2)
			if ok {
				return tt, ok, tt.TEnum2
			} else {
				return nil, ok, nil
			}
		}},

		{"t_bytes 1", []byte(`{"type": {"t_bytes": null} }`), []byte(nil), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TBytes)
			if ok {
				return tt, ok, tt.TBytes
			} else {
				return nil, ok, nil
			}
		}},
		{"t_bytes 2", []byte(`{"type": {"t_bytes": ""} }`), []byte(""), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TBytes)
			if ok {
				return tt, ok, tt.TBytes
			} else {
				return nil, ok, nil
			}
		}},
		{"t_bytes 3", []byte(fmt.Sprintf(`{"type": {"t_bytes": "%s"} }`, bb1)), b1, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TBytes)
			if ok {
				return tt, ok, tt.TBytes
			} else {
				return nil, ok, nil
			}
		}},
		{"t_bytes 4", []byte(fmt.Sprintf(`{"type": {"t_bytes": "%s"} }`, bb2)), b2, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TBytes)
			if ok {
				return tt, ok, tt.TBytes
			} else {
				return nil, ok, nil
			}
		}},

		{"t_aliases 1", []byte(`{"type": {"t_aliases": null }}`), (*gojsontest.UnmarshalOneofNotHide_Aliases)(nil), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TAliases)
			if ok {
				return tt, ok, tt.TAliases
			} else {
				return nil, ok, nil
			}
		}},
		{"t_aliases 2", []byte(`{"type": {"t_aliases": {} }}`), &gojsontest.UnmarshalOneofNotHide_Aliases{}, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TAliases)
			if ok {
				return tt, ok, tt.TAliases
			} else {
				return nil, ok, nil
			}
		}},

		{"t_config 1", []byte(`{"type": {"t_config": null }}`), (*gojsontest.UnmarshalOneofNotHide_Config)(nil), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TConfig)
			if ok {
				return tt, ok, tt.TConfig
			} else {
				return nil, ok, nil
			}
		}},
		{"t_config 2", []byte(`{"type": {"t_config": {} }}`), &gojsontest.UnmarshalOneofNotHide_Config{}, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TConfig)
			if ok {
				return tt, ok, tt.TConfig
			} else {
				return nil, ok, nil
			}
		}},
		{"t_config 3", []byte(`{"type": {"t_config": {"ip": "192.1", "port": 8081} }}`), &gojsontest.UnmarshalOneofNotHide_Config{Ip: "192.1", Port: 8081}, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofNotHide_TConfig)
			if ok {
				return tt, ok, tt.TConfig
			} else {
				return nil, ok, nil
			}
		}},
	}

	{
		b := []byte(`{"type": { }}`)
		err := data.UnmarshalJSON(b)
		require.Nil(t, err)
		require.Nil(t, data.Type)
	}

	{
		b := []byte(`{"type": null}`)
		err := data.UnmarshalJSON(b)
		require.Nil(t, err)
		require.Nil(t, data.Type)
	}

	for _, c := range cases {
		data.Type = nil

		err := data.UnmarshalJSON(c.B)
		require.Nil(t, err, c.Name, err)
		require.NotNil(t, data.Type, c.Name)

		types, ok, v := c.Actual()
		require.True(t, ok, c.Name)
		require.NotNil(t, types, c.Name)
		require.Equal(t, c.Expected, v, c.Name)
	}
}

func Test_GoJSON_UnmarshalOneofHide_CheckCorrect(t *testing.T) {
	type CaseDesc struct {
		Name     string
		B        []byte
		Expected interface{}                             // Expected value.
		Actual   func() (interface{}, bool, interface{}) // return oneof type, assert ok, value.
	}

	b1 := []byte("Hello Bytes 1")
	b2 := []byte("Hello Bytes 2")

	bb1 := make([]byte, base64.StdEncoding.EncodedLen(len(b1)))
	base64.StdEncoding.Encode(bb1, b1)

	bb2 := make([]byte, base64.StdEncoding.EncodedLen(len(b2)))
	base64.StdEncoding.Encode(bb2, b2)

	//var data *gojsontest.UnmarshalOneofNotHide
	data := &gojsontest.UnmarshalOneofHide{}

	cases := []*CaseDesc{
		{"t_string 1", []byte(`{"t_string": null}`), "", func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TString)
			if ok {
				return tt, ok, tt.TString
			} else {
				return nil, ok, nil
			}
		}},
		{"t_string 2", []byte(`{"t_string": "Hello C"}`), "Hello C", func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TString)
			if ok {
				return tt, ok, tt.TString
			} else {
				return nil, ok, nil
			}
		}},
		{"t_string 3", []byte(`{"t_string": null}`), "", func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TString)
			if ok {
				return tt, ok, tt.TString
			} else {
				return nil, ok, nil
			}
		}},

		{"t_int32 1", []byte(`{"t_int32": 11}`), int32(11), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TInt32)
			if ok {
				return tt, ok, tt.TInt32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_int64 1", []byte(`{"t_int64": 12}`), int64(12), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TInt64)
			if ok {
				return tt, ok, tt.TInt64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_uint32 1", []byte(`{"t_uint32": 13}`), uint32(13), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TUint32)
			if ok {
				return tt, ok, tt.TUint32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_uint64 1", []byte(`{"t_uint64": 14}`), uint64(14), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TUint64)
			if ok {
				return tt, ok, tt.TUint64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_sint32 1", []byte(`{"t_sint32": 15}`), int32(15), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TSint32)
			if ok {
				return tt, ok, tt.TSint32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_sint64 1", []byte(`{"t_sint64": 16}`), int64(16), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TSint64)
			if ok {
				return tt, ok, tt.TSint64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_sfixed32 1", []byte(`{"t_sfixed32": 17}`), int32(17), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TSfixed32)
			if ok {
				return tt, ok, tt.TSfixed32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_sfixed64 1", []byte(`{"t_sfixed64": 18}`), int64(18), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TSfixed64)
			if ok {
				return tt, ok, tt.TSfixed64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_fixed32 1", []byte(`{"t_fixed32": 19}`), uint32(19), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TFixed32)
			if ok {
				return tt, ok, tt.TFixed32
			} else {
				return nil, ok, nil
			}
		}},

		{"t_fixed64 1", []byte(`{"t_fixed64": 20}`), uint64(20), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TFixed64)
			if ok {
				return tt, ok, tt.TFixed64
			} else {
				return nil, ok, nil
			}
		}},

		{"t_float 1", []byte(`{"t_float": 32.32}`), float32(32.32), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TFloat)
			if ok {
				return tt, ok, tt.TFloat
			} else {
				return nil, ok, nil
			}
		}},

		{"t_double 1", []byte(`{"t_double": 64.64}`), float64(64.64), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TDouble)
			if ok {
				return tt, ok, tt.TDouble
			} else {
				return nil, ok, nil
			}
		}},

		{"t_bool 1", []byte(`{"t_bool": true}`), true, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TBool)
			if ok {
				return tt, ok, tt.TBool
			} else {
				return nil, ok, nil
			}
		}},

		{"t_enum1 1", []byte(`{"t_enum1": 1}`), gojsontest.UnmarshalOneofHide_stopped, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TEnum1)
			if ok {
				return tt, ok, tt.TEnum1
			} else {
				return nil, ok, nil
			}
		}},

		{"t_enum2 1", []byte(`{"t_enum2": "stopped"}`), gojsontest.UnmarshalOneofHide_stopped, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TEnum2)
			if ok {
				return tt, ok, tt.TEnum2
			} else {
				return nil, ok, nil
			}
		}},

		{"t_bytes 1", []byte(`{"t_bytes": null}`), []byte(nil), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TBytes)
			if ok {
				return tt, ok, tt.TBytes
			} else {
				return nil, ok, nil
			}
		}},
		{"t_bytes 2", []byte(`{"t_bytes": ""}`), []byte(""), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TBytes)
			if ok {
				return tt, ok, tt.TBytes
			} else {
				return nil, ok, nil
			}
		}},
		{"t_bytes 3", []byte(fmt.Sprintf(`{"t_bytes": "%s"}`, bb1)), b1, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TBytes)
			if ok {
				return tt, ok, tt.TBytes
			} else {
				return nil, ok, nil
			}
		}},
		{"t_bytes 4", []byte(fmt.Sprintf(`{"t_bytes": "%s"}`, bb2)), b2, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TBytes)
			if ok {
				return tt, ok, tt.TBytes
			} else {
				return nil, ok, nil
			}
		}},

		{"t_aliases 1", []byte(`{"t_aliases": null }`), (*gojsontest.UnmarshalOneofHide_Aliases)(nil), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TAliases)
			if ok {
				return tt, ok, tt.TAliases
			} else {
				return nil, ok, nil
			}
		}},
		{"t_aliases 2", []byte(`{"t_aliases": {} }`), &gojsontest.UnmarshalOneofHide_Aliases{}, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TAliases)
			if ok {
				return tt, ok, tt.TAliases
			} else {
				return nil, ok, nil
			}
		}},

		{"t_config 1", []byte(`{"t_config": null }`), (*gojsontest.UnmarshalOneofHide_Config)(nil), func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TConfig)
			if ok {
				return tt, ok, tt.TConfig
			} else {
				return nil, ok, nil
			}
		}},
		{"t_config 2", []byte(`{"t_config": {} }`), &gojsontest.UnmarshalOneofHide_Config{}, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TConfig)
			if ok {
				return tt, ok, tt.TConfig
			} else {
				return nil, ok, nil
			}
		}},
		{"t_config 3", []byte(`{"t_config": {"ip": "192.1", "port": 8081} }`), &gojsontest.UnmarshalOneofHide_Config{Ip: "192.1", Port: 8081}, func() (interface{}, bool, interface{}) {
			tt, ok := data.Type.(*gojsontest.UnmarshalOneofHide_TConfig)
			if ok {
				return tt, ok, tt.TConfig
			} else {
				return nil, ok, nil
			}
		}},
	}

	{
		b := []byte(`{"type": { }}`)
		err := data.UnmarshalJSON(b)
		require.Nil(t, err)
		require.Nil(t, data.Type)
	}

	{
		b := []byte(`{"type": null}`)
		err := data.UnmarshalJSON(b)
		require.Nil(t, err)
		require.Nil(t, data.Type)
	}

	for _, c := range cases {
		data.Type = nil

		err := data.UnmarshalJSON(c.B)
		require.Nil(t, err, c.Name, err)
		require.NotNil(t, data.Type, c.Name)

		types, ok, v := c.Actual()
		require.True(t, ok, c.Name)
		require.NotNil(t, types, c.Name)
		require.Equal(t, c.Expected, v, c.Name)
	}
}

func Test_GoJSON_UnmarshalData_CheckError(t *testing.T) {
	type CaseDesc struct {
		Name     string
		B        []byte
		Expected string // expected error message.
	}

	type UnmarshalData gojsontest.UnmarshalData

	data1 := &gojsontest.UnmarshalData{}
	data2 := &UnmarshalData{}

	cases := []*CaseDesc{
		{"Unexpected end of JSON input", []byte(`{`), "unexpected end of JSON input"},
		{"Invalid beginning character 1", []byte(`}`), "invalid character '}' looking for beginning of value"},
		{"Invalid beginning character 2", []byte(`x}`), "invalid character 'x' looking for beginning of value"},

		{"Invalid JSON format 1", []byte(``), `unexpected end of JSON input`},
		{"Invalid JSON format 2", []byte(`xxx`), `invalid character 'x' looking for beginning of value`},
		{"Invalid JSON format 3", []byte(`"abcdefg"`), `json: cannot unmarshal "abcdefg" into object`},

		{"Invalid null format 1", []byte(`nul`), `invalid character ' ' in literal null (expecting 'l')`},
		{"Invalid null format 2", []byte(`"null"`), `json: cannot unmarshal "null" into object`},
		{"Invalid null format 3", []byte(`NULL`), `invalid character 'N' looking for beginning of value`},

		{"Invalid key format 1", []byte(`{x": "v"}`), "invalid character 'x' looking for beginning of object key string"},
		{"Invalid key format 2", []byte(`{"x: "v"}`), "invalid character 'v' after object key"},
		{"Invalid key format 3", []byte(`{1: "v"}`), `invalid character '1' looking for beginning of object key string`},

		{"Invalid field Separator 1", []byte(`{"x"| "v"}`), "invalid character '|' after object key"},
		{"Invalid field Separator 2", []byte(`{"x", "v"}`), "invalid character ',' after object key"},
		{"Invalid array Separator 2", []byte(`{"x": [1:2]}`), "invalid character ':' after array element"},

		{"Invalid t_string 1", []byte(`{"t_string": 1}`), "json: cannot unmarshal 1 into field t_string of type string"},
		{"Invalid t_string 2", []byte(`{"t_string": ["1"]}`), `json: cannot unmarshal ["1"] into field t_string of type string`},
		{"Invalid t_string 3", []byte(`{"t_string": {"k1": "v1"}}`), `json: cannot unmarshal {"k1": "v1"} into field t_string of type string`},
		{"Invalid t_string 4", []byte(`{"t_string": true}`), `json: cannot unmarshal true into field t_string of type string`},
		{"Invalid t_string 5", []byte(`{"t_string": nul}`), `invalid character '}' in literal null (expecting 'l')`},
		{"Invalid t_string 6", []byte(`{"t_string": s1"}`), "invalid character 's' looking for beginning of value"},
		{"Invalid t_string 7", []byte(`{"t_string": 1"}`), `invalid character '"' after object key:value pair`},

		{"Invalid t_int32 1", []byte(`{"t_int32": "name"}`), `json: cannot unmarshal "name" into field t_int32 of type int32`},
		{"Invalid t_int32 2", []byte(`{"t_int32": 1.23}`), `json: cannot unmarshal 1.23 into field t_int32 of type int32`},
		{"Invalid t_int32 3", []byte(`{"t_int32": true}`), `json: cannot unmarshal true into field t_int32 of type int32`},
		{"Invalid t_int32 4", []byte(`{"t_int32": "123"}`), `json: cannot unmarshal "123" into field t_int32 of type int32`},
		{"Invalid t_int64 1", []byte(`{"t_int64": "123"}`), `json: cannot unmarshal "123" into field t_int64 of type int64`},
		{"Invalid t_uint32 1", []byte(`{"t_uint32": "123"}`), `json: cannot unmarshal "123" into field t_uint32 of type uint32`},
		{"Invalid t_uint64 1", []byte(`{"t_uint64": "123"}`), `json: cannot unmarshal "123" into field t_uint64 of type uint64`},
		{"Invalid t_sfixed32 1", []byte(`{"t_sfixed32": "123"}`), `json: cannot unmarshal "123" into field t_sfixed32 of type int32`},
		{"Invalid t_sfixed64 1", []byte(`{"t_sfixed64": "123"}`), `json: cannot unmarshal "123" into field t_sfixed64 of type int64`},
		{"Invalid t_fixed32 1", []byte(`{"t_fixed32": "123"}`), `json: cannot unmarshal "123" into field t_fixed32 of type uint32`},
		{"Invalid t_fixed64 1", []byte(`{"t_fixed64": "123"}`), `json: cannot unmarshal "123" into field t_fixed64 of type uint64`},

		{"Invalid t_float 1", []byte(`{"t_float": "123"}`), `json: cannot unmarshal "123" into field t_float of type float32`},
		{"Invalid t_double 1", []byte(`{"t_double": "123"}`), `json: cannot unmarshal "123" into field t_double of type float64`},
		{"Invalid t_bool 1", []byte(`{"t_bool": "true"}`), `json: cannot unmarshal "true" into field t_bool of type bool`},
		{"Invalid t_bool 2", []byte(`{"t_bool": TRUE}`), `invalid character 'T' looking for beginning of value`},
		{"Invalid t_bool 3", []byte(`{"t_bool": 12}`), `json: cannot unmarshal 12 into field t_bool of type bool`},

		{"Invalid t_enum1 1", []byte(`{"t_enum1": "1"}`), `json: cannot unmarshal "1" into field t_enum1 of type UnmarshalData_Enum`},

		{"Invalid t_bytes 1", []byte(`{"t_bytes": true}`), `json: cannot unmarshal true into field t_bytes of type []byte`},
		{"Invalid t_bytes 2", []byte(`{"t_bytes": 123}`), `json: cannot unmarshal 123 into field t_bytes of type []byte`},
		{"Invalid t_bytes 3", []byte(`{"t_bytes": "abcdefg"}`), `json: cannot unmarshal "abcdefg" into field t_bytes of type []byte`},

		{"Invalid t_aliases 1", []byte(`{"t_aliases": "abcdefg"}`), `json: cannot unmarshal "abcdefg" into object`},
		{"Invalid t_aliases 2", []byte(`{"t_aliases": {{{ }}} }`), `invalid character '{' looking for beginning of object key string`},

		{"Invalid t_config 1", []byte(`{"t_config": "abcdefg"}`), `json: cannot unmarshal "abcdefg" into object`},
		{"Invalid t_config 2", []byte(`{"t_config": {{{ }}} }`), `invalid character '{' looking for beginning of object key string`},

		{"Invalid array_double 1", []byte(`{"array_double": ["x"]}`), `json: cannot unmarshal "x" as array element into field array_double of type []float64`},
		{"Invalid array_double 2", []byte(`{"array_double": 12, "x"]}`), `invalid character ']' after object key`},
		{"Invalid array_double 3", []byte(`{"array_double": [12, "x"}`), `invalid character '}' after array element`},

		{"Invalid array_double 4", []byte(`{"array_double": "x"}`), `json: cannot unmarshal "x" as array into field array_double of type []float64`},

		{"Invalid array_double 5", []byte(`{"array_double": 64.64}`), `json: cannot unmarshal 64.64 as array into field array_double of type []float64`},
		{"Invalid array_double 6", []byte(`{"array_double": "64.64"}`), `json: cannot unmarshal "64.64" as array into field array_double of type []float64`},
		{"Invalid array_double 7", []byte(`{"array_double": {"k1": "v1"}}`), `json: cannot unmarshal {"k1": "v1"} as array into field array_double of type []float64`},
		{"Invalid array_double 8", []byte(`{"array_double": true}`), `json: cannot unmarshal true as array into field array_double of type []float64`},
		{"Invalid array_double 9", []byte(`{"array_double":          "x"    }`), `json: cannot unmarshal "x" as array into field array_double of type []float64`},
		{"Invalid array_double 10", []byte(`{"array_double":  {   "k1": "v1" }}`), `json: cannot unmarshal {   "k1": "v1" } as array into field array_double of type []float64`},
		{"Invalid array_double 11", []byte(`{"array_double":  [64.11, 64.22,]`), `invalid character ']' looking for beginning of value`},
		{"Invalid array_double 12", []byte(`{"array_double":  [, 64.11, 64.22]`), `invalid character ',' looking for beginning of value`},

		{"Invalid array_float 1", []byte(`{"array_float": ["x"]}`), `json: cannot unmarshal "x" as array element into field array_float of type []float32`},
		{"Invalid array_int32 1", []byte(`{"array_int32": ["x"]}`), `json: cannot unmarshal "x" as array element into field array_int32 of type []int32`},
		{"Invalid array_int32 2", []byte(`{"array_int32": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_int32 of type []int32`},
		{"Invalid array_int32 3", []byte(`{"array_int32": [true]}`), `json: cannot unmarshal true as array element into field array_int32 of type []int32`},
		{"Invalid array_int64 1", []byte(`{"array_int64": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_int64 of type []int64`},
		{"Invalid array_uint32 1", []byte(`{"array_uint32": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_uint32 of type []uint32`},
		{"Invalid array_uint64 1", []byte(`{"array_uint64": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_uint64 of type []uint64`},
		{"Invalid array_sint32 1", []byte(`{"array_sint32": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_sint32 of type []int32`},
		{"Invalid array_sint64 1", []byte(`{"array_sint64": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_sint64 of type []int64`},
		{"Invalid array_sfixed32 1", []byte(`{"array_sfixed32": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_sfixed32 of type []int32`},
		{"Invalid array_sfixed64 1", []byte(`{"array_sfixed64": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_sfixed64 of type []int64`},
		{"Invalid array_fixed32 1", []byte(`{"array_fixed32": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_fixed32 of type []uint32`},
		{"Invalid array_fixed64 1", []byte(`{"array_fixed64": [32.1]}`), `json: cannot unmarshal 32.1 as array element into field array_fixed64 of type []uint64`},

		{"Invalid array_bool 1", []byte(`{"array_bool": "true"}`), `json: cannot unmarshal "true" as array into field array_bool of type []bool`},
		{"Invalid array_bool 2", []byte(`{"array_bool": true}`), `json: cannot unmarshal true as array into field array_bool of type []bool`},
		{"Invalid array_bool 3", []byte(`{"array_bool": ["false"]}`), `json: cannot unmarshal "false" as array element into field array_bool of type []bool`},
		{"Invalid array_bool 4", []byte(`{"array_bool": {false}}`), `invalid character 'f' looking for beginning of object key string`},

		{"Invalid array_string 1", []byte(`{"array_string": [123]}`), `json: cannot unmarshal 123 as array element into field array_string of type []string`},

		{"Invalid array_enum1 1", []byte(`{"array_enum1": ["1"]}`), `json: cannot unmarshal "1" as array element into field array_enum1 of type []UnmarshalData_Enum`},

		{"Invalid array_bytes 1", []byte(`{"array_bytes": ["xxx", "yyy"]}`), `json: cannot unmarshal "xxx" as array element into field array_bytes of type [][]byte`},

		{"Invalid map_int32_double 1", []byte(`{"map_int32_double": {"k1": "123.3"}}`), `json: cannot unmarshal k1 as map key into field map_int32_double of type map[int32]float64`},
		{"Invalid map_int32_double 2", []byte(`{"map_int32_double": {32: "123.3"}}`), `invalid character '3' looking for beginning of object key string`},
		{"Invalid map_int32_double 3", []byte(`{"map_int32_double": {"32": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_double of type map[int32]float64`},
		{"Invalid map_int32_double 4", []byte(`{"map_int32_double": "null"}`), `json: cannot unmarshal "null" as map into field map_int32_double of type map[int32]float64`},
		{"Invalid map_int32_double 5", []byte(`{"map_int32_double": xxx}`), `invalid character 'x' looking for beginning of value`},
		{"Invalid map_int32_double 6", []byte(`{"map_int32_double": true}`), `json: cannot unmarshal true as map into field map_int32_double of type map[int32]float64`},
		{"Invalid map_int32_double 7", []byte(`{"map_int32_double": 123}`), `json: cannot unmarshal 123 as map into field map_int32_double of type map[int32]float64`},
		{"Invalid map_int32_double 8", []byte(`{"map_int32_double": "xxx"}`), `json: cannot unmarshal "xxx" as map into field map_int32_double of type map[int32]float64`},
		{"Invalid map_int32_double 9", []byte(`{"map_int32_double": [123]}`), `json: cannot unmarshal [123] as map into field map_int32_double of type map[int32]float64`},
		{"Invalid map_int32_double 10", []byte(`{"map_int32_double": {"32": 64.1,}}`), `invalid character '}' looking for beginning of object key string`},

		{"Invalid map_int32_float 1", []byte(`{"map_int32_float": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_float of type map[int32]float32`},
		{"Invalid map_int32_int32 1", []byte(`{"map_int32_int32": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_int32 of type map[int32]int32`},
		{"Invalid map_int32_int64 1", []byte(`{"map_int32_int64": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_int64 of type map[int32]int64`},
		{"Invalid map_int32_uint32 1", []byte(`{"map_int32_uint32": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_uint32 of type map[int32]uint32`},
		{"Invalid map_int32_uint64 1", []byte(`{"map_int32_uint64": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_uint64 of type map[int32]uint64`},
		{"Invalid map_int32_sint32 1", []byte(`{"map_int32_sint32": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_sint32 of type map[int32]int32`},
		{"Invalid map_int32_sint64 1", []byte(`{"map_int32_sint64": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_sint64 of type map[int32]int64`},
		{"Invalid map_int32_sfixed32 1", []byte(`{"map_int32_sfixed32": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_sfixed32 of type map[int32]int32`},
		{"Invalid map_int32_sfixed64 1", []byte(`{"map_int32_sfixed64": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_sfixed64 of type map[int32]int64`},
		{"Invalid map_int32_fixed32 1", []byte(`{"map_int32_fixed32": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_fixed32 of type map[int32]uint32`},
		{"Invalid map_int32_fixed64 1", []byte(`{"map_int32_fixed64": {"123": "123.3"}}`), `json: cannot unmarshal "123.3" as map value into field map_int32_fixed64 of type map[int32]uint64`},

		{"Invalid map_int32_bool 1", []byte(`{"map_int32_bool": {"123": "true"}}`), `json: cannot unmarshal "true" as map value into field map_int32_bool of type map[int32]bool`},
		{"Invalid map_int32_string 1", []byte(`{"map_int32_string": {"123": 123}}`), `json: cannot unmarshal 123 as map value into field map_int32_string of type map[int32]string`},

		{"Invalid map_int32_enum1 1", []byte(`{"map_int32_enum1": {"32": "1"}}`), `json: cannot unmarshal "1" as map value into field map_int32_enum1 of type map[int32]UnmarshalData_Enum`},
		{"Invalid map_int32_bytes 1", []byte(`{"map_int32_bytes": {"32": "xxx"}}`), `json: cannot unmarshal "xxx" as map value into field map_int32_bytes of type map[int32][]byte`},

		{"Invalid map_int64_int32 1", []byte(`{"map_int64_int32": {"xx": 32}}`), `json: cannot unmarshal xx as map key into field map_int64_int32 of type map[int64]int32`},
		{"Invalid map_int64_int32 2", []byte(`{"map_int64_int32": {"123.23": 32}}`), `json: cannot unmarshal 123.23 as map key into field map_int64_int32 of type map[int64]int32`},
		{"Invalid map_int64_int32 2", []byte(`{"map_int64_int32": {32: 32}}`), `invalid character '3' looking for beginning of object key string`},

		{"Invalid map_uint32_int32 1", []byte(`{"map_uint32_int32": {"xx": 32}}`), `json: cannot unmarshal xx as map key into field map_uint32_int32 of type map[uint32]int32`},
		{"Invalid map_uint64_int32 1", []byte(`{"map_uint64_int32": {"xx": 32}}`), `json: cannot unmarshal xx as map key into field map_uint64_int32 of type map[uint64]int32`},
		{"Invalid map_sint32_int32 1", []byte(`{"map_sint32_int32": {"xx": 32}}`), `json: cannot unmarshal xx as map key into field map_sint32_int32 of type map[int32]int32`},
		{"Invalid map_sint64_int32 1", []byte(`{"map_sint64_int32": {"xx": 32}}`), `json: cannot unmarshal xx as map key into field map_sint64_int32 of type map[int64]int32`},
		{"Invalid map_fixed32_int32 1", []byte(`{"map_fixed32_int32": {"xx": 32}}`), `json: cannot unmarshal xx as map key into field map_fixed32_int32 of type map[uint32]int32`},
		{"Invalid map_fixed64_int32 1", []byte(`{"map_fixed64_int32": {"xx": 32}}`), `json: cannot unmarshal xx as map key into field map_fixed64_int32 of type map[uint64]int32`},
		{"Invalid map_sfixed32_int32 1", []byte(`{"map_sfixed32_int32": {"xx": 32}}`), `json: cannot unmarshal xx as map key into field map_sfixed32_int32 of type map[int32]int32`},
		{"Invalid map_sfixed64_int32 1", []byte(`{"map_sfixed64_int32": {"xx": 32}}`), `json: cannot unmarshal xx as map key into field map_sfixed64_int32 of type map[int64]int32`},

		{"Invalid map_string_int32 1", []byte(`{"map_string_int32": {xx: 32}}`), `invalid character 'x' looking for beginning of object key string`},
	}

	// standard json
	for _, c := range cases {
		err := json.Unmarshal(c.B, data2)
		require.NotNil(t, err, c.Name)
		//fmt.Println(c.Name, ":", err.Error())
	}

	fmt.Println("------ split line ------")

	// gojson
	for _, c := range cases {
		err := data1.UnmarshalJSON(c.B)
		require.NotNil(t, err, c.Name)

		//fmt.Println(c.Name, ":", err.Error())
		require.Equal(t, c.Expected, err.Error(), c.Name)
	}

	enumCases := []*CaseDesc{
		{"Invalid t_enum1 2", []byte(`{"t_enum1": 2}`), `json: unknown enum value 2 in field t_enum1`},

		{"Invalid t_enum2 1", []byte(`{"t_enum2": 1}`), `json: cannot unmarshal 1 into field t_enum2 of type UnmarshalData_Enum`},
		{"Invalid t_enum2 2", []byte(`{"t_enum2": "xxxx"}`), `json: unknown enum value "xxxx" in field t_enum2`},

		{"Invalid array_enum1 2", []byte(`{"array_enum1": [3]}`), `json: unknown enum value 3 in field array_enum1`},

		{"Invalid array_enum2 1", []byte(`{"array_enum2": [1]}`), `json: cannot unmarshal 1 as array element into field array_enum2 of type []UnmarshalData_Enum`},
		{"Invalid array_enum2 2", []byte(`{"array_enum2": ["xxx"]}`), `json: unknown enum value "xxx" in field array_enum2`},

		{"Invalid map_int32_enum1 2", []byte(`{"map_int32_enum1": {"32": 3}}`), `json: unknown enum value 3 in field map_int32_enum1`},

		{"Invalid map_int32_enum2 1", []byte(`{"map_int32_enum2": {"32": "xxx"}}`), `json: unknown enum value "xxx" in field map_int32_enum2`},
		{"Invalid map_int32_enum2 2", []byte(`{"map_int32_enum2": {"32": 1}}`), `json: cannot unmarshal 1 as map value into field map_int32_enum2 of type map[int32]UnmarshalData_Enum`},
	}

	// gojson
	for _, c := range enumCases {
		err := data1.UnmarshalJSON(c.B)
		require.NotNil(t, err, c.Name)

		//fmt.Println(c.Name, ":", err.Error())
		require.Equal(t, c.Expected, err.Error(), c.Name)
	}
}

func Test_GoJSON_UnmarshalOneofNotHide_CheckError(t *testing.T) {
	type CaseDesc struct {
		Name     string
		B        []byte
		Expected string // expected error message.
	}

	data1 := &gojsontest.UnmarshalOneofNotHide{}

	cases := []*CaseDesc{
		//{"unknown", []byte(`{"type": { "t_unknown": 1 } }`), `json: unknown oneof field t_unknown`},

		{"Invalid t_string 1", []byte(`{"type": { "t_string": 1 } }`), `json: cannot unmarshal 1 into field type of type string`},
		{"Invalid t_string 8", []byte(`{"type": { "t_string": "1", "t_int32": 1 }}`), `json: unmarshal: the field type is type oneof, allow contains only one`},

		//{"NULL", []byte(``), `unexpected end of JSON input`},
	}

	for _, c := range cases {
		err := data1.UnmarshalJSON(c.B)
		require.NotNil(t, err, c.Name)

		//fmt.Println(c.Name, ":", err.Error())
		require.Equal(t, c.Expected, err.Error(), c.Name)
	}
}

func Test_GoJSON_UnmarshalOneofHide_CheckError(t *testing.T) {
	type CaseDesc struct {
		Name     string
		B        []byte
		Expected string // expected error message.
	}

	data1 := &gojsontest.UnmarshalOneofHide{}

	cases := []*CaseDesc{
		{"Invalid t_string 1", []byte(`{ "t_string": 1 }`), `json: cannot unmarshal 1 into field t_string of type string`},
		{"Invalid t_string 8", []byte(`{ "t_string": "1", "t_int32": 1 }`), `json: unmarshal: the field t_int32 is type oneof, allow contains only one`},

		//{"NULL", []byte(``), `unexpected end of JSON input`},
	}

	for _, c := range cases {
		err := data1.UnmarshalJSON(c.B)
		require.NotNil(t, err, c.Name)

		//fmt.Println(c.Name, ":", err.Error())
		require.Equal(t, c.Expected, err.Error(), c.Name)
	}
}

func Test_GoJSON_OptionalModel1(t *testing.T) {
	type OptionalModel1 gojsontest.OptionalModel1

	t32 := int32(0)
	t64 := int64(1)
	tBytes := make([]byte, 0)
	expected := []byte(`{"t_int32":0,"t_int64":1}`)

	data1 := &OptionalModel1{
		TString:   nil,
		TInt32:    &t32,
		TInt64:    &t64,
		TUint32:   nil,
		TUint64:   nil,
		TSint32:   nil,
		TSint64:   nil,
		TSfixed32: nil,
		TSfixed64: nil,
		TFixed32:  nil,
		TFixed64:  nil,
		TFloat:    nil,
		TDouble:   nil,
		TBool:     nil,
		TEnum1:    nil,
		TEnum2:    nil,
		TBytes:    tBytes,
		TAliases:  nil,
		TConfig:   nil,
	}
	b1, err := json.Marshal(data1)
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, b1, expected)

	data3 := &OptionalModel1{}
	err = json.Unmarshal(b1, data3)
	require.Nil(t, err)
	data1.TBytes = nil
	require.Equal(t, data1, data3)

	data2 := &gojsontest.OptionalModel1{
		TString:   nil,
		TInt32:    &t32,
		TInt64:    &t64,
		TUint32:   nil,
		TUint64:   nil,
		TSint32:   nil,
		TSint64:   nil,
		TSfixed32: nil,
		TSfixed64: nil,
		TFixed32:  nil,
		TFixed64:  nil,
		TFloat:    nil,
		TDouble:   nil,
		TBool:     nil,
		TEnum1:    nil,
		TEnum2:    nil,
		TBytes:    tBytes,
		TAliases:  nil,
		TConfig:   nil,
	}

	b2, err := data2.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b2))
	require.Equal(t, b2, expected)

	data4 := &gojsontest.OptionalModel1{}
	err = json.Unmarshal(b2, data4)
	require.Nil(t, err)
	data2.TBytes = nil
	require.Equal(t, data2, data4)
}

func Test_GoJSON_OptionalModel2(t *testing.T) {
	tString := "s1"
	tInt32 := int32(1)
	tInt64 := int64(2)
	tUint32 := uint32(3)
	tUint64 := uint64(4)
	tSint32 := int32(5)
	tSint64 := int64(6)
	tSfixed32 := int32(7)
	tSfixed64 := int64(8)
	tFixed32 := uint32(9)
	tFixed64 := uint64(10)
	tFloat := float32(11.11)
	tDouble := float64(12.12)
	tBool := true
	tEnum1 := gojsontest.OptionalModel2_running
	tEnum2 := gojsontest.OptionalModel2_stopped

	ip := "127.0.0.1"
	port := int32(80)

	data1 := &gojsontest.OptionalModel2{
		TString:   &tString,
		TInt32:    &tInt32,
		TInt64:    &tInt64,
		TUint32:   &tUint32,
		TUint64:   &tUint64,
		TSint32:   &tSint32,
		TSint64:   &tSint64,
		TSfixed32: &tSfixed32,
		TSfixed64: &tSfixed64,
		TFixed32:  &tFixed32,
		TFixed64:  &tFixed64,
		TFloat:    &tFloat,
		TDouble:   &tDouble,
		TBool:     &tBool,
		TEnum1:    &tEnum1,
		TEnum2:    &tEnum2,
		TBytes:    []byte("b1"),
		TAliases:  &gojsontest.OptionalModel2_Aliases{},
		TConfig: &gojsontest.OptionalModel2_Config{
			Ip:   &ip,
			Port: &port,
		},
	}

	expected := []byte(`{"t_string":"s1","t_int32":1,"t_int64":2,"t_uint32":3,"t_uint64":4,"t_sint32":5,"t_sint64":6,"t_sfixed32":7,"t_sfixed64":8,"t_fixed32":9,"t_fixed64":10,"t_float":11.11,"t_double":12.12,"t_bool":true,"t_enum1":0,"t_enum2":"stopped","t_bytes":"YjE=","t_aliases":{},"t_config":{"ip":"127.0.0.1","port":80}}`)

	b1, err := data1.MarshalJSON()
	require.Nil(t, err)
	//fmt.Println(string(b1))
	require.Equal(t, b1, expected)

	data2 := &gojsontest.OptionalModel2{}
	err = data2.UnmarshalJSON(b1)
	require.Nil(t, err)
	require.Equal(t, data1, data2)
}
