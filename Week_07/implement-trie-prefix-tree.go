type Trie struct {
	Data map[byte]Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{Data: make(map[byte]Trie)}
	t.Data['#'] = Trie{Data: make(map[byte]Trie)}
	return t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.Data
	for i := range word {
		if v, ok := node[word[i]]; !ok {
			node[word[i]] = Trie{Data: make(map[byte]Trie)}
			node = node[word[i]].Data
		} else {
			node = v.Data
		}
	}
	node['#'] = Trie{Data: make(map[byte]Trie)}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.Data
	for i := range word {
		if v, ok := node[word[i]]; !ok {
			fmt.Println("vivid")
			return false
		} else {
			node = v.Data
		}
	}
	if _, ok := node['#']; ok {
		return true
	}

	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.Data
	for i := range prefix {
		if v, ok := node[prefix[i]]; !ok {
			return false
		} else {
			node = v.Data
		}
	}
	return true
}