
#### Minimal example
```yaml
steps:
- restore-npm-cache@1: {}
- npm@1:
    inputs:
    - command: install
- save-npm-cache@1: {}
```

Check out [Workflow Recipes](https://github.com/bitrise-io/workflow-recipes#-key-based-caching-beta) for other caching examples!
