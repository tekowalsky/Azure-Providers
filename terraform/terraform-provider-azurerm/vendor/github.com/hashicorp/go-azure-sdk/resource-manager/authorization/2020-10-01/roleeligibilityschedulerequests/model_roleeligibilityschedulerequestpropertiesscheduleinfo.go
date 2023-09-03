package roleeligibilityschedulerequests

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleEligibilityScheduleRequestPropertiesScheduleInfo struct {
	Expiration    *RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration `json:"expiration,omitempty"`
	StartDateTime *string                                                         `json:"startDateTime,omitempty"`
}

func (o *RoleEligibilityScheduleRequestPropertiesScheduleInfo) GetStartDateTimeAsTime() (*time.Time, error) {
	if o.StartDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RoleEligibilityScheduleRequestPropertiesScheduleInfo) SetStartDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDateTime = &formatted
}
