package gocalc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var num int = 200

// TestMain() adalah sebuah fungsi spesial pada test yang akan dijalankan pertama kali ketika test dijalankan,
// biasanya digunakan untuk membuat instansi atau mengisi nilai variabel juga menyiapakan service-service yang
// mungkin akan digunakan oleh fungsi-fungsi test.
func TestMain(m *testing.M) {
	fmt.Println("Start test")
	num = 100
	m.Run()                    // m.Run() menjalankan semua fungsi test
	fmt.Println("Finish test") // akan diprint ketika semua test selesai dijalankan
}

func TestCobaMain(t *testing.T) {
	t.Log("TestCobaMain 100 x 10 = 1000")
	assert.Equal(t, 1000, New(num).Multiply(10).Result())
}
