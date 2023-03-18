package browserhistory

type BrowserHistory struct {
	Backwards stack
	Forwards  stack
	Current   string
}

func Constructor(homepage string) BrowserHistory {
	var bh = BrowserHistory{}
	bh.Backwards = stack{}
	bh.Forwards = stack{}
	bh.Current = homepage

	return bh
}

func (bh *BrowserHistory) Visit(url string) {
	bh.Backwards.Push(bh.Current)
	bh.Current = url

	// can't go forward any more - clear the forward
	// history
	bh.Forwards.Clear()
}

func (bh *BrowserHistory) Back(steps int) string {
	for i := steps; i > 0; i-- {
		if bh.Backwards.Count > 0 {
			bh.Forwards.Push(bh.Current)
			bh.Current = bh.Backwards.Pop()
		}
	}
	return bh.Current
}

func (bh *BrowserHistory) Forward(steps int) string {
	for i := steps; i > 0; i-- {
		if bh.Forwards.Count > 0 {
			bh.Backwards.Push(bh.Current)
			bh.Current = bh.Forwards.Pop()
		}
	}
	return bh.Current
}

// create a stack
type stack struct {
	Items []string
	Count int
}

func (s *stack) Push(str string) {
	if s.Items == nil {
		s.Items = make([]string, 1)
		s.Items[0] = str
		s.Count = 1
	} else {
		s.Items = append(s.Items, str)
		s.Count++
	}
}

func (s *stack) Pop() string {
	if s.Count > 0 {
		var item = s.Items[s.Count-1]

		// remove the last item from the array
		s.Items = s.Items[:s.Count-1]
		s.Count--

		return item
	}
	return ""
}

func (s *stack) Peek() string {
	if s.Count > 0 {
		return s.Items[s.Count-1]
	}
	return ""
}

func (s *stack) Clear() {
	s.Items = nil
	s.Count = 0
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
