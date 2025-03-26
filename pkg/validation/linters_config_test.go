package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/JoelSpeed/kal/pkg/config"
	"github.com/JoelSpeed/kal/pkg/validation"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

var _ = Describe("LintersConfig", func() {
	type validateLintersConfigTableInput struct {
		config      config.LintersConfig
		expectedErr string
	}

	DescribeTable("Validate Linters Configuration", func(in validateLintersConfigTableInput) {
		errs := validation.ValidateLintersConfig(in.config, field.NewPath("lintersConfig"))
		if len(in.expectedErr) > 0 {
			Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
		} else {
			Expect(errs).To(HaveLen(0), "No errors were expected")
		}

	},
		Entry("Empty config", validateLintersConfigTableInput{
			config:      config.LintersConfig{},
			expectedErr: "",
		}),

		// ConditionsConfig validation
		Entry("With a valid ConditionsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					IsFirstField: "",
					UseProtobuf:  "",
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig IsFirstField: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					IsFirstField: config.ConditionsFirstFieldWarn,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig IsFirstField: Ignore", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					IsFirstField: config.ConditionsFirstFieldIgnore,
				},
			},
			expectedErr: "",
		}),
		Entry("With an invalid ConditionsConfig IsFirstField", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					IsFirstField: "invalid",
				},
			},
			expectedErr: "lintersConfig.conditions.isFirstField: Invalid value: \"invalid\": invalid value, must be one of \"Warn\", \"Ignore\" or omitted",
		}),
		Entry("With a valid ConditionsConfig UseProtobuf: SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UseProtobuf: config.ConditionsUseProtobufSuggestFix,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UseProtobuf: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UseProtobuf: config.ConditionsUseProtobufWarn,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UseProtobuf: Ignore", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UseProtobuf: config.ConditionsUseProtobufIgnore,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UseProtobuf: Forbid", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UseProtobuf: config.ConditionsUseProtobufForbid,
				},
			},
			expectedErr: "",
		}),
		Entry("With an invalid ConditionsConfig UseProtobuf", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UseProtobuf: "invalid",
				},
			},
			expectedErr: "lintersConfig.conditions.useProtobuf: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\", \"Ignore\", \"Forbid\" or omitted",
		}),
		Entry("With a valid ConditionsConfig UsePatchStrategy: SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UsePatchStrategy: config.ConditionsUsePatchStrategySuggestFix,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UsePatchStrategy: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UsePatchStrategy: config.ConditionsUsePatchStrategyWarn,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UsePatchStrategy: Ignore", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UsePatchStrategy: config.ConditionsUsePatchStrategyIgnore,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UsePatchStrategy: Forbid", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UsePatchStrategy: config.ConditionsUsePatchStrategyForbid,
				},
			},
			expectedErr: "",
		}),
		Entry("With an invalid ConditionsConfig UsePatchStrategy", validateLintersConfigTableInput{
			config: config.LintersConfig{
				Conditions: config.ConditionsConfig{
					UsePatchStrategy: "invalid",
				},
			},
			expectedErr: "lintersConfig.conditions.usePatchStrategy: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\", \"Ignore\", \"Forbid\" or omitted",
		}),

		// JSONTagsConfig validation
		Entry("With a valid JSONTagsConfig JSONTagRegex", validateLintersConfigTableInput{
			config: config.LintersConfig{
				JSONTags: config.JSONTagsConfig{
					JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$",
				},
			},
			expectedErr: "",
		}),
		Entry("With an invalid JSONTagsConfig JSONTagRegex", validateLintersConfigTableInput{
			config: config.LintersConfig{
				JSONTags: config.JSONTagsConfig{
					JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*",
				},
			},
			expectedErr: "lintersConfig.jsonTags.jsonTagRegex: Invalid value: \"^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*\": invalid regex: error parsing regexp: missing closing ): `^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*`",
		}),

		// NoMapsConfig validation
		Entry("With a valid NoMapsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				NoMaps: config.NoMapsConfig{
					Policy: "",
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid NoMapsConfig: enforce is specified", validateLintersConfigTableInput{
			config: config.LintersConfig{
				NoMaps: config.NoMapsConfig{
					Policy: config.NoMapsEnforce,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid NoMapsConfig: allowStringToStringMaps is specified", validateLintersConfigTableInput{
			config: config.LintersConfig{
				NoMaps: config.NoMapsConfig{
					Policy: config.NoMapsAllowStringToStringMaps,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid NoMapsConfig: ignore is specified", validateLintersConfigTableInput{
			config: config.LintersConfig{
				NoMaps: config.NoMapsConfig{
					Policy: config.NoMapsIgnore,
				},
			},
			expectedErr: "",
		}),
		Entry("With a invalid NoMapsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				NoMaps: config.NoMapsConfig{
					Policy: "invalid",
				},
			},
			expectedErr: `lintersConfig.nomaps.policy: Invalid value: "invalid": invalid value, must be one of "Enforce", "AllowStringToStringMaps", "Ignore" or omitted`,
		}),

		// OptionalOrRequiredConfig validation
		Entry("With a valid OptionalOrRequiredConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				OptionalOrRequired: config.OptionalOrRequiredConfig{
					PreferredOptionalMarker: "optional",
					PreferredRequiredMarker: "required",
				},
			},
			expectedErr: "",
		}),
		Entry("With kubebuilder preferred markers", validateLintersConfigTableInput{
			config: config.LintersConfig{
				OptionalOrRequired: config.OptionalOrRequiredConfig{
					PreferredOptionalMarker: "kubebuilder:validation:Optional",
					PreferredRequiredMarker: "kubebuilder:validation:Required",
				},
			},
			expectedErr: "",
		}),
		Entry("With invalid preferred optional marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				OptionalOrRequired: config.OptionalOrRequiredConfig{
					PreferredOptionalMarker: "invalid",
				},
			},
			expectedErr: "lintersConfig.optionalOrRequired.preferredOptionalMarker: Invalid value: \"invalid\": invalid value, must be one of \"optional\", \"kubebuilder:validation:Optional\" or omitted",
		}),
		Entry("With invalid preferred required marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				OptionalOrRequired: config.OptionalOrRequiredConfig{
					PreferredRequiredMarker: "invalid",
				},
			},
			expectedErr: "lintersConfig.optionalOrRequired.preferredRequiredMarker: Invalid value: \"invalid\": invalid value, must be one of \"required\", \"kubebuilder:validation:Required\" or omitted",
		}),

		// RequiredFieldsConfig validation
		Entry("With a valid RequiredFieldsConfig: omitted", validateLintersConfigTableInput{
			config: config.LintersConfig{
				RequiredFields: config.RequiredFieldsConfig{
					PointerPolicy: "",
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid RequiredFieldsConfig: SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				RequiredFields: config.RequiredFieldsConfig{
					PointerPolicy: config.RequiredFieldPointerSuggestFix,
				},
			},
			expectedErr: "",
		}),
		Entry("With a valid RequiredFieldsConfig: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				RequiredFields: config.RequiredFieldsConfig{
					PointerPolicy: config.RequiredFieldPointerWarn,
				},
			},
			expectedErr: "",
		}),
		Entry("With an invalid RequiredFieldsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				RequiredFields: config.RequiredFieldsConfig{
					PointerPolicy: "invalid",
				},
			},
			expectedErr: "lintersConfig.requiredFields.pointerPolicy: Invalid value: \"invalid\": invalid value, must be one of \"Warn\", \"SuggestFix\" or omitted",
		}),
	)
})
