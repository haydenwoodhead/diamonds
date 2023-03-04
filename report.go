package main

type Definition struct {
	Name       string
	Schedule   string
	Recipients []string
	Currency   string
	Markups    []Markup
	Tax        int
	Shipping   int
	Variants   []Variant
}

type Variant struct {
	Type                    VariantType
	Shapes                  []Shape
	Color                   *FromToSelector[Color]
	Size                    *FromToSelector[float32]
	FancyColors             []FancyColor
	FancyColorIntensity     *FromToSelector[FancyColorIntensity]
	Clarity                 *FromToSelector[Clarity]
	Cur                     *FromToSelector[Cut]
	Polish                  *FromToSelector[Polish]
	Symmetry                *FromToSelector[Symmetry]
	Price                   *FromToSelector[int]
	Labs                    []Lab
	FluorescenceIntensities *FromToSelector[FluorescenceIntensities]
	FluorescenceColors      *FromToSelector[FluorescenceColors]
	DepthPercentage         *FromToSelector[float32]
	TablePercentage         *FromToSelector[float32]
	MeasLength              *FromToSelector[float32]
	MeasWidth               *FromToSelector[float32]
	MeasDepth               *FromToSelector[float32]
	Girdle                  *MinMaxSelector[float32]
	CutletSizes             []CutletSize
	EyeCleans               []EyeClean
}

type Markup struct {
	MinInclusive int // TODO: should this be > or >=
	MaxInclusive int
	Multiple     int
}

type FromToSelector[T any] struct {
	From T
	To   T
}

// TODO: patch over this fucking dogshit API requiring min/max for girdle
type MinMaxSelector[T any] struct {
	Min T
	Max T
}

type VariantType string

var (
	White VariantType = "white"
	Fancy VariantType = "fancy"
)

type Shape string

var (
	ClarityRound    Shape = "Round"
	ClarityPear     Shape = "Pear"
	ClarityPrincess Shape = "Princess"
	ClarityMarquise Shape = "Marquise"
	ClarityOval     Shape = "Oval"
	ClarityRaidant  Shape = "Radiant"
	ClarityEmerald  Shape = "Emerald"
	ClarityHeart    Shape = "Heart"
	ClarityCushion  Shape = "Cushion"
	ClarityAsscher  Shape = "Asscher"
)

type Color string

var (
	ColorD Color = "D"
	ColorE Color = "E"
	ColorF Color = "F"
	ColorG Color = "G"
	ColorH Color = "H"
	ColorI Color = "I"
	ColorJ Color = "J"
	ColorK Color = "K"
	ColorL Color = "L"
	ColorM Color = "M"
)

type FancyColor string

var (
	FancyColorYellow    FancyColor = "Yellow"
	FancyColorPink      FancyColor = "Pink"
	FancyColorOrange    FancyColor = "Orange"
	FancyColorGreen     FancyColor = "Green"
	FancyColorGray      FancyColor = "Gray"
	FancyColorBrown     FancyColor = "Brown"
	FancyColorBlue      FancyColor = "Blue"
	FancyColorBlack     FancyColor = "Black"
	FancyColorRed       FancyColor = "Red"
	FancyColorPurple    FancyColor = "Purple"
	FancyColorViolet    FancyColor = "Violet"
	FancyColorChampagne FancyColor = "Champagne"
	FancyColorCognac    FancyColor = "Cognac"
	FancyColorChameleon FancyColor = "Chameleon"
)

type FancyColorIntensity string

var (
	FancyColorIntensityFaint                 FancyColorIntensity = "Faint"
	FancyColorIntensityVeryLight             FancyColorIntensity = "Very Light"
	FancyColorIntensityLight                 FancyColorIntensity = "Light"
	FancyColorIntensityFancyIntensityLight   FancyColorIntensity = "Fancy Light"
	FancyColorIntensityFancyIntensityDark    FancyColorIntensity = "Fancy Dark"
	FancyColorIntensityFancyIntensityIntense FancyColorIntensity = "Fancy Intense"
	FancyColorIntensityFancyIntensityVivid   FancyColorIntensity = "Fancy Vivid"
	FancyColorIntensityFancyIntensityDeep    FancyColorIntensity = "Fancy Deep"
)

