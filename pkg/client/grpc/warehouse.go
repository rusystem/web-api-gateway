package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"github.com/rusystem/web-api-gateway/proto/crm_warehouse/warehouse"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WarehouseClient struct {
	conn            *grpc.ClientConn
	warehouseClient warehouse.WarehouseServiceClient
}

func NewWarehouseClient(addr string) (*WarehouseClient, error) {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opt...)
	if err != nil {
		return nil, err
	}

	return &WarehouseClient{
		conn:            conn,
		warehouseClient: warehouse.NewWarehouseServiceClient(conn),
	}, nil
}

func (w *WarehouseClient) Close() error {
	return w.conn.Close()
}

func (w *WarehouseClient) GetById(ctx context.Context, id int64) (domain.Warehouse, error) {
	if id <= 0 {
		return domain.Warehouse{}, errors.New("calls grpc: id can`t be zero")
	}

	resp, err := w.warehouseClient.GetById(ctx, &warehouse.Id{Id: id})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return domain.Warehouse{}, domain.ErrWarehouseNotFound
		}

		return domain.Warehouse{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return domain.Warehouse{}, err
	}

	return domain.Warehouse{
		ID:                resp.Id,
		Name:              resp.Name,
		Address:           resp.Address,
		ResponsiblePerson: resp.ResponsiblePerson,
		Phone:             resp.Phone,
		Email:             resp.Email,
		MaxCapacity:       resp.MaxCapacity,
		CurrentOccupancy:  resp.CurrentOccupancy,
		OtherFields:       otherFields,
		Country:           resp.Country,
	}, nil
}

func (w *WarehouseClient) Create(ctx context.Context, wh domain.Warehouse) (int64, error) {
	otherFieldsJSON, err := json.Marshal(wh.OtherFields)
	if err != nil {
		return 0, err
	}

	resp, err := w.warehouseClient.Create(ctx, &warehouse.Warehouse{
		Id:                wh.ID,
		Name:              wh.Name,
		Address:           wh.Address,
		ResponsiblePerson: wh.ResponsiblePerson,
		Phone:             wh.Phone,
		Email:             wh.Email,
		MaxCapacity:       wh.MaxCapacity,
		CurrentOccupancy:  wh.CurrentOccupancy,
		OtherFields:       string(otherFieldsJSON),
		Country:           wh.Country,
	})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}
