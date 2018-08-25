// +build js,wasm

package web

import (
	"log"
	"syscall/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D
type CanvasRenderingContext2D interface {
	// canvas
	Canvas() HTMLCanvasElement
	// fillStyle
	FillStyle() string
	SetFillStyle(string)
	// font
	Font() string
	SetFont(string)
	// globalAlpha
	GlobalAlpha() float64
	SetGlobalAlpha(float64)
	// globalCompositeOperation
	GlobalCompositeOperation() string
	SetGlobalCompositeOperation(s string)
	// lineCap
	LineCap() string
	SetLineCap(s string)
	// lineDashOffset
	LineDashOffset() float64
	SetLineDashOffset(s float64)
	// lineJoin
	LineJoin() string
	SetLineJoin(s string)
	// lineWidth
	LineWidth() float64
	SetLineWidth(s float64)
	// miterLimit
	MiterLimit() float64
	SetMiterLimit(s float64)
	// shadowBlur
	ShadowBlur() float64
	SetShadowBlur(s float64)
	// shadowColor
	ShadowColor() string
	SetShadowColor(s string)
	// shadowOffsetX
	ShadowOffsetX() float64
	SetShadowOffsetX(s float64)
	// shadowOffsetY
	ShadowOffsetY() float64
	SetShadowOffsetY(s float64)
	// strokeStyle
	StrokeStyle() string
	SetStrokeStyle(s string)
	// textAlign
	TextAlign() string
	SetTextAlign(s string)
	// textBaseline
	TextBaseline() string
	SetTextBaseline(s string)

	Arc(x, y, radius, startAngle, endAngle float64, anticlockwise ...bool)
	ArcTo(x1, y1, x2, y2, radius float64)
	BeginPath()
	BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64)
	ClearRect(x, y float64, width, height int)
	Clip(fillRule ...string)
	ClosePath()
	CreateImageData(width, height float64) ImageData
	CreateImageData2(img ImageData) ImageData
	CreateLinearGradient(x0, y0, x1, y1 float64) CanvasGradient

	Rect(x, y float64, width, height int)
	Fill()
	MoveTo(x, y float64)
	LineTo(x, y float64)
	FillText(s string, x, y float64)
	Stroke(path ...string)
}

type implCanvasRenderingContext2D struct {
	parent js.Value
	v      js.Value
}

func newCanvasRenderingContext2D(parent, v js.Value) implCanvasRenderingContext2D {
	return implCanvasRenderingContext2D{
		parent: parent,
		v:      v,
	}
}

func (r *implCanvasRenderingContext2D) Canvas() HTMLCanvasElement {
	return newHTMLCanvasElement(r.parent)
}

func (r *implCanvasRenderingContext2D) FillStyle() string {
	return r.v.Get("fillStyle").String()
}

// TODO param type: https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/fillStyle
func (r *implCanvasRenderingContext2D) SetFillStyle(s string) {
	r.v.Set("fillStyle", s)
}

func (r *implCanvasRenderingContext2D) Font() string {
	return r.v.Get("front").String()
}

func (r *implCanvasRenderingContext2D) SetFont(s string) {
	r.v.Set("front", s)
}

// The CanvasRenderingContext2D.globalAlpha property of the Canvas 2D API
// specifies the alpha value that is applied to shapes and images before they are drawn onto
// the canvas. The value is in the range from 0.0 (fully transparent) to 1.0 (fully opaque).
// 透明度，范围：0.0 - 1.0
func (r *implCanvasRenderingContext2D) GlobalAlpha() float64 {
	return r.v.Get("globalAlpha").Float()
}

func (r *implCanvasRenderingContext2D) SetGlobalAlpha(s float64) {
	r.v.Set("globalAlpha", s)
}

// The CanvasRenderingContext2D.globalCompositeOperation property of the
// Canvas 2D API sets the type of compositing operation to apply when drawing new shapes,
// where type is a string identifying which of the compositing or blending mode operations
// to use.
// 设置两个图片重合的绘制方式，很多选项
func (r *implCanvasRenderingContext2D) GlobalCompositeOperation() string {
	return r.v.Get("globalCompositeOperation").String()
}

func (r *implCanvasRenderingContext2D) SetGlobalCompositeOperation(s string) {
	r.v.Set("globalCompositeOperation", s)
}

// The CanvasRenderingContext2D.lineCap property of the Canvas 2D API determines
// how the end points of every line are drawn. There are three possible values for
// this property and those are: butt, round and square. By default this property is set to
// butt.
// 设置线的端点的形状：butt（平的）, round（圆） and square（平的，但是多出去一个圆的半径）
func (r *implCanvasRenderingContext2D) LineCap() string {
	return r.v.Get("lineCap").String()
}

func (r *implCanvasRenderingContext2D) SetLineCap(s string) {
	r.v.Set("lineCap", s)
}

// The CanvasRenderingContext2D.lineDashOffset property of the Canvas 2D API
// sets the line dash pattern offset or "phase" to achieve a "marching ants" effect, for
// example.
// 设置虚线（setLineDash）开始的位置，默认是0
func (r *implCanvasRenderingContext2D) LineDashOffset() float64 {
	return r.v.Get("lineDashOffset").Float()
}

