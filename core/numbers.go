package core

import (
	"math/big"
)

type (
	Number interface {
		Object
		Int() Int
		Double() Double
		BigInt() *big.Int
		BigFloat() *big.Float
		Ratio() *big.Rat
	}
	Ops interface {
		Combine(ops Ops) Ops
		Add(Number, Number) Number
		Subtract(Number, Number) Number
		Multiply(Number, Number) Number
		Divide(Number, Number) Number
		IsZero(Number) bool
		Lt(Number, Number) bool
		Lte(Number, Number) bool
		Gt(Number, Number) bool
		Gte(Number, Number) bool
		Eq(Number, Number) bool
		Quotient(Number, Number) Number
		Rem(Number, Number) Number
	}
	IntOps      struct{}
	DoubleOps   struct{}
	BigIntOps   struct{}
	BigFloatOps struct{}
	RatioOps    struct{}
)

const MAX_INT = int(^uint(0) >> 1)
const MIN_INT = -MAX_INT - 1
const MAX_RUNE = int(^uint32(0) >> 1)
const MIN_RUNE = -MAX_RUNE - 1

var (
	INT_OPS      = IntOps{}
	DOUBLE_OPS   = DoubleOps{}
	BIGINT_OPS   = BigIntOps{}
	BIGFLOAT_OPS = BigFloatOps{}
	RATIO_OPS    = RatioOps{}
)

func (ops IntOps) Combine(other Ops) Ops {
	return other
}

func (ops DoubleOps) Combine(other Ops) Ops {
	switch other.(type) {
	case BigFloatOps:
		return other
	default:
		return ops
	}
}

func (ops BigIntOps) Combine(other Ops) Ops {
	switch other.(type) {
	case IntOps:
		return ops
	default:
		return other
	}
}

func (ops BigFloatOps) Combine(other Ops) Ops {
	return ops
}

func (ops RatioOps) Combine(other Ops) Ops {
	switch other.(type) {
	case DoubleOps, BigFloatOps:
		return other
	default:
		return ops
	}
}

func GetOps(obj Object) Ops {
	switch obj.(type) {
	case Double:
		return DOUBLE_OPS
	case *BigInt:
		return BIGINT_OPS
	case *BigFloat:
		return BIGFLOAT_OPS
	case *Ratio:
		return RATIO_OPS
	default:
		return INT_OPS
	}
}

// Int conversions

func (i Int) Int() Int {
	return i
}

func (i Int) Double() Double {
	return Double{D: float64(i.I)}
}

func (i Int) BigInt() *big.Int {
	return big.NewInt(int64(i.I))
}

func (i Int) BigFloat() *big.Float {
	return big.NewFloat(float64(i.I))
}

func (i Int) Ratio() *big.Rat {
	return big.NewRat(int64(i.I), 1)
}

// Double conversions

func (d Double) Int() Int {
	return Int{I: int(d.D)}
}

func (d Double) BigInt() *big.Int {
	return big.NewInt(int64(d.D))
}

func (d Double) Double() Double {
	return d
}

func (d Double) BigFloat() *big.Float {
	return big.NewFloat(float64(d.D))
}

func (d Double) Ratio() *big.Rat {
	res := big.Rat{}
	return res.SetFloat64(float64(d.D))
}

// BigInt conversions

func (b *BigInt) Int() Int {
	return Int{I: int(b.BigInt().Int64())}
}

func (b *BigInt) BigInt() *big.Int {
	return &b.b
}

func (b *BigInt) Double() Double {
	return Double{D: float64(b.BigInt().Int64())}
}

func (b *BigInt) BigFloat() *big.Float {
	res := big.Float{}
	return res.SetInt(b.BigInt())
}

func (b *BigInt) Ratio() *big.Rat {
	res := big.Rat{}
	return res.SetInt(b.BigInt())
}

// BigFloat conversions

func (b *BigFloat) Int() Int {
	i, _ := b.BigFloat().Int64()
	return Int{I: int(i)}
}

