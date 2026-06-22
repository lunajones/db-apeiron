package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/repository/postgres"
)

type InventoryReader interface {
	GetInventoryByID(ctx context.Context, id string) (postgres.Inventory, error)
	GetInventoryByOwner(ctx context.Context, ownerType string, ownerID string, inventoryType string) (postgres.Inventory, error)
	GetInventoryWithItems(ctx context.Context, inventoryID string) (postgres.InventoryWithItems, error)
	GetItemTemplateByID(ctx context.Context, id string) (postgres.ItemTemplate, error)
	ListItemTemplatesByType(ctx context.Context, itemType string) ([]postgres.ItemTemplate, error)
	GetItemByID(ctx context.Context, id string) (postgres.InventoryItem, error)
	GetItemsByInventoryID(ctx context.Context, inventoryID string) ([]postgres.InventoryItem, error)
}

type InventoryDataHandler struct {
	apeironv1.UnimplementedInventoryDataServiceServer

	inventory InventoryReader
}

func NewInventoryDataHandler(inventory InventoryReader) *InventoryDataHandler {
	return &InventoryDataHandler{inventory: inventory}
}

func (h *InventoryDataHandler) GetInventory(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.InventoryResponse, error) {
	inventory, err := h.inventory.GetInventoryByID(ctx, req.GetId())
	if err != nil {
		return &apeironv1.InventoryResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.InventoryResponse{Found: true, Inventory: mapInventory(inventory)}, nil
}

func (h *InventoryDataHandler) GetInventoryByOwner(ctx context.Context, req *apeironv1.InventoryOwnerRequest) (*apeironv1.InventoryResponse, error) {
	inventory, err := h.inventory.GetInventoryByOwner(ctx, req.GetOwnerType(), req.GetOwnerId(), req.GetInventoryType())
	if err != nil {
		return &apeironv1.InventoryResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.InventoryResponse{Found: true, Inventory: mapInventory(inventory)}, nil
}

func (h *InventoryDataHandler) GetInventoryWithItems(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.InventoryWithItemsResponse, error) {
	inventory, err := h.inventory.GetInventoryWithItems(ctx, req.GetId())
	if err != nil {
		return &apeironv1.InventoryWithItemsResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.InventoryWithItemsResponse{Found: true, Inventory: mapInventoryWithItems(inventory)}, nil
}

func (h *InventoryDataHandler) GetItemTemplate(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.ItemTemplateResponse, error) {
	item, err := h.inventory.GetItemTemplateByID(ctx, req.GetId())
	if err != nil {
		return &apeironv1.ItemTemplateResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.ItemTemplateResponse{Found: true, ItemTemplate: mapItemTemplate(item)}, nil
}

func (h *InventoryDataHandler) ListItemTemplatesByType(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.ItemTemplatesResponse, error) {
	items, err := h.inventory.ListItemTemplatesByType(ctx, req.GetId())
	if err != nil {
		return &apeironv1.ItemTemplatesResponse{Error: err.Error()}, nil
	}

	out := make([]*apeironv1.ItemTemplate, 0, len(items))
	for _, item := range items {
		out = append(out, mapItemTemplate(item))
	}

	return &apeironv1.ItemTemplatesResponse{ItemTemplates: out}, nil
}

func (h *InventoryDataHandler) GetItem(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.InventoryItemResponse, error) {
	item, err := h.inventory.GetItemByID(ctx, req.GetId())
	if err != nil {
		return &apeironv1.InventoryItemResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.InventoryItemResponse{Found: true, Item: mapInventoryItem(item)}, nil
}

func (h *InventoryDataHandler) GetItemsByInventory(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.InventoryItemsResponse, error) {
	items, err := h.inventory.GetItemsByInventoryID(ctx, req.GetId())
	if err != nil {
		return &apeironv1.InventoryItemsResponse{Error: err.Error()}, nil
	}

	out := make([]*apeironv1.InventoryItem, 0, len(items))
	for _, item := range items {
		out = append(out, mapInventoryItem(item))
	}

	return &apeironv1.InventoryItemsResponse{Items: out}, nil
}

func mapInventory(inventory postgres.Inventory) *apeironv1.Inventory {
	return &apeironv1.Inventory{
		Id:            inventory.ID,
		OwnerType:     inventory.OwnerType,
		OwnerId:       inventory.OwnerID,
		InventoryType: inventory.InventoryType,
		MaxSlots:      int32(inventory.MaxSlots),
		MaxWeight:     nullFloat64(inventory.MaxWeight),
		CurrentWeight: inventory.CurrentWeight,
		IsLocked:      inventory.IsLocked,
	}
}

func mapItemTemplate(item postgres.ItemTemplate) *apeironv1.ItemTemplate {
	return &apeironv1.ItemTemplate{
		Id:            item.ID,
		Name:          item.Name,
		Description:   item.Description,
		ItemType:      item.ItemType,
		Rarity:        item.Rarity,
		MaxStack:      int32(item.MaxStack),
		Weight:        item.Weight,
		IsTradable:    item.IsTradable,
		IsDestroyable: item.IsDestroyable,
		BaseValue:     item.BaseValue,
	}
}

func mapInventoryItem(item postgres.InventoryItem) *apeironv1.InventoryItem {
	return &apeironv1.InventoryItem{
		Id:          item.ID,
		InventoryId: item.InventoryID,
		ItemId:      item.ItemID,
		SlotIndex:   int32(item.SlotIndex),
		Quantity:    int32(item.Quantity),
		Durability:  nullFloat64(item.Durability),
		IsEquipped:  item.IsEquipped,
	}
}

func mapInventoryWithItems(inventory postgres.InventoryWithItems) *apeironv1.InventoryWithItems {
	items := make([]*apeironv1.InventoryItem, 0, len(inventory.Items))
	for _, item := range inventory.Items {
		items = append(items, mapInventoryItem(item))
	}

	return &apeironv1.InventoryWithItems{
		Inventory: mapInventory(inventory.Inventory),
		Items:     items,
	}
}
