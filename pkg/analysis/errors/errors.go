package errors

import "errors"

var (
	ErrCouldNotCreateMarkers         = errors.New("could not create markers")
	ErrCouldNotCreateStructFieldTags = errors.New("could not create new structFieldTags")
	ErrCouldNotGetInspector          = errors.New("could not get inspector")
	ErrCouldNotGetJSONTags           = errors.New("could not get json tags")
	ErrCouldNotGetMarkers            = errors.New("could not get markers")
)
