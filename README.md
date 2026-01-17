# smpkg-repo

A repository containing plugins for sourcemod for use with [smpkg](https://github.com/leighmacdonald/smpkg).

## Plugin Package Definitions

Each plugin comes with a package definition file named `package.yaml`. This file contains metadata about the plugin, such as its name, version, authors, description, 
homepage, license, and dependencies.

Example smpkg.yaml

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/leighmacdonald/smpkg-repo/refs/heads/master/schema.yaml

name: plugin_name
version: 0.1.0
authors:
  - Someone Somewhere
description: |
  A plugin by someone somewhere.
homepage: https://github.com/username/plugin
license: MIT
dependencies: ["another-package-name"]
inputs: ["script.sp"]
```
