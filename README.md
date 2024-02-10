# Blog-Post 

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>



### Built With

These framework/Libraries and tools are require to build My-Favorite-Artist Service

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)




<!-- GETTING STARTED -->
## Getting Started

Given below prerequisite require to install yours system.

### Application 
There is Blog Post application 


### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Golang
  ```sh
  Linux:
  
  1. $ sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz
  2. export PATH=$PATH:/usr/local/go/bin
  
  MacOs:
  brew install go@1.21
   
  ```
* Windows
  [Golang Installation](https://go.dev/doc/install)

### Installation


4. Run Client
 There is three environment 
 1. Manual
    1. Local Environment
    ```
    go run blogPostClient/client.go -config config/local
      ```
    2. Non prod Environment

    ```
     go run blogPostClient/client.go -config config/non-prod
    ```
    3. Prod Environment

    ```
    go run blogPostClient/client.go -config config/prod
    ```

5. Run Server
 There is three environment 
 1. Manual

    1. Local Environment
    ```
    go run blogPostServer/server.go -config config/local
    ```

2. Non prod Environment

    ```
    go run blogPostServer/server.go -config config/non-prod
    ```
    3. Prod Environment

    ```
    go run blogPostServer/server.go -config config/prod
    ```






