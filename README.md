# plugins

A repository containing plugins for sourcemod for use with [sm-pkg](https://github.com/sm-pkg/sm-pkg).

## Plugin Package Definitions

Each plugin comes with a plugin definition file named `plugin.json`. This file contains metadata about the plugin, such as its name, version, authors, description, homepage, license, and dependencies.

Example `plugin.json`

```json
{
  "$schema": "https://raw.githubusercontent.com/sm-pkg/sm-pkg/refs/heads/master/schema/plugin.json",
  "name": "plugin_name",
  "homepage": "https://github.com/username/plugin",
  "description": "An example plugin definition.",
  "license": "MIT",
  "version": "0.1.2",
  "inputs": ["plugin_name.sp"],
  "authors": ["Someone Somewhere"],
  "dependencies": ["another_package_name"]
}
```
