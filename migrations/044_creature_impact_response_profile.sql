-- =========================================================
-- CREATURE IMPACT RESPONSE PROFILE
-- =========================================================
--
-- Final authority for target material / visual response profile.
-- Creature impact feedback must come from template data loaded through
-- DB/proto instead of runtime inference from entity kind alone.

ALTER TABLE apeiron.creature_template
    ADD COLUMN IF NOT EXISTS impact_response_profile TEXT NOT NULL DEFAULT 'creature_flesh_blood_red';

CREATE INDEX IF NOT EXISTS idx_creature_template_impact_response_profile
ON apeiron.creature_template(impact_response_profile);
