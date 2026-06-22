package postgres

import (
	"context"
	"database/sql"
	"time"

	"db-apeiron/internal/database"
)

type InventoryRepository struct {
	db database.TxManager
}

func NewInventoryRepository(db database.TxManager) *InventoryRepository {
	return &InventoryRepository{db: db}
}

//
// =========================
// INVENTORY
// =========================
//

func (r *InventoryRepository) GetInventoryByID(ctx context.Context, id string) (Inventory, error) {
	var inv Inventory

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			owner_type,
			owner_id,
			inventory_type,
			max_slots,
			max_weight,
			current_weight,
			is_locked,
			created_at,
			updated_at
		FROM apeiron.inventory
		WHERE id = $1
	`, id).Scan(
		&inv.ID,
		&inv.OwnerType,
		&inv.OwnerID,
		&inv.InventoryType,
		&inv.MaxSlots,
		&inv.MaxWeight,
		&inv.CurrentWeight,
		&inv.IsLocked,
		&inv.CreatedAt,
		&inv.UpdatedAt,
	)

	return inv, err
}

func (r *InventoryRepository) GetInventoryByOwner(
	ctx context.Context,
	ownerType string,
	ownerID string,
	inventoryType string,
) (Inventory, error) {
	var inv Inventory

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			owner_type,
			owner_id,
			inventory_type,
			max_slots,
			max_weight,
			current_weight,
			is_locked,
			created_at,
			updated_at
		FROM apeiron.inventory
		WHERE owner_type = $1
		  AND owner_id = $2
		  AND inventory_type = $3
	`, ownerType, ownerID, inventoryType).Scan(
		&inv.ID,
		&inv.OwnerType,
		&inv.OwnerID,
		&inv.InventoryType,
		&inv.MaxSlots,
		&inv.MaxWeight,
		&inv.CurrentWeight,
		&inv.IsLocked,
		&inv.CreatedAt,
		&inv.UpdatedAt,
	)

	return inv, err
}

func (r *InventoryRepository) CreateInventory(ctx context.Context, inv *Inventory) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO apeiron.inventory (
			id,
			owner_type,
			owner_id,
			inventory_type,
			max_slots,
			max_weight,
			current_weight,
			is_locked,
			created_at,
			updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`,
		inv.ID,
		inv.OwnerType,
		inv.OwnerID,
		inv.InventoryType,
		inv.MaxSlots,
		inv.MaxWeight,
		inv.CurrentWeight,
		inv.IsLocked,
		inv.CreatedAt,
		inv.UpdatedAt,
	)

	return err
}

func (r *InventoryRepository) UpdateInventoryWeight(ctx context.Context, id string, currentWeight float64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.inventory
		SET
			current_weight = $1,
			updated_at = NOW()
		WHERE id = $2
	`, currentWeight, id)

	return err
}

func (r *InventoryRepository) SetInventoryLocked(ctx context.Context, id string, locked bool) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.inventory
		SET
			is_locked = $1,
			updated_at = NOW()
		WHERE id = $2
	`, locked, id)

	return err
}

//
// =========================
// ITEM TEMPLATE
// =========================
//

func (r *InventoryRepository) GetItemTemplateByID(ctx context.Context, id string) (ItemTemplate, error) {
	var item ItemTemplate

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			name,
			description,
			item_type,
			rarity,
			max_stack,
			weight,
			is_tradable,
			is_destroyable,
			base_value,
			created_at,
			updated_at
		FROM apeiron.item_template
		WHERE id = $1
	`, id).Scan(
		&item.ID,
		&item.Name,
		&item.Description,
		&item.ItemType,
		&item.Rarity,
		&item.MaxStack,
		&item.Weight,
		&item.IsTradable,
		&item.IsDestroyable,
		&item.BaseValue,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	return item, err
}

func (r *InventoryRepository) ListItemTemplatesByType(ctx context.Context, itemType string) ([]ItemTemplate, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			name,
			description,
			item_type,
			rarity,
			max_stack,
			weight,
			is_tradable,
			is_destroyable,
			base_value,
			created_at,
			updated_at
		FROM apeiron.item_template
		WHERE item_type = $1
		ORDER BY id
	`, itemType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []ItemTemplate

	for rows.Next() {
		var item ItemTemplate

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.ItemType,
			&item.Rarity,
			&item.MaxStack,
			&item.Weight,
			&item.IsTradable,
			&item.IsDestroyable,
			&item.BaseValue,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, item)
	}

	return out, rows.Err()
}

func (r *InventoryRepository) CreateItemTemplate(ctx context.Context, item *ItemTemplate) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO apeiron.item_template (
			id,
			name,
			description,
			item_type,
			rarity,
			max_stack,
			weight,
			is_tradable,
			is_destroyable,
			base_value,
			created_at,
			updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
	`,
		item.ID,
		item.Name,
		item.Description,
		item.ItemType,
		item.Rarity,
		item.MaxStack,
		item.Weight,
		item.IsTradable,
		item.IsDestroyable,
		item.BaseValue,
		item.CreatedAt,
		item.UpdatedAt,
	)

	return err
}

