package main

import "fmt"

type Product struct {
	name    string
	price   int
	company string
}

func main() {
	//var companyName = "Apple"
	//fmt.Println("Company Name ", companyName)
	//x, y := check_odd_even2(11)
	//fmt.Println(x, y)

	//	p := Product{
	//		name:    "IPhone 14 Plus",
	//		price:   62000,
	//		company: "Apple Inc",
	//}
	//fmt.Println("Product Name : ", p.name)
	//fun(p)
	//fmt.Println("\nProduct Name : ", p.name)

	//demo_pointers()

	new_p := newProduct("IPhone 14 Plus", 62000, "Apple Inc")
	fmt.Println("Product Name : ", new_p.name)
	fmt.Println("Product price : ", new_p.price)
	fmt.Println("Company Name : ", new_p.company)
}

func newProduct(name string, price int, company string) *Product {
	p := Product{
		name:    name,
		price:   price,
		company: company,
	}
	return &p
}
func fun(copyOfP Product) {
	copyOfP.name = "Iphone18"
	print("Product new ", copyOfP.name)
}
func loops_demo() {

	for i := 0; i < 5; i++ {
		fmt.Println("Loop Iteration ", i)
	}

	for i := range 3 {
		fmt.Println("Range Iteration ", i)
	}

	for i, char := range "Sanket" {
		fmt.Println("String range loop", i, char)
	}

	j := 10
	for j > 0 {
		fmt.Println("While Loop Iteration ", j)
	}
}

func maps_demo() {
	productPrices := map[string]int{
		"Iphone14": 5000,
		"Iphone15": 10000,
		"Iphone16": 15000,
	}

	fmt.Println("Product Prices ", productPrices)

	customMap := map[string]string{}

	fmt.Println("Custom Map", customMap)

	emptyMap := make(map[string]int)

	fmt.Println("Empty Map ", emptyMap)

	emptyMap["key1"] = 100
	emptyMap["key2"] = 200
	emptyMap["key3"] = 300

	fmt.Println("Populated Empty Map ", emptyMap)

	for key, value := range emptyMap {
		fmt.Println("Key ", key, " Value ", value)
	}

}

func check_odd_even(num int) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func check_odd_even2(num int) (string, int) {
	if num%2 == 0 {
		return "Even", 0
	} else {
		return "Odd", 1
	}
}

func demo_pointers() {
	i := 120
	var ptr *int = &i
	ptr1 := &i
	fmt.Println(i, ptr, ptr1)

	fmt.Println("Value present at pointer ", *ptr)
	fmt.Println("Value present at pointer ", *ptr1)
}
