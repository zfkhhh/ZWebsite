package utils

import "fmt"

// error 封装处理

type Err struct {
	Pkg       string   // 标识抛出错误的包
	Info      string   // 描述错误
	Prev       error    // 包装error
}

func (e *Err) Error() string {
	if e.Prev == nil {
		return fmt.Sprintf("%s: %s", e.Pkg, e.Info)
	}
	return fmt.Sprintf("%s: %s\n%v", e.Pkg, e.Info, e.Prev)
}

//// 包装error
//func MakeErr(err error, format string, args ...interface{}) *Err {
//	if len(args) > 0 {
//		return &Err{
//			Pkg:  Pkg,
//			Info: fmt.Sprintf(format, args...),
//			Prev: err,
//		}
//	}
//	return &Err{
//		Pkg:  Pkg,
//		Info: format,
//		Prev: err,
//	}
//}
//
//// 检查error是否为空，不为空包装出Err结构抛出
//func CheckErr(err error, format string, args ...interface{}) {
//	if err != nil {
//		panic(MakeErr(err, format, args...))
//	}
//}
//
//// 抛出error后在边界recover
//func CatchErr(err *error) {
//	if p := recover(); p != nil {
//		if e, ok := p.(error); ok {
//			*err = e
//		} else {
//			panic(p)
//		}
//	}
//}
//
//func OriginErr(e error) error {
//	var ret error = e
//	for err, ok := ret.(*Err); ok && err.Prev != nil; err, ok = ret.(*Err) {
//		ret = err.Prev
//	}
//	return ret
//}
