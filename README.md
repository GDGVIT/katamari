<p align="center">
<a href="https://dscvit.com">
	<img src="https://user-images.githubusercontent.com/30529572/92081025-fabe6f00-edb1-11ea-9169-4a8a61a5dd45.png" alt="DSC VIT"/>
</a>
	<h1 align="center">katamari</h1>
	<h4 align="center">A projects page generator for GitHub users and organizations<h4>
</p>

---
[![DOCS](https://img.shields.io/badge/Documentation-see%20docs-green?style=flat-square&logo=appveyor)](https://pkg.go.dev/github.com/GDGVIT/katamari) 
[![Join Us](https://img.shields.io/badge/Join%20Us-Developer%20Student%20Clubs-red)](https://dsc.community.dev/vellore-institute-of-technology/)

# Functionality
- [x] Aggregate all your project READMEs into a single static site
- [x] Support for aggregating User READMEs also
- [x] Generate the static site using Hugo
- [x] Default theme set to "Ananke" (needs git)
- [x] Store GitHub Access Token in config
- [x] Available on the AUR
- [ ] Set Hugo configuration without having to edit `config.toml`
<br>


# Getting started
Follow the instructions to get started and run this project.

## Arch Linux
This package is available as an [AUR](https://aur.archlinux.org/packages/katamari/).

## Dependencies
-  [Go](https://golang.org/doc/install)
-  [Hugo](https://gohugo.io/)

## Installation
```bash
go get -u github.com/GDGVIT/katamari
```

Alternatively, you can also check out [releases](https://github.com/GDGVIT/katamari/releases/) to get a pre-compiled binary

## Generating a katamari project

```bash
katamari create <github organization name>
```
or
```bash
katamari create -u <your github username>
```

## Building the katamari project
```bash
cd <github organization name>
katamari build
```

## Fix rate limiting

Wherever you install katamari, it creates a `config.json` file in `.katamari` folder in the same directory as the executable. You need to add your GitHub Personal Access Token in `.katamari/config.json` in case you're being rate limited. If you require help in creating a Personal Access Token visit [this link](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token).


# Contributors

<table>
<tr align="center">


<td>

Amogh Lele

<p align="center">
<img src = "https://avatars3.githubusercontent.com/u/31761843" width="150" height="150" alt="Amogh Lele">
</p>
<p align="center">
<a href = "https://github.com/sphericalkat"><img src = "http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg" width="36" height = "36" alt="GitHub"/></a>
<a href = "https://www.linkedin.com/in/person1">
<img src = "http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg" width="36" height="36" alt="LinkedIn"/>
</a>
</p>
</td>

<td>

Siddhartha Varma

<p align="center">
<img src = "https://avatars0.githubusercontent.com/u/39856034" width="150" height="150" alt="Amogh Lele">
</p>
<p align="center">
<a href = "https://github.com/BRO3886"><img src = "http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg" width="36" height = "36" alt="GitHub"/></a>
<a href = "https://www.linkedin.com/in/siddharthav22/">
<img src = "http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg" width="36" height="36" alt="LinkedIn"/>
</a>
</p>
</td>

</tr>
  </table>

<p align="center">
	Made with :heart: by <a href="https://dscvit.com">DSC VIT</a>
</p>
