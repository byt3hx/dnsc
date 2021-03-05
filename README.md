# dnsc
[![GitHub issues](https://img.shields.io/github/issues/channyein1337/dnsc)](https://github.com/channyein1337/dnsc/issues)
[![GitHub forks](https://img.shields.io/github/forks/channyein1337/dnsc)](https://github.com/channyein1337/dnsc/network)
[![GitHub stars](https://img.shields.io/github/stars/channyein1337/dnsc)](https://github.com/channyein1337/dnsc/stargazers)
[![GitHub license](https://img.shields.io/github/license/channyein1337/dnsc)](https://github.com/channyein1337/dnsc)

## Introduction
This is the tool that I make when I was scanning the list of domains with nmap. Let's say you have a list of domains and  you want to know the IP address of each domain. Then you can use dnsc for that. It was written in go and it is super fast.

## Usage 

For the single domain
```
$ echo google.com | dnsc
```
![](https://raw.githubusercontent.com/channyein1337/dnsc/main/img/usage%20for%20single%20domain.png)

For multiple domains 
```
$ cat domain.txt | dnsc
```
![](https://raw.githubusercontent.com/channyein1337/dnsc/main/img/usage%20for%20multiple%20domains.png)

## Install:

```
go get github.com/channyein1337/dnsc
```
