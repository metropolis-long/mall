package testdb

import (
	"fmt"
	"mall/user/rpc/internal/data"
	"testing"
)

func Testdb() {
	fmt.Println("test db ...")
}
func TestAddLike(t *testing.T) {
	like := data.Like{}
	// if err := like.Add("1.1.1.1", "33er", "333"); err != nil {
	// 	t.Errorf("test fail %s",err)
	// }
	like.SumLikeByPerson()
	fmt.Println("end ")
}
