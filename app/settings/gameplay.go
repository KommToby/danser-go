package settings

var Gameplay = initGameplay()

func initGameplay() *gameplay {
	return &gameplay{
		HitErrorMeter: &hitError{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			PointFadeOutTime:     10,
			ShowPositionalMisses: true,
			PositionalMissScale:  1.5,
			ShowUnstableRate:     true,
			UnstableRateDecimals: 0,
			UnstableRateScale:    1.0,
			StaticUnstableRate:   false,
			ScaleWithSpeed:       false,
		},
		AimErrorMeter: &aimError{
			hudElementPosition: &hudElementPosition{
				hudElement: &hudElement{
					Show:    false,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XPosition: 1350,
				YPosition: 650,
			},
			PointFadeOutTime:     10,
			DotScale:             1,
			Align:                "Right",
			ShowUnstableRate:     false,
			UnstableRateScale:    1,
			UnstableRateDecimals: 0,
			StaticUnstableRate:   false,
			CapPositionalMisses:  true,
			AngleNormalized:      false,
		},
		Score: &score{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			ProgressBar:     "Pie",
			ShowGradeAlways: false,
			StaticScore:     false,
			StaticAccuracy:  false,
		},
		HpBar: &hudElementOffset{
			hudElement: &hudElement{
				Show:    true,
				Scale:   1.0,
				Opacity: 1.0,
			},
			XOffset: 0,
			YOffset: 0,
		},
		ComboCounter: &comboCounter{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			Static: false,
		},
		PPCounter: &ppCounter{
			hudElementPosition: &hudElementPosition{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XPosition: 5,
				YPosition: 150,
			},
			Color: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			Decimals:         0,
			Align:            "CentreLeft",
			ShowInResults:    true,
			ShowPPComponents: false,
			Static:           false,
		},
		HitCounter: &hitCounter{
			hudElementPosition: &hudElementPosition{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XPosition: 5,
				YPosition: 190,
			},
			Color300: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			Color100: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			Color50: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			ColorMiss: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			ColorSB: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			Spacing:          48,
			FontScale:        1,
			Align:            "Left",
			ValueAlign:       "Left",
			Vertical:         false,
			Show300:          false,
			ShowSliderBreaks: false,
		},
		StrainGraph: &strainGraph{
			Show:      true,
			Opacity:   1,
			XPosition: 5,
			YPosition: 310,
			Align:     "BottomLeft",
			Width:     130,
			Height:    70,
			BgColor: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      0.2,
			},
			FgColor: &HSV{
				Hue:        297,
				Saturation: 0.4,
				Value:      0.92,
			},
			Outline: &outline{
				Show:          false,
				Width:         2,
				InnerDarkness: 0.5,
				InnerOpacity:  0.5,
			},
		},
		KeyOverlay: &hudElementOffset{
			hudElement: &hudElement{
				Show:    true,
				Scale:   1.0,
				Opacity: 1.0,
			},
			XOffset: 0,
			YOffset: 0,
		},
		ScoreBoard: &scoreBoard{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			ModsOnly:       false,
			AlignRight:     false,
			HideOthers:     false,
			ShowAvatars:    false,
			ExplosionScale: 1.0,
		},
		Mods: &mods{
			hudElementOffset: &hudElementOffset{
				hudElement: &hudElement{
					Show:    true,
					Scale:   1.0,
					Opacity: 1.0,
				},
				XOffset: 0,
				YOffset: 0,
			},
			HideInReplays:     false,
			FoldInReplays:     false,
			AdditionalSpacing: 0,
		},
		ModMults: &modMults{
			NoFail:            0.5,
			Easy:              0.5,
			Hidden:            1.06,
			HardRock:          1.06,
			SuddenDeath:       1.00,
			DoubleTime:        1.2, // on lazer this is 1.1
			Relax:             0,   // on lazer this is 0.1
			HalfTime:          0.3, // on lazer this is 0.7
			Nightcore:         1.2, // on lazer this is 1.1
			Flashlight:        1.12,
			SpunOut:           0.9,
			Relax2:            0, // on lazer this is 0.1 - autopilot
			Perfect:           1.00,
			Random:            1.00,
			Daycore:           0.3, // on lazer this is 0.7 - matched with HT
			Blinds:            1.12,
			StrictTracking:    1.00,
			AccuracyChallenge: 1.00,
			DifficultyAdjust:  0.5,
			Classic:           1.00,
			Alternate:         1.00,
			SingleTap:         1.00,
			Transform:         1.00,
			Wiggle:            1.00,
			SpinIn:            1.00,
			Grow:              1.00,
			Deflate:           1.00,
			WindUp:            1.00,
			WindDown:          1.00,
			Traceable:         1.00,
			BarrelRoll:        1.00,
			ApproachDifferent: 1.00,
			Muted:             1.00,
			NoScope:           1.00,
			Magnetised:        1.00,
			Repel:             1.00,
			AdaptiveSpeed:     1.00,
			FreezeFrame:       1.00,
			Bubbles:           1.00,
			Synesthesia:       1.00,
		},
		Boundaries: &boundaries{
			Enabled:         true,
			BorderThickness: 1,
			BorderFill:      1,
			BorderColor: &HSV{
				Hue:        0,
				Saturation: 0,
				Value:      1,
			},
			BorderOpacity: 1,
			BackgroundColor: &HSV{
				Hue:        0,
				Saturation: 1,
				Value:      0,
			},
			BackgroundOpacity: 0.5,
		},
		Underlay: &underlay{
			Path:       "",
			AboveHpBar: false,
		},
		HUDFont:                 "",
		ShowResultsScreen:       true,
		ResultsScreenTime:       5,
		ResultsUseLocalTimeZone: false,
		ShowWarningArrows:       true,
		ShowHitLighting:         false,
		FlashlightDim:           1,
		PlayUsername:            "Guest",
		IgnoreFailsInReplays:    false,
		UseLazerPP:              false,
	}
}