//
// =========================
// INVENTORY ITEM
// =========================
//

func (r *InventoryRepository) GetItemByID(ctx context.Context, id string) (InventoryItem, error) {
	var item InventoryItem

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			inventory_id,
			item_id,
			slot_index,
			quantity,
			durability,
			is_equipped,
			created_at,
			updated_at
		FROM apeiron.inventory_item
		WHERE id = $1
	`, id).Scan(
		&item.ID,
		&item.InventoryID,
		&item.ItemID,
		&item.SlotIndex,
		&item.Quantity,
		&item.Durability,
		&item.IsEquipped,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	return item, err
}

func (r *InventoryRepository) GetItemsByInventoryID(ctx context.Context, inventoryID string) ([]InventoryItem, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			inventory_id,
			item_id,
			slot_index,
			quantity,
			durability,
			is_equipped,
			created_at,
			updated_at
		FROM apeiron.inventory_item
		WHERE inventory_id = $1
		ORDER BY slot_index ASC
	`, inventoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []InventoryItem

	for rows.Next() {
		var item InventoryItem

		if err := rows.Scan(
			&item.ID,
			&item.InventoryID,
			&item.ItemID,
			&item.SlotIndex,
			&item.Quantity,
			&item.Durability,
			&item.IsEquipped,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, item)
	}

	return out, rows.Err()
}

func (r *InventoryRepository) GetInventoryWithItems(ctx context.Context, inventoryID string) (InventoryWithItems, error) {
	inv, err := r.GetInventoryByID(ctx, inventoryID)
	if err != nil {
		return InventoryWithItems{}, err
	}

	items, err := r.GetItemsByInventoryID(ctx, inventoryID)
	if err != nil {
		return InventoryWithItems{}, err
	}

	return InventoryWithItems{
		Inventory: inv,
		Items:     items,
	}, nil
}

func (r *InventoryRepository) CreateItem(ctx context.Context, item *InventoryItem) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO apeiron.inventory_item (
			id,
			inventory_id,
			item_id,
			slot_index,
			quantity,
			durability,
			is_equipped,
			created_at,
			updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`,
		item.ID,
		item.InventoryID,
		item.ItemID,
		item.SlotIndex,
		item.Quantity,
		item.Durability,
		item.IsEquipped,
		item.CreatedAt,
		item.UpdatedAt,
	)

	return err
}

func (r *InventoryRepository) UpdateQuantity(ctx context.Context, id string, quantity int) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.inventory_item
		SET
			quantity = $1,
			updated_at = NOW()
		WHERE id = $2
	`, quantity, id)

	return err
}

func (r *InventoryRepository) MoveSlot(ctx context.Context, id string, slotIndex int) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.inventory_item
		SET
			slot_index = $1,
			updated_at = NOW()
		WHERE id = $2
	`, slotIndex, id)

	return err
}

func (r *InventoryRepository) UpdateDurability(ctx context.Context, id string, durability sql.NullFloat64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.inventory_item
		SET
			durability = $1,
			updated_at = NOW()
		WHERE id = $2
	`, durability, id)

	return err
}

func (r *InventoryRepository) SetEquipped(ctx context.Context, id string, equipped bool) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.inventory_item
		SET
			is_equipped = $1,
			updated_at = NOW()
		WHERE id = $2
	`, equipped, id)

	return err
}

func (r *InventoryRepository) DeleteItem(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM apeiron.inventory_item
		WHERE id = $1
	`, id)

	return err
}

//
// =========================
// MODELS
// =========================
//

type Inventory struct {
	ID            string
	OwnerType     string
	OwnerID       string
	InventoryType string

	MaxSlots      int
	MaxWeight     sql.NullFloat64
	CurrentWeight float64
	IsLocked      bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ItemTemplate struct {
	ID          string
	Name        string
	Description string

	ItemType string
	Rarity   string

	MaxStack int
	Weight   float64

	IsTradable    bool
	IsDestroyable bool

	BaseValue int64

	CreatedAt time.Time
	UpdatedAt time.Time
}

type InventoryItem struct {
	ID          string
	InventoryID string
	ItemID      string

	SlotIndex int
	Quantity  int

	Durability sql.NullFloat64
	IsEquipped bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type InventoryWithItems struct {
	Inventory Inventory
	Items     []InventoryItem
}
