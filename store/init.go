package store

// Init - initialise a store
func Init(path string) Store {
	s := Store{
		path: path,
	}
	return s
}
