# Git templates

Here are some templates for commit messages.

## Concise commit

```
<Type>: <Summary> <Ticket>
```

**Type**: [feat/fix/docs/chore/style/refactor/perf/test]

**Summary**: [Brief description of the change]

**Ticket**: [KAFKA/KAFKABETR-ticketNumber]

### Example

```
feat: Add user authentication for the login page KAFKA-1698
```

## Full commit message

```
<Type>: <Summary> <Ticket>

**Details:**: <Details>
```

**Type**: [feat/fix/docs/chore/style/refactor/perf/test]

**Summary**: [Brief description of the change]

**Ticket**: [KAFKA/KAFKABETR-ticketNumber]

**Details**: [Bullet point with details about the key changes of what has changed
and why, if necessary]

### Example

```
feat: Add user authentication for the login page KAFKA-1698

Details:
- Added login page form
- Added binding to backhand user authentication

```
