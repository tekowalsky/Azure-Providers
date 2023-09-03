// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = SqlPoolExtendedAuditingPolicyId{}

func TestSqlPoolExtendedAuditingPolicyIDFormatter(t *testing.T) {
	actual := NewSqlPoolExtendedAuditingPolicyID("12345678-1234-9876-4563-123456789012", "resGroup1", "workspace1", "sqlPool1", "default").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/extendedAuditingSettings/default"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestSqlPoolExtendedAuditingPolicyID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SqlPoolExtendedAuditingPolicyId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/",
			Error: true,
		},

		{
			// missing WorkspaceName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/",
			Error: true,
		},

		{
			// missing value for WorkspaceName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/",
			Error: true,
		},

		{
			// missing SqlPoolName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/",
			Error: true,
		},

		{
			// missing value for SqlPoolName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/",
			Error: true,
		},

		{
			// missing ExtendedAuditingSettingName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/",
			Error: true,
		},

		{
			// missing value for ExtendedAuditingSettingName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/extendedAuditingSettings/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/extendedAuditingSettings/default",
			Expected: &SqlPoolExtendedAuditingPolicyId{
				SubscriptionId:              "12345678-1234-9876-4563-123456789012",
				ResourceGroup:               "resGroup1",
				WorkspaceName:               "workspace1",
				SqlPoolName:                 "sqlPool1",
				ExtendedAuditingSettingName: "default",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/12345678-1234-9876-4563-123456789012/RESOURCEGROUPS/RESGROUP1/PROVIDERS/MICROSOFT.SYNAPSE/WORKSPACES/WORKSPACE1/SQLPOOLS/SQLPOOL1/EXTENDEDAUDITINGSETTINGS/DEFAULT",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := SqlPoolExtendedAuditingPolicyID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.WorkspaceName != v.Expected.WorkspaceName {
			t.Fatalf("Expected %q but got %q for WorkspaceName", v.Expected.WorkspaceName, actual.WorkspaceName)
		}
		if actual.SqlPoolName != v.Expected.SqlPoolName {
			t.Fatalf("Expected %q but got %q for SqlPoolName", v.Expected.SqlPoolName, actual.SqlPoolName)
		}
		if actual.ExtendedAuditingSettingName != v.Expected.ExtendedAuditingSettingName {
			t.Fatalf("Expected %q but got %q for ExtendedAuditingSettingName", v.Expected.ExtendedAuditingSettingName, actual.ExtendedAuditingSettingName)
		}
	}
}
