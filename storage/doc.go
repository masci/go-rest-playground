/*

The `storage` package implements the data layer for our API.

The idea behind the design of the package is to make it loosely coupled to the rest of the
application, in order to ease further development and refactoring as the requirements evolve.

In order to give a better idea of the design, the exercise provides two different ways of
storing data that are transparent to the web service logic serving our API. The application
can store data either in memory, using a simple map that will be lost on exit, or in a SQLite
database persisted on file.

┌───────────────────┐      ┌───────────────────┐
│                   │      │                   │
│    Application    │─────▶│ Storage interface │
│                   │      │                   │
└───────────────────┘      └───────────────────┘
                                     ▲
                       ┌─────────────┴─────────────┐
                       │                           │
             ┌───────────────────┐       ┌───────────────────┐
             │      In-mem       │       │      SQLite       │
             │  implementation   │       │  implementation   │
             └───────────────────┘       └───────────────────┘

*/
package storage
