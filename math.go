package main

type Expression interface {
	Evaluate() float64
}

type Number struct {
	value float64
}

func (n *Number) Evaluate() float64 {
	return n.value
}

type Add struct {
	Left  Expression
	Right Expression
}

func (a *Add) Evaluate() float64 {
	return a.Left.Evaluate() + a.Right.Evaluate()
}

type Subtract struct {
	Left  Expression
	Right Expression
}

func (s *Subtract) Evaluate() float64 {
	return s.Left.Evaluate() - s.Right.Evaluate()
}

type Multiply struct {
	Left  Expression
	Right Expression
}

func (m *Multiply) Evaluate() float64 {
	return m.Left.Evaluate() * m.Right.Evaluate()
}

type Divide struct {
	Left  Expression
	Right Expression
}

func (d *Divide) Evaluate() float64 {
	return d.Left.Evaluate() / d.Right.Evaluate()
}
