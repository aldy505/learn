# How to run the file

```sh
# make sure to have redis and postgres up and running
# make sure to set your pwd to this folder
$ julia

julia> ]

(@v1.6) pkg> activate .

(database) pkg> instantiate

# Ctrl + C or exit()

$ julia postgres.jl
$ julia redis.jl
```