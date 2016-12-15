package main

import "fmt"

// Recur

func AkkR(m,n int)  int{
  res := 0;
  switch {
    case m==0:
      res = n+1
    case m>0 && n==0:
      res = AkkR(m-1,1)
    case m>0 && n>0:
      res = AkkR(m-1,AkkR(m,n-1))
  }
  return res
}

var memo = make(map[string]int64)

func add(m,n,res int64) map[string]int64 {
  if find(m,n) != int64(-1) {
    return memo
  } else {
    id := fmt.Sprintf("%d",m) + "_" + fmt.Sprintf("%d",n)
    memo[id] = res
  }
  return memo
}

func find(m,n int64) int64 {
  id1 := fmt.Sprintf("%d",m)
  id2 := fmt.Sprintf("%d",n)
  id := id1+"_"+id2
  res := memo[id]
  if res == 0 {
    res = int64(-1)
  }
  return res
}

// Recur with memo?

func Akk(m,n int64) int64 {
  stack := [][]int64{}
  res := int64(0);
  //result := int64(0);
  switch {
    case m==int64(0):
      res = res + n+int64(1)
      fmt.Printf("case2: %v\n",stack)
      memo = add(m,n,res)
    case m>0 && n==0:
      res = find(m,n)
      if res == int64(-1) {
        stack = append(stack, []int64{m,1})
        res = Akk(m-int64(1),int64(1))
        fmt.Printf("case2: %v\n",stack)
        memo = add(m,n,res)
      }
    case m>0 && n>0:
      res = find(m,n)
      if res == int64(-1) {
        stack = append(stack, []int64{m,n})
        mid := Akk(m,n - int64(1))
        stack = append(stack, []int64{m,mid})
        res = Akk(m-int64(1),mid)
        memo = add(m,n,res)
      }
  }
  return res
}

//Trampoline version

type akkfun func(m, n, acc int64) (akkfun, int64, int64, int64)

func akk (m, n, acc int64) (akkfun, int64, int64, int64) {
  switch {
    case m==0:
      return nil, 0, 0, acc+n+1
    case m>0 && n==0:
      return akk, m-1, 1, acc
    default:
      internalRes := AkkTramp(m,n-int64(1))
      return akk, m-1, internalRes, acc
  }
}

func AkkTramp(m, n int64) int64 {
  fun := akkfun(akk)
  acc := int64(0)

  for fun !=nil {
    fun, m, n, acc = fun(m, n, acc)
  }
  return acc
}

//Iter version

// Iter version

type Stack struct{
  len int64
  stack [][]int64
}

func newStack() *Stack {
  return &Stack{0, [][]int64{}}
}

func (s *Stack) popS() (bool, []int64){
  if s.len!=0 {
      item := s.stack[s.len-1]
      s.stack = s.stack[:s.len-1]
      s.len = s.len-1
      //fmt.Printf("POP ELEM FROM STACK: %v STACK NOW IS: %v\n",item,stack)
      return true, item
  }
  return false, nil
}

func (s *Stack) pushS(m,n int64) {
  s.stack = append(s.stack, []int64{m,n})
  s.len = s.len+1
  //fmt.Printf("PUSH ELEM TO STACK: %v STACK NOW IS: %v\n",[]int64{m,n},stack)
  return
}

func pop() (bool,[]int64) {
  if len(stack)!=0 {
      item := stack[len(stack)-1]
      stack = stack[:len(stack)-1]
      //fmt.Printf("POP ELEM FROM STACK: %v STACK NOW IS: %v\n",item,stack)
      return true, item
  }
  //fmt.Println("STACK IS EMPTY")
  return false, []int64{}
}

func push(m,n int64) {
  stack = append(stack, []int64{m,n})
  //fmt.Printf("PUSH ELEM TO STACK: %v STACK NOW IS: %v\n",[]int64{m,n},stack)
  return
}


