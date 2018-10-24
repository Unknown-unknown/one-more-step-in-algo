//
// Created by Calios on 2018/10/14.
// Copyright (c) 2018 Calios. All rights reserved.
//


import Foundation

struct WrapperArray<Item> where Item: Equatable {
    private var data: [Item]
    private var capacity = 0
    private var count = 0
    
    // constructor
    init(defaultItem: Item, capacity: Int) {
        data = Array(repeating: defaultItem, count: capacity)
        self.capacity = capacity
    }
    
    init(data: [Item], capacity: Int) {
        self.data = data
        self.capacity = capacity
        count = data.count
    }
    
    // subscript
    subscript(index: Int) -> Item {
        get {
            return data[index]
        }
        set {
            data[index] = newValue
        }
    }
    
    // find by index
    func find(at index: Int) -> Item? {
        guard index >= 0 && index < count && count <= capacity else { return nil }
        
        return data[index]
    }
    
    // find by value
    func findIndex(by value: Item) -> Int? {
        return data.firstIndex { (itm) -> Bool in
            return itm == value
        }
    }
    
    // insert at index (move all)
    mutating func insert(at index: Int, item: Item) -> Bool {
        guard index >= 0 && index < count && count < capacity else { return false }
        
        if let last = data.last {
            data.append(last)
        }
        
        for i in (index ..< count - 1).reversed() {
            data[i + 1] = data[i]
        }
        
        data[index] = item
        count += 1
        return true
    }
    
    // insert at index (simple swap)
    mutating func insertSwap(at index: Int, item: Item) -> Bool {
        guard index >= 0 && index < count && count < capacity else { return false }
        
        let origin = data[index]
        data[index] = item
        data.append(origin)
        count += 1
        return true
    }
    
    // append
    mutating func append(_ item: Item) -> Bool {
        guard count < capacity else { return false }
        
        data.append(item)
        count += 1
        return true
    }
    
    // update at index
    mutating func update(at index: Int, item: Item) -> Bool {
        guard index >= 0 && index < count && count <= capacity else { return false }
        
        data[index] = item
        return true
    }
    
    // delete at index
    mutating func delete(at index: Int) -> Bool {
        guard index >= 0 && index < count && count <= capacity else { return false }
        
        for i in (index ..< count - 1) {
            data[i] = data[i + 1]
        }
        data = Array(data.dropLast())
        count -= 1
        return true
    }
    
    // check emptiness
    func isEmpty() -> Bool {
        return data.isEmpty
    }
}


/// Test cases:

var wrapper = WrapperArray(data: [1, 2, 3, 4, 5], capacity: 10)
var wrapper1 = wrapper
var wrapper2 = wrapper
var wrapper3 = wrapper
print(wrapper)
wrapper.find(at: 7)
wrapper.find(at: 4)
wrapper.findIndex(by: 10)
wrapper.findIndex(by: 5)
wrapper.insert(at: 2, item: 6)
print(wrapper)
wrapper1.insertSwap(at: 2, item: 6)
print(wrapper1)
wrapper2.delete(at: 2)
print(wrapper2)
wrapper3.update(at: 1, item: 9)
print(wrapper3)
