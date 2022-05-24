package master

import (
	"fmt"
	"context"
	"biostar/service/tenant"
)

func (c *MasterClient) InitTenant(tenantID, gatewayID string) error {
	tenantClient := tenant.NewTenantClient(c.conn)

	getReq := &tenant.GetRequest{
		TenantIDs: []string{ tenantID },
	}

	getResp, err := tenantClient.Get(context.Background(), getReq)

	if err == nil && len(getResp.TenantInfos) == 1 { // tenant is already initialized
		fmt.Printf("%v is already registered.\n", tenantID)
		return nil
	}

	tenantInfo := &tenant.TenantInfo{
		TenantID: tenantID,
		GatewayIDs: []string{ gatewayID },
	}

	addReq := &tenant.AddRequest{
		TenantInfos: []*tenant.TenantInfo {
			tenantInfo,
		},
	}

	_, err = tenantClient.Add(context.Background(), addReq)

	if err != nil {
		fmt.Printf("Cannot add the tenant: %v", err)
		return err
	}

	fmt.Printf("%v is registered successfully.\n", tenantID)

	return nil
}
