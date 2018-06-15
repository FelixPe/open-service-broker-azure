package postgresql

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (
	d *databaseManager,
) getProvisionParametersSchema() service.InputParametersSchema {
	return service.InputParametersSchema{
		PropertySchemas: map[string]service.PropertySchema{
			"extensions": dbExtensionsSchema,
		},
	}
}

type databaseInstanceDetails struct {
	ARMDeploymentName string `json:"armDeployment"`
	DatabaseName      string `json:"database"`
}

func (d *databaseManager) GetEmptyInstanceDetails() service.InstanceDetails {
	return &databaseInstanceDetails{}
}

func (d *databaseManager) GetEmptyBindingDetails() service.BindingDetails {
	return &bindingDetails{}
}
