# plugins

A repository containing plugins for sourcemod for use with [sm-pkg](https://github.com/sm-pkg/sm-pkg).

## Plugin Package Definitions

Each plugin comes with a package definition file named `package.json`. This file contains metadata about the plugin, such as its name, version, authors, description, homepage, license, and dependencies.

Example `package.json`

```json
{
  "$schema": "https://raw.githubusercontent.com/sm-pkg/plugins/refs/heads/master/schema.json",
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