func (r *implCanvasRenderingContext2D) SetLineDashOffset(s float64) {
	r.v.Set("lineDashOffset", s)
}

// The CanvasRenderingContext2D.lineJoin property of the Canvas 2D API
// determines how two connecting segments (of lines, arcs or curves) with non-zero lengths
// in a shape are joined together (degenerate segments with zero lengths, whose specified
// endpoints and control points are exactly at the same position, are skipped).
// 指定不平行的线连接处的形状：bevel（平）, round（圆）, miter（角），默认是miter
func (r *implCanvasRenderingContext2D) LineJoin() string {
	return r.v.Get("lineJoin").String()
}

func (r *implCanvasRenderingContext2D) SetLineJoin(s string) {
	r.v.Set("lineJoin", s)
}

// The CanvasRenderingContext2D.lineWidth property of the Canvas 2D API sets the
// thickness of lines in space units. When getting, it returns the current value (1.0 by
// default). When setting, zero, negative, Infinity and NaN values are ignored; otherwise
// the current value is set to the new value.
// 设置线条宽度，小于等于0等不合法的值将忽略
func (r *implCanvasRenderingContext2D) LineWidth() float64 {
	return r.v.Get("lineWidth").Float()
}

func (r *implCanvasRenderingContext2D) SetLineWidth(s float64) {
	r.v.Set("lineWidth", s)
}

// The CanvasRenderingContext2D.miterLimit property of the Canvas 2D API sets the
// miter limit ratio in space units. When getting, it returns the current value (10.0 by
// default). When setting, zero, negative, Infinity and NaN values are ignored; otherwise
// the current value is set to the new value.
func (r *implCanvasRenderingContext2D) MiterLimit() float64 {
	return r.v.Get("miterLimit").Float()
}

func (r *implCanvasRenderingContext2D) SetMiterLimit(s float64) {
	r.v.Set("miterLimit", s)
}

// The CanvasRenderingContext2D.shadowBlur property of the Canvas 2D API
// specifies the level of the blurring effect; this value doesn't correspond to a number of
// pixels and is not affected by the current transformation matrix. The default value is 0.
// 设定阴影级别，默认0，大于等于0
func (r *implCanvasRenderingContext2D) ShadowBlur() float64 {
	return r.v.Get("shadowBlur").Float()
}

func (r *implCanvasRenderingContext2D) SetShadowBlur(s float64) {
	r.v.Set("shadowBlur", s)
}

// The CanvasRenderingContext2D.shadowColor property of the Canvas 2D API
// specifies the color of the shadow.
// 指定阴影的颜色
func (r *implCanvasRenderingContext2D) ShadowColor() string {
	return r.v.Get("shadowColor").String()
}

func (r *implCanvasRenderingContext2D) SetShadowColor(s string) {
	r.v.Set("shadowColor", s)
}

// The CanvasRenderingContext2D.shadowOffsetX property of the Canvas 2D API
// specifies the distance that the shadow will be offset in horizontal distance.
func (r *implCanvasRenderingContext2D) ShadowOffsetX() float64 {
	return r.v.Get("shadowOffsetX").Float()
}

func (r *implCanvasRenderingContext2D) SetShadowOffsetX(s float64) {
	r.v.Set("shadowOffsetX", s)
}

// The CanvasRenderingContext2D.shadowOffsetY property of the Canvas 2D API
// specifies the distance that the shadow will be offset in vertical distance.
func (r *implCanvasRenderingContext2D) ShadowOffsetY() float64 {
	return r.v.Get("shadowOffsetY").Float()
}

func (r *implCanvasRenderingContext2D) SetShadowOffsetY(s float64) {
	r.v.Set("shadowOffsetY", s)
}

// The CanvasRenderingContext2D.strokeStyle property of the Canvas 2D API
// specifies the color or style to use for the lines around shapes. The default is #000 (black).
// 设定轮廓线的样式
func (r *implCanvasRenderingContext2D) StrokeStyle() string {
	return r.v.Get("strokeStyle").String()
}

func (r *implCanvasRenderingContext2D) SetStrokeStyle(s string) {
	r.v.Set("strokeStyle", s)
}

// The CanvasRenderingContext2D.textAlign property of the Canvas 2D API specifies
// the current text alignment being used when drawing text. Beware that the alignment is
// based on the x value of the fillText() method. So if textAlign is "center", then
// the text would be drawn at x - (width / 2).
// 设定文字对其方式
func (r *implCanvasRenderingContext2D) TextAlign() string {
	return r.v.Get("textAlign").String()
}

func (r *implCanvasRenderingContext2D) SetTextAlign(s string) {
	r.v.Set("textAlign", s)
}

// The CanvasRenderingContext2D.textBaseline property of the Canvas 2D API
// specifies the current text baseline being used when drawing text.
// 设定填充文字的时候，文字基于xy的位置
// 可选值：top, hanging, middle, alphabetic, ideographic, bottom
func (r *implCanvasRenderingContext2D) TextBaseline() string {
	return r.v.Get("textBaseline").String()
}