func (b *BigFloat) BigInt() *big.Int {
	bi, _ := b.BigFloat().Int(nil)
	return bi
}

func (b *BigFloat) Double() Double {
	f, _ := b.BigFloat().Float64()
	return Double{D: f}
}

func (b *BigFloat) BigFloat() *big.Float {
	return &b.b
}

func (b *BigFloat) Ratio() *big.Rat {
	res := big.Rat{}
	return res.SetFloat64(float64(b.Double().D))
}

// Ratio conversions

func (r *Ratio) Int() Int {
	f, _ := r.Ratio().Float64()
	return Int{I: int(f)}
}

func (r *Ratio) BigInt() *big.Int {
	f, _ := r.Ratio().Float64()
	return big.NewInt(int64(f))
}

func (r *Ratio) Double() Double {
	f, _ := r.Ratio().Float64()
	return Double{D: f}
}

func (r *Ratio) BigFloat() *big.Float {
	f, _ := r.Ratio().Float64()
	return big.NewFloat(f)
}

func (r *Ratio) Ratio() *big.Rat {
	return &r.r
}

// Ops

// Add

func (ops IntOps) Add(x, y Number) Number {
	return Int{I: x.Int().I + y.Int().I}
}

func (ops DoubleOps) Add(x, y Number) Number {
	return Double{D: x.Double().D + y.Double().D}
}

func (ops BigIntOps) Add(x, y Number) Number {
	b := big.Int{}
	b.Add(x.BigInt(), y.BigInt())
	res := BigInt{b: b}
	return &res
}

func (ops BigFloatOps) Add(x, y Number) Number {
	b := big.Float{}
	b.Add(x.BigFloat(), y.BigFloat())
	res := BigFloat{b: b}
	return &res
}

func (ops RatioOps) Add(x, y Number) Number {
	r := big.Rat{}
	r.Add(x.Ratio(), y.Ratio())
	res := Ratio{r: r}
	return &res
}

// Subtract

func (ops IntOps) Subtract(x, y Number) Number {
	return Int{I: x.Int().I - y.Int().I}
}

func (ops DoubleOps) Subtract(x, y Number) Number {
	return Double{D: x.Double().D - y.Double().D}
}

func (ops BigIntOps) Subtract(x, y Number) Number {
	b := big.Int{}
	b.Sub(x.BigInt(), y.BigInt())
	res := BigInt{b: b}
	return &res
}

func (ops BigFloatOps) Subtract(x, y Number) Number {
	b := big.Float{}
	b.Sub(x.BigFloat(), y.BigFloat())
	res := BigFloat{b: b}
	return &res
}

func (ops RatioOps) Subtract(x, y Number) Number {
	r := big.Rat{}
	r.Sub(x.Ratio(), y.Ratio())
	res := Ratio{r: r}
	return &res
}

// Multiply

func (ops IntOps) Multiply(x, y Number) Number {
	return Int{I: x.Int().I * y.Int().I}
}

func (ops DoubleOps) Multiply(x, y Number) Number {
	return Double{D: x.Double().D * y.Double().D}
}

func (ops BigIntOps) Multiply(x, y Number) Number {
	b := big.Int{}
	b.Mul(x.BigInt(), y.BigInt())
	res := BigInt{b: b}
	return &res
}

func (ops BigFloatOps) Multiply(x, y Number) Number {
	b := big.Float{}
	b.Mul(x.BigFloat(), y.BigFloat())
	res := BigFloat{b: b}
	return &res
}

func (ops RatioOps) Multiply(x, y Number) Number {
	r := big.Rat{}
	r.Mul(x.Ratio(), y.Ratio())
	res := Ratio{r: r}
	return &res
}

// Divide

func (ops IntOps) Divide(x, y Number) Number {
	if y.Int().I == 0 {
		panic(RT.NewError("Division by zero"))
	}
	b := big.NewRat(int64(x.Int().I), int64(y.Int().I))
	if b.IsInt() {
		return Int{I: int(b.Num().Int64())}
	}
	res := Ratio{r: *b}
	return &res
}

