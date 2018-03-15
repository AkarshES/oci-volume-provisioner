// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// LaunchInstanceRequest wrapper for the LaunchInstance operation
type LaunchInstanceRequest struct {
        
 // Instance details 
         LaunchInstanceDetails `contributesTo:"body"`
        
 // A token that uniquely identifies a request so it can be retried in case of a timeout or
 // server error without risk of executing that same action again. Retry tokens expire after 24
 // hours, but can be invalidated before then due to conflicting operations (for example, if a resource
 // has been deleted and purged from the system, then a retry of the original creation request
 // may be rejected). 
        OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`
        
 // Name of the pool in which to launch an instance. This feature is currently in preview and may change before public release. Do not use it for production workloads.
        OpcPoolName *string `mandatory:"false" contributesTo:"query" name:"opc-pool-name"`
}

func (request LaunchInstanceRequest) String() string {
    return common.PointerString(request)
}

// LaunchInstanceResponse wrapper for the LaunchInstance operation
type LaunchInstanceResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The Instance instance
     Instance `presentIn:"body"`

    
 // For optimistic concurrency control. See `if-match`.
    Etag *string `presentIn:"header" name:"etag"`
    
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
 // a particular request, please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`


}

func (response LaunchInstanceResponse) String() string {
    return common.PointerString(response)
}
