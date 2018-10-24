//
// Created by Calios on 2018/10/14.
// Copyright (c) 2018 Calios. All rights reserved.
//

import Foundation

struct WrapperArray<Item> {
    private var data: [Item]
    private var capacity = 0
    
    // constructor
    init(defaultItem: Item, capacity: Int) {
        data = Array(repeating: defaultItem, count: capacity)
        self.capacity = capacity
    }
    
    // subscript
    subscript(index: Int) -> Item { get set }
    
    // find by index
    func find(at index: Int) -> Item? {
        guard index >= 0, index < count else { return nil }
        return data[index]
    }
    
    // find by value
    func find(by value: Item) -> Int? {
        return data.index(value)
    }
}