func (ops DoubleOps) Divide(x, y Number) Number {
	return Double{D: x.Double().D / y.Double().D}
}

func (ops BigIntOps) Divide(x, y Number) Number {
	if y.Ratio().Num().Int64() == 0 {
		panic(RT.NewError("Division by zero"))
	}
	b := big.Rat{}
	b.Quo(x.Ratio(), y.Ratio())
	if b.IsInt() {
		res := BigInt{b: *b.Num()}
		return &res
	}
	res := Ratio{r: b}
	return &res
}

func (ops BigFloatOps) Divide(x, y Number) Number {
	b := big.Float{}
	b.Quo(x.BigFloat(), y.BigFloat())
	res := BigFloat{b: b}
	return &res
}

func (ops RatioOps) Divide(x, y Number) Number {
	if y.Ratio().Num().Int64() == 0 {
		panic(RT.NewError("Division by zero"))
	}
	r := big.Rat{}
	r.Quo(x.Ratio(), y.Ratio())
	res := Ratio{r: r}
	return &res
}

// Quotient

func (ops IntOps) Quotient(x, y Number) Number {
	return Int{I: x.Int().I / y.Int().I}
}

func (ops DoubleOps) Quotient(x, y Number) Number {
	z := x.Double().D / y.Double().D
	if z <= float64(MAX_INT) && z >= float64(MIN_INT) {
		return Int{I: int(z)}
	}
	return &BigInt{b: *big.NewInt(int64(z))}
}

func (ops BigIntOps) Quotient(x, y Number) Number {
	z := big.Int{}
	z.Quo(x.BigInt(), y.BigInt())
	return &BigInt{b: z}
}

func (ops BigFloatOps) Quotient(x, y Number) Number {
	z := big.Float{}
	i, _ := z.Quo(x.BigFloat(), y.BigFloat()).Int64()
	return &BigFloat{b: *z.SetInt64(i)}
}

func (ops RatioOps) Quotient(x, y Number) Number {
	z := big.Rat{}
	f, _ := z.Quo(x.Ratio(), y.Ratio()).Float64()
	return &BigInt{b: *big.NewInt(int64(f))}
}

// Remainder

func (ops IntOps) Rem(x, y Number) Number {
	return Int{I: x.Int().I % y.Int().I}
}

func (ops DoubleOps) Rem(x, y Number) Number {
	n := x.Double().D
	d := y.Double().D
	z := n / d
	if z <= float64(MAX_INT) && z >= float64(MIN_INT) {
		return Double{D: n - float64(int(z))*d}
	}
	return Double{D: n - float64(int64(z))*d}
}

func (ops BigIntOps) Rem(x, y Number) Number {
	z := big.Int{}
	z.Rem(x.BigInt(), y.BigInt())
	return &BigInt{b: z}
}

func (ops BigFloatOps) Rem(x, y Number) Number {
	n := x.BigFloat()
	d := y.BigFloat()
	z := big.Float{}
	i, _ := z.Quo(n, d).Int64()
	d.Mul(d, big.NewFloat(float64(i)))
	z.Sub(n, d)
	return &BigFloat{b: z}
}

func (ops RatioOps) Rem(x, y Number) Number {
	n := x.Ratio()
	d := y.Ratio()
	z := big.Rat{}
	f, _ := z.Quo(n, d).Float64()
	d.Mul(d, big.NewRat(int64(f), 1))
	z.Sub(n, d)
	return &Ratio{r: z}
}

// IsZero

func (ops IntOps) IsZero(x Number) bool {
	return x.Int().I == 0
}

func (ops DoubleOps) IsZero(x Number) bool {
	return x.Double().D == 0
}

func (ops BigIntOps) IsZero(x Number) bool {
	return x.BigInt().Sign() == 0
}

func (ops BigFloatOps) IsZero(x Number) bool {
	return x.BigFloat().Sign() == 0
}

