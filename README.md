# gupdate

> sync github repositories from different users and teams

<!-- TOC -->
* [gupdate](#gupdate)
  * [About The Project](#about-the-project)
    * [Built With](#built-with)
  * [Getting Started](#getting-started)
    * [Example](#example)
    * [Installation](#installation)
      * [use release](#use-release)
      * [Compile from source](#compile-from-source)
        * [Prerequisites](#prerequisites)
        * [create binary](#create-binary)
  * [Usage](#usage)
    * [Configuration](#configuration)
      * [Authentication](#authentication)
  * [Roadmap](#roadmap)
  * [Contributing](#contributing)
  * [License](#license)
  * [Contact](#contact)
<!-- TOC -->


## About The Project

### Built With

* [cobra](https://github.com/spf13/cobra)
* [yaml.v2](https://gopkg.in/yaml.v2)

## Getting Started

Follow the installation steps, add the configuration-file and you're ready to go.

### Example

```bash
$ gupdate
Updating user1, user2, team1 and team2...

Cloned 1 new repositories:
	new_repo:	/Users/user1/team2/new_repo
Local copies of archived repositories:
	archived_repo:	/Users/user1/team1/archived_repo
Errors happened in these repositories:
	not_git_repo:	/Users/user1/team1/not_git_repo - exit status 1
Pulled 132 repositories.
```

### Installation

#### use release
```bash
bash -c "$(curl -fsSL https://raw.githubusercontent.com/sejoharp/gupdate/refs/heads/main/install.sh)"
```

#### Compile from source
##### Prerequisites

In order to use `gupdate`, you need:

* git installed
* github access token (used for pulling/cloning)

##### create binary
1. Clone the repo
   ```bash
   git clone https://github.com/fr3dch3n/gupdate.git
   ```
2. Build the binary for your OS
   ```bash
   make build-linux-amd64
   make build-darwin-amd64
   make build-windows-amd64
   ```

## Usage

Put your configuration file in $HOME/gupdate.yaml or provide a specific path via `gupdate -c <path>`.

### Configuration

The configuration has to follow this format:
```yaml
authentication:
  token: token                              # github access token
  username: my_username                     # github_username for authentication
me:
  username: me                              # your github username (probably the same as authentication.username)
  dir: /home/user/repositories/me           # repos location
  clone_archived: false                     # should archived repos be cloned?
users:
  - username: username                      # specific user to clone
    dir: /home/user/repositories/username   # repos location
    clone_archived: false                   # should archived repos be cloned?
teams:
  - teamname: teamname                      # teamname to clone repos from
    prefix: remove_prefix                   # remove prefix from repos (e.g. team1_repo)
    dir: /home/user/repositories/teamname   # repos location
    org: organization                       # org for the team
    clone_archived: false                   # should archived repos be cloned?
```

#### Authentication

Authentication is done via a github-access-token. This token can be provided in three differeny ways:
* token: <plain text token>
* token_file: <path to a file containing the token>
* token_env_variable: <env variable containing the token>

## Roadmap

See the [open issues](https://github.com/fr3dch3n/gupdate/issues) for a list of proposed features (and
known issues).

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any
contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Twitter: [@fr3dch3n](https://twitter.com/fr3dch3n)
