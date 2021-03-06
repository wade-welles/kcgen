package kcgen

// Layer represents the layer on which a graphical element resides.
type Layer int

// Layers.
const (
	LayerFrontFab Layer = iota
	LayerBackFab
	LayerFrontSilkscreen
	LayerBackSilkscreen
	LayerEdgeCuts
	LayerAllCopper
	LayerAllMask
	LayerFrontCopper
	LayerFrontMask
)

// Strictname returns the string representing the layer used in the Kicad 4 (and probably later) formats.
func (l Layer) Strictname() string {
	switch l {
	case LayerFrontFab:
		return "F.Fab"
	case LayerBackFab:
		return "B.Fab"
	case LayerFrontSilkscreen:
		return "F.SilkS"
	case LayerBackSilkscreen:
		return "B.SilkS"
	case LayerEdgeCuts:
		return "Edge.Cuts"
	case LayerAllCopper:
		return "*.Cu"
	case LayerAllMask:
		return "*.Mask"
	case LayerFrontCopper:
		return "F.Cu"
	case LayerFrontMask:
		return "F.Mask"
	}
	panic("invalid layer")
}

// String returns a human representation of the layer
func (l Layer) String() string {
	switch l {
	case LayerFrontFab:
		return "Front Fabrication"
	case LayerBackFab:
		return "Back Fabrication"
	case LayerFrontSilkscreen:
		return "Front Silkscreen"
	case LayerBackSilkscreen:
		return "Back Silkscreen"
	case LayerEdgeCuts:
		return "Board Outline"
	case LayerAllCopper:
		return "All Copper"
	case LayerAllMask:
		return "All Mask"
	case LayerFrontCopper:
		return "Front Copper"
	case LayerFrontMask:
		return "Front Mask"
	}
	return "?"
}
