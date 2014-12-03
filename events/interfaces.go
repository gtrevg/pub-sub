/*
See http://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern for a description of the Publish-Subscribe pattern.
*/
package events

/*
The Publisher in the Publish-Subscribe pattern, or a shorthand for a function which you may call each time you want to inform about a particular event. Publishers can be created by any Topic instance. 
*/
type Publisher func(interface{})
/*
The Subscriber in the Publish-Subscribe pattern, or a shorthand for a function which you may provide, that is invoked whenever an event is published to the Topic. 
*/
type Subscriber func(interface{})
/*
A typical Topic used in a Pub-Sub pattern. The Topic has a name, which in theory should identify it uniquely among other topics. 
The implementation does not use this name, unless for informative reasons. Topics can create Publishers and Subscribers.  
Topics can be closed and this closes the Topic permanently.
*/
type Topic interface {
    //Allows you to create a new Publisher for a Topic. If you provide a callback (optional) here, it is guaranteed to be invoked during the publishing act. Depending on the 
    //implementation it may be also used in a defer() section or called when the the publish failed. This may happen if Publishing is using on a Closed Topic. 
    //Publishing may occur in its own go-routine, and it's not guaranteed that the order you call Publishers is preserved (especially if you write to multiple Topics). 
    NewPublisher() Publisher
    //Allows you to register an arbitrary Subscriber for events in the Topic
    //Subscribing may occur in its own go-routine, hence even if the act of subscribing 'blocks' (for example due to the waiting on channel), the remainders of the Topic still execute normally. 
    NewSubscriber(subscriber Subscriber)
    //Returns the topic's name
    String() string
    //Close frees the underlying resources, and depending on the implementation may render the Topic unusable
    Close() error
}