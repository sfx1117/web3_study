package main

import "fmt"

/*
	interface接口
1. 接口可以嵌套。
2. 接口中声明的方法，参数可以没有名称。
3. 如果函数参数使用 interface{}可以接受任何类型的实参。同样，可以接收任何类型的值也可以赋值给 interface{}类型的变量
*/
//声明两个接口
type Account interface {
	//获取余额
	GetBalance() int
}
type PayMethod interface {
	Account
	//付款
	Pay(int) bool
}

//声明两个结构体，并实现接口
//信用卡
type Cred struct {
	balance int //账单
	limit   int //额度
}

func (c *Cred) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡支付成功：%d\n", amount)
		return true
	}
	fmt.Println("信用卡支付失败：超出额度")
	return false
}
func (c *Cred) GetBalance() int {
	return c.balance
}

//借记卡
type Debit struct {
	balance int
}

func (d *Debit) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Printf("借记卡支付成功：%d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败：余额不足")
	return false
}
func (d *Debit) GetBalance() int {
	return d.balance
}

//定义购买商品函数
func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功，剩余余额=%d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

func main() {
	c := &Cred{0, 1000}
	//d := &Debit{500}
	fmt.Println(c.GetBalance())
	//先使用信用卡购买
	//purchaseItem(c, 800)
	//purchaseItem(c, 300)
	//purchaseItem(d, 400)
	//purchaseItem(d, 200)
}
