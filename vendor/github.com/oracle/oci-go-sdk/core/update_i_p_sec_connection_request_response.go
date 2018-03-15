// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// UpdateIPSecConnectionRequest wrapper for the UpdateIPSecConnection operation
type UpdateIPSecConnectionRequest struct {
        
 // The OCID of the IPSec connection. 
        IpscId *string `mandatory:"true" contributesTo:"path" name:"ipscId"`
        
 // Details object for updating a IPSec connection. 
         UpdateIpSecConnectionDetails `contributesTo:"body"`
        
 // For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
 // parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
 // will be updated or deleted only if the etag you provide matches the resource's current etag value. 
        IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request UpdateIPSecConnectionRequest) String() string {
    return common.PointerString(request)
}

// UpdateIPSecConnectionResponse wrapper for the UpdateIPSecConnection operation
type UpdateIPSecConnectionResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The IpSecConnection instance
     IpSecConnection `presentIn:"body"`

    
 // For optimistic concurrency control. See `if-match`.
    Etag *string `presentIn:"header" name:"etag"`
    
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
 // a particular request, please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`


}

func (response UpdateIPSecConnectionResponse) String() string {
    return common.PointerString(response)
}
