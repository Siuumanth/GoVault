
In GoVault (and in your Auth service), the rule is:

```
HTTP Handler  ⇄  DTO  ⇄  Service  ⇄  Model  ⇄  Repository
```

DTOs exist **only** at the **edge**:

* to parse JSON
* to shape responses
* to version APIs

Once the data enters the **service layer**, it should be:

* domain models
* primitives
* rich objects


Done during Upload testing:

- Debugged env collisions on Windows
- Fixed DSN parsing
- Designed clean DTOs
- Implemented chunked uploads correctly
- Handled filesystem paths + mkdir
- Thought about indexing, latency, and schema semantics