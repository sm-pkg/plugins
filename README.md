# plugins

A repository containing plugins for sourcemod for use with [sm-pkg](https://github.com/sm-pkg/sm-pkg).

## Plugin Package Definitions

Each plugin comes with a plugin definition file named `plugin.yaml`. This file contains metadata about the plugin, such as its name, version, authors, description, homepage, license, and dependencies.

Example `plugin.yaml`

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/sm-pkg/sm-pkg/refs/heads/master/schema/plugin.json
name: sm_vprofiler
description: Measures per-plugin performance and provides a log with various counters
homepage: "https://github.com/dragokas"
authors:
  - Moonly Days
inputs:
  - sm_vprofiler.sp
version: 1.0.0

```
