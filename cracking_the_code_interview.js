const linkedList = {



}

class LinkedList{
    constructor(next,data){
    this.next = next;
    this.data = data;
    }

    append(data){
        var child = new LinkedList(null,data);
        var current = this;
        while (current !== null){
            current = current.next
        }
        current.next = child;
    }
}


// head.next = new Node(null,"b")

function palidrome(head){
    let current = head;
    let stack = []
    while (current !== null){
        stack.push(current.data)
        current = curent.next
    }
    current = head
    for (i = 0; i < stack.length;i++){
        if(current.data !== stack.pop()){
            return false;
        }
        current = current.next;
    }
    return true;
}
let notPalindrome = new LinkedList(1,2,3)
