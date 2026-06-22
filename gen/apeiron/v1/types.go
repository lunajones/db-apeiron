package v1

type OperationResult struct {
	Success bool
	Message string
}

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
	ComboGroup          string
	ComboStep           int32
	ComboWindowMs       int32
	ComboResetMs        int32
	Interruptible       bool
	Tags                []string
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

func (s *Skill) GetMovementLockMs() int32 {
	if s == nil {
		return 0
	}
	return s.MovementLockMs
}

type SkillHitboxProfile struct {
	Id            string
	SkillId       string
	HitboxShape   string
	HitboxStartMs int32
	HitboxEndMs   int32
	OffsetX       float64
	OffsetY       float64
	OffsetZ       float64
	Length        float64
	Radius        float64
	SizeX         float64
	SizeY         float64
	SizeZ         float64
	MotionProfile *SkillHitboxMotionProfile
	DamageGroupId string
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
	GuardDamageMultiplier float64
}

type CombatCoreProfile struct {
	DamageDealtMultiplier   float64
	CanBlock                bool
	BlockDamageReduction    float64
	MaxPosture              float64
	PostureDamageMultiplier float64
	PostureBreakDurationMs  int32
}

type StatusEffect struct {
	Id         string
	DurationMs int32
}

type CreatureTemplate struct {
	Id string
}

type ItemTemplate struct {
	Id string
}
