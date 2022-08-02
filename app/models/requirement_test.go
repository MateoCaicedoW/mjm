package models

type test struct {
	input  []Requirement
	output string
}

func (ms *ModelSuite) Test_Create() {
	tests := []test{
		{
			input: []Requirement{
				{
					Title:       "",
					Description: "Description",
				},
			},
			output: "Title is required.",
		},
		{
			input: []Requirement{
				{
					Title:       "Title",
					Description: "",
				},
			},
			output: "Description is required.",
		},
		{
			input: []Requirement{
				{
					Title:       "TitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitle",
					Description: "Description",
				},
			},
			output: " Title must be less than 255 characters.",
		},
		{
			input: []Requirement{
				{
					Title:       "Title",
					Description: "TitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitleTitle",
				},
			},
			output: " Description must be less than 255 characters.",
		},
		{
			input: []Requirement{
				{
					Title:       "Title@",
					Description: "Description",
				},
			},
			output: " Title must be letters only.",
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