func (ops RatioOps) IsZero(x Number) bool {
	return x.Ratio().Sign() == 0
}

// Lt

func (ops IntOps) Lt(x Number, y Number) bool {
	return x.Int().I < y.Int().I
}

func (ops DoubleOps) Lt(x Number, y Number) bool {
	return x.Double().D < y.Double().D
}

func (ops BigIntOps) Lt(x Number, y Number) bool {
	return x.BigInt().Cmp(y.BigInt()) < 0
}

func (ops BigFloatOps) Lt(x Number, y Number) bool {
	return x.BigFloat().Cmp(y.BigFloat()) < 0
}

func (ops RatioOps) Lt(x Number, y Number) bool {
	return x.Ratio().Cmp(y.Ratio()) < 0
}

// Lte

func (ops IntOps) Lte(x Number, y Number) bool {
	return x.Int().I <= y.Int().I
}

func (ops DoubleOps) Lte(x Number, y Number) bool {
	return x.Double().D <= y.Double().D
}

func (ops BigIntOps) Lte(x Number, y Number) bool {
	return x.BigInt().Cmp(y.BigInt()) <= 0
}

func (ops BigFloatOps) Lte(x Number, y Number) bool {
	return x.BigFloat().Cmp(y.BigFloat()) <= 0
}

func (ops RatioOps) Lte(x Number, y Number) bool {
	return x.Ratio().Cmp(y.Ratio()) <= 0
}

// Gt

func (ops IntOps) Gt(x Number, y Number) bool {
	return x.Int().I > y.Int().I
}

func (ops DoubleOps) Gt(x Number, y Number) bool {
	return x.Double().D > y.Double().D
}

func (ops BigIntOps) Gt(x Number, y Number) bool {
	return x.BigInt().Cmp(y.BigInt()) > 0
}

func (ops BigFloatOps) Gt(x Number, y Number) bool {
	return x.BigFloat().Cmp(y.BigFloat()) > 0
}

func (ops RatioOps) Gt(x Number, y Number) bool {
	return x.Ratio().Cmp(y.Ratio()) > 0
}

// Gte

func (ops IntOps) Gte(x Number, y Number) bool {
	return x.Int().I >= y.Int().I
}

func (ops DoubleOps) Gte(x Number, y Number) bool {
	return x.Double().D >= y.Double().D
}

func (ops BigIntOps) Gte(x Number, y Number) bool {
	return x.BigInt().Cmp(y.BigInt()) >= 0
}

func (ops BigFloatOps) Gte(x Number, y Number) bool {
	return x.BigFloat().Cmp(y.BigFloat()) >= 0
}

func (ops RatioOps) Gte(x Number, y Number) bool {
	return x.Ratio().Cmp(y.Ratio()) >= 0
}

// Eq

func (ops IntOps) Eq(x Number, y Number) bool {
	return x.Int().I == y.Int().I
}

func (ops DoubleOps) Eq(x Number, y Number) bool {
	return x.Double().D == y.Double().D
}

func (ops BigIntOps) Eq(x Number, y Number) bool {
	return x.BigInt().Cmp(y.BigInt()) == 0
}

func (ops BigFloatOps) Eq(x Number, y Number) bool {
	return x.BigFloat().Cmp(y.BigFloat()) == 0
}

func (ops RatioOps) Eq(x Number, y Number) bool {
	return x.Ratio().Cmp(y.Ratio()) == 0
}

func CompareNumbers(x Number, y Number) int {
	ops := GetOps(x).Combine(GetOps(y))
	if ops.Lt(x, y) {
		return -1
	}
	if ops.Lt(y, x) {
		return 1
	}
	return 0
}

func Max(x Number, y Number) Number {
	ops := GetOps(x).Combine(GetOps(y))
	if ops.Lt(x, y) {
		return y
	}
	return x
}

func Min(x Number, y Number) Number {
	ops := GetOps(x).Combine(GetOps(y))
	if ops.Lt(x, y) {
		return x
	}
	return y
}
