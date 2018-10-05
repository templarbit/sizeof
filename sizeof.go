package sizeof

import (
	"reflect"
)

func SizeOf(obj interface{}) uint64 {
	t := reflect.TypeOf(obj)
	s := uint64(t.Size())
	return s + sizeOf(obj)
}

func sizeOf(obj interface{}) uint64 {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	s := uint64(0)

	switch t.Kind() {
	case reflect.Ptr:
		if !v.IsNil() && v.Elem().CanInterface() {
			s += uint64(reflect.TypeOf(v.Elem().Interface()).Size())
			s += sizeOf(v.Elem().Interface())
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			switch f.Kind() {
			case reflect.Ptr:
				if !f.IsNil() && f.Elem().CanInterface() {
					s += uint64(reflect.TypeOf(f.Elem().Interface()).Size())
					s += sizeOf(f.Elem().Interface())
				}
			case reflect.String:
				s += uint64(len(f.String()))
			}
		}
	case reflect.Map:
		s += 140 // aprox global overhead
		for _, b := range v.MapKeys() {
			s += 24 // aprox per key overhead
			switch b.Kind() {
			case reflect.Ptr:
				if !b.IsNil() && b.Elem().CanInterface() {
					s += uint64(reflect.TypeOf(b.Elem().Interface()).Size())
					s += 8 + sizeOf(b.Elem().Interface())
				}

			case reflect.String:
				s += 16 + uint64(len(b.String()))
			case reflect.Struct:
				s += uint64(reflect.TypeOf(b.Interface()).Size()) + sizeOf(b.Interface())
			case reflect.Array:
				s += uint64(reflect.TypeOf(b.Interface()).Size())
				for i := 0; i < b.Len(); i++ {
					f := b.Index(i)
					s += sizeOf(f.Interface())
				}
			}

			s += uint64(reflect.TypeOf(v.MapIndex(b).Interface()).Size())
			s += sizeOf(v.MapIndex(b).Interface())
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {

			f := v.Index(i)
			s += uint64(reflect.TypeOf(f.Interface()).Size())
			s += sizeOf(f.Interface())
		}
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			f := v.Index(i)
			s += sizeOf(f.Interface())
		}
	case reflect.String:
		s += uint64(len(v.Interface().(string)))
	}

	return s
}
