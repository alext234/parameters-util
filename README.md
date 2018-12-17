[![Build Status](https://travis-ci.org/alext234/parameters-util.svg?branch=master)](https://travis-ci.org/alext234/parameters-util)

## parameters-util
A small utility for retrieving AWS Parameter Store and output to files or directories


## Development

### To release a new version 

A git tag will trigger the Travis pipeline to build and upload new version

```
git tag <release-tag>  master
git push --tags

```


### How is api key for `.travis.yml`'s deploy generated ?

It was generated with the travis cli 

```
travis setup releases --org
```



## TODO

- create a dir aws-param-store

- could be used as a lib

- could be used as a cli utility

- output to file option

- decode from base64 option (to a directory hierarchy)