type gameplay struct {
	HitErrorMeter           *hitError
	AimErrorMeter           *aimError
	Score                   *score
	HpBar                   *hudElementOffset
	ComboCounter            *comboCounter
	PPCounter               *ppCounter
	HitCounter              *hitCounter
	StrainGraph             *strainGraph
	KeyOverlay              *hudElementOffset
	ScoreBoard              *scoreBoard
	Mods                    *mods
	ModMults                *modMults
	Boundaries              *boundaries
	Underlay                *underlay
	HUDFont                 string  `label:"Overlay (HUD) font" file:"Select HUD font" filter:"TrueType/OpenType Font (*.ttf, *.otf)|ttf,otf" tooltip:"Sets the font that will be used for PP/UR/hit counts" liveedit:"false"`
	ShowResultsScreen       bool    `liveedit:"false"`
	ResultsScreenTime       float64 `label:"Results screen duration" min:"1" max:"20" format:"%.1fs" liveedit:"false"`
	ResultsUseLocalTimeZone bool    `label:"Show PC's time zone instead of UTC"`
	ShowWarningArrows       bool
	ShowHitLighting         bool
	FlashlightDim           float64
	PlayUsername            string `liveedit:"false"`
	IgnoreFailsInReplays    bool
	UseLazerPP              bool `liveedit:"false" skip:"true"`
}

type boundaries struct {
	Enabled bool

	BorderThickness float64 `min:"0.5" max:"10" format:"%.1f o!px"`
	BorderFill      float64

	BorderColor   *HSV    `short:"true"`
	BorderOpacity float64 `scale:"100.0" format:"%.0f%%"`

	BackgroundColor   *HSV    `short:"true"`
	BackgroundOpacity float64 `scale:"100.0" format:"%.0f%%"`
}

type hudElement struct {
	Show    bool
	Scale   float64 `max:"3" scale:"100.0" format:"%.0f%%"`
	Opacity float64 `scale:"100.0" format:"%.0f%%"`
}

