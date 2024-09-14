package warehouse

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"github.com/rusystem/web-api-gateway/proto/crm_warehouse/materials"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MaterialsClient struct {
	conn            *grpc.ClientConn
	materialsClient materials.MaterialServiceClient
}

func NewMaterialsClient(addr string) (*MaterialsClient, error) {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opt...)
	if err != nil {
		return nil, err
	}

	return &MaterialsClient{
		conn:            conn,
		materialsClient: materials.NewMaterialServiceClient(conn),
	}, nil
}

func (mc *MaterialsClient) Close() error {
	return mc.conn.Close()
}

func (mc *MaterialsClient) CreatePlanning(ctx context.Context, material domain.Material) (int64, error) {
	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return 0, err
	}

	resp, err := mc.materialsClient.CreatePlanning(ctx, &materials.Material{
		WarehouseId:            material.WarehouseID,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVat:        material.PriceWithoutVAT,
		TotalWithoutVat:        material.TotalWithoutVAT,
		SupplierId:             material.SupplierID,
		Location:               material.Location,
		Contract:               timestamppb.New(material.Contract),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           timestamppb.New(material.ReceivedDate),
		LastUpdated:            timestamppb.New(material.LastUpdated),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         timestamppb.New(material.ExpirationDate),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            string(otherFieldsJSON),
		CompanyId:              material.CompanyID,
	})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = failed to insert planning material: pq: insert or update on table \"planning_materials\" violates foreign key constraint \"planning_materials_warehouse_id_fkey\"" {
			return 0, domain.ErrWarehouseNotFound
		} //todo make it better

		if err.Error() == "rpc error: code = Unknown desc = failed to insert planning material: pq: insert or update on table \"planning_materials\" violates foreign key constraint \"planning_materials_supplier_id_fkey\"" {
			return 0, domain.ErrSupplierNotFound
		} //todo make it better

		return 0, err
	}

	return resp.Id, nil
}

func (mc *MaterialsClient) UpdatePlanningById(ctx context.Context, material domain.Material) error {
	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return err
	}

	_, err = mc.materialsClient.UpdatePlanning(ctx, &materials.Material{
		Id:                     material.ID,
		WarehouseId:            material.WarehouseID,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVat:        material.PriceWithoutVAT,
		TotalWithoutVat:        material.TotalWithoutVAT,
		SupplierId:             material.SupplierID,
		Location:               material.Location,
		Contract:               timestamppb.New(material.Contract),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           timestamppb.New(material.ReceivedDate),
		LastUpdated:            timestamppb.New(material.LastUpdated),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         timestamppb.New(material.ExpirationDate),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            string(otherFieldsJSON),
	})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = pq: insert or update on table \"planning_materials\" violates foreign key constraint \"planning_materials_warehouse_id_fkey\"" {
			return domain.ErrWarehouseNotFound
		} //todo make it better

		if err.Error() == "rpc error: code = Unknown desc = pq: insert or update on table \"planning_materials\" violates foreign key constraint \"planning_materials_supplier_id_fkey\"" {
			return domain.ErrSupplierNotFound
		} //todo make it better

		return err
	}

	return nil
}

func (mc *MaterialsClient) DeletePlanningById(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.DeletePlanning(ctx, &materials.MaterialId{Id: id})
	return err
}

