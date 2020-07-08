package Self

/*
1. 合并k个有序数组 -- mergeKArray
*/

/*
创建一个大小为N的数组保存最后的结果
数组本身已经从小到大排好序，所以我们只需创建一个大小为k的最小堆，堆中初始元素为k个数组中的每个数组的第一个元素，
每次从堆中取出最小元素（堆顶元素），并将其存入输出数组中，将堆顶元素所在行的下一元素加入堆，重新排列出堆顶元素，
时间复杂度为logk，总共N个元素，所以总体时间复杂度是Nlogk
*/

type Ele struct {
	val      int
	idx      int
	arrayIdx int
}

type Stack []Ele

func (this *Stack) Add(e Ele) {
	if len(*this) == 0 {
		*this = append(*this, e)
	} else {
		insertIdx := -1
		for i := 0; i < len(*this); i++ {
			if (*this)[i].val > e.val {
				insertIdx = i
				break
			}
		}
		if insertIdx == 0 {
			tmp := append([]Ele{e}, *this...)
			*this = tmp
		} else if insertIdx == -1 {
			*this = append(*this, e)
		} else {
			*this = append(*this, Ele{})
			for i := len(*this) - 1; i > insertIdx; i-- {
				(*this)[i] = (*this)[i-1]
			}
			(*this)[insertIdx] = e
		}
	}
}

func (this *Stack) Pop() (e Ele) {
	e = (*this)[0]
	if len(*this) == 1 {
		*this = make([]Ele, 0)
	} else {
		*this = (*this)[1:]
	}
	return
}

func mergeKArray(input [][]int, k int) (ret []int) {
	stack := Stack{}
	ret = make([]int, 0)
	for i, arr := range input {
		e := Ele{
			val:      arr[0],
			idx:      0,
			arrayIdx: i,
		}
		stack.Add(e)
	}
	for len(stack) != 0 {
		val := stack.Pop()
		ret = append(ret, val.val)
		if val.idx < len(input[val.arrayIdx])-1 {
			newE := Ele{
				val:      input[val.arrayIdx][val.idx+1],
				arrayIdx: val.arrayIdx,
				idx:      val.idx + 1,
			}
			stack.Add(newE)
		}

	}
	return
}
