# pants-go-cyclic-dependency

The following command does not cause a `CycleException`:

```
pants package :bin0
```

The following command **does**:

```
pants package :docker0
```

It should not do this, because the only thing it needs to do is to copy the artifact of target `bin0`. There is no cyclic dependency.
