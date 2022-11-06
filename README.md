<h1 align="center">
    <p>Quality Assurance Platform</p>
    <img alt="Quality Assurance Platform" src="./image/qap_8_bit.png">
</h1>

<p align="center">
    <img src="https://img.shields.io/static/v1?label=go&message=1.18&color=blue&logo=go">
    <img src="https://img.shields.io/badge/version-1.0.0-lightgrey">
    <img src="https://img.shields.io/badge/tests-passed-brightgreen">
</p>

## Description
Quality Assurance Platform is a free API's that was developed to help IT roles to learn, test and/or use API's to make a tests for new positions.
Was simulated a car showrooms, following the image:

![Screenshot from 2022-11-06 13-43-21](https://user-images.githubusercontent.com/5350132/200174287-31fcfd54-dc43-44ef-8766-de7ff3c506ec.png)


## Status
<b> v1 - FINISHED </b> -> Simple Dealership API's with Authentication and Authorization

<b> v2 - BACKLOG </b> -> Implements WebSocket server

<b> v3 -  BACKLOG </b> -> Implements gRPC

## Functionalities
To look the functionalities, please access the <a href="https://qap-ws.herokuapp.com/swagger/index.html">QAP SWAGGER</a>

## How to work
- <a href="https://qap-ws.herokuapp.com/swagger/index.html#/users/post_user">Create user</a> with your email
  - If your user was created correctly, you receive a **token on your email**
- Call api`s using you token on header

## Technologies
- Golang 1.19
  - Gin
  - Swaggo
- Postgre Sql
- Heroku
- SendGrid

## Owners
<table>
    <tr>
        <td align=center>
            <img src="https://avatars.githubusercontent.com/u/5350132?v=4" width=115><br><a href='https://github.com/brenos'>Breno Souza</a>
        </td>
        <td align=center>
            <img src="https://avatars.githubusercontent.com/u/36016665?v=4" width=115><br><a href='https://github.com/luuizeduardo'>Luiz Eduardo</a>
        </td>
    </tr>
</table>

## Donate
[![](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?business=QBE4BAJYF5NRY&no_recurring=0&item_name=Help+us+keep+this+platform+online&currency_code=BRL)

