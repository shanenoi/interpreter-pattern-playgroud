package main

// Expression is AbstractExpression
type Expression interface {
	// Evaluate is interpret
	Evaluate() float64
}

// Number is TerminalExpression
type Number struct {
	value float64
}

// Evaluate is interpret for TerminalExpression
func (n *Number) Evaluate() float64 {
	return n.value
}

// Add is Non-terminalExpression
type Add struct {
	Left  Expression
	Right Expression
}

// Evaluate is interpret for TerminalExpression
func (a *Add) Evaluate() float64 {
	return a.Left.Evaluate() + a.Right.Evaluate()
}

// Subtract is Non-terminalExpression
type Subtract struct {
	Left  Expression
	Right Expression
}

// Evaluate is interpret for TerminalExpression
func (s *Subtract) Evaluate() float64 {
	return s.Left.Evaluate() - s.Right.Evaluate()
}

// Multiply is Non-terminalExpression
type Multiply struct {
	Left  Expression
	Right Expression
}

// Evaluate is interpret for TerminalExpression
func (m *Multiply) Evaluate() float64 {
	return m.Left.Evaluate() * m.Right.Evaluate()
}

// Divide is Non-terminalExpression
type Divide struct {
	Left  Expression
	Right Expression
}

// Evaluate is interpret for TerminalExpression
func (d *Divide) Evaluate() float64 {
	return d.Left.Evaluate() / d.Right.Evaluate()
}
