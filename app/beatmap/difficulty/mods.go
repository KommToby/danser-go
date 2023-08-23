package difficulty

import (
	"reflect"

	"github.com/kommtoby/rplpa"
	"github.com/wieku/danser-go/app/settings"
)

type Modifier int64

const (
	None   = Modifier(iota)
	NoFail = Modifier(1 << (iota - uint(1)))
	Easy
	TouchDevice
	Hidden
	HardRock
	SuddenDeath
	DoubleTime
	Relax
	HalfTime
	Nightcore // Only set along with DoubleTime. i.e: NC only gives 576
	Flashlight
	Autoplay
	SpunOut
	Relax2  // Autopilot
	Perfect // Only set along with SuddenDeath. i.e: PF only gives 16416
	Key4
	Key5
	Key6
	Key7
	Key8
	FadeIn
	Random
	Cinema
	Target
	Key9
	KeyCoop
	Key1
	Key3
	Key2
	ScoreV2
	LastMod
	Daycore
	Blinds
	StrictTracking
	AccuracyChallenge
	DifficultyAdjust
	Classic
	Alternate
	SingleTap
	Transform
	Wiggle
	SpinIn
	Grow
	Deflate
	WindUp
	WindDown
	Traceable
	BarrelRoll
	ApproachDifferent
	Muted
	NoScope
	Magnetised
	Repel
	AdaptiveSpeed
	FreezeFrame
	Bubbles
	Synesthesia

	// DifficultyAdjustMask is outdated, use GetDiffMaskedMods instead
	DifficultyAdjustMask    = HardRock | Easy | DoubleTime | Nightcore | HalfTime | Daycore | Flashlight | Relax
	difficultyAdjustMaskNew = HardRock | Easy | DoubleTime | HalfTime | Flashlight | Relax | TouchDevice
)

// GetDiffMaskedMods should be used instead of DifficultyAdjustMask. In 220930 deployment, HDFL is a separate mod difficulty wise
func GetDiffMaskedMods(mods Modifier) Modifier {
	//Probably redundant
	if mods.Active(Nightcore) {
		mods = (mods & (^Nightcore)) | DoubleTime
	}

	if mods.Active(Daycore) {
		mods = (mods & (^Daycore)) | HalfTime
	}

	base := difficultyAdjustMaskNew & mods

	if mods&(Hidden|Flashlight) == (Hidden | Flashlight) {
		base |= Hidden
	}

	return base
}

var modsString = [...]string{
	"NF",
	"EZ",
	"TD",
	"HD",
	"HR",
	"SD",
	"DT",
	"RX",
	"HT",
	"NC",
	"FL",
	"AT", // Auto.
	"SO",
	"AP", // Autopilot.
	"PF",
	"K4",
	"K5",
	"K6",
	"K7",
	"K8",
	"FI",
	"RN", // Random
	"CN",
	"TG",
	"K9",
	"K0",
	"K1",
	"K3",
	"K2",
	"V2",
	"LM",
	"DC",
	"BL", // Blinds
	"ST", // Strict Tracking (No Multiplier)
	"AC", // Accuracy Challenge (No Multiplier)
	"DA", // Difficulty Adjust (Always 0.5x)
	"CL", // Classic (No Multiplier)
	"AL", // Alternate (No Multiplier)
	"SG", // Singletap (No Multiplier)
	"TR", // Transform
	"WG", // Wiggle
	"SI", // Spin In
	"GR", // Grow
	"DF", // Deflate
	"WU", // Wind Up (Always 0.5x)
	"WD", // Wind Down (Always 0.5x)
	"TC", // Traceable
	"BR", // Barrel Roll
	"AD", // Approach Different
	"MU", // Muted
	"NS", // No Scope
	"MG", // Magnetised (Always 0.5x)
	"RP", // Repel
	"AS", // Adaptive Speed
	"FR", // Freeze Frame
	"BU", // Bubbles
	"SY", // Synesthesia
}

