// Code generated by goa v2.0.5, DO NOT EDIT.
//
// resume HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/multipart/design -o
// $(GOPATH)/src/goa.design/examples/multipart

package client

import (
	"encoding/json"
	"fmt"

	resume "goa.design/examples/multipart/gen/resume"
)

// BuildAddPayload builds the payload for the resume add endpoint from CLI
// flags.
func BuildAddPayload(resumeAddBody string) ([]*resume.Resume, error) {
	var err error
	var body []*ResumeRequestBody
	{
		err = json.Unmarshal([]byte(resumeAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'[\n      {\n         \"education\": [\n            {\n               \"institution\": \"Qui ut voluptates libero consectetur in.\",\n               \"major\": \"Aliquid nesciunt harum exercitationem qui reprehenderit.\"\n            },\n            {\n               \"institution\": \"Qui ut voluptates libero consectetur in.\",\n               \"major\": \"Aliquid nesciunt harum exercitationem qui reprehenderit.\"\n            }\n         ],\n         \"experience\": [\n            {\n               \"company\": \"Consectetur ut.\",\n               \"duration\": 9021502551841493030,\n               \"role\": \"Pariatur ad accusantium.\"\n            },\n            {\n               \"company\": \"Consectetur ut.\",\n               \"duration\": 9021502551841493030,\n               \"role\": \"Pariatur ad accusantium.\"\n            }\n         ],\n         \"name\": \"Earum rem.\"\n      },\n      {\n         \"education\": [\n            {\n               \"institution\": \"Qui ut voluptates libero consectetur in.\",\n               \"major\": \"Aliquid nesciunt harum exercitationem qui reprehenderit.\"\n            },\n            {\n               \"institution\": \"Qui ut voluptates libero consectetur in.\",\n               \"major\": \"Aliquid nesciunt harum exercitationem qui reprehenderit.\"\n            }\n         ],\n         \"experience\": [\n            {\n               \"company\": \"Consectetur ut.\",\n               \"duration\": 9021502551841493030,\n               \"role\": \"Pariatur ad accusantium.\"\n            },\n            {\n               \"company\": \"Consectetur ut.\",\n               \"duration\": 9021502551841493030,\n               \"role\": \"Pariatur ad accusantium.\"\n            }\n         ],\n         \"name\": \"Earum rem.\"\n      },\n      {\n         \"education\": [\n            {\n               \"institution\": \"Qui ut voluptates libero consectetur in.\",\n               \"major\": \"Aliquid nesciunt harum exercitationem qui reprehenderit.\"\n            },\n            {\n               \"institution\": \"Qui ut voluptates libero consectetur in.\",\n               \"major\": \"Aliquid nesciunt harum exercitationem qui reprehenderit.\"\n            }\n         ],\n         \"experience\": [\n            {\n               \"company\": \"Consectetur ut.\",\n               \"duration\": 9021502551841493030,\n               \"role\": \"Pariatur ad accusantium.\"\n            },\n            {\n               \"company\": \"Consectetur ut.\",\n               \"duration\": 9021502551841493030,\n               \"role\": \"Pariatur ad accusantium.\"\n            }\n         ],\n         \"name\": \"Earum rem.\"\n      },\n      {\n         \"education\": [\n            {\n               \"institution\": \"Qui ut voluptates libero consectetur in.\",\n               \"major\": \"Aliquid nesciunt harum exercitationem qui reprehenderit.\"\n            },\n            {\n               \"institution\": \"Qui ut voluptates libero consectetur in.\",\n               \"major\": \"Aliquid nesciunt harum exercitationem qui reprehenderit.\"\n            }\n         ],\n         \"experience\": [\n            {\n               \"company\": \"Consectetur ut.\",\n               \"duration\": 9021502551841493030,\n               \"role\": \"Pariatur ad accusantium.\"\n            },\n            {\n               \"company\": \"Consectetur ut.\",\n               \"duration\": 9021502551841493030,\n               \"role\": \"Pariatur ad accusantium.\"\n            }\n         ],\n         \"name\": \"Earum rem.\"\n      }\n   ]'")
		}
	}
	v := make([]*resume.Resume, len(body))
	for i, val := range body {
		v[i] = &resume.Resume{
			Name: val.Name,
		}
		if val.Experience != nil {
			v[i].Experience = make([]*resume.Experience, len(val.Experience))
			for j, val := range val.Experience {
				v[i].Experience[j] = marshalExperienceRequestBodyToResumeExperience(val)
			}
		}
		if val.Education != nil {
			v[i].Education = make([]*resume.Education, len(val.Education))
			for j, val := range val.Education {
				v[i].Education[j] = marshalEducationRequestBodyToResumeEducation(val)
			}
		}
	}
	return v, nil
}
