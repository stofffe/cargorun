package core

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func Run(kind, name string, extraArgs []string) error {
	args := []string{"run", fmt.Sprintf("--%s", kind), name}
	args = append(args, extraArgs...)
	cmd := exec.Command("cargo", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not run example: %v", err)
	}
	return nil
}

func GetRunnable(kind string) ([]string, error) {
	cmd := exec.Command("cargo", "read-manifest")
	bytes, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not run cargo read-manifest: %v", err)
	}

	var manifest manifest
	err = json.Unmarshal(bytes, &manifest)
	if err != nil {
		return nil, fmt.Errorf("could not parse json response: %v", err)
	}

	var examples []string
	for _, target := range manifest.Targets {
		for _, k := range target.Kind {
			if kind == k {
				examples = append(examples, target.Name)
				break
			}
		}
	}

	return examples, nil
}

type manifest struct {
	Targets []target `json:"targets"`
}

type target struct {
	Name string   `json:"name"`
	Kind []string `json:"kind"`
}
