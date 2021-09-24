package main
import "fmt"

type node struct {
    data int
    next *node
}
type linkedlist struct {
    head *node;
    length int
}

func (ll *linkedlist) insert(val int) {
    el:=&node{data:val};
    if(ll.length == 0) {
        ll.head = el
        ll.length++
    }else {
        last:=ll.head
        ll.length++
        for last.next!=nil {
            last = last.next
        }
        last.next = el
    }
}

func (ll *linkedlist) prepend(val int) {
    el:=&node{data:val};
    ll.length++
    el.next = ll.head;
    ll.head = el
}

func (ll *linkedlist) printList() {
    last:=ll.head
    for last!=nil {
        fmt.Print(last.data," ")
        last = last.next
    }
}
func (ll *linkedlist) reverse() {
    current:=ll.head
    var prev *node
    // next:=&node{}
    for current!=nil {
        next := current.next;
        current.next = prev
        prev = current
        current = next
    }
    ll.head = prev
}

func main(){
    root:=linkedlist{};
    
    root.insert(10);
    root.insert(20);
    root.insert(30);
    root.insert(40);
    root.insert(50);
    root.prepend(8);
    root.prepend(6);
    root.prepend(4);
    fmt.Println(root.length)
    // root.reverse()
    root.printList()
}
