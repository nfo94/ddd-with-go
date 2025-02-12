Excerpts from the book:

> In domain-driven design, entities are defined by their identity. Their attributes do
> not define them, and it is expected that although their attributes may change over
> time, their identity will not. While the entity may change so much that it is
> indistinguishable from where it started, it retains the same identity, and we treat it
> as the same object

> Anemic models have little or no domain behavior as part of their design. This means
> that you are not getting the full benefit of DDD

> If your model has mostly public getter and setter functions, no business logic, or
> depends on various clients to implement the business logic, you probably have an
> anemic model.

> If you want to use an ORM, ensure it does not control how you write your entities in
> your DDD context; otherwise, you may end up with an anemic model. We also want to keep
> the coupling between our entity and ORM to a minimum. Therefore, I recommend you use
> an adaptor layer to decouple your ORM and DDD entity layer.

> Value objects are, in some ways, the opposite of entities. With value objects, we want
> to assert that two objects are the same given their values. Value objects do not have
> identities and are often used in conjunction with entities and aggregates to enable
> us to build a rich model of our domain. We typically use them to measure, quantify,
> or describe something about our domain.

> It is recommended that value objects remain immutable to prevent any unexpected
> behavior.

> Value objects should be replaceable.

> If you care only about the values of an object, then it should preferably be a value
> object. Some other questions to ask yourself to ensure a value object is the right
> choice for you are:
>
> -   Is it possible for me to treat this object as immutable?
> -   Does it measure, quantify, or describe a domain concept?
> -   Can it be compared to other objects of the same type by its values?
>
> If the answers to all these questions are yes, a value object is probably right for
> your use case.

> In domain-driven design, the aggregate pattern refers to a group of domain objects
> that can be treated as one for some behaviors
