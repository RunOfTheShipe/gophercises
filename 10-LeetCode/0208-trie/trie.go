package trie

type Trie struct {
	Letter byte
	IsWord bool
	Next   []*Trie
}

func Constructor() Trie {
	t := Trie{}
	t.Next = nil
	t.IsWord = false
	return t
}

func (node *Trie) Insert(word string) {

	if len(word) > 0 {

		if node.Next == nil {
			// nothing in the list of nodes - create our list
			// and add the first node
			node.Next = []*Trie{create(word[0])}

			// move to the next index
			node.Next[0].Insert(word[1:])
		} else {
			// list contains something; see if we can find a
			// Trie that matches the current letter
			letter := word[0]

			for _, n := range node.Next {
				if n.Letter == letter {

					n.Insert(word[1:])

					// don't need to do anything else, so
					// exit early
					return
				}
			}

			// didn't find anything; create a new Trie at the
			// end of the list and continue adding the word
			node.Next = append(node.Next, create(letter))
			node.Next[len(node.Next)-1].Insert(word[1:])
		}
	} else {
		// no letters left in the word, so this node must
		// represent a word
		node.IsWord = true
	}
}

func create(letter byte) *Trie {
	t := Constructor()
	t.Letter = letter
	return &t
}

func (node *Trie) Search(word string) bool {
	return search(node, word, false)
}

func (node *Trie) StartsWith(prefix string) bool {
	return search(node, prefix, true)
}

func search(node *Trie, word string, startsWithOnly bool) bool {

	if len(word) == 0 {
		if startsWithOnly {
			// found something that matches all the letters
			// so it matches the prefix condition
			return true
		} else {
			// looking for a full word so as long as this node
			// represents a word, then the word has been found
			return node.IsWord
		}
	} else {
		for _, n := range node.Next {
			if n.Letter == word[0] {
				// found a Trie that matches the next letter and
				// can recurse to find the next letter
				return search(n, word[1:], startsWithOnly)
			}
		}

		// did not find any Trie that matches the
		// next letter, so it must not exist yet
		return false
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
