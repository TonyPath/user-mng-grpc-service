package dockertest

import (
	"context"
	"io"
	"net"

	// 3rd party
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const postgresImage = "postgres:14.4-alpine"

func SetupPostgres(envArgs []string) (teardown func(), hostAndPort string, err error) {
	ctx := context.Background()
	dc, _ := client.NewClientWithOpts(client.FromEnv)

	// pull the image
	reader, err := dc.ImagePull(ctx, postgresImage, types.ImagePullOptions{})
	if err != nil {
		return
	}
	_, _ = io.Copy(io.Discard, reader)

	// create container
	cc := container.Config{
		Image: postgresImage,
		Env:   envArgs,
	}
	hc := container.HostConfig{
		PublishAllPorts: true,
	}

	resp, err := dc.ContainerCreate(ctx, &cc, &hc, nil, nil, "")
	if err != nil {
		return
	}

	// run container
	err = dc.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return
	}

	containerID := resp.ID

	teardown = func() {
		if err := dc.ContainerStop(ctx, containerID, nil); err != nil {
			panic(err)
		}

		removeOptions := types.ContainerRemoveOptions{
			RemoveVolumes: true,
		}
		if err = dc.ContainerRemove(ctx, containerID, removeOptions); err != nil {
			panic(err)
		}
	}

	cnt, _ := dc.ContainerInspect(ctx, containerID)

	hostPorts := cnt.NetworkSettings.NetworkSettingsBase.Ports["5432/tcp"]
	host := hostPorts[0].HostIP
	port := hostPorts[0].HostPort

	hostAndPort = net.JoinHostPort(host, port)

	return
}
