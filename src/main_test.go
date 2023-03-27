package main

import (
	//GOPATH模式下，依赖包存储在$GOPATH/src，该目录下只保存特定依赖包的一个版本，
	//而在GOMODULE模式下，依赖包存储在$GOPATH/pkg/mod，该目录中可以存储特定依赖包的多个版本。
	//gomod通过export GO111MODULE=on开启
	/*
		在项目路径下git mod init hello初始化go.mod
		然后再go mod tidy，引用项目需要的依赖增加到go.mod文件，去掉go.mod文件中项目不需要的依赖。
	*/
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/*
跑单测的命令：go test main.go interfaceTest.go main_test.go --test.run TestRect
*/
// 对Rect实现的Shape接口做单测
func TestRect(t *testing.T) {
	var s Shape = &Rect{width: 5, height: 5}
	if s.Area() != 25 {
		t.Errorf("Area error %f\n", s.Area())
	}
	if s.Perimeter() != 20 {
		t.Errorf("Perimeter error %f\n", s.Perimeter())
	}
}

func TestRectByConvey(t *testing.T) {
	convey.Convey("TestRectByConvey", t, func() {
		var s Shape = &Rect{height: 10, width: 8}
		convey.So(s.Area(), convey.ShouldEqual, 80)
		convey.So(s.Perimeter(), convey.ShouldEqual, 36)
	})
}
