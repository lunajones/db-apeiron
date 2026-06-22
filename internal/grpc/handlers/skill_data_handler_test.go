package handlers

import (
	"database/sql"
	"testing"

	"db-apeiron/internal/repository/postgres"
)

func TestMapSkillHitboxProfileIncludesTemporalMotionProfile(t *testing.T) {
	profile := postgres.SkillHitboxProfile{
		ID:            "hitbox_bite_0",
		SkillID:       "bite",
		HitboxIndex:   0,
		HitboxShape:   "temporal_sweep",
		HitboxStartMS: 120,
		HitboxEndMS:   340,
		DamageGroupID: sql.NullString{String: "wolf_bite_damage", Valid: true},
		MotionProfile: &postgres.SkillHitboxMotionProfile{
			ID:            "motion_wolf_bite_melee_v1",
			Enabled:       true,
			MotionType:    "timeline_sweep",
			TimeBasis:     "hitbox_window_normalized",
			Interpolation: "linear",
			SweepShape:    "capsule_strip",
			DamageGroupID: sql.NullString{String: "wolf_bite_damage", Valid: true},
			Samples: []postgres.SkillHitboxMotionSample{
				{SampleIndex: 0, T: 0.0, Shape: "capsule_strip", OffsetY: 45, Radius: 45, Length: 70},
				{SampleIndex: 1, T: 1.0, Shape: "capsule_strip", OffsetY: 95, Radius: 48, Length: 145},
			},
		},
	}

	out := mapSkillHitboxProfile(profile)
	if out.GetDamageGroupId() != "wolf_bite_damage" {
		t.Fatalf("damage group = %q", out.GetDamageGroupId())
	}
	if out.GetMotionProfile() == nil {
		t.Fatal("expected motion profile")
	}
	if out.GetMotionProfile().GetId() != "motion_wolf_bite_melee_v1" {
		t.Fatalf("motion profile id = %q", out.GetMotionProfile().GetId())
	}
	if out.GetMotionProfile().GetSweepShape() != "capsule_strip" {
		t.Fatalf("sweep shape = %q", out.GetMotionProfile().GetSweepShape())
	}
	if got := len(out.GetMotionProfile().GetSamples()); got != 2 {
		t.Fatalf("samples = %d", got)
	}
}
