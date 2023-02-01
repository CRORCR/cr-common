package g_reflect

import (
	"errors"
	"fmt"
	"testing"
)

type Order struct {
	UserId             uint64  `db:"user_id" json:"user_id"`
	Type               uint64  `db:"type" json:"type"`
	OrderNumber        string  `db:"order_number" json:"order_number"`
	Price              float64 `db:"price" json:"price"`
	Amount             float64 `db:"amount" json:"amount"`
	TransferMemo       string  `db:"tranfer_memo" json:"tranfer_memo"`
	Status             uint64  `db:"status" json:"status"`
	PayerId            uint64  `db:"payer_id" json:"payer_id"`
	PayerName          string  `db:"payer_name" json:"payer_name"`
	PayerBankName      string  `db:"payer_bank_name" json:"payer_bank_name"`
	PayerBranchName    string  `db:"payer_branch_name" json:"payer_branch_name"`
	PayerBankNumber    string  `db:"payer_bank_number" json:"payer_bank_number"`
	ReceiverId         string  `db:"receiver_id" json:"receiver_id"`
	ReceiverName       string  `db:"receiver_name" json:"receiver_name"`
	ReceiverBankName   string  `db:"receiver_bank_name" json:"receiver_bank_name"`
	ReceiverBranchName string  `db:"receiver_branch_name" json:"receiver_branch_name"`
	ReceiverBankNumber string  `db:"receiver_bank_number" json:"receiver_bank_number"`
	RejectReason       *string `db:"reject_reason" json:"reject_reason"`
	PayCoinAt          *string `db:"pay_coin_at" json:"pay_coin_at"`
	ReceiveCoinAt      *string `db:"receive_coin_at" json:"receive_coin_at"`
	RejectAt           *string `db:"reject_at" json:"reject_at"`
	FreezeUid          *string `db:"freeze_uid" json:"freeze_uid"`
	ExpiredAt          *string `db:"expired_at" json:"expired_at"`

	BaseModel
}

type BaseModel struct {
	Id        uint64 `db:"id" json:"id"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func TestReflectClass_ToInt64(t *testing.T) {
	a := `1222222`
	if Reflect.ToInt64(a) != 1222222 {
		t.Error()
	}

	var b int = 12
	if Reflect.ToInt64(b) != 12 {
		t.Error()
	}

	var c int8 = 12
	if Reflect.ToInt64(c) != 12 {
		t.Error()
	}

	var d int16 = 12
	if Reflect.ToInt64(d) != 12 {
		t.Error()
	}

	var f int32 = 12
	if Reflect.ToInt64(f) != 12 {
		t.Error()
	}

	var g uint8 = 12
	if Reflect.ToInt64(g) != 12 {
		t.Error()
	}

	var h uint16 = 12
	if Reflect.ToInt64(h) != 12 {
		t.Error()
	}

	var i uint32 = 12
	if Reflect.ToInt64(i) != 12 {
		t.Error()
	}

	var j uint64 = 12
	if Reflect.ToInt64(j) != 12 {
		t.Error()
	}

	var k float32 = 12
	if Reflect.ToInt64(k) != 12 {
		t.Error()
	}

	var l float64 = 12
	if Reflect.ToInt64(l) != 12 {
		t.Error()
	}

	var m bool = true
	if Reflect.ToInt64(m) != 1 {
		t.Error()
	}
}

func TestReflectClass_ToUint64(t *testing.T) {
	map_ := map[string]interface{}{
		`haha`: 342842799,
	}
	fmt.Println(fmt.Sprintf(`%v`, map_))

	a := `1222222`
	if Reflect.ToUint64(a) != 1222222 {
		t.Error()
	}

	var b int = 12
	if Reflect.ToUint64(b) != 12 {
		t.Error()
	}

	var c int8 = 12
	if Reflect.ToUint64(c) != 12 {
		t.Error()
	}

	var d int16 = 12
	if Reflect.ToUint64(d) != 12 {
		t.Error()
	}

	var f int32 = 12
	if Reflect.ToUint64(f) != 12 {
		t.Error()
	}

	var g uint8 = 12
	if Reflect.ToUint64(g) != 12 {
		t.Error()
	}

	var h uint16 = 12
	if Reflect.ToUint64(h) != 12 {
		t.Error()
	}

	var i uint32 = 12
	if Reflect.ToUint64(i) != 12 {
		t.Error()
	}

	var j uint64 = 12
	if Reflect.ToUint64(j) != 12 {
		t.Error()
	}

	var k float32 = 12
	if Reflect.ToUint64(k) != 12 {
		t.Error()
	}

	var l float64 = 12
	if Reflect.ToUint64(l) != 12 {
		t.Error()
	}

	var m bool = true
	if Reflect.ToUint64(m) != 1 {
		t.Error()
	}
}

func TestReflectClass_ToString(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(Reflect.ToString(err))
		}
	}()

	var a *float64
	b := 0.34
	a = &b
	if Reflect.ToString(a) != `0.34` {
		t.Error()
	}

	type BType struct {
		B1 int
	}

	a1 := struct {
		A string
		B BType
	}{`1`, BType{2}}
	fmt.Println(Reflect.ToString(a1))

	type Test struct {
		Test1 *string `json:"test1"`
	}
	test := Test{}
	fmt.Println(Reflect.ToString(test.Test1))

	fmt.Println(Reflect.ToString(errors.New(`111`)))

	var c interface{} = `1`
	fmt.Println(c.(int))
}

func TestReflectClass_GetValuesInTagFromStruct(t *testing.T) {
	test := []*Order{}
	fields := Reflect.GetValuesInTagFromStruct(test, `json`)
	fmt.Println(fields)

	// Test{}
	//test1 := C2cOrder{}
	//fields = Reflect.GetValuesInTagFromStruct(test1, `json`)
	//if len(fields) != 23 {
	//	t.Error()
	//}

	// *Test{}
	//test2 := C2cOrder{}
	//fields = Reflect.GetValuesInTagFromStruct(&test2, `json`)
	//if len(fields) != 23 {
	//	t.Error()
	//}

	// []Test{}
	//test3 := []C2cOrder{}
	//fields = Reflect.GetValuesInTagFromStruct(test3, `json`)
	//if len(fields) != 23 {
	//	t.Error()
	//}

	// *[]Test{}
	//test4 := []C2cOrder{}
	//fields = Reflect.GetValuesInTagFromStruct(&test4, `json`)
	//if len(fields) != 23 {
	//	t.Error()
	//}
}

func TestReflectClass_ToBool(t *testing.T) {
	a := `true`
	fmt.Println(Reflect.ToBool(a))
}
