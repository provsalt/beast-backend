package provider

import (
	"fmt"
	"math"
	"strconv"
)

type Provider struct {
	// workerName is the name provided by the mining software via rig-id
	WorkerName string
	// hashrate is the current user hashrate.
	Hashrate int

	// Contact is the information used to contact the user.
	Contact string

	// This will check indicate if the worker went online or offline
	Online bool
}

// GetHashrate ...
func (p Provider) GetHashrate() int {
	return p.Hashrate
}

// GetWorkerName ...
func (p Provider) GetWorkerName() string {
	return p.WorkerName
}

// ContactInfo ...
func (p Provider) ContactInfo() string {
	return p.Contact
}

// Send ...
func (p Provider) Send() bool {
	return false
}

// ToHR converts numerical hashrate to a readable string
func ToHR(hashrate int) string {
	symbol := []string{" h/s", " Kh/s", " Mh/s", " Gh/s", " Th/s", " Ph/s", " Eh/s"}

	tier := math.Log10(math.Abs(float64(hashrate))) / 3

	if tier == 0 {
		return strconv.Itoa(hashrate) + " h/s"
	}
	suffix := symbol[int(tier)]
	scale := math.Pow10(int(tier * 3))

	scaled := float64(hashrate) / scale

	return fmt.Sprintf("%.2f %s", scaled, suffix)
}
