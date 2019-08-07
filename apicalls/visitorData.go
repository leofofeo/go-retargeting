package apicalls

import "fmt"

type visitorData []byte

func (vd visitorData) PrintByteSliceSize() {
	fmt.Println("Visitor data byte slice size is:", len(vd))
}
