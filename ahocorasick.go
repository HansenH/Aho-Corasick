package ahocorasick

import (
	"container/list"
)

type trieNode struct {
	count int
	fail  *trieNode
	child map[byte]*trieNode
	index int
}

func newTrieNode() *trieNode {
	return &trieNode{
		count: 0,
		fail:  nil,
		child: make(map[byte]*trieNode),
		index: -1,
	}
}

type ACAutomaton struct {
	root       *trieNode
	size       int
	dictionary []string
}

// Construct the Model with a certain dictionary
func NewACAutomaton(dictionary []string) *ACAutomaton {
	m := &ACAutomaton{
		root:       newTrieNode(),
		size:       0,
		dictionary: dictionary,
	}
	m.build()
	return m
}

// Initialize the Aho-Corasick Automaton
func (m *ACAutomaton) build() {
	for i := range m.dictionary {
		m.insert(m.dictionary[i])
	}
	ll := list.New()
	ll.PushBack(m.root)
	for ll.Len() > 0 {
		temp := ll.Remove(ll.Front()).(*trieNode)
		var p *trieNode = nil
		for i, v := range temp.child {
			if temp == m.root {
				v.fail = m.root
			} else {
				p = temp.fail
				for p != nil {
					if p.child[i] != nil {
						v.fail = p.child[i]
						break
					}
					p = p.fail
				}
				if p == nil {
					v.fail = m.root
				}
			}
			ll.PushBack(v)
		}
	}
}

func (m *ACAutomaton) insert(s string) {
	curNode := m.root
	for _, v := range []byte(s) {
		if curNode.child[v] == nil {
			curNode.child[v] = newTrieNode()
		}
		curNode = curNode.child[v]
	}
	curNode.count++
	curNode.index = m.size
	m.size++
}

// Search all the matched positions of all patterns
func (m *ACAutomaton) FindAllIndex(s string) (res map[string][]int) {
	curNode := m.root
	var p *trieNode = nil
	res = make(map[string][]int)

	for index, b := range []byte(s) {
		for curNode.child[b] == nil && curNode != m.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[b]
		if curNode == nil {
			curNode = m.root
		}
		p = curNode
		for p != m.root && p.count > 0 { //&& !mark[p.index]
			// mark[p.index] = true
			for i := 0; i < p.count; i++ {
				term := m.dictionary[p.index]
				startPos := index - len(term) + 1
				res[term] = append(res[term], startPos)
			}
			p = p.fail
		}
	}
	return res
}
