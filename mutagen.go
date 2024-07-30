// A generated module for Mutagen functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"dagger/mutagen/internal/dagger"
	"errors"
	"fmt"
	"os"
)

type Mutagen struct {
}

func (m *Mutagen) Agent(
	volume string,
	// +optional
	authorizedKeys *dagger.File,
	// +optional
	publicKey string,
) (*dagger.Service, error) {

	if publicKey != "" {
		authorizedKeys = dag.
			Directory().
			WithNewFile("authorized_keys", publicKey).
			File("authorized_keys")
	}
	if authorizedKeys == nil {
		return nil, errors.New("authorizedKeys or publicKey must be provided")
	}

	fmt.Fprintf(os.Stderr, "Agent created. Run:\n\n")
	fmt.Fprintf(os.Stderr, "\tmutagen sync create --name=%s <path> root@localhost:<port>:~/dagger\n", volume)
	fmt.Fprintf(os.Stderr, "\tmutagen sync flush %s\n\n", volume)
	return dag.Container().
		From("hermsi/alpine-sshd").
		WithEnvVariable("ROOT_KEYPAIR_LOGIN_ENABLED", "true").
		WithFile("/root/.ssh/authorized_keys", authorizedKeys, dagger.ContainerWithFileOpts{
			Permissions: 0600,
		}).
		WithMountedCache("/root/dagger", dag.CacheVolume(volume), dagger.ContainerWithMountedCacheOpts{
			Sharing: dagger.Shared,
		}).
		WithExec([]string{}, dagger.ContainerWithExecOpts{
			UseEntrypoint: true,
		}).
		WithExposedPort(22).
		AsService(), nil
}
