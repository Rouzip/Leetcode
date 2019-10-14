import "sort"

func getSkyline(buildings [][]int) [][]int {
	buildingNum := len(buildings)
	var (
		events []Event
		id     int
	)

	for _, building := range buildings {
		events = append(events, Event{id, building[0], building[2], 1})
		events = append(events, Event{id, building[1], building[2], -1})
		id++
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].x == events[j].x {
			return events[i].eventType*events[i].h > events[j].eventType*events[j].h
		}
		return events[i].x < events[j].x
	})

	init := make([]int, buildingNum)
	for i := 0; i < buildingNum; i++ {
		init[i] = -1
	}
	heap := &MaxHeap{init, make([]pair, buildingNum), 0}
	var ans [][]int

	for _, event := range events {
		if event.eventType == 1 {
			if event.h > heap.Max() {
				ans = append(ans, []int{event.x, event.h})
			}
			heap.Add(event.h, event.id)
		} else {
			heap.Remove(event.id)
			if event.h > heap.Max() {
				ans = append(ans, []int{event.x, heap.Max()})
			}
		}
	}
	return ans
}

// record the event of building changes
type Event struct {
	id        int
	x         int
	h         int
	eventType int
}

type MaxHeap struct {
	idx  []int
	vals []pair
	size int
}

type pair struct {
	first, second int
}

func (mh *MaxHeap) Add(key, id int) {
	mh.idx[id] = mh.size
	mh.vals[mh.size] = pair{first: key, second: id}
	mh.size++
	mh.heapIfUp(mh.idx[id])
}

func (mh *MaxHeap) Remove(id int) {
	idToRemove := mh.idx[id]
	mh.swapNode(idToRemove, mh.size-1)
	mh.size--
	mh.heapIfDown(idToRemove)
	mh.heapIfDown(idToRemove)
}

func (mh *MaxHeap) Empty() bool {
	return mh.size == 0
}

func (mh *MaxHeap) Max() int {
	if mh.Empty() {
		return 0
	}
	return mh.vals[0].first
}

func (mh *MaxHeap) swapNode(i, j int) {
	if i == j {
		return
	}
	mh.idx[mh.vals[i].second], mh.idx[mh.vals[j].second] = mh.idx[mh.vals[j].second], mh.idx[mh.vals[i].second]
	mh.vals[i], mh.vals[j] = mh.vals[j], mh.vals[i]
}

func (mh *MaxHeap) heapIfUp(i int) {
	for i != 0 {
		p := (i - 1) / 2
		if mh.vals[i].first <= mh.vals[p].first {
			return
		}
		mh.swapNode(i, p)
		i = p
	}
}

func (mh *MaxHeap) heapIfDown(i int) {
	for {
		c1 := i*2 + 1
		c2 := i*2 + 2

		// big than max size, return
		if c1 >= mh.size {
			return
		}

		// find a bigger child
		var c int
		if c2 < mh.size && mh.vals[c1].first < mh.vals[c2].first {
			c = c1
		} else {
			c = c2
		}

		// judge if it has become a BigHeap
		if mh.vals[c].first <= mh.vals[i].first {
			return
		}

		mh.swapNode(c, i)
		i = c
	}
}