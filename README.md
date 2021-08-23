# Packer Plugin HashiCups

This repo is part of the [Packer](https://learn.hashicorp.com/packer) Learn collection. The intent of this plugin is to help you create your own packer plugin.

Refer to the [documentation](docs) to learn about the HashiCups plugin and how it works.


## Test sample configuration

First, you will need the demo HashiCups API up and running.

```shell
$ make run-hashicups-api
```

This will run `docker-compose up -d` in [example/hashicups_api](example/hashicups_api).

Sign up to the HashiCups api.

```shell
$ curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}'
```

Then, navigate to the [example](exmaple) directory.

```shell
$ cd example
```

Run the following command to initialize and build the sample configuration.

```shell
$ packer init . && packer build .
```
