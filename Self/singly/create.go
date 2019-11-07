package singly

func HeadCreate(data []string) (chain *SinglyChain) {
	if data == nil {
		return
	}
	for _, name := range data {
		new := SinglyChain{Name: name}
		if chain == nil {
			chain = &new
		} else {
			tmp := chain
			new.Next = tmp
			chain = &new
		}
	}
	return
}

func HeadCreateExample1() {
	line := HeadCreate([]string{"a", "b", "c", "d"})
	line.Print()
	line = HeadCreate([]string{"a"})
	line.Print()
}

func TailCreate(data []string) (chain *SinglyChain) {
	if data == nil {
		return
	}
	var prev *SinglyChain
	for _, name := range data {
		new := SinglyChain{Name: name}
		if chain == nil {
			chain = &new
			prev = &new
			continue
		}
		prev.Next = &new
		prev = &new
	}
	return
}

func TailCreateExample1() {
	line := TailCreate([]string{"a", "b", "c", "d"})
	line.Print()
	line = TailCreate([]string{"a"})
	line.Print()
}
