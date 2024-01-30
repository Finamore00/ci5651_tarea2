package datatypes

type Interval struct {
	Start  uint32
	Length uint32
}

func (inter *Interval) End() uint32 {
	return inter.Start + inter.Length
}

func (inter1 *Interval) HasIntersectionWith(inter2 *Interval) bool {
	return inter1.End() > inter2.Start && inter1.Start < inter2.End()
}
