package client

import (
	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wire"
)

func (client *Client) CreateDirectory(path string, options *option.CreateDirectory) error {
	return client.post(join(pathsPrefix, path, createDirectory), nil, options, nil)
}

func (client *Client) CreateFile(path string, options *option.CreateFile) (int, error) {
	var result int
	if err := client.post(join(pathsPrefix, path, createFile), nil, options, &result); err != nil {
		return -1, err
	}
	return result, nil
}

func (client *Client) Delete(path string, options *option.Delete) error {
	return client.post(join(pathsPrefix, path, delete), nil, options, nil)
}

func (client *Client) Exists(path string, options *option.Exists) (bool, error) {
	var result bool
	if err := client.post(join(pathsPrefix, path, exists), nil, options, &result); err != nil {
		return false, err
	}
	return result, nil
}

func (client *Client) Free(path string, options *option.Free) error {
	return client.post(join(pathsPrefix, path, free), nil, options, nil)
}

func (client *Client) GetStatus(path string, options *option.GetStatus) (*wire.FileInfo, error) {
	var result wire.FileInfo
	if err := client.post(join(pathsPrefix, path, getStatus), nil, options, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *Client) ListStatus(path string, options *option.ListStatus) (wire.FileInfos, error) {
	var result wire.FileInfos
	if err := client.post(join(pathsPrefix, path, listStatus), nil, options, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) Mount(path, src string, options *option.Mount) error {
	params := map[string]string{
		"src": src,
	}
	return client.post(join(pathsPrefix, path, mount), params, options, nil)
}

func (client *Client) OpenFile(path string, options *option.OpenFile) (int, error) {
	var result int
	if err := client.post(join(pathsPrefix, path, openFile), nil, options, &result); err != nil {
		return -1, err
	}
	return result, nil
}

func (client *Client) Rename(path, dst string, options *option.Rename) error {
	params := map[string]string{
		"dst": dst,
	}
	return client.post(join(pathsPrefix, path, rename), params, options, nil)
}

func (client *Client) SetAttribute(path string, options *option.SetAttribute) error {
	return client.post(join(pathsPrefix, path, setAttribute), nil, options, nil)
}

func (client *Client) Unmount(path string, options *option.Unmount) error {
	return client.post(join(pathsPrefix, path, unmount), nil, options, nil)
}
