# Setup

Setup executable that prepares the database with some default data for the back-end to use.

> Note:
>
> This executable defines a flag by having more than 1 CLI argument being passed to it (the content doesn't matter).
> If there are no arguments being passed to it, it will read the database file from the current directory of the executable.
> If there are 1 or more arguments it will read from the parent directory.
>
> This kind of mechanism is necessary because there are cases that use both.
