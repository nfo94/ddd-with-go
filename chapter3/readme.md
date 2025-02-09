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
