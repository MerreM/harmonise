// harmonise project harmonise.go
package harmonise

// TODO For hashing look at https://github.com/stathat/consistent
import (
	"math"
)

const (
	k int = 6
)

var (
	MAX = int(math.Pow(2, float64(k)))
)

func decrease(value, size int) int {
	if size <= value {
		return value - size
	} else {
		return MAX - (size - value)
	}
}

func between(value, init, end int) bool {
	if init == end {
		return true
	} else if init > end {
		shift := MAX - init
		init = 0
		end = (end + shift) % MAX
		value = (value + shift) % MAX
	}
	result := init < value
	result = result && value < end
	return result
}

func Ebetween(value, init, end int) bool {
	if value == init {
		return true
	} else {
		return between(value, init, end)
	}
}

func betweenE(value, init, end int) bool {
	if value == end {
		return true
	} else {
		return between(value, init, end)
	}
}

type Node struct {
	id          int
	finger      map[int]*Node
	start       map[int]int
	predecessor *Node
}

func CreateStartId(id, i int) int {
	result := id + int(math.Pow(2, float64(i)))%int(math.Pow(2, float64(k)))
	return result
}
func CreateId(id int) int {
	return id
}

func Create(id int) Node {
	startMap := make(map[int]int)
	for i := 0; i < 10; i += 1 {
		startMap[i] = CreateStartId(id, i)
	}
	self := Node{
		CreateId(id),
		make(map[int]*Node),
		startMap,
		nil,
	}
	self.predecessor = nil
	return self
}

func (self *Node) Successor() *Node {
	return self.finger[0]
}

func (self *Node) Predecessor() *Node {
	if self.predecessor == nil {
		return self
	} else {
		return self.predecessor
	}

}

func (self *Node) FindSuccessor(id int) *Node {
	if betweenE(id, self.Predecessor().id, self.id) {
		return self
	}
	predecessor := self.FindPredecessor(id)
	return predecessor.Successor()
}

func (self *Node) FindPredecessor(id int) *Node {
	if self.id == id {
		return self.Predecessor()
	}
	n1 := self
	for !betweenE(id, n1.id, n1.Successor().id) {
		n1 = n1.ClosestPrecedingFinger(id)
	}
	return n1

}
func (self *Node) ClosestPrecedingFinger(id int) *Node {
	for i := k - 1; i < (-1); i -= 1 {
		if between(self.finger[id].id, self.id, id) {
			return self.finger[id]
		}
	}
	return self
}

func (self *Node) Join(node *Node) {
	if self == node {
		for i := 0; i < k; i += 1 {
			self.finger[i] = node
		}
	} else {
		self.InitFingerTable(node)
		self.UpdateOthers()
	}
}

func (self *Node) InitFingerTable(node *Node) {
	self.finger[0] = node.FindSuccessor(self.start[0])
	self.predecessor = self.Successor().Predecessor()
	self.Successor().predecessor = self
	self.Predecessor().finger[0] = self
	for i := 0; i < (k - 1); i += 1 {
		if Ebetween(self.start[i+1], self.id, self.finger[i].id) {
			self.finger[i+1] = self.finger[i]
		} else {
			self.finger[i+1] = node.FindSuccessor(self.start[i+1])
		}
	}

}
func (self *Node) UpdateOthers() {
	for i := 0; i < k; i += 1 {
		prev := decrease(self.id, int(math.Pow(2, float64(i))))
		p := self.FindPredecessor(prev)
		if prev == p.Successor().id {
			p = p.Successor()
		}
		p.UpdateFingerTable(self, i)
	}
}

func (self *Node) UpdateFingerTable(node *Node, index int) {
	if Ebetween(node.id, self.id, self.finger[index].id) &&
		self.id != node.id {
		self.finger[index] = node
	}
	p := self.Predecessor()
	p.UpdateFingerTable(node, index)
}

func (self *Node) UpdateOthersLeave() {
	for i := 0; i < k; i += 1 {
		prev := decrease(self.id, int(math.Pow(2, float64(i))))
		p := self.FindPredecessor(prev)
		p.UpdateFingerTable(self.Successor(), i)
	}
}
func (self *Node) Leave() {
	self.Successor().predecessor = self.predecessor
	self.Predecessor().SetSuccessor(self.Successor())
	self.UpdateOthersLeave()
}
func (self *Node) SetSuccessor(node *Node) {
	self.finger[0] = node
}
