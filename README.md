# boilr

The boilr boilerplating tool.


## How To Run

```
boilr -in <path to boilr.plate-file> -out <output folder>
```

## Defining Templates

Templates consist of a `.plate` file and an accompanying folder holding the
actual templates. 

### `boiler.plate` files

These Files hold the definition for input variables and the Path to you're
template folder. It is a simple yml file that only supports a single level.
They Look like this:

```
---
foo: string
bar: list
baz: string
TEMPLATE_ROOT: bar
```

`TEMPLATE_ROOT` is a special key that always needs to be set. There you specify
the path to the template folder relative to the .plate file. You don't need to
add `./` or a trailing `/`.

All Other keys are used as input variables. They support two values: `string`
and `list`. boilr will query the user for these values before rendering the
template. 


### Templates

boilr uses [pongo2](https://github.com/flosch/pongo2) to render its templates.
Every file that has .j2 as its ending will be rendered as a template. The
Resulting file will have the .j2 suffix stripped. 

boilr will make every variable you've defined available in each template.

#### Templating Folders and Files from lists

Use the following folder or file name to create multiple files based on a list var:

```
__for_var_in_listvar__
```
Let's say you have a list variable called `environments` defined in your service.plate. 
A folder called `__for_env_in_environments__` will create a new folder with the environment name containing the contents of that folder. 

All templates withing that folder will have access to the variable `env`, which holds the current environment. 

#### Templating in file

Boilr also supports templates in path names. This means, that arbitrary
templates are also allowed in file and folder names.

The following var is set for boilr: `env: production`.  A file in the template
called: `/environments/{{ env }}/deploy.py` will then be rendered to
`environments/productions/deploy.py`.


## Building 

This Project uses go-modules for building: [docs](https://github.com/golang/go/wiki/Modules).

This means, that when you checkout this repo in your `$GOPATH`, you'll need to set the env var:
```
export GO111MODULE=on
```
Directories outside of your $GOPATH already have go-modules enabled (unless explicitly disabled by the env var).

To build and install all requirements, just run the following:
```
go build
```

To make it available in your `$GOPATH/bin` run `go install`.

