package models

import "github.com/gofrs/uuid"

type test struct {
	input  []Requirement
	output string
}

func (ms *ModelSuite) Test_Create() {
	tests := []test{
		{
			input: []Requirement{
				{
					Title:                  "",
					Description:            "Description",
					CreatedByUserID:        uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					ServiceDepartmentID:    uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementTypeID:      uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementSubTypeID:   uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequestingDepartmentID: uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
				},
			},
			output: "Title is required.",
		},
		{
			input: []Requirement{
				{
					Title:                  "Title",
					Description:            "",
					CreatedByUserID:        uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					ServiceDepartmentID:    uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementTypeID:      uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementSubTypeID:   uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequestingDepartmentID: uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
				},
			},
			output: "Description is required.",
		},
		{
			input: []Requirement{
				{
					Title:                  "TitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitle",
					Description:            "Description",
					CreatedByUserID:        uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					ServiceDepartmentID:    uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementTypeID:      uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementSubTypeID:   uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequestingDepartmentID: uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
				},
			},
			output: " Title must be less than 255 characters.",
		},
		{
			input: []Requirement{
				{
					Title:                  "Title",
					CreatedByUserID:        uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					ServiceDepartmentID:    uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementTypeID:      uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementSubTypeID:   uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequestingDepartmentID: uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					Description:            "TitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitle",
				},
			},
			output: " Description must be less than 255 characters.",
		},
		{
			input: []Requirement{
				{
					Title:                  "Title@",
					Description:            "Description",
					CreatedByUserID:        uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					ServiceDepartmentID:    uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementTypeID:      uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementSubTypeID:   uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequestingDepartmentID: uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
				},
			},
			output: " Title must be letters only.",
		},
		{
			input: []Requirement{
				{
					Title:                  "Title",
					Description:            "Description",
					CreatedByUserID:        uuid.UUID{},
					ServiceDepartmentID:    uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementTypeID:      uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementSubTypeID:   uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequestingDepartmentID: uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
				},
			},
			output: "User is required.",
		},
		{
			input: []Requirement{
				{
					Title:                  "Title",
					Description:            "Description",
					CreatedByUserID:        uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					ServiceDepartmentID:    uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementTypeID:      uuid.UUID{},
					RequirementSubTypeID:   uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequestingDepartmentID: uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
				},
			},
			output: "Type is required.",
		},
		{
			input: []Requirement{
				{
					Title:                  "Title",
					Description:            "Description",
					CreatedByUserID:        uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					ServiceDepartmentID:    uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementTypeID:      uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementSubTypeID:   uuid.UUID{},
					RequestingDepartmentID: uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
				},
			},
			output: "Subtype is required.",
		},
		{
			input: []Requirement{
				{
					Title:                  "Title",
					Description:            "Description",
					CreatedByUserID:        uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					ServiceDepartmentID:    uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementTypeID:      uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequirementSubTypeID:   uuid.Must(uuid.FromString("2fa909bc-6c1e-4537-82de-b6bd72019ed6")),
					RequestingDepartmentID: uuid.UUID{},
				},
			},
			output: "Area is required.",
		},
	}

	for _, test := range tests {
		for _, requirement := range test.input {
			err, _ := requirement.Validate(ms.DB)
			if err != nil {
				ms.Equal(test.output, err.Error())
			}
		}
	}
}
