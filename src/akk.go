package main

import "fmt"

func Akk(m,n int)  int{
  res := 0;
  switch {
    case m==0:
      fmt.Println("M == 0. N ==",n)
      res = n+1
      fmt.Println("Middle result M=:",m,"N=",n,"res:", res)
    case m>0 && n==0:
      fmt.Println("M ==",m," N == 0")
      res = Akk(m-1,1)
      fmt.Println("Middle result M=:",m,"N=",n,"res:", res)
    case m>0 && n>0:
      fmt.Println("M ==",m, "N ==",n)
      res = Akk(m-1,Akk(m,n-1))
      fmt.Println("Middle result M=:",m,"N=",n,"res:", res)
  }
  fmt.Println("Middle result:", res)
  return res
}

var memo = make(map[string]int64)

func add(m,n,res int64) map[string]int64 {
  if find(m,n) != int64(-1) {
    return memo
  } else {
    id := fmt.Sprintf("%d",m) + "_" + fmt.Sprintf("%d",n)
    memo[id] = res
    //memo = append(memo, []int64{m,n,res})
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

func Akki(m,n int64) int64 {
  //stack := [][]int{}
  res := int64(0);
  //result := int64(0);
  switch {
    case m==int64(0):
      res = res + n+int64(1)
      memo = add(m,n,res)
    case m>0 && n==0:
      res = find(m,n)
      if res == int64(-1) {
        res = Akki(m-int64(1),int64(1))
        memo = add(m,n,res)
      }
    case m>0 && n>0:
      res = find(m,n)
      if res == int64(-1) {

        res = Akki(m-int64(1),Akki(m,n - int64(1)))
        memo = add(m,n,res)
      }
  }
  fmt.Println("res:", res)
  return res
}

  // switch {
  //   case m==int64(0):
  //     res = res + n+int64(1)
  //     memo = add(m,n,res)
  //   case m>0 && n==0:
  //     res = find(m,n)
  //     if res == int64(-1) {
  //       res = Akki(m-int64(1),int64(1))
  //       memo = add(m,n,res)
  //     }
  //   case m>0 && n>0:
  //     res = find(m,n)
  //     if res == int64(-1) {

  //       res = Akki(m-int64(1),Akki(m,n - int64(1)))
  //       memo = add(m,n,res)
  //     }
  // }

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