func (mc *MaterialsClient) GetPlanningById(ctx context.Context, id int64) (domain.Material, error) {
	if id <= 0 {
		return domain.Material{}, errors.New("materials, grpc client - invalid id")
	}

	resp, err := mc.materialsClient.GetPlanning(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
			return domain.Material{}, domain.ErrMaterialNotFound
		} //todo обработать корректно такие ошибки

		return domain.Material{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return domain.Material{}, err
	}

	return domain.Material{
		ID:                     resp.Id,
		WarehouseID:            resp.WarehouseId,
		ItemID:                 resp.ItemId,
		Name:                   resp.Name,
		ByInvoice:              resp.ByInvoice,
		Article:                resp.Article,
		ProductCategory:        resp.ProductCategory,
		Unit:                   resp.Unit,
		TotalQuantity:          resp.TotalQuantity,
		Volume:                 resp.Volume,
		PriceWithoutVAT:        resp.PriceWithoutVat,
		TotalWithoutVAT:        resp.TotalWithoutVat,
		SupplierID:             resp.SupplierId,
		Location:               resp.Location,
		Contract:               resp.Contract.AsTime(),
		File:                   resp.File,
		Status:                 resp.Status,
		Comments:               resp.Comments,
		Reserve:                resp.Reserve,
		ReceivedDate:           resp.ReceivedDate.AsTime(),
		LastUpdated:            resp.LastUpdated.AsTime(),
		MinStockLevel:          resp.MinStockLevel,
		ExpirationDate:         resp.ExpirationDate.AsTime(),
		ResponsiblePerson:      resp.ResponsiblePerson,
		StorageCost:            resp.StorageCost,
		WarehouseSection:       resp.WarehouseSection,
		IncomingDeliveryNumber: resp.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              resp.CompanyId,
	}, nil
}

func (mc *MaterialsClient) GetListPlanning(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	var mtrls []domain.Material

	resp, err := mc.materialsClient.GetListPlanning(ctx, &materials.MaterialParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		CompanyId: params.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	for _, mtrl := range resp.Materials {
		var otherFields map[string]interface{}
		if err = json.Unmarshal([]byte(mtrl.OtherFields), &otherFields); err != nil {
			return nil, err
		}

		mtrls = append(mtrls, domain.Material{
			ID:                     mtrl.Id,
			WarehouseID:            mtrl.WarehouseId,
			ItemID:                 mtrl.ItemId,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVAT:        mtrl.PriceWithoutVat,
			TotalWithoutVAT:        mtrl.TotalWithoutVat,
			SupplierID:             mtrl.SupplierId,
			Location:               mtrl.Location,
			Contract:               mtrl.Contract.AsTime(),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           mtrl.ReceivedDate.AsTime(),
			LastUpdated:            mtrl.LastUpdated.AsTime(),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         mtrl.ExpirationDate.AsTime(),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            otherFields,
			CompanyID:              mtrl.CompanyId,
		})
	}

	return mtrls, nil
}

func (mc *MaterialsClient) MovePlanningToPurchased(ctx context.Context, id int64) (int64, int64, error) {
	resp, err := mc.materialsClient.MovePlanningToPurchased(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		return 0, 0, err
	}

	return resp.Id, resp.ItemId, nil
}

func (mc *MaterialsClient) CreatePurchased(ctx context.Context, material domain.Material) (int64, int64, error) {
	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return 0, 0, err
	}

	resp, err := mc.materialsClient.CreatePurchased(ctx, &materials.Material{
		WarehouseId:            material.WarehouseID,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVat:        material.PriceWithoutVAT,
		TotalWithoutVat:        material.TotalWithoutVAT,
		SupplierId:             material.SupplierID,
		Location:               material.Location,
		Contract:               timestamppb.New(material.Contract),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           timestamppb.New(material.ReceivedDate),
		LastUpdated:            timestamppb.New(material.LastUpdated),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         timestamppb.New(material.ExpirationDate),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            string(otherFieldsJSON),
		CompanyId:              material.CompanyID,
	})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = failed to insert purchased material: pq: insert or update on table \"purchased_materials\" violates foreign key constraint \"purchased_materials_warehouse_id_fkey\"" {
			return 0, 0, domain.ErrWarehouseNotFound
		} //todo make it better

		if err.Error() == "rpc error: code = Unknown desc = failed to insert purchased material: pq: insert or update on table \"purchased_materials\" violates foreign key constraint \"purchased_materials_supplier_id_fkey\"" {
			return 0, 0, domain.ErrSupplierNotFound
		} //todo make it better

		return 0, 0, err
	}

	return resp.Id, resp.ItemId, nil
}

