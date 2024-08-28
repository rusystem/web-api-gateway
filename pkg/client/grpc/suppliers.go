package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"github.com/rusystem/web-api-gateway/proto/crm_warehouse/supplier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SuppliersClient struct {
	conn           *grpc.ClientConn
	supplierClient supplier.SupplierServiceClient
}

func NewSuppliersClient(addr string) (*SuppliersClient, error) {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(addr, opt...)
	if err != nil {
		return nil, err
	}

	return &SuppliersClient{
		conn:           conn,
		supplierClient: supplier.NewSupplierServiceClient(conn),
	}, nil
}

func (s *SuppliersClient) Close() error {
	return s.conn.Close()
}

func (s *SuppliersClient) GetById(ctx context.Context, id int64) (domain.Supplier, error) {
	if id <= 0 {
		return domain.Supplier{}, errors.New("calls grpc: id can`t be zero")
	}

	resp, err := s.supplierClient.GetById(ctx, &supplier.Id{Id: id})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return domain.Supplier{}, domain.ErrSupplierNotFound
		}

		return domain.Supplier{}, err
	}

	var otherFields map[string]interface{}
	if err = json.Unmarshal([]byte(resp.OtherFields), &otherFields); err != nil {
		return domain.Supplier{}, err
	}

	return domain.Supplier{
		ID:                resp.Id,
		Name:              resp.Name,
		LegalAddress:      resp.LegalAddress,
		ActualAddress:     resp.ActualAddress,
		WarehouseAddress:  resp.WarehouseAddress,
		ContactPerson:     resp.ContactPerson,
		Phone:             resp.Phone,
		Email:             resp.Email,
		Website:           resp.Website,
		ContractNumber:    resp.ContractNumber,
		ProductCategories: resp.ProductCategories,
		PurchaseAmount:    resp.PurchaseAmount,
		Balance:           resp.Balance,
		ProductTypes:      resp.ProductTypes,
		Comments:          resp.Comments,
		Files:             resp.Files,
		Country:           resp.Country,
		Region:            resp.Region,
		TaxID:             resp.TaxId,
		BankDetails:       resp.BankDetails,
		RegistrationDate:  resp.RegistrationDate.AsTime(),
		PaymentTerms:      resp.PaymentTerms,
		IsActive:          resp.IsActive,
		OtherFields:       otherFields,
	}, nil
}

func (s *SuppliersClient) Create(ctx context.Context, spl domain.Supplier) (int64, error) {
	otherFieldsJSON, err := json.Marshal(spl.OtherFields)
	if err != nil {
		return 0, err
	}

	resp, err := s.supplierClient.Create(ctx, &supplier.Supplier{
		Id:                spl.ID,
		Name:              spl.Name,
		LegalAddress:      spl.LegalAddress,
		ActualAddress:     spl.ActualAddress,
		WarehouseAddress:  spl.WarehouseAddress,
		ContactPerson:     spl.ContactPerson,
		Phone:             spl.Phone,
		Email:             spl.Email,
		Website:           spl.Website,
		ContractNumber:    spl.ContractNumber,
		ProductCategories: spl.ProductCategories,
		PurchaseAmount:    spl.PurchaseAmount,
		Balance:           spl.Balance,
		ProductTypes:      spl.ProductTypes,
		Comments:          spl.Comments,
		Files:             spl.Files,
		Country:           spl.Country,
		Region:            spl.Region,
		TaxId:             spl.TaxID,
		BankDetails:       spl.BankDetails,
		RegistrationDate:  timestamppb.New(spl.RegistrationDate),
		PaymentTerms:      spl.PaymentTerms,
		IsActive:          spl.IsActive,
		OtherFields:       string(otherFieldsJSON),
	})
	if err != nil {
		return 0, err
	}

	return resp.Id, nil
}
