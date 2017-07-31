# Go Bootstrap

# WARNING: It's under development

Web framework in Go focused on flexibility and scalability 
due to splitting by many lose coupled packages. 

**Main features:**

* Nested structure
* Lose coupled modules
* Pluggable DB layer
* Memcache support
* CLI commands support
* URL reverse
* Type safe config

After cloning repo you have to rename any `github.com/alehano/gobootstrap` to 
`{your project path}`

Project has a nested structure. Top level parts can use by many lower level parts.

For example, top level `/helpers` package should contain more common helpers 
than more specific lower level `/views/contacts/helpers.go`.
Also, files under `common` dir considered as a upper level. For example,
`/views/common/static` dir should contain files, used in multiple views, but
`/views/admin/static` should contain files used only in `/admin` section.

## Project Structure:

#### actions
Complex tasks usually involving multiple models interactions.    

#### config
Centralized storage of config parameters and some others items
using in many parts of the App, like Cache Keys.

#### helpers
Set of App specific helpers methods. Independent from other parts of the App.
For more common helpers use `github.com/alehano/gohelpers`.

#### models
Data structures and data access layer (Manager) with DB persistence layer (Storage). 
Models are independent from other parts instead of Config. Models can send PubSub 
messages to communicate with each other and trigger other actions.

#### services
Background jobs.

#### sys
Core components of the App. Used in many others components.
Can depend on Config, Helpers and external packages.
Often contain stores of some items like URL routes or CLI commands.

#### utils
Utils more complex than helpers, but also independent from other parts of the App. 

#### views
Representation layers split by modules. Contains web handlers, templates, 
static files (img, js, css) and more. Handlers mostly use models Managers to
get and save data. 
