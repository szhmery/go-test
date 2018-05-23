package src

import (
"log"
"fmt"
// 辅助库
"github.com/golang/protobuf/proto"

// prototest.pb.go 的路径
"prototest"
)

func main() {
	// 创建一个消息 Test
	fmt.Print("test")
	test := &prototest.Test{
		// 使用辅助函数设置域的值
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &prototest.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}

	fmt.Printf("Label:%s Type:%d\n", test.GetLabel(), test.GetType())

	*(test.Label) = "hello go"
	*(test.Type) = 18

	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("Binary Len:%d\n", len(data))

	// 进行解码
	newTest := &prototest.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Printf("Label:%s Type:%d\n", test.GetLabel(), test.GetType())

	// 测试结果
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
}