type hudElementOffset struct {
	*hudElement
	offset  string  `vector:"true" left:"XOffset" right:"YOffset"`
	XOffset float64 `min:"-10000" max:"10000"`
	YOffset float64 `min:"-10000" max:"10000"`
}

type hudElementPosition struct {
	*hudElement
	position  string  `vector:"true" left:"XPosition" right:"YPosition"`
	XPosition float64 `min:"-10000" max:"10000"`
	YPosition float64 `min:"-10000" max:"10000"`
}

type hitError struct {
	*hudElementOffset
	PointFadeOutTime     float64 `max:"10" format:"%.1fs"`
	ShowPositionalMisses bool
	PositionalMissScale  float64 `min:"1" max:"2" scale:"100" format:"%.0f%%"`
	ShowUnstableRate     bool
	UnstableRateDecimals int     `max:"5"`
	UnstableRateScale    float64 `min:"0.1" max:"5" scale:"100" format:"%.0f%%"`
	StaticUnstableRate   bool
	ScaleWithSpeed       bool
}

type aimError struct {
	*hudElementPosition
	PointFadeOutTime     float64 `max:"10" format:"%.1fs"`
	DotScale             float64 `min:"0.1" max:"5" scale:"100" format:"%.0f%%"`
	Align                string  `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`
	ShowUnstableRate     bool
	UnstableRateScale    float64 `min:"0.1" max:"5" scale:"100" format:"%.0f%%"`
	UnstableRateDecimals int     `max:"5"`
	StaticUnstableRate   bool
	CapPositionalMisses  bool
	AngleNormalized      bool
}

type score struct {
	*hudElementOffset
	ProgressBar     string `combo:"Pie,Bar,BottomRight,Bottom"`
	ShowGradeAlways bool   `label:"Always show grade"`
	StaticScore     bool
	StaticAccuracy  bool
}

type comboCounter struct {
	*hudElementOffset
	Static bool
}

type ppCounter struct {
	*hudElementPosition
	Color            *HSV   `short:"true"`
	Decimals         int    `max:"5"`
	Align            string `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`
	ShowInResults    bool
	ShowPPComponents bool `label:"Show PP breakdown"`
	Static           bool
}

type hitCounter struct {
	*hudElementPosition
	Color            []*HSV  `json:",omitempty" new:"InitHSV" label:"Color list" skip:"true"`
	Color300         *HSV    `label:"Color of 300s"`
	Color100         *HSV    `label:"Color of 100s"`
	Color50          *HSV    `label:"Color of 50s"`
	ColorMiss        *HSV    `label:"Color of misses"`
	ColorSB          *HSV    `label:"Color of slider breaks"`
	Spacing          float64 `string:"true" min:"0" max:"1366"`
	FontScale        float64 `min:"0.1" max:"5" scale:"100" format:"%.0f%%"`
	Align            string  `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`
	ValueAlign       string  `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`
	Vertical         bool
	Show300          bool `label:"Show perfect hits"`
	ShowSliderBreaks bool
}

type scoreBoard struct {
	*hudElementOffset
	ModsOnly       bool `label:"Show mod leaderboard"`
	AlignRight     bool `label:"Align to the right" label:"Simulates the second team of osu! multiplayer"`
	HideOthers     bool
	ShowAvatars    bool
	ExplosionScale float64 `min:"0.1" max:"2" scale:"100" format:"%.0f%%"`
}

type mods struct {
	*hudElementOffset
	HideInReplays     bool
	FoldInReplays     bool
	AdditionalSpacing float64 `string:"true" min:"-1366" max:"1366"`
}