func (mc *MaterialsClient) UpdatePurchasedById(ctx context.Context, material domain.Material) error {
	otherFieldsJSON, err := json.Marshal(material.OtherFields)
	if err != nil {
		return err
	}

	_, err = mc.materialsClient.UpdatePurchased(ctx, &materials.Material{
		Id:                     material.ID,
		WarehouseId:            material.WarehouseID,
		Name:                   material.Name,
		ByInvoice:              material.ByInvoice,
		Article:                material.Article,
		ProductCategory:        material.ProductCategory,
		Unit:                   material.Unit,
		TotalQuantity:          material.TotalQuantity,
		Volume:                 material.Volume,
		PriceWithoutVat:        material.PriceWithoutVAT,
		TotalWithoutVat:        material.TotalWithoutVAT,
		SupplierId:             material.SupplierID,
		Location:               material.Location,
		Contract:               timestamppb.New(material.Contract),
		File:                   material.File,
		Status:                 material.Status,
		Comments:               material.Comments,
		Reserve:                material.Reserve,
		ReceivedDate:           timestamppb.New(material.ReceivedDate),
		LastUpdated:            timestamppb.New(material.LastUpdated),
		MinStockLevel:          material.MinStockLevel,
		ExpirationDate:         timestamppb.New(material.ExpirationDate),
		ResponsiblePerson:      material.ResponsiblePerson,
		StorageCost:            material.StorageCost,
		WarehouseSection:       material.WarehouseSection,
		IncomingDeliveryNumber: material.IncomingDeliveryNumber,
		OtherFields:            string(otherFieldsJSON),
	})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = pq: insert or update on table \"purchased_materials\" violates foreign key constraint \"purchased_materials_warehouse_id_fkey\"" {
			return domain.ErrWarehouseNotFound
		} //todo make it better

		if err.Error() == "rpc error: code = Unknown desc = pq: insert or update on table \"purchased_materials\" violates foreign key constraint \"purchased_materials_supplier_id_fkey\"" {
			return domain.ErrSupplierNotFound
		} //todo make it better
		return err
	}

	return nil
}

func (mc *MaterialsClient) DeletePurchasedById(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.DeletePurchased(ctx, &materials.MaterialId{Id: id})
	return err
}

func (mc *MaterialsClient) GetPurchasedById(ctx context.Context, id int64) (domain.Material, error) {
	if id <= 0 {
		return domain.Material{}, errors.New("materials, grpc client - invalid id")
	}

	resp, err := mc.materialsClient.GetPurchased(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
			return domain.Material{}, domain.ErrMaterialNotFound
		}

		return domain.Material{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return domain.Material{}, err
	}

	return domain.Material{
		ID:                     resp.Id,
		WarehouseID:            resp.WarehouseId,
		ItemID:                 resp.ItemId,
		Name:                   resp.Name,
		ByInvoice:              resp.ByInvoice,
		Article:                resp.Article,
		ProductCategory:        resp.ProductCategory,
		Unit:                   resp.Unit,
		TotalQuantity:          resp.TotalQuantity,
		Volume:                 resp.Volume,
		PriceWithoutVAT:        resp.PriceWithoutVat,
		TotalWithoutVAT:        resp.TotalWithoutVat,
		SupplierID:             resp.SupplierId,
		Location:               resp.Location,
		Contract:               resp.Contract.AsTime(),
		File:                   resp.File,
		Status:                 resp.Status,
		Comments:               resp.Comments,
		Reserve:                resp.Reserve,
		ReceivedDate:           resp.ReceivedDate.AsTime(),
		LastUpdated:            resp.LastUpdated.AsTime(),
		MinStockLevel:          resp.MinStockLevel,
		ExpirationDate:         resp.ExpirationDate.AsTime(),
		ResponsiblePerson:      resp.ResponsiblePerson,
		StorageCost:            resp.StorageCost,
		WarehouseSection:       resp.WarehouseSection,
		IncomingDeliveryNumber: resp.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              resp.CompanyId,
	}, nil
}