func AkkStackEmul(m,n int64) int64 {
  res := int64(0)
  result := int64(0)
  //stack := newStack()
  internalRes := int64(0)
  var notEmpty bool
  var p []int64

  for {
    //fmt.Printf("M: %v. N: %v. res: %v\n",m,n,result)
    //fmt.Printf("MEMO:: %v\n",memo)
    switch {
      case m == 0:
        for {
          //notEmpty,p = stack.popS()
          notEmpty,p = pop()
          if notEmpty {
            if p[1]==0 {
              continue
            } else {
                res = n+1
                result = int64(0)
                //fmt.Println("Res: ",res)
                add(p[0],p[1]-1,res)
                m, n = p[0]-1, res
                //fmt.Println("Next call:",m,n)
                break
            }
          } else {
            return res + 1
          }
        }
      case m>0 && n==0:
        internalRes = find(m,n)
        if internalRes == -1 {
          //stack.pushS(m,n)
          push(m,n)
          m, n = m-1, 1
        } else{
          result = result+internalRes
          //notEmpty,p = stack.popS()
          notEmpty,p = pop()
          if notEmpty {
            m, n = p[0]-1, internalRes
          } else {
            return result
          }
        }
      case m>0 && n>0:
        internalRes = find(m,n)
        if internalRes == -1 {
          //stack.pushS(m,n)
          push(m,n)
          n = n-1
        } else {
          result = result+internalRes
          //notEmpty,p = stack.popS()
          notEmpty,p = pop()
          if notEmpty {
            m, n = p[0]-1, internalRes
          } else {
            return result
          }
        }
    }
  }
  if result==0{
    return res
  }
  return result
}

// CHEAT version

func hyper3(n *big.Int) *big.Int{
  return new(big.Int).Exp(big.NewInt(2), n, nil)
}

func hyper4(n *big.Int) *big.Int{
  res := new(big.Int).SetInt64(2)
  one := new(big.Int).SetInt64(1)
  for ;n.Cmp(one)==1;n.Sub(n, one) {
    res.Set(hyper3(res))
  }
  return res
}

func hyper5(n *big.Int) *big.Int{
  res := new(big.Int).SetInt64(2)
  one := new(big.Int).SetInt64(1)
  for ;n.Cmp(one)==1;n.Sub(n, one) {
    res.Set(hyper4(res))
  }
  return res
}

func AkkCheat(m, n int64) *big.Int {
  res := new(big.Int).SetInt64(0)
  var three *big.Int
  if m>2 {
    n = n+3
    three = new(big.Int).SetInt64(3)
  }
  switch {
    case m==int64(0):
      res.SetInt64(n+1)
    case m==int64(1):
      res.SetInt64(n+2)
    case m==2:
      res.SetInt64(int64(2)*n+int64(3))
    case m==3:
      res.Sub(hyper3(new(big.Int).SetInt64(n)), three)
    case m==4:
      res.Sub(hyper4(new(big.Int).SetInt64(n)), three)
    case m==5:
      res.Sub(hyper5(new(big.Int).SetInt64(n)), three)
  }
  return res
}

//===================

main () {
  fmt.Println(Akk(2,2))
}

  // for {
  //   switch {
  //     case m==int64(0):
  //       res = res + n+int64(1)
  //       memo = add(m,n,res)
  //     case m>0 && n==0:
  //       res = find(m,n)
  //       if res == int64(-1) {
  //         res = Akki(m-int64(1),int64(1))
  //         memo = add(m,n,res)
  //       }
  //     case m>0 && n>0:
  //       res = find(m,n)
  //       if res == int64(-1) {
  //         stack = append(stack, []int64{3,m,n})
  //         n = n - int64(1)
  //         stack = append(stack, []int64{3,m,n})

  //         res = Akki(m-int64(1),Akki(m,n - int64(1)))
  //         memo = add(m,n,res)
  //       }
  //       else {
  //         result = result + res
  //       }
  //   }
  // }