type modMults struct {
	NoFail float64 `min:"0" max:"2" string:"true"`
	Easy   float64 `min:"0" max:"2" string:"true"`
	// TouchDevice float64 `min:"0.1" max:"2"`
	Hidden      float64 `min:"0" max:"2" string:"true"`
	HardRock    float64 `min:"0" max:"2" string:"true"`
	SuddenDeath float64 `min:"0" max:"2" string:"true"`
	DoubleTime  float64 `min:"0" max:"2" string:"true"`
	Relax       float64 `min:"0" max:"2" string:"true"`
	HalfTime    float64 `min:"0" max:"2" string:"true"`
	Nightcore   float64 `min:"0" max:"2" string:"true"`
	Flashlight  float64 `min:"0" max:"2" string:"true"`
	// Autoplay    float64 `min:"0.1" max:"2"`
	SpunOut float64 `min:"0" max:"2" string:"true"`
	Relax2  float64 `min:"0" max:"2" string:"true"` // Autopilot
	Perfect float64 `min:"0" max:"2" string:"true"`
	// Key4              float64 `min:"0.1" max:"2"`
	// Key5              float64 `min:"0.1" max:"2"`
	// Key6              float64 `min:"0.1" max:"2"`
	// Key7              float64 `min:"0.1" max:"2"`
	// Key8              float64 `min:"0.1" max:"2"`
	// FadeIn            float64 `min:"0.1" max:"2"`
	Random float64 `min:"0" max:"2" string:"true"`
	// Cinema float64 `min:"0.1" max:"2"`
	// Target float64 `min:"0.1" max:"2"`
	// Key9              float64 `min:"0.1" max:"2"`
	// KeyCoop           float64 `min:"0.1" max:"2"`
	// Key1              float64 `min:"0.1" max:"2"`
	// Key3              float64 `min:"0.1" max:"2"`
	// Key2              float64 `min:"0.1" max:"2"`
	// ScoreV2           float64 `min:"0.1" max:"2"`
	// LastMod           float64 `min:"0.1" max:"2"` // Cinema
	Daycore           float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Blinds            float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	StrictTracking    float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	AccuracyChallenge float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	DifficultyAdjust  float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Classic           float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Alternate         float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	SingleTap         float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Transform         float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Wiggle            float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	SpinIn            float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Grow              float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Deflate           float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	WindUp            float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	WindDown          float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Traceable         float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	BarrelRoll        float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	ApproachDifferent float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Muted             float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	NoScope           float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Magnetised        float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Repel             float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	AdaptiveSpeed     float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	FreezeFrame       float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Bubbles           float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
	Synesthesia       float64 `min:"0" max:"2" string:"true" tooltip:"lazer-only mod"`
}

type strainGraph struct {
	Show    bool
	Opacity float64 `scale:"100.0" format:"%.0f%%"`

	position  string  `vector:"true" left:"XPosition" right:"YPosition"`
	XPosition float64 `min:"-10000" max:"10000"`
	YPosition float64 `min:"-10000" max:"10000"`

	Align string `combo:"TopLeft,Top,TopRight,Left,Centre,Right,BottomLeft,Bottom,BottomRight"`

	size   string  `vector:"true" left:"Width" right:"Height"`
	Width  float64 `string:"true" min:"1" max:"10000"`
	Height float64 `string:"true" min:"1" max:"768"`

	BgColor *HSV `label:"Background color" short:"true"`
	FgColor *HSV `label:"Foreground color" short:"true"`

	Outline *outline
}

type outline struct {
	Show          bool
	Width         float64 `min:"1" max:"5"`
	InnerDarkness float64 `scale:"100.0" format:"%.0f%%" tooltip:"Darkness of filled shape, only applicable when DrawOutline is enabled"`
	InnerOpacity  float64 `scale:"100.0" format:"%.0f%%" tooltip:"Opacity of filled shape, only applicable when DrawOutline is enabled"`
}

type underlay struct {
	Path       string `file:"Select underlay image" filter:"PNG file (*.png)|png" tooltip:"PNG file that will be used as HUD background (similar to custom HP bar backgrounds). It's scaled automatically to fit the screen vertically" liveedit:"false"`
	AboveHpBar bool   `label:"Show underlay above HP bar" tooltip:"Use this if HP bar background is large"`
}
