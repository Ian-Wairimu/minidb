package engine

type Tx struct {
	writes []func()
}

func (tx *Tx) Commit() {
	for _, w := range tx.writes {
		w()
	}
}