func (mc *MaterialsClient) GetListPurchased(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	var mtrls []domain.Material

	resp, err := mc.materialsClient.GetListPurchased(ctx, &materials.MaterialParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		CompanyId: params.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	for _, mtrl := range resp.Materials {
		var otherFields map[string]interface{}
		if err = json.Unmarshal([]byte(mtrl.OtherFields), &otherFields); err != nil {
			return nil, err
		}

		mtrls = append(mtrls, domain.Material{
			ID:                     mtrl.Id,
			WarehouseID:            mtrl.WarehouseId,
			ItemID:                 mtrl.ItemId,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVAT:        mtrl.PriceWithoutVat,
			TotalWithoutVAT:        mtrl.TotalWithoutVat,
			SupplierID:             mtrl.SupplierId,
			Location:               mtrl.Location,
			Contract:               mtrl.Contract.AsTime(),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           mtrl.ReceivedDate.AsTime(),
			LastUpdated:            mtrl.LastUpdated.AsTime(),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         mtrl.ExpirationDate.AsTime(),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            otherFields,
			CompanyID:              mtrl.CompanyId,
		})
	}

	return mtrls, nil
}

func (mc *MaterialsClient) MovePurchasedToArchive(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.MovePurchasedToArchive(ctx, &materials.MaterialId{Id: id})
	return err
}

func (mc *MaterialsClient) GetPlanningArchiveById(ctx context.Context, id int64) (domain.Material, error) {
	if id <= 0 {
		return domain.Material{}, errors.New("materials, grpc client - invalid id")
	}

	resp, err := mc.materialsClient.GetPlanningArchive(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
			return domain.Material{}, domain.ErrMaterialNotFound
		}

		return domain.Material{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return domain.Material{}, err
	}

	return domain.Material{
		ID:                     resp.Id,
		WarehouseID:            resp.WarehouseId,
		ItemID:                 resp.ItemId,
		Name:                   resp.Name,
		ByInvoice:              resp.ByInvoice,
		Article:                resp.Article,
		ProductCategory:        resp.ProductCategory,
		Unit:                   resp.Unit,
		TotalQuantity:          resp.TotalQuantity,
		Volume:                 resp.Volume,
		PriceWithoutVAT:        resp.PriceWithoutVat,
		TotalWithoutVAT:        resp.TotalWithoutVat,
		SupplierID:             resp.SupplierId,
		Location:               resp.Location,
		Contract:               resp.Contract.AsTime(),
		File:                   resp.File,
		Status:                 resp.Status,
		Comments:               resp.Comments,
		Reserve:                resp.Reserve,
		ReceivedDate:           resp.ReceivedDate.AsTime(),
		LastUpdated:            resp.LastUpdated.AsTime(),
		MinStockLevel:          resp.MinStockLevel,
		ExpirationDate:         resp.ExpirationDate.AsTime(),
		ResponsiblePerson:      resp.ResponsiblePerson,
		StorageCost:            resp.StorageCost,
		WarehouseSection:       resp.WarehouseSection,
		IncomingDeliveryNumber: resp.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              resp.CompanyId,
	}, nil
}

func (mc *MaterialsClient) GetPurchasedArchiveById(ctx context.Context, id int64) (domain.Material, error) {
	if id <= 0 {
		return domain.Material{}, errors.New("materials, grpc client - invalid id")
	}

	resp, err := mc.materialsClient.GetPurchasedArchive(ctx, &materials.MaterialId{Id: id})
	if err != nil {
		if err.Error() == "rpc error: code = Unknown desc = sql: no rows in result set" {
			return domain.Material{}, domain.ErrMaterialNotFound
		}

		return domain.Material{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return domain.Material{}, err
	}

	return domain.Material{
		ID:                     resp.Id,
		WarehouseID:            resp.WarehouseId,
		ItemID:                 resp.ItemId,
		Name:                   resp.Name,
		ByInvoice:              resp.ByInvoice,
		Article:                resp.Article,
		ProductCategory:        resp.ProductCategory,
		Unit:                   resp.Unit,
		TotalQuantity:          resp.TotalQuantity,
		Volume:                 resp.Volume,
		PriceWithoutVAT:        resp.PriceWithoutVat,
		TotalWithoutVAT:        resp.TotalWithoutVat,
		SupplierID:             resp.SupplierId,
		Location:               resp.Location,
		Contract:               resp.Contract.AsTime(),
		File:                   resp.File,
		Status:                 resp.Status,
		Comments:               resp.Comments,
		Reserve:                resp.Reserve,
		ReceivedDate:           resp.ReceivedDate.AsTime(),
		LastUpdated:            resp.LastUpdated.AsTime(),
		MinStockLevel:          resp.MinStockLevel,
		ExpirationDate:         resp.ExpirationDate.AsTime(),
		ResponsiblePerson:      resp.ResponsiblePerson,
		StorageCost:            resp.StorageCost,
		WarehouseSection:       resp.WarehouseSection,
		IncomingDeliveryNumber: resp.IncomingDeliveryNumber,
		OtherFields:            otherFields,
		CompanyID:              resp.CompanyId,
	}, nil
}

func (mc *MaterialsClient) GetListPlanningArchive(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	var mtrls []domain.Material

	resp, err := mc.materialsClient.GetListPlanningArchive(ctx, &materials.MaterialParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		CompanyId: params.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	for _, mtrl := range resp.Materials {
		var otherFields map[string]interface{}
		if err = json.Unmarshal([]byte(mtrl.OtherFields), &otherFields); err != nil {
			return nil, err
		}

		mtrls = append(mtrls, domain.Material{
			ID:                     mtrl.Id,
			WarehouseID:            mtrl.WarehouseId,
			ItemID:                 mtrl.ItemId,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVAT:        mtrl.PriceWithoutVat,
			TotalWithoutVAT:        mtrl.TotalWithoutVat,
			SupplierID:             mtrl.SupplierId,
			Location:               mtrl.Location,
			Contract:               mtrl.Contract.AsTime(),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           mtrl.ReceivedDate.AsTime(),
			LastUpdated:            mtrl.LastUpdated.AsTime(),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         mtrl.ExpirationDate.AsTime(),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            otherFields,
			CompanyID:              mtrl.CompanyId,
		})
	}

	return mtrls, nil
}

func (mc *MaterialsClient) GetListPurchasedArchive(ctx context.Context, params domain.MaterialParams) ([]domain.Material, error) {
	var mtrls []domain.Material

	resp, err := mc.materialsClient.GetListPurchasedArchive(ctx, &materials.MaterialParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		CompanyId: params.CompanyId,
	})
	if err != nil {
		return nil, err
	}

	for _, mtrl := range resp.Materials {
		var otherFields map[string]interface{}
		if err = json.Unmarshal([]byte(mtrl.OtherFields), &otherFields); err != nil {
			return nil, err
		}

		mtrls = append(mtrls, domain.Material{
			ID:                     mtrl.Id,
			WarehouseID:            mtrl.WarehouseId,
			ItemID:                 mtrl.ItemId,
			Name:                   mtrl.Name,
			ByInvoice:              mtrl.ByInvoice,
			Article:                mtrl.Article,
			ProductCategory:        mtrl.ProductCategory,
			Unit:                   mtrl.Unit,
			TotalQuantity:          mtrl.TotalQuantity,
			Volume:                 mtrl.Volume,
			PriceWithoutVAT:        mtrl.PriceWithoutVat,
			TotalWithoutVAT:        mtrl.TotalWithoutVat,
			SupplierID:             mtrl.SupplierId,
			Location:               mtrl.Location,
			Contract:               mtrl.Contract.AsTime(),
			File:                   mtrl.File,
			Status:                 mtrl.Status,
			Comments:               mtrl.Comments,
			Reserve:                mtrl.Reserve,
			ReceivedDate:           mtrl.ReceivedDate.AsTime(),
			LastUpdated:            mtrl.LastUpdated.AsTime(),
			MinStockLevel:          mtrl.MinStockLevel,
			ExpirationDate:         mtrl.ExpirationDate.AsTime(),
			ResponsiblePerson:      mtrl.ResponsiblePerson,
			StorageCost:            mtrl.StorageCost,
			WarehouseSection:       mtrl.WarehouseSection,
			IncomingDeliveryNumber: mtrl.IncomingDeliveryNumber,
			OtherFields:            otherFields,
			CompanyID:              mtrl.CompanyId,
		})
	}

	return mtrls, nil
}

func (mc *MaterialsClient) DeletePlanningArchiveById(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.DeletePlanningArchive(ctx, &materials.MaterialId{Id: id})
	return err
}

func (mc *MaterialsClient) DeletePurchasedArchiveById(ctx context.Context, id int64) error {
	_, err := mc.materialsClient.DeletePurchasedArchive(ctx, &materials.MaterialId{Id: id})
	return err
}
