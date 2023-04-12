# REDIS quick start in MAC

## Install
### install
```
brew install redis
```
### check Install
```
redis-server --version
```

## Start
### start in foreground 
```
redis-server
```

### start in background
```
// start
brew service start redis

// restart
brew service restart redis

// stop
brew service stop redis
```

## redis cli
### run
```
redis-cli
```
### insert, update
```
set [key] [value]

//ex
set name john
```

### select
```
get [key]

//ex
get name
```

### select all
```
keys *
```

### update key
```
rename [key] [key name to change]

//ex
rename name firstname
```

### key count
```
dbsize
```

### delete
```
del [key]

//ex
del firstname
```

### flush
```
flushall
```
