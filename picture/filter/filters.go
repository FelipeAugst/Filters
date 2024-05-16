package filter

import (
	"errors"

	"github.com/disintegration/gift"
)

var (
	noParamsOp map[string]any
	oneFloatOp map[string]any
	rgb        map[string]any
)

type Effect interface {
	generateFilter() gift.Filter
}

type noParams struct {
	name      string
	generator func() gift.Filter
}

func (n noParams) generateFilter() gift.Filter {
	return n.generator()
}

func (n noParams) getName() string {
	return n.name
}

type oneFloat struct {
	name      string
	param     float32
	generator func(float32) gift.Filter
}

func (o oneFloat) generateFilter() gift.Filter {
	return o.generator(o.param)
}

func (o oneFloat) getName() string {
	return o.name
}

type rgbOp struct {
	name      string
	r, g, b   float32
	generator func(float32, float32, float32) gift.Filter
}

func (rgb rgbOp) generateFilter() gift.Filter {
	return rgb.generator(rgb.r, rgb.g, rgb.b)
}

func (r rgbOp) getName() string {
	return r.name
}

func setMaps() {

	noParamsOp = map[string]any{"sobel": gift.Sobel, "invert": gift.Invert, "grayscale": gift.Grayscale, "transpose": gift.Transpose, "transverse": gift.Transverse, "flip-horizontal": gift.FlipHorizontal, "flip-vertical": gift.FlipVertical}

	oneFloatOp = map[string]any{"sepia": gift.Sepia, "brightness": gift.Brightness, "contrast": gift.Contrast, "gamma": gift.Gamma, "hue": gift.Hue,
		"gaussian": gift.GaussianBlur}

	rgb = map[string]any{"color-balance": gift.ColorBalance}

}

func NewEffect(name string, params ...float32) (gift.Filter, error) {
	switch len(params) {
	case 0:
		var n noParams
		n.name = name
		n.generator = noParamsOp[name].(func() gift.Filter)
		return n.generateFilter(), nil

	case 1:
		var o oneFloat
		o.name = name
		o.param = params[0]
		o.generator = oneFloatOp[name].(func(float32) gift.Filter)
		return o.generateFilter(), nil

	case 3:
		var r rgbOp
		r.name = name
		r.r = params[0]
		r.g = params[1]
		r.b = params[2]
		r.generator = rgb[name].(func(float32, float32, float32) gift.Filter)
		return r.generateFilter(), nil

	default:
		return nil, errors.New("invalid arguments size")

	}

}

func init() {
	setMaps()
}
