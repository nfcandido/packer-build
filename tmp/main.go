// This is an example plugin. In packer 0.9.0 and up, core plugins are compiled
// into the main binary so these files are no longer necessary for the packer
// project.
//
// However, it is still possible to create a third-party plugin for packer that
// is distributed independently from the packer distribution. These continue to
// work in the same way. They will be loaded from the same directory as packer
// by looking for packer-[builder|provisioner|post-processor]-plugin-name. For
// example:
//
//    packer-builder-docker
//
// Look at command/plugin.go to see how the core plugins are loaded now, but the
// format below was used for packer <= 0.8.6 and is forward-compatible.
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"math"
	"os"
	"github.com/mitchellh/packer/builder/amazon/chroot"
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/mitchellh/packer/post-processor/docker-push"
	"github.com/mitchellh/packer/provisioner/powershell"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	// Choose the appropriate type of plugin. You should only use one of these
	// at a time, which means you will have a separate plugin for each builder,
	// provisioner, or post-processor.
	server.RegisterBuilder(new(chroot.Builder))
	server.RegisterPostProcessor(new(dockerpush.PostProcessor))
	server.RegisterProvisioner(new(powershell.Provisioner))
	server.Serve()
}

// https://gist.github.com/josephspurrier/e714fa55ae4c5ddfa668

// 8KB
const filechunk = 8192

func hash() {
	// Ensure the file argument is passed
	if len(os.Args) != 2 {
		fmt.Println("Please use this syntax: sha256 file.txt")
		return
	}

	// Open the file for reading
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Cannot find file:", os.Args[1])
		return
	}

	defer file.Close()

	// Get file info
	info, err := file.Stat()
	if err != nil {
		fmt.Println("Cannot access file:", os.Args[1])
		return
	}

	// Get the filesize
	filesize := info.Size()

	// Calculate the number of blocks
	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	// Start hash
	hash := sha256.New()

	// Check each block
	for i := uint64(0); i < blocks; i++ {
		// Calculate block size
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))

		// Make a buffer
		buf := make([]byte, blocksize)

		// Make a buffer
		file.Read(buf)

		// Write to the buffer
		io.WriteString(hash, string(buf))
	}

	// Output the results
	fmt.Printf("%x\n", hash.Sum(nil))
}