// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// ListCrossConnectsRequest wrapper for the ListCrossConnects operation
type ListCrossConnectsRequest struct {
        
 // The OCID of the compartment. 
        CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`
        
 // The OCID of the cross-connect group. 
        CrossConnectGroupId *string `mandatory:"false" contributesTo:"query" name:"crossConnectGroupId"`
        
 // The maximum number of items to return in a paginated "List" call.
 // Example: `500` 
        Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`
        
 // The value of the `opc-next-page` response header from the previous "List" call. 
        Page *string `mandatory:"false" contributesTo:"query" name:"page"`
        
 // A filter to return only resources that match the given display name exactly. 
        DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`
        
 // The field to sort by. You can provide one sort order (`sortOrder`). Default order for
 // TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
 // sort order is case sensitive.
 // **Note:** In general, some "List" operations (for example, `ListInstances`) let you
 // optionally filter by Availability Domain if the scope of the resource type is within a
 // single Availability Domain. If you call one of these "List" operations without specifying
 // an Availability Domain, the resources are grouped by Availability Domain, then sorted. 
        SortBy ListCrossConnectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`
        
 // The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
 // is case sensitive. 
        SortOrder ListCrossConnectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`
        
 // A filter to return only resources that match the specified lifecycle state. The value is case insensitive. 
        LifecycleState CrossConnectLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`
}

func (request ListCrossConnectsRequest) String() string {
    return common.PointerString(request)
}

// ListCrossConnectsResponse wrapper for the ListCrossConnects operation
type ListCrossConnectsResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The []CrossConnect instance
    Items []CrossConnect `presentIn:"body"`

    
 // For pagination of a list of items. When paging through a list, if this header appears in the response,
 // then a partial list might have been returned. Include this value as the `page` parameter for the
 // subsequent GET request to get the next batch of items.
    OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
    
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
 // a particular request, please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`


}

func (response ListCrossConnectsResponse) String() string {
    return common.PointerString(response)
}

// ListCrossConnectsSortByEnum Enum with underlying type: string
type ListCrossConnectsSortByEnum string

// Set of constants representing the allowable values for ListCrossConnectsSortBy
const (
    ListCrossConnectsSortByTimecreated ListCrossConnectsSortByEnum = "TIMECREATED"
    ListCrossConnectsSortByDisplayname ListCrossConnectsSortByEnum = "DISPLAYNAME"
    ListCrossConnectsSortByUnknown ListCrossConnectsSortByEnum = "UNKNOWN"
)

var mappingListCrossConnectsSortBy = map[string]ListCrossConnectsSortByEnum { 
    "TIMECREATED": ListCrossConnectsSortByTimecreated,
    "DISPLAYNAME": ListCrossConnectsSortByDisplayname,
    "UNKNOWN": ListCrossConnectsSortByUnknown,
}

// GetListCrossConnectsSortByEnumValues Enumerates the set of values for ListCrossConnectsSortBy
func GetListCrossConnectsSortByEnumValues() []ListCrossConnectsSortByEnum {
   values := make([]ListCrossConnectsSortByEnum, 0)
   for _, v := range mappingListCrossConnectsSortBy {
      if v != ListCrossConnectsSortByUnknown {
         values = append(values, v)
      }
   }
   return values
}

// ListCrossConnectsSortOrderEnum Enum with underlying type: string
type ListCrossConnectsSortOrderEnum string

// Set of constants representing the allowable values for ListCrossConnectsSortOrder
const (
    ListCrossConnectsSortOrderAsc ListCrossConnectsSortOrderEnum = "ASC"
    ListCrossConnectsSortOrderDesc ListCrossConnectsSortOrderEnum = "DESC"
    ListCrossConnectsSortOrderUnknown ListCrossConnectsSortOrderEnum = "UNKNOWN"
)

var mappingListCrossConnectsSortOrder = map[string]ListCrossConnectsSortOrderEnum { 
    "ASC": ListCrossConnectsSortOrderAsc,
    "DESC": ListCrossConnectsSortOrderDesc,
    "UNKNOWN": ListCrossConnectsSortOrderUnknown,
}

// GetListCrossConnectsSortOrderEnumValues Enumerates the set of values for ListCrossConnectsSortOrder
func GetListCrossConnectsSortOrderEnumValues() []ListCrossConnectsSortOrderEnum {
   values := make([]ListCrossConnectsSortOrderEnum, 0)
   for _, v := range mappingListCrossConnectsSortOrder {
      if v != ListCrossConnectsSortOrderUnknown {
         values = append(values, v)
      }
   }
   return values
}
