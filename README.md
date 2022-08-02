Planor
======
[![license](http://img.shields.io/badge/license-GPLv3-red.svg?style=flat)](https://raw.githubusercontent.com/mrusme/go-fleek/master/LICENSE)

```

                 ===   T H E   C L O U D   A V I A T O R   ===

⠀⠀⠀⠀⠀⠀⠀⠀⢶⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢻⣿⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣿⣿⣶⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣿⣿⣿⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⢿⣿⣿⣷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣿⣿⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠻⣿⣿⣿⣿⣶⣄⠀⠀⠀⠀⠄⠄⠒⠒⠒⠒⠒⠒⠄⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⣀⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⢿⣿⣿⣿⣿⣦⣤⣶⡀⠀⠀⠀⠀⠀⠀⢀⣠⡴⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠈⠻⢿⣷⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣽⣿⣿⣿⣿⣿⣿⣿⣷⣶⣶⣶⣶⠿⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠰⣾⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣴⣾⣿⣿⣿⣿⣿⣿⡿⢿⣿⣿⣿⣿⣿⣯⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣿⣿⣿⣿⣟⠛⠀⠀⠀⠀⢀⣀⣤⣶⣾⣿⣿⠿⠿⠛⠛⠉⠉⠉⠀⠀⠀⠀⠈⠉⠛⠿⣿⣿⣿⣿⣷⣦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣆⣠⣤⣶⣿⠿⠟⠛⠋⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠿⣿⣿⣿⣿⣷⣦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢹⣿⠿⠿⣿⠛⠋⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⢿⣿⣿⣿⣷⣶⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠿⢿⣿⣿⣿⣶⣤⣀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠛⠿⣿⣿⣿⣶⣤⣀⠀
/\\\\\\\   /\\         /\\\    /\\\\    /\\     /\\\\     /\\\\\\\  ⠈⠙⠛⠿⣿⣿⣶⣤
\/\\////\\\\/\\       /\\\\\\  \/\\\\\  \/\\    /\\//\\   \/\\///\\\     ⠈⠉⠛⠿⠿⠶
 \/\\  \/\\\\/\\      /\\///\\\ \/\\/\\\ \/\\  /\\/ \///\\ \/\\   \/\\
  \/\\\\\\\/ \/\\     \/\\  \/\\ \/\\//\\ \/\\ /\\     \//\\\/\\\\\\\/
   \/\\////   \/\\     \/\\\\\\\\ \/\\\/ \\\/\\\/\\      \/\\\/\\///\\\
    \/\\       \/\\     \/\\////\\ \/\\ \/\\\/\\\//\\     /\\ \/\\ \// \\
     \/\\       \/\\     \/\\  \/\\ \/\\ \//\\\\\ \///\\ /\\   \/\\  \ /\\
      \/\\       \/\\\\\\\\/\\  \/\\ \/\\  \//\\\\   \///\\\    \/\\  \ /\\
       \//        \/////// \//   \//  \//    \////      \///     \//   \///

```

Planor is a text user interface for cloud services. It currently supports the
following cloud service providers and cloud services:

- Amazon Web Services (AWS)
  - [x] Elastic Cloud Compute
  - [x] CodePipeline
  - [x] CloudWatch Logs
- Vultr
  - [x] Cloud Instances
- Heroku
  - [x] Dynos
  - [x] Builds
  - [ ] Logs
- Render.com
  - [x] Services (+ deploys states)
  - [ ] Logs
- Fleek
  - [x] Sites
  - [ ] Logs

---

![Continuous Integration](screen-ci.png)
![Logging](screen-logs.png)


## Installation

Either download a build from the releases page or clone this repository and run:

```sh
go build
```

or

```sh
go install
```


## Configuration

Nothing to configure, enjoy!


## Usage

Make sure to have the cloud provider profile configured and upon launching 
planor press r to refresh the data.


### Amazon Web Services

Create ~/.aws/config and add your profile, e.g.:

```ini
[profile captain-baloo]
region = us-east-1
```

Create ~/.aws/credentials and add your profile credentials, e.g.:

```ini
[captain-baloo]
aws_access_key_id = AKXXXXXHDXXXGXXPXXHX
aws_secret_access_key = qWX0Xx0XxxDxxx+0XsqXXLX/XXdXsxxMXxXlxKXv
```

Then run planor and specify the cloud service and profile name:

```sh
planor -c aws -p captain-baloo
```

Library: https://github.com/aws/aws-sdk-go-v2


### Vultr

Run planor and specify the environment variable that holds the Vultr API key as
profile:

```sh
export VULTR_API_KEY='...'
planor -c vultr -p VULTR_API_KEY
```

Library: https://github.com/vultr/govultr


### Heroku

Run planor and specify the environment variable that holds the Heroku API key as
profile:

```sh
export HEROKU_API_KEY='...'
planor -c heroku -p HEROKU_API_KEY
```

Library: https://github.com/heroku/heroku-go


### Render

Run planor and specify the environment variable that holds the Render API key as
profile:

```sh
export RENDER_API_KEY='...'
planor -c render -p RENDER_API_KEY
```

Library: https://github.com/mrusme/go-render


### Fleek

Run planor and specify the environment variable that holds the Fleek API key as
profile, and export the Fleek Team ID as environment variable as well:

```sh
export FLEEK_TEAM_ID='my-team'
export RENDER_API_KEY='...'
planor -c render -p RENDER_API_KEY
```

Library: https://github.com/mrusme/go-fleek



## Navigation

The keyboard navigation:

```
       r: Refresh
     C-p: Previous tab/service
     C-n: Next tab/service
  F1-F12: Switch to tab <nr>
     tab: Switch focus
       k: Move up in list
       j: Move down in list
       g: Move to the beginning of list/text
       G: Move to the end of list/text
       q: Quit
```

