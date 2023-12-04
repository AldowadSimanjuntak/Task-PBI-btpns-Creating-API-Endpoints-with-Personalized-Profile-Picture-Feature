# Task BTPN - User Profile Image Upload API


- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [API Endpoints](#api-endpoints)
  - [Custom Configuration](#custom-configuration)

## Introduction

This Task was my final submission during my internship at BTPN . It aims to develop an API that allows users to upload and manage their profile images. This API is designed to enhance user engagement in a mobile banking application by offering a personalized user experience. Users can upload their profile photos to make the app feel more tailored to them.

## Features

- **User Registration:** Users can register and create their accounts.

- **User Login:** Registered users can log in to their accounts.

- **User Profile Image Upload:** Users can upload and manage their profile images.

- **User Authentication:** User actions are protected and require authentication.

## Getting Started

To get started with this project, follow the instructions below.

### Prerequisites

1. **Go Programming Language:** Ensure you have Go installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

2. **PostgreSQL Database:** You should have a PostgreSQL database up and running. Make sure to create a database and configure the URL in `config/config.go`.

3. **Secret Key:** Generate a secret key and update it in `config/config.go`.

### Installation

1. **Clone this repository:**

   ```
   git clone https://github.com/AldowadSimanjuntak/Task-PBI-btpns-Creating-API-Endpoints-with-Personalized-Profile-Picture-Feature.git
### Change to the project directory:
    cd Task-PBI-btpns-Creating-API-Endpoints-with-Personalized-Profile-Picture-Feature

### Install project dependencies:
Use [Go modules](https://golang.org/ref/mod) to install project dependencies.

### Run the project:
    go run app/main.go

The API will be accessible at http://localhost:8080.

## Usage

### API Endpoints

- **POST /users/register:** Register a new user.

- **POST /users/login:** Log in with an existing user.

- **POST /api/photos:** Upload a user profile image.

- **GET /api/photos:** Retrieve a list of user profile images.

- **DELETE /api/photos/:photoId:** Delete a user profile image.

## Custom Configuration

You can customize the configuration by updating the `config/config.go` file with your PostgreSQL database URL and secret key.
