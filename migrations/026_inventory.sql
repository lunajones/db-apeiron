-- =========================================================
-- INVENTORY
-- APEIRON MMO - INVENTORY CONTAINER
-- =========================================================

CREATE TABLE apeiron.inventory (
    id TEXT PRIMARY KEY,

    -- =========================
    -- OWNER
    -- =========================

    owner_type TEXT NOT NULL,
    -- player | creature | chest | bank | loot_container

    owner_id TEXT NOT NULL,
    -- id da entidade dona do inventário

    -- =========================
    -- INVENTORY CONFIG
    -- =========================

    inventory_type TEXT NOT NULL DEFAULT 'backpack',
    -- backpack | equipment | bank | loot | storage

    max_slots INT NOT NULL DEFAULT 30,

    max_weight FLOAT,
    current_weight FLOAT NOT NULL DEFAULT 0.0,

    is_locked BOOLEAN NOT NULL DEFAULT FALSE,

    -- =========================
    -- LIFECYCLE
    -- =========================

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT uq_inventory_owner_type_owner_id_inventory_type
        UNIQUE (owner_type, owner_id, inventory_type),

    CONSTRAINT chk_inventory_owner_type
        CHECK (owner_type IN (
            'player',
            'creature',
            'chest',
            'bank',
            'loot_container'
        )),

    CONSTRAINT chk_inventory_type
        CHECK (inventory_type IN (
            'backpack',
            'equipment',
            'bank',
            'loot',
            'storage'
        )),

    CONSTRAINT chk_inventory_max_slots
        CHECK (max_slots > 0),

    CONSTRAINT chk_inventory_max_weight
        CHECK (max_weight IS NULL OR max_weight >= 0),

    CONSTRAINT chk_inventory_current_weight
        CHECK (current_weight >= 0)
);

CREATE INDEX idx_inventory_owner
ON apeiron.inventory(owner_type, owner_id);

CREATE INDEX idx_inventory_type
ON apeiron.inventory(inventory_type);

-- =========================================================
-- ITEM TEMPLATE
-- APEIRON MMO - STATIC ITEM DEFINITION
-- =========================================================

CREATE TABLE apeiron.item_template (
    id TEXT PRIMARY KEY,

    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',

    item_type TEXT NOT NULL,
    -- weapon | armor | consumable | material | quest | currency | misc

    rarity TEXT NOT NULL DEFAULT 'common',
    -- common | uncommon | rare | epic | legendary | unique

    max_stack INT NOT NULL DEFAULT 1,

    weight FLOAT NOT NULL DEFAULT 0.0,

    is_tradable BOOLEAN NOT NULL DEFAULT TRUE,
    is_destroyable BOOLEAN NOT NULL DEFAULT TRUE,

    base_value BIGINT NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_item_template_item_type
        CHECK (item_type IN (
            'weapon',
            'armor',
            'consumable',
            'material',
            'quest',
            'currency',
            'misc'
        )),

    CONSTRAINT chk_item_template_rarity
        CHECK (rarity IN (
            'common',
            'uncommon',
            'rare',
            'epic',
            'legendary',
            'unique'
        )),

    CONSTRAINT chk_item_template_max_stack
        CHECK (max_stack >= 1),

    CONSTRAINT chk_item_template_weight
        CHECK (weight >= 0),

    CONSTRAINT chk_item_template_base_value
        CHECK (base_value >= 0)
);

CREATE INDEX idx_item_template_item_type
ON apeiron.item_template(item_type);

CREATE INDEX idx_item_template_rarity
ON apeiron.item_template(rarity);

-- =========================================================
-- INVENTORY ITEM
-- APEIRON MMO - ITEM INSTANCE
-- =========================================================

CREATE TABLE apeiron.inventory_item (
    id TEXT PRIMARY KEY,

    inventory_id TEXT NOT NULL,

    item_id TEXT NOT NULL,
    -- referência para item_template

    slot_index INT NOT NULL,
    -- posição no inventário

    quantity INT NOT NULL DEFAULT 1,

    durability FLOAT DEFAULT 1.0,

    is_equipped BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_inventory_item_inventory
        FOREIGN KEY (inventory_id)
        REFERENCES apeiron.inventory(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_inventory_item_item_template
        FOREIGN KEY (item_id)
        REFERENCES apeiron.item_template(id),

    CONSTRAINT uq_inventory_slot
        UNIQUE (inventory_id, slot_index),

    CONSTRAINT chk_inventory_item_slot_index
        CHECK (slot_index >= 0),

    CONSTRAINT chk_inventory_item_quantity
        CHECK (quantity > 0),

    CONSTRAINT chk_inventory_item_durability
        CHECK (durability IS NULL OR durability >= 0)
);

CREATE INDEX idx_inventory_item_inventory
ON apeiron.inventory_item(inventory_id);

CREATE INDEX idx_inventory_item_item_template
ON apeiron.inventory_item(item_id);