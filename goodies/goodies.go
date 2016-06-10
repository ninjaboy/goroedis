package goodies

// Goodies bag
type Goodies struct {
	Storage map[string]interface{}
}

// CreateGoodies creates new isntance of goodiebag
func CreateGoodies() *Goodies {
	return &Goodies{Storage: make(map[string]interface{})}
}

// Set Method
func (g *Goodies) Set(key string, value interface{}) {
	g.Storage[key] = value
}

// Get Method
func (g *Goodies) Get(key string) (interface{}, bool) {
	val, found := g.Storage[key]
	return val, found
}

// Update method (at the moment not clear how it should be different to Set)
func (g *Goodies) Update(key string, value string) {
	g.Set(key, value)
}

// Remove key from storage
func (g *Goodies) Remove(key string) {
	delete(g.Storage, key)
}

// Keys returns list of keys
func (g *Goodies) Keys() []string {
	keys := make([]string, len(g.Storage))
	i := 0
	for k := range g.Storage {
		keys[i] = k
		i++
	}
	return keys
}
