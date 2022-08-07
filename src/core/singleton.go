package core

import (
	"estj/src/exception"
	"reflect"
)

// Generic 연습..사용 안함.
// 아래와 같은 struct 및 variable을 기본으로 가지고 있어야 함.
// singleton 객체값(pointer)
// var instance *singleton

// singleton구현을 원하는 struct
//type singleton struct{
//}

// singleton 객체를 요청에 따라 반환하는 메소드
// nil일 경우 new()로 생성하고 instance에 담고 생성 실패 시, error 문구 반환.
func SetInstance[T any](instance *T) error {
	if instance == nil {
		instance = new(T)
	}

	if instance == nil {
		return exception.CreateInstanceCreationFailed(reflect.TypeOf(instance).String(), "")
	}
	return nil
}
