# cloud-project-request
Library to Support ServiceNow Cloud Project Request Table mapping to/from JSON for api/documentation/orchestration and API integration.

## Documentation

### Data model

|Project|||||
|---|---|---|---|---|
|ProjectName|string [5:40|required|must be a unique project name|`json:"projectName,omitempty" validate:"max=40,min=5,projunique" binding:"required"`|
|ProjectId|string [20]|optional|(counter provided by system)|`json:"projectId,omitempty" validate:"max"`|
|ProjectGroup|*OrgGroup|required project group expected name following naming conventions|`json:"projectGroup,omitempty" validate:"" binding:"required"`|
|ProjectOwner|*Person|required|project owner should be a registered person who will become the ADO project owner `json:"projectOwner" validate:"person" binding:"required"`|
|ProjectEngineer|*Person|required|project engineering lead should be a registered person who will own the cloud automation/environment and should be certified for role |`json:"projectEngineer" validate:"person" binding:"required"`|
|SalesId|SalesId|optional salesID for tracked clinet opportunity / engagement|`json:"salesId,omitempty" validate:"" binding:""`|
|Contract|Contract|optional contract finance reference / billing code for project|`json:"contract,omitempty" validate:"" binding:""`|
|Start|PTime|optional expected start time for use of environment|`json:"expectedStart,omitempty" validate:"" binding:""`|
|End|PTime|optional|expected end time for use of environment|`json:"expectectedEnd,omitempty" validate:"" binding:""`|
|Regulations|[]*string|optional|regulations likely in force within project implement as dictionary for validation|`json:"regulations" validate:"" binding:""`|
|Partners|[]*string|optional partners involved in development/sale|`json:"developmentPartners" validate:"" binding:""`|

|OrgGroup|key|na|abstract table of groups tied to projects|
|---|---|---|---|
|Name|string|required name of group|

|Person|key|na|person|
|---|---|---|---|
|Name|string [5:40]|optional|e.g. fname lname|
|Email|string [email]|required|valid email|

|Miscellaneous Type Definitions|na|abstract|  |
|---|---|---|---|
|SalesId|string [40]|optional|validate with CRM/SFA|
|Contract|string[40]|optional|may be a string the talks about the specific contract, eventually should hold a client pre-sales/project code|
|PTime|string[RFC3339]|optional|from/to date construct project, may want to validate upstream based upon some SLA (e.g. 24hr/1day response?) where used|



## License
Copyright 2012, Data Cloud, LLC and others. All rights reserved
Licensed under the MIT License (MIT)


