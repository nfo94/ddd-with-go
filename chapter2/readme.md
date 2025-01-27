Excerpts:

> domains are “a sphere of knowledge, influence, or activity.”

> Every time you read the phrase domain-driven design, you could read it as business problem-driven design

> Domains and sub-domains can be used almost interchangeably. We tend to use a sub-domain to signal that the domain we are talking about is a child of a higher-level domain

> Ubiquitous language is the overlap of the language that domain experts and technical experts use.

> This language should be used when discussing requirements and system design and should even be used in the source code itself.

> What if someone from a different area of a business came to discuss customers with us? The first thing we should do is define what a customer means to them as it may mean something different within their bounded context.

> Bounded contexts are all about dividing large models into smaller, easier-to-understand chunks and being explicit about how they relate to each other

> Since several bounded contexts often must communicate, we often apply patterns to
> ensure our models can maintain integrity. The three main patterns are as follows:

-   Open Host Service
-   Published language
-   Anti-corruption layer

> An Open Host Service is a means of giving other systems (or sub-systems) access to ours.

> A ubiquitous language is our team’s internal formally defined language; a published language is the opposite. If our team is going to expose some of our systems to other teams via an Open Host Service, we need to ensure the definition of what we expose to other teams in different bounded contexts is clear

> Two popular ways to present published language are via OpenAPI or gRPC.

> The downside to OpenAPI is that there are more performant alternatives out there.
> Furthermore, OpenAPI does not give any protection for breaking changes natively. For
> example, if you removed a field from your documentation, but another team depended on
> it, you would likely break their workflow

> gRPC was created at Google to handle remote communication at scale. It supports load
> balancing, tracing, health checks, bi-directional streaming, and authentication.

> Furthermore, gRPC uses binary serialization to compress the payload it sends, making
> it very efficient and fast.

> To call a method on a remote server, firstly, we must define our message protobuf.
> Protobufs are typically defined in a .proto file, and they are language-agnostic

> Sometimes called an adapter layer, an anti-corruption layer can be used to translate
> models from different systems
