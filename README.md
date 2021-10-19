# cloud-project-request
Library to Support ServiceNow Cloud Project Request Table mapping to/from JSON for api/documentation/orchestration and API integration.

## Documentation

### Data model
|Request|key|na|Complex Cloud Service Request|
|---|---|---|---|
|ProjectRequest|*Project|required|information about the project itself `json:"projectRequest"`|
|TenancyRequest| *Tenancy|required|information about the cloud tenancy(ies) required`json:"tenancyRequest"`|
|Integrations|*Integrations|required|requisite integrations like identity, github or other key services including front end firewall services`json:"integrations"|
|DevelopmentEnvironment|*Development|required|information around the developmnet environment, sdlc, toolchain and style  `json:"developmentEnvironment"`|

|Project|||||
|---|---|---|---|---|
|ProjectName|string [5:40|required|must be a unique project name|`json:"projectName,omitempty" validate:"max=40,min=5,projunique" binding:"required"`|
|ProjectId|string [20]|optional|(counter provided by system)|`json:"projectId,omitempty" validate:"max"`|
|ProjectGroup|*Group|required project group expected name following naming conventions|`json:"projectGroup,omitempty" validate:"" binding:"required"`|
|ProjectOwner|*Person|required|project owner should be a registered person who will become the ADO project owner `json:"projectOwner" validate:"person" binding:"required"`|
|ProjectEngineer|*Person|required|project engineering lead should be a registered person who will own the cloud automation/environment and should be certified for role |`json:"projectEngineer" validate:"person" binding:"required"`|
|SalesId|SalesId|optional salesID for tracked clinet opportunity / engagement|`json:"salesId,omitempty" validate:"" binding:""`|
|Contract|Contract|optional contract finance reference / billing code for project|`json:"contract,omitempty" validate:"" binding:""`|
|Start|PTime|optional expected start time for use of environment|`json:"expectedStart,omitempty" validate:"" binding:""`|
|End|PTime|optional|expected end time for use of environment|`json:"expectectedEnd,omitempty" validate:"" binding:""`|
|Regulations|[]*string|optional|regulations likely in force within project implement as dictionary for validation|`json:"regulations" validate:"" binding:""`|
|FedRamp|bool|required|is FedRAMP required for this application/service `json:"fedRamp" validate:"required" binding:"required"`|
|Partners|[]*string|optional partners involved in development/sale|`json:"developmentPartners" validate:"" binding:""`|

|Group|key|na|abstract table of groups tied to projects|
|---|---|---|---|
|Name|string|required name of group|
|Membership|[]*Person|optional|`json:"membership" validate:""`|

|Person|key|na|person|
|---|---|---|---|
|Name|string [5:40]|optional|e.g. fname lname|
|Email|string [email]|required|valid email|

|Miscellaneous Type Definitions|na|abstract|  |
|---|---|---|---|
|SalesId|string [40]|optional|validate with CRM/SFA|
|Contract|string[40]|optional|may be a string the talks about the specific contract, eventually should hold a client pre-sales/project code|
|PTime|string[RFC3339]|optional|from/to date construct project, may want to validate upstream based upon some SLA (e.g. 24hr/1day response?) where used|

|Region|key|na|cloud provider region information|
|---|---|---|---|
|DisplayName|string|required|Name e.g. US East 2 `json:"name"`|
|FullName|string|optional|Name e.g. (US) East 2 `json:"fullName"`|
|Latitude|string|optional|`json:"latitude,omitempty"`|
|Longitude|string|optional|`json:"longitude,omitempty"`|
|Name|string|required| for example useast2 `json:"code"`|
|Public|bool|required|public access vs. government `json:"public"`|
|Zones|[]string|optional|zones within region that can be selected `json:"zones"`|

|Provider|key|na|cloud provider|
|---|---|---|---|
|CloudProvider|string|required|one of aws, azure, awsgov, azuregov see @ProviderDict `json:"provider,omitempty" validate:"" binding:"required"`|
|PrimaryRegion|*Region|required|`json:"primaryRegion" validate:"" binding:""`|
|SecondaryRegion|*Region|optional|`json:"secondaryRegion,omitempty" validate:"" binding:""`|

var ProviderDict = []string{"unknown", "azure", "aws", "azuregov", "awsgov"} - valid

|Tenancy|key|na|tenant details at cloud provider|
|---|---|---|---|
|Provider|*Provider|required|`json:"provider" validate:"required"`|
|Subscription|string|optional| subscription identity - filled out at provisioning and might also be a resource group `json:"subscription" validate:"" binding:""`|
|Blueprint|string	|optional|name of catalog item used to setup environment, may be a github url containing terraform envs.`json:"blueprint" validate:"" binding:""`|
|SDLCStages|[]string|optional|sdlc stages required e.g. dev, test, stage, prod`json:"sdlcStages" validate:"" binding:""`|
|ExternalNetwork|string|required|statemetn on external networking... none,staticIP, corp intranet... may include CIDR block information `json:"network" validate:"" binding:""`|
|SharedServices|[]string|optional|shared services that must be bound e.g. logging via splunk, seim loggers, and other shared assets `json:"sharedServices" validate:"" binding:""`|

## License
Copyright 2021, Data Cloud, LLC and others. All rights reserved
Licensed under the MIT License (MIT)


