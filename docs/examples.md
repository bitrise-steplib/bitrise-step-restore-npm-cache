⚠️ **Beta status**: While this Step is in beta, everyone can use it without restrictions, quotas or costs.

### Examples

Check out [Workflow Recipes](https://github.com/bitrise-io/workflow-recipes#-key-based-caching-beta) for other platform-specific examples!

#### Minimal example
```yaml
steps:
- restore-npm-cache@1: {}
- npm@1:
    inputs:
    - command: install
- save-npm-cache@1: {}
```
