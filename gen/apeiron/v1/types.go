package apeironv1

type Skill struct {
	Id                  string
	BaseDamage          float64
	StaminaCost         float64
	ManaCost            float64
	HealthCost          float64
	CooldownMs          int32
	GlobalCooldownMs    int32
	MaxRange            float64
	RequiresTarget      bool
	RequiresLineOfSight bool
	AllowMovement       bool
	MovementLockMs      int32
	SkillType           string
	TargetType          string
	DamageMultiplier    float64
	PostureDamage       float64
	IsBlockable         bool
	IsParryable         bool
	MaxTargets          int32
	MovementDistance    float64
	ComboGroup          string
	ComboStep           int32
	ComboWindowMs       int32
	ComboResetMs        int32
	Interruptible       bool
	Tags                []string
	DamageType          string
	ElementalType       string
}

func (s *Skill) GetId() string {
	if s == nil {
		return ""
	}
	return s.Id
}

func (s *Skill) GetBaseDamage() float64 {
	if s == nil {
		return 0
	}
	return s.BaseDamage
}

func (s *Skill) GetCooldownMs() int32 {
	if s == nil {
		return 0
	}
	return s.CooldownMs
}

func (s *Skill) GetGlobalCooldownMs() int32 {
	if s == nil {
		return 0
	}
	return s.GlobalCooldownMs
}

func (s *Skill) GetComboIndex() int64 {
	if s == nil {
		return 0
	}
	return int64(s.ComboStep)
}

func (s *Skill) GetComboWindowMs() int32 {
	if s == nil {
		return 0
	}
	return s.ComboWindowMs
}

func (s *Skill) GetMovementLockMs() int32 {
	if s == nil {
		return 0
	}
	return s.MovementLockMs
}

func (s *Skill) GetMaxRange() float64 {
	if s == nil {
		return 0
	}
	return s.MaxRange
}

func (s *Skill) GetDamageType() string {
	if s == nil {
		return ""
	}
	return s.DamageType
}

func (s *Skill) GetElementalType() string {
	if s == nil {
		return ""
	}
	return s.ElementalType
}

func (s *Skill) GetSkillType() string {
	if s == nil {
		return ""
	}
	return s.SkillType
}

func (s *Skill) GetComboGroup() string {
	if s == nil {
		return ""
	}
	return s.ComboGroup
}

func (s *Skill) GetMaxTargets() int32 {
	if s == nil {
		return 0
	}
	return s.MaxTargets
}

func (s *Skill) GetMovementDistance() float64 {
	if s == nil {
		return 0
	}
	return s.MovementDistance
}

type SkillMovementProfile struct {
	Id                     string
	MovementType           string
	Distance               float64
	Speed                  float64
	DurationMs             int32
	MovementStartPhase     string
	MovementStartOffsetMs  int32
	TakeoffMs              int32
	LandingLockMs          int32
	ArcHeight              float64
	ArcCurve               string
	Bounds                 string
	SteeringPolicy         string
	MaxTurnDegPerSec       float64
	MaxTotalRedirectAngle  float64
	RedirectLockoutMs      int32
	CanPhaseThroughTargets bool
	MinLandingDistance     float64
	DesiredLandingDistance float64
	StopAtContactRatio     float64
	AppliesKnockback       bool
	KnockbackDistance      float64
	KnockbackSpeed         float64
}

func (p *SkillMovementProfile) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}

type SkillHitboxProfile struct {
	Id                  string
	SkillId             string
	HitboxShape         string
	HitboxStartMs       int32
	HitboxEndMs         int32
	OffsetX             float64
	OffsetY             float64
	OffsetZ             float64
	Length              float64
	Radius              float64
	SizeX               float64
	SizeY               float64
	SizeZ               float64
	MotionProfile       *SkillHitboxMotionProfile
	DamageGroupId       string
	HitboxIndex         int32
	Angle               float64
	TargetType          *string
	MaxTargets          *int32
	Priority            int32
	RequiresLineOfSight bool
	CanHitNeutral       bool
}

