/*:
 
# Linked list
 
 
 ### Categories:
 - Single linked list
 - Duplex linked list
 - Cycle linked list
 
 ### Use case:
 - LRU(Least Recently Used)
 
*/

import UIKit

// Single linked list

class Node<T> {
    var data: T?
    var next: Node?
    
    init() {}
    
    init(_ value: T) {
        self.data = value
    }
}

/// For initial state, what does it look like at the very beginning?
class SingleLinkedList<E: Equatable> {
    private var pointer = Node<E>()     // always pointing to the head, an outstander, holds no data

    func insertToHead(_ element: E) {
        let node = Node(element)
        node.next = pointer.next
        pointer.next = node
    }
    
    func delete(_ element: E) {
        guard pointer.next != nil else {
            print("empty linked list")
            return
        }
        var tmp = pointer
        var cursor = pointer.next
        while cursor != nil {
            if cursor?.data == element {
                tmp.next = cursor?.next
                break
            } else {
                tmp = cursor!
                cursor = cursor!.next
            }
        }
    }
    
    func deleteHead() {
        if pointer.next != nil {
            let head = pointer.next!
            pointer.next = head.next
        }
    }
    
    func find(_ element: E) -> Node<E>? {
        guard pointer.next != nil else {
            print("empty linked list")
            return nil
        }
        var cursor = pointer.next
        while cursor != nil {
            if cursor?.data! == element {
                print("Element ", element, " found")
                return cursor
            } else {
                cursor = cursor?.next
            }
        }
        print("Element ", element, " not found")
        return nil
    }

    func isEmpty() -> Bool {
        return pointer.next == nil
    }
    
    func description() {
        if pointer.next != nil {    // indicates that linked list is not empty
            var cursor = pointer.next
            var desc = ""
            while cursor != nil && cursor!.data != nil {
                desc.append("\(cursor!.data!)" + " -> ")
                cursor = cursor?.next
            }
            print(desc)
        }
    }
}

/// Test case:

let head = Node<Int>()
head.data = 9

let linkedList = SingleLinkedList<Int>()
linkedList.insertToHead(8)
linkedList.insertToHead(7)
linkedList.insertToHead(6)
linkedList.description()

//linkedList.deleteHead()
//linkedList.description()

//linkedList.find(8)
//linkedList.find(9)

linkedList.delete(7)
linkedList.description()
