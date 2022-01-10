package hhcm

import (
	"fmt"
  "math"
	"hash/fnv"
	"sort"
)

type HHCM struct {
	depth int
	width int
	numHitters int
  heavyHitters map[string]uint64
  smallest int
	full bool
	bins []uint64
	lowestTopFrequency uint64
	lowestTopFrequencyPath string
}

func NewHHCM (depth int, width int, numHitters int) *HHCM {
  hhcm := new(HHCM)
	hhcm.depth = depth
	hhcm.width = width
	hhcm.numHitters = numHitters
	hhcm.bins = make([]uint64, depth*width)
	hhcm.heavyHitters = make(map[string]uint64)
  return hhcm
}

/* insert value into HHCM underlying data structure,
 * using the Count-Min-Sketch method for processing
 * streaming data for approximate Heavy Hitters reporting
 *
 * This code was influenced by the code in this project
 * https://github.com/barrust/count-min-sketch, written
 * in Python
 */
func (h *HHCM) Insert (path string) error {
  depth := h.depth
  width := h.width

  tmp := path
  start_idx := 0
	res := uint64(math.MaxInt64)
  for i := 0; i < depth; i++ {
		hash := fnv.New64a()
		hash.Write([]byte(tmp))
		hval := hash.Sum64()

    tbin := int((hval % uint64(width))) + start_idx

    h.bins[tbin]++
    start_idx += width
    tmp = fmt.Sprintf("{0:/%x}", hval)

		b := h.bins[tbin]
		if res > b {
			res = b
		}
  }

	if res > h.lowestTopFrequency {
    hh := h.heavyHitters

		_, already_present := hh[path]
		hh[path] = res
		set_new_min := false

		if path == h.lowestTopFrequencyPath {
			set_new_min = true
		} else if !already_present{
			if h.full {
				delete(hh, h.lowestTopFrequencyPath)
			} else if len(hh) == h.numHitters {
				h.full = true
			}

			set_new_min = h.full
		}

    if set_new_min {
			min := uint64(0)
			min_p := ""
	    for p, f := range hh {
			  	if min == 0 || f < min {
						min = f
						min_p = p
					}
			}

	    h.lowestTopFrequency = min
			h.lowestTopFrequencyPath = min_p
		}
	}

	return nil
}

/* print the N most frequent observed paths, where
 * N = h.numHitters, in descending order by frequency
 */
func (h *HHCM) Report () error {
  fmt.Printf("Top %d paths:\n", h.numHitters)

	m := make(map[uint64][]string)
	vals := make([]uint64, 0)

	hh := h.heavyHitters
	for path, val := range hh {
		o_paths, present := m[val]
		if present {
			m[val] = append(o_paths, path)
		} else {
			m[val] = []string{path}
			vals = append(vals, val)
		}
	}

	sort.Slice(vals, func(i, j int) bool { return vals[i] > vals[j] })
	for _, val := range vals {
		for _, path := range m[val] {
			fmt.Println(path)
		}
	}

	return nil
}
