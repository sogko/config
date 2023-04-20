# config
My go-to Go package for my Go apps, based on my highly specific and opinionated configuration needs.

Built upon [github.com/spf13/viper](https://github.com/spf13/viper) 

It's probably not for everyone.

# Motivation

Been using spf13/viper for a while now, and I've been using it in a very specific way in many of my Go projects.

Instead of creating a new internal config package everytime, I decided to create a single package that I can use in all my projects.

This package might be a perfect match for your needs, or perhaps not quite. Either way, feel free to give it a try or pass it by.

Suggestions and feedback are always appreciated. 


## Installation
```shell
go get github.com/sogko/config
```

## Getting started

### Config file
Right out of the box, the package expects a `config.dev.json` file relative to the current working directory.

```json
{
  "foo": "bar"
}
```

### Usage
```go
package main

import (
	"fmt"
	"github.com/sogko/config"
)

func main() {

	// Load configuration
	cfg := config.Load()

	// Get configuration value
	foo := cfg.GetString("foo")

	// Print configuration value
	fmt.Printf("foo: %s\n", foo)
	// Output: foo: bar

}

```

---

## Configuration file
### File format
The package expects a configuration file in JSON format.

### File name
The package expects a configuration file named `config.<environment>.json`. 

By default, the environment is set to `dev` (i.e. it will look for `config.dev.json` file).

By setting the `ENV` environment variable, you can change the environment to `prod` or `staging` or whatever you want,
and the package will look for `config.prod.json` file, or `config.staging.json` file, etc.

You can also set the `CONFIG` environment variable to specify the name and location of the configuration file (accepts both relative and absolute paths).

### File location
The package expects a configuration file to exist in the current working directory.

You can also set the `CONFIG` environment variable to specify the location of the configuration file (accepts both relative and absolute paths).

## Built-in configuration keys
- `CONFIG`: Configuration file path
- `ENV`: Environment
- `ENV_PREFIX`: Environment variable prefix

### `CONFIG`
Default: `<not set>`

Accepts both relative and absolute paths.

For relative paths, it is relative to the current working directory.

#### Example
```json
// Default: config.dev.json
{
  "foo": "bar"
}
// Relative path: ./my_custom_config.json
{
  "foo": "custom_bar"
}
// Absolute path: /User/alice/config.prod.json
{
  "foo": "barber"
}
```
```go

func main() {
    cfg := config.Load()
    fmt.Printf("foo: %s\n", cfg.GetString("foo"))
}

```
```shell
$ go run main.go
# Output: foo: bar

$ env CONFIG=./my_custom_config go run main.go
# Output: foo: custom_bar

$ env CONFIG=/User/alice/config.prod.json go run main.go
# Output: foo: barber

```

### `ENV`

Default: `dev`

Alias: `ENVIRONMENT`

#### Example
Here's an example of setting the environment to `prod` via `ENV` environment variable.
```json
// config.dev.json
{
  "foo": "bar"
}
// config.prod.json
{
  "foo": "barber"
}
```
```go

func main() {
    cfg := config.Load()
    fmt.Printf("foo: %s\n", cfg.GetString("foo"))
}

```
```shell
$ env ENV=prod go run main.go
# Output: foo: barber
```


### `ENV_PREFIX`

Default: `<not set>`

You can set a prefix for your environment variables via `ENV_PREFIX`, 
so that they don't conflict with other apps.

By default, it is not set.

#### Example
Here's a basic example of how you can set `ENV_PREFIX` to  `MYAPP` and then set `MYAPP_FOO` environment variable to `barber`.
Note that environment variables takes precedence over config file values.

```json
// config.dev.json
{
  "foo": "bar"
}
```

```go

func main() {
    cfg := config.Load()
    fmt.Printf("foo: %s\n", cfg.GetString("foo"))
}

```
```shell
$ env ENV_PREFIX=MYAPP MYAPP_FOO=barber go run main.go
# Output: foo: barber
```

Once `ENV_PREFIX` is set, the built-in environment keys (`ENV` and `CONFIG`) will be prefixed as well.

```json
// config.dev.json
{
  "foo": "bar"
}
// config.prod.json
{
  "foo": "barber"
}
```
```go

func main() {
    cfg := config.Load()
    fmt.Printf("foo: %s\n", cfg.GetString("foo"))
}

```
```shell
$ env ENV_PREFIX=MYAPP MYAPP_ENV=prod go run main.go
# Output: foo: barber

$ env ENV_PREFIX=MYAPP MYAPP_CONFIG=./config.prod.json go run main.go
# Output: foo: barber
```


## API Reference

In addition to [github.com/spf13/viper API](https://github.com/spf13/viper), the package also exposes the following methods:

### `Load()`
Loads configuration from file and environment variables.

#### Example
```go  
func main() {
    
    // Load configuration
    cfg := config.Load()
	
	fmt.Printf("foo: %s\n", cfg.GetString("foo"))
    // Output: foo: bar
}
``` 

### `Load()`
Loads configuration from file and environment variables.

#### Example
```go  
func main() {
    
    // Load configuration
    cfg := config.Load()
	
	fmt.Printf("foo: %s\n", cfg.GetString("foo"))
    // Output: foo: bar
}
``` 

### `Reload()`
Reloads configuration from file and environment variables.

This is useful if you want to manually reload configuration after changing the configuration file.

### `Watch()`
Watches for changes in the configuration file and reloads configuration when changes are detected.

### `GetConfigFile()`
Returns the path of the configuration file.

### `Config` struct
Wraps `viper.Viper` struct.

All of Viper API methods are available.

#### `cfg.Save(key string, val interface{})`
Saves a configuration value to the configuration file.

Note: this method calls `cfg.WriteConfig()` internally, which will overwrite the configuration file and writes all configuration keys and its values.