# Yet Another Golang RESTfull API

### A quick story about this starter-kit
So, when I tried to write my first [golang][1] project, especially for `web api`, I realy don't know how to structure my project directory, file management, etc. Frequently I search for hours on [Github][9] or `googling` regarding the best practice on structuring our [Go][1] project, but the deeper I dive, the confused I get. Technically, most of the `Already Available` starter kit that I found is good, but sometimes it overkill. Some of them also great, but missing some part. Therefore I made my own so it will also aligned with code my writting style as well as minimalizing the 'confusion' I may get when depending on others starter-kit. This starter-kit may not for you as writting/ programming style is personal prefference and relative for every person, but feel free to use it.

## Table of Content
1. [Quick Review](#1.-quick-review)
2. [Directory Structure](#2.-directory-structure)
3. [Getting Started](#3.-getting-started)

### 1. Quick review
This `YET ANOTHER` [golang][1] RESTFull API is using below module:
1. [Gin][2] web framework, fast and easy
2. [Gorm][3] ORM, so you doesn't need to dive to deep on the `SQL`. Support for `PostgreSQL`, `MySQL` or `MariaDB`, and `Sqlite3`
3. [Viper][8] for easy configuration, support `yaml`, `toml`, `json`, and more...
4. [Golang JWT][5] for authentification
5. [Logrus][7] for nice and easy logging feature from [sirupsen][7]

Writting pattern: 
```bash
model-> repository-> service-> handler-> router (api endpoint)
```
Back to [Table of Content](#table-of-content) or back to [Top](#yet-another-golang-restfull-api)

### 2. Directory structure
```
project-directory
|-- cmd
|-- |-- app
|-- |-- | server
|-- |-- docker
|-- config
|-- internal
|-- |-- account
|-- |-- |-- handler
|-- |-- |-- model
|-- |-- |-- repository
|-- |-- |-- router
|-- |-- |-- service
|-- |-- blog
|-- |-- |-- handler
|-- |-- |-- model
|-- |-- |-- repository
|-- |-- |-- router
|-- |-- |-- service
|-- |-- database
|-- |-- |-- error 
|-- |-- |-- model
|-- |-- pkg
|-- pkg
|-- |-- logger
|-- |-- middleware
|-- vendor
|-- go.mod
|-- go.sum
|-- README.md
|-- Makefile
```
Back to [Table of Content](#table-of-content) or back to [Top](#yet-another-golang-restfull-api)

### 3. Getting started
To run the server:
```bash
make run
```
To build the project
```bash
make build
```
Back to [Table of Content](#table-of-content) or back to [Top](#yet-another-golang-restfull-api)

## LICENSE
[MIT](https://github.com/reshimahendra/gin-starter/blob/main/LICENSE)

[1]:https://golang.org
[2]:https://gin-gonic.com
[3]:https://gorm.io
[4]:https://github.com/
[5]:https://github.com/golang-jwt/jwt
[6]:https://github.com/google/uuid
[7]:https://github.com/sirupsen/logrus
[8]:https://github.com/spf13/viper
[9]:https://github.com
