package types

type File struct {
	Path    string
	IsDir   bool
	ModTime int64
	Mode    string
	Name    string
	Size    int64
	Content string
}