var modsStringFull = [...]string{
	"NoFail",
	"Easy",
	"TouchDevice",
	"Hidden",
	"HardRock",
	"SuddenDeath",
	"DoubleTime",
	"Relax",
	"HalfTime",
	"Nightcore",
	"Flashlight",
	"Autoplay",
	"SpunOut",
	"Relax2", // Autopilot
	"Perfect",
	"Key4",
	"Key5",
	"Key6",
	"Key7",
	"Key8",
	"FadeIn",
	"Random",
	"Cinema",
	"Target",
	"Key9",
	"KeyCoop",
	"Key1",
	"Key3",
	"Key2",
	"ScoreV2",
	"LastMod",
	"Daycore",
	"Blinds",
	"StrictTracking",
	"AccuracyChallenge",
	"DifficultyAdjust",
	"Classic",
	"Alternate",
	"SingleTap",
	"Transform",
	"Wiggle",
	"SpinIn",
	"Grow",
	"Deflate",
	"WindUp",
	"WindDown",
	"Traceable",
	"BarrelRoll",
	"ApproachDifferent",
	"Muted",
	"NoScope",
	"Magnetised",
	"Repel",
	"AdaptiveSpeed",
	"FreezeFrame",
	"Bubbles",
	"Synesthesia",
}

func (mods Modifier) GetScoreMultiplier() float64 {
	multiplier := 1.0

	for _, enabledMod := range mods.StringFull() { // checks every enabled mod against the modmult struct in gameplay settings
		if value, ok := getFieldValue(settings.Gameplay.ModMults, enabledMod); ok {
			multiplier *= value // if the mod has a defined multiplier in gameplay settings it applies it
		}
	}
	return multiplier
}

func getFieldValue(data interface{}, fieldName string) (float64, bool) {
	// get the mod multiplier if it exists in the modmult struct
	value := reflectValue(data)
	field := value.FieldByName(fieldName)

	if field.IsValid() && field.Kind() == reflect.Float64 {
		return field.Float(), true
	}
	// if not, it wont have an applied multiplier since not implemented. Currently mostly mania mods
	return 1.0, false
}

func reflectValue(data interface{}) reflect.Value {
	// create a reflect value from the modmults
	value := reflect.ValueOf(data)

	if value.Kind() == reflect.Ptr {
		// if the value is a pointer, return the value assigned to it
		value = value.Elem()
	}
	return value
}

func (mods Modifier) String() (s string) {
	if mods.Active(Nightcore) {
		mods &= ^DoubleTime
	}

	if mods.Active(Daycore) {
		mods &= ^HalfTime
	}

	if mods.Active(Perfect) {
		mods &= ^SuddenDeath
	}

	for i := 0; i < len(modsString); i++ {
		activated := mods&1 == 1
		if activated {
			s += modsString[i]
		}

		mods >>= 1
	}

	return
}

func (mods Modifier) StringFull() (s []string) {
	if mods.Active(Nightcore) {
		mods &= ^DoubleTime
	}

	if mods.Active(Daycore) {
		mods &= ^HalfTime
	}

	if mods.Active(Perfect) {
		mods &= ^SuddenDeath
	}

	for i := 0; i < len(modsString); i++ {
		activated := mods&1 == 1
		if activated {
			s = append(s, modsStringFull[i])
		}

		mods >>= 1
	}

	return
}

func ParseMods(mods string) (m Modifier) {
	modsSl := make([]string, len(mods)/2)
	for n, modPart := range mods {
		modsSl[n/2] += string(modPart)
	}

	for _, mod := range modsSl {
		for index, availableMod := range modsString {
			if availableMod == mod {
				m |= 1 << uint(index)
				break
			}
		}
	}

	if m.Active(Nightcore) {
		m |= DoubleTime
	}

	if m.Active(Perfect) {
		m |= SuddenDeath
	}

	if m.Active(Daycore) {
		m |= HalfTime
	}

	return
}

func ParseLazerMods(scoreInfo *rplpa.ScoreInfo) (m Modifier) {
	for _, mod := range scoreInfo.Mods {
		for index, availableMod := range modsString {
			if availableMod == mod.Acronym {
				m |= 1 << uint(index)
				break
			}
		}
	}

	return m
}

func (mods Modifier) Active(mod Modifier) bool {
	return mods&mod > 0
}

func (mods Modifier) Compatible() bool {
	if mods == None {
		return true
	}

	if mods.Active(Target) ||
		(mods.Active(HardRock) && mods.Active(Easy)) ||
		((mods.Active(Nightcore) || mods.Active(DoubleTime)) && (mods.Active(HalfTime) || mods.Active(Daycore))) ||
		((mods.Active(Perfect) || mods.Active(SuddenDeath)) && mods.Active(NoFail)) ||
		(mods.Active(Relax) && mods.Active(Relax2)) ||
		((mods.Active(Relax) || mods.Active(Relax2)) && (mods.Active(SuddenDeath) || mods.Active(Perfect) || mods.Active(Autoplay) || mods.Active(NoFail))) ||
		(mods.Active(Relax2) && mods.Active(SpunOut)) {
		return false
	}

	return true
}
