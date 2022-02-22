# Task Management Application

The goal was to create the **back-end** (REST API) part of the task management application (like [Trello](https://www.youtube.com/watch?v=noguPYxyv6g)) with no authentication and authorization required.

The main entity is a _Project_ (or _Board_) that always has its name and contains multiple _Columns_. _A Column_ always has a name, contains _Tasks_ and represents their status.

When a _Project_ is created at least one "default" _Column_ must be created in it as well.

A _Task_ can be created only inside the _Column_ and can be moved within the _Column_ (change priority) or across the _Columns_ (change status).

A _Task_ can have _Comments_ that could contain questions or _Task_ clarification information.

## Business Rules

_User_ must be able to manage (create, read, update, delete) _Projects_:

- _Projects_ in a list are sorted by name.

_Project_ must contain at least one column:

- the first column created by default when a _Project_ created;
- the last column cannot be deleted.

_User_ must be able to manage (create, read, update, delete) _Columns_:

- _Columns_ in a list are sorted by their position specified by _User_;
- _Column_ name **must** be unique;
- when a _Column_ is deleted its tasks are moved to the _Column_ to the left of the current.

_User_ must be able to move a _Column_ left or right.

_User_ must be able to rename a _Column_

_User_ must be able to manage (create, read, update, delete) _Tasks_:

- _Task_ can be created only within a _Column_;
- _User_ can view _Tasks_ in all _Columns_ of a _Project_;
- _User_ can update the name and the description of the _Task_;
- _User_ can delete the _Task_, with all _Comments_ related to this _Task_.

_User_ must be able to move a _Task_ across the _Columns_ (left or right) to change its status.

_User_ must be able to move a _Task_ within the _Column_ (up and down) to prioritize it.

_User_ must be able to manage (create, read, update, delete) _Comments_:

- _Comment_ can be created only within a _Task_;
- _Comments_ in a list are sorted by their creation date (from newest to oldest);
- _User_ can view _Comments_ related to a _Task_;
- _User_ can update the _Comment_ text;
- _User_ can delete the _Comment_.

## Fields and validation rules

**Project**

- Name (1-500 characters)
- Description (0-1000 characters)

**Column**

- Name/Status (1-255 characters, **unique** )

**Task**

- Name (1-500 characters)
- Description (0-5000 characters)

**Comment**

- Text (1-5000 characters)

## Using an application

To start Task-Management application locally, you need to:

```
make dev-up
make run
```

To start Task-Management application in simple production:

```
make prod-up
```

Also application uses these environment variables:

   - `PORT`: port to run app on. Example: PORT=8080
   - `ROOT_URL`: scheme and host of a running app. Example: `ROOT_URL="http://localhost"`
   - `DATABASE_URL`: data source name (DSN) for a PostgreSQL database. Example: `DATABASE_URL='postgres://user:password@localhost:5432/database'`
   - `MIGRATE`: specify if database needs to be migrated every time app is started. Example: `MIGRATE=true`