// The CanvasRenderingContext2D.arc() method of the Canvas 2D API adds an arc to
// the path which is centered at (x, y) position with radius r starting at startAngle and ending
// at endAngle going in the given direction by anticlockwise (defaulting to clockwise).
//
// 画圆弧，参数分别是圆心点，半径，开始和结束的弧度，最后一个默认false顺时针，true为逆时针
func (r *implCanvasRenderingContext2D) Arc(x, y, radius, startAngle, endAngle float64, anticlockwise ...bool) {
	if len(anticlockwise) > 0 {
		r.v.Call("arc", x, y, radius, startAngle, endAngle, anticlockwise[0])
	} else {
		r.v.Call("arc", x, y, radius, startAngle, endAngle)
	}
}

// The CanvasRenderingContext2D.arcTo() method of the Canvas 2D API adds an arc
// to the path with the given control points and radius.
//
// The arc drawn will be a part of a circle, never elliptical. Typical use could be making a
// rounded corner.
//
// One way to think about the arc drawn is to imagine two straight segments, from the
// starting point (latest point in current path) to the first control point, and then from the
// first control point to the second control point. These two segments form a sharp corner
// with the first control point being in the corner. Using arcTo, the corner will instead be an
// arc with the given radius.
//
// The arc is tangential to both segments, which can sometimes produce surprising results,
// e.g. if the radius given is larger than the distance between the starting point and the first
// control point.
//
// If the radius specified doesn't make the arc meet the starting point (latest point in the
// current path), the starting point is connected to the arc with a straight line segment.
func (r *implCanvasRenderingContext2D) ArcTo(x1, y1, x2, y2, radius float64) {
	r.v.Call("arcTo", x1, y1, x2, y2, radius)
}

// The CanvasRenderingContext2D.beginPath() method of the Canvas 2D API starts a
// new path by emptying the list of sub-paths. Call this method when you want to create a
// new path.
func (r *implCanvasRenderingContext2D) BeginPath() {
	r.v.Call("beginPath")
}

func (r *implCanvasRenderingContext2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	r.v.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// The CanvasRenderingContext2D.clearRect() method of the Canvas 2D API sets all
// pixels in the rectangle defined by starting point (x, y) and size (width, height) to transparent
// black, erasing any previously drawn content.
func (r *implCanvasRenderingContext2D) ClearRect(x, y float64, width, height int) {
	r.v.Call("clearRect", x, y, width, height)
}

// TODO: unsupport path
func (r *implCanvasRenderingContext2D) Clip(fillRule ...string) {
	if len(fillRule) > 0 {
		r.v.Call("clip", fillRule[0])
	} else {
		r.v.Call("clip")
	}
}

func (r *implCanvasRenderingContext2D) ClosePath() {
	r.v.Call("closePath")
}

func (r *implCanvasRenderingContext2D) CreateImageData(width, height float64) ImageData {
	v := r.v.Call("createImageData", width, height)
	t := newImplImageData(v)
	return &t
}

func (r *implCanvasRenderingContext2D) CreateImageData2(img ImageData) ImageData {
	v, ok := img.(*implImageData)
	if ok {
		v := r.v.Call("createImageData", v.v)
		t := newImplImageData(v)
		return &t
	} else {
		log.Println("ImageData is not implImageData")
		return img
	}
}

func (r *implCanvasRenderingContext2D) CreateLinearGradient(x0, y0, x1, y1 float64) CanvasGradient {
	v := r.v.Call("createLinearGradient", x0, y0, x1, y1)
	t := newImplCanvasGradient(v)
	return &t
}

func (r *implCanvasRenderingContext2D) createPattern(image CanvasImageSource, repetition string) CanvasPattern {
	v := r.v.Call("createPattern", image, repetition)
	t := newImplCanvasPattern(v)
	return &t
}

func (r *implCanvasRenderingContext2D) SetTextBaseline(s string) {
	r.v.Set("textBaseline", s)
}

func (r *implCanvasRenderingContext2D) Rect(x, y float64, width, height int) {
	r.v.Call("rect", x, y, width, height)
}

func (r *implCanvasRenderingContext2D) Fill() {
	r.v.Call("fill")
}

func (r *implCanvasRenderingContext2D) MoveTo(x, y float64) {
	r.v.Call("moveTo", x, y)
}

func (r *implCanvasRenderingContext2D) LineTo(x, y float64) {
	r.v.Call("lineTo", x, y)
}

func (r *implCanvasRenderingContext2D) FillText(s string, x, y float64) {
	r.v.Call("fillText", s, x, y)
}

// The CanvasRenderingContext2D.stroke() method of the Canvas 2D API strokes the
// current or given path with the current stroke style using the non-zero winding rule.
func (r *implCanvasRenderingContext2D) Stroke(path ...string) {
	if len(path) > 0 {
		r.v.Call("stroke", path[0])
	} else {
		r.v.Call("stroke")
	}
}
