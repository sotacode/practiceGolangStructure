package users

import (
	"fmt"
	"src/github.com/sotacode/practiceGolangStructure/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Nelson", time.Now(), true)
	fmt.Println(u)
}
