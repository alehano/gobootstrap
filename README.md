# Go Bootstrap

# WARNING: It's under development

Web framework in Go focused on simplicity and flexibility. 

#### Main features:

* Nested structure
* Lose coupled modules
* Pluggable DB layer
* Memcache support
* CLI commands support
* URL reverse
* JWT Auth support
* Type safe config

After cloning repo you have to rename any `github.com/alehano/gobootstrap` to 
`{your project path}`. Feel free to modify source code as you want.

## Project Structure:

Project has a nested structure. Top level parts can be used by many lower level parts.

For example, top level `/helpers` package should contain more common helpers 
than more specific lower level `/views/contacts/helpers.go`.
Also, files under `common` dirs considered as upper level's. For example,
`/views/common/tmpl` dir should contain files using by multiple views, but
`/views/admin/tmpl` should contain files using only from `/admin` section.

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
static files (img, js, css) and more. Handlers mostly use model Managers to
get and save data. 


## Example config file `config.yml`

```yaml
debug: true
port: 8000
project_path: "/Users/alehano/Development/go/src/github.com/alehano/gobootstrap"
website_protocol: "https://"
website_domain: "example.com"
admin_login: "admin"
admin_password_hash: "$2a$10$5gCLP.GlOVBVFLtrzhoxfO5wsT0eiH7IsjzupA7ukTyI/znLFotHu" # password is "admin"
jwt_secret: "secret"
```

## Some of used external packages:

* [Chi](https://github.com/go-chi/chi)
* [Cobra](https://github.com/spf13/cobra)
* [reverse](https://github.com/alehano/reverse)
* [gohelpers](https://github.com/alehano/gohelpers)
* [govalidator](https://github.com/asaskevich/govalidator)
* [gomemcache](https://github.com/bradfitz/gomemcache)
* [sqlx](https://github.com/jmoiron/sqlx)
* [jwt-go](https://github.com/dgrijalva/jwt-go)
