// Package main is the Canasta-CLI Go-to-Ansible bridge.
//
// Old Canasta-CLI Go binaries' selfupdate logic is hardcoded to query
// api.github.com/repos/CanastaWiki/Canasta-CLI/releases/latest and
// download canasta-<os>-<arch> from whichever release is "latest."
// Attaching this minimal bridge as a release asset (a) to v3.7.0 of
// this repo, and (b) to every release of the Ansible-based Canasta CLI
// (the renamed Canasta-Ansible repo, which now occupies the
// "Canasta-CLI" slot on GitHub) keeps the migration path working
// indefinitely.
//
// Any invocation of this binary prints the migration message and
// exits. Once installed via the old binary's selfupdate, running any
// canasta command becomes a forcing function for the user to run the
// install script.
package main

import "fmt"

func main() {
	fmt.Println(`The Go implementation of Canasta CLI that you have been using has been superseded by a python/Ansible implementation (version 4.0.0+).

Your registered instances and their data continue to work without changes.

To install the new Canasta CLI, run ONE of:

    # Docker mode (default — only Docker required on the host):
    curl -fsSL https://get.canasta.wiki | bash

    # Native mode (requires Python 3.10+ and git on the host;
    # Ansible is installed automatically into a Python venv):
    curl -fsSL https://get.canasta.wiki | bash -s -- --native

Docker mode is the simplest setup. Native mode runs Ansible directly
on the host (faster startup, easier playbook debugging).

See https://canasta.wiki/wiki/Help:Installation for details.

After install, run 'canasta version' to confirm you're on 4.0.0 or later.`)
}
