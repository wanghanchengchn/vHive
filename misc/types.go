// MIT License
//
// Copyright (c) 2020 Dmitrii Ustiugov and EASE lab
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package misc

import (
	"fmt"
	"github.com/ease-lab/vhive/networking"
	"sync"

	"github.com/containerd/containerd"
)

const (
	defaultVcpuCount = 1
	defaultMemsizeMib = 256
)

// VM type
type VM struct {
	ID               string
	ContainerSnapKey string
	SnapBooted       bool
	Image            *containerd.Image
	Container        *containerd.Container
	Task             *containerd.Task
	TaskCh           <-chan containerd.ExitStatus
	NetConfig        *networking.NetworkConfig
	VCPUCount        uint32
	MemSizeMib       uint32
}

// VMPool Pool of active VMs (can be in several states though)
type VMPool struct {
	vmMap      sync.Map
	networkManager *networking.NetworkManager
}

// NewVM Initialize a VM
func NewVM(vmID string) *VM {
	vm := new(VM)
	vm.ID = vmID
	vm.ContainerSnapKey = fmt.Sprintf("vm%s-containersnap", vmID)
	vm.SnapBooted = false
	vm.MemSizeMib = defaultMemsizeMib
	vm.VCPUCount = defaultVcpuCount

	return vm
}
