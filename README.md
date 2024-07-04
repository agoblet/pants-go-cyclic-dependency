# pants-go-cyclic-dependency

This repo is a minimal example highlighting the following Pants issue: https://github.com/pantsbuild/pants/issues/21131

The following command does not cause a `CycleException`:

```
pants package :bin0
```

The following command **does**:

```
pants package :docker0
```

It should not do this, because the only thing it needs to do is to copy the artifact of target `bin0`. There is no cyclic dependency.
