package main

func main() {

}

// TLE
//type Item struct {
//    Val int
//    PopVal int
//}
//
//type Stack struct {
//    data []*Item
//}
//
//func (s *Stack) Push(item *Item) {
//    s.data = append(s.data, item)
//}
//
//func (s *Stack) Len() int {
//    return len(s.data)
//}
//
//func (s *Stack) Pop() *Item {
//    if s.Len() == 0 {
//        return nil
//    }
//    item := s.data[len(s.data)-1]
//    s.data = s.data[:len(s.data)-1]
//    return item
//}
//
//func (s *Stack) Top() *Item {
//    if s.Len() == 0 {
//        return nil
//    }
//    return s.data[len(s.data)-1]
//}
//
//type StockSpanner struct {
//    sigStack *Stack
//    cacheStack *Stack
//}
//
//
//func Constructor() StockSpanner {
//    return StockSpanner{sigStack: &Stack{}, cacheStack: &Stack{}}
//}
//
//
//func (this *StockSpanner) Next(price int) int {
//    top := this.sigStack.Top()
//    item := &Item{Val: price}
//    result := 1
//    if top == nil || top.Val > price {
//        this.sigStack.Push(item)
//        return result
//    }
//    for this.sigStack.Len() > 0 && this.sigStack.Top().Val <= price {
//        p := this.sigStack.Pop()
//        if p.PopVal <= price {
//            result++
//            p.PopVal = price
//        }
//        this.cacheStack.Push(p)
//    }
//    this.sigStack.Push(item)
//    for this.cacheStack.Len() > 0 {
//        this.sigStack.Push(this.cacheStack.Pop())
//    }
//    return result
//}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
