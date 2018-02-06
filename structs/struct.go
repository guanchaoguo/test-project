package main

import(
	"encoding/json"
	"fmt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	user:= &User{Email:"123",Password:"456"}
	res,err:=json.Marshal(struct {
		*User
		Password bool `json:"password,omitempty"`
	}{
		User: user,
	})
	 fmt.Println(string(res),err)

}

