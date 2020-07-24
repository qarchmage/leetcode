package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	obj := Constructor()
	obj.Inserting(10)
	obj.GetRandoms(10)
	obj.Removing(10)
}

type RandomizedSet struct {
	Values []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{Values: make([]int, 0)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	for _, v := range this.Values {
		if v == val {
			return false
		}
	}
	this.Values = append(this.Values, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	for i, v := range this.Values {
		if v == val {
			copy(this.Values[i:], this.Values[i+1:])
			this.Values[len(this.Values)-1] = 0
			this.Values = this.Values[:len(this.Values)-1]
			return true
		}
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	if len(this.Values) == 1 {
		return this.Values[0]
	}
	return this.Values[rand.Int()%len(this.Values)]
}

// TESTING AREA

func (this *RandomizedSet) Inserting(to int) {
	for i := 0; i < to; i++ {
		this.Insert(i)
	}
	fmt.Println(this.Values)
}

func (this *RandomizedSet) Removing(from int) {
	for i := from; i > -1; i-- {
		this.Remove(i)
	}
	fmt.Println(this.Values)
}

func (this *RandomizedSet) GetRandoms(to int) {
	for i := 0; i < to; i++ {
		fmt.Println(this.GetRandom())
	}
}
