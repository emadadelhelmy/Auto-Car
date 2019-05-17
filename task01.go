package main

import(
	"bufio"
	"fmt"
	"os"
)
func rotate(text string, key int)string {
	key = (key%26 + 26) % 26
	b := make([]byte , len(text))
	for i := 0; i < len(text); i++ {
		t:= text[i]
		var a int
		switch  {
		case 'a' <= t && t <= 'z':
			a = 'a'
		case 'A' <= t && t <='z' :
			a = 'A'
		default:
			b[i]= t
			continue

		}
		b[i] = byte(a+((int(t)-a) +key)%26)

	}
	return string(b)

}

func decrypt (cipher string, key int )(plain string)  {
	return rotate(cipher,-key)

}

func main(){
	fmt.Println(  )

	fmt.Println("Decrypting:")
	fmt.Print("Enter text to decrypt: ")
	reader := bufio.NewReader(os.Stdin)
	value1, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Print("Enter encryption key: ")
	var dec_key int
	_,error1 := fmt.Scanf("%d",&dec_key)

	if(error1 != nil) {
		fmt.Println(error1)
	}

	value2 := decrypt (value1, dec_key)
	fmt.Println("Decrypted text: ",value2)
	fmt.Println()
	reader.ReadString('\n')
	fmt.Println()


}


/*package main

import(
	"bufio"
	"fmt"
	"os"
)
func rotate(text string, key int)string {
	key = (key%26 + 26) % 26
	b := make([]byte , len(text))
	for i := 0; i < len(text); i++ {
		t:= text[i]
		var a int
		switch  {
		case 'a' <= t && t <= 'z':
			a = 'a'
		case 'A' <= t && t <='z' :
			a = 'A'
		default:
			b[i]= t
			continue

		}
		b[i] = byte(a+((int(t)-a) +key)%26)

	}
	return string(b)

}

func decrypt (cipher string, key int )(plain string)  {
	return rotate(cipher,-key)

}

func main(){
	fmt.Println(  )

	fmt.Println("Decrypting:")
	fmt.Print("Enter text to decrypt: ")
	reader := bufio.NewReader(os.Stdin)
	value1, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Print("Enter encryption key: ")
	var dec_key int
	_,error1 := fmt.Scanf("%d",&dec_key)

	if(error1 != nil) {
		fmt.Println(error1)
	}

	value2 := decrypt (value1, dec_key)
	fmt.Println("Decrypted text: ",value2)
	fmt.Println()
	reader.ReadString('\n')
	fmt.Println()


*/