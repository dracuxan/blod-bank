package helper

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	blodBank "github.com/dracuxan/blod-bank/proto"
)

func GetConfig(id *blodBank.ConfigID, c blodBank.BlodBankServiceClient) (*blodBank.ConfigItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := c.GetConfig(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ListAllConfig(c blodBank.BlodBankServiceClient) ([]*blodBank.ConfigItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := c.ListAllConfig(ctx, &blodBank.NoParam{})
	var configs []*blodBank.ConfigItem

	if err != nil {
		log.Fatalf("Cannot get config list: %v", err)
	}

	for {
		config, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListConfigs failed: %v", err)
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func RegisterConfig(item *blodBank.ConfigItem, c blodBank.BlodBankServiceClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	status, err := c.RegisterConfig(ctx, item)
	if err != nil {
		return "", err
	}
	return status.Status, nil
}

func DeleteConfig(id *blodBank.ConfigID, c blodBank.BlodBankServiceClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	status, err := c.DeleteConfig(ctx, id)
	if err != nil {
		return "", err
	}
	return status.Status, nil
}

func UpdateConfig(item *blodBank.ConfigItem, c blodBank.BlodBankServiceClient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	status, err := c.UpdateConfig(ctx, item)
	if err != nil {
		return "", err
	}
	return status.Status, nil
}

func OpenEditor(filename string) (string, error) {
	editor := "vim"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return "", err
	}
	f.Close()

	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the editor and wait for it to close
	fmt.Printf("Opening %s with %s...\n", filename, editor)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return filename, nil
}

func CreateConfig(regName, filename string) (*blodBank.ConfigItem, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &blodBank.ConfigItem{
		Name:    regName,
		Content: string(content),
	}, nil
}