type Clarity string

var (
	ClarityIF   Clarity = "IF"
	ClarityVVS1 Clarity = "VVS1"
	ClarityVVS2 Clarity = "VVS2"
	ClarityVS1  Clarity = "VS1"
	ClarityVS2  Clarity = "VS2"
	ClaritySI1  Clarity = "SI1"
	ClaritySI2  Clarity = "SI2"
	ClaritySI3  Clarity = "SI3"
	ClarityI1   Clarity = "I1"
	ClarityI2   Clarity = "I2"
	ClarityI3   Clarity = "I3"
)

type Cut string

var (
	CutExcellent Cut = "Excellent"
	CutVeryGood  Cut = "VeryGood"
	CutGood      Cut = "Good"
	CutFair      Cut = "Fair"
	CutPoor      Cut = "Poor"
)

type Polish string

var (
	PolishExcellent Polish = "Excellent"
	PolishVeryGood  Polish = "VeryGood"
	PolishGood      Polish = "Good"
	PolishFair      Polish = "Fair"
	PolishPoor      Polish = "Poor"
)

type Symmetry string

var (
	SymmetryExcellent Symmetry = "Excellent"
	SymmetryVeryGood  Symmetry = "VeryGood"
	SymmetryGood      Symmetry = "Good"
	SymmetryFair      Symmetry = "Fair"
	SymmetryPoor      Symmetry = "Poor"
)

type Lab string

var (
	LabGIA   Lab = "GIA"
	LabIGI   Lab = "IGI"
	LabAGS   Lab = "AGS"
	LabHRD   Lab = "HRD"
	LabPGS   Lab = "PGS"
	LabDCLA  Lab = "DCLA"
	LabVGR   Lab = "VGR"
	LabGCAL  Lab = "GCAL"
	LabNGTC  Lab = "NGTC"
	LabGSI   Lab = "GSI"
	LabDBGIS Lab = "DBGIS"
	LabNONE  Lab = "NONE"
)

type FluorescenceIntensities string

var (
	FluorescenceIntensitiesVerySlight FluorescenceIntensities = "VerySlight"
	FluorescenceIntensitiesFaint      FluorescenceIntensities = "Faint"
	FluorescenceIntensitiesMedium     FluorescenceIntensities = "Medium"
	FluorescenceIntensitiesSlight     FluorescenceIntensities = "Slight"
	FluorescenceIntensitiesStrong     FluorescenceIntensities = "Strong"
	FluorescenceIntensitiesVeryStrong FluorescenceIntensities = "VeryStrong"
	FluorescenceIntensitiesNone       FluorescenceIntensities = "None"
)

type FluorescenceColors string

var (
	FluorescenceColorsBlue   FluorescenceColors = "Blue"
	FluorescenceColorsYellow FluorescenceColors = "Yellow"
	FluorescenceColorsGreen  FluorescenceColors = "Green"
	FluorescenceColorsRed    FluorescenceColors = "Red"
	FluorescenceColorsOrange FluorescenceColors = "Orange"
	FluorescenceColorsWhite  FluorescenceColors = "White"
)

type Girdle string

var (
	GirdleExtrThin      Girdle = "ExtrThin"
	GirdleVeryThin      Girdle = "VeryThin"
	GirdleThin          Girdle = "Thin"
	GirdleSlightlyThin  Girdle = "SlightlyThin"
	GirdleMedium        Girdle = "Medium"
	GirdleSlightlyThick Girdle = "SlightlyThick"
	GirdleThick         Girdle = "Thick"
	GirdleVeryThick     Girdle = "VeryThick"
	GirdleExtrThick     Girdle = "ExtrThick"
)

type CutletSize string

var (
	CuletSizeVerySmall CutletSize = "VerySmall"
	CuletSizeSmall     CutletSize = "Small"
	CuletSizeMedium    CutletSize = "Medium"
	CuletSizeLarge     CutletSize = "Large"
)

type EyeClean string

var (
	EyeCleanYes        EyeClean = "Yes"
	EyeCleanBorderline EyeClean = "Borderline"
	EyeCleanNo         EyeClean = "No"
)
