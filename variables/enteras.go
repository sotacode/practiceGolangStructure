package variables

import "fmt"

func MuestroEnteros() {
	//Por declaracion se usa Var
	var intComun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun = ", intComun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
