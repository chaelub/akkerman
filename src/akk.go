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

// Recur with memo

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

func popStack() (bool,[]int64) {
  if len(stack)!=0 {
      res := stack[len(stack)-1]
      stack = stack[:len(stack)-1]
      fmt.Printf("POP FROM STACK ELMENT: %v STACK NOW IS: %v\n",res,stack)
      return true, res
  }
  fmt.Println("STACK IS EMPTY")
  return false, []int64{}
}

func addToStack(m,n int64) {
  stack = append(stack, []int64{m,n})
  fmt.Printf("PUSH TO STACK ELEM: %v STACK NOW IS: %v\n",[]int64{m,n},stack)
  return
}

func AkkIter(m,n int64) int64 {
  res := int64(0);
  result := int64(0);
  var notEmpty bool
  var p []int64
  var nop bool = true
  for {
    fmt.Printf("STACK: %v. M: %v. N: %v. res: %v\n",stack,m,n,result)
    fmt.Printf("MEMO:: %v\n",memo)
    switch {
      case m == int64(0):
        res = n+int64(1)
        for ;nop; {
          notEmpty,p = popStack()
          if notEmpty==true{
            if p[1]==0{
              nop = false
            }
          }else {
            nop = false
          }
        }
        if notEmpty==true {
          result = result+res
          add(p[0],p[1],res)
        }
        nop = true
        notEmpty,p = popStack()
        if notEmpty==true {
          m, n = p[0]-int64(1), res
          if len(stack)==0 {
            continue
          }
        } else {
          return result
        }
      case m>0 && n==0:
        res = find(m,n)
        if res == int64(-1) {
          addToStack(m,n)
          m, n = m-int64(1), 1
        } else{
          result = result+res
          for ;nop; {
            notEmpty,p = popStack()
            if notEmpty==true{
              if p[1]==0{
                nop = false
              }
            }else {
              nop = false
            }
          }
          nop = true
          notEmpty,p = popStack()
          if notEmpty==true {
            addToStack(p[0]-int64(1),res)
            m, n = p[0]-int64(1), res
          } else {
            return result
          }
        }
      case m>0 && n>0:
        res = find(m,n)
        if res == int64(-1) {
          addToStack(m,n)
          n = n-1
        } else {
          result = result+res
          addToStack(m-int64(1),res)
          m, n = m-int64(1), res
        }
    }
    if len(stack)==0{
      return result
    }
  }
  return result
}

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