func (p *SkillHitboxProfile) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}
func (p *SkillHitboxProfile) GetSkillId() string {
	if p == nil {
		return ""
	}
	return p.SkillId
}
func (p *SkillHitboxProfile) GetHitboxShape() string {
	if p == nil {
		return ""
	}
	return p.HitboxShape
}
func (p *SkillHitboxProfile) GetHitboxStartMs() int32 {
	if p == nil {
		return 0
	}
	return p.HitboxStartMs
}
func (p *SkillHitboxProfile) GetHitboxEndMs() int32 {
	if p == nil {
		return 0
	}
	return p.HitboxEndMs
}
func (p *SkillHitboxProfile) GetOffsetX() float64 {
	if p == nil {
		return 0
	}
	return p.OffsetX
}
func (p *SkillHitboxProfile) GetOffsetY() float64 {
	if p == nil {
		return 0
	}
	return p.OffsetY
}
func (p *SkillHitboxProfile) GetOffsetZ() float64 {
	if p == nil {
		return 0
	}
	return p.OffsetZ
}
func (p *SkillHitboxProfile) GetLength() float64 {
	if p == nil {
		return 0
	}
	return p.Length
}
func (p *SkillHitboxProfile) GetRadius() float64 {
	if p == nil {
		return 0
	}
	return p.Radius
}
func (p *SkillHitboxProfile) GetSizeX() float64 {
	if p == nil {
		return 0
	}
	return p.SizeX
}
func (p *SkillHitboxProfile) GetSizeY() float64 {
	if p == nil {
		return 0
	}
	return p.SizeY
}
func (p *SkillHitboxProfile) GetSizeZ() float64 {
	if p == nil {
		return 0
	}
	return p.SizeZ
}
func (p *SkillHitboxProfile) GetMotionProfile() *SkillHitboxMotionProfile {
	if p == nil {
		return nil
	}
	return p.MotionProfile
}
func (p *SkillHitboxProfile) GetDamageGroupId() string {
	if p == nil {
		return ""
	}
	return p.DamageGroupId
}

func (p *SkillHitboxProfile) GetMaxTargets() int32 {
	if p == nil || p.MaxTargets == nil {
		return 0
	}
	return *p.MaxTargets
}

type SkillHitboxMotionProfile struct {
	Id            string
	Enabled       bool
	MotionType    string
	TimeBasis     string
	Interpolation string
	SweepShape    string
	DamageGroupId string
	Samples       []*SkillHitboxMotionSample
}

func (p *SkillHitboxMotionProfile) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}
func (p *SkillHitboxMotionProfile) GetEnabled() bool { return p != nil && p.Enabled }
func (p *SkillHitboxMotionProfile) GetMotionType() string {
	if p == nil {
		return ""
	}
	return p.MotionType
}
func (p *SkillHitboxMotionProfile) GetTimeBasis() string {
	if p == nil {
		return ""
	}
	return p.TimeBasis
}
func (p *SkillHitboxMotionProfile) GetInterpolation() string {
	if p == nil {
		return ""
	}
	return p.Interpolation
}
func (p *SkillHitboxMotionProfile) GetSweepShape() string {
	if p == nil {
		return ""
	}
	return p.SweepShape
}
func (p *SkillHitboxMotionProfile) GetDamageGroupId() string {
	if p == nil {
		return ""
	}
	return p.DamageGroupId
}
func (p *SkillHitboxMotionProfile) GetSamples() []*SkillHitboxMotionSample {
	if p == nil {
		return nil
	}
	return p.Samples
}

type SkillHitboxMotionSample struct {
	SampleIndex   int32
	T             float64
	OffsetX       float64
	OffsetY       float64
	OffsetZ       float64
	Length        float64
	Radius        float64
	SizeX         float64
	SizeY         float64
	SizeZ         float64
	StartAngleDeg float64
	EndAngleDeg   float64
}

func (s *SkillHitboxMotionSample) GetSampleIndex() int32 {
	if s == nil {
		return 0
	}
	return s.SampleIndex
}
func (s *SkillHitboxMotionSample) GetT() float64 {
	if s == nil {
		return 0
	}
	return s.T
}
func (s *SkillHitboxMotionSample) GetOffsetX() float64 {
	if s == nil {
		return 0
	}
	return s.OffsetX
}
func (s *SkillHitboxMotionSample) GetOffsetY() float64 {
	if s == nil {
		return 0
	}
	return s.OffsetY
}
func (s *SkillHitboxMotionSample) GetOffsetZ() float64 {
	if s == nil {
		return 0
	}
	return s.OffsetZ
}
func (s *SkillHitboxMotionSample) GetLength() float64 {
	if s == nil {
		return 0
	}
	return s.Length
}
func (s *SkillHitboxMotionSample) GetRadius() float64 {
	if s == nil {
		return 0
	}
	return s.Radius
}
func (s *SkillHitboxMotionSample) GetSizeX() float64 {
	if s == nil {
		return 0
	}
	return s.SizeX
}
func (s *SkillHitboxMotionSample) GetSizeY() float64 {
	if s == nil {
		return 0
	}
	return s.SizeY
}
func (s *SkillHitboxMotionSample) GetSizeZ() float64 {
	if s == nil {
		return 0
	}
	return s.SizeZ
}
func (s *SkillHitboxMotionSample) GetStartAngleDeg() float64 {
	if s == nil {
		return 0
	}
	return s.StartAngleDeg
}
func (s *SkillHitboxMotionSample) GetEndAngleDeg() float64 {
	if s == nil {
		return 0
	}
	return s.EndAngleDeg
}

