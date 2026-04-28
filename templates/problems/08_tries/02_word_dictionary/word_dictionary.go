package word_dictionary

// WordDictionary supports adding words and searching with '.' wildcards.
// Design the data structure yourself.
type WordDictionary struct {
}

// NewWordDictionary creates and returns a new WordDictionary.
func NewWordDictionary() WordDictionary {
	return WordDictionary{}
}

// AddWord adds a word to the dictionary.
func (wd *WordDictionary) AddWord(word string) {
}

// Search returns true if the word matches any string in the dictionary.
// '.' can match any single character.
func (wd *WordDictionary) Search(word string) bool {
	return false
}
