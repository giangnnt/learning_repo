package main

import (
	"fmt"
	"slices"
)

func main() {
	// state
	mergeArray := []int{}
	blockPerRun := 4 // 4 blocks equal to M = 4
	blockLoaded := 0
	b := buffer{}
	runs := []run{}
	blks := []block{
		{data: [4]int{17, 3, 25, 9}},
		{data: [4]int{8, 21, 14, 30}},
		{data: [4]int{11, 2, 18, 27}},
		{data: [4]int{6, 22, 13, 19}},
		{data: [4]int{5, 29, 1, 16}},
		{data: [4]int{24, 7, 12, 20}},
		{data: [4]int{4, 28, 15, 10}},
	}
	// stage 1: create runs
	// load block to buffer and sort to run
	for blockLoaded < len(blks) {
		end := blockLoaded + blockPerRun
		if end > len(blks) {
			end = len(blks)
		}
		b.loadBlock(blks[blockLoaded:end])
		run := b.sortToRun()
		runs = append(runs, run)
		b.flush()
		blockLoaded = end
		fmt.Printf("Run: %+v\n", run.data)
	}
	// stage 2: merge runs
	b.runCount = len(runs)
outer:
	for len(runs) > 0 {
		// constantly check if block is empty
		runIdx := b.checkEmptyBlock()
		if runIdx == -1 {
			// all block is not empty, take min value
			mergeArray = append(mergeArray, b.takeMinValue())
		} else {
			if runs[runIdx].isClosed || runs[runIdx].isEmpty() {
				runs[runIdx].isClosed = true
				b.data[runIdx] = block{}
				runs = append(runs[:runIdx], runs[runIdx+1:]...)
				b.runCount = len(runs)
				if len(runs) == 0 {
					break outer
				}
				continue
			}
			// load block from run to buffer
			if !b.loadRun(runs, runIdx) {
				// if run is empty or closed, remove it from runs
				runs[runIdx].isClosed = true
				b.data[runIdx] = block{}
				runs = append(runs[:runIdx], runs[runIdx+1:]...)
				b.runCount = len(runs)
				if len(runs) == 0 {
					break outer
				}
			}
		}
	}
	fmt.Printf("Merge result: %+v\n", mergeArray)
}

// simulate fix size buffer
type buffer struct {
	data     [4]block
	u_data   [4 * 4]int
	loaded   int
	runCount int
}

type run struct {
	data     []block
	isClosed bool
}

type block struct {
	data [4]int
}

func (r *run) isEmpty() bool {
	for _, blk := range r.data {
		if !blk.isEmpty() {
			return false
		}
	}
	return true
}

func (b *block) unmarshal() []int {
	return b.data[:]
}

func (b *block) isEmpty() bool {
	for _, v := range b.data {
		if v != 0 {
			return false
		}
	}
	return true
}

func (b *buffer) loadBlock(blks []block) {
	b.loaded = 0
	for _, blk := range blks {
		n := copy(b.u_data[b.loaded:], blk.unmarshal())
		b.loaded += n
	}
}

func (b *buffer) checkEmptyBlock() (blockIdx int) {
	for i := 0; i < b.runCount; i++ {
		if b.data[i].isEmpty() {
			return i
		}
	}
	return -1
}

// if load return true, else return false
func (b *buffer) loadRun(runs []run, runIdx int) bool {
	if runIdx < 0 || runIdx >= b.runCount || runIdx >= len(runs) {
		return false
	}

	run := &runs[runIdx]
	if run.isClosed || run.isEmpty() {
		run.isClosed = true
		return false
	}
	// load not empty block from run to buffer
	for i, blk := range run.data {
		if !blk.isEmpty() {
			b.data[runIdx] = blk
			// remove block from run
			run.data[i] = block{}
			return true
		}
	}

	run.isClosed = true
	return false
}

func (b *buffer) takeMinValue() int {
	min := 0
	firstValues := []int{}
	arrQueue := make([][]int, b.runCount)
	for i := 0; i < b.runCount; i++ {
		arrQueue[i] = append(arrQueue[i], b.data[i].unmarshal()...)
	}
	for _, arr := range arrQueue {
		// always take first value, block check is done in main loop
		firstValues = append(firstValues, arr[0])
	}
	min = slices.Min(firstValues)
	for i, v := range firstValues {
		if v == min {
			arrQueue[i] = arrQueue[i][1:] // remove first value
			// update block data
			clear(b.data[i].data[:])
			copy(b.data[i].data[:], arrQueue[i])
		}
	}
	return min
}

func (b *buffer) sortToRun() run {
	r := run{}
	if b.loaded == 0 {
		return r
	}
	slices.Sort(b.u_data[:b.loaded])
	for i := 0; i < b.loaded; i += 4 {
		var blk block
		end := i + 4
		if end > b.loaded {
			end = b.loaded
		}
		copy(blk.data[:], b.u_data[i:end])
		r.data = append(r.data, blk)
	}
	return r
}

func (b *buffer) flush() {
	b.data = [4]block{}
	b.u_data = [4 * 4]int{}
	b.loaded = 0
}
