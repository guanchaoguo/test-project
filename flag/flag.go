package main

import "flag"
import "fmt"
import "strconv"

type percentage float32

func (p *percentage) Set(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	*p = percentage(v)
	return err
}
func (p *percentage) String() string { return fmt.Sprintf("%f", *p) }


func main() {
	namePtr := flag.String("name", "lyh", "user's name")
	agePtr := flag.Int("age", 22, "user's age")
	vipPtr := flag.Bool("vip", true, "is a vip user")


	var email string
	flag.StringVar(&email, "email", "lyhopq@gmail.com", "user's email")

	var pop percentage
	flag.Var(&pop, "pop", "popularity")
	flag.Parse()
	others := flag.Args()
	fmt.Println("name:", *namePtr)
	fmt.Println("age:", *agePtr)
	fmt.Println("vip:", *vipPtr)
	fmt.Println("pop:", pop)
	fmt.Println("email:", email)
	fmt.Println("other:", others)
}
