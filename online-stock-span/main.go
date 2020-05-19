package main

import "fmt"

func main() {
	o := Constructor()
	o.Next(100)
	o.Next(80)
	o.Next(60)
	o.Next(70)
	o.Next(60)
	o.Next(75)
	o.Next(85)
	// control
	fmt.Println(o.price)
	fmt.Println(o.span)
}

type StockSpanner struct {
	price []int
	span  []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	c := 0
	this.price = append(this.price, price)
	for i := len(this.price) - 1; i >= 0; i-- {
		if this.price[i] <= price {
			c++
		} else {
			break
		}
	}
	this.span = append(this.span, c)
	return c
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
