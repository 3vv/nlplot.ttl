// h 20181225
//
// Convertor for translator

package translator

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
	"unsafe"
)

func Bytes2String(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2Result(b []byte) Result {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*Result)(unsafe.Pointer(&sh))
}

func Result2Bytes(r interface{}) []byte {
	buf := new(bytes.Buffer)
	_ = writeValue(buf /*reflect.Indirect(*/, reflect.ValueOf(r) /*)*/)
	return buf.Bytes()
}

// writeValue writes value
func writeValue(w io.Writer, v reflect.Value) (err error) {
	buf := &bytes.Buffer{}
	n := 0
	switch v.Kind() {
	case reflect.Struct:
		n = v.NumField()
	case reflect.String:
		buf.WriteString(v.String())
	default:
	}
	for i := 0; i < n; i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.Struct:
			err = writeValue(buf, v.Field(i))
			if err != nil {
				break
			}
		case reflect.String:
			_, err = buf.WriteString(v.Field(i).String())
			if err != nil {
				break
			}
		default:
			err = binary.Write(buf, binary.LittleEndian, v.Field(i).Interface())
			if err != nil {
				break
			}
		}
	}
	_, err = w.Write(buf.Bytes())
	return err
}
