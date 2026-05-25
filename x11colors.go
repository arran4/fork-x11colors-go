package x11colors

import "image/color"

// Name defines type for color name
type Name string

// X11Color defines single color
type X11Color struct {
	// Name is the color name as defined in <X11root>/lib/X11/rgb.txt (https://gitlab.freedesktop.org/xorg/app/rgb/-/raw/master/rgb.txt)
	Name Name
	// RGBA is the color RGB with A component always set to 0xFF (255)
	RGBA color.RGBA
	// CaptionBlack defines the text color that can be used on the current color background:
	// true for black and false for white - as it is in
	// https://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D1%86%D0%B2%D0%B5%D1%82%D0%BE%D0%B2_%D0%B2_X11
	CaptionBlack bool
}

var (
	AliceBlue            = X11Color{"Alice Blue", color.RGBA{0xF0, 0xF8, 0xFF, 0xFF}, true}
	Aliceblue            = X11Color{"AliceBlue", color.RGBA{0xF0, 0xF8, 0xFF, 0xFF}, true}
	AntiqueWhite         = X11Color{"Antique White", color.RGBA{0xFA, 0xEB, 0xD7, 0xFF}, true}
	Antiquewhite         = X11Color{"AntiqueWhite", color.RGBA{0xFA, 0xEB, 0xD7, 0xFF}, true}
	Antiquewhite1        = X11Color{"AntiqueWhite1", color.RGBA{0xFF, 0xEF, 0xDB, 0xFF}, true}
	Antiquewhite2        = X11Color{"AntiqueWhite2", color.RGBA{0xEE, 0xDF, 0xCC, 0xFF}, true}
	Antiquewhite3        = X11Color{"AntiqueWhite3", color.RGBA{0xCD, 0xC0, 0xB0, 0xFF}, true}
	Antiquewhite4        = X11Color{"AntiqueWhite4", color.RGBA{0x8B, 0x83, 0x78, 0xFF}, true}
	Aqua                 = X11Color{"Aqua", color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, false}
	Aquamarine           = X11Color{"Aquamarine", color.RGBA{0x7F, 0xFF, 0xD4, 0xFF}, false}
	Aquamarine1          = X11Color{"aquamarine1", color.RGBA{0x7F, 0xFF, 0xD4, 0xFF}, true}
	Aquamarine2          = X11Color{"aquamarine2", color.RGBA{0x76, 0xEE, 0xC6, 0xFF}, true}
	Aquamarine3          = X11Color{"aquamarine3", color.RGBA{0x66, 0xCD, 0xAA, 0xFF}, true}
	Aquamarine4          = X11Color{"aquamarine4", color.RGBA{0x45, 0x8B, 0x74, 0xFF}, true}
	Azure                = X11Color{"Azure", color.RGBA{0xF0, 0xFF, 0xFF, 0xFF}, true}
	Azure1               = X11Color{"azure1", color.RGBA{0xF0, 0xFF, 0xFF, 0xFF}, true}
	Azure2               = X11Color{"azure2", color.RGBA{0xE0, 0xEE, 0xEE, 0xFF}, true}
	Azure3               = X11Color{"azure3", color.RGBA{0xC1, 0xCD, 0xCD, 0xFF}, true}
	Azure4               = X11Color{"azure4", color.RGBA{0x83, 0x8B, 0x8B, 0xFF}, true}
	Beige                = X11Color{"Beige", color.RGBA{0xF5, 0xF5, 0xDC, 0xFF}, true}
	Bisque               = X11Color{"Bisque", color.RGBA{0xFF, 0xE4, 0xC4, 0xFF}, false}
	Bisque1              = X11Color{"bisque1", color.RGBA{0xFF, 0xE4, 0xC4, 0xFF}, true}
	Bisque2              = X11Color{"bisque2", color.RGBA{0xEE, 0xD5, 0xB7, 0xFF}, true}
	Bisque3              = X11Color{"bisque3", color.RGBA{0xCD, 0xB7, 0x9E, 0xFF}, true}
	Bisque4              = X11Color{"bisque4", color.RGBA{0x8B, 0x7D, 0x6B, 0xFF}, true}
	Black                = X11Color{"Black", color.RGBA{0x00, 0x00, 0x00, 0xFF}, false}
	BlanchedAlmond       = X11Color{"Blanched Almond", color.RGBA{0xFF, 0xEB, 0xCD, 0xFF}, true}
	Blanchedalmond       = X11Color{"BlanchedAlmond", color.RGBA{0xFF, 0xEB, 0xCD, 0xFF}, true}
	Blue                 = X11Color{"Blue", color.RGBA{0x00, 0x00, 0xFF, 0xFF}, false}
	Blue1                = X11Color{"blue1", color.RGBA{0x00, 0x00, 0xFF, 0xFF}, false}
	Blue2                = X11Color{"blue2", color.RGBA{0x00, 0x00, 0xEE, 0xFF}, false}
	Blue3                = X11Color{"blue3", color.RGBA{0x00, 0x00, 0xCD, 0xFF}, false}
	Blue4                = X11Color{"blue4", color.RGBA{0x00, 0x00, 0x8B, 0xFF}, false}
	BlueViolet           = X11Color{"Blue Violet", color.RGBA{0x8A, 0x2B, 0xE2, 0xFF}, false}
	Blueviolet           = X11Color{"BlueViolet", color.RGBA{0x8A, 0x2B, 0xE2, 0xFF}, false}
	Brown                = X11Color{"Brown", color.RGBA{0xA5, 0x2A, 0x2A, 0xFF}, false}
	Brown1               = X11Color{"brown1", color.RGBA{0xFF, 0x40, 0x40, 0xFF}, true}
	Brown2               = X11Color{"brown2", color.RGBA{0xEE, 0x3B, 0x3B, 0xFF}, true}
	Brown3               = X11Color{"brown3", color.RGBA{0xCD, 0x33, 0x33, 0xFF}, false}
	Brown4               = X11Color{"brown4", color.RGBA{0x8B, 0x23, 0x23, 0xFF}, false}
	Burlywood            = X11Color{"Burlywood", color.RGBA{0xDE, 0xB8, 0x87, 0xFF}, false}
	Burlywood1           = X11Color{"burlywood1", color.RGBA{0xFF, 0xD3, 0x9B, 0xFF}, true}
	Burlywood2           = X11Color{"burlywood2", color.RGBA{0xEE, 0xC5, 0x91, 0xFF}, true}
	Burlywood3           = X11Color{"burlywood3", color.RGBA{0xCD, 0xAA, 0x7D, 0xFF}, true}
	Burlywood4           = X11Color{"burlywood4", color.RGBA{0x8B, 0x73, 0x55, 0xFF}, true}
	CadetBlue            = X11Color{"Cadet Blue", color.RGBA{0x5F, 0x9E, 0xA0, 0xFF}, false}
	Cadetblue            = X11Color{"CadetBlue", color.RGBA{0x5F, 0x9E, 0xA0, 0xFF}, true}
	Cadetblue1           = X11Color{"CadetBlue1", color.RGBA{0x98, 0xF5, 0xFF, 0xFF}, true}
	Cadetblue2           = X11Color{"CadetBlue2", color.RGBA{0x8E, 0xE5, 0xEE, 0xFF}, true}
	Cadetblue3           = X11Color{"CadetBlue3", color.RGBA{0x7A, 0xC5, 0xCD, 0xFF}, true}
	Cadetblue4           = X11Color{"CadetBlue4", color.RGBA{0x53, 0x86, 0x8B, 0xFF}, true}
	Chartreuse           = X11Color{"Chartreuse", color.RGBA{0x7F, 0xFF, 0x00, 0xFF}, false}
	Chartreuse1          = X11Color{"chartreuse1", color.RGBA{0x7F, 0xFF, 0x00, 0xFF}, true}
	Chartreuse2          = X11Color{"chartreuse2", color.RGBA{0x76, 0xEE, 0x00, 0xFF}, true}
	Chartreuse3          = X11Color{"chartreuse3", color.RGBA{0x66, 0xCD, 0x00, 0xFF}, true}
	Chartreuse4          = X11Color{"chartreuse4", color.RGBA{0x45, 0x8B, 0x00, 0xFF}, true}
	Chocolate            = X11Color{"Chocolate", color.RGBA{0xD2, 0x69, 0x1E, 0xFF}, false}
	Chocolate1           = X11Color{"chocolate1", color.RGBA{0xFF, 0x7F, 0x24, 0xFF}, true}
	Chocolate2           = X11Color{"chocolate2", color.RGBA{0xEE, 0x76, 0x21, 0xFF}, true}
	Chocolate3           = X11Color{"chocolate3", color.RGBA{0xCD, 0x66, 0x1D, 0xFF}, true}
	Chocolate4           = X11Color{"chocolate4", color.RGBA{0x8B, 0x45, 0x13, 0xFF}, false}
	Coral                = X11Color{"Coral", color.RGBA{0xFF, 0x7F, 0x50, 0xFF}, false}
	Coral1               = X11Color{"coral1", color.RGBA{0xFF, 0x72, 0x56, 0xFF}, true}
	Coral2               = X11Color{"coral2", color.RGBA{0xEE, 0x6A, 0x50, 0xFF}, true}
	Coral3               = X11Color{"coral3", color.RGBA{0xCD, 0x5B, 0x45, 0xFF}, true}
	Coral4               = X11Color{"coral4", color.RGBA{0x8B, 0x3E, 0x2F, 0xFF}, false}
	Cornflower           = X11Color{"Cornflower", color.RGBA{0x64, 0x95, 0xED, 0xFF}, false}
	CornflowerBlue       = X11Color{"cornflower blue", color.RGBA{0x64, 0x95, 0xED, 0xFF}, true}
	Cornflowerblue       = X11Color{"CornflowerBlue", color.RGBA{0x64, 0x95, 0xED, 0xFF}, true}
	Cornsilk             = X11Color{"Cornsilk", color.RGBA{0xFF, 0xF8, 0xDC, 0xFF}, true}
	Cornsilk1            = X11Color{"cornsilk1", color.RGBA{0xFF, 0xF8, 0xDC, 0xFF}, true}
	Cornsilk2            = X11Color{"cornsilk2", color.RGBA{0xEE, 0xE8, 0xCD, 0xFF}, true}
	Cornsilk3            = X11Color{"cornsilk3", color.RGBA{0xCD, 0xC8, 0xB1, 0xFF}, true}
	Cornsilk4            = X11Color{"cornsilk4", color.RGBA{0x8B, 0x88, 0x78, 0xFF}, true}
	Crimson              = X11Color{"Crimson", color.RGBA{0xDC, 0x14, 0x3C, 0xFF}, false}
	Cyan                 = X11Color{"Cyan", color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, false}
	Cyan1                = X11Color{"cyan1", color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, true}
	Cyan2                = X11Color{"cyan2", color.RGBA{0x00, 0xEE, 0xEE, 0xFF}, true}
	Cyan3                = X11Color{"cyan3", color.RGBA{0x00, 0xCD, 0xCD, 0xFF}, true}
	Cyan4                = X11Color{"cyan4", color.RGBA{0x00, 0x8B, 0x8B, 0xFF}, true}
	DarkBlue             = X11Color{"Dark Blue", color.RGBA{0x00, 0x00, 0x8B, 0xFF}, false}
	DarkCyan             = X11Color{"Dark Cyan", color.RGBA{0x00, 0x8B, 0x8B, 0xFF}, false}
	DarkGoldenrod        = X11Color{"Dark Goldenrod", color.RGBA{0xB8, 0x86, 0x0B, 0xFF}, false}
	DarkGray             = X11Color{"Dark Gray", color.RGBA{0xA9, 0xA9, 0xA9, 0xFF}, false}
	DarkGreen            = X11Color{"Dark Green", color.RGBA{0x00, 0x64, 0x00, 0xFF}, false}
	DarkGrey             = X11Color{"dark grey", color.RGBA{0xA9, 0xA9, 0xA9, 0xFF}, true}
	DarkKhaki            = X11Color{"Dark Khaki", color.RGBA{0xBD, 0xB7, 0x6B, 0xFF}, false}
	DarkMagenta          = X11Color{"Dark Magenta", color.RGBA{0x8B, 0x00, 0x8B, 0xFF}, false}
	DarkOliveGreen       = X11Color{"Dark Olive Green", color.RGBA{0x55, 0x6B, 0x2F, 0xFF}, false}
	DarkOrange           = X11Color{"Dark Orange", color.RGBA{0xFF, 0x8C, 0x00, 0xFF}, false}
	DarkOrchid           = X11Color{"Dark Orchid", color.RGBA{0x99, 0x32, 0xCC, 0xFF}, false}
	DarkRed              = X11Color{"Dark Red", color.RGBA{0x8B, 0x00, 0x00, 0xFF}, false}
	DarkSalmon           = X11Color{"Dark Salmon", color.RGBA{0xE9, 0x96, 0x7A, 0xFF}, false}
	DarkSeaGreen         = X11Color{"Dark Sea Green", color.RGBA{0x8F, 0xBC, 0x8F, 0xFF}, false}
	DarkSlateBlue        = X11Color{"Dark Slate Blue", color.RGBA{0x48, 0x3D, 0x8B, 0xFF}, false}
	DarkSlateGray        = X11Color{"Dark Slate Gray", color.RGBA{0x2F, 0x4F, 0x4F, 0xFF}, false}
	DarkSlateGrey        = X11Color{"dark slate grey", color.RGBA{0x2F, 0x4F, 0x4F, 0xFF}, false}
	DarkTurquoise        = X11Color{"Dark Turquoise", color.RGBA{0x00, 0xCE, 0xD1, 0xFF}, false}
	DarkViolet           = X11Color{"Dark Violet", color.RGBA{0x94, 0x00, 0xD3, 0xFF}, false}
	Darkblue             = X11Color{"DarkBlue", color.RGBA{0x00, 0x00, 0x8B, 0xFF}, false}
	Darkcyan             = X11Color{"DarkCyan", color.RGBA{0x00, 0x8B, 0x8B, 0xFF}, true}
	Darkgoldenrod        = X11Color{"DarkGoldenrod", color.RGBA{0xB8, 0x86, 0x0B, 0xFF}, true}
	Darkgoldenrod1       = X11Color{"DarkGoldenrod1", color.RGBA{0xFF, 0xB9, 0x0F, 0xFF}, true}
	Darkgoldenrod2       = X11Color{"DarkGoldenrod2", color.RGBA{0xEE, 0xAD, 0x0E, 0xFF}, true}
	Darkgoldenrod3       = X11Color{"DarkGoldenrod3", color.RGBA{0xCD, 0x95, 0x0C, 0xFF}, true}
	Darkgoldenrod4       = X11Color{"DarkGoldenrod4", color.RGBA{0x8B, 0x65, 0x08, 0xFF}, false}
	Darkgray             = X11Color{"DarkGray", color.RGBA{0xA9, 0xA9, 0xA9, 0xFF}, true}
	Darkgreen            = X11Color{"DarkGreen", color.RGBA{0x00, 0x64, 0x00, 0xFF}, false}
	Darkgrey             = X11Color{"DarkGrey", color.RGBA{0xA9, 0xA9, 0xA9, 0xFF}, true}
	Darkkhaki            = X11Color{"DarkKhaki", color.RGBA{0xBD, 0xB7, 0x6B, 0xFF}, true}
	Darkmagenta          = X11Color{"DarkMagenta", color.RGBA{0x8B, 0x00, 0x8B, 0xFF}, false}
	Darkolivegreen       = X11Color{"DarkOliveGreen", color.RGBA{0x55, 0x6B, 0x2F, 0xFF}, false}
	Darkolivegreen1      = X11Color{"DarkOliveGreen1", color.RGBA{0xCA, 0xFF, 0x70, 0xFF}, true}
	Darkolivegreen2      = X11Color{"DarkOliveGreen2", color.RGBA{0xBC, 0xEE, 0x68, 0xFF}, true}
	Darkolivegreen3      = X11Color{"DarkOliveGreen3", color.RGBA{0xA2, 0xCD, 0x5A, 0xFF}, true}
	Darkolivegreen4      = X11Color{"DarkOliveGreen4", color.RGBA{0x6E, 0x8B, 0x3D, 0xFF}, true}
	Darkorange           = X11Color{"DarkOrange", color.RGBA{0xFF, 0x8C, 0x00, 0xFF}, true}
	Darkorange1          = X11Color{"DarkOrange1", color.RGBA{0xFF, 0x7F, 0x00, 0xFF}, true}
	Darkorange2          = X11Color{"DarkOrange2", color.RGBA{0xEE, 0x76, 0x00, 0xFF}, true}
	Darkorange3          = X11Color{"DarkOrange3", color.RGBA{0xCD, 0x66, 0x00, 0xFF}, true}
	Darkorange4          = X11Color{"DarkOrange4", color.RGBA{0x8B, 0x45, 0x00, 0xFF}, false}
	Darkorchid           = X11Color{"DarkOrchid", color.RGBA{0x99, 0x32, 0xCC, 0xFF}, false}
	Darkorchid1          = X11Color{"DarkOrchid1", color.RGBA{0xBF, 0x3E, 0xFF, 0xFF}, true}
	Darkorchid2          = X11Color{"DarkOrchid2", color.RGBA{0xB2, 0x3A, 0xEE, 0xFF}, true}
	Darkorchid3          = X11Color{"DarkOrchid3", color.RGBA{0x9A, 0x32, 0xCD, 0xFF}, false}
	Darkorchid4          = X11Color{"DarkOrchid4", color.RGBA{0x68, 0x22, 0x8B, 0xFF}, false}
	Darkred              = X11Color{"DarkRed", color.RGBA{0x8B, 0x00, 0x00, 0xFF}, false}
	Darksalmon           = X11Color{"DarkSalmon", color.RGBA{0xE9, 0x96, 0x7A, 0xFF}, true}
	Darkseagreen         = X11Color{"DarkSeaGreen", color.RGBA{0x8F, 0xBC, 0x8F, 0xFF}, true}
	Darkseagreen1        = X11Color{"DarkSeaGreen1", color.RGBA{0xC1, 0xFF, 0xC1, 0xFF}, true}
	Darkseagreen2        = X11Color{"DarkSeaGreen2", color.RGBA{0xB4, 0xEE, 0xB4, 0xFF}, true}
	Darkseagreen3        = X11Color{"DarkSeaGreen3", color.RGBA{0x9B, 0xCD, 0x9B, 0xFF}, true}
	Darkseagreen4        = X11Color{"DarkSeaGreen4", color.RGBA{0x69, 0x8B, 0x69, 0xFF}, true}
	Darkslateblue        = X11Color{"DarkSlateBlue", color.RGBA{0x48, 0x3D, 0x8B, 0xFF}, false}
	Darkslategray        = X11Color{"DarkSlateGray", color.RGBA{0x2F, 0x4F, 0x4F, 0xFF}, false}
	Darkslategray1       = X11Color{"DarkSlateGray1", color.RGBA{0x97, 0xFF, 0xFF, 0xFF}, true}
	Darkslategray2       = X11Color{"DarkSlateGray2", color.RGBA{0x8D, 0xEE, 0xEE, 0xFF}, true}
	Darkslategray3       = X11Color{"DarkSlateGray3", color.RGBA{0x79, 0xCD, 0xCD, 0xFF}, true}
	Darkslategray4       = X11Color{"DarkSlateGray4", color.RGBA{0x52, 0x8B, 0x8B, 0xFF}, true}
	Darkslategrey        = X11Color{"DarkSlateGrey", color.RGBA{0x2F, 0x4F, 0x4F, 0xFF}, false}
	Darkturquoise        = X11Color{"DarkTurquoise", color.RGBA{0x00, 0xCE, 0xD1, 0xFF}, true}
	Darkviolet           = X11Color{"DarkViolet", color.RGBA{0x94, 0x00, 0xD3, 0xFF}, false}
	DeepPink             = X11Color{"Deep Pink", color.RGBA{0xFF, 0x14, 0x93, 0xFF}, false}
	DeepSkyBlue          = X11Color{"Deep Sky Blue", color.RGBA{0x00, 0xBF, 0xFF, 0xFF}, false}
	Deeppink             = X11Color{"DeepPink", color.RGBA{0xFF, 0x14, 0x93, 0xFF}, true}
	Deeppink1            = X11Color{"DeepPink1", color.RGBA{0xFF, 0x14, 0x93, 0xFF}, true}
	Deeppink2            = X11Color{"DeepPink2", color.RGBA{0xEE, 0x12, 0x89, 0xFF}, true}
	Deeppink3            = X11Color{"DeepPink3", color.RGBA{0xCD, 0x10, 0x76, 0xFF}, false}
	Deeppink4            = X11Color{"DeepPink4", color.RGBA{0x8B, 0x0A, 0x50, 0xFF}, false}
	Deepskyblue          = X11Color{"DeepSkyBlue", color.RGBA{0x00, 0xBF, 0xFF, 0xFF}, true}
	Deepskyblue1         = X11Color{"DeepSkyBlue1", color.RGBA{0x00, 0xBF, 0xFF, 0xFF}, true}
	Deepskyblue2         = X11Color{"DeepSkyBlue2", color.RGBA{0x00, 0xB2, 0xEE, 0xFF}, true}
	Deepskyblue3         = X11Color{"DeepSkyBlue3", color.RGBA{0x00, 0x9A, 0xCD, 0xFF}, true}
	Deepskyblue4         = X11Color{"DeepSkyBlue4", color.RGBA{0x00, 0x68, 0x8B, 0xFF}, false}
	DimGray              = X11Color{"Dim Gray", color.RGBA{0x69, 0x69, 0x69, 0xFF}, false}
	DimGrey              = X11Color{"dim grey", color.RGBA{0x69, 0x69, 0x69, 0xFF}, false}
	Dimgray              = X11Color{"DimGray", color.RGBA{0x69, 0x69, 0x69, 0xFF}, false}
	Dimgrey              = X11Color{"DimGrey", color.RGBA{0x69, 0x69, 0x69, 0xFF}, false}
	DodgerBlue           = X11Color{"Dodger Blue", color.RGBA{0x1E, 0x90, 0xFF, 0xFF}, false}
	Dodgerblue           = X11Color{"DodgerBlue", color.RGBA{0x1E, 0x90, 0xFF, 0xFF}, true}
	Dodgerblue1          = X11Color{"DodgerBlue1", color.RGBA{0x1E, 0x90, 0xFF, 0xFF}, true}
	Dodgerblue2          = X11Color{"DodgerBlue2", color.RGBA{0x1C, 0x86, 0xEE, 0xFF}, true}
	Dodgerblue3          = X11Color{"DodgerBlue3", color.RGBA{0x18, 0x74, 0xCD, 0xFF}, false}
	Dodgerblue4          = X11Color{"DodgerBlue4", color.RGBA{0x10, 0x4E, 0x8B, 0xFF}, false}
	Firebrick            = X11Color{"Firebrick", color.RGBA{0xB2, 0x22, 0x22, 0xFF}, false}
	Firebrick1           = X11Color{"firebrick1", color.RGBA{0xFF, 0x30, 0x30, 0xFF}, true}
	Firebrick2           = X11Color{"firebrick2", color.RGBA{0xEE, 0x2C, 0x2C, 0xFF}, true}
	Firebrick3           = X11Color{"firebrick3", color.RGBA{0xCD, 0x26, 0x26, 0xFF}, false}
	Firebrick4           = X11Color{"firebrick4", color.RGBA{0x8B, 0x1A, 0x1A, 0xFF}, false}
	FloralWhite          = X11Color{"Floral White", color.RGBA{0xFF, 0xFA, 0xF0, 0xFF}, true}
	Floralwhite          = X11Color{"FloralWhite", color.RGBA{0xFF, 0xFA, 0xF0, 0xFF}, true}
	ForestGreen          = X11Color{"Forest Green", color.RGBA{0x22, 0x8B, 0x22, 0xFF}, false}
	Forestgreen          = X11Color{"ForestGreen", color.RGBA{0x22, 0x8B, 0x22, 0xFF}, true}
	Fuchsia              = X11Color{"Fuchsia", color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, false}
	Gainsboro            = X11Color{"Gainsboro", color.RGBA{0xDC, 0xDC, 0xDC, 0xFF}, true}
	GhostWhite           = X11Color{"Ghost White", color.RGBA{0xF8, 0xF8, 0xFF, 0xFF}, true}
	Ghostwhite           = X11Color{"GhostWhite", color.RGBA{0xF8, 0xF8, 0xFF, 0xFF}, true}
	Gold                 = X11Color{"Gold", color.RGBA{0xFF, 0xD7, 0x00, 0xFF}, false}
	Gold1                = X11Color{"gold1", color.RGBA{0xFF, 0xD7, 0x00, 0xFF}, true}
	Gold2                = X11Color{"gold2", color.RGBA{0xEE, 0xC9, 0x00, 0xFF}, true}
	Gold3                = X11Color{"gold3", color.RGBA{0xCD, 0xAD, 0x00, 0xFF}, true}
	Gold4                = X11Color{"gold4", color.RGBA{0x8B, 0x75, 0x00, 0xFF}, true}
	Goldenrod            = X11Color{"Goldenrod", color.RGBA{0xDA, 0xA5, 0x20, 0xFF}, false}
	Goldenrod1           = X11Color{"goldenrod1", color.RGBA{0xFF, 0xC1, 0x25, 0xFF}, true}
	Goldenrod2           = X11Color{"goldenrod2", color.RGBA{0xEE, 0xB4, 0x22, 0xFF}, true}
	Goldenrod3           = X11Color{"goldenrod3", color.RGBA{0xCD, 0x9B, 0x1D, 0xFF}, true}
	Goldenrod4           = X11Color{"goldenrod4", color.RGBA{0x8B, 0x69, 0x14, 0xFF}, false}
	Gray                 = X11Color{"gray", color.RGBA{0xBE, 0xBE, 0xBE, 0xFF}, true}
	Gray0                = X11Color{"gray0", color.RGBA{0x00, 0x00, 0x00, 0xFF}, false}
	Gray1                = X11Color{"gray1", color.RGBA{0x03, 0x03, 0x03, 0xFF}, false}
	Gray10               = X11Color{"gray10", color.RGBA{0x1A, 0x1A, 0x1A, 0xFF}, false}
	Gray100              = X11Color{"gray100", color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, true}
	Gray11               = X11Color{"gray11", color.RGBA{0x1C, 0x1C, 0x1C, 0xFF}, false}
	Gray12               = X11Color{"gray12", color.RGBA{0x1F, 0x1F, 0x1F, 0xFF}, false}
	Gray13               = X11Color{"gray13", color.RGBA{0x21, 0x21, 0x21, 0xFF}, false}
	Gray14               = X11Color{"gray14", color.RGBA{0x24, 0x24, 0x24, 0xFF}, false}
	Gray15               = X11Color{"gray15", color.RGBA{0x26, 0x26, 0x26, 0xFF}, false}
	Gray16               = X11Color{"gray16", color.RGBA{0x29, 0x29, 0x29, 0xFF}, false}
	Gray17               = X11Color{"gray17", color.RGBA{0x2B, 0x2B, 0x2B, 0xFF}, false}
	Gray18               = X11Color{"gray18", color.RGBA{0x2E, 0x2E, 0x2E, 0xFF}, false}
	Gray19               = X11Color{"gray19", color.RGBA{0x30, 0x30, 0x30, 0xFF}, false}
	Gray2                = X11Color{"gray2", color.RGBA{0x05, 0x05, 0x05, 0xFF}, false}
	Gray20               = X11Color{"gray20", color.RGBA{0x33, 0x33, 0x33, 0xFF}, false}
	Gray21               = X11Color{"gray21", color.RGBA{0x36, 0x36, 0x36, 0xFF}, false}
	Gray22               = X11Color{"gray22", color.RGBA{0x38, 0x38, 0x38, 0xFF}, false}
	Gray23               = X11Color{"gray23", color.RGBA{0x3B, 0x3B, 0x3B, 0xFF}, false}
	Gray24               = X11Color{"gray24", color.RGBA{0x3D, 0x3D, 0x3D, 0xFF}, false}
	Gray25               = X11Color{"gray25", color.RGBA{0x40, 0x40, 0x40, 0xFF}, false}
	Gray26               = X11Color{"gray26", color.RGBA{0x42, 0x42, 0x42, 0xFF}, false}
	Gray27               = X11Color{"gray27", color.RGBA{0x45, 0x45, 0x45, 0xFF}, false}
	Gray28               = X11Color{"gray28", color.RGBA{0x47, 0x47, 0x47, 0xFF}, false}
	Gray29               = X11Color{"gray29", color.RGBA{0x4A, 0x4A, 0x4A, 0xFF}, false}
	Gray3                = X11Color{"gray3", color.RGBA{0x08, 0x08, 0x08, 0xFF}, false}
	Gray30               = X11Color{"gray30", color.RGBA{0x4D, 0x4D, 0x4D, 0xFF}, false}
	Gray31               = X11Color{"gray31", color.RGBA{0x4F, 0x4F, 0x4F, 0xFF}, false}
	Gray32               = X11Color{"gray32", color.RGBA{0x52, 0x52, 0x52, 0xFF}, false}
	Gray33               = X11Color{"gray33", color.RGBA{0x54, 0x54, 0x54, 0xFF}, false}
	Gray34               = X11Color{"gray34", color.RGBA{0x57, 0x57, 0x57, 0xFF}, false}
	Gray35               = X11Color{"gray35", color.RGBA{0x59, 0x59, 0x59, 0xFF}, false}
	Gray36               = X11Color{"gray36", color.RGBA{0x5C, 0x5C, 0x5C, 0xFF}, false}
	Gray37               = X11Color{"gray37", color.RGBA{0x5E, 0x5E, 0x5E, 0xFF}, false}
	Gray38               = X11Color{"gray38", color.RGBA{0x61, 0x61, 0x61, 0xFF}, false}
	Gray39               = X11Color{"gray39", color.RGBA{0x63, 0x63, 0x63, 0xFF}, false}
	Gray4                = X11Color{"gray4", color.RGBA{0x0A, 0x0A, 0x0A, 0xFF}, false}
	Gray40               = X11Color{"gray40", color.RGBA{0x66, 0x66, 0x66, 0xFF}, false}
	Gray41               = X11Color{"gray41", color.RGBA{0x69, 0x69, 0x69, 0xFF}, false}
	Gray42               = X11Color{"gray42", color.RGBA{0x6B, 0x6B, 0x6B, 0xFF}, false}
	Gray43               = X11Color{"gray43", color.RGBA{0x6E, 0x6E, 0x6E, 0xFF}, false}
	Gray44               = X11Color{"gray44", color.RGBA{0x70, 0x70, 0x70, 0xFF}, false}
	Gray45               = X11Color{"gray45", color.RGBA{0x73, 0x73, 0x73, 0xFF}, false}
	Gray46               = X11Color{"gray46", color.RGBA{0x75, 0x75, 0x75, 0xFF}, false}
	Gray47               = X11Color{"gray47", color.RGBA{0x78, 0x78, 0x78, 0xFF}, true}
	Gray48               = X11Color{"gray48", color.RGBA{0x7A, 0x7A, 0x7A, 0xFF}, true}
	Gray49               = X11Color{"gray49", color.RGBA{0x7D, 0x7D, 0x7D, 0xFF}, true}
	Gray5                = X11Color{"gray5", color.RGBA{0x0D, 0x0D, 0x0D, 0xFF}, false}
	Gray50               = X11Color{"gray50", color.RGBA{0x7F, 0x7F, 0x7F, 0xFF}, true}
	Gray51               = X11Color{"gray51", color.RGBA{0x82, 0x82, 0x82, 0xFF}, true}
	Gray52               = X11Color{"gray52", color.RGBA{0x85, 0x85, 0x85, 0xFF}, true}
	Gray53               = X11Color{"gray53", color.RGBA{0x87, 0x87, 0x87, 0xFF}, true}
	Gray54               = X11Color{"gray54", color.RGBA{0x8A, 0x8A, 0x8A, 0xFF}, true}
	Gray55               = X11Color{"gray55", color.RGBA{0x8C, 0x8C, 0x8C, 0xFF}, true}
	Gray56               = X11Color{"gray56", color.RGBA{0x8F, 0x8F, 0x8F, 0xFF}, true}
	Gray57               = X11Color{"gray57", color.RGBA{0x91, 0x91, 0x91, 0xFF}, true}
	Gray58               = X11Color{"gray58", color.RGBA{0x94, 0x94, 0x94, 0xFF}, true}
	Gray59               = X11Color{"gray59", color.RGBA{0x96, 0x96, 0x96, 0xFF}, true}
	Gray6                = X11Color{"gray6", color.RGBA{0x0F, 0x0F, 0x0F, 0xFF}, false}
	Gray60               = X11Color{"gray60", color.RGBA{0x99, 0x99, 0x99, 0xFF}, true}
	Gray61               = X11Color{"gray61", color.RGBA{0x9C, 0x9C, 0x9C, 0xFF}, true}
	Gray62               = X11Color{"gray62", color.RGBA{0x9E, 0x9E, 0x9E, 0xFF}, true}
	Gray63               = X11Color{"gray63", color.RGBA{0xA1, 0xA1, 0xA1, 0xFF}, true}
	Gray64               = X11Color{"gray64", color.RGBA{0xA3, 0xA3, 0xA3, 0xFF}, true}
	Gray65               = X11Color{"gray65", color.RGBA{0xA6, 0xA6, 0xA6, 0xFF}, true}
	Gray66               = X11Color{"gray66", color.RGBA{0xA8, 0xA8, 0xA8, 0xFF}, true}
	Gray67               = X11Color{"gray67", color.RGBA{0xAB, 0xAB, 0xAB, 0xFF}, true}
	Gray68               = X11Color{"gray68", color.RGBA{0xAD, 0xAD, 0xAD, 0xFF}, true}
	Gray69               = X11Color{"gray69", color.RGBA{0xB0, 0xB0, 0xB0, 0xFF}, true}
	Gray7                = X11Color{"gray7", color.RGBA{0x12, 0x12, 0x12, 0xFF}, false}
	Gray70               = X11Color{"gray70", color.RGBA{0xB3, 0xB3, 0xB3, 0xFF}, true}
	Gray71               = X11Color{"gray71", color.RGBA{0xB5, 0xB5, 0xB5, 0xFF}, true}
	Gray72               = X11Color{"gray72", color.RGBA{0xB8, 0xB8, 0xB8, 0xFF}, true}
	Gray73               = X11Color{"gray73", color.RGBA{0xBA, 0xBA, 0xBA, 0xFF}, true}
	Gray74               = X11Color{"gray74", color.RGBA{0xBD, 0xBD, 0xBD, 0xFF}, true}
	Gray75               = X11Color{"gray75", color.RGBA{0xBF, 0xBF, 0xBF, 0xFF}, true}
	Gray76               = X11Color{"gray76", color.RGBA{0xC2, 0xC2, 0xC2, 0xFF}, true}
	Gray77               = X11Color{"gray77", color.RGBA{0xC4, 0xC4, 0xC4, 0xFF}, true}
	Gray78               = X11Color{"gray78", color.RGBA{0xC7, 0xC7, 0xC7, 0xFF}, true}
	Gray79               = X11Color{"gray79", color.RGBA{0xC9, 0xC9, 0xC9, 0xFF}, true}
	Gray8                = X11Color{"gray8", color.RGBA{0x14, 0x14, 0x14, 0xFF}, false}
	Gray80               = X11Color{"gray80", color.RGBA{0xCC, 0xCC, 0xCC, 0xFF}, true}
	Gray81               = X11Color{"gray81", color.RGBA{0xCF, 0xCF, 0xCF, 0xFF}, true}
	Gray82               = X11Color{"gray82", color.RGBA{0xD1, 0xD1, 0xD1, 0xFF}, true}
	Gray83               = X11Color{"gray83", color.RGBA{0xD4, 0xD4, 0xD4, 0xFF}, true}
	Gray84               = X11Color{"gray84", color.RGBA{0xD6, 0xD6, 0xD6, 0xFF}, true}
	Gray85               = X11Color{"gray85", color.RGBA{0xD9, 0xD9, 0xD9, 0xFF}, true}
	Gray86               = X11Color{"gray86", color.RGBA{0xDB, 0xDB, 0xDB, 0xFF}, true}
	Gray87               = X11Color{"gray87", color.RGBA{0xDE, 0xDE, 0xDE, 0xFF}, true}
	Gray88               = X11Color{"gray88", color.RGBA{0xE0, 0xE0, 0xE0, 0xFF}, true}
	Gray89               = X11Color{"gray89", color.RGBA{0xE3, 0xE3, 0xE3, 0xFF}, true}
	Gray9                = X11Color{"gray9", color.RGBA{0x17, 0x17, 0x17, 0xFF}, false}
	Gray90               = X11Color{"gray90", color.RGBA{0xE5, 0xE5, 0xE5, 0xFF}, true}
	Gray91               = X11Color{"gray91", color.RGBA{0xE8, 0xE8, 0xE8, 0xFF}, true}
	Gray92               = X11Color{"gray92", color.RGBA{0xEB, 0xEB, 0xEB, 0xFF}, true}
	Gray93               = X11Color{"gray93", color.RGBA{0xED, 0xED, 0xED, 0xFF}, true}
	Gray94               = X11Color{"gray94", color.RGBA{0xF0, 0xF0, 0xF0, 0xFF}, true}
	Gray95               = X11Color{"gray95", color.RGBA{0xF2, 0xF2, 0xF2, 0xFF}, true}
	Gray96               = X11Color{"gray96", color.RGBA{0xF5, 0xF5, 0xF5, 0xFF}, true}
	Gray97               = X11Color{"gray97", color.RGBA{0xF7, 0xF7, 0xF7, 0xFF}, true}
	Gray98               = X11Color{"gray98", color.RGBA{0xFA, 0xFA, 0xFA, 0xFF}, true}
	Gray99               = X11Color{"gray99", color.RGBA{0xFC, 0xFC, 0xFC, 0xFF}, true}
	GrayW3C              = X11Color{"Gray (W3C)", color.RGBA{0x7F, 0x7F, 0x7F, 0xFF}, false}
	GrayX11              = X11Color{"Gray (X11)", color.RGBA{0xBE, 0xBE, 0xBE, 0xFF}, false}
	Green                = X11Color{"green", color.RGBA{0x00, 0xFF, 0x00, 0xFF}, true}
	Green1               = X11Color{"green1", color.RGBA{0x00, 0xFF, 0x00, 0xFF}, true}
	Green2               = X11Color{"green2", color.RGBA{0x00, 0xEE, 0x00, 0xFF}, true}
	Green3               = X11Color{"green3", color.RGBA{0x00, 0xCD, 0x00, 0xFF}, true}
	Green4               = X11Color{"green4", color.RGBA{0x00, 0x8B, 0x00, 0xFF}, true}
	GreenW3C             = X11Color{"Green (W3C)", color.RGBA{0x00, 0x7F, 0x00, 0xFF}, false}
	GreenX11             = X11Color{"Green (X11)", color.RGBA{0x00, 0xFF, 0x00, 0xFF}, false}
	GreenYellow          = X11Color{"Green Yellow", color.RGBA{0xAD, 0xFF, 0x2F, 0xFF}, false}
	Greenyellow          = X11Color{"GreenYellow", color.RGBA{0xAD, 0xFF, 0x2F, 0xFF}, true}
	Grey                 = X11Color{"grey", color.RGBA{0xBE, 0xBE, 0xBE, 0xFF}, true}
	Grey0                = X11Color{"grey0", color.RGBA{0x00, 0x00, 0x00, 0xFF}, false}
	Grey1                = X11Color{"grey1", color.RGBA{0x03, 0x03, 0x03, 0xFF}, false}
	Grey10               = X11Color{"grey10", color.RGBA{0x1A, 0x1A, 0x1A, 0xFF}, false}
	Grey100              = X11Color{"grey100", color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, true}
	Grey11               = X11Color{"grey11", color.RGBA{0x1C, 0x1C, 0x1C, 0xFF}, false}
	Grey12               = X11Color{"grey12", color.RGBA{0x1F, 0x1F, 0x1F, 0xFF}, false}
	Grey13               = X11Color{"grey13", color.RGBA{0x21, 0x21, 0x21, 0xFF}, false}
	Grey14               = X11Color{"grey14", color.RGBA{0x24, 0x24, 0x24, 0xFF}, false}
	Grey15               = X11Color{"grey15", color.RGBA{0x26, 0x26, 0x26, 0xFF}, false}
	Grey16               = X11Color{"grey16", color.RGBA{0x29, 0x29, 0x29, 0xFF}, false}
	Grey17               = X11Color{"grey17", color.RGBA{0x2B, 0x2B, 0x2B, 0xFF}, false}
	Grey18               = X11Color{"grey18", color.RGBA{0x2E, 0x2E, 0x2E, 0xFF}, false}
	Grey19               = X11Color{"grey19", color.RGBA{0x30, 0x30, 0x30, 0xFF}, false}
	Grey2                = X11Color{"grey2", color.RGBA{0x05, 0x05, 0x05, 0xFF}, false}
	Grey20               = X11Color{"grey20", color.RGBA{0x33, 0x33, 0x33, 0xFF}, false}
	Grey21               = X11Color{"grey21", color.RGBA{0x36, 0x36, 0x36, 0xFF}, false}
	Grey22               = X11Color{"grey22", color.RGBA{0x38, 0x38, 0x38, 0xFF}, false}
	Grey23               = X11Color{"grey23", color.RGBA{0x3B, 0x3B, 0x3B, 0xFF}, false}
	Grey24               = X11Color{"grey24", color.RGBA{0x3D, 0x3D, 0x3D, 0xFF}, false}
	Grey25               = X11Color{"grey25", color.RGBA{0x40, 0x40, 0x40, 0xFF}, false}
	Grey26               = X11Color{"grey26", color.RGBA{0x42, 0x42, 0x42, 0xFF}, false}
	Grey27               = X11Color{"grey27", color.RGBA{0x45, 0x45, 0x45, 0xFF}, false}
	Grey28               = X11Color{"grey28", color.RGBA{0x47, 0x47, 0x47, 0xFF}, false}
	Grey29               = X11Color{"grey29", color.RGBA{0x4A, 0x4A, 0x4A, 0xFF}, false}
	Grey3                = X11Color{"grey3", color.RGBA{0x08, 0x08, 0x08, 0xFF}, false}
	Grey30               = X11Color{"grey30", color.RGBA{0x4D, 0x4D, 0x4D, 0xFF}, false}
	Grey31               = X11Color{"grey31", color.RGBA{0x4F, 0x4F, 0x4F, 0xFF}, false}
	Grey32               = X11Color{"grey32", color.RGBA{0x52, 0x52, 0x52, 0xFF}, false}
	Grey33               = X11Color{"grey33", color.RGBA{0x54, 0x54, 0x54, 0xFF}, false}
	Grey34               = X11Color{"grey34", color.RGBA{0x57, 0x57, 0x57, 0xFF}, false}
	Grey35               = X11Color{"grey35", color.RGBA{0x59, 0x59, 0x59, 0xFF}, false}
	Grey36               = X11Color{"grey36", color.RGBA{0x5C, 0x5C, 0x5C, 0xFF}, false}
	Grey37               = X11Color{"grey37", color.RGBA{0x5E, 0x5E, 0x5E, 0xFF}, false}
	Grey38               = X11Color{"grey38", color.RGBA{0x61, 0x61, 0x61, 0xFF}, false}
	Grey39               = X11Color{"grey39", color.RGBA{0x63, 0x63, 0x63, 0xFF}, false}
	Grey4                = X11Color{"grey4", color.RGBA{0x0A, 0x0A, 0x0A, 0xFF}, false}
	Grey40               = X11Color{"grey40", color.RGBA{0x66, 0x66, 0x66, 0xFF}, false}
	Grey41               = X11Color{"grey41", color.RGBA{0x69, 0x69, 0x69, 0xFF}, false}
	Grey42               = X11Color{"grey42", color.RGBA{0x6B, 0x6B, 0x6B, 0xFF}, false}
	Grey43               = X11Color{"grey43", color.RGBA{0x6E, 0x6E, 0x6E, 0xFF}, false}
	Grey44               = X11Color{"grey44", color.RGBA{0x70, 0x70, 0x70, 0xFF}, false}
	Grey45               = X11Color{"grey45", color.RGBA{0x73, 0x73, 0x73, 0xFF}, false}
	Grey46               = X11Color{"grey46", color.RGBA{0x75, 0x75, 0x75, 0xFF}, false}
	Grey47               = X11Color{"grey47", color.RGBA{0x78, 0x78, 0x78, 0xFF}, true}
	Grey48               = X11Color{"grey48", color.RGBA{0x7A, 0x7A, 0x7A, 0xFF}, true}
	Grey49               = X11Color{"grey49", color.RGBA{0x7D, 0x7D, 0x7D, 0xFF}, true}
	Grey5                = X11Color{"grey5", color.RGBA{0x0D, 0x0D, 0x0D, 0xFF}, false}
	Grey50               = X11Color{"grey50", color.RGBA{0x7F, 0x7F, 0x7F, 0xFF}, true}
	Grey51               = X11Color{"grey51", color.RGBA{0x82, 0x82, 0x82, 0xFF}, true}
	Grey52               = X11Color{"grey52", color.RGBA{0x85, 0x85, 0x85, 0xFF}, true}
	Grey53               = X11Color{"grey53", color.RGBA{0x87, 0x87, 0x87, 0xFF}, true}
	Grey54               = X11Color{"grey54", color.RGBA{0x8A, 0x8A, 0x8A, 0xFF}, true}
	Grey55               = X11Color{"grey55", color.RGBA{0x8C, 0x8C, 0x8C, 0xFF}, true}
	Grey56               = X11Color{"grey56", color.RGBA{0x8F, 0x8F, 0x8F, 0xFF}, true}
	Grey57               = X11Color{"grey57", color.RGBA{0x91, 0x91, 0x91, 0xFF}, true}
	Grey58               = X11Color{"grey58", color.RGBA{0x94, 0x94, 0x94, 0xFF}, true}
	Grey59               = X11Color{"grey59", color.RGBA{0x96, 0x96, 0x96, 0xFF}, true}
	Grey6                = X11Color{"grey6", color.RGBA{0x0F, 0x0F, 0x0F, 0xFF}, false}
	Grey60               = X11Color{"grey60", color.RGBA{0x99, 0x99, 0x99, 0xFF}, true}
	Grey61               = X11Color{"grey61", color.RGBA{0x9C, 0x9C, 0x9C, 0xFF}, true}
	Grey62               = X11Color{"grey62", color.RGBA{0x9E, 0x9E, 0x9E, 0xFF}, true}
	Grey63               = X11Color{"grey63", color.RGBA{0xA1, 0xA1, 0xA1, 0xFF}, true}
	Grey64               = X11Color{"grey64", color.RGBA{0xA3, 0xA3, 0xA3, 0xFF}, true}
	Grey65               = X11Color{"grey65", color.RGBA{0xA6, 0xA6, 0xA6, 0xFF}, true}
	Grey66               = X11Color{"grey66", color.RGBA{0xA8, 0xA8, 0xA8, 0xFF}, true}
	Grey67               = X11Color{"grey67", color.RGBA{0xAB, 0xAB, 0xAB, 0xFF}, true}
	Grey68               = X11Color{"grey68", color.RGBA{0xAD, 0xAD, 0xAD, 0xFF}, true}
	Grey69               = X11Color{"grey69", color.RGBA{0xB0, 0xB0, 0xB0, 0xFF}, true}
	Grey7                = X11Color{"grey7", color.RGBA{0x12, 0x12, 0x12, 0xFF}, false}
	Grey70               = X11Color{"grey70", color.RGBA{0xB3, 0xB3, 0xB3, 0xFF}, true}
	Grey71               = X11Color{"grey71", color.RGBA{0xB5, 0xB5, 0xB5, 0xFF}, true}
	Grey72               = X11Color{"grey72", color.RGBA{0xB8, 0xB8, 0xB8, 0xFF}, true}
	Grey73               = X11Color{"grey73", color.RGBA{0xBA, 0xBA, 0xBA, 0xFF}, true}
	Grey74               = X11Color{"grey74", color.RGBA{0xBD, 0xBD, 0xBD, 0xFF}, true}
	Grey75               = X11Color{"grey75", color.RGBA{0xBF, 0xBF, 0xBF, 0xFF}, true}
	Grey76               = X11Color{"grey76", color.RGBA{0xC2, 0xC2, 0xC2, 0xFF}, true}
	Grey77               = X11Color{"grey77", color.RGBA{0xC4, 0xC4, 0xC4, 0xFF}, true}
	Grey78               = X11Color{"grey78", color.RGBA{0xC7, 0xC7, 0xC7, 0xFF}, true}
	Grey79               = X11Color{"grey79", color.RGBA{0xC9, 0xC9, 0xC9, 0xFF}, true}
	Grey8                = X11Color{"grey8", color.RGBA{0x14, 0x14, 0x14, 0xFF}, false}
	Grey80               = X11Color{"grey80", color.RGBA{0xCC, 0xCC, 0xCC, 0xFF}, true}
	Grey81               = X11Color{"grey81", color.RGBA{0xCF, 0xCF, 0xCF, 0xFF}, true}
	Grey82               = X11Color{"grey82", color.RGBA{0xD1, 0xD1, 0xD1, 0xFF}, true}
	Grey83               = X11Color{"grey83", color.RGBA{0xD4, 0xD4, 0xD4, 0xFF}, true}
	Grey84               = X11Color{"grey84", color.RGBA{0xD6, 0xD6, 0xD6, 0xFF}, true}
	Grey85               = X11Color{"grey85", color.RGBA{0xD9, 0xD9, 0xD9, 0xFF}, true}
	Grey86               = X11Color{"grey86", color.RGBA{0xDB, 0xDB, 0xDB, 0xFF}, true}
	Grey87               = X11Color{"grey87", color.RGBA{0xDE, 0xDE, 0xDE, 0xFF}, true}
	Grey88               = X11Color{"grey88", color.RGBA{0xE0, 0xE0, 0xE0, 0xFF}, true}
	Grey89               = X11Color{"grey89", color.RGBA{0xE3, 0xE3, 0xE3, 0xFF}, true}
	Grey9                = X11Color{"grey9", color.RGBA{0x17, 0x17, 0x17, 0xFF}, false}
	Grey90               = X11Color{"grey90", color.RGBA{0xE5, 0xE5, 0xE5, 0xFF}, true}
	Grey91               = X11Color{"grey91", color.RGBA{0xE8, 0xE8, 0xE8, 0xFF}, true}
	Grey92               = X11Color{"grey92", color.RGBA{0xEB, 0xEB, 0xEB, 0xFF}, true}
	Grey93               = X11Color{"grey93", color.RGBA{0xED, 0xED, 0xED, 0xFF}, true}
	Grey94               = X11Color{"grey94", color.RGBA{0xF0, 0xF0, 0xF0, 0xFF}, true}
	Grey95               = X11Color{"grey95", color.RGBA{0xF2, 0xF2, 0xF2, 0xFF}, true}
	Grey96               = X11Color{"grey96", color.RGBA{0xF5, 0xF5, 0xF5, 0xFF}, true}
	Grey97               = X11Color{"grey97", color.RGBA{0xF7, 0xF7, 0xF7, 0xFF}, true}
	Grey98               = X11Color{"grey98", color.RGBA{0xFA, 0xFA, 0xFA, 0xFF}, true}
	Grey99               = X11Color{"grey99", color.RGBA{0xFC, 0xFC, 0xFC, 0xFF}, true}
	Honeydew             = X11Color{"Honeydew", color.RGBA{0xF0, 0xFF, 0xF0, 0xFF}, true}
	Honeydew1            = X11Color{"honeydew1", color.RGBA{0xF0, 0xFF, 0xF0, 0xFF}, true}
	Honeydew2            = X11Color{"honeydew2", color.RGBA{0xE0, 0xEE, 0xE0, 0xFF}, true}
	Honeydew3            = X11Color{"honeydew3", color.RGBA{0xC1, 0xCD, 0xC1, 0xFF}, true}
	Honeydew4            = X11Color{"honeydew4", color.RGBA{0x83, 0x8B, 0x83, 0xFF}, true}
	HotPink              = X11Color{"Hot Pink", color.RGBA{0xFF, 0x69, 0xB4, 0xFF}, false}
	Hotpink              = X11Color{"HotPink", color.RGBA{0xFF, 0x69, 0xB4, 0xFF}, true}
	Hotpink1             = X11Color{"HotPink1", color.RGBA{0xFF, 0x6E, 0xB4, 0xFF}, true}
	Hotpink2             = X11Color{"HotPink2", color.RGBA{0xEE, 0x6A, 0xA7, 0xFF}, true}
	Hotpink3             = X11Color{"HotPink3", color.RGBA{0xCD, 0x60, 0x90, 0xFF}, true}
	Hotpink4             = X11Color{"HotPink4", color.RGBA{0x8B, 0x3A, 0x62, 0xFF}, false}
	IndianRed            = X11Color{"Indian Red", color.RGBA{0xCD, 0x5C, 0x5C, 0xFF}, false}
	Indianred            = X11Color{"IndianRed", color.RGBA{0xCD, 0x5C, 0x5C, 0xFF}, true}
	Indianred1           = X11Color{"IndianRed1", color.RGBA{0xFF, 0x6A, 0x6A, 0xFF}, true}
	Indianred2           = X11Color{"IndianRed2", color.RGBA{0xEE, 0x63, 0x63, 0xFF}, true}
	Indianred3           = X11Color{"IndianRed3", color.RGBA{0xCD, 0x55, 0x55, 0xFF}, true}
	Indianred4           = X11Color{"IndianRed4", color.RGBA{0x8B, 0x3A, 0x3A, 0xFF}, false}
	Indigo               = X11Color{"Indigo", color.RGBA{0x4B, 0x00, 0x82, 0xFF}, false}
	Ivory                = X11Color{"Ivory", color.RGBA{0xFF, 0xFF, 0xF0, 0xFF}, true}
	Ivory1               = X11Color{"ivory1", color.RGBA{0xFF, 0xFF, 0xF0, 0xFF}, true}
	Ivory2               = X11Color{"ivory2", color.RGBA{0xEE, 0xEE, 0xE0, 0xFF}, true}
	Ivory3               = X11Color{"ivory3", color.RGBA{0xCD, 0xCD, 0xC1, 0xFF}, true}
	Ivory4               = X11Color{"ivory4", color.RGBA{0x8B, 0x8B, 0x83, 0xFF}, true}
	Khaki                = X11Color{"Khaki", color.RGBA{0xF0, 0xE6, 0x8C, 0xFF}, false}
	Khaki1               = X11Color{"khaki1", color.RGBA{0xFF, 0xF6, 0x8F, 0xFF}, true}
	Khaki2               = X11Color{"khaki2", color.RGBA{0xEE, 0xE6, 0x85, 0xFF}, true}
	Khaki3               = X11Color{"khaki3", color.RGBA{0xCD, 0xC6, 0x73, 0xFF}, true}
	Khaki4               = X11Color{"khaki4", color.RGBA{0x8B, 0x86, 0x4E, 0xFF}, true}
	Lavender             = X11Color{"Lavender", color.RGBA{0xE6, 0xE6, 0xFA, 0xFF}, false}
	LavenderBlush        = X11Color{"Lavender Blush", color.RGBA{0xFF, 0xF0, 0xF5, 0xFF}, true}
	Lavenderblush        = X11Color{"LavenderBlush", color.RGBA{0xFF, 0xF0, 0xF5, 0xFF}, true}
	Lavenderblush1       = X11Color{"LavenderBlush1", color.RGBA{0xFF, 0xF0, 0xF5, 0xFF}, true}
	Lavenderblush2       = X11Color{"LavenderBlush2", color.RGBA{0xEE, 0xE0, 0xE5, 0xFF}, true}
	Lavenderblush3       = X11Color{"LavenderBlush3", color.RGBA{0xCD, 0xC1, 0xC5, 0xFF}, true}
	Lavenderblush4       = X11Color{"LavenderBlush4", color.RGBA{0x8B, 0x83, 0x86, 0xFF}, true}
	LawnGreen            = X11Color{"Lawn Green", color.RGBA{0x7C, 0xFC, 0x00, 0xFF}, false}
	Lawngreen            = X11Color{"LawnGreen", color.RGBA{0x7C, 0xFC, 0x00, 0xFF}, true}
	LemonChiffon         = X11Color{"Lemon Chiffon", color.RGBA{0xFF, 0xFA, 0xCD, 0xFF}, true}
	Lemonchiffon         = X11Color{"LemonChiffon", color.RGBA{0xFF, 0xFA, 0xCD, 0xFF}, true}
	Lemonchiffon1        = X11Color{"LemonChiffon1", color.RGBA{0xFF, 0xFA, 0xCD, 0xFF}, true}
	Lemonchiffon2        = X11Color{"LemonChiffon2", color.RGBA{0xEE, 0xE9, 0xBF, 0xFF}, true}
	Lemonchiffon3        = X11Color{"LemonChiffon3", color.RGBA{0xCD, 0xC9, 0xA5, 0xFF}, true}
	Lemonchiffon4        = X11Color{"LemonChiffon4", color.RGBA{0x8B, 0x89, 0x70, 0xFF}, true}
	LightBlue            = X11Color{"Light Blue", color.RGBA{0xAD, 0xD8, 0xE6, 0xFF}, false}
	LightCoral           = X11Color{"Light Coral", color.RGBA{0xF0, 0x80, 0x80, 0xFF}, false}
	LightCyan            = X11Color{"Light Cyan", color.RGBA{0xE0, 0xFF, 0xFF, 0xFF}, true}
	LightGoldenrod       = X11Color{"Light Goldenrod", color.RGBA{0xFA, 0xFA, 0xD2, 0xFF}, true}
	LightGoldenrodYellow = X11Color{"light goldenrod yellow", color.RGBA{0xFA, 0xFA, 0xD2, 0xFF}, true}
	LightGray            = X11Color{"Light Gray", color.RGBA{0xD3, 0xD3, 0xD3, 0xFF}, false}
	LightGreen           = X11Color{"Light Green", color.RGBA{0x90, 0xEE, 0x90, 0xFF}, false}
	LightGrey            = X11Color{"light grey", color.RGBA{0xD3, 0xD3, 0xD3, 0xFF}, true}
	LightPink            = X11Color{"Light Pink", color.RGBA{0xFF, 0xB6, 0xC1, 0xFF}, false}
	LightSalmon          = X11Color{"Light Salmon", color.RGBA{0xFF, 0xA0, 0x7A, 0xFF}, false}
	LightSeaGreen        = X11Color{"Light Sea Green", color.RGBA{0x20, 0xB2, 0xAA, 0xFF}, false}
	LightSkyBlue         = X11Color{"Light Sky Blue", color.RGBA{0x87, 0xCE, 0xFA, 0xFF}, false}
	LightSlateBlue       = X11Color{"light slate blue", color.RGBA{0x84, 0x70, 0xFF, 0xFF}, true}
	LightSlateGray       = X11Color{"Light Slate Gray", color.RGBA{0x77, 0x88, 0x99, 0xFF}, false}
	LightSlateGrey       = X11Color{"light slate grey", color.RGBA{0x77, 0x88, 0x99, 0xFF}, true}
	LightSteelBlue       = X11Color{"Light Steel Blue", color.RGBA{0xB0, 0xC4, 0xDE, 0xFF}, false}
	LightYellow          = X11Color{"Light Yellow", color.RGBA{0xFF, 0xFF, 0xE0, 0xFF}, true}
	Lightblue            = X11Color{"LightBlue", color.RGBA{0xAD, 0xD8, 0xE6, 0xFF}, true}
	Lightblue1           = X11Color{"LightBlue1", color.RGBA{0xBF, 0xEF, 0xFF, 0xFF}, true}
	Lightblue2           = X11Color{"LightBlue2", color.RGBA{0xB2, 0xDF, 0xEE, 0xFF}, true}
	Lightblue3           = X11Color{"LightBlue3", color.RGBA{0x9A, 0xC0, 0xCD, 0xFF}, true}
	Lightblue4           = X11Color{"LightBlue4", color.RGBA{0x68, 0x83, 0x8B, 0xFF}, true}
	Lightcoral           = X11Color{"LightCoral", color.RGBA{0xF0, 0x80, 0x80, 0xFF}, true}
	Lightcyan            = X11Color{"LightCyan", color.RGBA{0xE0, 0xFF, 0xFF, 0xFF}, true}
	Lightcyan1           = X11Color{"LightCyan1", color.RGBA{0xE0, 0xFF, 0xFF, 0xFF}, true}
	Lightcyan2           = X11Color{"LightCyan2", color.RGBA{0xD1, 0xEE, 0xEE, 0xFF}, true}
	Lightcyan3           = X11Color{"LightCyan3", color.RGBA{0xB4, 0xCD, 0xCD, 0xFF}, true}
	Lightcyan4           = X11Color{"LightCyan4", color.RGBA{0x7A, 0x8B, 0x8B, 0xFF}, true}
	Lightgoldenrod       = X11Color{"LightGoldenrod", color.RGBA{0xEE, 0xDD, 0x82, 0xFF}, true}
	Lightgoldenrod1      = X11Color{"LightGoldenrod1", color.RGBA{0xFF, 0xEC, 0x8B, 0xFF}, true}
	Lightgoldenrod2      = X11Color{"LightGoldenrod2", color.RGBA{0xEE, 0xDC, 0x82, 0xFF}, true}
	Lightgoldenrod3      = X11Color{"LightGoldenrod3", color.RGBA{0xCD, 0xBE, 0x70, 0xFF}, true}
	Lightgoldenrod4      = X11Color{"LightGoldenrod4", color.RGBA{0x8B, 0x81, 0x4C, 0xFF}, true}
	Lightgoldenrodyellow = X11Color{"LightGoldenrodYellow", color.RGBA{0xFA, 0xFA, 0xD2, 0xFF}, true}
	Lightgray            = X11Color{"LightGray", color.RGBA{0xD3, 0xD3, 0xD3, 0xFF}, true}
	Lightgreen           = X11Color{"LightGreen", color.RGBA{0x90, 0xEE, 0x90, 0xFF}, true}
	Lightgrey            = X11Color{"LightGrey", color.RGBA{0xD3, 0xD3, 0xD3, 0xFF}, true}
	Lightpink            = X11Color{"LightPink", color.RGBA{0xFF, 0xB6, 0xC1, 0xFF}, true}
	Lightpink1           = X11Color{"LightPink1", color.RGBA{0xFF, 0xAE, 0xB9, 0xFF}, true}
	Lightpink2           = X11Color{"LightPink2", color.RGBA{0xEE, 0xA2, 0xAD, 0xFF}, true}
	Lightpink3           = X11Color{"LightPink3", color.RGBA{0xCD, 0x8C, 0x95, 0xFF}, true}
	Lightpink4           = X11Color{"LightPink4", color.RGBA{0x8B, 0x5F, 0x65, 0xFF}, false}
	Lightsalmon          = X11Color{"LightSalmon", color.RGBA{0xFF, 0xA0, 0x7A, 0xFF}, true}
	Lightsalmon1         = X11Color{"LightSalmon1", color.RGBA{0xFF, 0xA0, 0x7A, 0xFF}, true}
	Lightsalmon2         = X11Color{"LightSalmon2", color.RGBA{0xEE, 0x95, 0x72, 0xFF}, true}
	Lightsalmon3         = X11Color{"LightSalmon3", color.RGBA{0xCD, 0x81, 0x62, 0xFF}, true}
	Lightsalmon4         = X11Color{"LightSalmon4", color.RGBA{0x8B, 0x57, 0x42, 0xFF}, false}
	Lightseagreen        = X11Color{"LightSeaGreen", color.RGBA{0x20, 0xB2, 0xAA, 0xFF}, true}
	Lightskyblue         = X11Color{"LightSkyBlue", color.RGBA{0x87, 0xCE, 0xFA, 0xFF}, true}
	Lightskyblue1        = X11Color{"LightSkyBlue1", color.RGBA{0xB0, 0xE2, 0xFF, 0xFF}, true}
	Lightskyblue2        = X11Color{"LightSkyBlue2", color.RGBA{0xA4, 0xD3, 0xEE, 0xFF}, true}
	Lightskyblue3        = X11Color{"LightSkyBlue3", color.RGBA{0x8D, 0xB6, 0xCD, 0xFF}, true}
	Lightskyblue4        = X11Color{"LightSkyBlue4", color.RGBA{0x60, 0x7B, 0x8B, 0xFF}, true}
	Lightslateblue       = X11Color{"LightSlateBlue", color.RGBA{0x84, 0x70, 0xFF, 0xFF}, true}
	Lightslategray       = X11Color{"LightSlateGray", color.RGBA{0x77, 0x88, 0x99, 0xFF}, true}
	Lightslategrey       = X11Color{"LightSlateGrey", color.RGBA{0x77, 0x88, 0x99, 0xFF}, true}
	Lightsteelblue       = X11Color{"LightSteelBlue", color.RGBA{0xB0, 0xC4, 0xDE, 0xFF}, true}
	Lightsteelblue1      = X11Color{"LightSteelBlue1", color.RGBA{0xCA, 0xE1, 0xFF, 0xFF}, true}
	Lightsteelblue2      = X11Color{"LightSteelBlue2", color.RGBA{0xBC, 0xD2, 0xEE, 0xFF}, true}
	Lightsteelblue3      = X11Color{"LightSteelBlue3", color.RGBA{0xA2, 0xB5, 0xCD, 0xFF}, true}
	Lightsteelblue4      = X11Color{"LightSteelBlue4", color.RGBA{0x6E, 0x7B, 0x8B, 0xFF}, true}
	Lightyellow          = X11Color{"LightYellow", color.RGBA{0xFF, 0xFF, 0xE0, 0xFF}, true}
	Lightyellow1         = X11Color{"LightYellow1", color.RGBA{0xFF, 0xFF, 0xE0, 0xFF}, true}
	Lightyellow2         = X11Color{"LightYellow2", color.RGBA{0xEE, 0xEE, 0xD1, 0xFF}, true}
	Lightyellow3         = X11Color{"LightYellow3", color.RGBA{0xCD, 0xCD, 0xB4, 0xFF}, true}
	Lightyellow4         = X11Color{"LightYellow4", color.RGBA{0x8B, 0x8B, 0x7A, 0xFF}, true}
	Lime                 = X11Color{"Lime", color.RGBA{0x00, 0xFF, 0x00, 0xFF}, false}
	LimeGreen            = X11Color{"Lime Green", color.RGBA{0x32, 0xCD, 0x32, 0xFF}, false}
	Limegreen            = X11Color{"LimeGreen", color.RGBA{0x32, 0xCD, 0x32, 0xFF}, true}
	Linen                = X11Color{"Linen", color.RGBA{0xFA, 0xF0, 0xE6, 0xFF}, true}
	Magenta              = X11Color{"Magenta", color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, false}
	Magenta1             = X11Color{"magenta1", color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, true}
	Magenta2             = X11Color{"magenta2", color.RGBA{0xEE, 0x00, 0xEE, 0xFF}, true}
	Magenta3             = X11Color{"magenta3", color.RGBA{0xCD, 0x00, 0xCD, 0xFF}, false}
	Magenta4             = X11Color{"magenta4", color.RGBA{0x8B, 0x00, 0x8B, 0xFF}, false}
	Maroon               = X11Color{"maroon", color.RGBA{0xB0, 0x30, 0x60, 0xFF}, false}
	Maroon1              = X11Color{"maroon1", color.RGBA{0xFF, 0x34, 0xB3, 0xFF}, true}
	Maroon2              = X11Color{"maroon2", color.RGBA{0xEE, 0x30, 0xA7, 0xFF}, true}
	Maroon3              = X11Color{"maroon3", color.RGBA{0xCD, 0x29, 0x90, 0xFF}, false}
	Maroon4              = X11Color{"maroon4", color.RGBA{0x8B, 0x1C, 0x62, 0xFF}, false}
	MaroonW3C            = X11Color{"Maroon (W3C)", color.RGBA{0x7F, 0x00, 0x00, 0xFF}, false}
	MaroonX11            = X11Color{"Maroon (X11)", color.RGBA{0xB0, 0x30, 0x60, 0xFF}, false}
	MediumAquamarine     = X11Color{"Medium Aquamarine", color.RGBA{0x66, 0xCD, 0xAA, 0xFF}, false}
	MediumBlue           = X11Color{"Medium Blue", color.RGBA{0x00, 0x00, 0xCD, 0xFF}, false}
	MediumOrchid         = X11Color{"Medium Orchid", color.RGBA{0xBA, 0x55, 0xD3, 0xFF}, false}
	MediumPurple         = X11Color{"Medium Purple", color.RGBA{0x93, 0x70, 0xDB, 0xFF}, false}
	MediumSeaGreen       = X11Color{"Medium Sea Green", color.RGBA{0x3C, 0xB3, 0x71, 0xFF}, false}
	MediumSlateBlue      = X11Color{"Medium Slate Blue", color.RGBA{0x7B, 0x68, 0xEE, 0xFF}, false}
	MediumSpringGreen    = X11Color{"Medium Spring Green", color.RGBA{0x00, 0xFA, 0x9A, 0xFF}, false}
	MediumTurquoise      = X11Color{"Medium Turquoise", color.RGBA{0x48, 0xD1, 0xCC, 0xFF}, false}
	MediumVioletRed      = X11Color{"Medium Violet Red", color.RGBA{0xC7, 0x15, 0x85, 0xFF}, false}
	Mediumaquamarine     = X11Color{"MediumAquamarine", color.RGBA{0x66, 0xCD, 0xAA, 0xFF}, true}
	Mediumblue           = X11Color{"MediumBlue", color.RGBA{0x00, 0x00, 0xCD, 0xFF}, false}
	Mediumorchid         = X11Color{"MediumOrchid", color.RGBA{0xBA, 0x55, 0xD3, 0xFF}, true}
	Mediumorchid1        = X11Color{"MediumOrchid1", color.RGBA{0xE0, 0x66, 0xFF, 0xFF}, true}
	Mediumorchid2        = X11Color{"MediumOrchid2", color.RGBA{0xD1, 0x5F, 0xEE, 0xFF}, true}
	Mediumorchid3        = X11Color{"MediumOrchid3", color.RGBA{0xB4, 0x52, 0xCD, 0xFF}, true}
	Mediumorchid4        = X11Color{"MediumOrchid4", color.RGBA{0x7A, 0x37, 0x8B, 0xFF}, false}
	Mediumpurple         = X11Color{"MediumPurple", color.RGBA{0x93, 0x70, 0xDB, 0xFF}, true}
	Mediumpurple1        = X11Color{"MediumPurple1", color.RGBA{0xAB, 0x82, 0xFF, 0xFF}, true}
	Mediumpurple2        = X11Color{"MediumPurple2", color.RGBA{0x9F, 0x79, 0xEE, 0xFF}, true}
	Mediumpurple3        = X11Color{"MediumPurple3", color.RGBA{0x89, 0x68, 0xCD, 0xFF}, true}
	Mediumpurple4        = X11Color{"MediumPurple4", color.RGBA{0x5D, 0x47, 0x8B, 0xFF}, false}
	Mediumseagreen       = X11Color{"MediumSeaGreen", color.RGBA{0x3C, 0xB3, 0x71, 0xFF}, true}
	Mediumslateblue      = X11Color{"MediumSlateBlue", color.RGBA{0x7B, 0x68, 0xEE, 0xFF}, true}
	Mediumspringgreen    = X11Color{"MediumSpringGreen", color.RGBA{0x00, 0xFA, 0x9A, 0xFF}, true}
	Mediumturquoise      = X11Color{"MediumTurquoise", color.RGBA{0x48, 0xD1, 0xCC, 0xFF}, true}
	Mediumvioletred      = X11Color{"MediumVioletRed", color.RGBA{0xC7, 0x15, 0x85, 0xFF}, false}
	MidnightBlue         = X11Color{"Midnight Blue", color.RGBA{0x19, 0x19, 0x70, 0xFF}, false}
	Midnightblue         = X11Color{"MidnightBlue", color.RGBA{0x19, 0x19, 0x70, 0xFF}, false}
	MintCream            = X11Color{"Mint Cream", color.RGBA{0xF5, 0xFF, 0xFA, 0xFF}, true}
	Mintcream            = X11Color{"MintCream", color.RGBA{0xF5, 0xFF, 0xFA, 0xFF}, true}
	MistyRose            = X11Color{"Misty Rose", color.RGBA{0xFF, 0xE4, 0xE1, 0xFF}, true}
	Mistyrose            = X11Color{"MistyRose", color.RGBA{0xFF, 0xE4, 0xE1, 0xFF}, true}
	Mistyrose1           = X11Color{"MistyRose1", color.RGBA{0xFF, 0xE4, 0xE1, 0xFF}, true}
	Mistyrose2           = X11Color{"MistyRose2", color.RGBA{0xEE, 0xD5, 0xD2, 0xFF}, true}
	Mistyrose3           = X11Color{"MistyRose3", color.RGBA{0xCD, 0xB7, 0xB5, 0xFF}, true}
	Mistyrose4           = X11Color{"MistyRose4", color.RGBA{0x8B, 0x7D, 0x7B, 0xFF}, true}
	Moccasin             = X11Color{"Moccasin", color.RGBA{0xFF, 0xE4, 0xB5, 0xFF}, false}
	NavajoWhite          = X11Color{"Navajo White", color.RGBA{0xFF, 0xDE, 0xAD, 0xFF}, false}
	Navajowhite          = X11Color{"NavajoWhite", color.RGBA{0xFF, 0xDE, 0xAD, 0xFF}, true}
	Navajowhite1         = X11Color{"NavajoWhite1", color.RGBA{0xFF, 0xDE, 0xAD, 0xFF}, true}
	Navajowhite2         = X11Color{"NavajoWhite2", color.RGBA{0xEE, 0xCF, 0xA1, 0xFF}, true}
	Navajowhite3         = X11Color{"NavajoWhite3", color.RGBA{0xCD, 0xB3, 0x8B, 0xFF}, true}
	Navajowhite4         = X11Color{"NavajoWhite4", color.RGBA{0x8B, 0x79, 0x5E, 0xFF}, true}
	Navy                 = X11Color{"Navy", color.RGBA{0x00, 0x00, 0x80, 0xFF}, false}
	NavyBlue             = X11Color{"navy blue", color.RGBA{0x00, 0x00, 0x80, 0xFF}, false}
	Navyblue             = X11Color{"NavyBlue", color.RGBA{0x00, 0x00, 0x80, 0xFF}, false}
	OldLace              = X11Color{"Old Lace", color.RGBA{0xFD, 0xF5, 0xE6, 0xFF}, true}
	Oldlace              = X11Color{"OldLace", color.RGBA{0xFD, 0xF5, 0xE6, 0xFF}, true}
	Olive                = X11Color{"Olive", color.RGBA{0x80, 0x80, 0x00, 0xFF}, false}
	OliveDrab            = X11Color{"Olive Drab", color.RGBA{0x6B, 0x8E, 0x23, 0xFF}, false}
	Olivedrab            = X11Color{"OliveDrab", color.RGBA{0x6B, 0x8E, 0x23, 0xFF}, true}
	Olivedrab1           = X11Color{"OliveDrab1", color.RGBA{0xC0, 0xFF, 0x3E, 0xFF}, true}
	Olivedrab2           = X11Color{"OliveDrab2", color.RGBA{0xB3, 0xEE, 0x3A, 0xFF}, true}
	Olivedrab3           = X11Color{"OliveDrab3", color.RGBA{0x9A, 0xCD, 0x32, 0xFF}, true}
	Olivedrab4           = X11Color{"OliveDrab4", color.RGBA{0x69, 0x8B, 0x22, 0xFF}, true}
	Orange               = X11Color{"Orange", color.RGBA{0xFF, 0xA5, 0x00, 0xFF}, false}
	Orange1              = X11Color{"orange1", color.RGBA{0xFF, 0xA5, 0x00, 0xFF}, true}
	Orange2              = X11Color{"orange2", color.RGBA{0xEE, 0x9A, 0x00, 0xFF}, true}
	Orange3              = X11Color{"orange3", color.RGBA{0xCD, 0x85, 0x00, 0xFF}, true}
	Orange4              = X11Color{"orange4", color.RGBA{0x8B, 0x5A, 0x00, 0xFF}, false}
	OrangeRed            = X11Color{"Orange Red", color.RGBA{0xFF, 0x45, 0x00, 0xFF}, false}
	Orangered            = X11Color{"OrangeRed", color.RGBA{0xFF, 0x45, 0x00, 0xFF}, true}
	Orangered1           = X11Color{"OrangeRed1", color.RGBA{0xFF, 0x45, 0x00, 0xFF}, true}
	Orangered2           = X11Color{"OrangeRed2", color.RGBA{0xEE, 0x40, 0x00, 0xFF}, true}
	Orangered3           = X11Color{"OrangeRed3", color.RGBA{0xCD, 0x37, 0x00, 0xFF}, false}
	Orangered4           = X11Color{"OrangeRed4", color.RGBA{0x8B, 0x25, 0x00, 0xFF}, false}
	Orchid               = X11Color{"Orchid", color.RGBA{0xDA, 0x70, 0xD6, 0xFF}, false}
	Orchid1              = X11Color{"orchid1", color.RGBA{0xFF, 0x83, 0xFA, 0xFF}, true}
	Orchid2              = X11Color{"orchid2", color.RGBA{0xEE, 0x7A, 0xE9, 0xFF}, true}
	Orchid3              = X11Color{"orchid3", color.RGBA{0xCD, 0x69, 0xC9, 0xFF}, true}
	Orchid4              = X11Color{"orchid4", color.RGBA{0x8B, 0x47, 0x89, 0xFF}, false}
	PaleGoldenrod        = X11Color{"Pale Goldenrod", color.RGBA{0xEE, 0xE8, 0xAA, 0xFF}, false}
	PaleGreen            = X11Color{"Pale Green", color.RGBA{0x98, 0xFB, 0x98, 0xFF}, false}
	PaleTurquoise        = X11Color{"Pale Turquoise", color.RGBA{0xAF, 0xEE, 0xEE, 0xFF}, false}
	PaleVioletRed        = X11Color{"Pale Violet Red", color.RGBA{0xDB, 0x70, 0x93, 0xFF}, false}
	Palegoldenrod        = X11Color{"PaleGoldenrod", color.RGBA{0xEE, 0xE8, 0xAA, 0xFF}, true}
	Palegreen            = X11Color{"PaleGreen", color.RGBA{0x98, 0xFB, 0x98, 0xFF}, true}
	Palegreen1           = X11Color{"PaleGreen1", color.RGBA{0x9A, 0xFF, 0x9A, 0xFF}, true}
	Palegreen2           = X11Color{"PaleGreen2", color.RGBA{0x90, 0xEE, 0x90, 0xFF}, true}
	Palegreen3           = X11Color{"PaleGreen3", color.RGBA{0x7C, 0xCD, 0x7C, 0xFF}, true}
	Palegreen4           = X11Color{"PaleGreen4", color.RGBA{0x54, 0x8B, 0x54, 0xFF}, true}
	Paleturquoise        = X11Color{"PaleTurquoise", color.RGBA{0xAF, 0xEE, 0xEE, 0xFF}, true}
	Paleturquoise1       = X11Color{"PaleTurquoise1", color.RGBA{0xBB, 0xFF, 0xFF, 0xFF}, true}
	Paleturquoise2       = X11Color{"PaleTurquoise2", color.RGBA{0xAE, 0xEE, 0xEE, 0xFF}, true}
	Paleturquoise3       = X11Color{"PaleTurquoise3", color.RGBA{0x96, 0xCD, 0xCD, 0xFF}, true}
	Paleturquoise4       = X11Color{"PaleTurquoise4", color.RGBA{0x66, 0x8B, 0x8B, 0xFF}, true}
	Palevioletred        = X11Color{"PaleVioletRed", color.RGBA{0xDB, 0x70, 0x93, 0xFF}, true}
	Palevioletred1       = X11Color{"PaleVioletRed1", color.RGBA{0xFF, 0x82, 0xAB, 0xFF}, true}
	Palevioletred2       = X11Color{"PaleVioletRed2", color.RGBA{0xEE, 0x79, 0x9F, 0xFF}, true}
	Palevioletred3       = X11Color{"PaleVioletRed3", color.RGBA{0xCD, 0x68, 0x89, 0xFF}, true}
	Palevioletred4       = X11Color{"PaleVioletRed4", color.RGBA{0x8B, 0x47, 0x5D, 0xFF}, false}
	PapayaWhip           = X11Color{"Papaya Whip", color.RGBA{0xFF, 0xEF, 0xD5, 0xFF}, true}
	Papayawhip           = X11Color{"PapayaWhip", color.RGBA{0xFF, 0xEF, 0xD5, 0xFF}, true}
	PeachPuff            = X11Color{"Peach Puff", color.RGBA{0xFF, 0xDA, 0xB9, 0xFF}, false}
	Peachpuff            = X11Color{"PeachPuff", color.RGBA{0xFF, 0xDA, 0xB9, 0xFF}, true}
	Peachpuff1           = X11Color{"PeachPuff1", color.RGBA{0xFF, 0xDA, 0xB9, 0xFF}, true}
	Peachpuff2           = X11Color{"PeachPuff2", color.RGBA{0xEE, 0xCB, 0xAD, 0xFF}, true}
	Peachpuff3           = X11Color{"PeachPuff3", color.RGBA{0xCD, 0xAF, 0x95, 0xFF}, true}
	Peachpuff4           = X11Color{"PeachPuff4", color.RGBA{0x8B, 0x77, 0x65, 0xFF}, true}
	Peru                 = X11Color{"Peru", color.RGBA{0xCD, 0x85, 0x3F, 0xFF}, false}
	Pink                 = X11Color{"Pink", color.RGBA{0xFF, 0xC0, 0xCB, 0xFF}, false}
	Pink1                = X11Color{"pink1", color.RGBA{0xFF, 0xB5, 0xC5, 0xFF}, true}
	Pink2                = X11Color{"pink2", color.RGBA{0xEE, 0xA9, 0xB8, 0xFF}, true}
	Pink3                = X11Color{"pink3", color.RGBA{0xCD, 0x91, 0x9E, 0xFF}, true}
	Pink4                = X11Color{"pink4", color.RGBA{0x8B, 0x63, 0x6C, 0xFF}, false}
	Plum                 = X11Color{"Plum", color.RGBA{0xDD, 0xA0, 0xDD, 0xFF}, false}
	Plum1                = X11Color{"plum1", color.RGBA{0xFF, 0xBB, 0xFF, 0xFF}, true}
	Plum2                = X11Color{"plum2", color.RGBA{0xEE, 0xAE, 0xEE, 0xFF}, true}
	Plum3                = X11Color{"plum3", color.RGBA{0xCD, 0x96, 0xCD, 0xFF}, true}
	Plum4                = X11Color{"plum4", color.RGBA{0x8B, 0x66, 0x8B, 0xFF}, false}
	PowderBlue           = X11Color{"Powder Blue", color.RGBA{0xB0, 0xE0, 0xE6, 0xFF}, false}
	Powderblue           = X11Color{"PowderBlue", color.RGBA{0xB0, 0xE0, 0xE6, 0xFF}, true}
	Purple               = X11Color{"purple", color.RGBA{0xA0, 0x20, 0xF0, 0xFF}, false}
	Purple1              = X11Color{"purple1", color.RGBA{0x9B, 0x30, 0xFF, 0xFF}, false}
	Purple2              = X11Color{"purple2", color.RGBA{0x91, 0x2C, 0xEE, 0xFF}, false}
	Purple3              = X11Color{"purple3", color.RGBA{0x7D, 0x26, 0xCD, 0xFF}, false}
	Purple4              = X11Color{"purple4", color.RGBA{0x55, 0x1A, 0x8B, 0xFF}, false}
	PurpleW3C            = X11Color{"Purple (W3C)", color.RGBA{0x7F, 0x00, 0x7F, 0xFF}, false}
	PurpleX11            = X11Color{"Purple (X11)", color.RGBA{0xA0, 0x20, 0xF0, 0xFF}, false}
	RebeccaPurple        = X11Color{"rebecca purple", color.RGBA{0x66, 0x33, 0x99, 0xFF}, false}
	Rebeccapurple        = X11Color{"RebeccaPurple", color.RGBA{0x66, 0x33, 0x99, 0xFF}, false}
	Red                  = X11Color{"Red", color.RGBA{0xFF, 0x00, 0x00, 0xFF}, false}
	Red1                 = X11Color{"red1", color.RGBA{0xFF, 0x00, 0x00, 0xFF}, true}
	Red2                 = X11Color{"red2", color.RGBA{0xEE, 0x00, 0x00, 0xFF}, true}
	Red3                 = X11Color{"red3", color.RGBA{0xCD, 0x00, 0x00, 0xFF}, false}
	Red4                 = X11Color{"red4", color.RGBA{0x8B, 0x00, 0x00, 0xFF}, false}
	RosyBrown            = X11Color{"Rosy Brown", color.RGBA{0xBC, 0x8F, 0x8F, 0xFF}, false}
	Rosybrown            = X11Color{"RosyBrown", color.RGBA{0xBC, 0x8F, 0x8F, 0xFF}, true}
	Rosybrown1           = X11Color{"RosyBrown1", color.RGBA{0xFF, 0xC1, 0xC1, 0xFF}, true}
	Rosybrown2           = X11Color{"RosyBrown2", color.RGBA{0xEE, 0xB4, 0xB4, 0xFF}, true}
	Rosybrown3           = X11Color{"RosyBrown3", color.RGBA{0xCD, 0x9B, 0x9B, 0xFF}, true}
	Rosybrown4           = X11Color{"RosyBrown4", color.RGBA{0x8B, 0x69, 0x69, 0xFF}, false}
	RoyalBlue            = X11Color{"Royal Blue", color.RGBA{0x41, 0x69, 0xE1, 0xFF}, false}
	Royalblue            = X11Color{"RoyalBlue", color.RGBA{0x41, 0x69, 0xE1, 0xFF}, false}
	Royalblue1           = X11Color{"RoyalBlue1", color.RGBA{0x48, 0x76, 0xFF, 0xFF}, true}
	Royalblue2           = X11Color{"RoyalBlue2", color.RGBA{0x43, 0x6E, 0xEE, 0xFF}, true}
	Royalblue3           = X11Color{"RoyalBlue3", color.RGBA{0x3A, 0x5F, 0xCD, 0xFF}, false}
	Royalblue4           = X11Color{"RoyalBlue4", color.RGBA{0x27, 0x40, 0x8B, 0xFF}, false}
	SaddleBrown          = X11Color{"Saddle Brown", color.RGBA{0x8B, 0x45, 0x13, 0xFF}, false}
	Saddlebrown          = X11Color{"SaddleBrown", color.RGBA{0x8B, 0x45, 0x13, 0xFF}, false}
	Salmon               = X11Color{"Salmon", color.RGBA{0xFA, 0x80, 0x72, 0xFF}, false}
	Salmon1              = X11Color{"salmon1", color.RGBA{0xFF, 0x8C, 0x69, 0xFF}, true}
	Salmon2              = X11Color{"salmon2", color.RGBA{0xEE, 0x82, 0x62, 0xFF}, true}
	Salmon3              = X11Color{"salmon3", color.RGBA{0xCD, 0x70, 0x54, 0xFF}, true}
	Salmon4              = X11Color{"salmon4", color.RGBA{0x8B, 0x4C, 0x39, 0xFF}, false}
	SandyBrown           = X11Color{"Sandy Brown", color.RGBA{0xF4, 0xA4, 0x60, 0xFF}, false}
	Sandybrown           = X11Color{"SandyBrown", color.RGBA{0xF4, 0xA4, 0x60, 0xFF}, true}
	SeaGreen             = X11Color{"Sea Green", color.RGBA{0x2E, 0x8B, 0x57, 0xFF}, false}
	Seagreen             = X11Color{"SeaGreen", color.RGBA{0x2E, 0x8B, 0x57, 0xFF}, true}
	Seagreen1            = X11Color{"SeaGreen1", color.RGBA{0x54, 0xFF, 0x9F, 0xFF}, true}
	Seagreen2            = X11Color{"SeaGreen2", color.RGBA{0x4E, 0xEE, 0x94, 0xFF}, true}
	Seagreen3            = X11Color{"SeaGreen3", color.RGBA{0x43, 0xCD, 0x80, 0xFF}, true}
	Seagreen4            = X11Color{"SeaGreen4", color.RGBA{0x2E, 0x8B, 0x57, 0xFF}, true}
	Seashell             = X11Color{"Seashell", color.RGBA{0xFF, 0xF5, 0xEE, 0xFF}, true}
	Seashell1            = X11Color{"seashell1", color.RGBA{0xFF, 0xF5, 0xEE, 0xFF}, true}
	Seashell2            = X11Color{"seashell2", color.RGBA{0xEE, 0xE5, 0xDE, 0xFF}, true}
	Seashell3            = X11Color{"seashell3", color.RGBA{0xCD, 0xC5, 0xBF, 0xFF}, true}
	Seashell4            = X11Color{"seashell4", color.RGBA{0x8B, 0x86, 0x82, 0xFF}, true}
	Sienna               = X11Color{"Sienna", color.RGBA{0xA0, 0x52, 0x2D, 0xFF}, false}
	Sienna1              = X11Color{"sienna1", color.RGBA{0xFF, 0x82, 0x47, 0xFF}, true}
	Sienna2              = X11Color{"sienna2", color.RGBA{0xEE, 0x79, 0x42, 0xFF}, true}
	Sienna3              = X11Color{"sienna3", color.RGBA{0xCD, 0x68, 0x39, 0xFF}, true}
	Sienna4              = X11Color{"sienna4", color.RGBA{0x8B, 0x47, 0x26, 0xFF}, false}
	Silver               = X11Color{"Silver", color.RGBA{0xC0, 0xC0, 0xC0, 0xFF}, false}
	SkyBlue              = X11Color{"Sky Blue", color.RGBA{0x87, 0xCE, 0xEB, 0xFF}, false}
	Skyblue              = X11Color{"SkyBlue", color.RGBA{0x87, 0xCE, 0xEB, 0xFF}, true}
	Skyblue1             = X11Color{"SkyBlue1", color.RGBA{0x87, 0xCE, 0xFF, 0xFF}, true}
	Skyblue2             = X11Color{"SkyBlue2", color.RGBA{0x7E, 0xC0, 0xEE, 0xFF}, true}
	Skyblue3             = X11Color{"SkyBlue3", color.RGBA{0x6C, 0xA6, 0xCD, 0xFF}, true}
	Skyblue4             = X11Color{"SkyBlue4", color.RGBA{0x4A, 0x70, 0x8B, 0xFF}, false}
	SlateBlue            = X11Color{"Slate Blue", color.RGBA{0x6A, 0x5A, 0xCD, 0xFF}, false}
	SlateGray            = X11Color{"Slate Gray", color.RGBA{0x70, 0x80, 0x90, 0xFF}, false}
	SlateGrey            = X11Color{"slate grey", color.RGBA{0x70, 0x80, 0x90, 0xFF}, true}
	Slateblue            = X11Color{"SlateBlue", color.RGBA{0x6A, 0x5A, 0xCD, 0xFF}, false}
	Slateblue1           = X11Color{"SlateBlue1", color.RGBA{0x83, 0x6F, 0xFF, 0xFF}, true}
	Slateblue2           = X11Color{"SlateBlue2", color.RGBA{0x7A, 0x67, 0xEE, 0xFF}, true}
	Slateblue3           = X11Color{"SlateBlue3", color.RGBA{0x69, 0x59, 0xCD, 0xFF}, false}
	Slateblue4           = X11Color{"SlateBlue4", color.RGBA{0x47, 0x3C, 0x8B, 0xFF}, false}
	Slategray            = X11Color{"SlateGray", color.RGBA{0x70, 0x80, 0x90, 0xFF}, true}
	Slategray1           = X11Color{"SlateGray1", color.RGBA{0xC6, 0xE2, 0xFF, 0xFF}, true}
	Slategray2           = X11Color{"SlateGray2", color.RGBA{0xB9, 0xD3, 0xEE, 0xFF}, true}
	Slategray3           = X11Color{"SlateGray3", color.RGBA{0x9F, 0xB6, 0xCD, 0xFF}, true}
	Slategray4           = X11Color{"SlateGray4", color.RGBA{0x6C, 0x7B, 0x8B, 0xFF}, true}
	Slategrey            = X11Color{"SlateGrey", color.RGBA{0x70, 0x80, 0x90, 0xFF}, true}
	Snow                 = X11Color{"Snow", color.RGBA{0xFF, 0xFA, 0xFA, 0xFF}, true}
	Snow1                = X11Color{"snow1", color.RGBA{0xFF, 0xFA, 0xFA, 0xFF}, true}
	Snow2                = X11Color{"snow2", color.RGBA{0xEE, 0xE9, 0xE9, 0xFF}, true}
	Snow3                = X11Color{"snow3", color.RGBA{0xCD, 0xC9, 0xC9, 0xFF}, true}
	Snow4                = X11Color{"snow4", color.RGBA{0x8B, 0x89, 0x89, 0xFF}, true}
	SpringGreen          = X11Color{"Spring Green", color.RGBA{0x00, 0xFF, 0x7F, 0xFF}, false}
	Springgreen          = X11Color{"SpringGreen", color.RGBA{0x00, 0xFF, 0x7F, 0xFF}, true}
	Springgreen1         = X11Color{"SpringGreen1", color.RGBA{0x00, 0xFF, 0x7F, 0xFF}, true}
	Springgreen2         = X11Color{"SpringGreen2", color.RGBA{0x00, 0xEE, 0x76, 0xFF}, true}
	Springgreen3         = X11Color{"SpringGreen3", color.RGBA{0x00, 0xCD, 0x66, 0xFF}, true}
	Springgreen4         = X11Color{"SpringGreen4", color.RGBA{0x00, 0x8B, 0x45, 0xFF}, true}
	SteelBlue            = X11Color{"Steel Blue", color.RGBA{0x46, 0x82, 0xB4, 0xFF}, false}
	Steelblue            = X11Color{"SteelBlue", color.RGBA{0x46, 0x82, 0xB4, 0xFF}, true}
	Steelblue1           = X11Color{"SteelBlue1", color.RGBA{0x63, 0xB8, 0xFF, 0xFF}, true}
	Steelblue2           = X11Color{"SteelBlue2", color.RGBA{0x5C, 0xAC, 0xEE, 0xFF}, true}
	Steelblue3           = X11Color{"SteelBlue3", color.RGBA{0x4F, 0x94, 0xCD, 0xFF}, true}
	Steelblue4           = X11Color{"SteelBlue4", color.RGBA{0x36, 0x64, 0x8B, 0xFF}, false}
	Tan                  = X11Color{"Tan", color.RGBA{0xD2, 0xB4, 0x8C, 0xFF}, false}
	Tan1                 = X11Color{"tan1", color.RGBA{0xFF, 0xA5, 0x4F, 0xFF}, true}
	Tan2                 = X11Color{"tan2", color.RGBA{0xEE, 0x9A, 0x49, 0xFF}, true}
	Tan3                 = X11Color{"tan3", color.RGBA{0xCD, 0x85, 0x3F, 0xFF}, true}
	Tan4                 = X11Color{"tan4", color.RGBA{0x8B, 0x5A, 0x2B, 0xFF}, false}
	Teal                 = X11Color{"Teal", color.RGBA{0x00, 0x80, 0x80, 0xFF}, false}
	Thistle              = X11Color{"Thistle", color.RGBA{0xD8, 0xBF, 0xD8, 0xFF}, false}
	Thistle1             = X11Color{"thistle1", color.RGBA{0xFF, 0xE1, 0xFF, 0xFF}, true}
	Thistle2             = X11Color{"thistle2", color.RGBA{0xEE, 0xD2, 0xEE, 0xFF}, true}
	Thistle3             = X11Color{"thistle3", color.RGBA{0xCD, 0xB5, 0xCD, 0xFF}, true}
	Thistle4             = X11Color{"thistle4", color.RGBA{0x8B, 0x7B, 0x8B, 0xFF}, true}
	Tomato               = X11Color{"Tomato", color.RGBA{0xFF, 0x63, 0x47, 0xFF}, false}
	Tomato1              = X11Color{"tomato1", color.RGBA{0xFF, 0x63, 0x47, 0xFF}, true}
	Tomato2              = X11Color{"tomato2", color.RGBA{0xEE, 0x5C, 0x42, 0xFF}, true}
	Tomato3              = X11Color{"tomato3", color.RGBA{0xCD, 0x4F, 0x39, 0xFF}, true}
	Tomato4              = X11Color{"tomato4", color.RGBA{0x8B, 0x36, 0x26, 0xFF}, false}
	Turquoise            = X11Color{"Turquoise", color.RGBA{0x40, 0xE0, 0xD0, 0xFF}, false}
	Turquoise1           = X11Color{"turquoise1", color.RGBA{0x00, 0xF5, 0xFF, 0xFF}, true}
	Turquoise2           = X11Color{"turquoise2", color.RGBA{0x00, 0xE5, 0xEE, 0xFF}, true}
	Turquoise3           = X11Color{"turquoise3", color.RGBA{0x00, 0xC5, 0xCD, 0xFF}, true}
	Turquoise4           = X11Color{"turquoise4", color.RGBA{0x00, 0x86, 0x8B, 0xFF}, true}
	Violet               = X11Color{"Violet", color.RGBA{0xEE, 0x82, 0xEE, 0xFF}, false}
	VioletRed            = X11Color{"violet red", color.RGBA{0xD0, 0x20, 0x90, 0xFF}, false}
	Violetred            = X11Color{"VioletRed", color.RGBA{0xD0, 0x20, 0x90, 0xFF}, false}
	Violetred1           = X11Color{"VioletRed1", color.RGBA{0xFF, 0x3E, 0x96, 0xFF}, true}
	Violetred2           = X11Color{"VioletRed2", color.RGBA{0xEE, 0x3A, 0x8C, 0xFF}, true}
	Violetred3           = X11Color{"VioletRed3", color.RGBA{0xCD, 0x32, 0x78, 0xFF}, false}
	Violetred4           = X11Color{"VioletRed4", color.RGBA{0x8B, 0x22, 0x52, 0xFF}, false}
	WebGray              = X11Color{"web gray", color.RGBA{0x80, 0x80, 0x80, 0xFF}, true}
	WebGreen             = X11Color{"web green", color.RGBA{0x00, 0x80, 0x00, 0xFF}, false}
	WebGrey              = X11Color{"web grey", color.RGBA{0x80, 0x80, 0x80, 0xFF}, true}
	WebMaroon            = X11Color{"web maroon", color.RGBA{0x80, 0x00, 0x00, 0xFF}, false}
	WebPurple            = X11Color{"web purple", color.RGBA{0x80, 0x00, 0x80, 0xFF}, false}
	Webgray              = X11Color{"WebGray", color.RGBA{0x80, 0x80, 0x80, 0xFF}, true}
	Webgreen             = X11Color{"WebGreen", color.RGBA{0x00, 0x80, 0x00, 0xFF}, false}
	Webgrey              = X11Color{"WebGrey", color.RGBA{0x80, 0x80, 0x80, 0xFF}, true}
	Webmaroon            = X11Color{"WebMaroon", color.RGBA{0x80, 0x00, 0x00, 0xFF}, false}
	Webpurple            = X11Color{"WebPurple", color.RGBA{0x80, 0x00, 0x80, 0xFF}, false}
	Wheat                = X11Color{"Wheat", color.RGBA{0xF5, 0xDE, 0xB3, 0xFF}, false}
	Wheat1               = X11Color{"wheat1", color.RGBA{0xFF, 0xE7, 0xBA, 0xFF}, true}
	Wheat2               = X11Color{"wheat2", color.RGBA{0xEE, 0xD8, 0xAE, 0xFF}, true}
	Wheat3               = X11Color{"wheat3", color.RGBA{0xCD, 0xBA, 0x96, 0xFF}, true}
	Wheat4               = X11Color{"wheat4", color.RGBA{0x8B, 0x7E, 0x66, 0xFF}, true}
	White                = X11Color{"White", color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, true}
	WhiteSmoke           = X11Color{"White Smoke", color.RGBA{0xF5, 0xF5, 0xF5, 0xFF}, true}
	Whitesmoke           = X11Color{"WhiteSmoke", color.RGBA{0xF5, 0xF5, 0xF5, 0xFF}, true}
	X11Gray              = X11Color{"x11 gray", color.RGBA{0xBE, 0xBE, 0xBE, 0xFF}, true}
	X11Green             = X11Color{"x11 green", color.RGBA{0x00, 0xFF, 0x00, 0xFF}, true}
	X11Grey              = X11Color{"x11 grey", color.RGBA{0xBE, 0xBE, 0xBE, 0xFF}, true}
	X11Maroon            = X11Color{"x11 maroon", color.RGBA{0xB0, 0x30, 0x60, 0xFF}, false}
	X11Purple            = X11Color{"x11 purple", color.RGBA{0xA0, 0x20, 0xF0, 0xFF}, false}
	X11gray              = X11Color{"X11Gray", color.RGBA{0xBE, 0xBE, 0xBE, 0xFF}, true}
	X11green             = X11Color{"X11Green", color.RGBA{0x00, 0xFF, 0x00, 0xFF}, true}
	X11grey              = X11Color{"X11Grey", color.RGBA{0xBE, 0xBE, 0xBE, 0xFF}, true}
	X11maroon            = X11Color{"X11Maroon", color.RGBA{0xB0, 0x30, 0x60, 0xFF}, false}
	X11purple            = X11Color{"X11Purple", color.RGBA{0xA0, 0x20, 0xF0, 0xFF}, false}
	Yellow               = X11Color{"Yellow", color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, true}
	Yellow1              = X11Color{"yellow1", color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, true}
	Yellow2              = X11Color{"yellow2", color.RGBA{0xEE, 0xEE, 0x00, 0xFF}, true}
	Yellow3              = X11Color{"yellow3", color.RGBA{0xCD, 0xCD, 0x00, 0xFF}, true}
	Yellow4              = X11Color{"yellow4", color.RGBA{0x8B, 0x8B, 0x00, 0xFF}, true}
	YellowGreen          = X11Color{"Yellow Green", color.RGBA{0x9A, 0xCD, 0x32, 0xFF}, false}
	Yellowgreen          = X11Color{"YellowGreen", color.RGBA{0x9A, 0xCD, 0x32, 0xFF}, true}
)

