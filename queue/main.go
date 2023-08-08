package main

import "fmt"

type Queue struct{
    val[]int

}
func(q *Queue) enqueue(element int){
    q.val=append(q.val,element)
}
func(q *Queue) dequeue(){
    l:=len(q.val)
    q.val=q.val[1:l]
}
func main(){
    var noofele,ele,rem int
    Q:= Queue{}
    fmt.Print("no of ele to be added : ")
    fmt.Scan(&noofele)
    for i:=0;i<noofele;i++{
        fmt.Scan(&ele)
        Q.enqueue(ele)
    }
    fmt .Print("no of elements to be removed : ")
    fmt.Scan(&rem)
    for i:=0;i<rem;i++{
        Q.dequeue()
    }
    fmt.Print(Q)

}