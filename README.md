# Tickit

[![Build Status](https://travis-ci.org/itmecho/tickit.svg?branch=master)](https://travis-ci.org/itmecho/tickit)

Tickit is a simple command line todo list. It was written to get used to the Go programming language.

## Usage

To get started, run the init command to generate the list file:

```
tickit init
```

Next, add a few tasks and list them

```
$ tickit add "This is my first task"
$ tickit add "This is my second task"
$ tickit list
[ ] - 1: This is my first task []
[ ] - 2: This is my second task []
```
You can create a task with tags by using the -t flag on the add command

```
$ tickit add "This is my third task with tags" -t tag1 -t tag2
$ tickit list
[ ] - 1: This is my first task []
[ ] - 2: This is my second task []
[ ] - 3: This is my third task with tags [tag1,tag2]
```

You can mark a task as complete by using the close command and passing in the task id.

```
$ tickit close 1
$ tickit close 2
$ tickit list
[X] - 1: This is my first task []
[X] - 2: This is my second task []
[ ] - 3: This is my third task with tags [tag1,tag2]
```

If you need to undo this, you can use the reopen command

```
$ tickit reopen 2
$ tickit list
[X] - 1: This is my first task []
[ ] - 2: This is my second task []
[ ] - 3: This is my third task with tags [tag1,tag2]
```