var (
	colors = []X11Color{
		AliceBlue,
		Aliceblue,
		AntiqueWhite,
		Antiquewhite,
		Antiquewhite1,
		Antiquewhite2,
		Antiquewhite3,
		Antiquewhite4,
		Aqua,
		Aquamarine,
		Aquamarine1,
		Aquamarine2,
		Aquamarine3,
		Aquamarine4,
		Azure,
		Azure1,
		Azure2,
		Azure3,
		Azure4,
		Beige,
		Bisque,
		Bisque1,
		Bisque2,
		Bisque3,
		Bisque4,
		Black,
		BlanchedAlmond,
		Blanchedalmond,
		Blue,
		Blue1,
		Blue2,
		Blue3,
		Blue4,
		BlueViolet,
		Blueviolet,
		Brown,
		Brown1,
		Brown2,
		Brown3,
		Brown4,
		Burlywood,
		Burlywood1,
		Burlywood2,
		Burlywood3,
		Burlywood4,
		CadetBlue,
		Cadetblue,
		Cadetblue1,
		Cadetblue2,
		Cadetblue3,
		Cadetblue4,
		Chartreuse,
		Chartreuse1,
		Chartreuse2,
		Chartreuse3,
		Chartreuse4,
		Chocolate,
		Chocolate1,
		Chocolate2,
		Chocolate3,
		Chocolate4,
		Coral,
		Coral1,
		Coral2,
		Coral3,
		Coral4,
		Cornflower,
		CornflowerBlue,
		Cornflowerblue,
		Cornsilk,
		Cornsilk1,
		Cornsilk2,
		Cornsilk3,
		Cornsilk4,
		Crimson,
		Cyan,
		Cyan1,
		Cyan2,
		Cyan3,
		Cyan4,
		DarkBlue,
		DarkCyan,
		DarkGoldenrod,
		DarkGray,
		DarkGreen,
		DarkGrey,
		DarkKhaki,
		DarkMagenta,
		DarkOliveGreen,
		DarkOrange,
		DarkOrchid,
		DarkRed,
		DarkSalmon,
		DarkSeaGreen,
		DarkSlateBlue,
		DarkSlateGray,
		DarkSlateGrey,
		DarkTurquoise,
		DarkViolet,
		Darkblue,
		Darkcyan,
		Darkgoldenrod,
		Darkgoldenrod1,
		Darkgoldenrod2,
		Darkgoldenrod3,
		Darkgoldenrod4,
		Darkgray,
		Darkgreen,
		Darkgrey,
		Darkkhaki,
		Darkmagenta,
		Darkolivegreen,
		Darkolivegreen1,
		Darkolivegreen2,
		Darkolivegreen3,
		Darkolivegreen4,
		Darkorange,
		Darkorange1,
		Darkorange2,
		Darkorange3,
		Darkorange4,
		Darkorchid,
		Darkorchid1,
		Darkorchid2,
		Darkorchid3,
		Darkorchid4,
		Darkred,
		Darksalmon,
		Darkseagreen,
		Darkseagreen1,
		Darkseagreen2,
		Darkseagreen3,
		Darkseagreen4,
		Darkslateblue,
		Darkslategray,
		Darkslategray1,
		Darkslategray2,
		Darkslategray3,
		Darkslategray4,
		Darkslategrey,
		Darkturquoise,
		Darkviolet,
		DeepPink,
		DeepSkyBlue,
		Deeppink,
		Deeppink1,
		Deeppink2,
		Deeppink3,
		Deeppink4,
		Deepskyblue,
		Deepskyblue1,
		Deepskyblue2,
		Deepskyblue3,
		Deepskyblue4,
		DimGray,
		DimGrey,
		Dimgray,
		Dimgrey,
		DodgerBlue,
		Dodgerblue,
		Dodgerblue1,
		Dodgerblue2,
		Dodgerblue3,
		Dodgerblue4,
		Firebrick,
		Firebrick1,
		Firebrick2,
		Firebrick3,
		Firebrick4,
		FloralWhite,
		Floralwhite,
		ForestGreen,
		Forestgreen,
		Fuchsia,
		Gainsboro,
		GhostWhite,
		Ghostwhite,
		Gold,
		Gold1,
		Gold2,
		Gold3,
		Gold4,
		Goldenrod,
		Goldenrod1,
		Goldenrod2,
		Goldenrod3,
		Goldenrod4,
		Gray,
		Gray0,
		Gray1,
		Gray10,
		Gray100,
		Gray11,
		Gray12,
		Gray13,
		Gray14,
		Gray15,
		Gray16,
		Gray17,
		Gray18,
		Gray19,
		Gray2,
		Gray20,
		Gray21,
		Gray22,
		Gray23,
		Gray24,
		Gray25,
		Gray26,
		Gray27,
		Gray28,
		Gray29,
		Gray3,
		Gray30,
		Gray31,
		Gray32,
		Gray33,
		Gray34,
		Gray35,
		Gray36,
		Gray37,
		Gray38,
		Gray39,
		Gray4,
		Gray40,
		Gray41,
		Gray42,
		Gray43,
		Gray44,
		Gray45,
		Gray46,
		Gray47,
		Gray48,
		Gray49,
		Gray5,
		Gray50,
		Gray51,
		Gray52,
		Gray53,
		Gray54,
		Gray55,
		Gray56,
		Gray57,
		Gray58,
		Gray59,
		Gray6,
		Gray60,
		Gray61,
		Gray62,
		Gray63,
		Gray64,
		Gray65,
		Gray66,
		Gray67,
		Gray68,
		Gray69,
		Gray7,
		Gray70,
		Gray71,
		Gray72,
		Gray73,
		Gray74,
		Gray75,
		Gray76,
		Gray77,
		Gray78,
		Gray79,
		Gray8,
		Gray80,
		Gray81,
		Gray82,
		Gray83,
		Gray84,
		Gray85,
		Gray86,
		Gray87,
		Gray88,
		Gray89,
		Gray9,
		Gray90,
		Gray91,
		Gray92,
		Gray93,
		Gray94,
		Gray95,
		Gray96,
		Gray97,
		Gray98,
		Gray99,
		GrayW3C,
		GrayX11,
		Green,
		Green1,
		Green2,
		Green3,
		Green4,
		GreenW3C,
		GreenX11,
		GreenYellow,
		Greenyellow,
		Grey,
		Grey0,
		Grey1,
		Grey10,
		Grey100,
		Grey11,
		Grey12,
		Grey13,
		Grey14,
		Grey15,
		Grey16,
		Grey17,
		Grey18,
		Grey19,
		Grey2,
		Grey20,
		Grey21,
		Grey22,
		Grey23,
		Grey24,
		Grey25,
		Grey26,
		Grey27,
		Grey28,
		Grey29,
		Grey3,
		Grey30,
		Grey31,
		Grey32,
		Grey33,
		Grey34,
		Grey35,
		Grey36,
		Grey37,
		Grey38,
		Grey39,
		Grey4,
		Grey40,
		Grey41,
		Grey42,
		Grey43,
		Grey44,
		Grey45,
		Grey46,
		Grey47,
		Grey48,
		Grey49,
		Grey5,
		Grey50,
		Grey51,
		Grey52,
		Grey53,
		Grey54,
		Grey55,
		Grey56,
		Grey57,
		Grey58,
		Grey59,
		Grey6,
		Grey60,
		Grey61,
		Grey62,
		Grey63,
		Grey64,
		Grey65,
		Grey66,
		Grey67,
		Grey68,
		Grey69,
		Grey7,
		Grey70,
		Grey71,
		Grey72,
		Grey73,
		Grey74,
		Grey75,
		Grey76,
		Grey77,
		Grey78,
		Grey79,
		Grey8,
		Grey80,
		Grey81,
		Grey82,
		Grey83,
		Grey84,
		Grey85,
		Grey86,
		Grey87,
		Grey88,
		Grey89,
		Grey9,
		Grey90,
		Grey91,
		Grey92,
		Grey93,
		Grey94,
		Grey95,
		Grey96,
		Grey97,
		Grey98,
		Grey99,
		Honeydew,
		Honeydew1,
		Honeydew2,
		Honeydew3,
		Honeydew4,
		HotPink,
		Hotpink,
		Hotpink1,
		Hotpink2,
		Hotpink3,
		Hotpink4,
		IndianRed,
		Indianred,
		Indianred1,
		Indianred2,
		Indianred3,
		Indianred4,
		Indigo,
		Ivory,
		Ivory1,
		Ivory2,
		Ivory3,
		Ivory4,
		Khaki,
		Khaki1,
		Khaki2,
		Khaki3,
		Khaki4,
		Lavender,
		LavenderBlush,
		Lavenderblush,
		Lavenderblush1,
		Lavenderblush2,
		Lavenderblush3,
		Lavenderblush4,
		LawnGreen,
		Lawngreen,
		LemonChiffon,
		Lemonchiffon,
		Lemonchiffon1,
		Lemonchiffon2,
		Lemonchiffon3,
		Lemonchiffon4,
		LightBlue,
		LightCoral,
		LightCyan,
		LightGoldenrod,
		LightGoldenrodYellow,
		LightGray,
		LightGreen,
		LightGrey,
		LightPink,
		LightSalmon,
		LightSeaGreen,
		LightSkyBlue,
		LightSlateBlue,
		LightSlateGray,
		LightSlateGrey,
		LightSteelBlue,
		LightYellow,
		Lightblue,
		Lightblue1,
		Lightblue2,
		Lightblue3,
		Lightblue4,
		Lightcoral,
		Lightcyan,
		Lightcyan1,
		Lightcyan2,
		Lightcyan3,
		Lightcyan4,
		Lightgoldenrod,
		Lightgoldenrod1,
		Lightgoldenrod2,
		Lightgoldenrod3,
		Lightgoldenrod4,
		Lightgoldenrodyellow,
		Lightgray,
		Lightgreen,
		Lightgrey,
		Lightpink,
		Lightpink1,
		Lightpink2,
		Lightpink3,
		Lightpink4,
		Lightsalmon,
		Lightsalmon1,
		Lightsalmon2,
		Lightsalmon3,
		Lightsalmon4,
		Lightseagreen,
		Lightskyblue,
		Lightskyblue1,
		Lightskyblue2,
		Lightskyblue3,
		Lightskyblue4,
		Lightslateblue,
		Lightslategray,
		Lightslategrey,
		Lightsteelblue,
		Lightsteelblue1,
		Lightsteelblue2,
		Lightsteelblue3,
		Lightsteelblue4,
		Lightyellow,
		Lightyellow1,
		Lightyellow2,
		Lightyellow3,
		Lightyellow4,
		Lime,
		LimeGreen,
		Limegreen,
		Linen,
		Magenta,
		Magenta1,
		Magenta2,
		Magenta3,
		Magenta4,
		Maroon,
		Maroon1,
		Maroon2,
		Maroon3,
		Maroon4,
		MaroonW3C,
		MaroonX11,
		MediumAquamarine,
		MediumBlue,
		MediumOrchid,
		MediumPurple,
		MediumSeaGreen,
		MediumSlateBlue,
		MediumSpringGreen,
		MediumTurquoise,
		MediumVioletRed,
		Mediumaquamarine,
		Mediumblue,
		Mediumorchid,
		Mediumorchid1,
		Mediumorchid2,
		Mediumorchid3,
		Mediumorchid4,
		Mediumpurple,
		Mediumpurple1,
		Mediumpurple2,
		Mediumpurple3,
		Mediumpurple4,
		Mediumseagreen,
		Mediumslateblue,
		Mediumspringgreen,
		Mediumturquoise,
		Mediumvioletred,
		MidnightBlue,
		Midnightblue,
		MintCream,
		Mintcream,
		MistyRose,
		Mistyrose,
		Mistyrose1,
		Mistyrose2,
		Mistyrose3,
		Mistyrose4,
		Moccasin,
		NavajoWhite,
		Navajowhite,
		Navajowhite1,
		Navajowhite2,
		Navajowhite3,
		Navajowhite4,
		Navy,
		NavyBlue,
		Navyblue,
		OldLace,
		Oldlace,
		Olive,
		OliveDrab,
		Olivedrab,
		Olivedrab1,
		Olivedrab2,
		Olivedrab3,
		Olivedrab4,
		Orange,
		Orange1,
		Orange2,
		Orange3,
		Orange4,
		OrangeRed,
		Orangered,
		Orangered1,
		Orangered2,
		Orangered3,
		Orangered4,
		Orchid,
		Orchid1,
		Orchid2,
		Orchid3,
		Orchid4,
		PaleGoldenrod,
		PaleGreen,
		PaleTurquoise,
		PaleVioletRed,
		Palegoldenrod,
		Palegreen,
		Palegreen1,
		Palegreen2,
		Palegreen3,
		Palegreen4,
		Paleturquoise,
		Paleturquoise1,
		Paleturquoise2,
		Paleturquoise3,
		Paleturquoise4,
		Palevioletred,
		Palevioletred1,
		Palevioletred2,
		Palevioletred3,
		Palevioletred4,
		PapayaWhip,
		Papayawhip,
		PeachPuff,
		Peachpuff,
		Peachpuff1,
		Peachpuff2,
		Peachpuff3,
		Peachpuff4,
		Peru,
		Pink,
		Pink1,
		Pink2,
		Pink3,
		Pink4,
		Plum,
		Plum1,
		Plum2,
		Plum3,
		Plum4,
		PowderBlue,
		Powderblue,
		Purple,
		Purple1,
		Purple2,
		Purple3,
		Purple4,
		PurpleW3C,
		PurpleX11,
		RebeccaPurple,
		Rebeccapurple,
		Red,
		Red1,
		Red2,
		Red3,
		Red4,
		RosyBrown,
		Rosybrown,
		Rosybrown1,
		Rosybrown2,
		Rosybrown3,
		Rosybrown4,
		RoyalBlue,
		Royalblue,
		Royalblue1,
		Royalblue2,
		Royalblue3,
		Royalblue4,
		SaddleBrown,
		Saddlebrown,
		Salmon,
		Salmon1,
		Salmon2,
		Salmon3,
		Salmon4,
		SandyBrown,
		Sandybrown,
		SeaGreen,
		Seagreen,
		Seagreen1,
		Seagreen2,
		Seagreen3,
		Seagreen4,
		Seashell,
		Seashell1,
		Seashell2,
		Seashell3,
		Seashell4,
		Sienna,
		Sienna1,
		Sienna2,
		Sienna3,
		Sienna4,
		Silver,
		SkyBlue,
		Skyblue,
		Skyblue1,
		Skyblue2,
		Skyblue3,
		Skyblue4,
		SlateBlue,
		SlateGray,
		SlateGrey,
		Slateblue,
		Slateblue1,
		Slateblue2,
		Slateblue3,
		Slateblue4,
		Slategray,
		Slategray1,
		Slategray2,
		Slategray3,
		Slategray4,
		Slategrey,
		Snow,
		Snow1,
		Snow2,
		Snow3,
		Snow4,
		SpringGreen,
		Springgreen,
		Springgreen1,
		Springgreen2,
		Springgreen3,
		Springgreen4,
		SteelBlue,
		Steelblue,
		Steelblue1,
		Steelblue2,
		Steelblue3,
		Steelblue4,
		Tan,
		Tan1,
		Tan2,
		Tan3,
		Tan4,
		Teal,
		Thistle,
		Thistle1,
		Thistle2,
		Thistle3,
		Thistle4,
		Tomato,
		Tomato1,
		Tomato2,
		Tomato3,
		Tomato4,
		Turquoise,
		Turquoise1,
		Turquoise2,
		Turquoise3,
		Turquoise4,
		Violet,
		VioletRed,
		Violetred,
		Violetred1,
		Violetred2,
		Violetred3,
		Violetred4,
		WebGray,
		WebGreen,
		WebGrey,
		WebMaroon,
		WebPurple,
		Webgray,
		Webgreen,
		Webgrey,
		Webmaroon,
		Webpurple,
		Wheat,
		Wheat1,
		Wheat2,
		Wheat3,
		Wheat4,
		White,
		WhiteSmoke,
		Whitesmoke,
		X11Gray,
		X11Green,
		X11Grey,
		X11Maroon,
		X11Purple,
		X11gray,
		X11green,
		X11grey,
		X11maroon,
		X11purple,
		Yellow,
		Yellow1,
		Yellow2,
		Yellow3,
		Yellow4,
		YellowGreen,
		Yellowgreen,
	}

	names = map[string]X11Color{
		"Alice Blue":             AliceBlue,
		"AliceBlue":              Aliceblue,
		"Antique White":          AntiqueWhite,
		"AntiqueWhite":           Antiquewhite,
		"AntiqueWhite1":          Antiquewhite1,
		"AntiqueWhite2":          Antiquewhite2,
		"AntiqueWhite3":          Antiquewhite3,
		"AntiqueWhite4":          Antiquewhite4,
		"Aqua":                   Aqua,
		"Aquamarine":             Aquamarine,
		"aquamarine1":            Aquamarine1,
		"aquamarine2":            Aquamarine2,
		"aquamarine3":            Aquamarine3,
		"aquamarine4":            Aquamarine4,
		"Azure":                  Azure,
		"azure1":                 Azure1,
		"azure2":                 Azure2,
		"azure3":                 Azure3,
		"azure4":                 Azure4,
		"Beige":                  Beige,
		"Bisque":                 Bisque,
		"bisque1":                Bisque1,
		"bisque2":                Bisque2,
		"bisque3":                Bisque3,
		"bisque4":                Bisque4,
		"Black":                  Black,
		"Blanched Almond":        BlanchedAlmond,
		"BlanchedAlmond":         Blanchedalmond,
		"Blue":                   Blue,
		"blue1":                  Blue1,
		"blue2":                  Blue2,
		"blue3":                  Blue3,
		"blue4":                  Blue4,
		"Blue Violet":            BlueViolet,
		"BlueViolet":             Blueviolet,
		"Brown":                  Brown,
		"brown1":                 Brown1,
		"brown2":                 Brown2,
		"brown3":                 Brown3,
		"brown4":                 Brown4,
		"Burlywood":              Burlywood,
		"burlywood1":             Burlywood1,
		"burlywood2":             Burlywood2,
		"burlywood3":             Burlywood3,
		"burlywood4":             Burlywood4,
		"Cadet Blue":             CadetBlue,
		"CadetBlue":              Cadetblue,
		"CadetBlue1":             Cadetblue1,
		"CadetBlue2":             Cadetblue2,
		"CadetBlue3":             Cadetblue3,
		"CadetBlue4":             Cadetblue4,
		"Chartreuse":             Chartreuse,
		"chartreuse1":            Chartreuse1,
		"chartreuse2":            Chartreuse2,
		"chartreuse3":            Chartreuse3,
		"chartreuse4":            Chartreuse4,
		"Chocolate":              Chocolate,
		"chocolate1":             Chocolate1,
		"chocolate2":             Chocolate2,
		"chocolate3":             Chocolate3,
		"chocolate4":             Chocolate4,
		"Coral":                  Coral,
		"coral1":                 Coral1,
		"coral2":                 Coral2,
		"coral3":                 Coral3,
		"coral4":                 Coral4,
		"Cornflower":             Cornflower,
		"cornflower blue":        CornflowerBlue,
		"CornflowerBlue":         Cornflowerblue,
		"Cornsilk":               Cornsilk,
		"cornsilk1":              Cornsilk1,
		"cornsilk2":              Cornsilk2,
		"cornsilk3":              Cornsilk3,
		"cornsilk4":              Cornsilk4,
		"Crimson":                Crimson,
		"Cyan":                   Cyan,
		"cyan1":                  Cyan1,
		"cyan2":                  Cyan2,
		"cyan3":                  Cyan3,
		"cyan4":                  Cyan4,
		"Dark Blue":              DarkBlue,
		"Dark Cyan":              DarkCyan,
		"Dark Goldenrod":         DarkGoldenrod,
		"Dark Gray":              DarkGray,
		"Dark Green":             DarkGreen,
		"dark grey":              DarkGrey,
		"Dark Khaki":             DarkKhaki,
		"Dark Magenta":           DarkMagenta,
		"Dark Olive Green":       DarkOliveGreen,
		"Dark Orange":            DarkOrange,
		"Dark Orchid":            DarkOrchid,
		"Dark Red":               DarkRed,
		"Dark Salmon":            DarkSalmon,
		"Dark Sea Green":         DarkSeaGreen,
		"Dark Slate Blue":        DarkSlateBlue,
		"Dark Slate Gray":        DarkSlateGray,
		"dark slate grey":        DarkSlateGrey,
		"Dark Turquoise":         DarkTurquoise,
		"Dark Violet":            DarkViolet,
		"DarkBlue":               Darkblue,
		"DarkCyan":               Darkcyan,
		"DarkGoldenrod":          Darkgoldenrod,
		"DarkGoldenrod1":         Darkgoldenrod1,
		"DarkGoldenrod2":         Darkgoldenrod2,
		"DarkGoldenrod3":         Darkgoldenrod3,
		"DarkGoldenrod4":         Darkgoldenrod4,
		"DarkGray":               Darkgray,
		"DarkGreen":              Darkgreen,
		"DarkGrey":               Darkgrey,
		"DarkKhaki":              Darkkhaki,
		"DarkMagenta":            Darkmagenta,
		"DarkOliveGreen":         Darkolivegreen,
		"DarkOliveGreen1":        Darkolivegreen1,
		"DarkOliveGreen2":        Darkolivegreen2,
		"DarkOliveGreen3":        Darkolivegreen3,
		"DarkOliveGreen4":        Darkolivegreen4,
		"DarkOrange":             Darkorange,
		"DarkOrange1":            Darkorange1,
		"DarkOrange2":            Darkorange2,
		"DarkOrange3":            Darkorange3,
		"DarkOrange4":            Darkorange4,
		"DarkOrchid":             Darkorchid,
		"DarkOrchid1":            Darkorchid1,
		"DarkOrchid2":            Darkorchid2,
		"DarkOrchid3":            Darkorchid3,
		"DarkOrchid4":            Darkorchid4,
		"DarkRed":                Darkred,
		"DarkSalmon":             Darksalmon,
		"DarkSeaGreen":           Darkseagreen,
		"DarkSeaGreen1":          Darkseagreen1,
		"DarkSeaGreen2":          Darkseagreen2,
		"DarkSeaGreen3":          Darkseagreen3,
		"DarkSeaGreen4":          Darkseagreen4,
		"DarkSlateBlue":          Darkslateblue,
		"DarkSlateGray":          Darkslategray,
		"DarkSlateGray1":         Darkslategray1,
		"DarkSlateGray2":         Darkslategray2,
		"DarkSlateGray3":         Darkslategray3,
		"DarkSlateGray4":         Darkslategray4,
		"DarkSlateGrey":          Darkslategrey,
		"DarkTurquoise":          Darkturquoise,
		"DarkViolet":             Darkviolet,
		"Deep Pink":              DeepPink,
		"Deep Sky Blue":          DeepSkyBlue,
		"DeepPink":               Deeppink,
		"DeepPink1":              Deeppink1,
		"DeepPink2":              Deeppink2,
		"DeepPink3":              Deeppink3,
		"DeepPink4":              Deeppink4,
		"DeepSkyBlue":            Deepskyblue,
		"DeepSkyBlue1":           Deepskyblue1,
		"DeepSkyBlue2":           Deepskyblue2,
		"DeepSkyBlue3":           Deepskyblue3,
		"DeepSkyBlue4":           Deepskyblue4,
		"Dim Gray":               DimGray,
		"dim grey":               DimGrey,
		"DimGray":                Dimgray,
		"DimGrey":                Dimgrey,
		"Dodger Blue":            DodgerBlue,
		"DodgerBlue":             Dodgerblue,
		"DodgerBlue1":            Dodgerblue1,
		"DodgerBlue2":            Dodgerblue2,
		"DodgerBlue3":            Dodgerblue3,
		"DodgerBlue4":            Dodgerblue4,
		"Firebrick":              Firebrick,
		"firebrick1":             Firebrick1,
		"firebrick2":             Firebrick2,
		"firebrick3":             Firebrick3,
		"firebrick4":             Firebrick4,
		"Floral White":           FloralWhite,
		"FloralWhite":            Floralwhite,
		"Forest Green":           ForestGreen,
		"ForestGreen":            Forestgreen,
		"Fuchsia":                Fuchsia,
		"Gainsboro":              Gainsboro,
		"Ghost White":            GhostWhite,
		"GhostWhite":             Ghostwhite,
		"Gold":                   Gold,
		"gold1":                  Gold1,
		"gold2":                  Gold2,
		"gold3":                  Gold3,
		"gold4":                  Gold4,
		"Goldenrod":              Goldenrod,
		"goldenrod1":             Goldenrod1,
		"goldenrod2":             Goldenrod2,
		"goldenrod3":             Goldenrod3,
		"goldenrod4":             Goldenrod4,
		"gray":                   Gray,
		"gray0":                  Gray0,
		"gray1":                  Gray1,
		"gray10":                 Gray10,
		"gray100":                Gray100,
		"gray11":                 Gray11,
		"gray12":                 Gray12,
		"gray13":                 Gray13,
		"gray14":                 Gray14,
		"gray15":                 Gray15,
		"gray16":                 Gray16,
		"gray17":                 Gray17,
		"gray18":                 Gray18,
		"gray19":                 Gray19,
		"gray2":                  Gray2,
		"gray20":                 Gray20,
		"gray21":                 Gray21,
		"gray22":                 Gray22,
		"gray23":                 Gray23,
		"gray24":                 Gray24,
		"gray25":                 Gray25,
		"gray26":                 Gray26,
		"gray27":                 Gray27,
		"gray28":                 Gray28,
		"gray29":                 Gray29,
		"gray3":                  Gray3,
		"gray30":                 Gray30,
		"gray31":                 Gray31,
		"gray32":                 Gray32,
		"gray33":                 Gray33,
		"gray34":                 Gray34,
		"gray35":                 Gray35,
		"gray36":                 Gray36,
		"gray37":                 Gray37,
		"gray38":                 Gray38,
		"gray39":                 Gray39,
		"gray4":                  Gray4,
		"gray40":                 Gray40,
		"gray41":                 Gray41,
		"gray42":                 Gray42,
		"gray43":                 Gray43,
		"gray44":                 Gray44,
		"gray45":                 Gray45,
		"gray46":                 Gray46,
		"gray47":                 Gray47,
		"gray48":                 Gray48,
		"gray49":                 Gray49,
		"gray5":                  Gray5,
		"gray50":                 Gray50,
		"gray51":                 Gray51,
		"gray52":                 Gray52,
		"gray53":                 Gray53,
		"gray54":                 Gray54,
		"gray55":                 Gray55,
		"gray56":                 Gray56,
		"gray57":                 Gray57,
		"gray58":                 Gray58,
		"gray59":                 Gray59,
		"gray6":                  Gray6,
		"gray60":                 Gray60,
		"gray61":                 Gray61,
		"gray62":                 Gray62,
		"gray63":                 Gray63,
		"gray64":                 Gray64,
		"gray65":                 Gray65,
		"gray66":                 Gray66,
		"gray67":                 Gray67,
		"gray68":                 Gray68,
		"gray69":                 Gray69,
		"gray7":                  Gray7,
		"gray70":                 Gray70,
		"gray71":                 Gray71,
		"gray72":                 Gray72,
		"gray73":                 Gray73,
		"gray74":                 Gray74,
		"gray75":                 Gray75,
		"gray76":                 Gray76,
		"gray77":                 Gray77,
		"gray78":                 Gray78,
		"gray79":                 Gray79,
		"gray8":                  Gray8,
		"gray80":                 Gray80,
		"gray81":                 Gray81,
		"gray82":                 Gray82,
		"gray83":                 Gray83,
		"gray84":                 Gray84,
		"gray85":                 Gray85,
		"gray86":                 Gray86,
		"gray87":                 Gray87,
		"gray88":                 Gray88,
		"gray89":                 Gray89,
		"gray9":                  Gray9,
		"gray90":                 Gray90,
		"gray91":                 Gray91,
		"gray92":                 Gray92,
		"gray93":                 Gray93,
		"gray94":                 Gray94,
		"gray95":                 Gray95,
		"gray96":                 Gray96,
		"gray97":                 Gray97,
		"gray98":                 Gray98,
		"gray99":                 Gray99,
		"Gray (W3C)":             GrayW3C,
		"Gray (X11)":             GrayX11,
		"green":                  Green,
		"green1":                 Green1,
		"green2":                 Green2,
		"green3":                 Green3,
		"green4":                 Green4,
		"Green (W3C)":            GreenW3C,
		"Green (X11)":            GreenX11,
		"Green Yellow":           GreenYellow,
		"GreenYellow":            Greenyellow,
		"grey":                   Grey,
		"grey0":                  Grey0,
		"grey1":                  Grey1,
		"grey10":                 Grey10,
		"grey100":                Grey100,
		"grey11":                 Grey11,
		"grey12":                 Grey12,
		"grey13":                 Grey13,
		"grey14":                 Grey14,
		"grey15":                 Grey15,
		"grey16":                 Grey16,
		"grey17":                 Grey17,
		"grey18":                 Grey18,
		"grey19":                 Grey19,
		"grey2":                  Grey2,
		"grey20":                 Grey20,
		"grey21":                 Grey21,
		"grey22":                 Grey22,
		"grey23":                 Grey23,
		"grey24":                 Grey24,
		"grey25":                 Grey25,
		"grey26":                 Grey26,
		"grey27":                 Grey27,
		"grey28":                 Grey28,
		"grey29":                 Grey29,
		"grey3":                  Grey3,
		"grey30":                 Grey30,
		"grey31":                 Grey31,
		"grey32":                 Grey32,
		"grey33":                 Grey33,
		"grey34":                 Grey34,
		"grey35":                 Grey35,
		"grey36":                 Grey36,
		"grey37":                 Grey37,
		"grey38":                 Grey38,
		"grey39":                 Grey39,
		"grey4":                  Grey4,
		"grey40":                 Grey40,
		"grey41":                 Grey41,
		"grey42":                 Grey42,
		"grey43":                 Grey43,
		"grey44":                 Grey44,
		"grey45":                 Grey45,
		"grey46":                 Grey46,
		"grey47":                 Grey47,
		"grey48":                 Grey48,
		"grey49":                 Grey49,
		"grey5":                  Grey5,
		"grey50":                 Grey50,
		"grey51":                 Grey51,
		"grey52":                 Grey52,
		"grey53":                 Grey53,
		"grey54":                 Grey54,
		"grey55":                 Grey55,
		"grey56":                 Grey56,
		"grey57":                 Grey57,
		"grey58":                 Grey58,
		"grey59":                 Grey59,
		"grey6":                  Grey6,
		"grey60":                 Grey60,
		"grey61":                 Grey61,
		"grey62":                 Grey62,
		"grey63":                 Grey63,
		"grey64":                 Grey64,
		"grey65":                 Grey65,
		"grey66":                 Grey66,
		"grey67":                 Grey67,
		"grey68":                 Grey68,
		"grey69":                 Grey69,
		"grey7":                  Grey7,
		"grey70":                 Grey70,
		"grey71":                 Grey71,
		"grey72":                 Grey72,
		"grey73":                 Grey73,
		"grey74":                 Grey74,
		"grey75":                 Grey75,
		"grey76":                 Grey76,
		"grey77":                 Grey77,
		"grey78":                 Grey78,
		"grey79":                 Grey79,
		"grey8":                  Grey8,
		"grey80":                 Grey80,
		"grey81":                 Grey81,
		"grey82":                 Grey82,
		"grey83":                 Grey83,
		"grey84":                 Grey84,
		"grey85":                 Grey85,
		"grey86":                 Grey86,
		"grey87":                 Grey87,
		"grey88":                 Grey88,
		"grey89":                 Grey89,
		"grey9":                  Grey9,
		"grey90":                 Grey90,
		"grey91":                 Grey91,
		"grey92":                 Grey92,
		"grey93":                 Grey93,
		"grey94":                 Grey94,
		"grey95":                 Grey95,
		"grey96":                 Grey96,
		"grey97":                 Grey97,
		"grey98":                 Grey98,
		"grey99":                 Grey99,
		"Honeydew":               Honeydew,
		"honeydew1":              Honeydew1,
		"honeydew2":              Honeydew2,
		"honeydew3":              Honeydew3,
		"honeydew4":              Honeydew4,
		"Hot Pink":               HotPink,
		"HotPink":                Hotpink,
		"HotPink1":               Hotpink1,
		"HotPink2":               Hotpink2,
		"HotPink3":               Hotpink3,
		"HotPink4":               Hotpink4,
		"Indian Red":             IndianRed,
		"IndianRed":              Indianred,
		"IndianRed1":             Indianred1,
		"IndianRed2":             Indianred2,
		"IndianRed3":             Indianred3,
		"IndianRed4":             Indianred4,
		"Indigo":                 Indigo,
		"Ivory":                  Ivory,
		"ivory1":                 Ivory1,
		"ivory2":                 Ivory2,
		"ivory3":                 Ivory3,
		"ivory4":                 Ivory4,
		"Khaki":                  Khaki,
		"khaki1":                 Khaki1,
		"khaki2":                 Khaki2,
		"khaki3":                 Khaki3,
		"khaki4":                 Khaki4,
		"Lavender":               Lavender,
		"Lavender Blush":         LavenderBlush,
		"LavenderBlush":          Lavenderblush,
		"LavenderBlush1":         Lavenderblush1,
		"LavenderBlush2":         Lavenderblush2,
		"LavenderBlush3":         Lavenderblush3,
		"LavenderBlush4":         Lavenderblush4,
		"Lawn Green":             LawnGreen,
		"LawnGreen":              Lawngreen,
		"Lemon Chiffon":          LemonChiffon,
		"LemonChiffon":           Lemonchiffon,
		"LemonChiffon1":          Lemonchiffon1,
		"LemonChiffon2":          Lemonchiffon2,
		"LemonChiffon3":          Lemonchiffon3,
		"LemonChiffon4":          Lemonchiffon4,
		"Light Blue":             LightBlue,
		"Light Coral":            LightCoral,
		"Light Cyan":             LightCyan,
		"Light Goldenrod":        LightGoldenrod,
		"light goldenrod yellow": LightGoldenrodYellow,
		"Light Gray":             LightGray,
		"Light Green":            LightGreen,
		"light grey":             LightGrey,
		"Light Pink":             LightPink,
		"Light Salmon":           LightSalmon,
		"Light Sea Green":        LightSeaGreen,
		"Light Sky Blue":         LightSkyBlue,
		"light slate blue":       LightSlateBlue,
		"Light Slate Gray":       LightSlateGray,
		"light slate grey":       LightSlateGrey,
		"Light Steel Blue":       LightSteelBlue,
		"Light Yellow":           LightYellow,
		"LightBlue":              Lightblue,
		"LightBlue1":             Lightblue1,
		"LightBlue2":             Lightblue2,
		"LightBlue3":             Lightblue3,
		"LightBlue4":             Lightblue4,
		"LightCoral":             Lightcoral,
		"LightCyan":              Lightcyan,
		"LightCyan1":             Lightcyan1,
		"LightCyan2":             Lightcyan2,
		"LightCyan3":             Lightcyan3,
		"LightCyan4":             Lightcyan4,
		"LightGoldenrod":         Lightgoldenrod,
		"LightGoldenrod1":        Lightgoldenrod1,
		"LightGoldenrod2":        Lightgoldenrod2,
		"LightGoldenrod3":        Lightgoldenrod3,
		"LightGoldenrod4":        Lightgoldenrod4,
		"LightGoldenrodYellow":   Lightgoldenrodyellow,
		"LightGray":              Lightgray,
		"LightGreen":             Lightgreen,
		"LightGrey":              Lightgrey,
		"LightPink":              Lightpink,
		"LightPink1":             Lightpink1,
		"LightPink2":             Lightpink2,
		"LightPink3":             Lightpink3,
		"LightPink4":             Lightpink4,
		"LightSalmon":            Lightsalmon,
		"LightSalmon1":           Lightsalmon1,
		"LightSalmon2":           Lightsalmon2,
		"LightSalmon3":           Lightsalmon3,
		"LightSalmon4":           Lightsalmon4,
		"LightSeaGreen":          Lightseagreen,
		"LightSkyBlue":           Lightskyblue,
		"LightSkyBlue1":          Lightskyblue1,
		"LightSkyBlue2":          Lightskyblue2,
		"LightSkyBlue3":          Lightskyblue3,
		"LightSkyBlue4":          Lightskyblue4,
		"LightSlateBlue":         Lightslateblue,
		"LightSlateGray":         Lightslategray,
		"LightSlateGrey":         Lightslategrey,
		"LightSteelBlue":         Lightsteelblue,
		"LightSteelBlue1":        Lightsteelblue1,
		"LightSteelBlue2":        Lightsteelblue2,
		"LightSteelBlue3":        Lightsteelblue3,
		"LightSteelBlue4":        Lightsteelblue4,
		"LightYellow":            Lightyellow,
		"LightYellow1":           Lightyellow1,
		"LightYellow2":           Lightyellow2,
		"LightYellow3":           Lightyellow3,
		"LightYellow4":           Lightyellow4,
		"Lime":                   Lime,
		"Lime Green":             LimeGreen,
		"LimeGreen":              Limegreen,
		"Linen":                  Linen,
		"Magenta":                Magenta,
		"magenta1":               Magenta1,
		"magenta2":               Magenta2,
		"magenta3":               Magenta3,
		"magenta4":               Magenta4,
		"maroon":                 Maroon,
		"maroon1":                Maroon1,
		"maroon2":                Maroon2,
		"maroon3":                Maroon3,
		"maroon4":                Maroon4,
		"Maroon (W3C)":           MaroonW3C,
		"Maroon (X11)":           MaroonX11,
		"Medium Aquamarine":      MediumAquamarine,
		"Medium Blue":            MediumBlue,
		"Medium Orchid":          MediumOrchid,
		"Medium Purple":          MediumPurple,
		"Medium Sea Green":       MediumSeaGreen,
		"Medium Slate Blue":      MediumSlateBlue,
		"Medium Spring Green":    MediumSpringGreen,
		"Medium Turquoise":       MediumTurquoise,
		"Medium Violet Red":      MediumVioletRed,
		"MediumAquamarine":       Mediumaquamarine,
		"MediumBlue":             Mediumblue,
		"MediumOrchid":           Mediumorchid,
		"MediumOrchid1":          Mediumorchid1,
		"MediumOrchid2":          Mediumorchid2,
		"MediumOrchid3":          Mediumorchid3,
		"MediumOrchid4":          Mediumorchid4,
		"MediumPurple":           Mediumpurple,
		"MediumPurple1":          Mediumpurple1,
		"MediumPurple2":          Mediumpurple2,
		"MediumPurple3":          Mediumpurple3,
		"MediumPurple4":          Mediumpurple4,
		"MediumSeaGreen":         Mediumseagreen,
		"MediumSlateBlue":        Mediumslateblue,
		"MediumSpringGreen":      Mediumspringgreen,
		"MediumTurquoise":        Mediumturquoise,
		"MediumVioletRed":        Mediumvioletred,
		"Midnight Blue":          MidnightBlue,
		"MidnightBlue":           Midnightblue,
		"Mint Cream":             MintCream,
		"MintCream":              Mintcream,
		"Misty Rose":             MistyRose,
		"MistyRose":              Mistyrose,
		"MistyRose1":             Mistyrose1,
		"MistyRose2":             Mistyrose2,
		"MistyRose3":             Mistyrose3,
		"MistyRose4":             Mistyrose4,
		"Moccasin":               Moccasin,
		"Navajo White":           NavajoWhite,
		"NavajoWhite":            Navajowhite,
		"NavajoWhite1":           Navajowhite1,
		"NavajoWhite2":           Navajowhite2,
		"NavajoWhite3":           Navajowhite3,
		"NavajoWhite4":           Navajowhite4,
		"Navy":                   Navy,
		"navy blue":              NavyBlue,
		"NavyBlue":               Navyblue,
		"Old Lace":               OldLace,
		"OldLace":                Oldlace,
		"Olive":                  Olive,
		"Olive Drab":             OliveDrab,
		"OliveDrab":              Olivedrab,
		"OliveDrab1":             Olivedrab1,
		"OliveDrab2":             Olivedrab2,
		"OliveDrab3":             Olivedrab3,
		"OliveDrab4":             Olivedrab4,
		"Orange":                 Orange,
		"orange1":                Orange1,
		"orange2":                Orange2,
		"orange3":                Orange3,
		"orange4":                Orange4,
		"Orange Red":             OrangeRed,
		"OrangeRed":              Orangered,
		"OrangeRed1":             Orangered1,
		"OrangeRed2":             Orangered2,
		"OrangeRed3":             Orangered3,
		"OrangeRed4":             Orangered4,
		"Orchid":                 Orchid,
		"orchid1":                Orchid1,
		"orchid2":                Orchid2,
		"orchid3":                Orchid3,
		"orchid4":                Orchid4,
		"Pale Goldenrod":         PaleGoldenrod,
		"Pale Green":             PaleGreen,
		"Pale Turquoise":         PaleTurquoise,
		"Pale Violet Red":        PaleVioletRed,
		"PaleGoldenrod":          Palegoldenrod,
		"PaleGreen":              Palegreen,
		"PaleGreen1":             Palegreen1,
		"PaleGreen2":             Palegreen2,
		"PaleGreen3":             Palegreen3,
		"PaleGreen4":             Palegreen4,
		"PaleTurquoise":          Paleturquoise,
		"PaleTurquoise1":         Paleturquoise1,
		"PaleTurquoise2":         Paleturquoise2,
		"PaleTurquoise3":         Paleturquoise3,
		"PaleTurquoise4":         Paleturquoise4,
		"PaleVioletRed":          Palevioletred,
		"PaleVioletRed1":         Palevioletred1,
		"PaleVioletRed2":         Palevioletred2,
		"PaleVioletRed3":         Palevioletred3,
		"PaleVioletRed4":         Palevioletred4,
		"Papaya Whip":            PapayaWhip,
		"PapayaWhip":             Papayawhip,
		"Peach Puff":             PeachPuff,
		"PeachPuff":              Peachpuff,
		"PeachPuff1":             Peachpuff1,
		"PeachPuff2":             Peachpuff2,
		"PeachPuff3":             Peachpuff3,
		"PeachPuff4":             Peachpuff4,
		"Peru":                   Peru,
		"Pink":                   Pink,
		"pink1":                  Pink1,
		"pink2":                  Pink2,
		"pink3":                  Pink3,
		"pink4":                  Pink4,
		"Plum":                   Plum,
		"plum1":                  Plum1,
		"plum2":                  Plum2,
		"plum3":                  Plum3,
		"plum4":                  Plum4,
		"Powder Blue":            PowderBlue,
		"PowderBlue":             Powderblue,
		"purple":                 Purple,
		"purple1":                Purple1,
		"purple2":                Purple2,
		"purple3":                Purple3,
		"purple4":                Purple4,
		"Purple (W3C)":           PurpleW3C,
		"Purple (X11)":           PurpleX11,
		"rebecca purple":         RebeccaPurple,
		"RebeccaPurple":          Rebeccapurple,
		"Red":                    Red,
		"red1":                   Red1,
		"red2":                   Red2,
		"red3":                   Red3,
		"red4":                   Red4,
		"Rosy Brown":             RosyBrown,
		"RosyBrown":              Rosybrown,
		"RosyBrown1":             Rosybrown1,
		"RosyBrown2":             Rosybrown2,
		"RosyBrown3":             Rosybrown3,
		"RosyBrown4":             Rosybrown4,
		"Royal Blue":             RoyalBlue,
		"RoyalBlue":              Royalblue,
		"RoyalBlue1":             Royalblue1,
		"RoyalBlue2":             Royalblue2,
		"RoyalBlue3":             Royalblue3,
		"RoyalBlue4":             Royalblue4,
		"Saddle Brown":           SaddleBrown,
		"SaddleBrown":            Saddlebrown,
		"Salmon":                 Salmon,
		"salmon1":                Salmon1,
		"salmon2":                Salmon2,
		"salmon3":                Salmon3,
		"salmon4":                Salmon4,
		"Sandy Brown":            SandyBrown,
		"SandyBrown":             Sandybrown,
		"Sea Green":              SeaGreen,
		"SeaGreen":               Seagreen,
		"SeaGreen1":              Seagreen1,
		"SeaGreen2":              Seagreen2,
		"SeaGreen3":              Seagreen3,
		"SeaGreen4":              Seagreen4,
		"Seashell":               Seashell,
		"seashell1":              Seashell1,
		"seashell2":              Seashell2,
		"seashell3":              Seashell3,
		"seashell4":              Seashell4,
		"Sienna":                 Sienna,
		"sienna1":                Sienna1,
		"sienna2":                Sienna2,
		"sienna3":                Sienna3,
		"sienna4":                Sienna4,
		"Silver":                 Silver,
		"Sky Blue":               SkyBlue,
		"SkyBlue":                Skyblue,
		"SkyBlue1":               Skyblue1,
		"SkyBlue2":               Skyblue2,
		"SkyBlue3":               Skyblue3,
		"SkyBlue4":               Skyblue4,
		"Slate Blue":             SlateBlue,
		"Slate Gray":             SlateGray,
		"slate grey":             SlateGrey,
		"SlateBlue":              Slateblue,
		"SlateBlue1":             Slateblue1,
		"SlateBlue2":             Slateblue2,
		"SlateBlue3":             Slateblue3,
		"SlateBlue4":             Slateblue4,
		"SlateGray":              Slategray,
		"SlateGray1":             Slategray1,
		"SlateGray2":             Slategray2,
		"SlateGray3":             Slategray3,
		"SlateGray4":             Slategray4,
		"SlateGrey":              Slategrey,
		"Snow":                   Snow,
		"snow1":                  Snow1,
		"snow2":                  Snow2,
		"snow3":                  Snow3,
		"snow4":                  Snow4,
		"Spring Green":           SpringGreen,
		"SpringGreen":            Springgreen,
		"SpringGreen1":           Springgreen1,
		"SpringGreen2":           Springgreen2,
		"SpringGreen3":           Springgreen3,
		"SpringGreen4":           Springgreen4,
		"Steel Blue":             SteelBlue,
		"SteelBlue":              Steelblue,
		"SteelBlue1":             Steelblue1,
		"SteelBlue2":             Steelblue2,
		"SteelBlue3":             Steelblue3,
		"SteelBlue4":             Steelblue4,
		"Tan":                    Tan,
		"tan1":                   Tan1,
		"tan2":                   Tan2,
		"tan3":                   Tan3,
		"tan4":                   Tan4,
		"Teal":                   Teal,
		"Thistle":                Thistle,
		"thistle1":               Thistle1,
		"thistle2":               Thistle2,
		"thistle3":               Thistle3,
		"thistle4":               Thistle4,
		"Tomato":                 Tomato,
		"tomato1":                Tomato1,
		"tomato2":                Tomato2,
		"tomato3":                Tomato3,
		"tomato4":                Tomato4,
		"Turquoise":              Turquoise,
		"turquoise1":             Turquoise1,
		"turquoise2":             Turquoise2,
		"turquoise3":             Turquoise3,
		"turquoise4":             Turquoise4,
		"Violet":                 Violet,
		"violet red":             VioletRed,
		"VioletRed":              Violetred,
		"VioletRed1":             Violetred1,
		"VioletRed2":             Violetred2,
		"VioletRed3":             Violetred3,
		"VioletRed4":             Violetred4,
		"web gray":               WebGray,
		"web green":              WebGreen,
		"web grey":               WebGrey,
		"web maroon":             WebMaroon,
		"web purple":             WebPurple,
		"WebGray":                Webgray,
		"WebGreen":               Webgreen,
		"WebGrey":                Webgrey,
		"WebMaroon":              Webmaroon,
		"WebPurple":              Webpurple,
		"Wheat":                  Wheat,
		"wheat1":                 Wheat1,
		"wheat2":                 Wheat2,
		"wheat3":                 Wheat3,
		"wheat4":                 Wheat4,
		"White":                  White,
		"White Smoke":            WhiteSmoke,
		"WhiteSmoke":             Whitesmoke,
		"x11 gray":               X11Gray,
		"x11 green":              X11Green,
		"x11 grey":               X11Grey,
		"x11 maroon":             X11Maroon,
		"x11 purple":             X11Purple,
		"X11Gray":                X11gray,
		"X11Green":               X11green,
		"X11Grey":                X11grey,
		"X11Maroon":              X11maroon,
		"X11Purple":              X11purple,
		"Yellow":                 Yellow,
		"yellow1":                Yellow1,
		"yellow2":                Yellow2,
		"yellow3":                Yellow3,
		"yellow4":                Yellow4,
		"Yellow Green":           YellowGreen,
		"YellowGreen":            Yellowgreen,
	}
)

// Random returns random color
