#!/usr/local/bin/python

import sys, os

areYouSure1 = raw_input("Are you sure you want to complete this action? This will seriously alter this repository (Y/n): ")
if areYouSure1 != "Y":
    sys.exit(1)

newRepoName = raw_input("Please enter your new repository name: ")
if newRepoName == "":
    print("Can't be empty.")
    sys.exit(1)

newGithubOriginURL = raw_input("Please enter your new GitHub Repository origin URL (should look something like: \"https://github.com/Javin-Ambridge/zoom.eat.git\"): ")
if newGithubOriginURL == "":
    print("Can't be empty.")
    sys.exit(1)
elif not newGithubOriginURL.startswith("https://github.com/"):
    print("Needs to start with \"https://github.com/\".")
    sys.exit(1)
elif not newGithubOriginURL.endswith(".git"):
    print("Needs to end with \".git\".")
    sys.exit(1)

newReadme = raw_input("Please enter a small description for your Readme: ")

print("\nNew Settings. Please Confirm.")
print("Repository Name: " + newRepoName)
print("GitHub Origin URL: " + newGithubOriginURL)
print("Readme Description: " + newReadme)
acceptSettings = raw_input("\nDo these settings look correct to you? (Y/n): ")

if acceptSettings != "Y":
    sys.exit(1)

print("Converting Repository!")
print("Renaming all instances of 'go.base' with \'" + newRepoName + "\' in .go files")
os.system("find . -type f -name \"*.go\" -print0 | xargs -0 sed -i '' -e 's/go\.base/" + newRepoName + "/g'")
print("Renaming all instances of 'go.base' with \'" + newRepoName + "\' in Makefile")
os.system("find . -type f -name \"Makefile\" -print0 | xargs -0 sed -i '' -e 's/go\.base/" + newRepoName + "/g'")
print("Updating README.md")
os.system("echo '# " + newRepoName + " \n\n" + newReadme + "' > README.md")
print("Setting new github origin URL")
os.system("git remote set-url origin " + newGithubOriginURL)
print("Renaming directory")
os.system("cd ..")
os.system("mv go.base/ " + newRepoName + "/")

print("Done. Please add (git add -A), commit (git commit -m ...), and push (git push) to your new repository now.")
