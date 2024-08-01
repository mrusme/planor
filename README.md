Planor
------
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

[![Static 
Badge](https://img.shields.io/badge/Join_on_Matrix-green?style=for-the-badge&logo=element&logoColor=%23ffffff&label=Chat&labelColor=%23333&color=%230DBD8B&link=https%3A%2F%2Fmatrix.to%2F%23%2F%2521PHlbgZTdrhjkCJrfVY%253Amatrix.org)](https://matrix.to/#/%21PHlbgZTdrhjkCJrfVY%3Amatrix.org)

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

### OpenBSD

Planor is available through the `-current` and `-stable` ports on OpenBSD >=
7.2, as well as via `pkg_add planor` on `-current` and `-stable` > 7.2. 
The package is being maintained by [gonzalo-](https://github.com/gonzalo-)
(see [this issue](https://github.com/mrusme/planor/issues/2)).

For more info on using the anoncvs ports checkout, see
[here](https://www.openbsd.org/anoncvs.html#updating) and
[here](https://www.openbsd.org/faq/ports/guide.html).


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

