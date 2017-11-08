package provisioner

import (
	"testing"

	"github.com/kubernetes-incubator/external-storage/lib/controller"
)

func TestGetOptionsForVolume(t *testing.T) {
	volume := getOptionsForVolume(100, "", "myvolume")
	if volume.DisplayName != "myvolume" {
		t.Fatalf("Incorrect display name. Expecting %s but got %s", "myvolume", volume.DisplayName)
	}
}

func TestGetGetOptionsForVolumeCustomDisplayName(t *testing.T) {
	volume := getOptionsForVolume(100, "XXX", "myvolume")

	if volume.DisplayName != "XXXmyvolume" {
		t.Fatalf("Incorrect display name. Expecting %s but got %s", "XXXmyvolume", volume.DisplayName)
	}
}

func TestResolveFSTypeWhenNotConfigured(t *testing.T) {
	options := controller.VolumeOptions{Parameters: make(map[string]string)}
	// test default fsType of 'ext4' is always returned.
	fst := resolveFSType(options)
	if fst != "ext4" {
		t.Fatalf("Unexpected filesystem type: '%s'.", fst)
	}
}

func TestResolveFSTypeWhenConfigured(t *testing.T) {
	// test default fsType of 'ext3' is always returned when configured.
	options := controller.VolumeOptions{Parameters: map[string]string{fsType: "ext3"}}
	fst := resolveFSType(options)
	if fst != "ext3" {
		t.Fatalf("Unexpected filesystem type: '%s'.", fst)
	}
}
