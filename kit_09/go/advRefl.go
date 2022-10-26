package main

//type t1 int
//type t2 int
//
//type a0 struct {
//	X    int
//	Y    float64
//	Text string
//}
//
//func (a1 a0) compareStruct(a2 a) bool {
//	r1 := reflect.ValueOf(&a1).Elem()
//	r2 := reflect.ValueOf(&a2).Elem()
//
//	for i := 0; i < r1.NumField(); i++ {
//		if r1.Field(i).Interface() != r2.Field(i).Interface() {
//			return false
//		}
//	}
//	return true
//}
//
//func printMethods(i interface{}) {
//	r := reflect.ValueOf(i)
//	t := r.Type()
//	fmt.Printf("Type to examine: %s\n", t)
//
//	for j := 0; j < r.NumMethod(); j++ {
//		m := r.Method(j).Type()
//		fmt.Println(t.Method(j).Name, "-->", m)
//	}
//
//}