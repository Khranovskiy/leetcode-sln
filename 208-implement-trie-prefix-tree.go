package main

import "fmt"

func main() {
	var t Trie

	ops := []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}
	args := [][]string{{""}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}
	trueConst := true
	falseConst := false
	result := []*bool{nil, nil, &trueConst, &falseConst, &trueConst, nil, &trueConst}

	for i, _ := range ops {
		op := ops[i]
		arg := args[i][0]
		expected := result[i]

		var result *bool

		switch op {
		case "Trie":
			t = Constructor()
		case "insert":
			t.Insert(arg)
		case "search":
			res := t.Search(arg)
			result = &res
		case "startsWith":
			res := t.StartsWith(arg)
			result = &res
		}
		if result != nil && expected != nil {
			if *result != *expected {
				fmt.Println("different")
			} else {
				fmt.Println("same")
			}
		}
		if result == nil && expected != nil ||
			result != nil && expected == nil {
			fmt.Println("different")
		}
	}
}

type Trie struct {
	children   map[rune]*Trie
	isFullWord bool
}

func Constructor() Trie {
	return Trie{
		children: nil, //make(map[rune]*Trie),
	}
}

func (t *Trie) Insert(word string) {
	t.insert([]rune(word))
}

func (t *Trie) Search(word string) bool {
	node, exist := t.traverse([]rune(word))
	if !exist {
		return false
	}
	return node.isFullWord
}

func (t *Trie) StartsWith(prefix string) bool {
	_, exist := t.traverse([]rune(prefix))

	return exist
}

func (t *Trie) insert(word []rune) {
	if len(word) == 0 {
		t.isFullWord = true
		return
	}
	child, exist := t.children[word[0]]
	if !exist {
		child = &Trie{
			children: nil, //make(map[rune]*Trie),
		}
		if t.children == nil {
			t.children = make(map[rune]*Trie)
		}
		t.children[word[0]] = child
	}
	child.insert(word[1:])
}

func (t *Trie) traverse(prefix []rune) (*Trie, bool) {
	if len(prefix) == 0 {
		return t, true
	}
	child, exist := t.children[prefix[0]]
	if !exist {
		return nil, false
	}
	return child.traverse(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
