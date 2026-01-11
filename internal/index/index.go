package index

type Index interface {
	Insert(value any, rowID int) error
	Search(value any) []int
}
