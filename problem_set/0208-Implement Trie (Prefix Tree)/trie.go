type Trie struct {
    children [26]*Trie
    isEnd    bool
}

// Constructor initializes an empty Trie.
func Constructor() Trie {
    return Trie{}
}

// Insert adds a word into the Trie.
func (this *Trie) Insert(word string) {
    for i := range word {
        c := word[i] - 'a'
        if child := this.children[c]; child == nil {
            this.children[c] = &Trie{}
        }
        this = this.children[c]
    }
    this.isEnd = true
}

// Search returns if the word is in the Trie.
func (this *Trie) Search(word string) bool {
    for i := range word {
        c := word[i] - 'a'
        if child := this.children[c]; child == nil {
            return false
        }
        this = this.children[c]
    }
    return this.isEnd
}

// StartsWith returns if there is any word in the Trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool {
    for i := range prefix {
        c := prefix[i] - 'a'
        if child := this.children[c]; child == nil {
            return false
        }
        this = this.children[c]
    }
    return true
}