type SkillImpactProfile struct {
	SkillId               string
	ImpactType            string
	PoiseDamage           float64
	StaggerPower          float64
	InterruptPower        float64
	HitReaction           string
	GuardDamageMultiplier float64
}

type SkillTimingProfile struct {
	WindupMs           int32
	ActiveStartMs      int32
	ActiveEndMs        int32
	RecoveryMs         int32
	ActionLockMs       int32
	GlobalCooldownMs   int32
	MovementLockPolicy string
}

func (p *SkillTimingProfile) GetWindupMs() int32 {
	if p == nil {
		return 0
	}
	return p.WindupMs
}

func (p *SkillTimingProfile) GetActiveStartMs() int32 {
	if p == nil {
		return 0
	}
	return p.ActiveStartMs
}

func (p *SkillTimingProfile) GetActiveEndMs() int32 {
	if p == nil {
		return 0
	}
	return p.ActiveEndMs
}

func (p *SkillTimingProfile) GetRecoveryMs() int32 {
	if p == nil {
		return 0
	}
	return p.RecoveryMs
}

func (p *SkillTimingProfile) GetActionLockMs() int32 {
	if p == nil {
		return 0
	}
	return p.ActionLockMs
}

func (p *SkillTimingProfile) GetGlobalCooldownMs() int32 {
	if p == nil {
		return 0
	}
	return p.GlobalCooldownMs
}

func (p *SkillTimingProfile) GetMovementLockPolicy() string {
	if p == nil {
		return ""
	}
	return p.MovementLockPolicy
}

func (p *SkillImpactProfile) GetImpactType() string {
	if p == nil {
		return ""
	}
	return p.ImpactType
}

func (p *SkillImpactProfile) GetPoiseDamage() float64 {
	if p == nil {
		return 0
	}
	return p.PoiseDamage
}

func (p *SkillImpactProfile) GetStaggerPower() float64 {
	if p == nil {
		return 0
	}
	return p.StaggerPower
}

func (p *SkillImpactProfile) GetInterruptPower() float64 {
	if p == nil {
		return 0
	}
	return p.InterruptPower
}

func (p *SkillImpactProfile) GetHitReaction() string {
	if p == nil {
		return ""
	}
	return p.HitReaction
}

type CombatCoreProfile struct {
	DamageDealtMultiplier   float64
	CanBlock                bool
	BlockDamageReduction    float64
	MaxPosture              float64
	PostureDamageMultiplier float64
	PostureBreakDurationMs  int32
}

func (p *CombatCoreProfile) GetDamageDealtMultiplier() float64 {
	if p == nil {
		return 0
	}
	return p.DamageDealtMultiplier
}

func (p *CombatCoreProfile) GetPostureDamageMultiplier() float64 {
	if p == nil {
		return 0
	}
	return p.PostureDamageMultiplier
}

type StatusEffect struct {
	Id             string
	Name           string
	EffectType     string
	EffectCategory string
	ControlType    string
	StackingMode   string
	MaxStacks      int32
	DurationMs     int32
	IsPvpEnabled   bool
	BlocksMovement bool
	BlocksActions  bool
	BlocksSkills   bool
}

func (s *StatusEffect) GetId() string {
	if s == nil {
		return ""
	}
	return s.Id
}

func (s *StatusEffect) GetDurationMs() int32 {
	if s == nil {
		return 0
	}
	return s.DurationMs
}

type SkillControlEffect struct {
	Id              string
	Enabled         bool
	StatusEffectId  string
	DurationMs      int32
	ControlType     string
	ReleasePolicyId string
}

func (e *SkillControlEffect) GetId() string {
	if e == nil {
		return ""
	}
	return e.Id
}

func (e *SkillControlEffect) GetEnabled() bool {
	return e != nil && e.Enabled
}

func (e *SkillControlEffect) GetStatusEffectId() string {
	if e == nil {
		return ""
	}
	return e.StatusEffectId
}

func (e *SkillControlEffect) GetDurationMs() int32 {
	if e == nil {
		return 0
	}
	return e.DurationMs
}

func (e *SkillControlEffect) GetControlType() string {
	if e == nil {
		return ""
	}
	return e.ControlType
}

func (e *SkillControlEffect) GetReleasePolicyId() string {
	if e == nil {
		return ""
	}
	return e.ReleasePolicyId
}

type ItemTemplate struct {
	Id string
}
