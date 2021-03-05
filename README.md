# dnsc

## Introduction
This is the tool that I make when I was scanning the list of domains with nmap. Let's say you have a list of domains and  you want to know the IP address of each domain. Then you can use dnsc for that. It was written in go and it is super fast.

## Usage 

For the single domain
```
$ echo google.com | dnsc
```
For multiple domains 
```
$ cat domain.txt | dnsc
```
