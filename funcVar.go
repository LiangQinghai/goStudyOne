package main

import "fmt"

type BankAccount struct {
	Name      string
	LeftMoney float64
}

type SaveMoney func(bankLeft float64, save float64) (left float64)

func Work(money SaveMoney, account *BankAccount, sal float64) (left float64) {

	res := money(account.LeftMoney, sal)
	fmt.Println("Work....", res, account.LeftMoney)

	account.LeftMoney = res

	return res

}
// 返回一个计算函数func(int, int) int
// flag: add/sub
func count(flag string) func(int, int) int {
	switch flag {
	case "add":
		return func(a int, b int) int{
			return a + b
		}
	case "sub":
		return func(a int, b int) int {
			return a - b
		}
	default:
		return nil
	}
}

func main() {

	var saveMoney SaveMoney
	saveMoney = func(bankLeft float64, save float64) (left float64) {
		res := bankLeft + save
		fmt.Println("SaveMoney...", res, bankLeft)
		return res
	}

	ba := BankAccount{Name: "Liang", LeftMoney: 888.00}

	fmt.Println(Work(saveMoney, &ba, 5000.00))

	add := count("add")
	fmt.Println(add(1, 2))

	sub := count("sub")(1, 2)
	fmt.Println(sub)

}
