/**
This package offers some casting and flattening logic which can be useful when you start dealing with compound Topic results. 

Topic aggregation functions like And and Or provide you with events that have the form map[string][]interface{}. Quite often, 
the interfaces used in the values are themselves arrays or maps as well. This makes transformations tedious and problematic. 

For example, within a call to NewSubscriber(func(event interface{}) you could use this snippet of code:

result, err := New(event).First to obtain the first element in the array, or an error if the event is not []interface{}

Many of the methods here have their XXXAndWrap variant, which does not return the interface{}, but 'wraps' the interface{} with a call 
to New. This allows for a Fluent API approach can be taken.
*/
package results

import (
    "errors"
    "fmt"
)

/* A wrapper around a Topic result which allows for casting and flattening functionalities 
*/
type Results struct {
    result interface{}
}

func New(result interface{}) *Results {
    return &Results { result }
}

func (r *Results) First() (interface{}, error) {
    switch r.result.(type) {
    case []interface{}:
        return r.result.([]interface{})[0], nil
    default:
    }
    return nil, r.toError("This is not an array")
}

func (r *Results) toError(message string) error {
    return errors.New(fmt.Sprintf(message+": %v %T", r.result, r.result))
}
