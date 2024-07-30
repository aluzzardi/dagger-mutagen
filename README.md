# Dagger Mutagen Integration

Sync local directories to dagger using mutagen.

- Under the hood, it uses the [SSH Transport](https://mutagen.io/documentation/transports/ssh) to transfer data.
- Directories are synced to a `dagger.CacheVolume` which is accessible in dagger containers
- Currently this requires running a mutagen agent as a dagger service

## Usage

**Start a mutagen agent service for the `MyCode` cache volume**

```
dagger call -m github.com/aluzzardi/dagger-mutagen agent \
  --volume MyCode \
  --authorized_keys ~/.ssh/id_ed25519.pub \
  up --ports 1222:22
```

**Create a mutagen sync and monitor progress**

```
mutagen sync create --name=MyCode ../app/cloud root@localhost:1222:~/dagger
mutagen sync monitor MyCode
```

**Use the `MyCode` cache volume in your own service**

```go
func (m *Livedev) Dev(ctx context.Context) *dagger.Service {
	code := dag.CacheVolume("MyCode")

	return dag.Container().
		From("python").
		WithWorkdir("/code").
		WithMountedCache("/code", code, dagger.ContainerWithMountedCacheOpts{
			Sharing: dagger.Shared,
		}).
		WithExec([]string{"python", "-m", "http.server", "8080"}).
		WithExposedPort(8080).
		AsService()
}
```

```
dagger call dev up --ports 8080:8080
```
