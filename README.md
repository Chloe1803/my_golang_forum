# My Forum in Golang - 'Le Forum du jardinier'

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Features](#features)
5. [Architecture](#architecture)
6. [License](#license)

## Introduction

The aim of this project was to build a web forum entirely in Go, with no JavaScript except for Bootstrap's accordion animation.
Key features include user authentication, post/comment functionalities, and likes/dislikes, all backed by SQLite for data storage. The focus is on web basics, SQL, encryption, and database manipulation. No frontend frameworks were used.
For personalization, I chose the theme of gardening for the forum, and the content is written in French.

## Installation

1. Installation of SQLite3:

Ensure SQLite3 is installed on your system.
On Ubuntu/Linux:

```bash
sudo apt-get install sqlite3
```

On macOS (using Homebrew):

```bash
brew install sqlite3
```

On Windows:
Download and install SQLite3 from SQLite Download Page.

2. Clone the Source Code:

```bash
git clone https://github.com/your-username/your-project.git
```

3. Run the Application:

Open a terminal, navigate to the root of the project, and execute the following command to launch the application:

```bash
go run cmd/main.go
```

4. Access the Application:

Open your browser and navigate to https://localhost:8080.

## Usage

1. User Registration:

Navigate to the registration page and provide a valid email, username, and password.

2. User Login:

Log in using your registered email and password.
Ensure the login session is established.

3. Create a Post:

After logging in, navigate to the "Create Post" section.
Write the content of your post and optionally associate it with one or more categories, or add an image.
Submit the post.

4. Add Comments:

View existing posts.
Add comments to posts by navigating to the respective post's page.
Submit your comment.

5. Like and Dislike Posts/Comments:

As a registered user, interact with the like and dislike buttons on posts and comments.
Observe the visible count of likes and dislikes.

6. Filtering Posts:

Utilize the filtering mechanism to view posts based on categories, created posts, and liked posts.

7. Logout:

Log out of the application when done.

## Features

1. SQLite Database:
   Utilizes SQLite for efficient and reliable data storage.

2. HTTPS Protocol:
   Implements secure communication through the HTTPS protocol.

3. Image Uploads in Posts:
   Enables users to upload and showcase images within their posts.

4. Rate Limiting:
   Implements rate limiting to control and manage client requests.

5. Encrypted Passwords:
   Enhances security by encrypting and securely storing user passwords.

6. Unique Client Sessions and Cookies:
   Ensures secure and unique client sessions using cookies with expiration dates.

7. Dockerfile:
   Provides a Dockerfile for easy deployment and isolation of the forum application.

8. Authentication with Google and GitHub:
   Offers user authentication options through Google and GitHub accounts.

These features collectively contribute to a robust and secure forum application, providing users with a seamless and protected experience.

## Architecture

The file structure of the application is organized for clarity and modularity. Here's a brief overview of the main directories:

1. 'cmd':

• 'main.go': Loads CSS, images, configures the TLS server, and launches the application.

2. 'front':

• 'scripts': Bootstrap JS files for accordion functionality.
• 'static': Contains CSS files and images.
• 'tmpl': HTML files.

3. 'back':

• 'create_db': Initializes and populates the database if it does not exist.

• 'database': Contains the SQLite database.

• 'func':
•'queries': Files with SQL queries.
• Other files containing functions called by handlers for cookie management, email verification, handling likes/dislikes, login, and rate limiting.

• 'server': Contains handlers for each page/request.

This structure promotes a clear separation of concerns, with frontend resources in front, backend logic in back, and the main application entry point in cmd. The func directory encapsulates various functions essential for handling different aspects of the application. Each component is organized to enhance maintainability and ease of understanding.

## License

MIT License

Copyright (c) 2024 Chloé Masse
