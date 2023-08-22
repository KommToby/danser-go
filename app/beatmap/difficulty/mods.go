package difficulty

import (
	"github.com/kommtoby/rplpa"
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
	"DA", // Difficulty Adjust
	"CL", // Classic
	"AL", // Alternate
	"SG", // Singletap
	"TR", // Transform
	"WG", // Wiggle
	"SI", // Spin In
	"GR", // Grow
	"DF", // Deflate
	"WU", // Wind Up
	"WD", // Wind Down
	"TC", // Traceable
	"BR", // Barrel Roll
	"AD", // Approach Different
	"MU", // Muted
	"NS", // No Scope
	"MG", // Magnetised
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
	"Relax2",
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

	if mods&NoFail > 0 && mods&ScoreV2 == 0 {
		multiplier *= 0.5
	}

	if mods&Easy > 0 {
		multiplier *= 0.5
	}

	if mods&HalfTime > 0 {
		multiplier *= 0.3
	}

	if mods&Hidden > 0 {
		multiplier *= 1.06
	}

	if mods&HardRock > 0 {
		if mods&ScoreV2 > 0 {
			multiplier *= 1.10
		} else {
			multiplier *= 1.06
		}
	}

	if mods&DoubleTime > 0 {
		if mods&ScoreV2 > 0 {
			multiplier *= 1.20
		} else {
			multiplier *= 1.12
		}
	}

	if mods&Flashlight > 0 {
		multiplier *= 1.12
	}

	if (mods&Relax | mods&Relax2) > 0 {
		multiplier = 0
	}

	if mods&SpunOut > 0 {
		multiplier *= 0.9
	}

	if containsMod(mods.StringFull(), "DifficultyAdjust") {
		multiplier *= 0.1 // this SHOULD be 0.5, but for the case of osu!collab series (which is why im doing this) - im setting it to 0.1 temporarily.
	}

	return multiplier
}

// might be redundant
func containsMod(mods []string, targetMod string) bool {
	for _, mod := range mods {
		if mod == targetMod {
			return true
		}
	}
	return false
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
