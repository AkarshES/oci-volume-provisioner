// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // CreateDrgDetails The representation of CreateDrgDetails
type CreateDrgDetails struct {
    
 // The OCID of the compartment to contain the DRG.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
    DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateDrgDetails) String() string {
    return common.PointerString(m)
}




