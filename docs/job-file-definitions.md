1 - Main Section
-----------------

- `version`: initially '1' but might evolve based on how the project evolve
  - this is **optional** for now (default to the latest version number if absent, which is '1' for now)
- `name`: the name of the job. This will be used for job listing and some description commands (informative commands)
  - this is **mandatory**
- `description`: a more detailed explanation of the purpose of the job
  - this is **optional**
- `type`: 3 possible values: `sequential`, `parallel`, `orchestrator`
  - this is **optional** since hte default value (if not provided) is `sequential`
  - for now, the `orchestrator` value can be ignored since it won't be implemented soon.
    It could also become a more *cicrcle ci* like approach (all parallel, just put *require*
    to let the user define himself what kind of workflow is expected)
- `default`: define a list of setup that can be shared among different tasks.
  - this is a list of default. Each default has either:
    - `name` + `interpreter`
    - `name` + `docker` definition (`credential`, `registry`, `image`)
  - each name have to be unique
  - if there is only one default, the name can be omitted (will be replaced by the name `default`)
  - each parameter defined here can be over-written in each task
- `profiles`: a list of string that will be the name of the profile. Each task will hold a profile
  section to detail wait should be injected based on profile selection
- `tasks`: a list of task (define later)

2 - Tasks
---------

### 2.1 - Interpreter Task

- `name`: the name of the task. Should be unique
  - this is **mandatory**
- `interpreter`: what is the shell that should execute the command (or the source)
  - this is **mandatory**
- `source`: the script source location relative to the scrip app folder
  - this is **optional**
  - this is in conflict with **command**
  - one of `source` or `command` is **mandatory**
- `command`: the command to execute
  - this is **optional**
  - this is in conflict with **source**
  - one of `source` or `command` is **mandatory**
- `execute-in-current-dir`: default to `false` and use the project root dir
  - if set to `true` will use the current dir for the task execution
  - this is **optional** since it default to `false`
- `default`: if there is multiple default available for this kind or Task, then give 
   the name of the profile to use
  - this is **optional** unless an ambiguity is possible
- `termination-condition`: *manual* or *automatic*
  - default value is `automatic`
  - this is hence **optional** since a default is possible
- `profiles`: a list of profile object that can be used to inject parameters into the different builds