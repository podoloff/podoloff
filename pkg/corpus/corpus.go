package corpus

// Corpus contains the set of restricted content
type Corpus struct {
	Size  int
	words map[string]bool
}

// NewCorpus creates a new corpus
func NewCorpus(basis []string) *Corpus {
	w := make(map[string]bool)
	for _, v := range basis {
		if _, ok := w[v]; !ok {
			w[v] = true
		}
	}
	return &Corpus{
		Size:  len(w),
		words: w,
	}
}

// Add appends an arbitrary number of words to the corpus
func (c *Corpus) Add(vals ...string) bool {
	for _, v := range vals {
		if _, ok := c.words[v]; !ok {
			c.Size++
			c.words[v] = true
		}
	}
	return true
}

// Remove deletes an arbitrary number of words from the corpus
func (c *Corpus) Remove(vals ...string) bool {
	for _, v := range vals {
		if _, ok := c.words[v]; ok {
			c.Size--
			c.words[v] = false
		}
	}
	return true
}

// Has returns true if the word already exists in the corpus and false if not
func (c *Corpus) Has(val string) bool {
	return c.words[val]
}
