package filter

import (
	"fmt"

	"github.com/disintegration/gift"
)

var (
	effects map[string]effect
)

type Filter struct {
	Name   string `json:"name"`
	Params []any  `json:"params"`
}

func (f *Filter) Generate() (gift.Filter, error) {
	f.ConvertFloat()
	e, ok := effects[f.Name]
	if !ok {
		return nil, fmt.Errorf("invalid filter name %s", f.Name)
	}

	if err := e.setParams(f.Params...); err != nil {
		return nil, err
	}

	return e.generateFilter(), nil

}

func (f *Filter) ConvertFloat() {
	for idx, val := range f.Params {
		s, ok := val.([]float64)
		if ok {
			f.Params[idx] = sliceConvert(s)

		}
		v, ok := val.(float64)
		if ok {
			f.Params[idx] = float32(v)
		}

	}

}

func sliceConvert(s []float64) []float32 {
	var converted []float32
	for _, v := range s {
		converted = append(converted, float32(v))
	}
	return converted
}

type effect interface {
	generateFilter() gift.Filter
	getName() string
	setParams(params ...any) error
}

type noParams struct {
	name      string
	generator func() gift.Filter
}

func (n *noParams) generateFilter() gift.Filter {
	return n.generator()
}

func (n *noParams) getName() string {
	return n.name
}

func (n *noParams) setParams(params ...any) error {
	if len(params) > 0 {
		return fmt.Errorf("%s take no arguments", n.getName())

	}
	return nil

}

type oneFloat struct {
	name      string
	param     float32
	generator func(float32) gift.Filter
}

func (o *oneFloat) generateFilter() gift.Filter {
	return o.generator(o.param)
}

func (o *oneFloat) getName() string {
	return o.name
}

func (o *oneFloat) setParams(params ...any) error {
	if len(params) != 1 {
		return fmt.Errorf("%s require one argument", o.getName())

	}

	param, ok := params[0].(float32)
	if !ok {
		return fmt.Errorf("argument for %s must be a float 32", o.getName())
	}
	o.param = param
	return nil

}

type rgbOp struct {
	name      string
	r, g, b   float32
	generator func(float32, float32, float32) gift.Filter
}

func (rgb *rgbOp) generateFilter() gift.Filter {
	return rgb.generator(rgb.r, rgb.g, rgb.b)
}

func (r *rgbOp) getName() string {
	return r.name
}

func (r *rgbOp) setParams(params ...any) error {
	if len(params) != 3 {
		return fmt.Errorf("%s require three arguments", r.getName())

	}
	var rgb = make([]float32, 3)
	for _, p := range params {
		arg, ok := p.(float32)
		if !ok {
			return fmt.Errorf("argument for %s must be a float 32", r.getName())
		} else {
			rgb = append(rgb, arg)
		}

	}
	r.r, r.g, r.b = rgb[0], rgb[1], rgb[2]
	return nil

}

type convolution struct {
	kernel    []float32
	normalize bool
	alpha     bool
	abs       bool
	delta     float32
	generator func([]float32, bool, bool, bool, float32) gift.Filter
}

func (c *convolution) generateFilter() gift.Filter {
	return c.generator(c.kernel, c.normalize, c.alpha, c.abs, c.delta)
}

func (c *convolution) getName() string {
	return "convolution"
}

func (c *convolution) setParams(params ...any) error {
	if len(params) != 5 {
		return fmt.Errorf("%s require 5 arguments", c.getName())
	}
	s, ok := params[0].([]float32)
	if !ok {
		return fmt.Errorf("%s require a float 32 slice for kernel", c.getName())
	}
	c.kernel = s
	var args []bool
	for idx, arg := range params[1:4] {
		v, ok := arg.(bool)
		if !ok {
			return fmt.Errorf("%s require a bool for argument %d", c.getName(), idx+1)
		}
		args = append(args, v)
	}
	c.normalize, c.alpha, c.abs = args[0], args[1], args[2]

	v, ok := params[4].(float32)
	if !ok {
		return fmt.Errorf("%s require float 32 for argument 5", c.getName())
	}
	c.delta = v
	return nil
}

func setMaps() {

	effects = map[string]effect{
		"sobel":           &noParams{name: "sobel", generator: gift.Sobel},
		"invert":          &noParams{name: "invert", generator: gift.Invert},
		"grayscale":       &noParams{name: "grayscale", generator: gift.Grayscale},
		"transpose":       &noParams{name: "transpose", generator: gift.Transpose},
		"transverse":      &noParams{name: "transverse", generator: gift.Transverse},
		"flip-horizontal": &noParams{name: "flip-horizontal", generator: gift.FlipHorizontal},
		"flip-vertical":   &noParams{name: "fip-vertical", generator: gift.FlipVertical},
		"sepia":           &oneFloat{name: "sepia", generator: gift.Sepia},
		"brightness":      &oneFloat{name: "brightness", generator: gift.Brightness},
		"contrast":        &oneFloat{name: "contrast", generator: gift.Contrast},
		"gamma":           &oneFloat{name: "gamma", generator: gift.Gamma},
		"hue":             &oneFloat{name: "hue", generator: gift.Hue},
		"gaussian":        &oneFloat{name: "gaussian-blur", generator: gift.GaussianBlur},
		"gaussian-blur":   &oneFloat{name: "gaussian-blur", generator: gift.GaussianBlur},
		"color-balance":   &rgbOp{name: "color-balance", generator: gift.ColorBalance},
		"convolution":     &convolution{generator: gift.Convolution},
	}

}

func init() {
	setMaps()
}
