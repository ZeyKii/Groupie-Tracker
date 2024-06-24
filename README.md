# SUMMARY
1. 🔍 [Introduction](#introduction)
2. 📋 [Requirements](#requirements)
3. 🚀 [How does the website work](#how-does-the-website-work)
4. 💡 [What features are available to you](#what-features-are-available-to-you)
5. 🛠️ [If you need to edit anything](#if-you-need-to-edit-anything)

<br>

# Introduction
* Name: Groupie Tracker
* Contributors: Ferran Maxance, Lopez Doniban, Bourdoncle Alan

The Groupie Tracker project aims to develop a website in Golang for the backend server part and with the Tailwind framework for the frontend part. This website brings together several artists with precise information on each of them.

<br>

# Requirements

* Golang
* NodeJs (LTS) 
* AZERTY or QWERTY Keyboard 


<br>

# How does the website work
1. Clone this repo `git clone [url]`

2. Go to the repo you just cloned using your terminal: `cd .\groupie-tracker\`

3. Initialize the Go module with `go mod init groupie-tracker`

4. You must go to the SERVER file: `cd .\SERVER\`

5. To launch the website you need to execute server.go file: `go run .\server.go`

6. Open a page on your browser and put in the search bar: `localhost:8081`

7. Let's Go explore! 🚀

<br>

# If you need to edit anything
* If you want to clone the project and edit any css content you will need to use the tailwind scanning command inside the static folder: `npx tailwindcss -i ./src/style.css -o ./dist/output.css --watch`

* If you face an error just type `npm i` inside your static folder to install everything the site needs to work properly, then do the above command again 🔧


<br>

# What features are available to you
The website includes several features, each more useful than the other for you.

- A functional search bar that include filters 🔍
- Clickable cards for each artist with a redirection that returns more informations 📇
- You can easily edit the site using the Tailwind Framework ✏️

<br>

## Here you know everything. ENJOY! 😎
