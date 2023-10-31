class Node:
    def __init__(self, data):
        self.data = data
        self.next_node = None

class LinkedList:
    def __init__(self):
        self.head = None
        self.size = 0

    def add_node(self, data):
        new_node = Node(data)
        if self.head is None:
            self.head = new_node
            self.size += 1
        else:
            current_node = self.head
            while current_node.next_node is not None:
                current_node = current_node.next_node
            current_node.next_node = new_node
            self.size += 1

    def print_list(self):
        current_node = self.head
        while current_node is not None:
            print(current_node.data)
            current_node = current_node.next_node

    def get_size(self):
        return self.size
    
    def remove_middle_node(self):
        num_node = self.get_size()//2
        current_node = self.head
        i = 0
        while current_node is not None and i != num_node:
            i += 1
            prev_node = current_node
            current_node = current_node.next_node
        prev_node.next_node = current_node.next_node

def sum_list(LinkedList: List1, LinkedList: List2) -> int:
    current_node1 = List1.head
    current_node2 = List2.head
    string_sum1 = ""
    string_sum2 = ""
    while current_node1 is not None:
        string_sum1 += str(current_node1.data)
        current_node1 = current_node1.next_node
    while current_node2 is not None:
        string_sum2 += str(current_node2.data)
        current_node2 = current_node2.next_node
    return int(string_sum1) + int(string_sum2)

if __name__ == "__main__":
    l = LinkedList()
    l.add_node(3)
    l.add_node(0)
    l.add_node(0)
    l.add_node(0)

    q = LinkedList()
    q.add_node(1)
    q.add_node(0)
    q.add_node(0)
    q.add_node(0)
    print(sum_list(l,q))
    